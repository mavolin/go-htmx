<div align="center">
<h1>go-htmx</h1>

[![Go Reference](https://pkg.go.dev/badge/github.com/mavolin/go-htmx.svg)](https://pkg.go.dev/github.com/mavolin/go-htmx)
[![Go Report Card](https://goreportcard.com/badge/github.com/mavolin/corgi)](https://goreportcard.com/report/github.com/mavolin/corgi)
[![License MIT](https://img.shields.io/github/license/mavolin/corgi)](./LICENSE)
</div>

---

## About

go-htmx is a helper library for setting and reading [htmx](https://htmx.org) headers.

## Main Features

* 🏔️ High-Level wrapper around the htmx headers
* ✏️ Overwrite header values by calling setters multiple times
* 📖 Fully documented—so you read the header's docs directly in your IDE
* ✅ Proper JSON handling

## Examples

### 👓️ Reading Request Headers

The headers set by htmx can be retrieved by calling `htmx.Request`.

If `htmx.Request` returns nil `htmx.RequestHeaders`, the request was not made
by htmx.

```go
func retrieveHeaders(r *http.Request, w http.ResponseWriter) {
    fmt.Println("boosted:", htmx.Request(r).Boosted)
    fmt.Println("current url:", htmx.Request(r).CurrentURL)
    // you get the idea...
}
```

### ✏️ Setting Response Headers

To add response headers, you first need to add the htmx middleware.

By using a middleware instead of setting headers directly
you can overwrite response headers at a later point in your code.
This is useful if you have a default value for a header that only changes in certain cases.

It also means you can add event triggers one by one and don't have to set them at once.

For [chi](https://github.com/go-chi/chi), adding the middleware could look like this:

```go
r := chi.NewRouter()
r.Use(htmx.NewMiddleware())
```

The middleware will add the headers once the first call to `http.ResponseWriter.Write` is made.

After you've added the middleware, you can start setting headers:

```go
type reloadNavData struct {
	ActiveEntry string
}

func setHeaders(r *http.Request, w http.ResponseWriter) {
    htmx.Retarget(r, "#main")
    htmx.Trigger(r, "reload-nav", reloadNavData{ActiveEntry: "foo"})
    htmx.Trigger(r, "update-cart", nil)
	
    // HX-Retarget: #main 
    // HX-Trigger: {"reload-nav": {ActiveEntry: "foo"}, "update-cart": null}
}
```

You can find the full list of setters on [pkg.go.dev](https://pkg.go.dev/github.com/mavolin/go-htmx).

## License

Built with ❤ by [Maximilian von Lindern](https://github.com/mavolin).
Available under the [MIT License](./LICENSE).
