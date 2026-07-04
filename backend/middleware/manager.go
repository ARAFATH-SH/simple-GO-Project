package middleware

import (
	"net/http"
)

type Middleware func(next http.Handler) http.Handler

type Manager struct {
	globarMiddlewares []Middleware
}

func NewManager() *Manager {
	return &Manager{
		globarMiddlewares: make([]Middleware, 0),
	}
}

func (mngr *Manager) Use(middlewares ...Middleware) {
	mngr.globarMiddlewares = append(mngr.globarMiddlewares, middlewares...)
}

func (mngr *Manager) With(next http.Handler, middlewares ...Middleware) http.Handler {
	n := next

	// logger(http.HandleFunc(GetProducts))
	for _, middleware := range middlewares {
		n = middleware(n)
	}

	//m.globalMiddlewares = [logger, hudai]
	// for _, globalMiddlewares := range mngr.globarMiddlewares {
	// 	n = globalMiddlewares(n)
	// }

	return n
}

func (mngr *Manager) WrapMux(mux http.Handler) http.Handler {
	n := mux

	// logger(http.HandleFunc(GetProducts))
	for _, middleware := range mngr.globarMiddlewares {
		n = middleware(n)
	}

	return n
}
