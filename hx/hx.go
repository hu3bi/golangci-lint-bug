// Package hx provides utilities for handling HTMX requests.
//
// HTMX response code handling: ⇒ TLDR: use 2xx code
//   - 2xx: SWAP — HTMX will process the response and update the DOM.
//   - 204: NO SWAP	— [http.StatusNoContent] by default does nothing, but is not an [error].
//   - 3xx: WEIRD — If the response includes a valid Location header,
//     HTMX will swap content using the response from the redirected URL.
//     The browser will intercept the redirection internally,
//     retrieve the headers and body from the new location, and HTMX swaps in the redirected response.
//   - 4xx, 5xx: NO SWAP — HTMX ignores the response and does not update the DOM. These are treated as errors.
//   - other code: NO SWAP — HTMX ignores the response and does not update the DOM.
package hx

import (
	"net/http"
)

// IsRequest reports if the request is an HTMX request by looking for the Hx-Request header.
//
// In order to reject non-HTMX requests, use [httpe.RejectNonHXRequest].
//
// See [hxRequestHeader].
func IsRequest(r *http.Request) bool {
	return r.Header.Get(hxRequestHeader) == "true"
}


// bug exists only if the constant is in a separate file. enable "go.lintOnSave": "package" in vscode settings.