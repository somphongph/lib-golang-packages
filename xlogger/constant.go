package xlogger

type AppLogLevel string

const (
	AppLogLevelInfo  AppLogLevel = "INFO"
	AppLogLevelWarn  AppLogLevel = "WARN"
	AppLogLevelError AppLogLevel = "ERROR"
)

type HttpMethod string

const (
	HttpMethodGet    HttpMethod = "GET"
	HttpMethodPost   HttpMethod = "POST"
	HttpMethodPut    HttpMethod = "PUT"
	HttpMethodPatch  HttpMethod = "PATCH"
	HttpMethodDelete HttpMethod = "DELETE"
)
