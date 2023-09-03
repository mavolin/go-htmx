package htmx

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// LocationData is a location used as the HX-LocationData response header.
//
// See: https://htmx.org/headers/hx-location
type LocationData struct {
	// Path is required and is url to load the response from.
	Path URL
	// Source is the source element of the request.
	Source Selector
	// Event is an event that “triggered” the request.
	Event Event
	// Handler is a callback that will handle the response HTML.
	Handler JS
	// Target is the target to swap the response into.
	Target Selector
	// Swap determines how the response will be swapped in relative to the
	// target.
	Swap SwapStrategy
	// Values are the values to submit with the request.
	Values any
	// Headers are the headers to submit with the request.
	Headers Headers
}

func (d *LocationData) toHeader() (LocationHeader, error) {
	h := LocationHeader{
		Path:    d.Path,
		Source:  d.Source,
		Event:   d.Event,
		Handler: d.Handler,
		Target:  d.Target,
		Swap:    d.Swap,
		Headers: d.Headers,
	}
	if h.Values != nil {
		values, err := json.Marshal(d.Values)
		if err != nil {
			return h, fmt.Errorf("HX-Location: Values: %w", err)
		}

		h.Values = values
	}

	return h, nil
}

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
//	HX-LocationData: /test
//
// Which would push the client to test as if the user had clicked on
// <a href="/test" hx-boost="true">.
//
// If you want to redirect to a specific target on the page rather than
// the default of document.body, you can pass more details along with
// the event, by setting the remaining Location values.
//
// If LocationData.Path is an empty string, the location is not
// included.
//
// Previous values are overwritten.
func Location(r *http.Request, loc LocationData) error {
	h, err := loc.toHeader()
	if err != nil {
		return err
	}

	Response(r).Location = h
	return nil
}

// LocationPath is a shorthand for Location(r, LocationData{Path: path}).
func LocationPath(r *http.Request, path URL) {
	Response(r).Location = LocationHeader{Path: path}
}

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
// "false".
//
// The URL must be from the same origin as the request.
//
// Previous values are overwritten.
func PushURL(r *http.Request, u SameOriginURL) {
	resp := Response(r)
	resp.PushURL = u
}

// PreventPushURL sets the HX-PushURL Header to "false".
//
// It is equivalent to calling PushURL(r, "false").
//
// Previous values are overwritten.
func PreventPushURL(r *http.Request) {
	resp := Response(r)
	resp.PushURL = "false"
}

// Redirect can be used to do a client-side redirect to a new location.
//
// Previous values are overwritten.
func Redirect(r *http.Request, u URL) {
	Response(r).Redirect = u
}

// Refresh if set to “true”, will do a full refresh of the page on the
// client side.
//
// Previous values are overwritten.
func Refresh(r *http.Request, refresh bool) {
	Response(r).Refresh = refresh
}

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
//
// Previous values are overwritten.
func ReplaceURL(r *http.Request, u SameOriginURL) {
	Response(r).ReplaceURL = u
}

// PreventReplaceURL sets the HX-ReplaceURL Header to "false".
//
// It is equivalent to calling ReplaceURL(r, "false").
//
// Previous values are overwritten.
func PreventReplaceURL(r *http.Request) {
	resp := Response(r)
	resp.ReplaceURL = "false"
}

// Reswap allows you to specify how the response will be swapped.
//
// Previous values are overwritten.
func Reswap(r *http.Request, strategy SwapStrategy) {
	Response(r).Reswap = strategy
}

// Retarget is a CSS selector that updates the target of the content
// update to a different element on the page.
//
// Previous values are overwritten.
func Retarget(r *http.Request, sel Selector) {
	Response(r).Retarget = sel
}

// Reselect is a CSS selector that allows you to choose which part of
// the response is used to be swapped in.
//
// Previous values are overwritten.
func Reselect(r *http.Request, sel Selector) {
	Response(r).Reselect = sel
}

// Trigger triggers the passed event as soon as the response is received.
//
// If a there already is a trigger for that event, it will be overwritten.
//
// An error will only be returned if data can't be marshalled to json.
// It is guaranteed that Trigger will never return an error for nil data.
func Trigger(r *http.Request, name Event, data any) error {
	var jsonData JSON
	if data != nil {
		var err error
		jsonData, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	Response(r).Trigger[name] = jsonData
	return nil
}

// TriggerAfterSettle triggers the passed event after the settling step.
//
// If a there already is an after-settle trigger for that event, it will be
// overwritten.
//
// An error will only be returned if data can't be marshalled to json.
// It is guaranteed that TriggerAfterSettle will never return an error for nil
// data.
func TriggerAfterSettle(r *http.Request, name Event, data any) error {
	var jsonData JSON
	if data != nil {
		var err error
		jsonData, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	Response(r).TriggerAfterSettle[name] = jsonData
	return nil
}

// TriggerAfterSwap triggers the passed event after the swap step.
//
// If a there already is an after-swap trigger for that event, it will be
// overwritten.
//
// An error will only be returned if data can't be marshalled to json.
// It is guaranteed that TriggerAfterSwap will never return an error for nil
// data.
func TriggerAfterSwap(r *http.Request, name Event, data any) error {
	var jsonData JSON
	if data != nil {
		var err error
		jsonData, err = json.Marshal(data)
		if err != nil {
			return err
		}
	}

	Response(r).TriggerAfterSwap[name] = jsonData
	return nil
}
