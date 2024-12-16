// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package waf_signatures_changelog

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

func (m *ReleaseSignatures) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReleaseSignatures) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReleaseSignatures) DeepCopy() *ReleaseSignatures {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReleaseSignatures{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReleaseSignatures) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReleaseSignatures) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReleaseSignaturesValidator().Validate(ctx, m, opts...)
}

type ValidateReleaseSignatures struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReleaseSignatures) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReleaseSignatures)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReleaseSignatures got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["added_signature_ids"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signature_ids"))
		for idx, item := range m.GetAddedSignatureIds() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["release_date"]; exists {

		vOpts := append(opts, db.WithValidateField("release_date"))
		if err := fv(ctx, m.GetReleaseDate(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["updated_signature_ids"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signature_ids"))
		for idx, item := range m.GetUpdatedSignatureIds() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReleaseSignaturesValidator = func() *ValidateReleaseSignatures {
	v := &ValidateReleaseSignatures{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ReleaseSignaturesValidator() db.Validator {
	return DefaultReleaseSignaturesValidator
}

// augmented methods on protoc/std generated struct

func (m *ReleasedSignaturesReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReleasedSignaturesReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReleasedSignaturesReq) DeepCopy() *ReleasedSignaturesReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReleasedSignaturesReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReleasedSignaturesReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReleasedSignaturesReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReleasedSignaturesReqValidator().Validate(ctx, m, opts...)
}

type ValidateReleasedSignaturesReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReleasedSignaturesReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReleasedSignaturesReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReleasedSignaturesReq got type %s", t)
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

	if fv, exists := v.FldValidators["vh_name"]; exists {

		vOpts := append(opts, db.WithValidateField("vh_name"))
		if err := fv(ctx, m.GetVhName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReleasedSignaturesReqValidator = func() *ValidateReleasedSignaturesReq {
	v := &ValidateReleasedSignaturesReq{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ReleasedSignaturesReqValidator() db.Validator {
	return DefaultReleasedSignaturesReqValidator
}

// augmented methods on protoc/std generated struct

func (m *ReleasedSignaturesRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReleasedSignaturesRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReleasedSignaturesRsp) DeepCopy() *ReleasedSignaturesRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReleasedSignaturesRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReleasedSignaturesRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReleasedSignaturesRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReleasedSignaturesRspValidator().Validate(ctx, m, opts...)
}

type ValidateReleasedSignaturesRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReleasedSignaturesRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReleasedSignaturesRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReleasedSignaturesRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["release_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("release_signatures"))
		for idx, item := range m.GetReleaseSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReleasedSignaturesRspValidator = func() *ValidateReleasedSignaturesRsp {
	v := &ValidateReleasedSignaturesRsp{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func ReleasedSignaturesRspValidator() db.Validator {
	return DefaultReleasedSignaturesRspValidator
}
