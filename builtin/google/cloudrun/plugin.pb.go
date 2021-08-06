// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.15.8
// source: waypoint/builtin/google/cloudrun/plugin.proto

package cloudrun

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

	Url           string               `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
	Resource      *Deployment_Resource `protobuf:"bytes,2,opt,name=resource,proto3" json:"resource,omitempty"`
	RevisionId    string               `protobuf:"bytes,3,opt,name=revision_id,json=revisionId,proto3" json:"revision_id,omitempty"`
	Region        string               `protobuf:"bytes,4,opt,name=region,proto3" json:"region,omitempty"`
	Id            string               `protobuf:"bytes,5,opt,name=id,proto3" json:"id,omitempty"`
	ResourceState *anypb.Any           `protobuf:"bytes,6,opt,name=resource_state,json=resourceState,proto3" json:"resource_state,omitempty"`
}

func (x *Deployment) Reset() {
	*x = Deployment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment) ProtoMessage() {}

func (x *Deployment) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[0]
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
	return file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescGZIP(), []int{0}
}

func (x *Deployment) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

func (x *Deployment) GetResource() *Deployment_Resource {
	if x != nil {
		return x.Resource
	}
	return nil
}

func (x *Deployment) GetRevisionId() string {
	if x != nil {
		return x.RevisionId
	}
	return ""
}

func (x *Deployment) GetRegion() string {
	if x != nil {
		return x.Region
	}
	return ""
}

func (x *Deployment) GetId() string {
	if x != nil {
		return x.Id
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

	Url string `protobuf:"bytes,1,opt,name=url,proto3" json:"url,omitempty"`
}

func (x *Release) Reset() {
	*x = Release{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Release) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Release) ProtoMessage() {}

func (x *Release) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[1]
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
	return file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescGZIP(), []int{1}
}

func (x *Release) GetUrl() string {
	if x != nil {
		return x.Url
	}
	return ""
}

type Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *Resource) Reset() {
	*x = Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource) ProtoMessage() {}

func (x *Resource) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[2]
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
	return file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescGZIP(), []int{2}
}

type Deployment_Resource struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Location string `protobuf:"bytes,1,opt,name=location,proto3" json:"location,omitempty"`
	Project  string `protobuf:"bytes,2,opt,name=project,proto3" json:"project,omitempty"`
	Name     string `protobuf:"bytes,3,opt,name=name,proto3" json:"name,omitempty"`
}

func (x *Deployment_Resource) Reset() {
	*x = Deployment_Resource{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Deployment_Resource) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Deployment_Resource) ProtoMessage() {}

func (x *Deployment_Resource) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Deployment_Resource.ProtoReflect.Descriptor instead.
func (*Deployment_Resource) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescGZIP(), []int{0, 0}
}

func (x *Deployment_Resource) GetLocation() string {
	if x != nil {
		return x.Location
	}
	return ""
}

func (x *Deployment_Resource) GetProject() string {
	if x != nil {
		return x.Project
	}
	return ""
}

func (x *Deployment_Resource) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

type Resource_Service struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ApiName      string `protobuf:"bytes,1,opt,name=api_name,json=apiName,proto3" json:"api_name,omitempty"`
	RevisionName string `protobuf:"bytes,2,opt,name=revision_name,json=revisionName,proto3" json:"revision_name,omitempty"`
}

func (x *Resource_Service) Reset() {
	*x = Resource_Service{}
	if protoimpl.UnsafeEnabled {
		mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Resource_Service) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Resource_Service) ProtoMessage() {}

func (x *Resource_Service) ProtoReflect() protoreflect.Message {
	mi := &file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Resource_Service.ProtoReflect.Descriptor instead.
func (*Resource_Service) Descriptor() ([]byte, []int) {
	return file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescGZIP(), []int{2, 0}
}

func (x *Resource_Service) GetApiName() string {
	if x != nil {
		return x.ApiName
	}
	return ""
}

func (x *Resource_Service) GetRevisionName() string {
	if x != nil {
		return x.RevisionName
	}
	return ""
}

var File_waypoint_builtin_google_cloudrun_plugin_proto protoreflect.FileDescriptor

var file_waypoint_builtin_google_cloudrun_plugin_proto_rawDesc = []byte{
	0x0a, 0x2d, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x74,
	0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72,
	0x75, 0x6e, 0x2f, 0x70, 0x6c, 0x75, 0x67, 0x69, 0x6e, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12,
	0x0f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x75, 0x6e,
	0x1a, 0x19, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2f, 0x61, 0x6e, 0x79, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0xbc, 0x02, 0x0a, 0x0a,
	0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72,
	0x6c, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x12, 0x40, 0x0a, 0x08,
	0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x24,
	0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72, 0x75, 0x6e,
	0x2e, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x52, 0x65, 0x73, 0x6f,
	0x75, 0x72, 0x63, 0x65, 0x52, 0x08, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x12, 0x1f,
	0x0a, 0x0b, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x69, 0x64, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x0a, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x49, 0x64, 0x12,
	0x16, 0x0a, 0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52,
	0x06, 0x72, 0x65, 0x67, 0x69, 0x6f, 0x6e, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x05, 0x20,
	0x01, 0x28, 0x09, 0x52, 0x02, 0x69, 0x64, 0x12, 0x3b, 0x0a, 0x0e, 0x72, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x5f, 0x73, 0x74, 0x61, 0x74, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0b, 0x32,
	0x14, 0x2e, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x75,
	0x66, 0x2e, 0x41, 0x6e, 0x79, 0x52, 0x0d, 0x72, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65, 0x53,
	0x74, 0x61, 0x74, 0x65, 0x1a, 0x54, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75, 0x72, 0x63, 0x65,
	0x12, 0x1a, 0x0a, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x08, 0x6c, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x18, 0x0a, 0x07,
	0x70, 0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x70,
	0x72, 0x6f, 0x6a, 0x65, 0x63, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x22, 0x1b, 0x0a, 0x07, 0x52, 0x65,
	0x6c, 0x65, 0x61, 0x73, 0x65, 0x12, 0x10, 0x0a, 0x03, 0x75, 0x72, 0x6c, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x75, 0x72, 0x6c, 0x22, 0x55, 0x0a, 0x08, 0x52, 0x65, 0x73, 0x6f, 0x75,
	0x72, 0x63, 0x65, 0x1a, 0x49, 0x0a, 0x07, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x12, 0x19,
	0x0a, 0x08, 0x61, 0x70, 0x69, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x07, 0x61, 0x70, 0x69, 0x4e, 0x61, 0x6d, 0x65, 0x12, 0x23, 0x0a, 0x0d, 0x72, 0x65, 0x76,
	0x69, 0x73, 0x69, 0x6f, 0x6e, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09,
	0x52, 0x0c, 0x72, 0x65, 0x76, 0x69, 0x73, 0x69, 0x6f, 0x6e, 0x4e, 0x61, 0x6d, 0x65, 0x42, 0x22,
	0x5a, 0x20, 0x77, 0x61, 0x79, 0x70, 0x6f, 0x69, 0x6e, 0x74, 0x2f, 0x62, 0x75, 0x69, 0x6c, 0x74,
	0x69, 0x6e, 0x2f, 0x67, 0x6f, 0x6f, 0x67, 0x6c, 0x65, 0x2f, 0x63, 0x6c, 0x6f, 0x75, 0x64, 0x72,
	0x75, 0x6e, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescOnce sync.Once
	file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescData = file_waypoint_builtin_google_cloudrun_plugin_proto_rawDesc
)

func file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescGZIP() []byte {
	file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescOnce.Do(func() {
		file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescData = protoimpl.X.CompressGZIP(file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescData)
	})
	return file_waypoint_builtin_google_cloudrun_plugin_proto_rawDescData
}

var file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes = make([]protoimpl.MessageInfo, 5)
var file_waypoint_builtin_google_cloudrun_plugin_proto_goTypes = []interface{}{
	(*Deployment)(nil),          // 0: google.cloudrun.Deployment
	(*Release)(nil),             // 1: google.cloudrun.Release
	(*Resource)(nil),            // 2: google.cloudrun.Resource
	(*Deployment_Resource)(nil), // 3: google.cloudrun.Deployment.Resource
	(*Resource_Service)(nil),    // 4: google.cloudrun.Resource.Service
	(*anypb.Any)(nil),           // 5: google.protobuf.Any
}
var file_waypoint_builtin_google_cloudrun_plugin_proto_depIdxs = []int32{
	3, // 0: google.cloudrun.Deployment.resource:type_name -> google.cloudrun.Deployment.Resource
	5, // 1: google.cloudrun.Deployment.resource_state:type_name -> google.protobuf.Any
	2, // [2:2] is the sub-list for method output_type
	2, // [2:2] is the sub-list for method input_type
	2, // [2:2] is the sub-list for extension type_name
	2, // [2:2] is the sub-list for extension extendee
	0, // [0:2] is the sub-list for field type_name
}

func init() { file_waypoint_builtin_google_cloudrun_plugin_proto_init() }
func file_waypoint_builtin_google_cloudrun_plugin_proto_init() {
	if File_waypoint_builtin_google_cloudrun_plugin_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
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
		file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
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
		file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
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
		file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Deployment_Resource); i {
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
		file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Resource_Service); i {
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
			RawDescriptor: file_waypoint_builtin_google_cloudrun_plugin_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   5,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_waypoint_builtin_google_cloudrun_plugin_proto_goTypes,
		DependencyIndexes: file_waypoint_builtin_google_cloudrun_plugin_proto_depIdxs,
		MessageInfos:      file_waypoint_builtin_google_cloudrun_plugin_proto_msgTypes,
	}.Build()
	File_waypoint_builtin_google_cloudrun_plugin_proto = out.File
	file_waypoint_builtin_google_cloudrun_plugin_proto_rawDesc = nil
	file_waypoint_builtin_google_cloudrun_plugin_proto_goTypes = nil
	file_waypoint_builtin_google_cloudrun_plugin_proto_depIdxs = nil
}
