package errcodes

type errorCode uint8

const (
	DBUnknown errorCode = iota
	DBNoData
	DBConflict
	Validation
	Unknown
	NoData
)

func (e errorCode) Error() string { return "" }
