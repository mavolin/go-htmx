package htmx

import "encoding/json"

// aliases for easier understanding of fields
type (
	Element       = string
	ID            = string
	Selector      = string
	Event         = string
	URL           = string
	SameOriginURL = string
	JS            = string

	JSON = json.RawMessage
)

type Headers map[string]string

type SwapStrategy string

const (
	SwapInnerHTML   SwapStrategy = "innerHTML"
	SwapOuterHTML   SwapStrategy = "outerHTML"
	SwapBeforeBegin SwapStrategy = "beforebegin"
	SwapAfterBegin  SwapStrategy = "afterbegin"
	SwapBeforeEnd   SwapStrategy = "beforeend"
	SwapAfterEnd    SwapStrategy = "afterend"
	SwapDelete      SwapStrategy = "delete"
	SwapNone        SwapStrategy = "none"
)
