package descriptor

import (
	"google.golang.org/protobuf/types/descriptorpb"
	"google.golang.org/protobuf/types/pluginpb"
)

type SourceFile struct {
	*descriptorpb.FileDescriptorProto

	Package                 string
	GeneratedFilenamePrefix string
}

type TargetFile struct {
	*pluginpb.CodeGeneratorResponse_File
}
