package packet

type OnlyEndOfBlockError struct {
	Content string
}

func (err *OnlyEndOfBlockError) Error() string {
	return "packet: error: has only end of block (GS): " + err.Content
}

type NoEndOfBlockError struct {
	Content string
}

func (err *NoEndOfBlockError) Error() string {
	return "packet: error: no end of block (GS): " + err.Content
}

type NoHeaderError struct {
	Content string
}

func (err *NoHeaderError) Error() string {
	return "packet: error: no header: " + err.Content
}
