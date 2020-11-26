//
// Copyright (c) 2020 Volterra, Inc. Licensed under APACHE LICENSE, VERSION 2.0
//

package volterra

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"strings"
	"time"

	"github.com/google/uuid"
	"github.com/hashicorp/terraform-plugin-sdk/helper/schema"
	"github.com/hashicorp/terraform-plugin-sdk/helper/validation"
	"gopkg.volterra.us/stdlib/codec"
	"gopkg.volterra.us/stdlib/errors"

	ves_io_schema_tf_params "gopkg.volterra.us/terraform-provider-volterra/pbgo/extschema/schema/views/terraform_parameters"
)

const (
	tfParamsActionRunRPCFQN = "ves.io.schema.views.terraform_parameters.CustomActionAPI.Run"
	tfParamsGetStatusRPCFQN = "ves.io.schema.views.terraform_parameters.CustomAPI.GetStatus"
	tfParamsActionRunURI    = "/public/namespaces/%s/terraform/%s/%s/run"
	tfParamsGetStatusURI    = "/public/namespaces/%s/terraform_parameters/%s/%s/status"
)

type tfActionParams struct {
	siteName string
	siteKind string
	action   string

	ignoreOnUpdate bool
	waitForAction  bool
}

// resourceVolterraTFParamsRunAction is implementation of Volterra's Namespace resources
func resourceVolterraTFParamsRunAction() *schema.Resource {
	return &schema.Resource{
		Create: resourceVolterraTFParamsRunActionCreate,
		Read:   resourceVolterraTFParamsRunActionRead,
		Update: resourceVolterraTFParamsRunActionUpdate,
		Delete: resourceVolterraTFParamsRunActionDelete,

		Schema: map[string]*schema.Schema{
			"site_name": {
				Type:     schema.TypeString,
				Required: true,
			},
			"site_kind": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"aws_vpc_site",
					"aws_tgw_site",
					"gcp_vpc_site",
					"azure_vnet_site",
				}, false),
			},
			"action": {
				Type:     schema.TypeString,
				Required: true,
				ValidateFunc: validation.StringInSlice([]string{
					"plan",
					"apply",
				}, false),
			},
			"ignore_on_update": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  false,
			},
			"wait_for_action": {
				Type:     schema.TypeBool,
				Optional: true,
				Default:  true,
			},
			"tf_output": {
				Type:     schema.TypeString,
				Computed: true,
			},
		},
	}
}

// resourceVolterraTFParamsRunActionCreate creates Namespace resource
func resourceVolterraTFParamsRunActionCreate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	actionParams := &tfActionParams{}
	if v, ok := d.GetOk("site_name"); ok {
		actionParams.siteName = v.(string)
	}
	if v, ok := d.GetOk("site_kind"); ok {
		actionParams.siteKind = v.(string)
	}
	if v, ok := d.GetOk("action"); ok {
		actionParams.action = v.(string)
	}
	if v, ok := d.GetOk("wait_for_action"); ok {
		actionParams.waitForAction = v.(bool)
	}

	actionVal, ok := ves_io_schema_tf_params.RunAction_value[strings.ToUpper(actionParams.action)]
	if !ok {
		return fmt.Errorf("Invalid action input: %s", actionParams.action)
	}

	req := &ves_io_schema_tf_params.RunRequest{
		Namespace: "system",
		ViewKind:  actionParams.siteKind,
		ViewName:  actionParams.siteName,
		Action:    ves_io_schema_tf_params.RunAction(actionVal),
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Running terraform parameter action struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost,
		fmt.Sprintf(tfParamsActionRunURI, "system", actionParams.siteKind, actionParams.siteName), tfParamsActionRunRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Running terraform parameter action struct: %s", err)
	}

	var tfOutput string
	if actionParams.waitForAction {
		tfOutput, err = waitForActionToComplete(client, actionParams.action,
			actionParams.siteName, actionParams.siteKind)
		if err != nil {
			return err
		}
	}
	d.Set("tf_output", tfOutput)

	d.SetId(uuid.New().String())
	return nil
}

func resourceVolterraTFParamsRunActionRead(d *schema.ResourceData, meta interface{}) error {
	// Ignore dont do anything when refresh/get is trigerred
	return nil
}

