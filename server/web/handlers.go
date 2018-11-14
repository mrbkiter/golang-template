package web

import (
	"context"
	"fmt"
	"net/http"
	"time"

	"template.github.com/server/utils"

	"github.com/go-http-utils/headers"
)

type ContextKey string

const (
	IPAddress ContextKey = "IP_ADDRESS"
	UserConst ContextKey = "USER"
)

//Handler handler function builder for controller
type Handler struct {
	HandleFunc func(*context.Context, http.ResponseWriter, *http.Request)
}

func (h Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	now := time.Now()
	ctx := r.Context()
	fmt.Println(now)
	w.Header().Add(headers.ContentType, "application/json")
	//build context here
	context.WithValue(r.Context(), utils.String("AAA"), utils.String("1234"))
	//call Handler function
	h.HandleFunc(&ctx, w, r)
	//after finish
}
