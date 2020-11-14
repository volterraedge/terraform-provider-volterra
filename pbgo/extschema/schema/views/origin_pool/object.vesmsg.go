//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.
//
package origin_pool

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

func (m *SpecType) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SpecType) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

// Redact squashes sensitive info in m (in-place)
func (m *SpecType) Redact(ctx context.Context) error {
	// clear fields with confidential option set (at message or field level)
	if m == nil {
		return nil
	}

	if err := m.GetGcSpec().Redact(ctx); err != nil {
		return errors.Wrapf(err, "Redacting SpecType.gc_spec")
	}

	return nil
}

func (m *SpecType) DeepCopy() *SpecType {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SpecType{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SpecType) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SpecType) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SpecTypeValidator().Validate(ctx, m, opts...)
}

func (m *SpecType) GetDRefInfo() ([]db.DRefInfo, error) {
	var drInfos []db.DRefInfo
	if fdrInfos, err := m.GetGcSpecDRefInfo(); err != nil {
		return nil, err
	} else {
		drInfos = append(drInfos, fdrInfos...)
	}

	return drInfos, nil
}

// GetDRefInfo for the field's type
func (m *SpecType) GetGcSpecDRefInfo() ([]db.DRefInfo, error) {
	var (
		drInfos, driSet []db.DRefInfo
		err             error
	)
	_ = driSet
	if m.GcSpec == nil {
		return []db.DRefInfo{}, nil
	}

	driSet, err = m.GcSpec.GetDRefInfo()
	if err != nil {
		return nil, err
	}
	for _, dri := range driSet {
		dri.DRField = "gc_spec." + dri.DRField
		drInfos = append(drInfos, dri)
	}

	return drInfos, err
}

type ValidateSpecType struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSpecType) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SpecType)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SpecType got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["gc_spec"]; exists {

		vOpts := append(opts, db.WithValidateField("gc_spec"))
		if err := fv(ctx, m.GetGcSpec(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultSpecTypeValidator = func() *ValidateSpecType {
	v := &ValidateSpecType{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["gc_spec"] = GlobalSpecTypeValidator().Validate

	return v
}()

func SpecTypeValidator() db.Validator {
	return DefaultSpecTypeValidator
}
