// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package site_analysis

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

func (m *AnalysisAndAction) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AnalysisAndAction) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AnalysisAndAction) DeepCopy() *AnalysisAndAction {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AnalysisAndAction{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AnalysisAndAction) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AnalysisAndAction) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AnalysisAndActionValidator().Validate(ctx, m, opts...)
}

type ValidateAnalysisAndAction struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAnalysisAndAction) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AnalysisAndAction)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AnalysisAndAction got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["action"]; exists {

		vOpts := append(opts, db.WithValidateField("action"))
		if err := fv(ctx, m.GetAction(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["analysis"]; exists {

		vOpts := append(opts, db.WithValidateField("analysis"))
		if err := fv(ctx, m.GetAnalysis(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAnalysisAndActionValidator = func() *ValidateAnalysisAndAction {
	v := &ValidateAnalysisAndAction{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func AnalysisAndActionValidator() db.Validator {
	return DefaultAnalysisAndActionValidator
}

// augmented methods on protoc/std generated struct

func (m *SiteAnalysisResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *SiteAnalysisResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *SiteAnalysisResponse) DeepCopy() *SiteAnalysisResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &SiteAnalysisResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *SiteAnalysisResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *SiteAnalysisResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return SiteAnalysisResponseValidator().Validate(ctx, m, opts...)
}

type ValidateSiteAnalysisResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateSiteAnalysisResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*SiteAnalysisResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *SiteAnalysisResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["analysis_and_actions"]; exists {

		vOpts := append(opts, db.WithValidateField("analysis_and_actions"))
		for idx, item := range m.GetAnalysisAndActions() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["external_link"]; exists {

		vOpts := append(opts, db.WithValidateField("external_link"))
		if err := fv(ctx, m.GetExternalLink(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["internal_link"]; exists {

		vOpts := append(opts, db.WithValidateField("internal_link"))
		if err := fv(ctx, m.GetInternalLink(), vOpts...); err != nil {
			return err
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
var DefaultSiteAnalysisResponseValidator = func() *ValidateSiteAnalysisResponse {
	v := &ValidateSiteAnalysisResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["external_link"] = ves_io_schema_ai_assistant_common.LinkValidator().Validate

	v.FldValidators["internal_link"] = ves_io_schema_ai_assistant_common.LinkValidator().Validate

	return v
}()

func SiteAnalysisResponseValidator() db.Validator {
	return DefaultSiteAnalysisResponseValidator
}
