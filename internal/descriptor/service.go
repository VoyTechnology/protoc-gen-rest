package descriptor

type Method struct {
	Name         string
	InternalName string
	RequestType  string
	ResponseType string
	Path         string
	Method       string
	Body         string
}

type Service struct {
	Name         string
	InternalName string

	Methods map[string]map[string]*Method
}
