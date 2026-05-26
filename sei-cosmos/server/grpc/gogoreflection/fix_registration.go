package gogoreflection

import (
	"reflect"

	_ "github.com/gogo/protobuf/gogoproto" // required so it does register the gogoproto file descriptor
	gogoproto "github.com/gogo/protobuf/proto"

	// nolint: staticcheck
	"github.com/golang/protobuf/proto"
	dpb "github.com/golang/protobuf/protoc-gen-go/descriptor"
	_ "github.com/regen-network/cosmos-proto" // look above
)

var importsToFix = map[string]string{
	"gogo.proto":   "gogoproto/gogo.proto",
	"cosmos.proto": "cosmos_proto/cosmos.proto",
}

// fixRegistration is required because certain files register themselves in a way
// but are imported by other files in a different way.
// NOTE(fdymylja): This fix should not be needed and should be addressed in some CI.
// Currently every cosmos-sdk proto file is importing gogo.proto as gogoproto/gogo.proto,
// but gogo.proto registers itself as gogo.proto, same goes for cosmos.proto.
func fixRegistration(registeredAs, importedAs string) error { _ = "STUB: not implemented"; return nil }

// fix name

func init() {
	// we need to fix the gogoproto filedesc to match the import path
	// in theory this shouldn't be required, generally speaking
	// proto files should be imported as their registration path

	for registeredAs, importedAs := range importsToFix {
		err := fixRegistration(registeredAs, importedAs)
		if err != nil {
			panic(err)
		}
	}
}

// compress compresses the given file descriptor

func compress(fd *dpb.FileDescriptorProto) ([]byte, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getFileDescriptor(filePath string) []byte {
	_ = "STUB: not implemented"
	// since we got well known descriptors which are not registered into gogoproto registry
	// but are instead registered into the proto one, we need to check both
	return nil
}

// nolint: staticcheck

func getMessageType(name string) reflect.Type { _ = "STUB: not implemented"; return *new(reflect.Type) }

// nolint: staticcheck

func getExtension(extID int32, m proto.Message) *gogoproto.ExtensionDesc {
	_ = "STUB: not implemented"
	// check first in gogoproto registry
	return nil
}

// check into proto registry
// nolint: staticcheck

func getExtensionsNumbers(m proto.Message) []int32 { _ = "STUB: not implemented"; return nil }

// nolint: staticcheck
