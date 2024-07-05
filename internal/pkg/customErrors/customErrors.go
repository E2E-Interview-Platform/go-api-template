package customerrors

type Error struct {
	Code          int
	CustomMessage string
	InternalError error
}

func (err Error) Error() string {
	return err.InternalError.Error()
}
