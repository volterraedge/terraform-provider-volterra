//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"strings"

	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"gopkg.volterra.us/stdlib/client/vesapi"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_alert_policy "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/alert_policy"
)

// resourceVolterraAlertPolicy is implementation of Volterra's AlertPolicy resources
func resourceVolterraAlertPolicy() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraAlertPolicyCreate,
		Read:   resourceVolterraAlertPolicyRead,
		Update: resourceVolterraAlertPolicyUpdate,
		Delete: resourceVolterraAlertPolicyDelete,

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
			},

			"namespace": {
				Type:     schema.TypeString,
				Required: true,
			},

			"notification_parameters": {

				Type:     schema.TypeSet,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"custom": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"labels": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"default": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"individual": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"ves_io_group": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"group_interval": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"group_wait": {
							Type:     schema.TypeString,
							Optional: true,
						},

						"repeat_interval": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"receivers": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{
						"kind": {
							Type:     schema.TypeString,
							Computed: true,
						},
						"name": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"namespace": {
							Type:     schema.TypeString,
							Optional: true,
						},
						"tenant": {
							Type:     schema.TypeString,
							Optional: true,
						},
					},
				},
			},

			"routes": {

				Type:     schema.TypeList,
				Optional: true,
				Elem: &schema.Resource{
					Schema: map[string]*schema.Schema{

						"dont_send": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"send": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"alertname": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"alertname_regex": {

							Type:     schema.TypeString,
							Optional: true,
						},

						"any": {

							Type:     schema.TypeBool,
							Optional: true,
						},

						"custom": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"alertname": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"exact_match": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"regex_match": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"group": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"exact_match": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"regex_match": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},

									"severity": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"exact_match": {

													Type:     schema.TypeString,
													Optional: true,
												},

												"regex_match": {

													Type:     schema.TypeString,
													Optional: true,
												},
											},
										},
									},
								},
							},
						},

						"group": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"groups": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"severity": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"severities": {

										Type: schema.TypeList,

										Optional: true,
										Elem: &schema.Schema{
											Type: schema.TypeString,
										},
									},
								},
							},
						},

						"notification_parameters": {

							Type:     schema.TypeSet,
							Optional: true,
							Elem: &schema.Resource{
								Schema: map[string]*schema.Schema{

									"custom": {

										Type:     schema.TypeSet,
										Optional: true,
										Elem: &schema.Resource{
											Schema: map[string]*schema.Schema{

												"labels": {

													Type: schema.TypeList,

													Optional: true,
													Elem: &schema.Schema{
														Type: schema.TypeString,
													},
												},
											},
										},
									},

									"default": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"individual": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"ves_io_group": {

										Type:     schema.TypeBool,
										Optional: true,
									},

									"group_interval": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"group_wait": {
										Type:     schema.TypeString,
										Optional: true,
									},

									"repeat_interval": {
										Type:     schema.TypeString,
										Optional: true,
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

// resourceVolterraAlertPolicyCreate creates AlertPolicy resource
func resourceVolterraAlertPolicyCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	createMeta := &ves_io_schema.ObjectCreateMetaType{}
	createSpec := &ves_io_schema_alert_policy.CreateSpecType{}
	createReq := &ves_io_schema_alert_policy.CreateRequest{
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

	if v, ok := d.GetOk("notification_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		notificationParameters := &ves_io_schema_alert_policy.NotificationParameters{}
		createSpec.NotificationParameters = notificationParameters
		for _, set := range sl {

			notificationParametersMapStrToI := set.(map[string]interface{})

			groupByTypeFound := false

			if v, ok := notificationParametersMapStrToI["custom"]; ok && !isIntfNil(v) && !groupByTypeFound {

				groupByTypeFound = true
				groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Custom{}
				groupByInt.Custom = &ves_io_schema_alert_policy.CustomGroupBy{}
				notificationParameters.GroupBy = groupByInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["labels"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						groupByInt.Custom.Labels = ls

					}

				}

			}

			if v, ok := notificationParametersMapStrToI["default"]; ok && !isIntfNil(v) && !groupByTypeFound {

				groupByTypeFound = true

				if v.(bool) {
					groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Default{}
					groupByInt.Default = &ves_io_schema.Empty{}
					notificationParameters.GroupBy = groupByInt
				}

			}

			if v, ok := notificationParametersMapStrToI["individual"]; ok && !isIntfNil(v) && !groupByTypeFound {

				groupByTypeFound = true

				if v.(bool) {
					groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Individual{}
					groupByInt.Individual = &ves_io_schema.Empty{}
					notificationParameters.GroupBy = groupByInt
				}

			}

			if v, ok := notificationParametersMapStrToI["ves_io_group"]; ok && !isIntfNil(v) && !groupByTypeFound {

				groupByTypeFound = true

				if v.(bool) {
					groupByInt := &ves_io_schema_alert_policy.NotificationParameters_VesIoGroup{}
					groupByInt.VesIoGroup = &ves_io_schema.Empty{}
					notificationParameters.GroupBy = groupByInt
				}

			}

			if w, ok := notificationParametersMapStrToI["group_interval"]; ok && !isIntfNil(w) {
				notificationParameters.GroupInterval = w.(string)
			}

			if w, ok := notificationParametersMapStrToI["group_wait"]; ok && !isIntfNil(w) {
				notificationParameters.GroupWait = w.(string)
			}

			if w, ok := notificationParametersMapStrToI["repeat_interval"]; ok && !isIntfNil(w) {
				notificationParameters.RepeatInterval = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("receivers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		receiversInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		createSpec.Receivers = receiversInt
		for i, ps := range sl {

			rMapToStrVal := ps.(map[string]interface{})
			receiversInt[i] = &ves_io_schema.ObjectRefType{}

			receiversInt[i].Kind = "alert_receiver"

			if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
				receiversInt[i].Name = v.(string)
			}

			if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				receiversInt[i].Namespace = v.(string)
			}

			if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				receiversInt[i].Tenant = v.(string)
			}

			if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
				receiversInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		routes := make([]*ves_io_schema_alert_policy.Route, len(sl))
		createSpec.Routes = routes
		for i, set := range sl {
			routes[i] = &ves_io_schema_alert_policy.Route{}

			routesMapStrToI := set.(map[string]interface{})

			actionTypeFound := false

			if v, ok := routesMapStrToI["dont_send"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true

				if v.(bool) {
					actionInt := &ves_io_schema_alert_policy.Route_DontSend{}
					actionInt.DontSend = &ves_io_schema.Empty{}
					routes[i].Action = actionInt
				}

			}

			if v, ok := routesMapStrToI["send"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true

				if v.(bool) {
					actionInt := &ves_io_schema_alert_policy.Route_Send{}
					actionInt.Send = &ves_io_schema.Empty{}
					routes[i].Action = actionInt
				}

			}

			matcherTypeFound := false

			if v, ok := routesMapStrToI["alertname"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_Alertname{}

				routes[i].Matcher = matcherInt

				matcherInt.Alertname = ves_io_schema_alert_policy.AlertName(ves_io_schema_alert_policy.AlertName_value[v.(string)])

			}

			if v, ok := routesMapStrToI["alertname_regex"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_AlertnameRegex{}

				routes[i].Matcher = matcherInt

				matcherInt.AlertnameRegex = v.(string)

			}

			if v, ok := routesMapStrToI["any"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true

				if v.(bool) {
					matcherInt := &ves_io_schema_alert_policy.Route_Any{}
					matcherInt.Any = &ves_io_schema.Empty{}
					routes[i].Matcher = matcherInt
				}

			}

			if v, ok := routesMapStrToI["custom"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_Custom{}
				matcherInt.Custom = &ves_io_schema_alert_policy.CustomMatcher{}
				routes[i].Matcher = matcherInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["alertname"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						alertname := &ves_io_schema_alert_policy.LabelMatcher{}
						matcherInt.Custom.Alertname = alertname
						for _, set := range sl {

							alertnameMapStrToI := set.(map[string]interface{})

							matcherTypeTypeFound := false

							if v, ok := alertnameMapStrToI["exact_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_ExactMatch{}

								alertname.MatcherType = matcherTypeInt

								matcherTypeInt.ExactMatch = v.(string)

							}

							if v, ok := alertnameMapStrToI["regex_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_RegexMatch{}

								alertname.MatcherType = matcherTypeInt

								matcherTypeInt.RegexMatch = v.(string)

							}

						}

					}

					if v, ok := cs["group"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						group := &ves_io_schema_alert_policy.LabelMatcher{}
						matcherInt.Custom.Group = group
						for _, set := range sl {

							groupMapStrToI := set.(map[string]interface{})

							matcherTypeTypeFound := false

							if v, ok := groupMapStrToI["exact_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_ExactMatch{}

								group.MatcherType = matcherTypeInt

								matcherTypeInt.ExactMatch = v.(string)

							}

							if v, ok := groupMapStrToI["regex_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_RegexMatch{}

								group.MatcherType = matcherTypeInt

								matcherTypeInt.RegexMatch = v.(string)

							}

						}

					}

					if v, ok := cs["severity"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						severity := &ves_io_schema_alert_policy.LabelMatcher{}
						matcherInt.Custom.Severity = severity
						for _, set := range sl {

							severityMapStrToI := set.(map[string]interface{})

							matcherTypeTypeFound := false

							if v, ok := severityMapStrToI["exact_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_ExactMatch{}

								severity.MatcherType = matcherTypeInt

								matcherTypeInt.ExactMatch = v.(string)

							}

							if v, ok := severityMapStrToI["regex_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_RegexMatch{}

								severity.MatcherType = matcherTypeInt

								matcherTypeInt.RegexMatch = v.(string)

							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["group"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_Group{}
				matcherInt.Group = &ves_io_schema_alert_policy.GroupMatcher{}
				routes[i].Matcher = matcherInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["groups"]; ok && !isIntfNil(v) {

						groupsList := []ves_io_schema_alert_policy.Group{}
						for _, j := range v.([]interface{}) {
							groupsList = append(groupsList, ves_io_schema_alert_policy.Group(ves_io_schema_alert_policy.Group_value[j.(string)]))
						}
						matcherInt.Group.Groups = groupsList

					}

				}

			}

			if v, ok := routesMapStrToI["severity"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_Severity{}
				matcherInt.Severity = &ves_io_schema_alert_policy.SeverityMatcher{}
				routes[i].Matcher = matcherInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["severities"]; ok && !isIntfNil(v) {

						severitiesList := []ves_io_schema_alert_policy.Severity{}
						for _, j := range v.([]interface{}) {
							severitiesList = append(severitiesList, ves_io_schema_alert_policy.Severity(ves_io_schema_alert_policy.Severity_value[j.(string)]))
						}
						matcherInt.Severity.Severities = severitiesList

					}

				}

			}

			if v, ok := routesMapStrToI["notification_parameters"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				notificationParameters := &ves_io_schema_alert_policy.NotificationParameters{}
				routes[i].NotificationParameters = notificationParameters
				for _, set := range sl {

					notificationParametersMapStrToI := set.(map[string]interface{})

					groupByTypeFound := false

					if v, ok := notificationParametersMapStrToI["custom"]; ok && !isIntfNil(v) && !groupByTypeFound {

						groupByTypeFound = true
						groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Custom{}
						groupByInt.Custom = &ves_io_schema_alert_policy.CustomGroupBy{}
						notificationParameters.GroupBy = groupByInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["labels"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								groupByInt.Custom.Labels = ls

							}

						}

					}

					if v, ok := notificationParametersMapStrToI["default"]; ok && !isIntfNil(v) && !groupByTypeFound {

						groupByTypeFound = true

						if v.(bool) {
							groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Default{}
							groupByInt.Default = &ves_io_schema.Empty{}
							notificationParameters.GroupBy = groupByInt
						}

					}

					if v, ok := notificationParametersMapStrToI["individual"]; ok && !isIntfNil(v) && !groupByTypeFound {

						groupByTypeFound = true

						if v.(bool) {
							groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Individual{}
							groupByInt.Individual = &ves_io_schema.Empty{}
							notificationParameters.GroupBy = groupByInt
						}

					}

					if v, ok := notificationParametersMapStrToI["ves_io_group"]; ok && !isIntfNil(v) && !groupByTypeFound {

						groupByTypeFound = true

						if v.(bool) {
							groupByInt := &ves_io_schema_alert_policy.NotificationParameters_VesIoGroup{}
							groupByInt.VesIoGroup = &ves_io_schema.Empty{}
							notificationParameters.GroupBy = groupByInt
						}

					}

					if w, ok := notificationParametersMapStrToI["group_interval"]; ok && !isIntfNil(w) {
						notificationParameters.GroupInterval = w.(string)
					}

					if w, ok := notificationParametersMapStrToI["group_wait"]; ok && !isIntfNil(w) {
						notificationParameters.GroupWait = w.(string)
					}

					if w, ok := notificationParametersMapStrToI["repeat_interval"]; ok && !isIntfNil(w) {
						notificationParameters.RepeatInterval = w.(string)
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Creating Volterra AlertPolicy object with struct: %+v", createReq)

	createAlertPolicyResp, err := client.CreateObject(context.Background(), ves_io_schema_alert_policy.ObjectType, createReq)
	if err != nil {
		return fmt.Errorf("error creating AlertPolicy: %s", err)
	}
	d.SetId(createAlertPolicyResp.GetObjSystemMetadata().GetUid())

	return resourceVolterraAlertPolicyRead(d, meta)
}

func resourceVolterraAlertPolicyRead(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	resp, err := client.GetObject(context.Background(), ves_io_schema_alert_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AlertPolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AlertPolicy %q: %s", d.Id(), err)
	}
	return setAlertPolicyFields(client, d, resp)
}

func setAlertPolicyFields(client *APIClient, d *schema.ResourceData, resp vesapi.GetObjectResponse) error {
	metadata := resp.GetObjMetadata()

	d.Set("annotations", metadata.GetAnnotations())

	d.Set("description", metadata.GetDescription())

	d.Set("disable", metadata.GetDisable())

	d.Set("labels", metadata.GetLabels())

	d.Set("name", metadata.GetName())

	d.Set("namespace", metadata.GetNamespace())

	return nil
}

// resourceVolterraAlertPolicyUpdate updates AlertPolicy resource
func resourceVolterraAlertPolicyUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	updateMeta := &ves_io_schema.ObjectReplaceMetaType{}
	updateSpec := &ves_io_schema_alert_policy.ReplaceSpecType{}
	updateReq := &ves_io_schema_alert_policy.ReplaceRequest{
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

	if v, ok := d.GetOk("notification_parameters"); ok && !isIntfNil(v) {

		sl := v.(*schema.Set).List()
		notificationParameters := &ves_io_schema_alert_policy.NotificationParameters{}
		updateSpec.NotificationParameters = notificationParameters
		for _, set := range sl {

			notificationParametersMapStrToI := set.(map[string]interface{})

			groupByTypeFound := false

			if v, ok := notificationParametersMapStrToI["custom"]; ok && !isIntfNil(v) && !groupByTypeFound {

				groupByTypeFound = true
				groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Custom{}
				groupByInt.Custom = &ves_io_schema_alert_policy.CustomGroupBy{}
				notificationParameters.GroupBy = groupByInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["labels"]; ok && !isIntfNil(v) {

						ls := make([]string, len(v.([]interface{})))
						for i, v := range v.([]interface{}) {
							ls[i] = v.(string)
						}
						groupByInt.Custom.Labels = ls

					}

				}

			}

			if v, ok := notificationParametersMapStrToI["default"]; ok && !isIntfNil(v) && !groupByTypeFound {

				groupByTypeFound = true

				if v.(bool) {
					groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Default{}
					groupByInt.Default = &ves_io_schema.Empty{}
					notificationParameters.GroupBy = groupByInt
				}

			}

			if v, ok := notificationParametersMapStrToI["individual"]; ok && !isIntfNil(v) && !groupByTypeFound {

				groupByTypeFound = true

				if v.(bool) {
					groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Individual{}
					groupByInt.Individual = &ves_io_schema.Empty{}
					notificationParameters.GroupBy = groupByInt
				}

			}

			if v, ok := notificationParametersMapStrToI["ves_io_group"]; ok && !isIntfNil(v) && !groupByTypeFound {

				groupByTypeFound = true

				if v.(bool) {
					groupByInt := &ves_io_schema_alert_policy.NotificationParameters_VesIoGroup{}
					groupByInt.VesIoGroup = &ves_io_schema.Empty{}
					notificationParameters.GroupBy = groupByInt
				}

			}

			if w, ok := notificationParametersMapStrToI["group_interval"]; ok && !isIntfNil(w) {
				notificationParameters.GroupInterval = w.(string)
			}

			if w, ok := notificationParametersMapStrToI["group_wait"]; ok && !isIntfNil(w) {
				notificationParameters.GroupWait = w.(string)
			}

			if w, ok := notificationParametersMapStrToI["repeat_interval"]; ok && !isIntfNil(w) {
				notificationParameters.RepeatInterval = w.(string)
			}

		}

	}

	if v, ok := d.GetOk("receivers"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		receiversInt := make([]*ves_io_schema.ObjectRefType, len(sl))
		updateSpec.Receivers = receiversInt
		for i, ps := range sl {

			rMapToStrVal := ps.(map[string]interface{})
			receiversInt[i] = &ves_io_schema.ObjectRefType{}

			receiversInt[i].Kind = "alert_receiver"

			if v, ok := rMapToStrVal["name"]; ok && !isIntfNil(v) {
				receiversInt[i].Name = v.(string)
			}

			if v, ok := rMapToStrVal["namespace"]; ok && !isIntfNil(v) {
				receiversInt[i].Namespace = v.(string)
			}

			if v, ok := rMapToStrVal["tenant"]; ok && !isIntfNil(v) {
				receiversInt[i].Tenant = v.(string)
			}

			if v, ok := rMapToStrVal["uid"]; ok && !isIntfNil(v) {
				receiversInt[i].Uid = v.(string)
			}

		}

	}

	if v, ok := d.GetOk("routes"); ok && !isIntfNil(v) {

		sl := v.([]interface{})
		routes := make([]*ves_io_schema_alert_policy.Route, len(sl))
		updateSpec.Routes = routes
		for i, set := range sl {
			routes[i] = &ves_io_schema_alert_policy.Route{}

			routesMapStrToI := set.(map[string]interface{})

			actionTypeFound := false

			if v, ok := routesMapStrToI["dont_send"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true

				if v.(bool) {
					actionInt := &ves_io_schema_alert_policy.Route_DontSend{}
					actionInt.DontSend = &ves_io_schema.Empty{}
					routes[i].Action = actionInt
				}

			}

			if v, ok := routesMapStrToI["send"]; ok && !isIntfNil(v) && !actionTypeFound {

				actionTypeFound = true

				if v.(bool) {
					actionInt := &ves_io_schema_alert_policy.Route_Send{}
					actionInt.Send = &ves_io_schema.Empty{}
					routes[i].Action = actionInt
				}

			}

			matcherTypeFound := false

			if v, ok := routesMapStrToI["alertname"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_Alertname{}

				routes[i].Matcher = matcherInt

				matcherInt.Alertname = ves_io_schema_alert_policy.AlertName(ves_io_schema_alert_policy.AlertName_value[v.(string)])

			}

			if v, ok := routesMapStrToI["alertname_regex"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_AlertnameRegex{}

				routes[i].Matcher = matcherInt

				matcherInt.AlertnameRegex = v.(string)

			}

			if v, ok := routesMapStrToI["any"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true

				if v.(bool) {
					matcherInt := &ves_io_schema_alert_policy.Route_Any{}
					matcherInt.Any = &ves_io_schema.Empty{}
					routes[i].Matcher = matcherInt
				}

			}

			if v, ok := routesMapStrToI["custom"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_Custom{}
				matcherInt.Custom = &ves_io_schema_alert_policy.CustomMatcher{}
				routes[i].Matcher = matcherInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["alertname"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						alertname := &ves_io_schema_alert_policy.LabelMatcher{}
						matcherInt.Custom.Alertname = alertname
						for _, set := range sl {

							alertnameMapStrToI := set.(map[string]interface{})

							matcherTypeTypeFound := false

							if v, ok := alertnameMapStrToI["exact_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_ExactMatch{}

								alertname.MatcherType = matcherTypeInt

								matcherTypeInt.ExactMatch = v.(string)

							}

							if v, ok := alertnameMapStrToI["regex_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_RegexMatch{}

								alertname.MatcherType = matcherTypeInt

								matcherTypeInt.RegexMatch = v.(string)

							}

						}

					}

					if v, ok := cs["group"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						group := &ves_io_schema_alert_policy.LabelMatcher{}
						matcherInt.Custom.Group = group
						for _, set := range sl {

							groupMapStrToI := set.(map[string]interface{})

							matcherTypeTypeFound := false

							if v, ok := groupMapStrToI["exact_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_ExactMatch{}

								group.MatcherType = matcherTypeInt

								matcherTypeInt.ExactMatch = v.(string)

							}

							if v, ok := groupMapStrToI["regex_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_RegexMatch{}

								group.MatcherType = matcherTypeInt

								matcherTypeInt.RegexMatch = v.(string)

							}

						}

					}

					if v, ok := cs["severity"]; ok && !isIntfNil(v) {

						sl := v.(*schema.Set).List()
						severity := &ves_io_schema_alert_policy.LabelMatcher{}
						matcherInt.Custom.Severity = severity
						for _, set := range sl {

							severityMapStrToI := set.(map[string]interface{})

							matcherTypeTypeFound := false

							if v, ok := severityMapStrToI["exact_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_ExactMatch{}

								severity.MatcherType = matcherTypeInt

								matcherTypeInt.ExactMatch = v.(string)

							}

							if v, ok := severityMapStrToI["regex_match"]; ok && !isIntfNil(v) && !matcherTypeTypeFound {

								matcherTypeTypeFound = true
								matcherTypeInt := &ves_io_schema_alert_policy.LabelMatcher_RegexMatch{}

								severity.MatcherType = matcherTypeInt

								matcherTypeInt.RegexMatch = v.(string)

							}

						}

					}

				}

			}

			if v, ok := routesMapStrToI["group"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_Group{}
				matcherInt.Group = &ves_io_schema_alert_policy.GroupMatcher{}
				routes[i].Matcher = matcherInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["groups"]; ok && !isIntfNil(v) {

						groupsList := []ves_io_schema_alert_policy.Group{}
						for _, j := range v.([]interface{}) {
							groupsList = append(groupsList, ves_io_schema_alert_policy.Group(ves_io_schema_alert_policy.Group_value[j.(string)]))
						}
						matcherInt.Group.Groups = groupsList

					}

				}

			}

			if v, ok := routesMapStrToI["severity"]; ok && !isIntfNil(v) && !matcherTypeFound {

				matcherTypeFound = true
				matcherInt := &ves_io_schema_alert_policy.Route_Severity{}
				matcherInt.Severity = &ves_io_schema_alert_policy.SeverityMatcher{}
				routes[i].Matcher = matcherInt

				sl := v.(*schema.Set).List()
				for _, set := range sl {
					cs := set.(map[string]interface{})

					if v, ok := cs["severities"]; ok && !isIntfNil(v) {

						severitiesList := []ves_io_schema_alert_policy.Severity{}
						for _, j := range v.([]interface{}) {
							severitiesList = append(severitiesList, ves_io_schema_alert_policy.Severity(ves_io_schema_alert_policy.Severity_value[j.(string)]))
						}
						matcherInt.Severity.Severities = severitiesList

					}

				}

			}

			if v, ok := routesMapStrToI["notification_parameters"]; ok && !isIntfNil(v) {

				sl := v.(*schema.Set).List()
				notificationParameters := &ves_io_schema_alert_policy.NotificationParameters{}
				routes[i].NotificationParameters = notificationParameters
				for _, set := range sl {

					notificationParametersMapStrToI := set.(map[string]interface{})

					groupByTypeFound := false

					if v, ok := notificationParametersMapStrToI["custom"]; ok && !isIntfNil(v) && !groupByTypeFound {

						groupByTypeFound = true
						groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Custom{}
						groupByInt.Custom = &ves_io_schema_alert_policy.CustomGroupBy{}
						notificationParameters.GroupBy = groupByInt

						sl := v.(*schema.Set).List()
						for _, set := range sl {
							cs := set.(map[string]interface{})

							if v, ok := cs["labels"]; ok && !isIntfNil(v) {

								ls := make([]string, len(v.([]interface{})))
								for i, v := range v.([]interface{}) {
									ls[i] = v.(string)
								}
								groupByInt.Custom.Labels = ls

							}

						}

					}

					if v, ok := notificationParametersMapStrToI["default"]; ok && !isIntfNil(v) && !groupByTypeFound {

						groupByTypeFound = true

						if v.(bool) {
							groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Default{}
							groupByInt.Default = &ves_io_schema.Empty{}
							notificationParameters.GroupBy = groupByInt
						}

					}

					if v, ok := notificationParametersMapStrToI["individual"]; ok && !isIntfNil(v) && !groupByTypeFound {

						groupByTypeFound = true

						if v.(bool) {
							groupByInt := &ves_io_schema_alert_policy.NotificationParameters_Individual{}
							groupByInt.Individual = &ves_io_schema.Empty{}
							notificationParameters.GroupBy = groupByInt
						}

					}

					if v, ok := notificationParametersMapStrToI["ves_io_group"]; ok && !isIntfNil(v) && !groupByTypeFound {

						groupByTypeFound = true

						if v.(bool) {
							groupByInt := &ves_io_schema_alert_policy.NotificationParameters_VesIoGroup{}
							groupByInt.VesIoGroup = &ves_io_schema.Empty{}
							notificationParameters.GroupBy = groupByInt
						}

					}

					if w, ok := notificationParametersMapStrToI["group_interval"]; ok && !isIntfNil(w) {
						notificationParameters.GroupInterval = w.(string)
					}

					if w, ok := notificationParametersMapStrToI["group_wait"]; ok && !isIntfNil(w) {
						notificationParameters.GroupWait = w.(string)
					}

					if w, ok := notificationParametersMapStrToI["repeat_interval"]; ok && !isIntfNil(w) {
						notificationParameters.RepeatInterval = w.(string)
					}

				}

			}

		}

	}

	log.Printf("[DEBUG] Updating Volterra AlertPolicy obj with struct: %+v", updateReq)

	err := client.ReplaceObject(context.Background(), ves_io_schema_alert_policy.ObjectType, updateReq)
	if err != nil {
		return fmt.Errorf("error updating AlertPolicy: %s", err)
	}

	return resourceVolterraAlertPolicyRead(d, meta)
}

func resourceVolterraAlertPolicyDelete(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)
	name := d.Get("name").(string)
	namespace := d.Get("namespace").(string)

	_, err := client.GetObject(context.Background(), ves_io_schema_alert_policy.ObjectType, namespace, name)
	if err != nil {
		if strings.Contains(err.Error(), "status code 404") {
			log.Printf("[INFO] AlertPolicy %s no longer exists", d.Id())
			d.SetId("")
			return nil
		}
		return fmt.Errorf("Error finding Volterra AlertPolicy before deleting %q: %s", d.Id(), err)
	}

	log.Printf("[DEBUG] Deleting Volterra AlertPolicy obj with name %+v in namespace %+v", name, namespace)
	return client.DeleteObject(context.Background(), ves_io_schema_alert_policy.ObjectType, namespace, name)
}