// resourceVolterraTFParamsRunActionUpdate updates Namespace resource
func resourceVolterraTFParamsRunActionUpdate(d *schema.ResourceData, meta interface{}) error {
	client := meta.(*APIClient)

	actionParams := &tfActionParams{}
	if v, ok := d.GetOk("ignore_on_update"); ok {
		actionParams.ignoreOnUpdate = v.(bool)
	}
	if !actionParams.ignoreOnUpdate {
		if v, ok := d.GetOk("site_name"); ok {
			actionParams.siteName = v.(string)
		}
		if v, ok := d.GetOk("site_kind"); ok {
			actionParams.siteKind = v.(string)
		}
		if v, ok := d.GetOk("action"); ok {
			actionParams.action = v.(string)
		}
		actionVal, ok := ves_io_schema_tf_params.RunAction_value[strings.ToUpper(actionParams.action)]
		if !ok {
			return fmt.Errorf("Invalid action input: %s", actionParams.action)
		}

		req := &ves_io_schema_tf_params.RunRequest{
			Namespace: "system",
			ViewKind:  actionParams.siteKind,
			ViewName:  actionParams.siteName,
			Action:    ves_io_schema_tf_params.RunAction(actionVal),
		}

		yamlReq, err := codec.ToYAML(req)
		if err != nil {
			return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
		}

		log.Printf("[DEBUG] Running terraform parameter action struct: %+v", req)
		_, err = client.CustomAPI(context.Background(), http.MethodPost,
			fmt.Sprintf(tfParamsActionRunURI, "system", actionParams.siteKind, actionParams.siteName), tfParamsActionRunRPCFQN, yamlReq)
		if err != nil {
			return fmt.Errorf("Error Running terraform parameter action struct: %s", err)
		}

		var tfOutput string
		if actionParams.waitForAction {
			tfOutput, err = waitForActionToComplete(client, actionParams.action,
				actionParams.siteName, actionParams.siteKind)
			if err != nil {
				return err
			}
		}
		d.Set("tf_output", tfOutput)
	}
	return nil
}

func resourceVolterraTFParamsRunActionDelete(d *schema.ResourceData, meta interface{}) error {
	// Remove the VPN tunnel config
	client := meta.(*APIClient)

	actionParams := &tfActionParams{}
	if v, ok := d.GetOk("site_name"); ok {
		actionParams.siteName = v.(string)
	}
	if v, ok := d.GetOk("site_kind"); ok {
		actionParams.siteKind = v.(string)
	}

	actionVal, ok := ves_io_schema_tf_params.RunAction_value["DESTROY"]
	if !ok {
		return fmt.Errorf("Invalid action input: %s", actionParams.action)
	}

	req := &ves_io_schema_tf_params.RunRequest{
		Namespace: "system",
		ViewKind:  actionParams.siteKind,
		ViewName:  actionParams.siteName,
		Action:    ves_io_schema_tf_params.RunAction(actionVal),
	}

	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return fmt.Errorf("Error marshalling rpc response to yaml: %s", err)
	}

	log.Printf("[DEBUG] Running terraform parameter action struct: %+v", req)
	_, err = client.CustomAPI(context.Background(), http.MethodPost,
		fmt.Sprintf(tfParamsActionRunURI, "system", actionParams.siteKind, actionParams.siteName), tfParamsActionRunRPCFQN, yamlReq)
	if err != nil {
		return fmt.Errorf("Error Running terraform parameter action struct: %s", err)
	}
	d.SetId("")
	return nil
}

