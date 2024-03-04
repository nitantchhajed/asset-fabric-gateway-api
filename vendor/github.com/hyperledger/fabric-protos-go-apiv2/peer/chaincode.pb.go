// Copyright the Hyperledger Fabric contributors. All rights reserved.
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.0
// 	protoc        (unknown)
// source: peer/chaincode.proto

package peer

import (
	common "github.com/hyperledger/fabric-protos-go-apiv2/common"
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

type ChaincodeSpec_Type int32

const (
	ChaincodeSpec_UNDEFINED ChaincodeSpec_Type = 0
	ChaincodeSpec_GOLANG    ChaincodeSpec_Type = 1
	ChaincodeSpec_NODE      ChaincodeSpec_Type = 2
	ChaincodeSpec_CAR       ChaincodeSpec_Type = 3
	ChaincodeSpec_JAVA      ChaincodeSpec_Type = 4
)

// Enum value maps for ChaincodeSpec_Type.
var (
	ChaincodeSpec_Type_name = map[int32]string{
		0: "UNDEFINED",
		1: "GOLANG",
		2: "NODE",
		3: "CAR",
		4: "JAVA",
	}
	ChaincodeSpec_Type_value = map[string]int32{
		"UNDEFINED": 0,
		"GOLANG":    1,
		"NODE":      2,
		"CAR":       3,
		"JAVA":      4,
	}
)

func (x ChaincodeSpec_Type) Enum() *ChaincodeSpec_Type {
	p := new(ChaincodeSpec_Type)
	*p = x
	return p
}

func (x ChaincodeSpec_Type) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ChaincodeSpec_Type) Descriptor() protoreflect.EnumDescriptor {
	return file_peer_chaincode_proto_enumTypes[0].Descriptor()
}

func (ChaincodeSpec_Type) Type() protoreflect.EnumType {
	return &file_peer_chaincode_proto_enumTypes[0]
}

func (x ChaincodeSpec_Type) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ChaincodeSpec_Type.Descriptor instead.
func (ChaincodeSpec_Type) EnumDescriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{2, 0}
}

//ChaincodeID contains the path as specified by the deploy transaction
//that created it as well as the hashCode that is generated by the
//system for the path. From the user level (ie, CLI, REST API and so on)
//deploy transaction is expected to provide the path and other requests
//are expected to provide the hashCode. The other value will be ignored.
//Internally, the structure could contain both values. For instance, the
//hashCode will be set when first generated using the path
type ChaincodeID struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	//deploy transaction will use the path
	Path string `protobuf:"bytes,1,opt,name=path,proto3" json:"path,omitempty"`
	//all other requests will use the name (really a hashcode) generated by
	//the deploy transaction
	Name string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`
	//user friendly version name for the chaincode
	Version string `protobuf:"bytes,3,opt,name=version,proto3" json:"version,omitempty"`
}

func (x *ChaincodeID) Reset() {
	*x = ChaincodeID{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeID) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeID) ProtoMessage() {}

func (x *ChaincodeID) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeID.ProtoReflect.Descriptor instead.
func (*ChaincodeID) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{0}
}

func (x *ChaincodeID) GetPath() string {
	if x != nil {
		return x.Path
	}
	return ""
}

func (x *ChaincodeID) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChaincodeID) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

// Carries the chaincode function and its arguments.
// UnmarshalJSON in transaction.go converts the string-based REST/JSON input to
// the []byte-based current ChaincodeInput structure.
type ChaincodeInput struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Args        [][]byte          `protobuf:"bytes,1,rep,name=args,proto3" json:"args,omitempty"`
	Decorations map[string][]byte `protobuf:"bytes,2,rep,name=decorations,proto3" json:"decorations,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	// is_init is used for the application to signal that an invocation is to be routed
	// to the legacy 'Init' function for compatibility with chaincodes which handled
	// Init in the old way.  New applications should manage their initialized state
	// themselves.
	IsInit bool `protobuf:"varint,3,opt,name=is_init,json=isInit,proto3" json:"is_init,omitempty"`
}

func (x *ChaincodeInput) Reset() {
	*x = ChaincodeInput{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInput) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInput) ProtoMessage() {}

func (x *ChaincodeInput) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInput.ProtoReflect.Descriptor instead.
func (*ChaincodeInput) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{1}
}

func (x *ChaincodeInput) GetArgs() [][]byte {
	if x != nil {
		return x.Args
	}
	return nil
}

