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

func (m *CreateSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CreateSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CreateSpecType) DeepCopy() *CreateSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CreateSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CreateSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CreateSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CreateSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateCreateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCreateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CreateSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CreateSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["added_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signatures"))
		for idx, item := range m.GetAddedSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["added_signatures_data"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signatures_data"))
		for idx, item := range m.GetAddedSignaturesData() {
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

	if fv, exists := v.FldValidators["updated_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signatures"))
		for idx, item := range m.GetUpdatedSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["updated_signatures_data"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signatures_data"))
		for idx, item := range m.GetUpdatedSignaturesData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["added_signatures_data"] = SignatureValidator().Validate

	v.FldValidators["updated_signatures_data"] = SignatureValidator().Validate

	return v
}()

func CreateSpecTypeValidator() db.Validator {
	return DefaultCreateSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GetSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GetSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GetSpecType) DeepCopy() *GetSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GetSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GetSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GetSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GetSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GetSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GetSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["added_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signatures"))
		for idx, item := range m.GetAddedSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["added_signatures_data"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signatures_data"))
		for idx, item := range m.GetAddedSignaturesData() {
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

	if fv, exists := v.FldValidators["updated_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signatures"))
		for idx, item := range m.GetUpdatedSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["updated_signatures_data"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signatures_data"))
		for idx, item := range m.GetUpdatedSignaturesData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["added_signatures_data"] = SignatureValidator().Validate

	v.FldValidators["updated_signatures_data"] = SignatureValidator().Validate

	return v
}()

func GetSpecTypeValidator() db.Validator {
	return DefaultGetSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *GlobalSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GlobalSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GlobalSpecType) DeepCopy() *GlobalSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GlobalSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GlobalSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GlobalSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GlobalSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GlobalSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GlobalSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["added_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signatures"))
		for idx, item := range m.GetAddedSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["added_signatures_data"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signatures_data"))
		for idx, item := range m.GetAddedSignaturesData() {
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

	if fv, exists := v.FldValidators["updated_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signatures"))
		for idx, item := range m.GetUpdatedSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["updated_signatures_data"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signatures_data"))
		for idx, item := range m.GetUpdatedSignaturesData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["added_signatures_data"] = SignatureValidator().Validate

	v.FldValidators["updated_signatures_data"] = SignatureValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *ReplaceSpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *ReplaceSpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *ReplaceSpecType) DeepCopy() *ReplaceSpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &ReplaceSpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *ReplaceSpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *ReplaceSpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return ReplaceSpecTypeValidator().Validate(ctx, m, opts...)
}

type ValidateReplaceSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateReplaceSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*ReplaceSpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *ReplaceSpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["added_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signatures"))
		for idx, item := range m.GetAddedSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["added_signatures_data"]; exists {

		vOpts := append(opts, db.WithValidateField("added_signatures_data"))
		for idx, item := range m.GetAddedSignaturesData() {
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

	if fv, exists := v.FldValidators["updated_signatures"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signatures"))
		for idx, item := range m.GetUpdatedSignatures() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["updated_signatures_data"]; exists {

		vOpts := append(opts, db.WithValidateField("updated_signatures_data"))
		for idx, item := range m.GetUpdatedSignaturesData() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["added_signatures_data"] = SignatureValidator().Validate

	v.FldValidators["updated_signatures_data"] = SignatureValidator().Validate

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *Signature) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Signature) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Signature) DeepCopy() *Signature {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Signature{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Signature) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Signature) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SignatureValidator().Validate(ctx, m, opts...)
}

type ValidateSignature struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSignature) LastUpdateValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for last_update")
	}

	return validatorFn, nil
}

