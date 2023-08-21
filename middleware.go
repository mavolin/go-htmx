package htmx

import (
	"context"
	"net/http"
)

type ctxKey struct{}

type responseWriterWrapper struct {
	http.ResponseWriter
	h            *ResponseHeaders
	wroteHeaders bool
	errHandler   func(error)
}

func (r *responseWriterWrapper) Write(data []byte) (int, error) {
	if !r.wroteHeaders {
		r.h.AddHeaders(r.ResponseWriter.Header())
		r.wroteHeaders = true
	}

	return r.ResponseWriter.Write(data)
}

// NewMiddleware returns a new middleware that adds htmx headers, set by
// handlers called after this middleware, to the response.
func NewMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			var h ResponseHeaders
			*r = *r.WithContext(context.WithValue(r.Context(), ctxKey{}, &h))

			ww := &responseWriterWrapper{ResponseWriter: w, h: &h}
			next.ServeHTTP(ww, r)
		})
	}
}

// Response returns a pointer to the response headers that will be sent back.
//
// It must be called after the middleware has executed.
func Response(r *http.Request) *ResponseHeaders {
	return r.Context().Value(ctxKey{}).(*ResponseHeaders)
}
