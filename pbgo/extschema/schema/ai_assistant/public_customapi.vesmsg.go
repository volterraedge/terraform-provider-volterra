// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package ai_assistant

import (
	"context"
	"fmt"
	"strings"

	"github.com/gogo/protobuf/proto"

	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_ai_assistant_gen_dashboard_filter "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ai_assistant/gen_dashboard_filter"
	ves_io_schema_ai_assistant_list "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ai_assistant/list"
	ves_io_schema_ai_assistant_site_analysis "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ai_assistant/site_analysis"
	ves_io_schema_ai_assistant_widget "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/ai_assistant/widget"
)

var (
	// dummy imports in case file has no message with Refs
	_ db.Interface
	_ = errors.Wrap
	_ = strings.Split
)

// augmented methods on protoc/std generated struct

func (m *AIAssistantQueryFeedbackRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AIAssistantQueryFeedbackRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AIAssistantQueryFeedbackRequest) DeepCopy() *AIAssistantQueryFeedbackRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AIAssistantQueryFeedbackRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AIAssistantQueryFeedbackRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AIAssistantQueryFeedbackRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AIAssistantQueryFeedbackRequestValidator().Validate(ctx, m, opts...)
}

type ValidateAIAssistantQueryFeedbackRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAIAssistantQueryFeedbackRequest) QueryIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for query_id")
	}

	return validatorFn, nil
}

func (v *ValidateAIAssistantQueryFeedbackRequest) CommentValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for comment")
	}

	return validatorFn, nil
}

