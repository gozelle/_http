package _http

type Method string

const (
	POST   Method = "POST"
	GET    Method = "GET"
	PUT    Method = "PUT"
	DELETE Method = "DELETE"
	HEAD   Method = "HEAD"
)
const (
	CONTENT_TYPE                       = "Content-Type"
	CONTENT_TYPE_JSON                  = "application/json"
	CONTENT_TYPE_JSON_UTF8             = "application/json; charset=utf-8"
	CONTENT_TYPE_X_WWW_FORM_URLENCODED = "application/x-www-form-urlencoded"
	CONTENT_TYPE_FORM_DATA             = "multipart/form-data"
	CONTENT_TYPE_TEXT_XML              = "text/xml"
)