func (x *ChaincodeInput) GetDecorations() map[string][]byte {
	if x != nil {
		return x.Decorations
	}
	return nil
}

func (x *ChaincodeInput) GetIsInit() bool {
	if x != nil {
		return x.IsInit
	}
	return false
}

// Carries the chaincode specification. This is the actual metadata required for
// defining a chaincode.
type ChaincodeSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Type        ChaincodeSpec_Type `protobuf:"varint,1,opt,name=type,proto3,enum=protos.ChaincodeSpec_Type" json:"type,omitempty"`
	ChaincodeId *ChaincodeID       `protobuf:"bytes,2,opt,name=chaincode_id,json=chaincodeId,proto3" json:"chaincode_id,omitempty"`
	Input       *ChaincodeInput    `protobuf:"bytes,3,opt,name=input,proto3" json:"input,omitempty"`
	Timeout     int32              `protobuf:"varint,4,opt,name=timeout,proto3" json:"timeout,omitempty"`
}

func (x *ChaincodeSpec) Reset() {
	*x = ChaincodeSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeSpec) ProtoMessage() {}

func (x *ChaincodeSpec) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeSpec.ProtoReflect.Descriptor instead.
func (*ChaincodeSpec) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{2}
}

func (x *ChaincodeSpec) GetType() ChaincodeSpec_Type {
	if x != nil {
		return x.Type
	}
	return ChaincodeSpec_UNDEFINED
}

func (x *ChaincodeSpec) GetChaincodeId() *ChaincodeID {
	if x != nil {
		return x.ChaincodeId
	}
	return nil
}

func (x *ChaincodeSpec) GetInput() *ChaincodeInput {
	if x != nil {
		return x.Input
	}
	return nil
}

func (x *ChaincodeSpec) GetTimeout() int32 {
	if x != nil {
		return x.Timeout
	}
	return 0
}

// Specify the deployment of a chaincode.
// TODO: Define `codePackage`.
type ChaincodeDeploymentSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec,proto3" json:"chaincode_spec,omitempty"`
	CodePackage   []byte         `protobuf:"bytes,3,opt,name=code_package,json=codePackage,proto3" json:"code_package,omitempty"`
}

func (x *ChaincodeDeploymentSpec) Reset() {
	*x = ChaincodeDeploymentSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeDeploymentSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeDeploymentSpec) ProtoMessage() {}

func (x *ChaincodeDeploymentSpec) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeDeploymentSpec.ProtoReflect.Descriptor instead.
func (*ChaincodeDeploymentSpec) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{3}
}

func (x *ChaincodeDeploymentSpec) GetChaincodeSpec() *ChaincodeSpec {
	if x != nil {
		return x.ChaincodeSpec
	}
	return nil
}

func (x *ChaincodeDeploymentSpec) GetCodePackage() []byte {
	if x != nil {
		return x.CodePackage
	}
	return nil
}

// Carries the chaincode function and its arguments.
type ChaincodeInvocationSpec struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChaincodeSpec *ChaincodeSpec `protobuf:"bytes,1,opt,name=chaincode_spec,json=chaincodeSpec,proto3" json:"chaincode_spec,omitempty"`
}

func (x *ChaincodeInvocationSpec) Reset() {
	*x = ChaincodeInvocationSpec{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeInvocationSpec) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeInvocationSpec) ProtoMessage() {}

func (x *ChaincodeInvocationSpec) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeInvocationSpec.ProtoReflect.Descriptor instead.
func (*ChaincodeInvocationSpec) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{4}
}

func (x *ChaincodeInvocationSpec) GetChaincodeSpec() *ChaincodeSpec {
	if x != nil {
		return x.ChaincodeSpec
	}
	return nil
}

// LifecycleEvent is used as the payload of the chaincode event emitted by LSCC
type LifecycleEvent struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	ChaincodeName string `protobuf:"bytes,1,opt,name=chaincode_name,json=chaincodeName,proto3" json:"chaincode_name,omitempty"`
}

func (x *LifecycleEvent) Reset() {
	*x = LifecycleEvent{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *LifecycleEvent) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*LifecycleEvent) ProtoMessage() {}

func (x *LifecycleEvent) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use LifecycleEvent.ProtoReflect.Descriptor instead.
func (*LifecycleEvent) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{5}
}

