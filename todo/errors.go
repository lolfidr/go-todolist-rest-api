package todo

import "errors"

type ErrDTO struct {
	Message string
	Code    int
}

var ErrTaskNotFound = errors.New("task not found")
var ErrTaskAlreadyExists = errors.New("task already exists")
