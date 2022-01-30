// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: waypoint/builtin/nomad/jobspec/plugin.proto

package jobspec

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	anypb "google.golang.org/protobuf/types/known/anypb"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type Deployment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id            string     `protobuf:"bytes,1,opt,name=id,proto3" json:"id,omitempty"`
	Name          string     `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	ResourceState *anypb.Any `protobuf:"bytes,3,opt,name=resource_state,json=resourceState,proto3" json:"resource_state,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment.ProtoReflect.Descriptor instead.
func (*Deployment) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Deployment) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Deployment) GetResourceState() *anypb.Any {
	if x != nil {
		return x.ResourceState
	}
	return nil
}

type Release struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Url           string     `protobuf:"bytes,1,opt,name=Url,proto3" json:"Url,omitempty"`
	Id            string     `protobuf:"bytes,2,opt,name=id,proto3" json:"id,omitempty"`
	Name          string     `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
	ResourceState *anypb.Any `protobuf:"bytes,4,opt,name=resource_state,json=resourceState,proto3" json:"resource_state,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Release.ProtoReflect.Descriptor instead.
func (*Release) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *Release) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Release) GetId() string {
	if x != nil {
		return x.Id
	}
	return ""
}

func (x *Release) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Release) GetResourceState() *anypb.Any {
	if x != nil {
		return x.ResourceState
	}
	return nil
}

// Resource contains the internal resource states.
type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource.ProtoReflect.Descriptor instead.
func (*Resource) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescGZIP(), []int{2}
}

type Resource_Job struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Resource_Job) Reset() {
	*x = Resource_Job{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource_Job) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource_Job) ProtoMessage() {}

func (x *Resource_Job) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource_Job.ProtoReflect.Descriptor instead.
func (*Resource_Job) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Resource_Job) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

var File_waypoint_builtin_nomad_jobspec_plugin_proto protoreflect.FileDescriptor

var file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDesc = []byte{
	0x0a, 0x2b, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x74,
	0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x2f, 0x6a, 0x6f, 0x62, 0x73, 0x70, 0x65, 0x63,
	0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x07, 0x6a,
	0x6f, 0x62, 0x73, 0x70, 0x65, 0x63, 0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x22, 0x6d, 0x0a, 0x0a, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12,
	0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12,
	0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e,
	0x61, 0x6d, 0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f,
	0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f,
	0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e,
	0x79, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65,
	0x22, 0x7c, 0x0a, 0x07, 0x52, 0x65, 0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x55,
	0x72, 0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x55, 0x72, 0x6c, 0x12, 0x0e, 0x0a,
	0x02, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d,
	0x65, 0x12, 0x3b, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x5f, 0x73, 0x74,
	0x61, 0x74, 0x65, 0x18, 0x04, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67,
	0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75, 0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52,
	0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53, 0x74, 0x61, 0x74, 0x65, 0x22, 0x25,
	0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x1a, 0x19, 0x0a, 0x03, 0x4a, 0x6f,
	0x62, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x04, 0x6e, 0x61, 0x6d, 0x65, 0x42, 0x20, 0x5a, 0x1e, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e,
	0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x74, 0x69, 0x6e, 0x2f, 0x6e, 0x6f, 0x6d, 0x61, 0x64, 0x2f,
	0x6a, 0x6f, 0x62, 0x73, 0x70, 0x65, 0x63, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescOnce sync.Once
	file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescData = file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDesc
)

func file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescGZIP() []byte {
	file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescOnce.Do(func() {
		file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescData)
	})
	return file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDescData
}

var file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 4)
var file_waypoint_builtin_nomad_jobspec_plugin_proto_goTypes = []interface{}{
	(*Deployment)(nil),   // 0: jobspec.Deployment
	(*Release)(nil),      // 1: jobspec.Release
	(*Resource)(nil),     // 2: jobspec.Resource
	(*Resource_Job)(nil), // 3: jobspec.Resource.Job
	(*anypb.Any)(nil),    // 4: google.protobuf.Any
}
var file_waypoint_builtin_nomad_jobspec_plugin_proto_depIdxs = []int32{
	4, // 0: jobspec.Deployment.resource_state:type_name -> google.protobuf.Any
	4, // 1: jobspec.Release.resource_state:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_waypoint_builtin_nomad_jobspec_plugin_proto_init() }
func file_waypoint_builtin_nomad_jobspec_plugin_proto_init() {
	if File_waypoint_builtin_nomad_jobspec_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Release); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource_Job); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   4,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waypoint_builtin_nomad_jobspec_plugin_proto_goTypes,
		DependencyIndexes: file_waypoint_builtin_nomad_jobspec_plugin_proto_depIdxs,
		MessageInfos:      file_waypoint_builtin_nomad_jobspec_plugin_proto_msgTypes,
	}.Build()
	File_waypoint_builtin_nomad_jobspec_plugin_proto = out.File
	file_waypoint_builtin_nomad_jobspec_plugin_proto_rawDesc = nil
	file_waypoint_builtin_nomad_jobspec_plugin_proto_goTypes = nil
	file_waypoint_builtin_nomad_jobspec_plugin_proto_depIdxs = nil
}
