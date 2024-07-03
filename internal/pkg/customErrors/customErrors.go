package customerrors

type CustomError struct {
	Message string
	Code    int
}

func (err CustomError) Error() string {
	return err.Message
}

type InternalServerError struct {
	Message string
}

func (err InternalServerError) Error() string {
	return err.Message
}

type BadRequestError struct {
	Message string
}

func (err BadRequestError) Error() string {
	return err.Message
}

type DuplicateKeyError struct {
	Message string
}

func (err DuplicateKeyError) Error() string {
	return err.Message
}

type InvalidCredentialError struct {
	Message string
}

func (err InvalidCredentialError) Error() string {
	return err.Message
}

type NotFoundError struct {
	Message string
}

func (err NotFoundError) Error() string {
	return err.Message
}
