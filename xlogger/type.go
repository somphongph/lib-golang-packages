package xlogger

type LogStruct struct {
	Timestamp   string         `json:"timestamp"`
	Level       AppLogLevel    `json:"level"`
	ServiceName string         `json:"service_name"`
	MethodName  string         `json:"method_name"`
	Subject     string         `json:"subject"`
	Message     string         `json:"message"`
	Environment string         `json:"environment,omitempty"`
	RequestId   string         `json:"request_id,omitempty"` // in case of microservices
	TraceId     string         `json:"trace_id,omitempty"`   // distributed tracing
	SpanId      string         `json:"span_id,omitempty"`    // span of distributed tracing
	Additional  map[string]any `json:"additional,omitempty"`
	LogRequest  *LogRequest    `json:"log_request"`
}

type LogRequest struct {
	HttpMethod    HttpMethod `json:"http_method,omitempty"`
	Endpoint      string     `json:"endpoint,omitempty"`
	Request       string     `json:"request,omitempty"`
	StatusCode    string     `json:"status_code,omitempty"`
	Response      string     `json:"response,omitempty"`
	ExecutionTime int64      `json:"execution_time,omitempty"`
}
