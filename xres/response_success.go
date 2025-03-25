package xres

func Success(i interface{}) (res Response) {
	return Response{
		Code:    "success",
		Message: "Success",
		Data:    i,
	}
}

func SuccessCached(i interface{}, c bool) (res ResponseCached) {
	res = ResponseCached{
		IsCached: c,
	}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i

	return res
}

func SuccessPaging(i interface{}, p Paging) (res ResponsePaging) {
	res = ResponsePaging{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i
	res.Page = p.Page
	res.Limit = p.Limit
	res.Total = p.Total

	return res
}
