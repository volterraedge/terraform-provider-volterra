//
// Copyright (c) 2018 Volterra, Inc. All rights reserved.
// Code generated by ves-gen-tf-provider. DO NOT EDIT.
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"
	google_protobuf "github.com/gogo/protobuf/types"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_filter_set "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/filter_set"
)

// resourceVolterraFilterSet is implementation of Volterra's FilterSet resources
func resourceVolterraFilterSet() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraFilterSetCreate,
		Read:   resourceVolterraFilterSetRead,
		Update: resourceVolterraFilterSetUpdate,
		Delete: resourceVolterraFilterSetDelete,

		Schema: map[string]*schema.Schema{

			"annotations": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"description": {
				Type:     schema.TypeString,
				Optional: true,
			},

			"disable": {
				Type:     schema.TypeBool,
				Optional: true,
			},

			"labels": {
				Type:     schema.TypeMap,
				Optional: true,
			},

			"name": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
				ForceNew: true,
			},

			"context_key": {
				Type:     schema.TypeString,
				Required: true,
			},

			"filter_fields": {

				Type:     schema.TypeList,
				Required: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"field_id": {
							Type:     schema.TypeString,
							Required: true,
						},

						"date_field": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"absolute": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"end_date": {
													Type:     schema.TypeString,
													Required: true,
												},

												"start_date": {
													Type:     schema.TypeString,
													Required: true,
												},
											},
										},
									},

									"relative": {

										Type:     schema.TypeString,
										Optional: true,
									},
								},
							},
						},

						"filter_expression_field": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"expression": {
										Type:     schema.TypeString,
										Required: true,
									},
								},
							},
						},

						"label_selector_field": {

							Type:       schema.TypeSet,
							Optional:   true,
							Deprecated: "This field is deprecated and will be removed in future release.",
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"expressions": {

										Type: schema.TypeList,

										Required:   true,
										Deprecated: "This field is deprecated and will be removed in future release.",
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"string_field": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"field_values": {

										Type: schema.TypeList,

										Required: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},
					},
				},
			},
		},
	}
}

