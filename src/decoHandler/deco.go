package decoHandler

import "net/http"

type DecoratorFunc func(w http.ResponseWriter, r *http.Request, handler http.Handler)

type DecoHandler struct {
	fn DecoratorFunc
	h  http.Handler
}

func (self *DecoHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	self.fn(w, r, self.h)
}

func NewDecoHandler(handler http.Handler, fn DecoratorFunc) http.Handler {
	return &DecoHandler{
		fn: fn,
		h:  handler,
	}
}
