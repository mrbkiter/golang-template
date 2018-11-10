package web

import (
	"net/http"

	"template.github.com/server/utils"
)

//Handle404 handler 404
func Handle404(w http.ResponseWriter, r *http.Request) {

}

//ReturnStatusOK Return status 200
func ReturnStatusOK(w http.ResponseWriter) {
	m := make(map[string]string)
	m["status"] = "200"
	w.Write([]byte(utils.MapToJson(m)))
}
