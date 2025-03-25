package xres

func CannotBindData() (res Response) {
	return Response{
		Code:    "cannot_bind_data",
		Message: "Cannot bind data",
		Data:    nil,
	}
}

func DataNotFound() (res Response) {
	return Response{
		Code:    "data_not_found",
		Message: "Data not found.",
		Data:    nil,
	}
}

func OperationFailed() (res Response) {
	return Response{
		Code:    "operation_failed",
		Message: "The operation failed.",
		Data:    nil,
	}
}
