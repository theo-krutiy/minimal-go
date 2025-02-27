package core

type Error uint8

const (
	ErrDBUnknown Error = iota
	ErrDBNoData
	ErrDBConflict
	ErrValidation
	ErrUnknown
	ErrNoData
)

func (e Error) Error() string { return "" }
