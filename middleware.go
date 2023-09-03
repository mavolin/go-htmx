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
}

func (w *responseWriterWrapper) Write(data []byte) (int, error) {
	w.writeHXHeader()
	return w.ResponseWriter.Write(data)
}

func (w *responseWriterWrapper) WriteHeader(statusCode int) {
	w.writeHXHeader()
	w.ResponseWriter.WriteHeader(statusCode)
}

func (w *responseWriterWrapper) writeHXHeader() {
	if w.wroteHeaders {
		return
	}

	w.h.AddHeaders(w.ResponseWriter.Header())
	w.wroteHeaders = true
}

// NewMiddleware returns a new middleware that adds htmx headers, set by
// handlers called after this middleware, to the response.
func NewMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			h := ResponseHeaders{
				Trigger:            make(map[Event]JSON),
				TriggerAfterSettle: make(map[Event]JSON),
				TriggerAfterSwap:   make(map[Event]JSON),
			}
			*r = *r.WithContext(context.WithValue(r.Context(), ctxKey{}, &h))

			ww := &responseWriterWrapper{ResponseWriter: w, h: &h}
			next.ServeHTTP(ww, r)
			ww.writeHXHeader()
		})
	}
}

// Response returns a pointer to the response headers that will be sent back.
//
// It must be called after the middleware has executed.
func Response(r *http.Request) *ResponseHeaders {
	return r.Context().Value(ctxKey{}).(*ResponseHeaders)
}
