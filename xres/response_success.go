package xres

func Success(i any) (res Response) {
	return Response{
		Code:    "success",
		Message: "Success",
		Data:    i,
	}
}

func SuccessCached(i any, c bool) (res ResponseCached) {
	res = ResponseCached{
		IsCached: c,
	}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i

	return res
}

func SuccessPaging(i any, p Paging) (res ResponsePaging) {
	res = ResponsePaging{}
	res.Code = "success"
	res.Message = "Success"
	res.Data = i
	res.Page = p.Page
	res.Limit = p.Limit
	res.Total = p.Total

	return res
}
