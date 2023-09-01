package htmx

import "net/http"

// RequestHeaders contains the headers set in an htmx request.
//
// The docs of the fields are copied from the website.
//
// See: https://htmx.org/reference/#request_headers
type RequestHeaders struct {
	// Boosted indicates that the request is via an element using hx-boost.
	Boosted bool
	// CurrentURL is the current URL of the browser.
	CurrentURL URL
	// HistoryRestoreRequest is set to true, if the request is for history
	// restoration after a miss in the local history cache.
	HistoryRestoreRequest bool
	// Prompt is the user response to an hx-prompt.
	Prompt string
	// Target is the id of the target element, if it exists.
	Target ID
	// TriggerName is the name of the triggered element if it exists.
	TriggerName Element
	// Trigger is the id of the triggered element if it exists.
	Trigger ID
}

// Request returns the htmx [RequestHeaders] for the current request.
//
// If the request was not made by htmx (as determined by the lack of the
// "HX-Request" header), Request returns nil.
//
// This function works without the middleware in place.
func Request(r *http.Request) *RequestHeaders {
	if r.Header.Get("HX-Request") != "true" {
		return nil
	}

	return &RequestHeaders{
		Boosted:               r.Header.Get("HX-Boosted") == "true",
		CurrentURL:            r.Header.Get("HX-Current-Url"),
		HistoryRestoreRequest: r.Header.Get("HX-History-Restore-Request") == "true",
		Prompt:                r.Header.Get("HX-Prompt"),
		Target:                r.Header.Get("HX-Target"),
		TriggerName:           r.Header.Get("HX-Trigger-Name"),
		Trigger:               r.Header.Get("HX-Trigger"),
	}
}
