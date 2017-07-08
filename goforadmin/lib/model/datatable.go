package datatable

import (
	"bytes"
	"encoding/json"
	"fmt"
	"reflect"
	"strconv"
	"strings"

	. "github.com/Yaoguais/gadmin/lib/db"
	"github.com/Yaoguais/gadmin/lib/log"
	"github.com/Yaoguais/gadmin/lib/slice"
	"github.com/labstack/echo"
)

var defaultLimits []int

func init() {
	defaultLimits = []int{10, 20, 50}
}

type DtResult struct {
	Draw            int             `json:"draw"`
	RecordsTotal    int             `json:"recordsTotal"`
	RecordsFiltered int             `json:"recordsFiltered"`
	Data            [][]interface{} `json:"data"`
}

type DtConfig struct {
	Table   string              `json:"-"`
	Limits  []int               `json:"limits"`
	Columns []DtColConfig       `json:"columns"`
	RowCB   func([]interface{}) `json:"-"`
}

type DtColConfig struct {
	Index      int    `json:"index"`
	Sortable   bool   `json:"sortable"`
	Searchable bool   `json:"-"`
	Order      string `json:"order"`
	Type       string `json:"type"`
	Name       string `json:"-"`
}

func GetModelDtData(c echo.Context, conf *DtConfig) *DtResult {
	var (
		result DtResult
		orders bytes.Buffer
		sql    string
		err    error
	)

	var size int
	for _, colConf := range conf.Columns {
		if colConf.Index+1 > size {
			size = colConf.Index + 1
		}
	}

	var columns []string = make([]string, size, size)
	var columnMap []int = make([]int, len(conf.Columns), len(conf.Columns))
	var columnMap2 []int = make([]int, size, size)
	var likes []string = make([]string, 0, size)
	var keywords []interface{} = make([]interface{}, 0, size)

	draw, _ := strconv.Atoi(c.FormValue("draw"))
	total, _ := strconv.Atoi(c.FormValue("total"))
	offset, _ := strconv.Atoi(c.FormValue("start"))
	limit, _ := strconv.Atoi(c.FormValue("length"))
	keyword := c.FormValue("search[value]")
	pattern := "%" + keyword + "%"

	if offset < 0 {
		offset = 0
	}
	if !slice.InIntArray(limit, conf.Limits) {
		limit = conf.Limits[0]
	}

	result.Draw = draw
	result.RecordsFiltered = total
	result.RecordsTotal = total

	for i, colConf := range conf.Columns {
		columnMap[i] = colConf.Index
		columnMap2[colConf.Index] = i
		columns[colConf.Index] = colConf.Name
		if len(keyword) > 0 && colConf.Searchable {
			likes = append(likes, fmt.Sprintf("%s like ?", wrapperColumn(colConf.Name)))
			keywords = append(keywords, pattern)
		}
	}

	for i := 0; i < size; i++ {
		ck := fmt.Sprintf("order[%d][column]", i)
		dk := fmt.Sprintf("order[%d][dir]", i)
		cv := c.FormValue(ck)
		dv := c.FormValue(dk)

		if cv == "" {
			break
		}
		if ci, err := strconv.Atoi(cv); err == nil && ci >= 0 && ci < size {
			if dv == "asc" || dv == "desc" {
				c := conf.Columns[columnMap2[ci]]
				if c.Sortable {
					if orders.Len() > 0 {
						orders.WriteString(",")
					}
					orders.WriteString(fmt.Sprintf("%s %s", wrapperColumn(c.Name), dv))
				}
			}
		}
	}

	var column bytes.Buffer
	var writeComma = false
	for _, v := range columns {
		if len(v) > 0 {
			if writeComma {
				column.WriteString(",")
			} else {
				writeComma = true
			}
			column.WriteString(wrapperColumn(v))
		}
	}

	var where string
	if len(likes) > 0 {
		where = "WHERE " + (strings.Join(likes, " OR "))
	}

	var orderBy string
	if orders.Len() > 0 {
		orderBy = "ORDER BY " + orders.String()
	}

	var limits string
	if limit > 0 {
		limits = fmt.Sprintf("LIMIT %d OFFSET %d", limit, offset)
	}

	if total == 0 {
		sqlTotal := fmt.Sprintf("SELECT count(*) FROM %s %s",
			conf.Table,
			where,
		)
		err := Db.Get(&total, sqlTotal, keywords...)
		if err != nil {
			log.Error("datatable get total failed ", sqlTotal, err)
			result.Data = make([][]interface{}, 0, 1)
			return &result
		} else {
			result.RecordsFiltered = total
			result.RecordsTotal = total
		}
	}

	sql = fmt.Sprintf("SELECT %s FROM %s %s %s %s",
		column.String(),
		conf.Table,
		where,
		orderBy,
		limits,
	)

	if limit < 0 {
		result.Data = make([][]interface{}, 0, total)
	} else {
		result.Data = make([][]interface{}, 0, limit)
	}

	rows, err := Db.Queryx(sql, keywords...)
	if err != nil {
		log.Error("datatable get data failed ", sql, err)
		return &result
	}

	for rows.Next() {
		row, err := rows.SliceScan()
		if err != nil {
			log.Error("datatable get data failed ", sql, err)
			break
		} else {
			for i, v := range row {
				if v2, ok := v.([]byte); ok {
					row[i] = string(v2)
				}
			}
			newRow := make([]interface{}, size, size)
			for i, v := range columnMap {
				newRow[v] = row[i]
			}

			if conf.RowCB != nil {
				conf.RowCB(newRow)
			}
			result.Data = append(result.Data, newRow)
		}
	}

	return &result
}

