// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package uztna_active_sessions

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

func (m *GetActiveSessionRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetActiveSessionRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetActiveSessionRequest) DeepCopy() *GetActiveSessionRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetActiveSessionRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetActiveSessionRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetActiveSessionRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetActiveSessionRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetActiveSessionRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetActiveSessionRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetActiveSessionRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetActiveSessionRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
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
var DefaultGetActiveSessionRequestValidator = func() *ValidateGetActiveSessionRequest {
	v := &ValidateGetActiveSessionRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetActiveSessionRequestValidator() db.Validator {
	return DefaultGetActiveSessionRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetActiveSessionResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetActiveSessionResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetActiveSessionResponse) DeepCopy() *GetActiveSessionResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetActiveSessionResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetActiveSessionResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetActiveSessionResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetActiveSessionResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetActiveSessionResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetActiveSessionResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetActiveSessionResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetActiveSessionResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["client_ip"]; exists {

		vOpts := append(opts, db.WithValidateField("client_ip"))
		if err := fv(ctx, m.GetClientIp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["expiration_time"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_time"))
		if err := fv(ctx, m.GetExpirationTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["username"]; exists {

		vOpts := append(opts, db.WithValidateField("username"))
		if err := fv(ctx, m.GetUsername(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["variables"]; exists {

		vOpts := append(opts, db.WithValidateField("variables"))
		for idx, item := range m.GetVariables() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetActiveSessionResponseValidator = func() *ValidateGetActiveSessionResponse {
	v := &ValidateGetActiveSessionResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetActiveSessionResponseValidator() db.Validator {
	return DefaultGetActiveSessionResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *ListActiveSessionsItem) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListActiveSessionsItem) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListActiveSessionsItem) DeepCopy() *ListActiveSessionsItem {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListActiveSessionsItem{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListActiveSessionsItem) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListActiveSessionsItem) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListActiveSessionsItemValidator().Validate(ctx, m, opts...)
}

type ValidateListActiveSessionsItem struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListActiveSessionsItem) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListActiveSessionsItem)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListActiveSessionsItem got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["client_ip"]; exists {

		vOpts := append(opts, db.WithValidateField("client_ip"))
		if err := fv(ctx, m.GetClientIp(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["expiration_time"]; exists {

		vOpts := append(opts, db.WithValidateField("expiration_time"))
		if err := fv(ctx, m.GetExpirationTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["start_time"]; exists {

		vOpts := append(opts, db.WithValidateField("start_time"))
		if err := fv(ctx, m.GetStartTime(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["status"]; exists {

		vOpts := append(opts, db.WithValidateField("status"))
		if err := fv(ctx, m.GetStatus(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["username"]; exists {

		vOpts := append(opts, db.WithValidateField("username"))
		if err := fv(ctx, m.GetUsername(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListActiveSessionsItemValidator = func() *ValidateListActiveSessionsItem {
	v := &ValidateListActiveSessionsItem{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListActiveSessionsItemValidator() db.Validator {
	return DefaultListActiveSessionsItemValidator
}

// augmented methods on protoc/std generated struct

func (m *ListActiveSessionsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListActiveSessionsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListActiveSessionsRequest) DeepCopy() *ListActiveSessionsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListActiveSessionsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListActiveSessionsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListActiveSessionsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListActiveSessionsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateListActiveSessionsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListActiveSessionsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListActiveSessionsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListActiveSessionsRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["limit"]; exists {

		vOpts := append(opts, db.WithValidateField("limit"))
		if err := fv(ctx, m.GetLimit(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["offset"]; exists {

		vOpts := append(opts, db.WithValidateField("offset"))
		if err := fv(ctx, m.GetOffset(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListActiveSessionsRequestValidator = func() *ValidateListActiveSessionsRequest {
	v := &ValidateListActiveSessionsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListActiveSessionsRequestValidator() db.Validator {
	return DefaultListActiveSessionsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *ListActiveSessionsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ListActiveSessionsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ListActiveSessionsResponse) DeepCopy() *ListActiveSessionsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ListActiveSessionsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ListActiveSessionsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ListActiveSessionsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ListActiveSessionsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateListActiveSessionsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateListActiveSessionsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ListActiveSessionsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ListActiveSessionsResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["count"]; exists {

		vOpts := append(opts, db.WithValidateField("count"))
		if err := fv(ctx, m.GetCount(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["items"]; exists {

		vOpts := append(opts, db.WithValidateField("items"))
		for idx, item := range m.GetItems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultListActiveSessionsResponseValidator = func() *ValidateListActiveSessionsResponse {
	v := &ValidateListActiveSessionsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ListActiveSessionsResponseValidator() db.Validator {
	return DefaultListActiveSessionsResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *SessionVariable) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SessionVariable) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SessionVariable) DeepCopy() *SessionVariable {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SessionVariable{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SessionVariable) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SessionVariable) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SessionVariableValidator().Validate(ctx, m, opts...)
}

type ValidateSessionVariable struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSessionVariable) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SessionVariable)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SessionVariable got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["value"]; exists {

		vOpts := append(opts, db.WithValidateField("value"))
		if err := fv(ctx, m.GetValue(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["variable"]; exists {

		vOpts := append(opts, db.WithValidateField("variable"))
		if err := fv(ctx, m.GetVariable(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSessionVariableValidator = func() *ValidateSessionVariable {
	v := &ValidateSessionVariable{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func SessionVariableValidator() db.Validator {
	return DefaultSessionVariableValidator
}