package beego

import "net/http"

var ErrorMaps = make(map[string]*errorInfo, 10)

// show 401 unauthorized error.
func unauthorized(rw http.ResponseWriter, r *http.Request) {

}

// show 402 Payment Required
func paymentRequired(rw http.ResponseWriter, r *http.Request) {

}

// show 403 forbidden error
func forbidden(rw http.ResponseWriter, r *http.Request) {

}

// show 422 missing xsrf token
func missingxsrf(rw http.ResponseWriter, r *http.Request) {

}

// show 417 invalid xsrf token
func invalidxsrf(rw http.ResponseWriter, r *http.Request) {

}

// show 404 not found error
func notFound(rw http.ResponseWriter, r *http.Request) {

}

// show 405 Method not Allowed
func methodNotAllowed(rw http.ResponseWriter, r *http.Request) {

}

// show 500 internal server error
func internalServerError(rw http.ResponseWriter, r *http.Request) {

}

// show 501 Not Implemented
func notImplemented(rw http.ResponseWriter, r *http.Request) {

}

// show 502 Bad Gateway
func badGateway(rw http.ResponseWriter, r *http.Request) {

}

// show 503 service unavailable error
func serviceUnavailable(rw http.ResponseWriter, r *http.Request) {

}

// show 504 Gateway
func gatewayTimeout(rw http.ResponseWriter, r *http.Request) {

}

func ErrorHandler(code string, h http.HandlerFunc) *App {

}