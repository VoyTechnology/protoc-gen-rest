package registry

import (
	"fmt"
	"net/http"
	"path/filepath"
	"strings"

	"github.com/voytechnology/protoc-gen-rest/internal/descriptor"
	options "google.golang.org/genproto/googleapis/api/annotations"
	"google.golang.org/protobuf/compiler/protogen"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/descriptorpb"
)

type Registry struct {
	files map[string]*protogen.File
}

// NewRegistry creates a new registry which contains all the information
// about all the protobufs
func NewRegistry(gen *protogen.Plugin) *Registry {
	return &Registry{gen.FilesByPath}
}

func (r *Registry) Lookup(fileName string) (*descriptor.SourceFile, error) {
	file, ok := r.files[fileName]
	if !ok {
		return nil, fmt.Errorf("file %s not found", fileName)
	}

	pathName, fileName := filepath.Split(fileName)
	fileName = strings.TrimSuffix(fileName, filepath.Ext(fileName))

	var services []*descriptor.Service
	for _, service := range file.Proto.Service {
		svc := &descriptor.Service{
			Name:         service.GetName(),
			InternalName: strings.ToLower(service.GetName()),
			Methods:      make(map[string]map[string]*descriptor.Method),
		}
		for _, method := range service.Method {
			httpMethod, httpPath, err := httpOptions(method)
			if err != nil {
				return nil, err
			}

			methods, exists := svc.Methods[httpPath]
			if !exists {
				methods = make(map[string]*descriptor.Method)

			}
			methods[httpMethod] = &descriptor.Method{
				Name:         method.GetName(),
				InternalName: strings.ToLower(method.GetName()),
				RequestType:  method.GetInputType(),
				ResponseType: method.GetOutputType(),
				Path:         httpPath,
				Method:       httpMethod,
			}

			svc.Methods[httpPath] = methods

		}
		services = append(services, svc)
	}

	return &descriptor.SourceFile{
		FileDescriptorProto:     file.Proto,
		Package:                 string(file.GoPackageName),
		GeneratedFilenamePrefix: filepath.Join(pathName, fileName),
		Services:                services,
	}, nil
}

func httpOptions(method *descriptorpb.MethodDescriptorProto) (string, string, error) {
	if method.Options == nil {
		return "", "", nil
	}

	if !proto.HasExtension(method.Options, options.E_Http) {
		return "", "", nil
	}
	ext := proto.GetExtension(method.Options, options.E_Http)
	opts, ok := ext.(*options.HttpRule)
	if !ok {
		return "", "", fmt.Errorf("extention is %T, want HttpRule", ext)
	}
	var httpMethod, httpPath string
	switch opts.Pattern.(type) {
	case *options.HttpRule_Get:
		httpMethod = http.MethodGet
		httpPath = opts.GetGet()
	case *options.HttpRule_Put:
		httpMethod = http.MethodPut
		httpPath = opts.GetPut()
	case *options.HttpRule_Post:
		httpMethod = http.MethodPost
		httpPath = opts.GetPost()
	case *options.HttpRule_Delete:
		httpMethod = http.MethodDelete
		httpPath = opts.GetDelete()
	case *options.HttpRule_Patch:
		httpMethod = http.MethodPatch
		httpPath = opts.GetPatch()
	case *options.HttpRule_Custom:
		httpMethod = opts.GetCustom().Kind
		httpPath = opts.GetCustom().Path
	}
	return httpMethod, httpPath, nil
}
