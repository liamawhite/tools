// Code generated by protoc-gen-gogo. DO NOT EDIT.
// source: external.proto

package generated

import bytes "bytes"
import github_com_gogo_protobuf_jsonpb "github.com/gogo/protobuf/jsonpb"
import proto "github.com/gogo/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// MarshalJSON is a custom marshaler supporting oneof fields for ExternalOneof
func (this *ExternalOneof) MarshalJSON() ([]byte, error) {
	str, err := ExternalMarshaler.MarshalToString(this)
	return []byte(str), err
}

// UnmarshalJSON is a custom unmarshaler supporting oneof fields for ExternalOneof
func (this *ExternalOneof) UnmarshalJSON(b []byte) error {
	return ExternalUnmarshaler.Unmarshal(bytes.NewReader(b), this)
}

var (
	ExternalMarshaler   = &github_com_gogo_protobuf_jsonpb.Marshaler{}
	ExternalUnmarshaler = &github_com_gogo_protobuf_jsonpb.Unmarshaler{}
)
