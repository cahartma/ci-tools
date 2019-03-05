// Code generated by protoc-gen-go. DO NOT EDIT.
// source: extra_actions_base_proto/extra_actions_base.proto

package extra_actions_base_proto

import proto "github.com/golang/protobuf/proto"
import fmt "fmt"
import math "math"

// Reference imports to suppress errors if they are not otherwise used.
var _ = proto.Marshal
var _ = fmt.Errorf
var _ = math.Inf

// This is a compile-time assertion to ensure that this generated file
// is compatible with the proto package it is being compiled against.
// A compilation error at this line likely means your copy of the
// proto package needs to be updated.
const _ = proto.ProtoPackageIsVersion2 // please upgrade the proto package

type ExtraActionSummary struct {
	Action               []*DetailedExtraActionInfo `protobuf:"bytes,1,rep,name=action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}                   `json:"-"`
	XXX_unrecognized     []byte                     `json:"-"`
	XXX_sizecache        int32                      `json:"-"`
}

func (m *ExtraActionSummary) Reset()         { *m = ExtraActionSummary{} }
func (m *ExtraActionSummary) String() string { return proto.CompactTextString(m) }
func (*ExtraActionSummary) ProtoMessage()    {}
func (*ExtraActionSummary) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{0}
}
func (m *ExtraActionSummary) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionSummary.Unmarshal(m, b)
}
func (m *ExtraActionSummary) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionSummary.Marshal(b, m, deterministic)
}
func (dst *ExtraActionSummary) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionSummary.Merge(dst, src)
}
func (m *ExtraActionSummary) XXX_Size() int {
	return xxx_messageInfo_ExtraActionSummary.Size(m)
}
func (m *ExtraActionSummary) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionSummary.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionSummary proto.InternalMessageInfo

func (m *ExtraActionSummary) GetAction() []*DetailedExtraActionInfo {
	if m != nil {
		return m.Action
	}
	return nil
}

type DetailedExtraActionInfo struct {
	TriggeringFile       *string          `protobuf:"bytes,1,opt,name=triggering_file,json=triggeringFile" json:"triggering_file,omitempty"`
	Action               *ExtraActionInfo `protobuf:"bytes,2,req,name=action" json:"action,omitempty"`
	XXX_NoUnkeyedLiteral struct{}         `json:"-"`
	XXX_unrecognized     []byte           `json:"-"`
	XXX_sizecache        int32            `json:"-"`
}

func (m *DetailedExtraActionInfo) Reset()         { *m = DetailedExtraActionInfo{} }
func (m *DetailedExtraActionInfo) String() string { return proto.CompactTextString(m) }
func (*DetailedExtraActionInfo) ProtoMessage()    {}
func (*DetailedExtraActionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{1}
}
func (m *DetailedExtraActionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_DetailedExtraActionInfo.Unmarshal(m, b)
}
func (m *DetailedExtraActionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_DetailedExtraActionInfo.Marshal(b, m, deterministic)
}
func (dst *DetailedExtraActionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_DetailedExtraActionInfo.Merge(dst, src)
}
func (m *DetailedExtraActionInfo) XXX_Size() int {
	return xxx_messageInfo_DetailedExtraActionInfo.Size(m)
}
func (m *DetailedExtraActionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_DetailedExtraActionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_DetailedExtraActionInfo proto.InternalMessageInfo

func (m *DetailedExtraActionInfo) GetTriggeringFile() string {
	if m != nil && m.TriggeringFile != nil {
		return *m.TriggeringFile
	}
	return ""
}

func (m *DetailedExtraActionInfo) GetAction() *ExtraActionInfo {
	if m != nil {
		return m.Action
	}
	return nil
}

type ExtraActionInfo struct {
	Owner                        *string                                `protobuf:"bytes,1,opt,name=owner" json:"owner,omitempty"`
	AspectName                   *string                                `protobuf:"bytes,6,opt,name=aspect_name,json=aspectName" json:"aspect_name,omitempty"`                                                                                             // Deprecated: Do not use.
	AspectParameters             map[string]*ExtraActionInfo_StringList `protobuf:"bytes,7,rep,name=aspect_parameters,json=aspectParameters" json:"aspect_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"` // Deprecated: Do not use.
	Aspects                      []*ExtraActionInfo_AspectDescriptor    `protobuf:"bytes,8,rep,name=aspects" json:"aspects,omitempty"`
	Id                           *string                                `protobuf:"bytes,2,opt,name=id" json:"id,omitempty"`
	Mnemonic                     *string                                `protobuf:"bytes,5,opt,name=mnemonic" json:"mnemonic,omitempty"`
	XXX_NoUnkeyedLiteral         struct{}                               `json:"-"`
	proto.XXX_InternalExtensions `json:"-"`
	XXX_unrecognized             []byte `json:"-"`
	XXX_sizecache                int32  `json:"-"`
}

func (m *ExtraActionInfo) Reset()         { *m = ExtraActionInfo{} }
func (m *ExtraActionInfo) String() string { return proto.CompactTextString(m) }
func (*ExtraActionInfo) ProtoMessage()    {}
func (*ExtraActionInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{2}
}

var extRange_ExtraActionInfo = []proto.ExtensionRange{
	{Start: 1000, End: 536870911},
}

func (*ExtraActionInfo) ExtensionRangeArray() []proto.ExtensionRange {
	return extRange_ExtraActionInfo
}
func (m *ExtraActionInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionInfo.Unmarshal(m, b)
}
func (m *ExtraActionInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionInfo.Marshal(b, m, deterministic)
}
func (dst *ExtraActionInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionInfo.Merge(dst, src)
}
func (m *ExtraActionInfo) XXX_Size() int {
	return xxx_messageInfo_ExtraActionInfo.Size(m)
}
func (m *ExtraActionInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionInfo.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionInfo proto.InternalMessageInfo

func (m *ExtraActionInfo) GetOwner() string {
	if m != nil && m.Owner != nil {
		return *m.Owner
	}
	return ""
}

// Deprecated: Do not use.
func (m *ExtraActionInfo) GetAspectName() string {
	if m != nil && m.AspectName != nil {
		return *m.AspectName
	}
	return ""
}

// Deprecated: Do not use.
func (m *ExtraActionInfo) GetAspectParameters() map[string]*ExtraActionInfo_StringList {
	if m != nil {
		return m.AspectParameters
	}
	return nil
}

func (m *ExtraActionInfo) GetAspects() []*ExtraActionInfo_AspectDescriptor {
	if m != nil {
		return m.Aspects
	}
	return nil
}

func (m *ExtraActionInfo) GetId() string {
	if m != nil && m.Id != nil {
		return *m.Id
	}
	return ""
}

func (m *ExtraActionInfo) GetMnemonic() string {
	if m != nil && m.Mnemonic != nil {
		return *m.Mnemonic
	}
	return ""
}

// Deprecated: Do not use.
type ExtraActionInfo_StringList struct {
	Value                []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraActionInfo_StringList) Reset()         { *m = ExtraActionInfo_StringList{} }
func (m *ExtraActionInfo_StringList) String() string { return proto.CompactTextString(m) }
func (*ExtraActionInfo_StringList) ProtoMessage()    {}
func (*ExtraActionInfo_StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{2, 1}
}
func (m *ExtraActionInfo_StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionInfo_StringList.Unmarshal(m, b)
}
func (m *ExtraActionInfo_StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionInfo_StringList.Marshal(b, m, deterministic)
}
func (dst *ExtraActionInfo_StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionInfo_StringList.Merge(dst, src)
}
func (m *ExtraActionInfo_StringList) XXX_Size() int {
	return xxx_messageInfo_ExtraActionInfo_StringList.Size(m)
}
func (m *ExtraActionInfo_StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionInfo_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionInfo_StringList proto.InternalMessageInfo

func (m *ExtraActionInfo_StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type ExtraActionInfo_AspectDescriptor struct {
	AspectName           *string                                                 `protobuf:"bytes,1,opt,name=aspect_name,json=aspectName" json:"aspect_name,omitempty"`
	AspectParameters     map[string]*ExtraActionInfo_AspectDescriptor_StringList `protobuf:"bytes,2,rep,name=aspect_parameters,json=aspectParameters" json:"aspect_parameters,omitempty" protobuf_key:"bytes,1,opt,name=key" protobuf_val:"bytes,2,opt,name=value"`
	XXX_NoUnkeyedLiteral struct{}                                                `json:"-"`
	XXX_unrecognized     []byte                                                  `json:"-"`
	XXX_sizecache        int32                                                   `json:"-"`
}

func (m *ExtraActionInfo_AspectDescriptor) Reset()         { *m = ExtraActionInfo_AspectDescriptor{} }
func (m *ExtraActionInfo_AspectDescriptor) String() string { return proto.CompactTextString(m) }
func (*ExtraActionInfo_AspectDescriptor) ProtoMessage()    {}
func (*ExtraActionInfo_AspectDescriptor) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{2, 2}
}
func (m *ExtraActionInfo_AspectDescriptor) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor.Unmarshal(m, b)
}
func (m *ExtraActionInfo_AspectDescriptor) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor.Marshal(b, m, deterministic)
}
func (dst *ExtraActionInfo_AspectDescriptor) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionInfo_AspectDescriptor.Merge(dst, src)
}
func (m *ExtraActionInfo_AspectDescriptor) XXX_Size() int {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor.Size(m)
}
func (m *ExtraActionInfo_AspectDescriptor) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionInfo_AspectDescriptor.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionInfo_AspectDescriptor proto.InternalMessageInfo

func (m *ExtraActionInfo_AspectDescriptor) GetAspectName() string {
	if m != nil && m.AspectName != nil {
		return *m.AspectName
	}
	return ""
}

func (m *ExtraActionInfo_AspectDescriptor) GetAspectParameters() map[string]*ExtraActionInfo_AspectDescriptor_StringList {
	if m != nil {
		return m.AspectParameters
	}
	return nil
}

type ExtraActionInfo_AspectDescriptor_StringList struct {
	Value                []string `protobuf:"bytes,1,rep,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *ExtraActionInfo_AspectDescriptor_StringList) Reset() {
	*m = ExtraActionInfo_AspectDescriptor_StringList{}
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) String() string {
	return proto.CompactTextString(m)
}
func (*ExtraActionInfo_AspectDescriptor_StringList) ProtoMessage() {}
func (*ExtraActionInfo_AspectDescriptor_StringList) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{2, 2, 1}
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.Unmarshal(m, b)
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.Marshal(b, m, deterministic)
}
func (dst *ExtraActionInfo_AspectDescriptor_StringList) XXX_Merge(src proto.Message) {
	xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.Merge(dst, src)
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_Size() int {
	return xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.Size(m)
}
func (m *ExtraActionInfo_AspectDescriptor_StringList) XXX_DiscardUnknown() {
	xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList.DiscardUnknown(m)
}

var xxx_messageInfo_ExtraActionInfo_AspectDescriptor_StringList proto.InternalMessageInfo

func (m *ExtraActionInfo_AspectDescriptor_StringList) GetValue() []string {
	if m != nil {
		return m.Value
	}
	return nil
}

type EnvironmentVariable struct {
	Name                 *string  `protobuf:"bytes,1,req,name=name" json:"name,omitempty"`
	Value                *string  `protobuf:"bytes,2,req,name=value" json:"value,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *EnvironmentVariable) Reset()         { *m = EnvironmentVariable{} }
func (m *EnvironmentVariable) String() string { return proto.CompactTextString(m) }
func (*EnvironmentVariable) ProtoMessage()    {}
func (*EnvironmentVariable) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{3}
}
func (m *EnvironmentVariable) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_EnvironmentVariable.Unmarshal(m, b)
}
func (m *EnvironmentVariable) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_EnvironmentVariable.Marshal(b, m, deterministic)
}
func (dst *EnvironmentVariable) XXX_Merge(src proto.Message) {
	xxx_messageInfo_EnvironmentVariable.Merge(dst, src)
}
func (m *EnvironmentVariable) XXX_Size() int {
	return xxx_messageInfo_EnvironmentVariable.Size(m)
}
func (m *EnvironmentVariable) XXX_DiscardUnknown() {
	xxx_messageInfo_EnvironmentVariable.DiscardUnknown(m)
}

var xxx_messageInfo_EnvironmentVariable proto.InternalMessageInfo

func (m *EnvironmentVariable) GetName() string {
	if m != nil && m.Name != nil {
		return *m.Name
	}
	return ""
}

func (m *EnvironmentVariable) GetValue() string {
	if m != nil && m.Value != nil {
		return *m.Value
	}
	return ""
}

type SpawnInfo struct {
	Argument             []string               `protobuf:"bytes,1,rep,name=argument" json:"argument,omitempty"`
	Variable             []*EnvironmentVariable `protobuf:"bytes,2,rep,name=variable" json:"variable,omitempty"`
	InputFile            []string               `protobuf:"bytes,4,rep,name=input_file,json=inputFile" json:"input_file,omitempty"`
	OutputFile           []string               `protobuf:"bytes,5,rep,name=output_file,json=outputFile" json:"output_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *SpawnInfo) Reset()         { *m = SpawnInfo{} }
func (m *SpawnInfo) String() string { return proto.CompactTextString(m) }
func (*SpawnInfo) ProtoMessage()    {}
func (*SpawnInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{4}
}
func (m *SpawnInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_SpawnInfo.Unmarshal(m, b)
}
func (m *SpawnInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_SpawnInfo.Marshal(b, m, deterministic)
}
func (dst *SpawnInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_SpawnInfo.Merge(dst, src)
}
func (m *SpawnInfo) XXX_Size() int {
	return xxx_messageInfo_SpawnInfo.Size(m)
}
func (m *SpawnInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_SpawnInfo.DiscardUnknown(m)
}

var xxx_messageInfo_SpawnInfo proto.InternalMessageInfo

func (m *SpawnInfo) GetArgument() []string {
	if m != nil {
		return m.Argument
	}
	return nil
}

func (m *SpawnInfo) GetVariable() []*EnvironmentVariable {
	if m != nil {
		return m.Variable
	}
	return nil
}

func (m *SpawnInfo) GetInputFile() []string {
	if m != nil {
		return m.InputFile
	}
	return nil
}

func (m *SpawnInfo) GetOutputFile() []string {
	if m != nil {
		return m.OutputFile
	}
	return nil
}

var E_SpawnInfo_SpawnInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*SpawnInfo)(nil),
	Field:         1003,
	Name:          "blaze.SpawnInfo.spawn_info",
	Tag:           "bytes,1003,opt,name=spawn_info,json=spawnInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

type CppCompileInfo struct {
	Tool                 *string                `protobuf:"bytes,1,opt,name=tool" json:"tool,omitempty"`
	CompilerOption       []string               `protobuf:"bytes,2,rep,name=compiler_option,json=compilerOption" json:"compiler_option,omitempty"`
	SourceFile           *string                `protobuf:"bytes,3,opt,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	OutputFile           *string                `protobuf:"bytes,4,opt,name=output_file,json=outputFile" json:"output_file,omitempty"`
	SourcesAndHeaders    []string               `protobuf:"bytes,5,rep,name=sources_and_headers,json=sourcesAndHeaders" json:"sources_and_headers,omitempty"`
	Variable             []*EnvironmentVariable `protobuf:"bytes,6,rep,name=variable" json:"variable,omitempty"`
	XXX_NoUnkeyedLiteral struct{}               `json:"-"`
	XXX_unrecognized     []byte                 `json:"-"`
	XXX_sizecache        int32                  `json:"-"`
}

func (m *CppCompileInfo) Reset()         { *m = CppCompileInfo{} }
func (m *CppCompileInfo) String() string { return proto.CompactTextString(m) }
func (*CppCompileInfo) ProtoMessage()    {}
func (*CppCompileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{5}
}
func (m *CppCompileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CppCompileInfo.Unmarshal(m, b)
}
func (m *CppCompileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CppCompileInfo.Marshal(b, m, deterministic)
}
func (dst *CppCompileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppCompileInfo.Merge(dst, src)
}
func (m *CppCompileInfo) XXX_Size() int {
	return xxx_messageInfo_CppCompileInfo.Size(m)
}
func (m *CppCompileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CppCompileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CppCompileInfo proto.InternalMessageInfo

func (m *CppCompileInfo) GetTool() string {
	if m != nil && m.Tool != nil {
		return *m.Tool
	}
	return ""
}

func (m *CppCompileInfo) GetCompilerOption() []string {
	if m != nil {
		return m.CompilerOption
	}
	return nil
}

func (m *CppCompileInfo) GetSourceFile() string {
	if m != nil && m.SourceFile != nil {
		return *m.SourceFile
	}
	return ""
}

func (m *CppCompileInfo) GetOutputFile() string {
	if m != nil && m.OutputFile != nil {
		return *m.OutputFile
	}
	return ""
}

func (m *CppCompileInfo) GetSourcesAndHeaders() []string {
	if m != nil {
		return m.SourcesAndHeaders
	}
	return nil
}

func (m *CppCompileInfo) GetVariable() []*EnvironmentVariable {
	if m != nil {
		return m.Variable
	}
	return nil
}

var E_CppCompileInfo_CppCompileInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*CppCompileInfo)(nil),
	Field:         1001,
	Name:          "blaze.CppCompileInfo.cpp_compile_info",
	Tag:           "bytes,1001,opt,name=cpp_compile_info,json=cppCompileInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

type CppLinkInfo struct {
	InputFile               []string `protobuf:"bytes,1,rep,name=input_file,json=inputFile" json:"input_file,omitempty"`
	OutputFile              *string  `protobuf:"bytes,2,opt,name=output_file,json=outputFile" json:"output_file,omitempty"`
	InterfaceOutputFile     *string  `protobuf:"bytes,3,opt,name=interface_output_file,json=interfaceOutputFile" json:"interface_output_file,omitempty"`
	LinkTargetType          *string  `protobuf:"bytes,4,opt,name=link_target_type,json=linkTargetType" json:"link_target_type,omitempty"`
	LinkStaticness          *string  `protobuf:"bytes,5,opt,name=link_staticness,json=linkStaticness" json:"link_staticness,omitempty"`
	LinkStamp               []string `protobuf:"bytes,6,rep,name=link_stamp,json=linkStamp" json:"link_stamp,omitempty"`
	BuildInfoHeaderArtifact []string `protobuf:"bytes,7,rep,name=build_info_header_artifact,json=buildInfoHeaderArtifact" json:"build_info_header_artifact,omitempty"`
	LinkOpt                 []string `protobuf:"bytes,8,rep,name=link_opt,json=linkOpt" json:"link_opt,omitempty"`
	XXX_NoUnkeyedLiteral    struct{} `json:"-"`
	XXX_unrecognized        []byte   `json:"-"`
	XXX_sizecache           int32    `json:"-"`
}

func (m *CppLinkInfo) Reset()         { *m = CppLinkInfo{} }
func (m *CppLinkInfo) String() string { return proto.CompactTextString(m) }
func (*CppLinkInfo) ProtoMessage()    {}
func (*CppLinkInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{6}
}
func (m *CppLinkInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_CppLinkInfo.Unmarshal(m, b)
}
func (m *CppLinkInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_CppLinkInfo.Marshal(b, m, deterministic)
}
func (dst *CppLinkInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_CppLinkInfo.Merge(dst, src)
}
func (m *CppLinkInfo) XXX_Size() int {
	return xxx_messageInfo_CppLinkInfo.Size(m)
}
func (m *CppLinkInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_CppLinkInfo.DiscardUnknown(m)
}

var xxx_messageInfo_CppLinkInfo proto.InternalMessageInfo

func (m *CppLinkInfo) GetInputFile() []string {
	if m != nil {
		return m.InputFile
	}
	return nil
}

func (m *CppLinkInfo) GetOutputFile() string {
	if m != nil && m.OutputFile != nil {
		return *m.OutputFile
	}
	return ""
}

func (m *CppLinkInfo) GetInterfaceOutputFile() string {
	if m != nil && m.InterfaceOutputFile != nil {
		return *m.InterfaceOutputFile
	}
	return ""
}

func (m *CppLinkInfo) GetLinkTargetType() string {
	if m != nil && m.LinkTargetType != nil {
		return *m.LinkTargetType
	}
	return ""
}

func (m *CppLinkInfo) GetLinkStaticness() string {
	if m != nil && m.LinkStaticness != nil {
		return *m.LinkStaticness
	}
	return ""
}

func (m *CppLinkInfo) GetLinkStamp() []string {
	if m != nil {
		return m.LinkStamp
	}
	return nil
}

func (m *CppLinkInfo) GetBuildInfoHeaderArtifact() []string {
	if m != nil {
		return m.BuildInfoHeaderArtifact
	}
	return nil
}

func (m *CppLinkInfo) GetLinkOpt() []string {
	if m != nil {
		return m.LinkOpt
	}
	return nil
}

var E_CppLinkInfo_CppLinkInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*CppLinkInfo)(nil),
	Field:         1002,
	Name:          "blaze.CppLinkInfo.cpp_link_info",
	Tag:           "bytes,1002,opt,name=cpp_link_info,json=cppLinkInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

type JavaCompileInfo struct {
	Outputjar            *string  `protobuf:"bytes,1,opt,name=outputjar" json:"outputjar,omitempty"`
	Classpath            []string `protobuf:"bytes,2,rep,name=classpath" json:"classpath,omitempty"`
	Sourcepath           []string `protobuf:"bytes,3,rep,name=sourcepath" json:"sourcepath,omitempty"`
	SourceFile           []string `protobuf:"bytes,4,rep,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	JavacOpt             []string `protobuf:"bytes,5,rep,name=javac_opt,json=javacOpt" json:"javac_opt,omitempty"`
	Processor            []string `protobuf:"bytes,6,rep,name=processor" json:"processor,omitempty"`
	Processorpath        []string `protobuf:"bytes,7,rep,name=processorpath" json:"processorpath,omitempty"`
	Bootclasspath        []string `protobuf:"bytes,8,rep,name=bootclasspath" json:"bootclasspath,omitempty"`
	Argument             []string `protobuf:"bytes,9,rep,name=argument" json:"argument,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *JavaCompileInfo) Reset()         { *m = JavaCompileInfo{} }
func (m *JavaCompileInfo) String() string { return proto.CompactTextString(m) }
func (*JavaCompileInfo) ProtoMessage()    {}
func (*JavaCompileInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{7}
}
func (m *JavaCompileInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_JavaCompileInfo.Unmarshal(m, b)
}
func (m *JavaCompileInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_JavaCompileInfo.Marshal(b, m, deterministic)
}
func (dst *JavaCompileInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_JavaCompileInfo.Merge(dst, src)
}
func (m *JavaCompileInfo) XXX_Size() int {
	return xxx_messageInfo_JavaCompileInfo.Size(m)
}
func (m *JavaCompileInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_JavaCompileInfo.DiscardUnknown(m)
}

var xxx_messageInfo_JavaCompileInfo proto.InternalMessageInfo

func (m *JavaCompileInfo) GetOutputjar() string {
	if m != nil && m.Outputjar != nil {
		return *m.Outputjar
	}
	return ""
}

func (m *JavaCompileInfo) GetClasspath() []string {
	if m != nil {
		return m.Classpath
	}
	return nil
}

func (m *JavaCompileInfo) GetSourcepath() []string {
	if m != nil {
		return m.Sourcepath
	}
	return nil
}

func (m *JavaCompileInfo) GetSourceFile() []string {
	if m != nil {
		return m.SourceFile
	}
	return nil
}

func (m *JavaCompileInfo) GetJavacOpt() []string {
	if m != nil {
		return m.JavacOpt
	}
	return nil
}

func (m *JavaCompileInfo) GetProcessor() []string {
	if m != nil {
		return m.Processor
	}
	return nil
}

func (m *JavaCompileInfo) GetProcessorpath() []string {
	if m != nil {
		return m.Processorpath
	}
	return nil
}

func (m *JavaCompileInfo) GetBootclasspath() []string {
	if m != nil {
		return m.Bootclasspath
	}
	return nil
}

func (m *JavaCompileInfo) GetArgument() []string {
	if m != nil {
		return m.Argument
	}
	return nil
}

var E_JavaCompileInfo_JavaCompileInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*JavaCompileInfo)(nil),
	Field:         1000,
	Name:          "blaze.JavaCompileInfo.java_compile_info",
	Tag:           "bytes,1000,opt,name=java_compile_info,json=javaCompileInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

type PythonInfo struct {
	SourceFile           []string `protobuf:"bytes,1,rep,name=source_file,json=sourceFile" json:"source_file,omitempty"`
	DepFile              []string `protobuf:"bytes,2,rep,name=dep_file,json=depFile" json:"dep_file,omitempty"`
	XXX_NoUnkeyedLiteral struct{} `json:"-"`
	XXX_unrecognized     []byte   `json:"-"`
	XXX_sizecache        int32    `json:"-"`
}

func (m *PythonInfo) Reset()         { *m = PythonInfo{} }
func (m *PythonInfo) String() string { return proto.CompactTextString(m) }
func (*PythonInfo) ProtoMessage()    {}
func (*PythonInfo) Descriptor() ([]byte, []int) {
	return fileDescriptor_extra_actions_base_a50b304fd81425db, []int{8}
}
func (m *PythonInfo) XXX_Unmarshal(b []byte) error {
	return xxx_messageInfo_PythonInfo.Unmarshal(m, b)
}
func (m *PythonInfo) XXX_Marshal(b []byte, deterministic bool) ([]byte, error) {
	return xxx_messageInfo_PythonInfo.Marshal(b, m, deterministic)
}
func (dst *PythonInfo) XXX_Merge(src proto.Message) {
	xxx_messageInfo_PythonInfo.Merge(dst, src)
}
func (m *PythonInfo) XXX_Size() int {
	return xxx_messageInfo_PythonInfo.Size(m)
}
func (m *PythonInfo) XXX_DiscardUnknown() {
	xxx_messageInfo_PythonInfo.DiscardUnknown(m)
}

var xxx_messageInfo_PythonInfo proto.InternalMessageInfo

func (m *PythonInfo) GetSourceFile() []string {
	if m != nil {
		return m.SourceFile
	}
	return nil
}

func (m *PythonInfo) GetDepFile() []string {
	if m != nil {
		return m.DepFile
	}
	return nil
}

var E_PythonInfo_PythonInfo = &proto.ExtensionDesc{
	ExtendedType:  (*ExtraActionInfo)(nil),
	ExtensionType: (*PythonInfo)(nil),
	Field:         1005,
	Name:          "blaze.PythonInfo.python_info",
	Tag:           "bytes,1005,opt,name=python_info,json=pythonInfo",
	Filename:      "extra_actions_base_proto/extra_actions_base.proto",
}

func init() {
	proto.RegisterType((*ExtraActionSummary)(nil), "blaze.ExtraActionSummary")
	proto.RegisterType((*DetailedExtraActionInfo)(nil), "blaze.DetailedExtraActionInfo")
	proto.RegisterType((*ExtraActionInfo)(nil), "blaze.ExtraActionInfo")
	proto.RegisterMapType((map[string]*ExtraActionInfo_StringList)(nil), "blaze.ExtraActionInfo.AspectParametersEntry")
	proto.RegisterType((*ExtraActionInfo_StringList)(nil), "blaze.ExtraActionInfo.StringList")
	proto.RegisterType((*ExtraActionInfo_AspectDescriptor)(nil), "blaze.ExtraActionInfo.AspectDescriptor")
	proto.RegisterMapType((map[string]*ExtraActionInfo_AspectDescriptor_StringList)(nil), "blaze.ExtraActionInfo.AspectDescriptor.AspectParametersEntry")
	proto.RegisterType((*ExtraActionInfo_AspectDescriptor_StringList)(nil), "blaze.ExtraActionInfo.AspectDescriptor.StringList")
	proto.RegisterType((*EnvironmentVariable)(nil), "blaze.EnvironmentVariable")
	proto.RegisterType((*SpawnInfo)(nil), "blaze.SpawnInfo")
	proto.RegisterType((*CppCompileInfo)(nil), "blaze.CppCompileInfo")
	proto.RegisterType((*CppLinkInfo)(nil), "blaze.CppLinkInfo")
	proto.RegisterType((*JavaCompileInfo)(nil), "blaze.JavaCompileInfo")
	proto.RegisterType((*PythonInfo)(nil), "blaze.PythonInfo")
	proto.RegisterExtension(E_SpawnInfo_SpawnInfo)
	proto.RegisterExtension(E_CppCompileInfo_CppCompileInfo)
	proto.RegisterExtension(E_CppLinkInfo_CppLinkInfo)
	proto.RegisterExtension(E_JavaCompileInfo_JavaCompileInfo)
	proto.RegisterExtension(E_PythonInfo_PythonInfo)
}

func init() {
	proto.RegisterFile("extra_actions_base_proto/extra_actions_base.proto", fileDescriptor_extra_actions_base_a50b304fd81425db)
}

var fileDescriptor_extra_actions_base_a50b304fd81425db = []byte{
	// 1045 bytes of a gzipped FileDescriptorProto
	0x1f, 0x8b, 0x08, 0x00, 0x00, 0x00, 0x00, 0x00, 0x02, 0xff, 0x8c, 0x56, 0xdd, 0x6e, 0xe3, 0x44,
	0x14, 0x96, 0x9d, 0xa6, 0x8d, 0x4f, 0xb4, 0x69, 0x3a, 0xa5, 0x6c, 0x36, 0xc0, 0x6e, 0x09, 0x88,
	0xad, 0x00, 0x79, 0x45, 0x2e, 0x16, 0x04, 0x42, 0xa8, 0xdb, 0x2d, 0x2a, 0x50, 0xd1, 0xca, 0x5d,
	0x21, 0x24, 0x84, 0xac, 0x89, 0x3d, 0x4d, 0xa7, 0xb5, 0x3d, 0xa3, 0xf1, 0x24, 0x25, 0x5c, 0xf5,
	0x09, 0xb8, 0xe3, 0x8e, 0x47, 0xe1, 0x25, 0x78, 0x8b, 0x5d, 0x10, 0x77, 0x3c, 0x00, 0x9a, 0x1f,
	0xdb, 0x75, 0x9a, 0x76, 0x73, 0x37, 0xf3, 0x9d, 0xcf, 0x67, 0xce, 0xf9, 0xbe, 0x99, 0x93, 0xc0,
	0x27, 0xe4, 0x17, 0x29, 0x70, 0x88, 0x23, 0x49, 0x59, 0x96, 0x87, 0x23, 0x9c, 0x93, 0x90, 0x0b,
	0x26, 0xd9, 0x93, 0x9b, 0x01, 0x5f, 0x07, 0x50, 0x73, 0x94, 0xe0, 0x5f, 0xc9, 0xe0, 0x10, 0xd0,
	0xbe, 0xa2, 0xec, 0x6a, 0xc6, 0xc9, 0x24, 0x4d, 0xb1, 0x98, 0xa1, 0xa7, 0xb0, 0x6a, 0x3e, 0xe9,
	0x39, 0xdb, 0x8d, 0x9d, 0xf6, 0xf0, 0xa1, 0xaf, 0xd9, 0xfe, 0x73, 0x22, 0x31, 0x4d, 0x48, 0x7c,
	0xed, 0x93, 0x6f, 0xb2, 0x53, 0x16, 0x58, 0xf6, 0x40, 0xc0, 0xfd, 0x5b, 0x28, 0xe8, 0x31, 0xac,
	0x4b, 0x41, 0xc7, 0x63, 0x22, 0x68, 0x36, 0x0e, 0x4f, 0x69, 0x42, 0x7a, 0xce, 0xb6, 0xb3, 0xe3,
	0x05, 0x9d, 0x0a, 0xfe, 0x9a, 0x26, 0x04, 0xf9, 0xe5, 0xd9, 0xee, 0xb6, 0xbb, 0xd3, 0x1e, 0xbe,
	0x69, 0xcf, 0xbe, 0xed, 0xcc, 0xff, 0x9a, 0xb0, 0x3e, 0x7f, 0xd8, 0x1b, 0xd0, 0x64, 0x97, 0x19,
	0x11, 0xf6, 0x08, 0xb3, 0x41, 0xef, 0x41, 0x1b, 0xe7, 0x9c, 0x44, 0x32, 0xcc, 0x70, 0x4a, 0x7a,
	0xab, 0x2a, 0xf6, 0xcc, 0xed, 0x39, 0x01, 0x18, 0xf8, 0x7b, 0x9c, 0x12, 0xf4, 0x33, 0x6c, 0x58,
	0x12, 0xc7, 0x02, 0xa7, 0x44, 0x12, 0x91, 0xf7, 0xd6, 0xb4, 0x0a, 0x1f, 0x2f, 0xae, 0xc4, 0xdf,
	0xd5, 0xfc, 0xe3, 0x92, 0xbe, 0x9f, 0x49, 0x31, 0xd3, 0x89, 0xbb, 0x78, 0x2e, 0x84, 0x76, 0x61,
	0xcd, 0x60, 0x79, 0xaf, 0xa5, 0x93, 0x3e, 0xbe, 0x33, 0xe9, 0x73, 0x92, 0x47, 0x82, 0x72, 0xc9,
	0x44, 0x50, 0x7c, 0x87, 0x3a, 0xe0, 0xd2, 0xb8, 0xe7, 0xea, 0xce, 0x5c, 0x1a, 0xa3, 0x3e, 0xb4,
	0xd2, 0x8c, 0xa4, 0x2c, 0xa3, 0x51, 0xaf, 0xa9, 0xd1, 0x72, 0xdf, 0x3f, 0x85, 0xad, 0x85, 0xd5,
	0xa1, 0x2e, 0x34, 0x2e, 0xc8, 0xcc, 0xea, 0xa3, 0x96, 0xe8, 0x53, 0x68, 0x4e, 0x71, 0x32, 0x21,
	0x3a, 0x73, 0x7b, 0xf8, 0xee, 0x2d, 0x75, 0x9d, 0x48, 0xe5, 0xd4, 0x21, 0xcd, 0x65, 0x60, 0xf8,
	0x9f, 0xbb, 0x9f, 0x39, 0xfd, 0x0f, 0x00, 0xaa, 0x80, 0x92, 0xdf, 0xa4, 0x52, 0xb7, 0xc7, 0x2b,
	0x79, 0x3d, 0xa7, 0xff, 0xa7, 0x0b, 0xdd, 0xf9, 0xce, 0xd0, 0xa3, 0xba, 0x2f, 0xa6, 0xa6, 0xeb,
	0x9e, 0x9c, 0x2f, 0xf2, 0xc4, 0xd5, 0xf2, 0x7d, 0xb9, 0xa4, 0x7c, 0x8b, 0x4d, 0xba, 0x69, 0x50,
	0xff, 0x72, 0x79, 0xc5, 0x0e, 0xea, 0x8a, 0x0d, 0x97, 0x2d, 0x65, 0xb1, 0x84, 0x83, 0xd7, 0x4b,
	0xf8, 0xa1, 0xd7, 0x7a, 0xb9, 0xd6, 0xbd, 0xba, 0xba, 0xba, 0x72, 0x07, 0x5f, 0xc1, 0xe6, 0x7e,
	0x36, 0xa5, 0x82, 0x65, 0x29, 0xc9, 0xe4, 0x0f, 0x58, 0x50, 0x3c, 0x4a, 0x08, 0x42, 0xb0, 0x62,
	0x45, 0x74, 0x77, 0xbc, 0x40, 0xaf, 0xab, 0x5c, 0xae, 0x06, 0xcd, 0x66, 0xf0, 0xca, 0x01, 0xef,
	0x84, 0xe3, 0x4b, 0xf3, 0x62, 0xfa, 0xd0, 0xc2, 0x62, 0x3c, 0x51, 0xb9, 0xec, 0x91, 0xe5, 0x1e,
	0x3d, 0x85, 0xd6, 0xd4, 0xe6, 0xb7, 0xaa, 0xf7, 0x8b, 0x56, 0x6f, 0x56, 0x10, 0x94, 0x5c, 0xf4,
	0x0e, 0x00, 0xcd, 0xf8, 0x44, 0x9a, 0xd7, 0xbe, 0xa2, 0xb3, 0x7a, 0x1a, 0xd1, 0x0f, 0xfd, 0x11,
	0xb4, 0xd9, 0x44, 0x96, 0xf1, 0xa6, 0x8e, 0x83, 0x81, 0x14, 0x61, 0x78, 0x00, 0x90, 0xab, 0x02,
	0x43, 0xaa, 0x2a, 0xbc, 0x65, 0x0e, 0xf4, 0xfe, 0x59, 0xd3, 0xea, 0x77, 0x6d, 0xb8, 0x6c, 0x29,
	0xf0, 0xf2, 0x62, 0x39, 0xf8, 0xcb, 0x85, 0xce, 0x1e, 0xe7, 0x7b, 0x2c, 0xe5, 0x34, 0x21, 0xba,
	0x61, 0x04, 0x2b, 0x92, 0xb1, 0xc4, 0xfa, 0xa9, 0xd7, 0x6a, 0x46, 0x45, 0x86, 0x22, 0x42, 0xc6,
	0xed, 0x0c, 0x52, 0x55, 0x75, 0x0a, 0xf8, 0x48, 0xa3, 0xaa, 0xf4, 0x9c, 0x4d, 0x44, 0x44, 0x4c,
	0xe9, 0x0d, 0x73, 0x63, 0x0d, 0xb4, 0xa8, 0xb7, 0x15, 0x43, 0xa8, 0x7a, 0x43, 0x3e, 0x6c, 0x1a,
	0x7a, 0x1e, 0xe2, 0x2c, 0x0e, 0xcf, 0x08, 0x8e, 0xd5, 0xa5, 0x36, 0x22, 0x6c, 0xd8, 0xd0, 0x6e,
	0x16, 0x1f, 0x98, 0x40, 0xcd, 0x83, 0xd5, 0xe5, 0x3d, 0x18, 0xfe, 0x08, 0xdd, 0x88, 0xf3, 0xd0,
	0xd6, 0x7f, 0xb7, 0x92, 0xaf, 0x8c, 0x92, 0x5b, 0x36, 0x5c, 0x17, 0x2c, 0xe8, 0x44, 0xb5, 0xfd,
	0xe0, 0x8f, 0x06, 0xb4, 0xf7, 0x38, 0x3f, 0xa4, 0xd9, 0x85, 0x16, 0xb4, 0xee, 0xb6, 0xf3, 0x1a,
	0xb7, 0xdd, 0x1b, 0x8a, 0x0c, 0x61, 0x8b, 0x66, 0x92, 0x88, 0x53, 0x1c, 0x91, 0xf0, 0x3a, 0xd5,
	0xa8, 0xbb, 0x59, 0x06, 0x8f, 0xaa, 0x6f, 0x76, 0xa0, 0x9b, 0xd0, 0xec, 0x22, 0x94, 0x58, 0x8c,
	0x89, 0x0c, 0xe5, 0x8c, 0x17, 0x5a, 0x77, 0x14, 0xfe, 0x42, 0xc3, 0x2f, 0x66, 0x9c, 0x28, 0x6b,
	0x35, 0x33, 0x97, 0x58, 0xd2, 0x28, 0x23, 0x79, 0x6e, 0x67, 0xa5, 0x26, 0x9e, 0x94, 0xa8, 0x6a,
	0xa3, 0x20, 0xa6, 0x5c, 0x4b, 0xed, 0x05, 0x9e, 0xe5, 0xa4, 0x1c, 0x7d, 0x01, 0xfd, 0xd1, 0x84,
	0x26, 0xb1, 0x56, 0xd2, 0xda, 0x16, 0x62, 0x21, 0xe9, 0x29, 0x8e, 0xa4, 0xfe, 0x9d, 0xf0, 0x82,
	0xfb, 0x9a, 0xa1, 0x44, 0x31, 0xee, 0xed, 0xda, 0x30, 0x7a, 0x00, 0x2d, 0x9d, 0x9b, 0x71, 0xa9,
	0xa7, 0xbf, 0x17, 0xac, 0xa9, 0xfd, 0x11, 0x97, 0xc3, 0x23, 0xb8, 0xa7, 0x7c, 0xd2, 0xe1, 0x3b,
	0x4d, 0xfa, 0xdb, 0x98, 0x84, 0x2a, 0x93, 0x0a, 0x07, 0x82, 0x76, 0x54, 0x6d, 0x06, 0xbf, 0x35,
	0x60, 0xfd, 0x5b, 0x3c, 0xc5, 0xd7, 0xef, 0xfc, 0xdb, 0xe0, 0x19, 0x61, 0xcf, 0x71, 0xf1, 0xd3,
	0x58, 0x01, 0x2a, 0x1a, 0x25, 0x38, 0xcf, 0x39, 0x96, 0x67, 0xf6, 0xde, 0x57, 0x00, 0x7a, 0x08,
	0xf6, 0x7e, 0xeb, 0x70, 0xc3, 0x3c, 0xd6, 0x0a, 0x99, 0x7f, 0x12, 0x2b, 0xd7, 0x09, 0xda, 0xab,
	0xb7, 0xc0, 0x3b, 0xc7, 0x53, 0x1c, 0xe9, 0xee, 0xcd, 0x3d, 0x6f, 0x69, 0xe0, 0x88, 0x4b, 0x75,
	0x36, 0x17, 0x2c, 0x22, 0x79, 0xce, 0x44, 0x21, 0x7a, 0x09, 0xa0, 0xf7, 0xe1, 0x5e, 0xb9, 0xd1,
	0xc7, 0x1b, 0x9d, 0xeb, 0xa0, 0x62, 0x8d, 0x18, 0x93, 0x55, 0x0f, 0x46, 0xe2, 0x3a, 0x58, 0x1b,
	0x74, 0x5e, 0x7d, 0xd0, 0x0d, 0x7f, 0x82, 0x0d, 0x55, 0xd1, 0x72, 0xaf, 0xe5, 0xa5, 0x31, 0xa2,
	0x08, 0xcf, 0x69, 0x1d, 0xac, 0x9f, 0xd7, 0x81, 0xc1, 0xef, 0x0e, 0xc0, 0xf1, 0x4c, 0x9e, 0xd9,
	0xbf, 0x28, 0x73, 0x7a, 0x39, 0x37, 0xf4, 0x7a, 0x00, 0xad, 0x98, 0xf0, 0xe2, 0xb5, 0xe8, 0xcb,
	0x12, 0x13, 0xae, 0x07, 0xe3, 0x77, 0xd0, 0xe6, 0x3a, 0xd3, 0xdd, 0x15, 0xfe, 0x6b, 0x2a, 0xdc,
	0xb0, 0xe1, 0xea, 0xf0, 0x00, 0x78, 0xb9, 0x7e, 0xf6, 0x04, 0x3e, 0x8a, 0x58, 0xea, 0x8f, 0x19,
	0x1b, 0x27, 0xc4, 0x8f, 0xc9, 0x54, 0x8d, 0xc2, 0xdc, 0xd7, 0x77, 0xd8, 0x4f, 0xe8, 0xc8, 0xb7,
	0x7f, 0x1e, 0x7d, 0xfd, 0x57, 0xf2, 0xd8, 0xf9, 0x3f, 0x00, 0x00, 0xff, 0xff, 0xbe, 0x5f, 0x76,
	0x0c, 0x6d, 0x0a, 0x00, 0x00,
}
