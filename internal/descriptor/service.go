package descriptor

type Method struct {
	Name         string
	RequestType  string
	ResponseType string
	Path         string
	Method       string
	Body         string
}

type Service struct {
	Name    string
	Methods []Method
}
