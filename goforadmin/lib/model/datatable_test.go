package datatable

import (
	"fmt"
	"github.com/Yaoguais/gadmin/lib/db"
	"github.com/Yaoguais/gadmin/lib/json"
	"github.com/Yaoguais/gadmin/lib/random"
	"github.com/labstack/echo"
	"math/rand"
	"net/http"
	"strings"
	"testing"
)

type Model struct {
	_         string
	Uid       string `db:"uid" dt:"index:0;sortable:true;searchable:true;order:asc;type:string"`
	Username  string `dt:"index:1;sortable:true;"`
	Avatar    string `dt:"index:2"`
	CreatedAt int    `db:"created_at" dt:"index:3"`
}

const ModelJson = `{"limits":[10,20,50],"columns":[{"index":0,"sortable":true,"order":"asc","type":"string"},{"index":1,"sortable":true,"order":"","type":""},{"index":2,"sortable":false,"order":"","type":""},{"index":3,"sortable":false,"order":"","type":""}]}`

func TestNewDtConfig(t *testing.T) {
	ret := json.JsonEncode(NewDtConfig(Model{}))
	if ret != ModelJson {
		t.Error("model json not match")
	}
}

// Single Table
type User struct {
	Uid      int64  `db:"id" dt:"index:0;searchable:true;sortable:true;order:asc;type:string"`
	Username string `db:"username" dt:"index:1;searchable:true"`
	Age      int8   `db:"age" dt:"index:2"`
	_        string `dt:"table:test_users;limits:[5, 10, 50]"`
}

// Multi Table with Join
type UserWallet struct {
	Uid      int64  `db:"u.id" dt:"index:0;searchable:true;sortable:true;order:asc;type:string"`
	Username string `db:"u.username" dt:"index:1;searchable:true"`
	Age      int8   `db:"u.age" dt:"index:2"`
	Money    int64  `db:"w.money" dt:"index:3"`
	_        string `dt:"table:test_users u inner join test_user_wallets w on u.id=w.uid;limits:[5, 10, 50]"`
}

var schemaSql = `
CREATE TABLE IF NOT EXISTS test_users (
  id bigint(20) unsigned NOT NULL AUTO_INCREMENT,
  username varchar(32) COLLATE utf8mb4_unicode_ci NOT NULL DEFAULT '',
  age tinyint(3) unsigned NOT NULL DEFAULT '0',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

CREATE TABLE IF NOT EXISTS test_user_wallets (
  uid bigint(20) unsigned NOT NULL,
  money bigint(20) unsigned NOT NULL DEFAULT 0,
  PRIMARY KEY (uid)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;
`

func TestDataTableResult(t *testing.T) {
	dbConf := db.DbConf{
		Driver: "mysql",
		DSN:    "root:@tcp(localhost:3306)/test?charset=utf8mb4",
	}
	db := db.Connect(&dbConf)

	db.Exec(schemaSql)

	for i := 0; i < 1000; i++ {
		db.Exec("INSERT INTO test_users(username, age) VALUES(?,?);",
			random.RandowmString(6),
			18+rand.Intn(20),
		)
	}

	db.Exec("INSERT INTO test_user_wallets(uid) (select id from test_users where id > (select max(uid) from test_user_wallets));")

	models := []interface{}{
		User{},
		UserWallet{},
	}

	for _, model := range models {
		dtConf := NewDtConfig(model)
		e := echo.New()
		req, _ := http.NewRequest("POST", "/", strings.NewReader("search[value]=NE"))
		req.Header.Set(echo.HeaderContentType, echo.MIMEApplicationForm)
		c := e.NewContext(req, nil)
		c.Request().ParseForm()

		result := GetModelDtData(c, dtConf)
		fmt.Println(json.JsonEncode(result))
	}
}
