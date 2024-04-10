// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.26.0
// 	protoc        v3.19.1
// source: flyt/preparation-time-requested.proto

package flyt

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type PreparationTimeRequested struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Define your proto fields here
	Version int32 `protobuf:"varint,1,opt,name=version,proto3" json:"version,omitempty"`
	// Used the wrong ident type here, annoyingly.
	//
	// Deprecated: Do not use.
	Restaurant *RestaurantIdent                     `protobuf:"bytes,2,opt,name=restaurant,proto3" json:"restaurant,omitempty"`
	ItemCount  int32                                `protobuf:"varint,3,opt,name=item_count,json=itemCount,proto3" json:"item_count,omitempty"`
	Items      []*PreparationTimeRequested_MenuItem `protobuf:"bytes,4,rep,name=items,proto3" json:"items,omitempty"`
	// Add correct restaurant identifier field
	RestaurantId *Ident `protobuf:"bytes,5,opt,name=restaurant_id,json=restaurantId,proto3" json:"restaurant_id,omitempty"`
	// Add external order ID reference
	ExternalReference string `protobuf:"bytes,6,opt,name=external_reference,json=externalReference,proto3" json:"external_reference,omitempty"`
}

func (x *PreparationTimeRequested) Reset() {
	*x = PreparationTimeRequested{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_preparation_time_requested_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparationTimeRequested) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparationTimeRequested) ProtoMessage() {}

func (x *PreparationTimeRequested) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_preparation_time_requested_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparationTimeRequested.ProtoReflect.Descriptor instead.
func (*PreparationTimeRequested) Descriptor() ([]byte, []int) {
	return file_flyt_preparation_time_requested_proto_rawDescGZIP(), []int{0}
}

func (x *PreparationTimeRequested) GetVersion() int32 {
	if x != nil {
		return x.Version
	}
	return 0
}

// Deprecated: Do not use.
func (x *PreparationTimeRequested) GetRestaurant() *RestaurantIdent {
	if x != nil {
		return x.Restaurant
	}
	return nil
}

func (x *PreparationTimeRequested) GetItemCount() int32 {
	if x != nil {
		return x.ItemCount
	}
	return 0
}

func (x *PreparationTimeRequested) GetItems() []*PreparationTimeRequested_MenuItem {
	if x != nil {
		return x.Items
	}
	return nil
}

func (x *PreparationTimeRequested) GetRestaurantId() *Ident {
	if x != nil {
		return x.RestaurantId
	}
	return nil
}

func (x *PreparationTimeRequested) GetExternalReference() string {
	if x != nil {
		return x.ExternalReference
	}
	return ""
}

// Send an array of chosen Menu Items and their associated
// UUIDs and the UUIDs of any nested pick options & modifiers.
// This follows the same idea as basket totals, but is a seperate
// call flow from Skip. So doesn't logically sit in quite the same domain.
type PreparationTimeRequested_MenuItem struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Uuid        string                               `protobuf:"bytes,1,opt,name=uuid,proto3" json:"uuid,omitempty"`
	PickOptions []*PreparationTimeRequested_MenuItem `protobuf:"bytes,2,rep,name=pick_options,json=pickOptions,proto3" json:"pick_options,omitempty"`
	Modifiers   []*PreparationTimeRequested_MenuItem `protobuf:"bytes,3,rep,name=modifiers,proto3" json:"modifiers,omitempty"`
}

func (x *PreparationTimeRequested_MenuItem) Reset() {
	*x = PreparationTimeRequested_MenuItem{}
	if protoimpl.UnsafeEnabled {
		mi := &file_flyt_preparation_time_requested_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *PreparationTimeRequested_MenuItem) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*PreparationTimeRequested_MenuItem) ProtoMessage() {}

