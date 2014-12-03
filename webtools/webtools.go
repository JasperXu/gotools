/*
 *
 */
package webtools

import (
	"net/http"
)

func IsAjax(r *http.Request) bool {
	result := false
	if r.Header.Get("X-Requested-With") == "XMLHttpRequest" {
		return true
	}
	return result
}
