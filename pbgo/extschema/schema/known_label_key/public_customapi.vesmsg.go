// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package known_label_key

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *CreateRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateRequest) DeepCopy() *CreateRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateRequestValidator().Validate(ctx, m, opts...)
}

type ValidateCreateRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["key"]; exists {

		vOpts := append(opts, db.WithValidateField("key"))
		if err := fv(ctx, m.GetKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateRequestValidator = func() *ValidateCreateRequest {
	v := &ValidateCreateRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CreateRequestValidator() db.Validator {
	return DefaultCreateRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *CreateResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateResponse) DeepCopy() *CreateResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateResponseValidator().Validate(ctx, m, opts...)
}

type ValidateCreateResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["label_key"]; exists {

		vOpts := append(opts, db.WithValidateField("label_key"))
		if err := fv(ctx, m.GetLabelKey(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateResponseValidator = func() *ValidateCreateResponse {
	v := &ValidateCreateResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func CreateResponseValidator() db.Validator {
	return DefaultCreateResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *DeleteRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DeleteRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DeleteRequest) DeepCopy() *DeleteRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DeleteRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DeleteRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DeleteRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DeleteRequestValidator().Validate(ctx, m, opts...)
}

type ValidateDeleteRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDeleteRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DeleteRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DeleteRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["key"]; exists {

		vOpts := append(opts, db.WithValidateField("key"))
		if err := fv(ctx, m.GetKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDeleteRequestValidator = func() *ValidateDeleteRequest {
	v := &ValidateDeleteRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DeleteRequestValidator() db.Validator {
	return DefaultDeleteRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *DeleteResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *DeleteResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *DeleteResponse) DeepCopy() *DeleteResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &DeleteResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *DeleteResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *DeleteResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return DeleteResponseValidator().Validate(ctx, m, opts...)
}

type ValidateDeleteResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateDeleteResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*DeleteResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *DeleteResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultDeleteResponseValidator = func() *ValidateDeleteResponse {
	v := &ValidateDeleteResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func DeleteResponseValidator() db.Validator {
	return DefaultDeleteResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *GetRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetRequest) DeepCopy() *GetRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["key"]; exists {

		vOpts := append(opts, db.WithValidateField("key"))
		if err := fv(ctx, m.GetKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query"]; exists {

		vOpts := append(opts, db.WithValidateField("query"))
		if err := fv(ctx, m.GetQuery(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetRequestValidator = func() *ValidateGetRequest {
	v := &ValidateGetRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetRequestValidator() db.Validator {
	return DefaultGetRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetResponse) DeepCopy() *GetResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["label_key"]; exists {

		vOpts := append(opts, db.WithValidateField("label_key"))
		for idx, item := range m.GetLabelKey() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetResponseValidator = func() *ValidateGetResponse {
	v := &ValidateGetResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetResponseValidator() db.Validator {
	return DefaultGetResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *LabelKeyType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *LabelKeyType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *LabelKeyType) DeepCopy() *LabelKeyType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &LabelKeyType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *LabelKeyType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *LabelKeyType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return LabelKeyTypeValidator().Validate(ctx, m, opts...)
}

type ValidateLabelKeyType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateLabelKeyType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*LabelKeyType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *LabelKeyType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["key"]; exists {

		vOpts := append(opts, db.WithValidateField("key"))
		if err := fv(ctx, m.GetKey(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultLabelKeyTypeValidator = func() *ValidateLabelKeyType {
	v := &ValidateLabelKeyType{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func LabelKeyTypeValidator() db.Validator {
	return DefaultLabelKeyTypeValidator
}
