package httpserver

type Error struct {
	Error string `json:"error"`
}

func ErrorResponse(err error) Error {
	error := Error{
		Error: err.Error(),
	}
	return error
}
