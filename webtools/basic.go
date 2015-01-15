package webtools

import (
	"net/http"
)

/*
 * 判断是否为Ajax请求
 */
func IsAjax(r *http.Request) bool {
	result := false
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		return true
	}
	return result
}