// resourceVolterraFilterSetCreate creates FilterSet resource
func resourceVolterraFilterSetCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_filter_set.CreateSpecType{}
	createReq := &ves_io_schema_filter_set.CreateRequest{
		Metadata: createMeta,
		Spec:     createSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		createMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		createMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		createMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		createMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		createMeta.Namespace =
			v.(string)
	}

	//context_key
	if v, ok := d.GetOk("context_key"); ok && !isIntfNil(v) {

		createSpec.ContextKey =
			v.(string)

	}

	//filter_fields
	if v, ok := d.GetOk("filter_fields"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		filterFields := make([]*ves_io_schema_filter_set.FilterSetField, len(sl))
		createSpec.FilterFields = filterFields
		for i, set := range sl {
			filterFields[i] = &ves_io_schema_filter_set.FilterSetField{}
			filterFieldsMapStrToI := set.(map[string]interface{})

			if w, ok := filterFieldsMapStrToI["field_id"]; ok && !isIntfNil(w) {
				filterFields[i].FieldId = w.(string)
			}

			fieldValueTypeFound := false

			if v, ok := filterFieldsMapStrToI["date_field"]; ok && !isIntfNil(v) && !fieldValueTypeFound {

				fieldValueTypeFound = true
				fieldValueInt := &ves_io_schema_filter_set.FilterSetField_DateField{}
				fieldValueInt.DateField = &ves_io_schema_filter_set.FilterTimeRangeField{}
				filterFields[i].FieldValue = fieldValueInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					rangeTypeTypeFound := false

					if v, ok := cs["absolute"]; ok && !isIntfNil(v) && !rangeTypeTypeFound {

						rangeTypeTypeFound = true
						rangeTypeInt := &ves_io_schema_filter_set.FilterTimeRangeField_Absolute{}
						rangeTypeInt.Absolute = &ves_io_schema.DateRange{}
						fieldValueInt.DateField.RangeType = rangeTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["end_date"]; ok && !isIntfNil(v) {

								ts, err := parseTime(v.(string))
								if err != nil {
									return fmt.Errorf("error creating EndDate, timestamp format is wrong: %s", err)
								}
								rangeTypeInt.Absolute.EndDate = ts
							}

							if v, ok := cs["start_date"]; ok && !isIntfNil(v) {

								ts, err := parseTime(v.(string))
								if err != nil {
									return fmt.Errorf("error creating StartDate, timestamp format is wrong: %s", err)
								}
								rangeTypeInt.Absolute.StartDate = ts
							}

						}

					}

					if _, ok := cs["relative"]; ok && !rangeTypeTypeFound {

						rangeTypeTypeFound = true
						rangeTypeInt := &ves_io_schema_filter_set.FilterTimeRangeField_Relative{}
						rangeTypeInt.Relative = &google_protobuf.Duration{}
						fieldValueInt.DateField.RangeType = rangeTypeInt

					}

				}

			}

			if v, ok := filterFieldsMapStrToI["filter_expression_field"]; ok && !isIntfNil(v) && !fieldValueTypeFound {

				fieldValueTypeFound = true
				fieldValueInt := &ves_io_schema_filter_set.FilterSetField_FilterExpressionField{}
				fieldValueInt.FilterExpressionField = &ves_io_schema_filter_set.FilterExpressionField{}
				filterFields[i].FieldValue = fieldValueInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expression"]; ok && !isIntfNil(v) {

						fieldValueInt.FilterExpressionField.Expression = v.(string)

					}

				}

			}

			if v, ok := filterFieldsMapStrToI["label_selector_field"]; ok && !isIntfNil(v) && !fieldValueTypeFound {

				fieldValueTypeFound = true
				fieldValueInt := &ves_io_schema_filter_set.FilterSetField_LabelSelectorField{}
				fieldValueInt.LabelSelectorField = &ves_io_schema.LabelSelectorType{}
				filterFields[i].FieldValue = fieldValueInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						fieldValueInt.LabelSelectorField.Expressions = ls

					}

				}

			}

			if v, ok := filterFieldsMapStrToI["string_field"]; ok && !isIntfNil(v) && !fieldValueTypeFound {

				fieldValueTypeFound = true
				fieldValueInt := &ves_io_schema_filter_set.FilterSetField_StringField{}
				fieldValueInt.StringField = &ves_io_schema_filter_set.FilterStringField{}
				filterFields[i].FieldValue = fieldValueInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["field_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						fieldValueInt.StringField.FieldValues = ls

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra FilterSet object with struct: %+v", createReq)

	createFilterSetResp, err := client.CreateObject(context.Background(), ves_io_schema_filter_set.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating FilterSet: %s", err)
	}
	d.SetId(createFilterSetResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraFilterSetRead(d, meta)
}

func resourceVolterraFilterSetRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_filter_set.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] FilterSet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra FilterSet %q: %s", d.Id(), err)
	}
	return setFilterSetFields(client, d, resp)
}

func setFilterSetFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraFilterSetUpdate updates FilterSet resource
func resourceVolterraFilterSetUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_filter_set.ReplaceSpecType{}
	updateReq := &ves_io_schema_filter_set.ReplaceRequest{
		Metadata: updateMeta,
		Spec:     updateSpec,
	}

	if v, ok := d.GetOk("annotations"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Annotations = ms
	}

	if v, ok := d.GetOk("description"); ok && !isIntfNil(v) {
		updateMeta.Description =
			v.(string)
	}

	if v, ok := d.GetOk("disable"); ok && !isIntfNil(v) {
		updateMeta.Disable =
			v.(bool)
	}

	if v, ok := d.GetOk("labels"); ok && !isIntfNil(v) {

		ms := map[string]string{}

		for k, v := range v.(map[string]interface{}) {
			val := v.(string)
			ms[k] = val
		}
		updateMeta.Labels = ms
	}

	if v, ok := d.GetOk("name"); ok && !isIntfNil(v) {
		updateMeta.Name =
			v.(string)
	}

	if v, ok := d.GetOk("namespace"); ok && !isIntfNil(v) {
		updateMeta.Namespace =
			v.(string)
	}

	if v, ok := d.GetOk("context_key"); ok && !isIntfNil(v) {

		updateSpec.ContextKey =
			v.(string)

	}

	if v, ok := d.GetOk("filter_fields"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		filterFields := make([]*ves_io_schema_filter_set.FilterSetField, len(sl))
		updateSpec.FilterFields = filterFields
		for i, set := range sl {
			filterFields[i] = &ves_io_schema_filter_set.FilterSetField{}
			filterFieldsMapStrToI := set.(map[string]interface{})

			if w, ok := filterFieldsMapStrToI["field_id"]; ok && !isIntfNil(w) {
				filterFields[i].FieldId = w.(string)
			}

			fieldValueTypeFound := false

			if v, ok := filterFieldsMapStrToI["date_field"]; ok && !isIntfNil(v) && !fieldValueTypeFound {

				fieldValueTypeFound = true
				fieldValueInt := &ves_io_schema_filter_set.FilterSetField_DateField{}
				fieldValueInt.DateField = &ves_io_schema_filter_set.FilterTimeRangeField{}
				filterFields[i].FieldValue = fieldValueInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					rangeTypeTypeFound := false

					if v, ok := cs["absolute"]; ok && !isIntfNil(v) && !rangeTypeTypeFound {

						rangeTypeTypeFound = true
						rangeTypeInt := &ves_io_schema_filter_set.FilterTimeRangeField_Absolute{}
						rangeTypeInt.Absolute = &ves_io_schema.DateRange{}
						fieldValueInt.DateField.RangeType = rangeTypeInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["end_date"]; ok && !isIntfNil(v) {

								ts, err := parseTime(v.(string))
								if err != nil {
									return fmt.Errorf("error creating EndDate, timestamp format is wrong: %s", err)
								}
								rangeTypeInt.Absolute.EndDate = ts
							}

							if v, ok := cs["start_date"]; ok && !isIntfNil(v) {

								ts, err := parseTime(v.(string))
								if err != nil {
									return fmt.Errorf("error creating StartDate, timestamp format is wrong: %s", err)
								}
								rangeTypeInt.Absolute.StartDate = ts
							}

						}

					}

					if _, ok := cs["relative"]; ok && !rangeTypeTypeFound {

						rangeTypeTypeFound = true
						rangeTypeInt := &ves_io_schema_filter_set.FilterTimeRangeField_Relative{}
						rangeTypeInt.Relative = &google_protobuf.Duration{}
						fieldValueInt.DateField.RangeType = rangeTypeInt

					}

				}

			}

			if v, ok := filterFieldsMapStrToI["filter_expression_field"]; ok && !isIntfNil(v) && !fieldValueTypeFound {

				fieldValueTypeFound = true
				fieldValueInt := &ves_io_schema_filter_set.FilterSetField_FilterExpressionField{}
				fieldValueInt.FilterExpressionField = &ves_io_schema_filter_set.FilterExpressionField{}
				filterFields[i].FieldValue = fieldValueInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expression"]; ok && !isIntfNil(v) {

						fieldValueInt.FilterExpressionField.Expression = v.(string)

					}

				}

			}

			if v, ok := filterFieldsMapStrToI["label_selector_field"]; ok && !isIntfNil(v) && !fieldValueTypeFound {

				fieldValueTypeFound = true
				fieldValueInt := &ves_io_schema_filter_set.FilterSetField_LabelSelectorField{}
				fieldValueInt.LabelSelectorField = &ves_io_schema.LabelSelectorType{}
				filterFields[i].FieldValue = fieldValueInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["expressions"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						fieldValueInt.LabelSelectorField.Expressions = ls

					}

				}

			}

			if v, ok := filterFieldsMapStrToI["string_field"]; ok && !isIntfNil(v) && !fieldValueTypeFound {

				fieldValueTypeFound = true
				fieldValueInt := &ves_io_schema_filter_set.FilterSetField_StringField{}
				fieldValueInt.StringField = &ves_io_schema_filter_set.FilterStringField{}
				filterFields[i].FieldValue = fieldValueInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["field_values"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						fieldValueInt.StringField.FieldValues = ls

					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra FilterSet obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_filter_set.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating FilterSet: %s", err)
	}

	return resourceVolterraFilterSetRead(d, meta)
}

func resourceVolterraFilterSetDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_filter_set.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] FilterSet %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra FilterSet before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra FilterSet obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_filter_set.ObjectType, namespace, name)
}
