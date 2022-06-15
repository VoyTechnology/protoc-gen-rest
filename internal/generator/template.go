package generator

import "text/template"

var header = template.Must(template.New("header").Parse(`
// Code generated by protoc-gen-rest. DO NOT EDIT.
// source: {{ .Source }}

{{ range .Services }}
type {{ .Name}}Server interface {
	{{ range .Methods }}
	{{ .Name }}(ctx context.Context, req *{{ .RequestType }}) (*{{ .ResponseType }}, error)
	{{ end }}
}
{{ end}}
`))
