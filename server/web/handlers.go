package web

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-http-utils/headers"
)

//Handler handler function builder for controller
type Handler struct {
	HandleFunc func(*Context, http.ResponseWriter, *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	c := &Context{}
	fmt.Println(now)
	w.Header().Add(headers.ContentType, "application/json")
	//build context here

	//call Handler function
	h.HandleFunc(c, w, r)
}
