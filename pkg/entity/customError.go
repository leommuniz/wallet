package entity

type CustomError struct {
	Message string `json:"message"`
}

func NewCustomError(m string) *CustomError {
	err := &CustomError{Message: m}
	return err
}
