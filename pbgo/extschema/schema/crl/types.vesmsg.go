// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package crl

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

func (v *ValidateCreateSpecType) ServerAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for server_address")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) ServerPortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for server_port")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) RefreshIntervalValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for refresh_interval")
	}

	return validatorFn, nil
}

func (v *ValidateCreateSpecType) TimeoutValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for timeout")
	}

	return validatorFn, nil
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

	switch m.GetAccessInfoChoice().(type) {
	case *CreateSpecType_HttpAccess:
		if fv, exists := v.FldValidators["access_info_choice.http_access"]; exists {
			val := m.GetAccessInfoChoice().(*CreateSpecType_HttpAccess).HttpAccess
			vOpts := append(opts,
				db.WithValidateField("access_info_choice"),
				db.WithValidateField("http_access"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["refresh_interval"]; exists {

		vOpts := append(opts, db.WithValidateField("refresh_interval"))
		if err := fv(ctx, m.GetRefreshInterval(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_address"]; exists {

		vOpts := append(opts, db.WithValidateField("server_address"))
		if err := fv(ctx, m.GetServerAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_port"]; exists {

		vOpts := append(opts, db.WithValidateField("server_port"))
		if err := fv(ctx, m.GetServerPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["timeout"]; exists {

		vOpts := append(opts, db.WithValidateField("timeout"))
		if err := fv(ctx, m.GetTimeout(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["verify_all_certs_with_crl"]; exists {

		vOpts := append(opts, db.WithValidateField("verify_all_certs_with_crl"))
		if err := fv(ctx, m.GetVerifyAllCertsWithCrl(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCreateSpecTypeValidator = func() *ValidateCreateSpecType {
	v := &ValidateCreateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhServerAddress := v.ServerAddressValidationRuleHandler
	rulesServerAddress := map[string]string{
		"ves.io.schema.rules.message.required":      "true",
		"ves.io.schema.rules.string.hostname_or_ip": "true",
		"ves.io.schema.rules.string.max_len":        "255",
	}
	vFn, err = vrhServerAddress(rulesServerAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.server_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["server_address"] = vFn

	vrhServerPort := v.ServerPortValidationRuleHandler
	rulesServerPort := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "65535",
	}
	vFn, err = vrhServerPort(rulesServerPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.server_port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["server_port"] = vFn

	vrhRefreshInterval := v.RefreshIntervalValidationRuleHandler
	rulesRefreshInterval := map[string]string{
		"ves.io.schema.rules.uint32.gte": "6",
		"ves.io.schema.rules.uint32.lte": "168",
	}
	vFn, err = vrhRefreshInterval(rulesRefreshInterval)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.refresh_interval: %s", err)
		panic(errMsg)
	}
	v.FldValidators["refresh_interval"] = vFn

	vrhTimeout := v.TimeoutValidationRuleHandler
	rulesTimeout := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "180",
	}
	vFn, err = vrhTimeout(rulesTimeout)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CreateSpecType.timeout: %s", err)
		panic(errMsg)
	}
	v.FldValidators["timeout"] = vFn

	v.FldValidators["access_info_choice.http_access"] = HTTPAccessInfoValidator().Validate

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

func (v *ValidateGetSpecType) ServerAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for server_address")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) ServerPortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for server_port")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) RefreshIntervalValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for refresh_interval")
	}

	return validatorFn, nil
}

func (v *ValidateGetSpecType) TimeoutValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for timeout")
	}

	return validatorFn, nil
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

	switch m.GetAccessInfoChoice().(type) {
	case *GetSpecType_HttpAccess:
		if fv, exists := v.FldValidators["access_info_choice.http_access"]; exists {
			val := m.GetAccessInfoChoice().(*GetSpecType_HttpAccess).HttpAccess
			vOpts := append(opts,
				db.WithValidateField("access_info_choice"),
				db.WithValidateField("http_access"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["refresh_interval"]; exists {

		vOpts := append(opts, db.WithValidateField("refresh_interval"))
		if err := fv(ctx, m.GetRefreshInterval(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_address"]; exists {

		vOpts := append(opts, db.WithValidateField("server_address"))
		if err := fv(ctx, m.GetServerAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_port"]; exists {

		vOpts := append(opts, db.WithValidateField("server_port"))
		if err := fv(ctx, m.GetServerPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["timeout"]; exists {

		vOpts := append(opts, db.WithValidateField("timeout"))
		if err := fv(ctx, m.GetTimeout(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["verify_all_certs_with_crl"]; exists {

		vOpts := append(opts, db.WithValidateField("verify_all_certs_with_crl"))
		if err := fv(ctx, m.GetVerifyAllCertsWithCrl(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGetSpecTypeValidator = func() *ValidateGetSpecType {
	v := &ValidateGetSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhServerAddress := v.ServerAddressValidationRuleHandler
	rulesServerAddress := map[string]string{
		"ves.io.schema.rules.message.required":      "true",
		"ves.io.schema.rules.string.hostname_or_ip": "true",
		"ves.io.schema.rules.string.max_len":        "255",
	}
	vFn, err = vrhServerAddress(rulesServerAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.server_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["server_address"] = vFn

	vrhServerPort := v.ServerPortValidationRuleHandler
	rulesServerPort := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "65535",
	}
	vFn, err = vrhServerPort(rulesServerPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.server_port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["server_port"] = vFn

	vrhRefreshInterval := v.RefreshIntervalValidationRuleHandler
	rulesRefreshInterval := map[string]string{
		"ves.io.schema.rules.uint32.gte": "6",
		"ves.io.schema.rules.uint32.lte": "168",
	}
	vFn, err = vrhRefreshInterval(rulesRefreshInterval)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.refresh_interval: %s", err)
		panic(errMsg)
	}
	v.FldValidators["refresh_interval"] = vFn

	vrhTimeout := v.TimeoutValidationRuleHandler
	rulesTimeout := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "180",
	}
	vFn, err = vrhTimeout(rulesTimeout)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.timeout: %s", err)
		panic(errMsg)
	}
	v.FldValidators["timeout"] = vFn

	v.FldValidators["access_info_choice.http_access"] = HTTPAccessInfoValidator().Validate

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

func (v *ValidateGlobalSpecType) ServerAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for server_address")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) ServerPortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for server_port")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) RefreshIntervalValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for refresh_interval")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) TimeoutValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for timeout")
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

	switch m.GetAccessInfoChoice().(type) {
	case *GlobalSpecType_HttpAccess:
		if fv, exists := v.FldValidators["access_info_choice.http_access"]; exists {
			val := m.GetAccessInfoChoice().(*GlobalSpecType_HttpAccess).HttpAccess
			vOpts := append(opts,
				db.WithValidateField("access_info_choice"),
				db.WithValidateField("http_access"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["refresh_interval"]; exists {

		vOpts := append(opts, db.WithValidateField("refresh_interval"))
		if err := fv(ctx, m.GetRefreshInterval(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_address"]; exists {

		vOpts := append(opts, db.WithValidateField("server_address"))
		if err := fv(ctx, m.GetServerAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_port"]; exists {

		vOpts := append(opts, db.WithValidateField("server_port"))
		if err := fv(ctx, m.GetServerPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["timeout"]; exists {

		vOpts := append(opts, db.WithValidateField("timeout"))
		if err := fv(ctx, m.GetTimeout(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["verify_all_certs_with_crl"]; exists {

		vOpts := append(opts, db.WithValidateField("verify_all_certs_with_crl"))
		if err := fv(ctx, m.GetVerifyAllCertsWithCrl(), vOpts...); err != nil {
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

	vrhServerAddress := v.ServerAddressValidationRuleHandler
	rulesServerAddress := map[string]string{
		"ves.io.schema.rules.message.required":      "true",
		"ves.io.schema.rules.string.hostname_or_ip": "true",
		"ves.io.schema.rules.string.max_len":        "255",
	}
	vFn, err = vrhServerAddress(rulesServerAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.server_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["server_address"] = vFn

	vrhServerPort := v.ServerPortValidationRuleHandler
	rulesServerPort := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "65535",
	}
	vFn, err = vrhServerPort(rulesServerPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.server_port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["server_port"] = vFn

	vrhRefreshInterval := v.RefreshIntervalValidationRuleHandler
	rulesRefreshInterval := map[string]string{
		"ves.io.schema.rules.uint32.gte": "6",
		"ves.io.schema.rules.uint32.lte": "168",
	}
	vFn, err = vrhRefreshInterval(rulesRefreshInterval)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.refresh_interval: %s", err)
		panic(errMsg)
	}
	v.FldValidators["refresh_interval"] = vFn

	vrhTimeout := v.TimeoutValidationRuleHandler
	rulesTimeout := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "180",
	}
	vFn, err = vrhTimeout(rulesTimeout)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.timeout: %s", err)
		panic(errMsg)
	}
	v.FldValidators["timeout"] = vFn

	v.FldValidators["access_info_choice.http_access"] = HTTPAccessInfoValidator().Validate

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

// augmented methods on protoc/std generated struct

func (m *HTTPAccessInfo) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *HTTPAccessInfo) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *HTTPAccessInfo) DeepCopy() *HTTPAccessInfo {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &HTTPAccessInfo{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *HTTPAccessInfo) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *HTTPAccessInfo) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return HTTPAccessInfoValidator().Validate(ctx, m, opts...)
}

type ValidateHTTPAccessInfo struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateHTTPAccessInfo) PathValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for path")
	}

	return validatorFn, nil
}

func (v *ValidateHTTPAccessInfo) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*HTTPAccessInfo)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *HTTPAccessInfo got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["path"]; exists {

		vOpts := append(opts, db.WithValidateField("path"))
		if err := fv(ctx, m.GetPath(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultHTTPAccessInfoValidator = func() *ValidateHTTPAccessInfo {
	v := &ValidateHTTPAccessInfo{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhPath := v.PathValidationRuleHandler
	rulesPath := map[string]string{
		"ves.io.schema.rules.string.http_path": "true",
		"ves.io.schema.rules.string.max_len":   "256",
	}
	vFn, err = vrhPath(rulesPath)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for HTTPAccessInfo.path: %s", err)
		panic(errMsg)
	}
	v.FldValidators["path"] = vFn

	return v
}()

func HTTPAccessInfoValidator() db.Validator {
	return DefaultHTTPAccessInfoValidator
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

func (v *ValidateReplaceSpecType) ServerAddressValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for server_address")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) ServerPortValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for server_port")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) RefreshIntervalValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for refresh_interval")
	}

	return validatorFn, nil
}

func (v *ValidateReplaceSpecType) TimeoutValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewUint32ValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for timeout")
	}

	return validatorFn, nil
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

	switch m.GetAccessInfoChoice().(type) {
	case *ReplaceSpecType_HttpAccess:
		if fv, exists := v.FldValidators["access_info_choice.http_access"]; exists {
			val := m.GetAccessInfoChoice().(*ReplaceSpecType_HttpAccess).HttpAccess
			vOpts := append(opts,
				db.WithValidateField("access_info_choice"),
				db.WithValidateField("http_access"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["refresh_interval"]; exists {

		vOpts := append(opts, db.WithValidateField("refresh_interval"))
		if err := fv(ctx, m.GetRefreshInterval(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_address"]; exists {

		vOpts := append(opts, db.WithValidateField("server_address"))
		if err := fv(ctx, m.GetServerAddress(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["server_port"]; exists {

		vOpts := append(opts, db.WithValidateField("server_port"))
		if err := fv(ctx, m.GetServerPort(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["timeout"]; exists {

		vOpts := append(opts, db.WithValidateField("timeout"))
		if err := fv(ctx, m.GetTimeout(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["verify_all_certs_with_crl"]; exists {

		vOpts := append(opts, db.WithValidateField("verify_all_certs_with_crl"))
		if err := fv(ctx, m.GetVerifyAllCertsWithCrl(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultReplaceSpecTypeValidator = func() *ValidateReplaceSpecType {
	v := &ValidateReplaceSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhServerAddress := v.ServerAddressValidationRuleHandler
	rulesServerAddress := map[string]string{
		"ves.io.schema.rules.message.required":      "true",
		"ves.io.schema.rules.string.hostname_or_ip": "true",
		"ves.io.schema.rules.string.max_len":        "255",
	}
	vFn, err = vrhServerAddress(rulesServerAddress)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.server_address: %s", err)
		panic(errMsg)
	}
	v.FldValidators["server_address"] = vFn

	vrhServerPort := v.ServerPortValidationRuleHandler
	rulesServerPort := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "65535",
	}
	vFn, err = vrhServerPort(rulesServerPort)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.server_port: %s", err)
		panic(errMsg)
	}
	v.FldValidators["server_port"] = vFn

	vrhRefreshInterval := v.RefreshIntervalValidationRuleHandler
	rulesRefreshInterval := map[string]string{
		"ves.io.schema.rules.uint32.gte": "6",
		"ves.io.schema.rules.uint32.lte": "168",
	}
	vFn, err = vrhRefreshInterval(rulesRefreshInterval)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.refresh_interval: %s", err)
		panic(errMsg)
	}
	v.FldValidators["refresh_interval"] = vFn

	vrhTimeout := v.TimeoutValidationRuleHandler
	rulesTimeout := map[string]string{
		"ves.io.schema.rules.uint32.gte": "1",
		"ves.io.schema.rules.uint32.lte": "180",
	}
	vFn, err = vrhTimeout(rulesTimeout)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for ReplaceSpecType.timeout: %s", err)
		panic(errMsg)
	}
	v.FldValidators["timeout"] = vFn

	v.FldValidators["access_info_choice.http_access"] = HTTPAccessInfoValidator().Validate

	return v
}()

func ReplaceSpecTypeValidator() db.Validator {
	return DefaultReplaceSpecTypeValidator
}

// create setters in CreateSpecType from GlobalSpecType for oneof fields
func (r *CreateSpecType) SetAccessInfoChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.AccessInfoChoice.(type) {
	case nil:
		o.AccessInfoChoice = nil

	case *CreateSpecType_HttpAccess:
		o.AccessInfoChoice = &GlobalSpecType_HttpAccess{HttpAccess: of.HttpAccess}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *CreateSpecType) GetAccessInfoChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.AccessInfoChoice.(type) {
	case nil:
		r.AccessInfoChoice = nil

	case *GlobalSpecType_HttpAccess:
		r.AccessInfoChoice = &CreateSpecType_HttpAccess{HttpAccess: of.HttpAccess}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *CreateSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.GetAccessInfoChoiceFromGlobalSpecType(f)
	m.RefreshInterval = f.GetRefreshInterval()
	m.ServerAddress = f.GetServerAddress()
	m.ServerPort = f.GetServerPort()
	m.Timeout = f.GetTimeout()
	m.VerifyAllCertsWithCrl = f.GetVerifyAllCertsWithCrl()
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

	m1.SetAccessInfoChoiceToGlobalSpecType(f)
	f.RefreshInterval = m1.RefreshInterval
	f.ServerAddress = m1.ServerAddress
	f.ServerPort = m1.ServerPort
	f.Timeout = m1.Timeout
	f.VerifyAllCertsWithCrl = m1.VerifyAllCertsWithCrl
}

func (m *CreateSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *CreateSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

// create setters in GetSpecType from GlobalSpecType for oneof fields
func (r *GetSpecType) SetAccessInfoChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.AccessInfoChoice.(type) {
	case nil:
		o.AccessInfoChoice = nil

	case *GetSpecType_HttpAccess:
		o.AccessInfoChoice = &GlobalSpecType_HttpAccess{HttpAccess: of.HttpAccess}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *GetSpecType) GetAccessInfoChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.AccessInfoChoice.(type) {
	case nil:
		r.AccessInfoChoice = nil

	case *GlobalSpecType_HttpAccess:
		r.AccessInfoChoice = &GetSpecType_HttpAccess{HttpAccess: of.HttpAccess}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *GetSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.GetAccessInfoChoiceFromGlobalSpecType(f)
	m.RefreshInterval = f.GetRefreshInterval()
	m.ServerAddress = f.GetServerAddress()
	m.ServerPort = f.GetServerPort()
	m.Timeout = f.GetTimeout()
	m.VerifyAllCertsWithCrl = f.GetVerifyAllCertsWithCrl()
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

	m1.SetAccessInfoChoiceToGlobalSpecType(f)
	f.RefreshInterval = m1.RefreshInterval
	f.ServerAddress = m1.ServerAddress
	f.ServerPort = m1.ServerPort
	f.Timeout = m1.Timeout
	f.VerifyAllCertsWithCrl = m1.VerifyAllCertsWithCrl
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *GetSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}

// create setters in ReplaceSpecType from GlobalSpecType for oneof fields
func (r *ReplaceSpecType) SetAccessInfoChoiceToGlobalSpecType(o *GlobalSpecType) error {
	switch of := r.AccessInfoChoice.(type) {
	case nil:
		o.AccessInfoChoice = nil

	case *ReplaceSpecType_HttpAccess:
		o.AccessInfoChoice = &GlobalSpecType_HttpAccess{HttpAccess: of.HttpAccess}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (r *ReplaceSpecType) GetAccessInfoChoiceFromGlobalSpecType(o *GlobalSpecType) error {
	switch of := o.AccessInfoChoice.(type) {
	case nil:
		r.AccessInfoChoice = nil

	case *GlobalSpecType_HttpAccess:
		r.AccessInfoChoice = &ReplaceSpecType_HttpAccess{HttpAccess: of.HttpAccess}

	default:
		return fmt.Errorf("Unknown oneof field %T", of)
	}
	return nil
}

func (m *ReplaceSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.GetAccessInfoChoiceFromGlobalSpecType(f)
	m.RefreshInterval = f.GetRefreshInterval()
	m.ServerAddress = f.GetServerAddress()
	m.ServerPort = f.GetServerPort()
	m.Timeout = f.GetTimeout()
	m.VerifyAllCertsWithCrl = f.GetVerifyAllCertsWithCrl()
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

	m1.SetAccessInfoChoiceToGlobalSpecType(f)
	f.RefreshInterval = m1.RefreshInterval
	f.ServerAddress = m1.ServerAddress
	f.ServerPort = m1.ServerPort
	f.Timeout = m1.Timeout
	f.VerifyAllCertsWithCrl = m1.VerifyAllCertsWithCrl
}

func (m *ReplaceSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *ReplaceSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
