//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package terraform_parameters

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

func (m *ForceDeleteRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForceDeleteRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForceDeleteRequest) DeepCopy() *ForceDeleteRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForceDeleteRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForceDeleteRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForceDeleteRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForceDeleteRequestValidator().Validate(ctx, m, opts...)
}

type ValidateForceDeleteRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForceDeleteRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForceDeleteRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForceDeleteRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["view_kind"]; exists {

		vOpts := append(opts, db.WithValidateField("view_kind"))
		if err := fv(ctx, m.GetViewKind(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["view_name"]; exists {

		vOpts := append(opts, db.WithValidateField("view_name"))
		if err := fv(ctx, m.GetViewName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultForceDeleteRequestValidator = func() *ValidateForceDeleteRequest {
	v := &ValidateForceDeleteRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ForceDeleteRequestValidator() db.Validator {
	return DefaultForceDeleteRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ForceDeleteResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ForceDeleteResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ForceDeleteResponse) DeepCopy() *ForceDeleteResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ForceDeleteResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ForceDeleteResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ForceDeleteResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ForceDeleteResponseValidator().Validate(ctx, m, opts...)
}

type ValidateForceDeleteResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateForceDeleteResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ForceDeleteResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ForceDeleteResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultForceDeleteResponseValidator = func() *ValidateForceDeleteResponse {
	v := &ValidateForceDeleteResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ForceDeleteResponseValidator() db.Validator {
	return DefaultForceDeleteResponseValidator
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

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["view_kind"]; exists {

		vOpts := append(opts, db.WithValidateField("view_kind"))
		if err := fv(ctx, m.GetViewKind(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["view_name"]; exists {

		vOpts := append(opts, db.WithValidateField("view_name"))
		if err := fv(ctx, m.GetViewName(), vOpts...); err != nil {
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

	if fv, exists := v.FldValidators["terraform_parameters"]; exists {

		vOpts := append(opts, db.WithValidateField("terraform_parameters"))
		if err := fv(ctx, m.GetTerraformParameters(), vOpts...); err != nil {
			return err
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

func (m *GetStatusResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetStatusResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetStatusResponse) DeepCopy() *GetStatusResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetStatusResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetStatusResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetStatusResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetStatusResponseValidator().Validate(ctx, m, opts...)
}

func (m *GetStatusResponse) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo

	return drInfos, nil
}

type ValidateGetStatusResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetStatusResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetStatusResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetStatusResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetStatusResponseValidator = func() *ValidateGetStatusResponse {
	v := &ValidateGetStatusResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetStatusResponseValidator() db.Validator {
	return DefaultGetStatusResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *RunRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RunRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RunRequest) DeepCopy() *RunRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RunRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RunRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RunRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RunRequestValidator().Validate(ctx, m, opts...)
}

type ValidateRunRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRunRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RunRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RunRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["action"]; exists {

		vOpts := append(opts, db.WithValidateField("action"))
		if err := fv(ctx, m.GetAction(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["view_kind"]; exists {

		vOpts := append(opts, db.WithValidateField("view_kind"))
		if err := fv(ctx, m.GetViewKind(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["view_name"]; exists {

		vOpts := append(opts, db.WithValidateField("view_name"))
		if err := fv(ctx, m.GetViewName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRunRequestValidator = func() *ValidateRunRequest {
	v := &ValidateRunRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RunRequestValidator() db.Validator {
	return DefaultRunRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *RunResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RunResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RunResponse) DeepCopy() *RunResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RunResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RunResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RunResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RunResponseValidator().Validate(ctx, m, opts...)
}

type ValidateRunResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRunResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RunResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RunResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRunResponseValidator = func() *ValidateRunResponse {
	v := &ValidateRunResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RunResponseValidator() db.Validator {
	return DefaultRunResponseValidator
}