func wrapperColumn(c string) string {
	var r bytes.Buffer
	s := strings.Split(c, ".")
	for _, v := range s {
		if r.Len() > 0 {
			r.WriteString(".")
		}
		r.WriteString("`")
		r.WriteString(v)
		r.WriteString("`")
	}

	return r.String()
}

func NewDtConfig(model interface{}) *DtConfig {
	var conf DtConfig
	v := reflect.ValueOf(model)
	t := reflect.Indirect(v).Type()
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		tag := f.Tag.Get("dt")
		if f.Name == "_" {
			parseDtConfig(tag, &conf)
		} else {
			colConf := parseDtColConfig(tag)

			if colConf.Index >= 0 {
				name := f.Tag.Get("db")
				if name == "" {
					name = strings.ToLower(f.Name)
				}
				colConf.Name = name

				conf.Columns = append(conf.Columns, *colConf)
			}
		}
	}

	if len(conf.Limits) == 0 {
		conf.Limits = defaultLimits
	}

	return &conf
}

func parseDtConfig(str string, config *DtConfig) bool {
	var parsed bool
	paris := strings.Split(str, ";")
	for _, v := range paris {
		kv := strings.Split(v, ":")
		if len(kv) == 2 {
			if kv[0] == "limits" {
				if kv[1] == "" || nil != json.Unmarshal([]byte(kv[1]), &config.Limits) {
					config.Limits = defaultLimits
				}
				parsed = true
			} else if kv[0] == "table" {
				config.Table = kv[1]
			}
		}
	}

	if !parsed {
		return false
	}

	return true
}

func parseDtColConfig(str string) *DtColConfig {
	var config DtColConfig

	config.Index = -1
	config.Sortable = true
	config.Searchable = true

	paris := strings.Split(str, ";")
	for _, v := range paris {
		kv := strings.Split(v, ":")
		if len(kv) == 2 {
			k2 := strings.Trim(kv[0], " ")
			v2 := strings.Trim(kv[1], " ")
			if k2 == "index" {
				config.Index, _ = strconv.Atoi(v2)
			} else if k2 == "sortable" {
				if v2 == "false" || v2 == "0" {
					config.Sortable = false
				}
			} else if k2 == "searchable" {
				if v2 == "false" || v2 == "0" {
					config.Searchable = false
				}
			} else if k2 == "order" {
				if v2 == "asc" || v2 == "desc" {
					config.Order = v2
				}
			} else if k2 == "type" {
				if v2 == "string" || v2 == "image" || v2 == "avatar" {
					config.Type = v2
				} else {
					config.Type = "string"
				}
			}
		}
	}

	return &config
}
