package hx

// Non exhaustive constants of [HTMX request headers].
//
// [HTMX request headers]: https://htmx.org/docs/#request-headers
const (
	// hxRequestHeader is always “true”
	// except on history restore requests if `htmx.config.historyRestoreAsHxRequest’ disabled.
	hxRequestHeader = "Hx-Request"
)