func (x *LifecycleEvent) GetChaincodeName() string {
	if x != nil {
		return x.ChaincodeName
	}
	return ""
}

// CDSData is data stored in the LSCC on instantiation of a CC
// for CDSPackage.  This needs to be serialized for ChaincodeData
// hence the protobuf format
type CDSData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash         []byte `protobuf:"bytes,1,opt,name=hash,proto3" json:"hash,omitempty"`                 // hash of ChaincodeDeploymentSpec.code_package
	Metadatahash []byte `protobuf:"bytes,2,opt,name=metadatahash,proto3" json:"metadatahash,omitempty"` // hash of ChaincodeID.name + ChaincodeID.version
}

func (x *CDSData) Reset() {
	*x = CDSData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CDSData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CDSData) ProtoMessage() {}

func (x *CDSData) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CDSData.ProtoReflect.Descriptor instead.
func (*CDSData) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{6}
}

func (x *CDSData) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *CDSData) GetMetadatahash() []byte {
	if x != nil {
		return x.Metadatahash
	}
	return nil
}

// ChaincodeData defines the datastructure for chaincodes to be serialized by proto
// Type provides an additional check by directing to use a specific package after instantiation
// Data is Type specific (see CDSPackage and SignedCDSPackage)
type ChaincodeData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name of the chaincode
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Version of the chaincode
	Version string `protobuf:"bytes,2,opt,name=version,proto3" json:"version,omitempty"`
	// Escc for the chaincode instance
	Escc string `protobuf:"bytes,3,opt,name=escc,proto3" json:"escc,omitempty"`
	// Vscc for the chaincode instance
	Vscc string `protobuf:"bytes,4,opt,name=vscc,proto3" json:"vscc,omitempty"`
	// Policy endorsement policy for the chaincode instance
	Policy *common.SignaturePolicyEnvelope `protobuf:"bytes,5,opt,name=policy,proto3" json:"policy,omitempty"`
	// Data data specific to the package
	Data []byte `protobuf:"bytes,6,opt,name=data,proto3" json:"data,omitempty"`
	// Id of the chaincode that's the unique fingerprint for the CC This is not
	// currently used anywhere but serves as a good eyecatcher
	Id []byte `protobuf:"bytes,7,opt,name=id,proto3" json:"id,omitempty"`
	// InstantiationPolicy for the chaincode
	InstantiationPolicy *common.SignaturePolicyEnvelope `protobuf:"bytes,8,opt,name=instantiation_policy,json=instantiationPolicy,proto3" json:"instantiation_policy,omitempty"`
}

func (x *ChaincodeData) Reset() {
	*x = ChaincodeData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_peer_chaincode_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *ChaincodeData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*ChaincodeData) ProtoMessage() {}

func (x *ChaincodeData) ProtoReflect() protoreflect.Message {
	mi := &file_peer_chaincode_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use ChaincodeData.ProtoReflect.Descriptor instead.
func (*ChaincodeData) Descriptor() ([]byte, []int) {
	return file_peer_chaincode_proto_rawDescGZIP(), []int{7}
}

func (x *ChaincodeData) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *ChaincodeData) GetVersion() string {
	if x != nil {
		return x.Version
	}
	return ""
}

func (x *ChaincodeData) GetEscc() string {
	if x != nil {
		return x.Escc
	}
	return ""
}

func (x *ChaincodeData) GetVscc() string {
	if x != nil {
		return x.Vscc
	}
	return ""
}

func (x *ChaincodeData) GetPolicy() *common.SignaturePolicyEnvelope {
	if x != nil {
		return x.Policy
	}
	return nil
}

func (x *ChaincodeData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *ChaincodeData) GetId() []byte {
	if x != nil {
		return x.Id
	}
	return nil
}

func (x *ChaincodeData) GetInstantiationPolicy() *common.SignaturePolicyEnvelope {
	if x != nil {
		return x.InstantiationPolicy
	}
	return nil
}

var File_peer_chaincode_proto protoreflect.FileDescriptor

