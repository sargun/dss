package dsscommon

import proto "code.google.com/p/goprotobuf/proto"

type DataTypeMethod struct {
	Func func(proto.Message) proto.Message
}
type DataType struct {
	Name    string
	Methods map[string]DataTypeMethod
}
