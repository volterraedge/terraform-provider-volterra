// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package filter_set

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

func (m *FindFilterSetsReq) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FindFilterSetsReq) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FindFilterSetsReq) DeepCopy() *FindFilterSetsReq {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FindFilterSetsReq{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FindFilterSetsReq) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FindFilterSetsReq) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FindFilterSetsReqValidator().Validate(ctx, m, opts...)
}

type ValidateFindFilterSetsReq struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFindFilterSetsReq) NamespaceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for namespace")
	}

	return validatorFn, nil
}

func (v *ValidateFindFilterSetsReq) ContextKeysValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for context_keys")
	}
	itemsValidatorFn := func(ctx context.Context, elems []string, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for context_keys")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]string)
		if !ok {
			return fmt.Errorf("Repeated validation expected []string, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated context_keys")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items context_keys")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateFindFilterSetsReq) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FindFilterSetsReq)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FindFilterSetsReq got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["context_keys"]; exists {
		vOpts := append(opts, db.WithValidateField("context_keys"))
		if err := fv(ctx, m.GetContextKeys(), vOpts...); err != nil {
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
var DefaultFindFilterSetsReqValidator = func() *ValidateFindFilterSetsReq {
	v := &ValidateFindFilterSetsReq{FldValidators: map[string]db.ValidatorFunc{}}

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
		errMsg := fmt.Sprintf("ValidationRuleHandler for FindFilterSetsReq.namespace: %s", err)
		panic(errMsg)
	}
	v.FldValidators["namespace"] = vFn

	vrhContextKeys := v.ContextKeysValidationRuleHandler
	rulesContextKeys := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.min_items": "1",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhContextKeys(rulesContextKeys)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for FindFilterSetsReq.context_keys: %s", err)
		panic(errMsg)
	}
	v.FldValidators["context_keys"] = vFn

	return v
}()

func FindFilterSetsReqValidator() db.Validator {
	return DefaultFindFilterSetsReqValidator
}

// augmented methods on protoc/std generated struct

func (m *FindFilterSetsRsp) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *FindFilterSetsRsp) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *FindFilterSetsRsp) DeepCopy() *FindFilterSetsRsp {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &FindFilterSetsRsp{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *FindFilterSetsRsp) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *FindFilterSetsRsp) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return FindFilterSetsRspValidator().Validate(ctx, m, opts...)
}

func (m *FindFilterSetsRsp) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return nil, nil

}

type ValidateFindFilterSetsRsp struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateFindFilterSetsRsp) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*FindFilterSetsRsp)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *FindFilterSetsRsp got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["filter_sets"]; exists {

		vOpts := append(opts, db.WithValidateField("filter_sets"))
		for idx, item := range m.GetFilterSets() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultFindFilterSetsRspValidator = func() *ValidateFindFilterSetsRsp {
	v := &ValidateFindFilterSetsRsp{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["filter_sets"] = ObjectValidator().Validate

	return v
}()

func FindFilterSetsRspValidator() db.Validator {
	return DefaultFindFilterSetsRspValidator
}