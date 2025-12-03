package packet

type NullContentError struct{}

func (err *NullContentError) Error() string {
	return "null content"
}