func (v *ValidateSignature) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Signature)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Signature got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["accuracy"]; exists {

		vOpts := append(opts, db.WithValidateField("accuracy"))
		if err := fv(ctx, m.GetAccuracy(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["applies_to"]; exists {

		vOpts := append(opts, db.WithValidateField("applies_to"))
		if err := fv(ctx, m.GetAppliesTo(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["attack_type"]; exists {

		vOpts := append(opts, db.WithValidateField("attack_type"))
		if err := fv(ctx, m.GetAttackType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["description"]; exists {

		vOpts := append(opts, db.WithValidateField("description"))
		if err := fv(ctx, m.GetDescription(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["id"]; exists {

		vOpts := append(opts, db.WithValidateField("id"))
		if err := fv(ctx, m.GetId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_update"]; exists {

		vOpts := append(opts, db.WithValidateField("last_update"))
		if err := fv(ctx, m.GetLastUpdate(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["references"]; exists {

		vOpts := append(opts, db.WithValidateField("references"))
		for idx, item := range m.GetReferences() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["risk"]; exists {

		vOpts := append(opts, db.WithValidateField("risk"))
		if err := fv(ctx, m.GetRisk(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["systems"]; exists {

		vOpts := append(opts, db.WithValidateField("systems"))
		for idx, item := range m.GetSystems() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSignatureValidator = func() *ValidateSignature {
	v := &ValidateSignature{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhLastUpdate := v.LastUpdateValidationRuleHandler
	rulesLastUpdate := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhLastUpdate(rulesLastUpdate)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Signature.last_update: %s", err)
		panic(errMsg)
	}
	v.FldValidators["last_update"] = vFn

	return v
}()

func SignatureValidator() db.Validator {
	return DefaultSignatureValidator
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.AddedSignatures = f.GetAddedSignatures()
	m.AddedSignaturesData = f.GetAddedSignaturesData()
	m.ReleaseDate = f.GetReleaseDate()
	m.UpdatedSignatures = f.GetUpdatedSignatures()
	m.UpdatedSignaturesData = f.GetUpdatedSignaturesData()
}

func (m *CreateSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *CreateSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *CreateSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.AddedSignatures = m1.AddedSignatures
	f.AddedSignaturesData = m1.AddedSignaturesData
	f.ReleaseDate = m1.ReleaseDate
	f.UpdatedSignatures = m1.UpdatedSignatures
	f.UpdatedSignaturesData = m1.UpdatedSignaturesData
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *CreateSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

func (m *GetSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.AddedSignatures = f.GetAddedSignatures()
	m.AddedSignaturesData = f.GetAddedSignaturesData()
	m.ReleaseDate = f.GetReleaseDate()
	m.UpdatedSignatures = f.GetUpdatedSignatures()
	m.UpdatedSignaturesData = f.GetUpdatedSignaturesData()
}

func (m *GetSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *GetSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *GetSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.AddedSignatures = m1.AddedSignatures
	f.AddedSignaturesData = m1.AddedSignaturesData
	f.ReleaseDate = m1.ReleaseDate
	f.UpdatedSignatures = m1.UpdatedSignatures
	f.UpdatedSignaturesData = m1.UpdatedSignaturesData
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *GetSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

func (m *ReplaceSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.AddedSignatures = f.GetAddedSignatures()
	m.AddedSignaturesData = f.GetAddedSignaturesData()
	m.ReleaseDate = f.GetReleaseDate()
	m.UpdatedSignatures = f.GetUpdatedSignatures()
	m.UpdatedSignaturesData = f.GetUpdatedSignaturesData()
}

func (m *ReplaceSpecType) FromGlobalSpecType(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) FromGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.fromGlobalSpecType(f, false)
}

func (m *ReplaceSpecType) toGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	m1 := m
	if withDeepCopy {
		m1 = m.DeepCopy()
	}
	_ = m1

	f.AddedSignatures = m1.AddedSignatures
	f.AddedSignaturesData = m1.AddedSignaturesData
	f.ReleaseDate = m1.ReleaseDate
	f.UpdatedSignatures = m1.UpdatedSignatures
	f.UpdatedSignaturesData = m1.UpdatedSignaturesData
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
