// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package ticket

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

	return m.GetTicketTrackingSystemDRefInfo()

}

func (m *GlobalSpecType) GetTicketTrackingSystemDRefInfo() ([]db.DRefInfo, error) {
	refs := m.GetTicketTrackingSystem()
	if len(refs) == 0 {
		return nil, nil
	}
	drInfos := make([]db.DRefInfo, 0, len(refs))
	for i, ref := range refs {
		if ref == nil {
			return nil, fmt.Errorf("GlobalSpecType.ticket_tracking_system[%d] has a nil value", i)
		}
		// resolve kind to type if needed at DBObject.GetDRefInfo()
		drInfos = append(drInfos, db.DRefInfo{
			RefdType:   "ticket_tracking_system.Object",
			RefdUID:    ref.Uid,
			RefdTenant: ref.Tenant,
			RefdNS:     ref.Namespace,
			RefdName:   ref.Name,
			DRField:    "ticket_tracking_system",
			Ref:        ref,
		})
	}
	return drInfos, nil

}

// GetTicketTrackingSystemDBEntries returns the db.Entry corresponding to the ObjRefType from the default Table
func (m *GlobalSpecType) GetTicketTrackingSystemDBEntries(ctx context.Context, d db.Interface) ([]db.Entry, error) {
	var entries []db.Entry
	refdType, err := d.TypeForEntryKind("", "", "ticket_tracking_system.Object")
	if err != nil {
		return nil, errors.Wrap(err, "Cannot find type for kind: ticket_tracking_system")
	}
	for _, ref := range m.GetTicketTrackingSystem() {
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

func (v *ValidateGlobalSpecType) ServiceFeatureValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for service_feature")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) ExternalIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for external_id")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) TicketTrackingSystemValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for ticket_tracking_system")
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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for ticket_tracking_system")
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
			return errors.Wrap(err, "repeated ticket_tracking_system")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items ticket_tracking_system")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) AuthorValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for author")
	}

	return validatorFn, nil
}

func (v *ValidateGlobalSpecType) IssueKeyValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for issue_key")
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

	if fv, exists := v.FldValidators["author"]; exists {

		vOpts := append(opts, db.WithValidateField("author"))
		if err := fv(ctx, m.GetAuthor(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["external_id"]; exists {

		vOpts := append(opts, db.WithValidateField("external_id"))
		if err := fv(ctx, m.GetExternalId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["issue_key"]; exists {

		vOpts := append(opts, db.WithValidateField("issue_key"))
		if err := fv(ctx, m.GetIssueKey(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service_feature"]; exists {

		vOpts := append(opts, db.WithValidateField("service_feature"))
		if err := fv(ctx, m.GetServiceFeature(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["ticket_tracking_system"]; exists {
		vOpts := append(opts, db.WithValidateField("ticket_tracking_system"))
		if err := fv(ctx, m.GetTicketTrackingSystem(), vOpts...); err != nil {
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

	vrhServiceFeature := v.ServiceFeatureValidationRuleHandler
	rulesServiceFeature := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhServiceFeature(rulesServiceFeature)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.service_feature: %s", err)
		panic(errMsg)
	}
	v.FldValidators["service_feature"] = vFn

	vrhExternalId := v.ExternalIdValidationRuleHandler
	rulesExternalId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhExternalId(rulesExternalId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.external_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["external_id"] = vFn

	vrhTicketTrackingSystem := v.TicketTrackingSystemValidationRuleHandler
	rulesTicketTrackingSystem := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.max_items": "1",
		"ves.io.schema.rules.repeated.min_items": "1",
	}
	vFn, err = vrhTicketTrackingSystem(rulesTicketTrackingSystem)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.ticket_tracking_system: %s", err)
		panic(errMsg)
	}
	v.FldValidators["ticket_tracking_system"] = vFn

	vrhAuthor := v.AuthorValidationRuleHandler
	rulesAuthor := map[string]string{
		"ves.io.schema.rules.string.email": "true",
	}
	vFn, err = vrhAuthor(rulesAuthor)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.author: %s", err)
		panic(errMsg)
	}
	v.FldValidators["author"] = vFn

	vrhIssueKey := v.IssueKeyValidationRuleHandler
	rulesIssueKey := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhIssueKey(rulesIssueKey)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for GlobalSpecType.issue_key: %s", err)
		panic(errMsg)
	}
	v.FldValidators["issue_key"] = vFn

	return v
}()

func GlobalSpecTypeValidator() db.Validator {
	return DefaultGlobalSpecTypeValidator
}
