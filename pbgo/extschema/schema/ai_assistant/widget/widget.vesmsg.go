// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package widget

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

func (m *Cell) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Cell) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Cell) DeepCopy() *Cell {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Cell{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Cell) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Cell) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return CellValidator().Validate(ctx, m, opts...)
}

type ValidateCell struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateCell) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Cell)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Cell got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["link"]; exists {

		vOpts := append(opts, db.WithValidateField("link"))
		if err := fv(ctx, m.GetLink(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["properties"]; exists {

		vOpts := append(opts, db.WithValidateField("properties"))
		if err := fv(ctx, m.GetProperties(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["value"]; exists {

		vOpts := append(opts, db.WithValidateField("value"))
		if err := fv(ctx, m.GetValue(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultCellValidator = func() *ValidateCell {
	v := &ValidateCell{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["link"] = ves_io_schema_ai_assistant_common.LinkValidator().Validate

	return v
}()

func CellValidator() db.Validator {
	return DefaultCellValidator
}

// augmented methods on protoc/std generated struct

func (m *Row) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Row) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Row) DeepCopy() *Row {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Row{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Row) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Row) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return RowValidator().Validate(ctx, m, opts...)
}

type ValidateRow struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateRow) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Row)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Row got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["cells"]; exists {

		vOpts := append(opts, db.WithValidateField("cells"))
		for idx, item := range m.GetCells() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultRowValidator = func() *ValidateRow {
	v := &ValidateRow{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["cells"] = CellValidator().Validate

	return v
}()

func RowValidator() db.Validator {
	return DefaultRowValidator
}

// augmented methods on protoc/std generated struct

func (m *Table) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *Table) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *Table) DeepCopy() *Table {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &Table{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *Table) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *Table) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return TableValidator().Validate(ctx, m, opts...)
}

type ValidateTable struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateTable) FieldPropertiesValidationRuleHandler(rules map[string]string) (db.ValidatorFunc, error) {

	itemRules := db.GetRepMessageItemRules(rules)
	itemValFn, err := db.NewMessageValidationRuleHandler(itemRules)
	if err != nil {
		return nil, errors.Wrap(err, "Message ValidationRuleHandler for field_properties")
	}
	itemsValidatorFn := func(ctx context.Context, elems []*ves_io_schema_ai_assistant_common.FieldProperties, opts ...db.ValidateOpt) error {
		for i, el := range elems {
			if err := itemValFn(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
			if err := ves_io_schema_ai_assistant_common.FieldPropertiesValidator().Validate(ctx, el, opts...); err != nil {
				return errors.Wrap(err, fmt.Sprintf("element %d", i))
			}
		}
		return nil
	}
	repValFn, err := db.NewRepeatedValidationRuleHandler(rules)
	if err != nil {
		return nil, errors.Wrap(err, "Repeated ValidationRuleHandler for field_properties")
	}

	validatorFn := func(ctx context.Context, val interface{}, opts ...db.ValidateOpt) error {
		elems, ok := val.([]*ves_io_schema_ai_assistant_common.FieldProperties)
		if !ok {
			return fmt.Errorf("Repeated validation expected []*ves_io_schema_ai_assistant_common.FieldProperties, got %T", val)
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
			return errors.Wrap(err, "repeated field_properties")
		}
		if err := itemsValidatorFn(ctx, elems, opts...); err != nil {
			return errors.Wrap(err, "items field_properties")
		}
		return nil
	}

	return validatorFn, nil
}

func (v *ValidateTable) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*Table)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *Table got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["field_properties"]; exists {
		vOpts := append(opts, db.WithValidateField("field_properties"))
		if err := fv(ctx, m.GetFieldProperties(), vOpts...); err != nil {
			return err
		}

	}

	if fv, exists := v.FldValidators["rows"]; exists {

		vOpts := append(opts, db.WithValidateField("rows"))
		for idx, item := range m.GetRows() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["widget_type"]; exists {

		vOpts := append(opts, db.WithValidateField("widget_type"))
		if err := fv(ctx, m.GetWidgetType(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultTableValidator = func() *ValidateTable {
	v := &ValidateTable{FldValidators: map[string]db.ValidatorFunc{}}

	var (
		err error
		vFn db.ValidatorFunc
	)
	_, _ = err, vFn
	vFnMap := map[string]db.ValidatorFunc{}
	_ = vFnMap

	vrhFieldProperties := v.FieldPropertiesValidationRuleHandler
	rulesFieldProperties := map[string]string{
		"ves.io.schema.rules.repeated.unique": "true",
	}
	vFn, err = vrhFieldProperties(rulesFieldProperties)
	if err != nil {
		errMsg := fmt.Sprintf("ValidationRuleHandler for Table.field_properties: %s", err)
		panic(errMsg)
	}
	v.FldValidators["field_properties"] = vFn

	v.FldValidators["rows"] = RowValidator().Validate

	return v
}()

func TableValidator() db.Validator {
	return DefaultTableValidator
}

// augmented methods on protoc/std generated struct

func (m *WidgetResponse) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *WidgetResponse) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *WidgetResponse) DeepCopy() *WidgetResponse {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &WidgetResponse{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *WidgetResponse) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *WidgetResponse) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return WidgetResponseValidator().Validate(ctx, m, opts...)
}

type ValidateWidgetResponse struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateWidgetResponse) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*WidgetResponse)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *WidgetResponse got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["item_links"]; exists {

		vOpts := append(opts, db.WithValidateField("item_links"))
		for idx, item := range m.GetItemLinks() {
			vOpts := append(vOpts, db.WithValidateRepItem(idx), db.WithValidateIsRepItem(true))
			if err := fv(ctx, item, vOpts...); err != nil {
				return err
			}
		}

	}

	if fv, exists := v.FldValidators["items"]; exists {

		vOpts := append(opts, db.WithValidateField("items"))
		for idx, item := range m.GetItems() {
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
var DefaultWidgetResponseValidator = func() *ValidateWidgetResponse {
	v := &ValidateWidgetResponse{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["items"] = WidgetViewValidator().Validate

	v.FldValidators["item_links"] = ves_io_schema_ai_assistant_common.LinkValidator().Validate

	return v
}()

func WidgetResponseValidator() db.Validator {
	return DefaultWidgetResponseValidator
}

// augmented methods on protoc/std generated struct

func (m *WidgetView) ToJSON() (string, error) {
	return codec.ToJSON(m)
}

func (m *WidgetView) ToYAML() (string, error) {
	return codec.ToYAML(m)
}

func (m *WidgetView) DeepCopy() *WidgetView {
	if m == nil {
		return nil
	}
	ser, err := m.Marshal()
	if err != nil {
		return nil
	}
	c := &WidgetView{}
	err = c.Unmarshal(ser)
	if err != nil {
		return nil
	}
	return c
}

func (m *WidgetView) DeepCopyProto() proto.Message {
	if m == nil {
		return nil
	}
	return m.DeepCopy()
}

func (m *WidgetView) Validate(ctx context.Context, opts ...db.ValidateOpt) error {
	return WidgetViewValidator().Validate(ctx, m, opts...)
}

type ValidateWidgetView struct {
	FldValidators map[string]db.ValidatorFunc
}

func (v *ValidateWidgetView) Validate(ctx context.Context, pm interface{}, opts ...db.ValidateOpt) error {
	m, ok := pm.(*WidgetView)
	if !ok {
		switch t := pm.(type) {
		case nil:
			return nil
		default:
			return fmt.Errorf("Expected type *WidgetView got type %s", t)
		}
	}
	if m == nil {
		return nil
	}

	if fv, exists := v.FldValidators["item"]; exists {

		vOpts := append(opts, db.WithValidateField("item"))
		if err := fv(ctx, m.GetItem(), vOpts...); err != nil {
			return err
		}

	}

	return nil
}

// Well-known symbol for default validator implementation
var DefaultWidgetViewValidator = func() *ValidateWidgetView {
	v := &ValidateWidgetView{FldValidators: map[string]db.ValidatorFunc{}}

	v.FldValidators["item"] = TableValidator().Validate

	return v
}()

func WidgetViewValidator() db.Validator {
	return DefaultWidgetViewValidator
}
