// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
package stored_object

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

func (v *ValidateGlobalSpecType) ObjectTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for object_type")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) VersionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for version")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) ContentHashValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for content_hash")
	}

	return validatorFn, nil
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

	if fv, exists := v.FldValidators["content_format"]; exists {

		vOpts := append(opts, db.WithValidateField("content_format"))
		if err := fv(ctx, m.GetContentFormat(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["content_hash"]; exists {

		vOpts := append(opts, db.WithValidateField("content_hash"))
		if err := fv(ctx, m.GetContentHash(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["content_transport"]; exists {

		vOpts := append(opts, db.WithValidateField("content_transport"))
		if err := fv(ctx, m.GetContentTransport(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["content_url_in_storage"]; exists {

		vOpts := append(opts, db.WithValidateField("content_url_in_storage"))
		if err := fv(ctx, m.GetContentUrlInStorage(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	switch m.GetObjectAttributes().(type) {
	case *GlobalSpecType_NoAttributes:
		if fv, exists := v.FldValidators["object_attributes.no_attributes"]; exists {
			val := m.GetObjectAttributes().(*GlobalSpecType_NoAttributes).NoAttributes
			vOpts := append(opts,
				db.WithValidateField("object_attributes"),
				db.WithValidateField("no_attributes"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GlobalSpecType_MobileSdk:
		if fv, exists := v.FldValidators["object_attributes.mobile_sdk"]; exists {
			val := m.GetObjectAttributes().(*GlobalSpecType_MobileSdk).MobileSdk
			vOpts := append(opts,
				db.WithValidateField("object_attributes"),
				db.WithValidateField("mobile_sdk"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *GlobalSpecType_MobileIntegrator:
		if fv, exists := v.FldValidators["object_attributes.mobile_integrator"]; exists {
			val := m.GetObjectAttributes().(*GlobalSpecType_MobileIntegrator).MobileIntegrator
			vOpts := append(opts,
				db.WithValidateField("object_attributes"),
				db.WithValidateField("mobile_integrator"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["object_type"]; exists {

		vOpts := append(opts, db.WithValidateField("object_type"))
		if err := fv(ctx, m.GetObjectType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["version"]; exists {

		vOpts := append(opts, db.WithValidateField("version"))
		if err := fv(ctx, m.GetVersion(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGlobalSpecTypeValidator = func() *ValidateGlobalSpecType {
	v := &ValidateGlobalSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhObjectType := v.ObjectTypeValidationRuleHandler
	rulesObjectType := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhObjectType(rulesObjectType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.object_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["object_type"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhVersion := v.VersionValidationRuleHandler
	rulesVersion := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhVersion(rulesVersion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.version: %s", err)
		panic(errMsg)
	}
	v.FldValidators["version"] = vFn

	vrhContentHash := v.ContentHashValidationRuleHandler
	rulesContentHash := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhContentHash(rulesContentHash)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.content_hash: %s", err)
		panic(errMsg)
	}
	v.FldValidators["content_hash"] = vFn

	v.FldValidators["object_attributes.mobile_sdk"] = MobileSDKAttributesValidator().Validate
	v.FldValidators["object_attributes.mobile_integrator"] = MobileIntegratorAttributesValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}