var file_peer_chaincode_proto_rawDesc = []byte{
	0x0a, 0x14, 0x70, 0x65, 0x65, 0x72, 0x2f, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x12, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x1a, 0x15,
	0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2f, 0x70, 0x6f, 0x6c, 0x69, 0x63, 0x69, 0x65, 0x73, 0x2e,
	0x70, 0x72, 0x6f, 0x74, 0x6f, 0x22, 0x4f, 0x0a, 0x0b, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x49, 0x44, 0x12, 0x12, 0x0a, 0x04, 0x70, 0x61, 0x74, 0x68, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x70, 0x61, 0x74, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07,
	0x76, 0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76,
	0x65, 0x72, 0x73, 0x69, 0x6f, 0x6e, 0x22, 0xc8, 0x01, 0x0a, 0x0e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x12, 0x0a, 0x04, 0x61, 0x72, 0x67,
	0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x04, 0x61, 0x72, 0x67, 0x73, 0x12, 0x49, 0x0a,
	0x0b, 0x64, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x18, 0x02, 0x20, 0x03,
	0x28, 0x0b, 0x32, 0x27, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69,
	0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x2e, 0x44, 0x65, 0x63, 0x6f, 0x72,
	0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x0b, 0x64, 0x65, 0x63,
	0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73, 0x12, 0x17, 0x0a, 0x07, 0x69, 0x73, 0x5f, 0x69,
	0x6e, 0x69, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x08, 0x52, 0x06, 0x69, 0x73, 0x49, 0x6e, 0x69,
	0x74, 0x1a, 0x3e, 0x0a, 0x10, 0x44, 0x65, 0x63, 0x6f, 0x72, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x73,
	0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a, 0x02, 0x38,
	0x01, 0x22, 0xff, 0x01, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53,
	0x70, 0x65, 0x63, 0x12, 0x2e, 0x0a, 0x04, 0x74, 0x79, 0x70, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0e, 0x32, 0x1a, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e,
	0x63, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x2e, 0x54, 0x79, 0x70, 0x65, 0x52, 0x04, 0x74,
	0x79, 0x70, 0x65, 0x12, 0x36, 0x0a, 0x0c, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65,
	0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x13, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x44, 0x52, 0x0b,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x64, 0x12, 0x2c, 0x0a, 0x05, 0x69,
	0x6e, 0x70, 0x75, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x16, 0x2e, 0x70, 0x72, 0x6f,
	0x74, 0x6f, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x49, 0x6e, 0x70,
	0x75, 0x74, 0x52, 0x05, 0x69, 0x6e, 0x70, 0x75, 0x74, 0x12, 0x18, 0x0a, 0x07, 0x74, 0x69, 0x6d,
	0x65, 0x6f, 0x75, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x05, 0x52, 0x07, 0x74, 0x69, 0x6d, 0x65,
	0x6f, 0x75, 0x74, 0x22, 0x3e, 0x0a, 0x04, 0x54, 0x79, 0x70, 0x65, 0x12, 0x0d, 0x0a, 0x09, 0x55,
	0x4e, 0x44, 0x45, 0x46, 0x49, 0x4e, 0x45, 0x44, 0x10, 0x00, 0x12, 0x0a, 0x0a, 0x06, 0x47, 0x4f,
	0x4c, 0x41, 0x4e, 0x47, 0x10, 0x01, 0x12, 0x08, 0x0a, 0x04, 0x4e, 0x4f, 0x44, 0x45, 0x10, 0x02,
	0x12, 0x07, 0x0a, 0x03, 0x43, 0x41, 0x52, 0x10, 0x03, 0x12, 0x08, 0x0a, 0x04, 0x4a, 0x41, 0x56,
	0x41, 0x10, 0x04, 0x22, 0xa0, 0x01, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64,
	0x65, 0x44, 0x65, 0x70, 0x6c, 0x6f, 0x79, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x70, 0x65, 0x63, 0x12,
	0x3c, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73, 0x70, 0x65,
	0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x52, 0x0d,
	0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x12, 0x21, 0x0a,
	0x0c, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x70, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x0b, 0x63, 0x6f, 0x64, 0x65, 0x50, 0x61, 0x63, 0x6b, 0x61, 0x67, 0x65,
	0x4a, 0x04, 0x08, 0x02, 0x10, 0x03, 0x4a, 0x04, 0x08, 0x04, 0x10, 0x05, 0x52, 0x0e, 0x65, 0x66,
	0x66, 0x65, 0x63, 0x74, 0x69, 0x76, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x52, 0x08, 0x65, 0x78,
	0x65, 0x63, 0x5f, 0x65, 0x6e, 0x76, 0x22, 0x70, 0x0a, 0x17, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63,
	0x6f, 0x64, 0x65, 0x49, 0x6e, 0x76, 0x6f, 0x63, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x70, 0x65,
	0x63, 0x12, 0x3c, 0x0a, 0x0e, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x73,
	0x70, 0x65, 0x63, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x15, 0x2e, 0x70, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0x2e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63,
	0x52, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x53, 0x70, 0x65, 0x63, 0x4a,
	0x04, 0x08, 0x02, 0x10, 0x03, 0x52, 0x11, 0x69, 0x64, 0x5f, 0x67, 0x65, 0x6e, 0x65, 0x72, 0x61,
	0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x61, 0x6c, 0x67, 0x22, 0x37, 0x0a, 0x0e, 0x4c, 0x69, 0x66, 0x65,
	0x63, 0x79, 0x63, 0x6c, 0x65, 0x45, 0x76, 0x65, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x63, 0x68,
	0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x5f, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x0d, 0x63, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x4e, 0x61, 0x6d,
	0x65, 0x22, 0x41, 0x0a, 0x07, 0x43, 0x44, 0x53, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04,
	0x68, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x68, 0x61, 0x73, 0x68,
	0x12, 0x22, 0x0a, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0x68, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x6d, 0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61,
	0x68, 0x61, 0x73, 0x68, 0x22, 0x96, 0x02, 0x0a, 0x0d, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f,
	0x64, 0x65, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65, 0x12, 0x18, 0x0a, 0x07, 0x76, 0x65,
	0x72, 0x73, 0x69, 0x6f, 0x6e, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x76, 0x65, 0x72,
	0x73, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x65, 0x73, 0x63, 0x63, 0x18, 0x03, 0x20, 0x01,
	0x28, 0x09, 0x52, 0x04, 0x65, 0x73, 0x63, 0x63, 0x12, 0x12, 0x0a, 0x04, 0x76, 0x73, 0x63, 0x63,
	0x18, 0x04, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x76, 0x73, 0x63, 0x63, 0x12, 0x37, 0x0a, 0x06,
	0x70, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x6f, 0x6e, 0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x06, 0x70,
	0x6f, 0x6c, 0x69, 0x63, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x64, 0x61, 0x74, 0x61, 0x18, 0x06, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x64, 0x61, 0x74, 0x61, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18,
	0x07, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x02, 0x69, 0x64, 0x12, 0x52, 0x0a, 0x14, 0x69, 0x6e, 0x73,
	0x74, 0x61, 0x6e, 0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x5f, 0x70, 0x6f, 0x6c, 0x69, 0x63,
	0x79, 0x18, 0x08, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x6f, 0x6e,
	0x2e, 0x53, 0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79,
	0x45, 0x6e, 0x76, 0x65, 0x6c, 0x6f, 0x70, 0x65, 0x52, 0x13, 0x69, 0x6e, 0x73, 0x74, 0x61, 0x6e,
	0x74, 0x69, 0x61, 0x74, 0x69, 0x6f, 0x6e, 0x50, 0x6f, 0x6c, 0x69, 0x63, 0x79, 0x42, 0xa0, 0x01,
	0x0a, 0x22, 0x6f, 0x72, 0x67, 0x2e, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65,
	0x72, 0x2e, 0x66, 0x61, 0x62, 0x72, 0x69, 0x63, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e,
	0x70, 0x65, 0x65, 0x72, 0x42, 0x0e, 0x43, 0x68, 0x61, 0x69, 0x6e, 0x63, 0x6f, 0x64, 0x65, 0x50,
	0x72, 0x6f, 0x74, 0x6f, 0x50, 0x01, 0x5a, 0x32, 0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63,
	0x6f, 0x6d, 0x2f, 0x68, 0x79, 0x70, 0x65, 0x72, 0x6c, 0x65, 0x64, 0x67, 0x65, 0x72, 0x2f, 0x66,
	0x61, 0x62, 0x72, 0x69, 0x63, 0x2d, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2d, 0x67, 0x6f, 0x2d,
	0x61, 0x70, 0x69, 0x76, 0x32, 0x2f, 0x70, 0x65, 0x65, 0x72, 0xa2, 0x02, 0x03, 0x50, 0x58, 0x58,
	0xaa, 0x02, 0x06, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0xca, 0x02, 0x06, 0x50, 0x72, 0x6f, 0x74,
	0x6f, 0x73, 0xe2, 0x02, 0x12, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x5c, 0x47, 0x50, 0x42, 0x4d,
	0x65, 0x74, 0x61, 0x64, 0x61, 0x74, 0x61, 0xea, 0x02, 0x06, 0x50, 0x72, 0x6f, 0x74, 0x6f, 0x73,
	0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_peer_chaincode_proto_rawDescOnce sync.Once
	file_peer_chaincode_proto_rawDescData = file_peer_chaincode_proto_rawDesc
)

