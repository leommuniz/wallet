package entity

type CustomError struct {
	Message string `json:"message"`
}

func (e *CustomError) Error(m string) *CustomError {
	err := &CustomError{Message: m}
	return err
}
