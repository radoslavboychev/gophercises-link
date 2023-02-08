package parsingerrors

const (
	ErrOpeningFile  = Error("Error opening file")
	ErrEOF          = Error("End of file")
	ErrNoLinksFound = Error("No links found")
)

type Error string

func (e Error) Error() string {
	return string(e)
}
