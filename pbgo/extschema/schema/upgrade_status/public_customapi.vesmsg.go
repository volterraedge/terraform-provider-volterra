// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package upgrade_status

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

func (m *GetUpgradeStatusRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetUpgradeStatusRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetUpgradeStatusRequest) DeepCopy() *GetUpgradeStatusRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetUpgradeStatusRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetUpgradeStatusRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetUpgradeStatusRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetUpgradeStatusRequestValidator().Validate(ctx, m, opts...)
}

type ValidateGetUpgradeStatusRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetUpgradeStatusRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetUpgradeStatusRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetUpgradeStatusRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
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
var DefaultGetUpgradeStatusRequestValidator = func() *ValidateGetUpgradeStatusRequest {
	v := &ValidateGetUpgradeStatusRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func GetUpgradeStatusRequestValidator() db.Validator {
	return DefaultGetUpgradeStatusRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *GetUpgradeStatusResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetUpgradeStatusResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetUpgradeStatusResponse) DeepCopy() *GetUpgradeStatusResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetUpgradeStatusResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetUpgradeStatusResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetUpgradeStatusResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetUpgradeStatusResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGetUpgradeStatusResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetUpgradeStatusResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetUpgradeStatusResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetUpgradeStatusResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["upgrade_status"]; exists {

		vOpts := append(opts, db.WithValidateField("upgrade_status"))
		if err := fv(ctx, m.GetUpgradeStatus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetUpgradeStatusResponseValidator = func() *ValidateGetUpgradeStatusResponse {
	v := &ValidateGetUpgradeStatusResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["upgrade_status"] = GlobalSpecTypeValidator().Validate

	return v
}()

func GetUpgradeStatusResponseValidator() db.Validator {
	return DefaultGetUpgradeStatusResponseValidator
}