func file_peer_chaincode_proto_rawDescGZIP() []byte {
	file_peer_chaincode_proto_rawDescOnce.Do(func() {
		file_peer_chaincode_proto_rawDescData = protoimpl.X.CompressGZIP(file_peer_chaincode_proto_rawDescData)
	})
	return file_peer_chaincode_proto_rawDescData
}

var file_peer_chaincode_proto_enumTypes = make([]protoimpl.EnumInfo, 1)
var file_peer_chaincode_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_peer_chaincode_proto_goTypes = []interface{}{
	(ChaincodeSpec_Type)(0),                // 0: protos.ChaincodeSpec.Type
	(*ChaincodeID)(nil),                    // 1: protos.ChaincodeID
	(*ChaincodeInput)(nil),                 // 2: protos.ChaincodeInput
	(*ChaincodeSpec)(nil),                  // 3: protos.ChaincodeSpec
	(*ChaincodeDeploymentSpec)(nil),        // 4: protos.ChaincodeDeploymentSpec
	(*ChaincodeInvocationSpec)(nil),        // 5: protos.ChaincodeInvocationSpec
	(*LifecycleEvent)(nil),                 // 6: protos.LifecycleEvent
	(*CDSData)(nil),                        // 7: protos.CDSData
	(*ChaincodeData)(nil),                  // 8: protos.ChaincodeData
	nil,                                    // 9: protos.ChaincodeInput.DecorationsEntry
	(*common.SignaturePolicyEnvelope)(nil), // 10: common.SignaturePolicyEnvelope
}
var file_peer_chaincode_proto_depIdxs = []int32{
	9,  // 0: protos.ChaincodeInput.decorations:type_name -> protos.ChaincodeInput.DecorationsEntry
	0,  // 1: protos.ChaincodeSpec.type:type_name -> protos.ChaincodeSpec.Type
	1,  // 2: protos.ChaincodeSpec.chaincode_id:type_name -> protos.ChaincodeID
	2,  // 3: protos.ChaincodeSpec.input:type_name -> protos.ChaincodeInput
	3,  // 4: protos.ChaincodeDeploymentSpec.chaincode_spec:type_name -> protos.ChaincodeSpec
	3,  // 5: protos.ChaincodeInvocationSpec.chaincode_spec:type_name -> protos.ChaincodeSpec
	10, // 6: protos.ChaincodeData.policy:type_name -> common.SignaturePolicyEnvelope
	10, // 7: protos.ChaincodeData.instantiation_policy:type_name -> common.SignaturePolicyEnvelope
	8,  // [8:8] is the sub-list for method output_type
	8,  // [8:8] is the sub-list for method input_type
	8,  // [8:8] is the sub-list for extension type_name
	8,  // [8:8] is the sub-list for extension extendee
	0,  // [0:8] is the sub-list for field type_name
}

func init() { file_peer_chaincode_proto_init() }
func file_peer_chaincode_proto_init() {
	if File_peer_chaincode_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_peer_chaincode_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeID); i {
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
		file_peer_chaincode_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInput); i {
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
		file_peer_chaincode_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeSpec); i {
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
		file_peer_chaincode_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeDeploymentSpec); i {
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
		file_peer_chaincode_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeInvocationSpec); i {
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
		file_peer_chaincode_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*LifecycleEvent); i {
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
		file_peer_chaincode_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CDSData); i {
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
		file_peer_chaincode_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*ChaincodeData); i {
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
			RawDescriptor: file_peer_chaincode_proto_rawDesc,
			NumEnums:      1,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_peer_chaincode_proto_goTypes,
		DependencyIndexes: file_peer_chaincode_proto_depIdxs,
		EnumInfos:         file_peer_chaincode_proto_enumTypes,
		MessageInfos:      file_peer_chaincode_proto_msgTypes,
	}.Build()
	File_peer_chaincode_proto = out.File
	file_peer_chaincode_proto_rawDesc = nil
	file_peer_chaincode_proto_goTypes = nil
	file_peer_chaincode_proto_depIdxs = nil
}
