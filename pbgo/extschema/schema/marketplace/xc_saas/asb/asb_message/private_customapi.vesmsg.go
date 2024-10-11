// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package asb_message

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

func (m *Admin) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Admin) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Admin) DeepCopy() *Admin {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Admin{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Admin) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Admin) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return AdminValidator().Validate(ctx, m, opts...)
}

type ValidateAdmin struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateAdmin) FirstNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for first_name")
	}

	return validatorFn, nil
}

func (v *ValidateAdmin) LastNameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for last_name")
	}

	return validatorFn, nil
}

func (v *ValidateAdmin) EmailValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for email")
	}

	return validatorFn, nil
}

func (v *ValidateAdmin) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Admin)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Admin got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["email"]; exists {

		vOpts := append(opts, db.WithValidateField("email"))
		if err := fv(ctx, m.GetEmail(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["first_name"]; exists {

		vOpts := append(opts, db.WithValidateField("first_name"))
		if err := fv(ctx, m.GetFirstName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["last_name"]; exists {

		vOpts := append(opts, db.WithValidateField("last_name"))
		if err := fv(ctx, m.GetLastName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultAdminValidator = func() *ValidateAdmin {
	v := &ValidateAdmin{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhFirstName := v.FirstNameValidationRuleHandler
	rulesFirstName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "256",
		"ves.io.schema.rules.string.min_len":   "1",
	}
	vFn, err = vrhFirstName(rulesFirstName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Admin.first_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["first_name"] = vFn

	vrhLastName := v.LastNameValidationRuleHandler
	rulesLastName := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "256",
		"ves.io.schema.rules.string.min_len":   "1",
	}
	vFn, err = vrhLastName(rulesLastName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Admin.last_name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["last_name"] = vFn

	vrhEmail := v.EmailValidationRuleHandler
	rulesEmail := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "256",
		"ves.io.schema.rules.string.min_len":   "3",
	}
	vFn, err = vrhEmail(rulesEmail)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Admin.email: %s", err)
		panic(errMsg)
	}
	v.FldValidators["email"] = vFn

	return v
}()

func AdminValidator() db.Validator {
	return DefaultAdminValidator
}

// augmented methods on protoc/std generated struct

func (m *CustomerMetadata) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *CustomerMetadata) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *CustomerMetadata) DeepCopy() *CustomerMetadata {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &CustomerMetadata{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *CustomerMetadata) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *CustomerMetadata) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CustomerMetadataValidator().Validate(ctx, m, opts...)
}

type ValidateCustomerMetadata struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCustomerMetadata) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateCustomerMetadata) AddressOneValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for address_one")
	}

	return validatorFn, nil
}

func (v *ValidateCustomerMetadata) CityValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for city")
	}

	return validatorFn, nil
}

func (v *ValidateCustomerMetadata) CountryValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for country")
	}

	return validatorFn, nil
}

func (v *ValidateCustomerMetadata) ZipCodeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for zip_code")
	}

	return validatorFn, nil
}