func (x *PreparationTimeRequested_MenuItem) ProtoReflect() protoreflect.Message {
	mi := &file_flyt_preparation_time_requested_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use PreparationTimeRequested_MenuItem.ProtoReflect.Descriptor instead.
func (*PreparationTimeRequested_MenuItem) Descriptor() ([]byte, []int) {
	return file_flyt_preparation_time_requested_proto_rawDescGZIP(), []int{0, 0}
}

func (x *PreparationTimeRequested_MenuItem) GetUuid() string {
	if x != nil {
		return x.Uuid
	}
	return ""
}

func (x *PreparationTimeRequested_MenuItem) GetPickOptions() []*PreparationTimeRequested_MenuItem {
	if x != nil {
		return x.PickOptions
	}
	return nil
}

func (x *PreparationTimeRequested_MenuItem) GetModifiers() []*PreparationTimeRequested_MenuItem {
	if x != nil {
		return x.Modifiers
	}
	return nil
}

var File_flyt_preparation_time_requested_proto protoreflect.FileDescriptor

var file_flyt_preparation_time_requested_proto_rawDesc = []byte{
	0x0a, 0x25, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69,
	0x6f, 0x6e, 0x2d, 0x74, 0x69, 0x6d, 0x65, 0x2d, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65,
	0x64, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x04, 0x66, 0x6c, 0x79, 0x74, 0x1a, 0x15, 0x66,
	0x6c, 0x79, 0x74, 0x2f, 0x64, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x6f, 0x72, 0x2e, 0x70,
	0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x10, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x69, 0x64, 0x65, 0x6e, 0x74,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x1a, 0x15, 0x66, 0x6c, 0x79, 0x74, 0x2f, 0x72, 0x65, 0x73,
	0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x82, 0x04,
	0x0a, 0x18, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d,
	0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x39, 0x0a, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61,
	0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e,
	0x52, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x65, 0x6e, 0x74, 0x42,
	0x02, 0x18, 0x01, 0x52, 0x0a, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x12,
	0x1d, 0x0a, 0x0a, 0x69, 0x74, 0x65, 0x6d, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x05, 0x52, 0x09, 0x69, 0x74, 0x65, 0x6d, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x3d,
	0x0a, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x18, 0x04, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e,
	0x66, 0x6c, 0x79, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e,
	0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x4d, 0x65,
	0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x05, 0x69, 0x74, 0x65, 0x6d, 0x73, 0x12, 0x30, 0x0a,
	0x0d, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x05,
	0x20, 0x01, 0x28, 0x0b, 0x32, 0x0b, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e, 0x49, 0x64, 0x65, 0x6e,
	0x74, 0x52, 0x0c, 0x72, 0x65, 0x73, 0x74, 0x61, 0x75, 0x72, 0x61, 0x6e, 0x74, 0x49, 0x64, 0x12,
	0x2d, 0x0a, 0x12, 0x65, 0x78, 0x74, 0x65, 0x72, 0x6e, 0x61, 0x6c, 0x5f, 0x72, 0x65, 0x66, 0x65,
	0x72, 0x65, 0x6e, 0x63, 0x65, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x11, 0x65, 0x78, 0x74,
	0x65, 0x72, 0x6e, 0x61, 0x6c, 0x52, 0x65, 0x66, 0x65, 0x72, 0x65, 0x6e, 0x63, 0x65, 0x1a, 0xb1,
	0x01, 0x0a, 0x08, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x12, 0x12, 0x0a, 0x04, 0x75,
	0x75, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x75, 0x75, 0x69, 0x64, 0x12,
	0x4a, 0x0a, 0x0c, 0x70, 0x69, 0x63, 0x6b, 0x5f, 0x6f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18,
	0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e, 0x50, 0x72, 0x65,
	0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75,
	0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x4d, 0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x0b,
	0x70, 0x69, 0x63, 0x6b, 0x4f, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x45, 0x0a, 0x09, 0x6d,
	0x6f, 0x64, 0x69, 0x66, 0x69, 0x65, 0x72, 0x73, 0x18, 0x03, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x27,
	0x2e, 0x66, 0x6c, 0x79, 0x74, 0x2e, 0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f,
	0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x2e, 0x4d,
	0x65, 0x6e, 0x75, 0x49, 0x74, 0x65, 0x6d, 0x52, 0x09, 0x6d, 0x6f, 0x64, 0x69, 0x66, 0x69, 0x65,
	0x72, 0x73, 0x3a, 0x1e, 0x82, 0xb5, 0x18, 0x1a, 0x70, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74,
	0x69, 0x6f, 0x6e, 0x2e, 0x74, 0x69, 0x6d, 0x65, 0x2e, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74,
	0x65, 0x64, 0x42, 0x7c, 0x0a, 0x08, 0x63, 0x6f, 0x6d, 0x2e, 0x66, 0x6c, 0x79, 0x74, 0x42, 0x1d,
	0x50, 0x72, 0x65, 0x70, 0x61, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x54, 0x69, 0x6d, 0x65, 0x52,
	0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x65, 0x64, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a,
	0x21, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x66, 0x6c, 0x79, 0x70,
	0x61, 0x79, 0x2f, 0x65, 0x76, 0x65, 0x6e, 0x74, 0x73, 0x2f, 0x70, 0x6b, 0x67, 0x2f, 0x66, 0x6c,
	0x79, 0x74, 0xa2, 0x02, 0x03, 0x46, 0x58, 0x58, 0xaa, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xca,
	0x02, 0x04, 0x46, 0x6c, 0x79, 0x74, 0xe2, 0x02, 0x10, 0x46, 0x6c, 0x79, 0x74, 0x5c, 0x47, 0x50,
	0x42, 0x4d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x04, 0x46, 0x6c, 0x79, 0x74,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_flyt_preparation_time_requested_proto_rawDescOnce sync.Once
	file_flyt_preparation_time_requested_proto_rawDescData = file_flyt_preparation_time_requested_proto_rawDesc
)

func file_flyt_preparation_time_requested_proto_rawDescGZIP() []byte {
	file_flyt_preparation_time_requested_proto_rawDescOnce.Do(func() {
		file_flyt_preparation_time_requested_proto_rawDescData = protoimpl.X.CompressGZIP(file_flyt_preparation_time_requested_proto_rawDescData)
	})
	return file_flyt_preparation_time_requested_proto_rawDescData
}

var file_flyt_preparation_time_requested_proto_msgTypes = make([]protoimpl.MessageInfo, 2)
var file_flyt_preparation_time_requested_proto_goTypes = []interface{}{
	(*PreparationTimeRequested)(nil),          // 0: flyt.PreparationTimeRequested
	(*PreparationTimeRequested_MenuItem)(nil), // 1: flyt.PreparationTimeRequested.MenuItem
	(*RestaurantIdent)(nil),                   // 2: flyt.RestaurantIdent
	(*Ident)(nil),                             // 3: flyt.Ident
}
var file_flyt_preparation_time_requested_proto_depIdxs = []int32{
	2, // 0: flyt.PreparationTimeRequested.restaurant:type_name -> flyt.RestaurantIdent
	1, // 1: flyt.PreparationTimeRequested.items:type_name -> flyt.PreparationTimeRequested.MenuItem
	3, // 2: flyt.PreparationTimeRequested.restaurant_id:type_name -> flyt.Ident
	1, // 3: flyt.PreparationTimeRequested.MenuItem.pick_options:type_name -> flyt.PreparationTimeRequested.MenuItem
	1, // 4: flyt.PreparationTimeRequested.MenuItem.modifiers:type_name -> flyt.PreparationTimeRequested.MenuItem
	5, // [5:5] is the sub-list for method output_type
	5, // [5:5] is the sub-list for method input_type
	5, // [5:5] is the sub-list for extension type_name
	5, // [5:5] is the sub-list for extension extendee
	0, // [0:5] is the sub-list for field type_name
}

func init() { file_flyt_preparation_time_requested_proto_init() }
func file_flyt_preparation_time_requested_proto_init() {
	if File_flyt_preparation_time_requested_proto != nil {
		return
	}
	file_flyt_descriptor_proto_init()
	file_flyt_ident_proto_init()
	file_flyt_restaurant_proto_init()
	if !protoimpl.UnsafeEnabled {
		file_flyt_preparation_time_requested_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparationTimeRequested); i {
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
		file_flyt_preparation_time_requested_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*PreparationTimeRequested_MenuItem); i {
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
			RawDescriptor: file_flyt_preparation_time_requested_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   2,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_flyt_preparation_time_requested_proto_goTypes,
		DependencyIndexes: file_flyt_preparation_time_requested_proto_depIdxs,
		MessageInfos:      file_flyt_preparation_time_requested_proto_msgTypes,
	}.Build()
	File_flyt_preparation_time_requested_proto = out.File
	file_flyt_preparation_time_requested_proto_rawDesc = nil
	file_flyt_preparation_time_requested_proto_goTypes = nil
	file_flyt_preparation_time_requested_proto_depIdxs = nil
}
