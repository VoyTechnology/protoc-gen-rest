package registry

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/voytechnology/protoc-gen-rest/internal/descriptor"
	"google.golang.org/protobuf/compiler/protogen"
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

	return &descriptor.SourceFile{
		FileDescriptorProto:     file.Proto,
		GeneratedFilenamePrefix: filepath.Join(pathName, fileName),
	}, nil
}
