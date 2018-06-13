package main

import (
	"fmt"
	"net/http"
	"strconv"
)

// Render write response to http connection
func Render(w http.ResponseWriter, r *http.Request, body []byte, code int) {
	w.WriteHeader(code)
	w.Write(body)
}

// ParseFileSize convert file size string to integer
func ParseFileSize(s string) (int64, error) {
	p := -1
	for i, v := range s {
		if v < '0' || v > '9' {
			p = i
			break
		}
	}

	if p == -1 {
		return strconv.ParseInt(s, 10, 64)
	}

	n, err := strconv.ParseInt(s[0:p], 10, 64)
	if err != nil {
		return 0, err
	}

	switch s[p] {
	case 'k', 'K':
		return n * 1024, nil
	case 'm', 'M':
		return n * 1024 * 1024, nil
	case 'g', 'G':
		return n * 1024 * 1024 * 1024, nil
	default:
		return 0, fmt.Errorf("error format, should be 1k/1kb/1KB")
	}
}