func (v *ValidateAIAssistantQueryFeedbackRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AIAssistantQueryFeedbackRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AIAssistantQueryFeedbackRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["comment"]; exists {

		vOpts := append(opts, db.WithValidateField("comment"))
		if err := fv(ctx, m.GetComment(), vOpts...); err != nil {
			return err
		}

	}

	switch m.GetFeedbackChoice().(type) {
	case *AIAssistantQueryFeedbackRequest_PositiveFeedback:
		if fv, exists := v.FldValidators["feedback_choice.positive_feedback"]; exists {
			val := m.GetFeedbackChoice().(*AIAssistantQueryFeedbackRequest_PositiveFeedback).PositiveFeedback
			vOpts := append(opts,
				db.WithValidateField("feedback_choice"),
				db.WithValidateField("positive_feedback"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AIAssistantQueryFeedbackRequest_NegativeFeedback:
		if fv, exists := v.FldValidators["feedback_choice.negative_feedback"]; exists {
			val := m.GetFeedbackChoice().(*AIAssistantQueryFeedbackRequest_NegativeFeedback).NegativeFeedback
			vOpts := append(opts,
				db.WithValidateField("feedback_choice"),
				db.WithValidateField("negative_feedback"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["namespace"]; exists {

		vOpts := append(opts, db.WithValidateField("namespace"))
		if err := fv(ctx, m.GetNamespace(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query"]; exists {

		vOpts := append(opts, db.WithValidateField("query"))
		if err := fv(ctx, m.GetQuery(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query_id"]; exists {

		vOpts := append(opts, db.WithValidateField("query_id"))
		if err := fv(ctx, m.GetQueryId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAIAssistantQueryFeedbackRequestValidator = func() *ValidateAIAssistantQueryFeedbackRequest {
	v := &ValidateAIAssistantQueryFeedbackRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhQueryId := v.QueryIdValidationRuleHandler
	rulesQueryId := map[string]string{
		"ves.io.schema.rules.string.pattern": "^([0-9a-f]{8}-){1}([0-9a-f]{4}-){3}([0-9a-f]{12})$",
	}
	vFn, err = vrhQueryId(rulesQueryId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AIAssistantQueryFeedbackRequest.query_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["query_id"] = vFn

	vrhComment := v.CommentValidationRuleHandler
	rulesComment := map[string]string{
		"ves.io.schema.rules.string.max_len": "4096",
	}
	vFn, err = vrhComment(rulesComment)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AIAssistantQueryFeedbackRequest.comment: %s", err)
		panic(errMsg)
	}
	v.FldValidators["comment"] = vFn

	v.FldValidators["feedback_choice.negative_feedback"] = NegativeFeedbackDetailsValidator().Validate

	return v
}()

func AIAssistantQueryFeedbackRequestValidator() db.Validator {
	return DefaultAIAssistantQueryFeedbackRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *AIAssistantQueryRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AIAssistantQueryRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AIAssistantQueryRequest) DeepCopy() *AIAssistantQueryRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AIAssistantQueryRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AIAssistantQueryRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AIAssistantQueryRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AIAssistantQueryRequestValidator().Validate(ctx, m, opts...)
}

type ValidateAIAssistantQueryRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAIAssistantQueryRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AIAssistantQueryRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AIAssistantQueryRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["current_query"]; exists {

		vOpts := append(opts, db.WithValidateField("current_query"))
		if err := fv(ctx, m.GetCurrentQuery(), vOpts...); err != nil {
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
var DefaultAIAssistantQueryRequestValidator = func() *ValidateAIAssistantQueryRequest {
	v := &ValidateAIAssistantQueryRequest{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func AIAssistantQueryRequestValidator() db.Validator {
	return DefaultAIAssistantQueryRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *AIAssistantQueryResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *AIAssistantQueryResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *AIAssistantQueryResponse) DeepCopy() *AIAssistantQueryResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &AIAssistantQueryResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *AIAssistantQueryResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *AIAssistantQueryResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AIAssistantQueryResponseValidator().Validate(ctx, m, opts...)
}

type ValidateAIAssistantQueryResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAIAssistantQueryResponse) QueryIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for query_id")
	}

	return validatorFn, nil
}

func (v *ValidateAIAssistantQueryResponse) FollowUpQueriesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for follow_up_queries")
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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for follow_up_queries")
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
			return errors.Wrap(err, "repeated follow_up_queries")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items follow_up_queries")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateAIAssistantQueryResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*AIAssistantQueryResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *AIAssistantQueryResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["current_query"]; exists {

		vOpts := append(opts, db.WithValidateField("current_query"))
		if err := fv(ctx, m.GetCurrentQuery(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["follow_up_queries"]; exists {
		vOpts := append(opts, db.WithValidateField("follow_up_queries"))
		if err := fv(ctx, m.GetFollowUpQueries(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["query_id"]; exists {

		vOpts := append(opts, db.WithValidateField("query_id"))
		if err := fv(ctx, m.GetQueryId(), vOpts...); err != nil {
			return err
		}

	}

	switch m.GetResponseChoice().(type) {
	case *AIAssistantQueryResponse_ExplainLog:
		if fv, exists := v.FldValidators["response_choice.explain_log"]; exists {
			val := m.GetResponseChoice().(*AIAssistantQueryResponse_ExplainLog).ExplainLog
			vOpts := append(opts,
				db.WithValidateField("response_choice"),
				db.WithValidateField("explain_log"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AIAssistantQueryResponse_GenDashboardFilter:
		if fv, exists := v.FldValidators["response_choice.gen_dashboard_filter"]; exists {
			val := m.GetResponseChoice().(*AIAssistantQueryResponse_GenDashboardFilter).GenDashboardFilter
			vOpts := append(opts,
				db.WithValidateField("response_choice"),
				db.WithValidateField("gen_dashboard_filter"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AIAssistantQueryResponse_GenericResponse:
		if fv, exists := v.FldValidators["response_choice.generic_response"]; exists {
			val := m.GetResponseChoice().(*AIAssistantQueryResponse_GenericResponse).GenericResponse
			vOpts := append(opts,
				db.WithValidateField("response_choice"),
				db.WithValidateField("generic_response"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AIAssistantQueryResponse_SiteAnalysisResponse:
		if fv, exists := v.FldValidators["response_choice.site_analysis_response"]; exists {
			val := m.GetResponseChoice().(*AIAssistantQueryResponse_SiteAnalysisResponse).SiteAnalysisResponse
			vOpts := append(opts,
				db.WithValidateField("response_choice"),
				db.WithValidateField("site_analysis_response"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AIAssistantQueryResponse_WidgetResponse:
		if fv, exists := v.FldValidators["response_choice.widget_response"]; exists {
			val := m.GetResponseChoice().(*AIAssistantQueryResponse_WidgetResponse).WidgetResponse
			vOpts := append(opts,
				db.WithValidateField("response_choice"),
				db.WithValidateField("widget_response"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}
	case *AIAssistantQueryResponse_ListResponse:
		if fv, exists := v.FldValidators["response_choice.list_response"]; exists {
			val := m.GetResponseChoice().(*AIAssistantQueryResponse_ListResponse).ListResponse
			vOpts := append(opts,
				db.WithValidateField("response_choice"),
				db.WithValidateField("list_response"),
			)
			if err := fv(ctx, val, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAIAssistantQueryResponseValidator = func() *ValidateAIAssistantQueryResponse {
	v := &ValidateAIAssistantQueryResponse{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhQueryId := v.QueryIdValidationRuleHandler
	rulesQueryId := map[string]string{
		"ves.io.schema.rules.string.pattern": "^([0-9a-f]{8}-){1}([0-9a-f]{4}-){3}([0-9a-f]{12})$",
	}
	vFn, err = vrhQueryId(rulesQueryId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AIAssistantQueryResponse.query_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["query_id"] = vFn

	vrhFollowUpQueries := v.FollowUpQueriesValidationRuleHandler
	rulesFollowUpQueries := map[string]string{
		"ves.io.schema.rules.repeated.unique": "true",
	}
	vFn, err = vrhFollowUpQueries(rulesFollowUpQueries)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for AIAssistantQueryResponse.follow_up_queries: %s", err)
		panic(errMsg)
	}
	v.FldValidators["follow_up_queries"] = vFn

	v.FldValidators["response_choice.gen_dashboard_filter"] = ves_io_schema_ai_assistant_gen_dashboard_filter.GenDashboardFilterResponseValidator().Validate
	v.FldValidators["response_choice.site_analysis_response"] = ves_io_schema_ai_assistant_site_analysis.SiteAnalysisResponseValidator().Validate
	v.FldValidators["response_choice.widget_response"] = ves_io_schema_ai_assistant_widget.WidgetResponseValidator().Validate
	v.FldValidators["response_choice.list_response"] = ves_io_schema_ai_assistant_list.ListResponseValidator().Validate

	return v
}()

func AIAssistantQueryResponseValidator() db.Validator {
	return DefaultAIAssistantQueryResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *NegativeFeedbackDetails) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *NegativeFeedbackDetails) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *NegativeFeedbackDetails) DeepCopy() *NegativeFeedbackDetails {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &NegativeFeedbackDetails{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *NegativeFeedbackDetails) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *NegativeFeedbackDetails) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return NegativeFeedbackDetailsValidator().Validate(ctx, m, opts...)
}

type ValidateNegativeFeedbackDetails struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateNegativeFeedbackDetails) RemarksValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepEnumItemRules(rules)
	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(NegativeFeedbackType)
		return int32(i)
	}
	// NegativeFeedbackType_name is generated in .pb.go
	itemValFn, err := db.NewEnumValidationRuleHandler(itemRules, NegativeFeedbackType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for remarks")
	}
	itemsValidatorFn := func(ctx context.Context, elems []NegativeFeedbackType, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for remarks")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]NegativeFeedbackType)
		if !ok {
			return fmt.Errorf("Repeated validation expected []NegativeFeedbackType, got %T", val)
		}
		l := []string{}
		for _, elem := range elems {
			strVal := fmt.Sprintf("%v", elem)
			l = append(l, strVal)
		}
		if err := repValFn(ctx, l, opts...); err != nil {
			return errors.Wrap(err, "repeated remarks")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items remarks")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateNegativeFeedbackDetails) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*NegativeFeedbackDetails)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *NegativeFeedbackDetails got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["remarks"]; exists {
		vOpts := append(opts, db.WithValidateField("remarks"))
		if err := fv(ctx, m.GetRemarks(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultNegativeFeedbackDetailsValidator = func() *ValidateNegativeFeedbackDetails {
	v := &ValidateNegativeFeedbackDetails{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhRemarks := v.RemarksValidationRuleHandler
	rulesRemarks := map[string]string{
		"ves.io.schema.rules.repeated.items.enum.defined_only": "true",
		"ves.io.schema.rules.repeated.unique":                  "true",
	}
	vFn, err = vrhRemarks(rulesRemarks)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for NegativeFeedbackDetails.remarks: %s", err)
		panic(errMsg)
	}
	v.FldValidators["remarks"] = vFn

	return v
}()

func NegativeFeedbackDetailsValidator() db.Validator {
	return DefaultNegativeFeedbackDetailsValidator
}
