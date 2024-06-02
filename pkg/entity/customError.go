package entity

type CustomError struct {
	Message string `json:"message"`
}

// Error implements error.
func (c *CustomError) Error() string {
	panic("unimplemented")
}

func NewCustomError(m string) *CustomError {
	err := &CustomError{Message: m}
	return err
}