func (v *ValidateCustomerMetadata) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*CustomerMetadata)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *CustomerMetadata got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["address_one"]; exists {

		vOpts := append(opts, db.WithValidateField("address_one"))
		if err := fv(ctx, m.GetAddressOne(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["address_two"]; exists {

		vOpts := append(opts, db.WithValidateField("address_two"))
		if err := fv(ctx, m.GetAddressTwo(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["city"]; exists {

		vOpts := append(opts, db.WithValidateField("city"))
		if err := fv(ctx, m.GetCity(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["country"]; exists {

		vOpts := append(opts, db.WithValidateField("country"))
		if err := fv(ctx, m.GetCountry(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["zip_code"]; exists {

		vOpts := append(opts, db.WithValidateField("zip_code"))
		if err := fv(ctx, m.GetZipCode(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCustomerMetadataValidator = func() *ValidateCustomerMetadata {
	v := &ValidateCustomerMetadata{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.string.max_len": "256",
		"ves.io.schema.rules.string.min_len": "1",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CustomerMetadata.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhAddressOne := v.AddressOneValidationRuleHandler
	rulesAddressOne := map[string]string{
		"ves.io.schema.rules.string.min_len": "1",
	}
	vFn, err = vrhAddressOne(rulesAddressOne)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CustomerMetadata.address_one: %s", err)
		panic(errMsg)
	}
	v.FldValidators["address_one"] = vFn

	vrhCity := v.CityValidationRuleHandler
	rulesCity := map[string]string{
		"ves.io.schema.rules.string.min_len": "1",
	}
	vFn, err = vrhCity(rulesCity)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CustomerMetadata.city: %s", err)
		panic(errMsg)
	}
	v.FldValidators["city"] = vFn

	vrhCountry := v.CountryValidationRuleHandler
	rulesCountry := map[string]string{
		"ves.io.schema.rules.string.min_len": "1",
	}
	vFn, err = vrhCountry(rulesCountry)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CustomerMetadata.country: %s", err)
		panic(errMsg)
	}
	v.FldValidators["country"] = vFn

	vrhZipCode := v.ZipCodeValidationRuleHandler
	rulesZipCode := map[string]string{
		"ves.io.schema.rules.string.min_len": "5",
	}
	vFn, err = vrhZipCode(rulesZipCode)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for CustomerMetadata.zip_code: %s", err)
		panic(errMsg)
	}
	v.FldValidators["zip_code"] = vFn

	return v
}()

func CustomerMetadataValidator() db.Validator {
	return DefaultCustomerMetadataValidator
}

// augmented methods on protoc/std generated struct

func (m *Entitlement) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Entitlement) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Entitlement) DeepCopy() *Entitlement {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Entitlement{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Entitlement) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Entitlement) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EntitlementValidator().Validate(ctx, m, opts...)
}

type ValidateEntitlement struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEntitlement) EntitlementMetadataValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for entitlement_metadata")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EntitlementMetadataValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateEntitlement) TenantMetadataValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for tenant_metadata")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := TenantMetadataValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateEntitlement) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Entitlement)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Entitlement got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["customer_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("customer_metadata"))
		if err := fv(ctx, m.GetCustomerMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["entitlement_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("entitlement_metadata"))
		if err := fv(ctx, m.GetEntitlementMetadata(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["tenant_metadata"]; exists {

		vOpts := append(opts, db.WithValidateField("tenant_metadata"))
		if err := fv(ctx, m.GetTenantMetadata(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEntitlementValidator = func() *ValidateEntitlement {
	v := &ValidateEntitlement{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhEntitlementMetadata := v.EntitlementMetadataValidationRuleHandler
	rulesEntitlementMetadata := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEntitlementMetadata(rulesEntitlementMetadata)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Entitlement.entitlement_metadata: %s", err)
		panic(errMsg)
	}
	v.FldValidators["entitlement_metadata"] = vFn

	vrhTenantMetadata := v.TenantMetadataValidationRuleHandler
	rulesTenantMetadata := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhTenantMetadata(rulesTenantMetadata)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Entitlement.tenant_metadata: %s", err)
		panic(errMsg)
	}
	v.FldValidators["tenant_metadata"] = vFn

	v.FldValidators["customer_metadata"] = CustomerMetadataValidator().Validate

	return v
}()

func EntitlementValidator() db.Validator {
	return DefaultEntitlementValidator
}

// augmented methods on protoc/std generated struct

func (m *EntitlementMetadata) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *EntitlementMetadata) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *EntitlementMetadata) DeepCopy() *EntitlementMetadata {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &EntitlementMetadata{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *EntitlementMetadata) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *EntitlementMetadata) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return EntitlementMetadataValidator().Validate(ctx, m, opts...)
}

type ValidateEntitlementMetadata struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateEntitlementMetadata) EntitledSkusValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepStringItemRules(rules)
	itemValFn, err := db.NewStringValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Item ValidationRuleHandler for entitled_skus")
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
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for entitled_skus")
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
			return errors.Wrap(err, "repeated entitled_skus")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items entitled_skus")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateEntitlementMetadata) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*EntitlementMetadata)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *EntitlementMetadata got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["entitled_skus"]; exists {
		vOpts := append(opts, db.WithValidateField("entitled_skus"))
		if err := fv(ctx, m.GetEntitledSkus(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultEntitlementMetadataValidator = func() *ValidateEntitlementMetadata {
	v := &ValidateEntitlementMetadata{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhEntitledSkus := v.EntitledSkusValidationRuleHandler
	rulesEntitledSkus := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.min_items": "1",
	}
	vFn, err = vrhEntitledSkus(rulesEntitledSkus)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for EntitlementMetadata.entitled_skus: %s", err)
		panic(errMsg)
	}
	v.FldValidators["entitled_skus"] = vFn

	return v
}()

func EntitlementMetadataValidator() db.Validator {
	return DefaultEntitlementMetadataValidator
}

// augmented methods on protoc/std generated struct

func (m *Order) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Order) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Order) DeepCopy() *Order {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Order{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Order) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Order) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return OrderValidator().Validate(ctx, m, opts...)
}

type ValidateOrder struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateOrder) ServiceValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ServiceType)
		return int32(i)
	}
	// ServiceType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ServiceType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for service")
	}

	return validatorFn, nil
}

func (v *ValidateOrder) AccountIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for account_id")
	}

	return validatorFn, nil
}