func waitForActionToComplete(apiCl *APIClient, action, siteName, siteType string) (string, error) {
	log.Printf("[INFO] Waiting for action %s to complete on site %s", action, siteName)

	var prevState string
	for i := 0; i <= 30; i++ {
		statusResp, err := getTerraformStatus(apiCl, action, siteName, siteType)
		if err != nil {
			if action == "destroy" {
				if strings.Contains(err.Error(), "status code 404") {
					log.Printf("[DEBUG] Terraform job completed and also the site is deleted")
					return "", nil
				}
			}
			log.Printf("[DEBUG] GetStatus failed for site %s", siteName)
			time.Sleep(10 * time.Second)
			continue
		}
		log.Printf("[DEBUG] Status response: %#v", statusResp)
		switch action {
		case "plan":
			if statusResp.Status.PlanStatus == nil {
				log.Printf("[DEBUG] Plan State is not yet, wait and try again")
				time.Sleep(10 * time.Second)
				continue
			}
			switch statusResp.Status.PlanStatus.State.(type) {
			case *ves_io_schema_tf_params.PlanStatus_PlanState:
				planState := statusResp.Status.PlanStatus.State.(*ves_io_schema_tf_params.PlanStatus_PlanState).PlanState
				if planState == ves_io_schema_tf_params.PLANNING {
					if prevState != ves_io_schema_tf_params.PlanStageState_name[int32(planState)] {
						prevState = ves_io_schema_tf_params.PlanStageState_name[int32(planState)]
						log.Printf("[INFO] Plan State changed to %s", prevState)
					}
					time.Sleep(10 * time.Second)
					continue
				} else if planState == ves_io_schema_tf_params.PLAN_ERRORED || planState == ves_io_schema_tf_params.PLAN_INIT_ERRORED {
					log.Printf("[DEBUG] Unexpected plan state: %s", ves_io_schema_tf_params.PlanStageState_name[int32(planState)])
					log.Printf("[DEBUG] Terraform job stdout: \n%s", statusResp.Status.PlanStatus.TfPlanOutput)
					return "", fmt.Errorf("Error output from the terraform job: \n%s", statusResp.Status.PlanStatus.ErrorOutput)
				} else {
					if prevState != ves_io_schema_tf_params.PlanStageState_name[int32(planState)] {
						state := ves_io_schema_tf_params.PlanStageState_name[int32(planState)]
						log.Printf("[DEBUG] Plan Stage State changed to expected state %s", state)
						log.Printf("[DEBUG] Terraform job stdout: \n%s", statusResp.Status.PlanStatus.TfPlanOutput)
						return "", nil
					}
				}

			case *ves_io_schema_tf_params.PlanStatus_InfraState:
				infraState := statusResp.Status.PlanStatus.State.(*ves_io_schema_tf_params.PlanStatus_InfraState).InfraState
				var retry bool
				retry, prevState, err = parseInfraState(prevState, infraState)
				if retry {
					continue
				}
				return "", err
			}

		case "apply", "destroy":
			if statusResp.Status.ApplyStatus == nil {
				log.Printf("[DEBUG] Apply State is not yet set, wait and try again")
				time.Sleep(30 * time.Second)
				continue
			}
			switch statusResp.Status.ApplyStatus.State.(type) {
			case *ves_io_schema_tf_params.ApplyStatus_ApplyState:
				applyState := statusResp.Status.ApplyStatus.State.(*ves_io_schema_tf_params.ApplyStatus_ApplyState).ApplyState
				if applyState == ves_io_schema_tf_params.APPLY_PLANNING || applyState == ves_io_schema_tf_params.APPLYING {
					if prevState != ves_io_schema_tf_params.ApplyStageState_name[int32(applyState)] {
						prevState = ves_io_schema_tf_params.ApplyStageState_name[int32(applyState)]
						log.Printf("[INFO] Apply State changed to %s", prevState)
					}
					time.Sleep(30 * time.Second)
					continue
				} else if applyState == ves_io_schema_tf_params.APPLY_ERRORED ||
					applyState == ves_io_schema_tf_params.APPLY_INIT_ERRORED ||
					applyState == ves_io_schema_tf_params.APPLY_PLAN_ERRORED {
					log.Printf("[DEBUG] Unexpected apply state: %s", ves_io_schema_tf_params.ApplyStageState_name[int32(applyState)])
					log.Printf("[DEBUG] Terraform job stdout: \n%s", statusResp.Status.ApplyStatus.TfStdout)
					log.Printf("[DEBUG] Terraform output for the site: \n%s", statusResp.Status.ApplyStatus.TfOutput)
					return "", fmt.Errorf("Error output from the terraform job: \n%s", statusResp.Status.ApplyStatus.ErrorOutput)
				}
				state := ves_io_schema_tf_params.ApplyStageState_name[int32(applyState)]
				log.Printf("[DEBUG] Apply State changed to expected state %s", state)
				log.Printf("[DEBUG] Terraform job stdout: \n%s", statusResp.Status.ApplyStatus.TfStdout)
				return statusResp.Status.ApplyStatus.TfOutput, nil
			case *ves_io_schema_tf_params.ApplyStatus_DestroyState:
				destroyState := statusResp.Status.ApplyStatus.State.(*ves_io_schema_tf_params.ApplyStatus_DestroyState).DestroyState
				var retry bool
				retry, prevState, err = parseDestroyState(prevState, destroyState, statusResp.Status.ApplyStatus)
				if retry {
					continue
				}
				return "", err

			case *ves_io_schema_tf_params.ApplyStatus_InfraState:
				infraState := statusResp.Status.ApplyStatus.State.(*ves_io_schema_tf_params.ApplyStatus_InfraState).InfraState
				var retry bool
				retry, prevState, err = parseInfraState(prevState, infraState)
				if retry {
					continue
				}
				return "", err

			}
		}
	}
	log.Printf("[DEBUG] Timeout waiting for correct status")
	return "", nil
}

