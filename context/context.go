package context

import "net/http"

// Context Http request context struct including BeegoInpout, BeegoOutput, http.Request and http.ResponseWriter.
// BeegoInput and BeegoOutput provides some api to operate request and response more easily.
type Context struct {
	Input   *BeegoInput
	Output  *BeegoOutput
	Request	*http.Request
	ResponseWriter *Response
	_xsrfToken string
}