// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package geo_config

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

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

func (m *GetSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetPreferredRegionDRefInfo()

}

func (m *GetSpecType) GetPreferredRegionDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetPreferredRegion()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GetSpecType.preferred_region[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "data_privacy.lma_region.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "preferred_region",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetPreferredRegionDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GetSpecType) GetPreferredRegionDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "data_privacy.lma_region.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: data_privacy.lma_region")
	}
	for _, ref := range m.GetPreferredRegion() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateGetSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGetSpecType) PreferredRegionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for preferred_region")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for preferred_region")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated preferred_region")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items preferred_region")
		}
		return nil
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

	if fv, exists := v.FldValidators["preferred_region"]; exists {
		vOpts := append(opts, db.WithValidateField("preferred_region"))
		if err := fv(ctx, m.GetPreferredRegion(), vOpts...); err != nil {
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

	vrhPreferredRegion := v.PreferredRegionValidationRuleHandler
	rulesPreferredRegion := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhPreferredRegion(rulesPreferredRegion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GetSpecType.preferred_region: %s", err)
		panic(errMsg)
	}
	v.FldValidators["preferred_region"] = vFn

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

func (m *GlobalSpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	if m == nil {
		return nil, nil
	}

	return m.GetPreferredRegionDRefInfo()

}

func (m *GlobalSpecType) GetPreferredRegionDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetPreferredRegion()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.preferred_region[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "data_privacy.lma_region.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "preferred_region",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetPreferredRegionDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetPreferredRegionDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "data_privacy.lma_region.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: data_privacy.lma_region")
	}
	for _, ref := range m.GetPreferredRegion() {
		refdEnt, err := d.GetReferredEntry(ctx, refdType, ref, db.WithRefOpOptions(db.OpWithReadRefFromInternalTable()))
		if err != nil {
			return nil, errors.Wrap(err, "Getting referred entry")
		}
		if refdEnt != nil {
			entries = append(entries, refdEnt)
		}
	}

	return entries, nil
}

type ValidateGlobalSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGlobalSpecType) PreferredRegionValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for preferred_region")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema.ObjectRefType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema.ObjectRefTypeValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for preferred_region")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema.ObjectRefType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema.ObjectRefType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal, err := codec.ToJSON(elem, codec.ToWithUseProtoFieldName())
			if err != nil {
				return errors.Wrapf(err, "Converting %v to JSON", elem)
			}
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated preferred_region")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items preferred_region")
		}
		return nil
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

	if fv, exists := v.FldValidators["preferred_region"]; exists {
		vOpts := append(opts, db.WithValidateField("preferred_region"))
		if err := fv(ctx, m.GetPreferredRegion(), vOpts...); err != nil {
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

	vrhPreferredRegion := v.PreferredRegionValidationRuleHandler
	rulesPreferredRegion := map[string]string{
		"ves.io.schema.rules.repeated.max_items": "1",
	}
	vFn, err = vrhPreferredRegion(rulesPreferredRegion)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.preferred_region: %s", err)
		panic(errMsg)
	}
	v.FldValidators["preferred_region"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}

func (m *GetSpecType) fromGlobalSpecType(f *GlobalSpecType, withDeepCopy bool) {
	if f == nil {
		return
	}
	m.PreferredRegion = f.GetPreferredRegion()
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

	f.PreferredRegion = m1.PreferredRegion
}

func (m *GetSpecType) ToGlobalSpecType(f *GlobalSpecType) {
	m.toGlobalSpecType(f, true)
}

func (m *GetSpecType) ToGlobalSpecTypeWithoutDeepCopy(f *GlobalSpecType) {
	m.toGlobalSpecType(f, false)
}
