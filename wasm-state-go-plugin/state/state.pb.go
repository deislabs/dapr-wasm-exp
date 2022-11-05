// Code generated by protoc-gen-go-plugin. DO NOT EDIT.
// versions:
// 	protoc-gen-go-plugin v0.1.0
// 	protoc               v3.21.5
// source: state/state.proto

//import "google/protobuf/any.proto";

package state

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

// The request message containing the user's name.
type InitRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Metadata *Metadata `protobuf:"bytes,1,opt,name=metadata,proto3" json:"metadata,omitempty"`
}

func (x *InitRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *InitRequest) GetMetadata() *Metadata {
	if x != nil {
		return x.Metadata
	}
	return nil
}

type Metadata struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	// Name is the name of the component that's being used.
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	// Properties is the metadata properties.
	Properties map[string]string `protobuf:"bytes,2,rep,name=properties,proto3" json:"properties,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *Metadata) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *Metadata) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *Metadata) GetProperties() map[string]string {
	if x != nil {
		return x.Properties
	}
	return nil
}

// The response message for the init
type InitResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *InitResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

type FeaturesRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *FeaturesRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

type FeaturesResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Features []string `protobuf:"bytes,1,rep,name=features,proto3" json:"features,omitempty"`
}

func (x *FeaturesResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *FeaturesResponse) GetFeatures() []string {
	if x != nil {
		return x.Features
	}
	return nil
}

// GetRequest is the object describing a state fetch request.
type GetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Metadata map[string]string `protobuf:"bytes,2,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Options  *GetStateOption   `protobuf:"bytes,3,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *GetRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *GetRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GetRequest) GetOptions() *GetStateOption {
	if x != nil {
		return x.Options
	}
	return nil
}

// GetStateOption controls how a state store reacts to a get request.
type GetStateOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Consistency string `protobuf:"bytes,1,opt,name=consistency,proto3" json:"consistency,omitempty"`
}

func (x *GetStateOption) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetStateOption) GetConsistency() string {
	if x != nil {
		return x.Consistency
	}
	return ""
}

// GetResponse is the response object for getting state.
type GetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Data        []byte            `protobuf:"bytes,1,opt,name=data,proto3" json:"data,omitempty"`
	Etag        string            `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	ContentType string            `protobuf:"bytes,4,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *GetResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *GetResponse) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *GetResponse) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *GetResponse) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *GetResponse) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

// GetRequest is the object describing a state fetch request.
type SetRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key         string            `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Data        []byte            `protobuf:"bytes,2,opt,name=data,proto3" json:"data,omitempty"`
	Etag        string            `protobuf:"bytes,3,opt,name=etag,proto3" json:"etag,omitempty"`
	Metadata    map[string]string `protobuf:"bytes,4,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Options     *SetStateOption   `protobuf:"bytes,5,opt,name=options,proto3" json:"options,omitempty"`
	ContentType string            `protobuf:"bytes,6,opt,name=content_type,json=contentType,proto3" json:"content_type,omitempty"`
}

func (x *SetRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SetRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *SetRequest) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *SetRequest) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *SetRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *SetRequest) GetOptions() *SetStateOption {
	if x != nil {
		return x.Options
	}
	return nil
}

func (x *SetRequest) GetContentType() string {
	if x != nil {
		return x.ContentType
	}
	return ""
}

// SetStateOption controls how a state store reacts to a set request.
type SetStateOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Concurrency string `protobuf:"bytes,1,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Consistency string `protobuf:"bytes,2,opt,name=consistency,proto3" json:"consistency,omitempty"`
}

func (x *SetStateOption) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *SetStateOption) GetConcurrency() string {
	if x != nil {
		return x.Concurrency
	}
	return ""
}

func (x *SetStateOption) GetConsistency() string {
	if x != nil {
		return x.Consistency
	}
	return ""
}

type SetResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *SetResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

// GetRequest is the object describing a state fetch request.
type DeleteRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Key      string             `protobuf:"bytes,1,opt,name=key,proto3" json:"key,omitempty"`
	Etag     string             `protobuf:"bytes,2,opt,name=etag,proto3" json:"etag,omitempty"`
	Metadata map[string]string  `protobuf:"bytes,3,rep,name=metadata,proto3" json:"metadata,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
	Options  *DeleteStateOption `protobuf:"bytes,4,opt,name=options,proto3" json:"options,omitempty"`
}

func (x *DeleteRequest) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *DeleteRequest) GetKey() string {
	if x != nil {
		return x.Key
	}
	return ""
}

func (x *DeleteRequest) GetEtag() string {
	if x != nil {
		return x.Etag
	}
	return ""
}

func (x *DeleteRequest) GetMetadata() map[string]string {
	if x != nil {
		return x.Metadata
	}
	return nil
}

func (x *DeleteRequest) GetOptions() *DeleteStateOption {
	if x != nil {
		return x.Options
	}
	return nil
}

// DeleteStateOption controls how a state store reacts to a delete request.
type DeleteStateOption struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Concurrency string `protobuf:"bytes,1,opt,name=concurrency,proto3" json:"concurrency,omitempty"`
	Consistency string `protobuf:"bytes,2,opt,name=consistency,proto3" json:"consistency,omitempty"`
}

func (x *DeleteStateOption) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

func (x *DeleteStateOption) GetConcurrency() string {
	if x != nil {
		return x.Concurrency
	}
	return ""
}

func (x *DeleteStateOption) GetConsistency() string {
	if x != nil {
		return x.Consistency
	}
	return ""
}

type DeleteResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields
}

func (x *DeleteResponse) ProtoReflect() protoreflect.Message {
	panic(`not implemented`)
}

// The Store service definition.
// go:plugin type=plugin version=1
type Store interface {
	// Sends a greeting
	Init(context.Context, InitRequest) (InitResponse, error)
	Features(context.Context, FeaturesRequest) (FeaturesResponse, error)
	Get(context.Context, GetRequest) (GetResponse, error)
	Set(context.Context, SetRequest) (SetResponse, error)
	Delete(context.Context, DeleteRequest) (DeleteResponse, error)
}
