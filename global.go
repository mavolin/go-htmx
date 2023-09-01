package htmx

import (
	"encoding/json"
	"net/http"
)

// LocationPath sets the [ResponseHeaders.Location.Path], discarding all other
// values set for the location.
func LocationPath(r *http.Request, path URL) {
	Response(r).Location = Location{Path: path}
}

// FullLocation allows sets [ResponseHeaders.Location].
//
// Previous values are overwritten.
func FullLocation(r *http.Request, loc Location) {
	Response(r).Location = loc
}

// PushURL sets [ResponseHeaders.PushURL] to the specified same-origin url.
//
// Previous values are overwritten.
func PushURL(r *http.Request, u SameOriginURL) {
	resp := Response(r)
	resp.PushURL = u
}

// PreventPushURL sets [ResponseHeaders.PushURL] to false.
//
// Previous values are overwritten.
func PreventPushURL(r *http.Request) {
	resp := Response(r)
	resp.PushURL = "false"
}

// Redirect sets [ResponseHeaders.Redirect] to the specified url.
//
// Previous values are overwritten.
func Redirect(r *http.Request, u URL) {
	Response(r).Redirect = u
}

// Refresh sets [ResponseHeaders.Refresh] to the specified value.
//
// Previous values are overwritten.
func Refresh(r *http.Request, refresh bool) {
	Response(r).Refresh = refresh
}

// ReplaceURL sets [ResponseHeaders.ReplaceURL] to the specified same-origin
// url.
//
// Previous values are overwritten.
func ReplaceURL(r *http.Request, u SameOriginURL) {
	Response(r).ReplaceURL = u
}

// PreventReplaceURL sets [ResponseHeaders.ReplaceURL] to false.
//
// Previous values are overwritten.
func PreventReplaceURL(r *http.Request) {
	resp := Response(r)
	resp.ReplaceURL = "false"
}

// Reswap sets [ResponseHeaders.Reswap] to the passed [SwapStrategy].
//
// Previous values are overwritten.
func Reswap(r *http.Request, strategy SwapStrategy) {
	Response(r).Reswap = strategy
}

// Retarget sets [ResponseHeaders.Retarget] to the passed selector.
//
// Previous values are overwritten.
func Retarget(r *http.Request, sel Selector) {
	Response(r).Retarget = sel
}

// Reselect sets [ResponseHeaders.Reselect] to the passed selector.
//
// Previous values are overwritten.
func Reselect(r *http.Request, sel Selector) {
	Response(r).Reselect = sel
}

// Trigger adds a new event trigger.
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

// TriggerAfterSettle adds a new event trigger run after settle.
//
// If a there already is a trigger for that event, it will be overwritten.
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

// TriggerAfterSwap adds a new event trigger run after swap.
//
// If a there already is a trigger for that event, it will be overwritten.
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
