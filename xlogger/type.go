package xlogger

type AppLog struct {
	MethodName    string         `json:"method_name"`
	Subject       string         `json:"subject"`
	Message       string         `json:"message"`
	SpanId        string         `json:"span_id,omitempty"` // span of distributed tracing
	Additional    map[string]any `json:"additional,omitempty"`
	AppLogRequest *AppLogRequest `json:"app_log_request"`
}

type AppLogRequest struct {
	HttpMethod    HttpMethod `json:"http_method,omitempty"`
	Endpoint      string     `json:"endpoint,omitempty"`
	Request       string     `json:"request,omitempty"`
	Response      string     `json:"response,omitempty"`
	StatusCode    string     `json:"status_code,omitempty"`
	ExecutionTime int64      `json:"execution_time,omitempty"`
}

type logRecord struct {
	AppLog `json:",inline" bson:",inline"`

	Timestamp   string      `json:"timestamp"`
	Level       AppLogLevel `json:"level"`
	TraceId     string      `json:"trace_id"`
	ServiceName string      `json:"service_name"`
	Environment string      `json:"environment"`
}
