package xres

type Paging struct {
	Page  int   `json:"page"`
	Limit int   `json:"limit"`
	Total int64 `json:"total"`
}

type Response struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Data    any    `json:"data"`
}

type ResponseCached struct {
	IsCached bool `json:"is_cached"`
	Response
}

type ResponsePaging struct {
	Response
	Paging
}
