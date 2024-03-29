// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package cdn_loadbalancer

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

func (m *UpdateServiceDomainsRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpdateServiceDomainsRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpdateServiceDomainsRequest) DeepCopy() *UpdateServiceDomainsRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpdateServiceDomainsRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpdateServiceDomainsRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpdateServiceDomainsRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpdateServiceDomainsRequestValidator().Validate(ctx, m, opts...)
}

type ValidateUpdateServiceDomainsRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpdateServiceDomainsRequest) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateUpdateServiceDomainsRequest) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateUpdateServiceDomainsRequest) TenantValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for tenant")
	}

	return validatorFn, nil
}

func (v *ValidateUpdateServiceDomainsRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpdateServiceDomainsRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpdateServiceDomainsRequest got type %s", t)
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

	if fv, exists := v.FldValidators["service_domains"]; exists {

		vOpts := append(opts, db.WithValidateField("service_domains"))
		for idx, item := range m.GetServiceDomains() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["tenant"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant"))
		if err := fv(ctx, m.GetTenant(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpdateServiceDomainsRequestValidator = func() *ValidateUpdateServiceDomainsRequest {
	v := &ValidateUpdateServiceDomainsRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhNamespace := v.NamespaceValidationRuleHandler
	rulesNamespace := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhNamespace(rulesNamespace)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpdateServiceDomainsRequest.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpdateServiceDomainsRequest.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhTenant := v.TenantValidationRuleHandler
	rulesTenant := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhTenant(rulesTenant)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for UpdateServiceDomainsRequest.tenant: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tenant"] = vFn

	return v
}()

func UpdateServiceDomainsRequestValidator() db.Validator {
	return DefaultUpdateServiceDomainsRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *UpdateServiceDomainsResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *UpdateServiceDomainsResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *UpdateServiceDomainsResponse) DeepCopy() *UpdateServiceDomainsResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &UpdateServiceDomainsResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *UpdateServiceDomainsResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *UpdateServiceDomainsResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return UpdateServiceDomainsResponseValidator().Validate(ctx, m, opts...)
}

type ValidateUpdateServiceDomainsResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateUpdateServiceDomainsResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*UpdateServiceDomainsResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *UpdateServiceDomainsResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultUpdateServiceDomainsResponseValidator = func() *ValidateUpdateServiceDomainsResponse {
	v := &ValidateUpdateServiceDomainsResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func UpdateServiceDomainsResponseValidator() db.Validator {
	return DefaultUpdateServiceDomainsResponseValidator
}
