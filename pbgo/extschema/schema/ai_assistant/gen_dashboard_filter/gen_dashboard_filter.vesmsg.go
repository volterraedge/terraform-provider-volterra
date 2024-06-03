// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package gen_dashboard_filter

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_ai_assistant_common "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ai_assistant/common"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *GenDashboardFilterResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *GenDashboardFilterResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *GenDashboardFilterResponse) DeepCopy() *GenDashboardFilterResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &GenDashboardFilterResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *GenDashboardFilterResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *GenDashboardFilterResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return GenDashboardFilterResponseValidator().Validate(ctx, m, opts...)
}

type ValidateGenDashboardFilterResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateGenDashboardFilterResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*GenDashboardFilterResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *GenDashboardFilterResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["actions"]; exists {

		vOpts := append(opts, db.WithValidateField("actions"))
		if err := fv(ctx, m.GetActions(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["links"]; exists {

		vOpts := append(opts, db.WithValidateField("links"))
		for idx, item := range m.GetLinks() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["summary"]; exists {

		vOpts := append(opts, db.WithValidateField("summary"))
		if err := fv(ctx, m.GetSummary(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultGenDashboardFilterResponseValidator = func() *ValidateGenDashboardFilterResponse {
	v := &ValidateGenDashboardFilterResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["links"] = ves_io_schema_ai_assistant_common.LinkValidator().Validate

	return v
}()

func GenDashboardFilterResponseValidator() db.Validator {
	return DefaultGenDashboardFilterResponseValidator
}