func getTerraformStatus(apiCl *APIClient, action, siteName,
	siteType string) (*ves_io_schema_tf_params.GetStatusResponse, error) {

	method := http.MethodGet
	uri := fmt.Sprintf(tfParamsGetStatusURI, "system", siteType, siteName)
	req := &ves_io_schema_tf_params.GetRequest{
		Namespace: "system",
		ViewKind:  siteType,
		ViewName:  siteName,
	}

	log.Printf("[DEBUG] Get status req: %#v", req)
	yamlReq, err := codec.ToYAML(req)
	if err != nil {
		return nil, errors.Wrap(err, "Converting request to yaml")
	}
	respProto, err := apiCl.CustomAPI(context.Background(), method, uri, tfParamsGetStatusRPCFQN, yamlReq)
	if err != nil {
		return nil, err
	}
	return respProto.(*ves_io_schema_tf_params.GetStatusResponse), nil
}

func parseInfraState(prevState string, infraState ves_io_schema_tf_params.InfraState) (bool, string, error) {
	if infraState == ves_io_schema_tf_params.PROVISIONING {
		if prevState != ves_io_schema_tf_params.InfraState_name[int32(infraState)] {
			prevState = ves_io_schema_tf_params.InfraState_name[int32(infraState)]
			log.Printf("[DEBUG] Infra State changed to %s", prevState)
		}
		time.Sleep(30 * time.Second)
		return true, prevState, nil
	} else if infraState == ves_io_schema_tf_params.ERRORED || infraState == ves_io_schema_tf_params.TIMED_OUT {
		log.Printf("[DEBUG] Unexpected infra state: %s", ves_io_schema_tf_params.InfraState_name[int32(infraState)])
		return false, prevState, fmt.Errorf("Error infra state: %s", ves_io_schema_tf_params.InfraState_name[int32(infraState)])
	}
	state := ves_io_schema_tf_params.InfraState_name[int32(infraState)]
	log.Printf("[DEBUG] Infra State changed to expected state %s", state)
	return false, prevState, nil
}

func parseDestroyState(prevState string,
	destroyState ves_io_schema_tf_params.DestroyStageState,
	applyStatus *ves_io_schema_tf_params.ApplyStatus) (bool, string, error) {
	if destroyState == ves_io_schema_tf_params.DESTROYING {
		if prevState != ves_io_schema_tf_params.DestroyStageState_name[int32(destroyState)] {
			prevState = ves_io_schema_tf_params.DestroyStageState_name[int32(destroyState)]
			log.Printf("[DEBUG] Destroy State changed to %s", prevState)
		}
		time.Sleep(30 * time.Second)
		return true, prevState, nil
	} else if destroyState == ves_io_schema_tf_params.DESTROY_ERRORED {
		log.Printf("[DEBUG] Unexpected destroy state: %s", ves_io_schema_tf_params.DestroyStageState_name[int32(destroyState)])
		log.Print(fmt.Sprintf("Terraform job stdout: \n%s", applyStatus.TfStdout))
		return false, prevState, fmt.Errorf("Error output from the terraform job: \n%s", applyStatus.ErrorOutput)
	}
	state := ves_io_schema_tf_params.DestroyStageState_name[int32(destroyState)]
	log.Printf("[DEBUG] Destroy State changed to expected state %s", state)
	return false, prevState, nil
}