func (v *ValidateOrder) SubscriptionIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for subscription_id")
	}

	return validatorFn, nil
}

func (v *ValidateOrder) EntitlementIdValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for entitlement_id")
	}

	return validatorFn, nil
}

func (v *ValidateOrder) ActionTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ActionType)
		return int32(i)
	}
	// ActionType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ActionType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for action_type")
	}

	return validatorFn, nil
}

func (v *ValidateOrder) EntitlementValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for entitlement")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := EntitlementValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateOrder) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Order)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Order got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["account_id"]; exists {

		vOpts := append(opts, db.WithValidateField("account_id"))
		if err := fv(ctx, m.GetAccountId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["action_type"]; exists {

		vOpts := append(opts, db.WithValidateField("action_type"))
		if err := fv(ctx, m.GetActionType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["entitlement"]; exists {

		vOpts := append(opts, db.WithValidateField("entitlement"))
		if err := fv(ctx, m.GetEntitlement(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["entitlement_id"]; exists {

		vOpts := append(opts, db.WithValidateField("entitlement_id"))
		if err := fv(ctx, m.GetEntitlementId(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["order_type"]; exists {

		vOpts := append(opts, db.WithValidateField("order_type"))
		if err := fv(ctx, m.GetOrderType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service"]; exists {

		vOpts := append(opts, db.WithValidateField("service"))
		if err := fv(ctx, m.GetService(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["subscription_id"]; exists {

		vOpts := append(opts, db.WithValidateField("subscription_id"))
		if err := fv(ctx, m.GetSubscriptionId(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultOrderValidator = func() *ValidateOrder {
	v := &ValidateOrder{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhService := v.ServiceValidationRuleHandler
	rulesService := map[string]string{
		"ves.io.schema.rules.enum.defined_only": "true",
		"ves.io.schema.rules.message.required":  "true",
	}
	vFn, err = vrhService(rulesService)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Order.service: %s", err)
		panic(errMsg)
	}
	v.FldValidators["service"] = vFn

	vrhAccountId := v.AccountIdValidationRuleHandler
	rulesAccountId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "256",
		"ves.io.schema.rules.string.min_len":   "1",
	}
	vFn, err = vrhAccountId(rulesAccountId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Order.account_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["account_id"] = vFn

	vrhSubscriptionId := v.SubscriptionIdValidationRuleHandler
	rulesSubscriptionId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "256",
		"ves.io.schema.rules.string.min_len":   "1",
	}
	vFn, err = vrhSubscriptionId(rulesSubscriptionId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Order.subscription_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["subscription_id"] = vFn

	vrhEntitlementId := v.EntitlementIdValidationRuleHandler
	rulesEntitlementId := map[string]string{
		"ves.io.schema.rules.message.required": "true",
		"ves.io.schema.rules.string.max_len":   "256",
		"ves.io.schema.rules.string.min_len":   "1",
	}
	vFn, err = vrhEntitlementId(rulesEntitlementId)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Order.entitlement_id: %s", err)
		panic(errMsg)
	}
	v.FldValidators["entitlement_id"] = vFn

	vrhActionType := v.ActionTypeValidationRuleHandler
	rulesActionType := map[string]string{
		"ves.io.schema.rules.enum.defined_only": "true",
		"ves.io.schema.rules.message.required":  "true",
	}
	vFn, err = vrhActionType(rulesActionType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Order.action_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["action_type"] = vFn

	vrhEntitlement := v.EntitlementValidationRuleHandler
	rulesEntitlement := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhEntitlement(rulesEntitlement)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Order.entitlement: %s", err)
		panic(errMsg)
	}
	v.FldValidators["entitlement"] = vFn

	return v
}()

func OrderValidator() db.Validator {
	return DefaultOrderValidator
}

// augmented methods on protoc/std generated struct

func (m *RegisterXCSaaSRequest) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RegisterXCSaaSRequest) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RegisterXCSaaSRequest) DeepCopy() *RegisterXCSaaSRequest {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RegisterXCSaaSRequest{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RegisterXCSaaSRequest) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RegisterXCSaaSRequest) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RegisterXCSaaSRequestValidator().Validate(ctx, m, opts...)
}

type ValidateRegisterXCSaaSRequest struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRegisterXCSaaSRequest) DocumentTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(DocumentType)
		return int32(i)
	}
	// DocumentType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, DocumentType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for document_type")
	}

	return validatorFn, nil
}

func (v *ValidateRegisterXCSaaSRequest) ServiceOrderTypeValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	var conv db.EnumConvFn
	conv = func(v interface{}) int32 {
		i := v.(ServiceOrderType)
		return int32(i)
	}
	// ServiceOrderType_name is generated in .pb.go
	validatorFn, err := db.NewEnumValidationRuleHandler(rules, ServiceOrderType_name, conv)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for service_order_type")
	}

	return validatorFn, nil
}

func (v *ValidateRegisterXCSaaSRequest) OrderValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	reqdValidatorFn, err := db.NewMessageValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "MessageValidationRuleHandler for order")
	}
	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		if err := reqdValidatorFn(ctx, val, opts...); err != nil {
			return err
		}

		if err := OrderValidator().Validate(ctx, val, opts...); err != nil {
			return err
		}

		return nil
	}

	return validatorFn, nil
}

func (v *ValidateRegisterXCSaaSRequest) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RegisterXCSaaSRequest)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RegisterXCSaaSRequest got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["document_type"]; exists {

		vOpts := append(opts, db.WithValidateField("document_type"))
		if err := fv(ctx, m.GetDocumentType(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["document_version"]; exists {

		vOpts := append(opts, db.WithValidateField("document_version"))
		if err := fv(ctx, m.GetDocumentVersion(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["order"]; exists {

		vOpts := append(opts, db.WithValidateField("order"))
		if err := fv(ctx, m.GetOrder(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["service_order_type"]; exists {

		vOpts := append(opts, db.WithValidateField("service_order_type"))
		if err := fv(ctx, m.GetServiceOrderType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRegisterXCSaaSRequestValidator = func() *ValidateRegisterXCSaaSRequest {
	v := &ValidateRegisterXCSaaSRequest{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhDocumentType := v.DocumentTypeValidationRuleHandler
	rulesDocumentType := map[string]string{
		"ves.io.schema.rules.enum.defined_only": "true",
		"ves.io.schema.rules.message.required":  "true",
	}
	vFn, err = vrhDocumentType(rulesDocumentType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RegisterXCSaaSRequest.document_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["document_type"] = vFn

	vrhServiceOrderType := v.ServiceOrderTypeValidationRuleHandler
	rulesServiceOrderType := map[string]string{
		"ves.io.schema.rules.enum.defined_only": "true",
		"ves.io.schema.rules.message.required":  "true",
	}
	vFn, err = vrhServiceOrderType(rulesServiceOrderType)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RegisterXCSaaSRequest.service_order_type: %s", err)
		panic(errMsg)
	}
	v.FldValidators["service_order_type"] = vFn

	vrhOrder := v.OrderValidationRuleHandler
	rulesOrder := map[string]string{
		"ves.io.schema.rules.message.required": "true",
	}
	vFn, err = vrhOrder(rulesOrder)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for RegisterXCSaaSRequest.order: %s", err)
		panic(errMsg)
	}
	v.FldValidators["order"] = vFn

	return v
}()

func RegisterXCSaaSRequestValidator() db.Validator {
	return DefaultRegisterXCSaaSRequestValidator
}

// augmented methods on protoc/std generated struct

func (m *RegisterXCSaaSResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *RegisterXCSaaSResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *RegisterXCSaaSResponse) DeepCopy() *RegisterXCSaaSResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &RegisterXCSaaSResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *RegisterXCSaaSResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *RegisterXCSaaSResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RegisterXCSaaSResponseValidator().Validate(ctx, m, opts...)
}

type ValidateRegisterXCSaaSResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRegisterXCSaaSResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*RegisterXCSaaSResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *RegisterXCSaaSResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRegisterXCSaaSResponseValidator = func() *ValidateRegisterXCSaaSResponse {
	v := &ValidateRegisterXCSaaSResponse{FldValidators: map[string]db.ValidatorFunc{}}

	return v
}()

func RegisterXCSaaSResponseValidator() db.Validator {
	return DefaultRegisterXCSaaSResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *TenantMetadata) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *TenantMetadata) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *TenantMetadata) DeepCopy() *TenantMetadata {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &TenantMetadata{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *TenantMetadata) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *TenantMetadata) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return TenantMetadataValidator().Validate(ctx, m, opts...)
}

type ValidateTenantMetadata struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateTenantMetadata) NameValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	validatorFn, err := db.NewStringValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "ValidationRuleHandler for name")
	}

	return validatorFn, nil
}

func (v *ValidateTenantMetadata) AdminsValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for admins")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*Admin, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := AdminValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for admins")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*Admin)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*Admin, got %T", val)
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
			return errors.Wrap(err, "repeated admins")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items admins")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateTenantMetadata) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*TenantMetadata)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *TenantMetadata got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["admins"]; exists {
		vOpts := append(opts, db.WithValidateField("admins"))
		if err := fv(ctx, m.GetAdmins(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["name"]; exists {

		vOpts := append(opts, db.WithValidateField("name"))
		if err := fv(ctx, m.GetName(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultTenantMetadataValidator = func() *ValidateTenantMetadata {
	v := &ValidateTenantMetadata{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhName := v.NameValidationRuleHandler
	rulesName := map[string]string{
		"ves.io.schema.rules.string.max_len":         "17",
		"ves.io.schema.rules.string.min_len":         "1",
		"ves.io.schema.rules.string.not_in":          "[\"ves\",\"volterra\",\"ves-io\",\"f5xc\"]",
		"ves.io.schema.rules.string.ves_object_name": "true",
	}
	vFn, err = vrhName(rulesName)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for TenantMetadata.name: %s", err)
		panic(errMsg)
	}
	v.FldValidators["name"] = vFn

	vrhAdmins := v.AdminsValidationRuleHandler
	rulesAdmins := map[string]string{
		"ves.io.schema.rules.message.required":   "true",
		"ves.io.schema.rules.repeated.min_items": "1",
		"ves.io.schema.rules.repeated.unique":    "true",
	}
	vFn, err = vrhAdmins(rulesAdmins)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for TenantMetadata.admins: %s", err)
		panic(errMsg)
	}
	v.FldValidators["admins"] = vFn

	return v
}()

func TenantMetadataValidator() db.Validator {
	return DefaultTenantMetadataValidator
}