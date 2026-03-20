package http

import (
	"encoding/json"
	"errors"
	"time"
)

type CompleteTaskDTO struct {
	Complete bool
}

type TaskDTO struct {
	Title       string
	Description string
}

func (t TaskDTO) ValidForCreate() error {
	if t.Title == "" {
		return errors.New("title is empty")
	}

	if t.Description == "" {
		return errors.New("description is empty")
	}

	return nil
}

type ErrDTO struct {
	Message string
	Time    time.Time
}

func NewErrDTO(message error) *ErrDTO {
	errDTO := &ErrDTO{
		Message: message.Error(),
		Time:    time.Now(),
	}

	return errDTO
}

func (e ErrDTO) ToString() string {
	b, err := json.MarshalIndent(e, "", "    ")
	if err != nil {
		panic(err)
	}

	return string(b)
}
