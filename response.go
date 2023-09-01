package htmx

import (
	"encoding/json"
	"net/http"
	"strings"
)

type (
	// ResponseHeaders are the headers that can be sent in response to an htmx
	// request.
	//
	// The docs of the fields are copied from the website.
	//
	// See: https://htmx.org/reference#response_headers
	ResponseHeaders struct {
		// Location allows you to do a client-side redirect that does not do a
		// full page reload.
		//
		// This response header can be used to trigger a client side
		// redirection without reloading the whole page.
		// Instead of changing the page’s location it will act like following a
		// hx-boost link, creating a new history entry, issuing an ajax request
		// to the value of the header and pushing the path into history.
		//
		// A sample response would be:
		//
		//	HX-Location: /test
		//
		// Which would push the client to test as if the user had clicked on
		// <a href="/test" hx-boost="true">.
		//
		// If you want to redirect to a specific target on the page rather than
		// the default of document.body, you can pass more details along with
		// the event, by setting the remaining Location values.
		//
		// If Location.Path is an empty string, the location is not included.
		Location Location
		// PushURL pushes a new url into the history stack:
		//
		// The HX-Push-Url header allows you to push a URL into the browser
		// location history.
		// This creates a new history entry, allowing navigation with the
		// browser’s back and forward buttons.
		// This is similar to the hx-push-url attribute.
		//
		// If present, this header overrides any behavior defined with
		// attributes.
		//
		// To prevent the browser's history from being updated set PushURL to
		// "false"/
		//
		// The URL must be from the same origin as the request.
		PushURL SameOriginURL
		// Redirect can be used to do a client-side redirect to a new location.
		Redirect URL
		// Refresh, if set to “true”, will do a full refresh of the page on the
		// client side.
		Refresh bool
		// ReplaceURL allows you to replace the current URL in the browser
		// location history.
		// This does not create a new history entry; in effect, it removes the
		// previous current URL from the browser’s history.
		// This is similar to the hx-replace-url attribute.
		//
		// If present, this header overrides any behavior defined with
		// attributes.
		//
		// To prevent the browser's current URL from being updated, set
		// ReplaceURL to "false".
		//
		// The URL must be from the same origin as the request.
		ReplaceURL SameOriginURL
		// Reswap allows you to specify how the response will be swapped.
		Reswap SwapStrategy
		// Retarget is a CSS selector that updates the target of the content
		// update to a different element on the page.
		Retarget Selector
		// Reselect is a CSS selector that allows you to choose which part of
		// the response is used to be swapped in.
		//
		// Overrides an existing hx-select on the triggering element.
		Reselect Selector
		// Trigger triggers events as soon as the response is received.
		Trigger map[Event]JSON
		// TriggerAfterSettle triggers events after the settling step.
		TriggerAfterSettle map[Event]JSON
		// TriggerAfterSwap triggers JSON after the swap step.
		TriggerAfterSwap map[Event]JSON
	}

	// Location is a location used as the HX-Location response header.
	//
	// See: https://htmx.org/headers/hx-location
	Location struct {
		// Path is required and is url to load the response from.
		Path URL `json:"path,omitempty"`
		// Source is the source element of the request.
		Source Selector `json:"source,omitempty"`
		// Event is an event that “triggered” the request.
		Event Event `json:"event,omitempty"`
		// Handler is a callback that will handle the response HTML.
		Handler JS `json:"handler,omitempty"`
		// Target is the target to swap the response into.
		Target Selector `json:"target,omitempty"`
		// Swap determines how the response will be swapped in relative to the
		// target.
		Swap SwapStrategy `json:"swap,omitempty"`
		// Values are the values to submit with the request.
		Values JSON `json:"values,omitempty"`
		// Headers are the headers to submit with the request.
		Headers Headers `json:"headers,omitempty"`
	}
)

func (h *ResponseHeaders) AddHeaders(header http.Header) {
	if h.Location.Path != "" {
		header.Add("HX-Location", h.Location.HeaderValue())
	}
	if h.PushURL != "" {
		header.Add("HX-Push-Url", h.PushURL)
	}
	if h.Redirect != "" {
		header.Add("HX-Redirect", h.Redirect)
	}
	if h.Refresh {
		header.Add("HX-Refresh", "true")
	}
	if h.ReplaceURL != "" {
		header.Add("HX-Replace-Url", h.ReplaceURL)
	}
	if h.Reswap != "" {
		header.Add("HX-Reswap", string(h.Reswap))
	}
	if h.Retarget != "" {
		header.Add("HX-Retarget", h.Retarget)
	}
	if h.Reselect != "" {
		header.Add("HX-Reselect", h.Reselect)
	}
	if len(h.Trigger) > 0 {
		header.Add("HX-Trigger", eventTriggersToHeaderValue(h.Trigger))
	}
	if len(h.TriggerAfterSettle) > 0 {
		header.Add("HX-Trigger-After-Settle", eventTriggersToHeaderValue(h.TriggerAfterSettle))
	}
	if len(h.TriggerAfterSwap) > 0 {
		header.Add("HX-Trigger-After-Swap", eventTriggersToHeaderValue(h.TriggerAfterSwap))
	}
}

func (loc *Location) HeaderValue() string {
	if loc.Source == "" && loc.Event == "" && loc.Handler == "" && loc.Target == "" &&
		loc.Swap == "" && loc.Values == nil && len(loc.Headers) == 0 {
		return loc.Path
	}

	val, err := json.Marshal(loc)
	if err != nil {
		panic(err) // this should never happen
	}

	return string(val)
}

func eventTriggersToHeaderValue(ts map[Event]JSON) string {
	var eventLen int

	var hasData bool
	for event, data := range ts {
		eventLen += len(event) + len(",")
		if data != nil {
			hasData = true
		}
	}

	if hasData {
		data, err := json.Marshal(ts)
		if err != nil {
			panic(err) // this should never happen
		}
		return string(data)
	}

	var b strings.Builder
	b.Grow(eventLen - 1) // minus one comma that we don't need
	for event := range ts {
		if b.Len() > 0 {
			b.WriteByte(',')
		}
		b.WriteString(event)
	}
	return b.String()
}
