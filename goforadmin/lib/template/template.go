package template

import (
	"bytes"
	"html/template"
	"sync"
)

type Vars map[string]interface{}

type tplData struct {
	templates map[string]*template.Template
	sync.RWMutex
}

var (
	tplCache   tplData
	globalVars Vars
)

func init() {
	tplCache.templates = make(map[string]*template.Template)
	globalVars = make(Vars)
}

func RegisterGlobalVars(vars Vars) {
	for k, v := range vars {
		globalVars[k] = v
	}
}

func ParseFile(file string) (*template.Template, error) {
	tplCache.RLock()
	t, ok := tplCache.templates[file]
	tplCache.RUnlock()
	if ok {
		// loading template every time for development mode
		if _, ok := globalVars["EnableCache"]; ok {
			return t, nil
		}
	}

	t, err := template.ParseFiles(file)
	if err != nil {
		return nil, err
	}

	tplCache.Lock()
	tplCache.templates[file] = t
	tplCache.Unlock()

	return t, nil
}

func ParseFileToString(file string, data interface{}) (string, error) {
	t, err := ParseFile(file)
	if err != nil {
		return "", err
	}

	if data == nil {
		data = make(Vars)
	}

	if m, ok := data.(Vars); ok {
		if len(globalVars) > 0 {
			for k, v := range globalVars {
				m[k] = v
			}
		}
		if layout, ok := m["layout"]; ok {
			t.ParseFiles(layout.(string))
		}
		data = m
	}

	var buffer bytes.Buffer
	if err := t.Execute(&buffer, data); err != nil {
		return "", err
	}

	return buffer.String(), nil
}
