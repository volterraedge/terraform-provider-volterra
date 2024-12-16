// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"context"
	"fmt"
	"math/rand"
	"testing"
	"time"

	"github.com/gogo/protobuf/types"
	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/acctest"

	"gopkg.volterra.us/stdlib/client/vesapi"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
	svcfw_test "gopkg.volterra.us/stdlib/svcfw/test"
	"gopkg.volterra.us/stdlib/svcfw/test/generic"
	"gopkg.volterra.us/stdlib/testutil/fixtures/vesenv"

	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_api_credential "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/api_credential"
	ves_io_schema_combined "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/combined"
	ves_io_schema_dns_zone_rrset "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/dns_zone/rrset"
	ves_io_schema_known_label "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/known_label"
	ves_io_schema_known_label_key "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/known_label_key"
	ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
	ves_io_schema_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/site"
	ves_io_schema_tenant "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/tenant"
	"github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views"
	ves_io_schema_aws_tgw_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_tgw_site"
	ves_io_schema_aws_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/aws_vpc_site"
	ves_io_schema_azure_vnet_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/azure_vnet_site"
	ves_io_schema_gcp_vpc_site "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/gcp_vpc_site"
	ves_io_schema_tf_params "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/views/terraform_parameters"
	ves_io_schema_vh "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/virtual_host"
)

var (
	tdRoot = "../testdata"
	// APIStimulusSet holds all API stimulus set in init() functions per object type, keyed by API FQN
	APIStimulusSet = map[string]APIStimulus{}
	// tenant in which all tests run, can be changed before running test
	tenant = "acmecorp"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

// APIStimulus is a type holding all RPC stimulus keyed by RPC Name
type APIStimulus map[string]*RPCStimulus

// RPCStimulus holds the APIOpertor s to exercise the RPC and associated documentation
type RPCStimulus struct {
	APIOps []vesenv.APIOperator
}

type ctxSvcFKey struct{}

// ves.io.schema.api_credential.CustomAPI handling - start

type apiCredentialCustomAPIServer struct {
	sf svcfw.Service
}

func randomFunc() string {
	source := rand.NewSource(time.Now().UnixNano())
	random := rand.New(source)
	length := 7
	result := make([]byte, length)
	for i := 0; i < length; i++ {
		result[i] = charset[random.Intn(len(charset))]
	}
	randomString := string(result)
	return randomString
}

func newAPICredentialCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &apiCredentialCustomAPIServer{sf: sf}
}

var (
	rspCreateResponse = ves_io_schema_api_credential.CreateResponse{}
)

func getCreateResponseAPICredential(name string) {
	timestamp := &types.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(time.Now().Nanosecond()),
	}

	rspCreateResponse = ves_io_schema_api_credential.CreateResponse{
		Name:                fmt.Sprintf("%s-%s", name, randomFunc()),
		ExpirationTimestamp: timestamp,
	}
}

func (s *apiCredentialCustomAPIServer) Create(ctx context.Context, req *ves_io_schema_api_credential.CreateRequest) (*ves_io_schema_api_credential.CreateResponse, error) {
	return &rspCreateResponse, nil
}

func (s *apiCredentialCustomAPIServer) Get(ctx context.Context, req *ves_io_schema_api_credential.GetRequest) (*ves_io_schema_api_credential.GetResponse, error) {
	timestamp := &types.Timestamp{
		Seconds: time.Now().Unix(),
		Nanos:   int32(time.Now().Nanosecond()),
	}
	return &ves_io_schema_api_credential.GetResponse{
		Object: &ves_io_schema_api_credential.Object{
			Spec: &ves_io_schema_api_credential.SpecType{
				GcSpec: &ves_io_schema_api_credential.GlobalSpecType{
					ExpirationTimestamp: timestamp,
				},
			},
		},
	}, nil
}

func (s *apiCredentialCustomAPIServer) List(ctx context.Context, req *ves_io_schema_api_credential.ListRequest) (*ves_io_schema_api_credential.ListResponse, error) {
	return &ves_io_schema_api_credential.ListResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) Revoke(ctx context.Context, req *ves_io_schema_api_credential.GetRequest) (*ves_io_schema_api_credential.StatusResponse, error) {
	return &ves_io_schema_api_credential.StatusResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) Activate(ctx context.Context, req *ves_io_schema_api_credential.GetRequest) (*ves_io_schema_api_credential.StatusResponse, error) {
	return &ves_io_schema_api_credential.StatusResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) Renew(ctx context.Context, req *ves_io_schema_api_credential.RenewRequest) (*ves_io_schema_api_credential.StatusResponse, error) {
	return &ves_io_schema_api_credential.StatusResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) CreateServiceCredentials(ctx context.Context, req *ves_io_schema_api_credential.CreateServiceCredentialsRequest) (*ves_io_schema_api_credential.CreateResponse, error) {
	return &ves_io_schema_api_credential.CreateResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) ListServiceCredentials(ctx context.Context, req *ves_io_schema_api_credential.ListRequest) (*ves_io_schema_api_credential.ListResponse, error) {
	return &ves_io_schema_api_credential.ListResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) ActivateServiceCredentials(ctx context.Context, req *ves_io_schema_api_credential.GetRequest) (*ves_io_schema_api_credential.StatusResponse, error) {
	return &ves_io_schema_api_credential.StatusResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) RenewServiceCredentials(ctx context.Context, req *ves_io_schema_api_credential.RenewRequest) (*ves_io_schema_api_credential.StatusResponse, error) {
	return &ves_io_schema_api_credential.StatusResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) RevokeServiceCredentials(ctx context.Context, req *ves_io_schema_api_credential.GetRequest) (*ves_io_schema_api_credential.StatusResponse, error) {
	return &ves_io_schema_api_credential.StatusResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) RecreateScimToken(ctx context.Context, req *ves_io_schema_api_credential.RecreateScimTokenRequest) (*ves_io_schema_api_credential.CreateResponse, error) {
	return &ves_io_schema_api_credential.CreateResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) RevokeScimToken(ctx context.Context, req *ves_io_schema_api_credential.ScimTokenRequest) (*ves_io_schema_api_credential.StatusResponse, error) {
	return &ves_io_schema_api_credential.StatusResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) GetScimToken(ctx context.Context, req *ves_io_schema_api_credential.ScimTokenRequest) (*ves_io_schema_api_credential.GetResponse, error) {
	return &ves_io_schema_api_credential.GetResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) GetServiceCredentials(ctx context.Context, req *ves_io_schema_api_credential.GetRequest) (*ves_io_schema_api_credential.GetServiceCredentialsResponse, error) {
	return &ves_io_schema_api_credential.GetServiceCredentialsResponse{}, nil
}

func (s *apiCredentialCustomAPIServer) ReplaceServiceCredentials(context.Context, *ves_io_schema_api_credential.ReplaceServiceCredentialsRequest) (*ves_io_schema_api_credential.ReplaceServiceCredentialsResponse, error) {
	return &ves_io_schema_api_credential.ReplaceServiceCredentialsResponse{}, nil
}

var _ ves_io_schema_api_credential.CustomAPIServer = &apiCredentialCustomAPIServer{}

// ves.io.schema.virtual_host.CustomAPI handling - start

type vhCustomAPIServer struct {
	sf svcfw.Service
}

func newVHCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &vhCustomAPIServer{sf: sf}
}

// GetDnsInfo custom api implementation
func (s *vhCustomAPIServer) GetDnsInfo(ctx context.Context,
	req *ves_io_schema_vh.GetDnsInfoRequest) (*ves_io_schema_vh.GetDnsInfoResponse, error) {
	return &ves_io_schema_vh.GetDnsInfoResponse{}, nil
}

func (s *vhCustomAPIServer) AssignAPIDefinition(context.Context, *ves_io_schema_vh.AssignAPIDefinitionReq) (*ves_io_schema_vh.AssignAPIDefinitionResp, error) {
	return &ves_io_schema_vh.AssignAPIDefinitionResp{}, nil
}

var _ ves_io_schema_vh.CustomAPIServer = &vhCustomAPIServer{}

// ves.io.schema.virtual_host.CustomAPI handling - end

// ves.io.schema.site.CustomAPI handling - start
type siteCustomAPIServer struct {
	sf svcfw.Service
}

func newSiteCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &siteCustomAPIServer{sf: sf}
}

func (s *siteCustomAPIServer) SetState(ctx context.Context,
	req *ves_io_schema_site.SetStateReq) (*ves_io_schema_site.SetStateResp, error) {
	return &ves_io_schema_site.SetStateResp{}, nil
}

var _ ves_io_schema_site.CustomStateAPIServer = &siteCustomAPIServer{}

// ves.io.schema.site.CustomAPI handling - end

// ves.io.schema.namespace.CustomAPI handling - start
type nsCustomAPIServer struct {
	sf svcfw.Service
}

func newNSCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &nsCustomAPIServer{sf: sf}
}

func (n *nsCustomAPIServer) CascadeDelete(context.Context,
	*ves_io_schema_ns.CascadeDeleteRequest) (*ves_io_schema_ns.CascadeDeleteResponse, error) {
	return &ves_io_schema_ns.CascadeDeleteResponse{}, nil
}

func (n *nsCustomAPIServer) EvaluateAPIAccess(context.Context,
	*ves_io_schema_ns.EvaluateAPIAccessReq) (*ves_io_schema_ns.EvaluateAPIAccessResp, error) {
	return &ves_io_schema_ns.EvaluateAPIAccessResp{}, nil
}

func (n *nsCustomAPIServer) EvaluateBatchAPIAccess(context.Context, *ves_io_schema_ns.EvaluateBatchAPIAccessReq) (*ves_io_schema_ns.EvaluateBatchAPIAccessResp, error) {
	return &ves_io_schema_ns.EvaluateBatchAPIAccessResp{}, nil
}

func (n *nsCustomAPIServer) SuggestValues(context.Context,
	*ves_io_schema_ns.SuggestValuesReq) (*ves_io_schema_ns.SuggestValuesResp, error) {
	return &ves_io_schema_ns.SuggestValuesResp{}, nil
}

func (n *nsCustomAPIServer) LookupUserRoles(context.Context,
	*ves_io_schema_ns.LookupUserRolesReq) (*ves_io_schema_ns.LookupUserRolesResp, error) {
	return &ves_io_schema_ns.LookupUserRolesResp{}, nil
}

var _ ves_io_schema_ns.CustomAPIServer = &nsCustomAPIServer{}

// ves.io.schema.namespace.CustomAPI handling - end

// ves.io.schema.views.aws_vpc_site.CustomAPI handling - start
type vpcCustomAPIServer struct {
	sf svcfw.Service
}

func newVPCCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &vpcCustomAPIServer{sf: sf}
}

func (t *vpcCustomAPIServer) SetVPCK8SHostnames(context.Context,
	*ves_io_schema_aws_vpc_site.SetVPCK8SHostnamesRequest) (*ves_io_schema_aws_vpc_site.SetVPCK8SHostnamesResponse, error) {
	return &ves_io_schema_aws_vpc_site.SetVPCK8SHostnamesResponse{}, nil

}

func (t *vpcCustomAPIServer) SetVIPInfo(context.Context,
	*ves_io_schema_aws_vpc_site.SetVIPInfoRequest) (*ves_io_schema_aws_vpc_site.SetVIPInfoResponse, error) {
	return &ves_io_schema_aws_vpc_site.SetVIPInfoResponse{}, nil

}

func (t *vpcCustomAPIServer) SetCloudSiteInfo(context.Context,
	*ves_io_schema_aws_vpc_site.SetCloudSiteInfoRequest) (*ves_io_schema_aws_vpc_site.SetCloudSiteInfoResponse, error) {
	return &ves_io_schema_aws_vpc_site.SetCloudSiteInfoResponse{}, nil

}

func (t *vpcCustomAPIServer) ValidateConfig(context.Context, *views.ValidateConfigRequest) (*views.ValidateConfigResponse, error) {
	return &views.ValidateConfigResponse{}, nil
}

var _ ves_io_schema_aws_vpc_site.CustomAPIServer = &vpcCustomAPIServer{}

// ves.io.schema.views.azure_vnet_site.CustomAPI handling - start
type vnetCustomAPIServer struct {
	sf svcfw.Service
}

func newVnetCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &vnetCustomAPIServer{sf: sf}
}

func (t *vnetCustomAPIServer) SetVIPInfo(context.Context,
	*ves_io_schema_azure_vnet_site.SetVIPInfoRequest) (*ves_io_schema_azure_vnet_site.SetVIPInfoResponse, error) {
	return &ves_io_schema_azure_vnet_site.SetVIPInfoResponse{}, nil

}

func (t *vnetCustomAPIServer) SetCloudSiteInfo(context.Context,
	*ves_io_schema_azure_vnet_site.SetCloudSiteInfoRequest) (*ves_io_schema_azure_vnet_site.SetCloudSiteInfoResponse, error) {
	return &ves_io_schema_azure_vnet_site.SetCloudSiteInfoResponse{}, nil

}

func (t *vnetCustomAPIServer) ValidateConfig(context.Context, *views.ValidateConfigRequest) (*views.ValidateConfigResponse, error) {
	return &views.ValidateConfigResponse{}, nil
}

var _ ves_io_schema_azure_vnet_site.CustomAPIServer = &vnetCustomAPIServer{}

// ves.io.schema.views.azure_vnet_site.CustomAPI handling - end

// ves.io.schema.views.aws_tgw_site.CustomAPI handling - start
type tgwCustomAPIServer struct {
	sf svcfw.Service
}

func newTGWCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &tgwCustomAPIServer{sf: sf}
}

func (t *tgwCustomAPIServer) SetVPCIpPrefixes(ctx context.Context,
	r *ves_io_schema_aws_tgw_site.SetVPCIpPrefixesRequest) (*ves_io_schema_aws_tgw_site.SetVPCIpPrefixesResponse, error) {
	tenant := server.TenantFromContext(ctx)
	if tenant == "" {
		return nil, fmt.Errorf("Tenant missing in request")
	}

	e, ok, err := t.sf.FindEntry(ctx, ves_io_schema_aws_tgw_site.ObjectDefTblName, "", db.WithFindUsingNSIndex(tenant, "system", r.Name))
	if err != nil {
		return nil, fmt.Errorf("cannot find aws_tgw_site entry %s, err: %s", r.Name, err)
	}
	if !ok {
		return nil, fmt.Errorf("aws_tgw_site %s does not exist", r.Name)
	}
	obj := e.(*ves_io_schema_aws_tgw_site.DBObject)
	obj.Spec.GcSpec.VpcIpPrefixes = r.GetVpcIpPrefixes()
	if _, err := t.sf.ChgEntry(ctx, obj.ToEntry()); err != nil {
		return nil, fmt.Errorf("Could not update aws_tgw_site %s", err)
	}
	return &ves_io_schema_aws_tgw_site.SetVPCIpPrefixesResponse{}, nil

}
func (t *tgwCustomAPIServer) SetVPNTunnels(context.Context,
	*ves_io_schema_aws_tgw_site.SetVPNTunnelsRequest) (*ves_io_schema_aws_tgw_site.SetVPNTunnelsResponse, error) {
	return &ves_io_schema_aws_tgw_site.SetVPNTunnelsResponse{}, nil
}

func (t *tgwCustomAPIServer) SetTGWInfo(context.Context,
	*ves_io_schema_aws_tgw_site.SetTGWInfoRequest) (*ves_io_schema_aws_tgw_site.SetTGWInfoResponse, error) {
	return &ves_io_schema_aws_tgw_site.SetTGWInfoResponse{}, nil
}

func (t *tgwCustomAPIServer) SetVIPInfo(context.Context,
	*ves_io_schema_aws_tgw_site.SetVIPInfoRequest) (*ves_io_schema_aws_tgw_site.SetVIPInfoResponse, error) {
	return &ves_io_schema_aws_tgw_site.SetVIPInfoResponse{}, nil
}

func (t *tgwCustomAPIServer) ValidateConfig(context.Context, *views.ValidateConfigRequest) (*views.ValidateConfigResponse, error) {
	return &views.ValidateConfigResponse{}, nil
}

var _ ves_io_schema_aws_tgw_site.CustomAPIServer = &tgwCustomAPIServer{}

// ves.io.schema.views.aws_tgw_site.CustomAPI handling - end

// ves.io.schema.views.terraform_parameters.CustomAPI handling - start
type tfCustomAPIServer struct {
	sf svcfw.Service
}

func newTFCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &tfCustomAPIServer{sf: sf}
}

func (t *tfCustomAPIServer) Get(ctx context.Context, req *ves_io_schema_tf_params.GetRequest) (*ves_io_schema_tf_params.GetResponse, error) {
	return &ves_io_schema_tf_params.GetResponse{}, nil

}
func (t *tfCustomAPIServer) GetStatus(ctx context.Context, req *ves_io_schema_tf_params.GetRequest) (*ves_io_schema_tf_params.GetStatusResponse, error) {
	return &ves_io_schema_tf_params.GetStatusResponse{}, nil
}

var _ ves_io_schema_tf_params.CustomAPIServer = &tfCustomAPIServer{}

// ves.io.schema.views.terraform_parameters.CustomAPI handling - end

// ves.io.schema.views.terraform_parameters.CustomActionAPI handling - start
type tfCustomActionAPIServer struct {
	sf svcfw.Service
}

func newTFCustomActionAPIServer(sf svcfw.Service) server.APIHandler {
	return &tfCustomActionAPIServer{sf: sf}
}

func (t *tfCustomActionAPIServer) Run(ctx context.Context, req *ves_io_schema_tf_params.RunRequest) (*ves_io_schema_tf_params.RunResponse, error) {
	return &ves_io_schema_tf_params.RunResponse{}, nil

}
func (t *tfCustomActionAPIServer) ForceDelete(ctx context.Context, req *ves_io_schema_tf_params.ForceDeleteRequest) (*ves_io_schema_tf_params.ForceDeleteResponse, error) {
	return &ves_io_schema_tf_params.ForceDeleteResponse{}, nil
}

var _ ves_io_schema_tf_params.CustomActionAPIServer = &tfCustomActionAPIServer{}

// ves.io.schema.views.terraform_parameters.CustomAPI handling - end

// ves.io.schema.namespace.NamespaceCustomAPI handling - start
type namespaceCustomAPIServer struct {
	sf svcfw.Service
}

func newNamespaceCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &namespaceCustomAPIServer{sf: sf}
}

func (t *namespaceCustomAPIServer) SetFastACLsForInternetVIPs(ctx context.Context, req *ves_io_schema_ns.SetFastACLsForInternetVIPsRequest) (*ves_io_schema_ns.SetFastACLsForInternetVIPsResponse, error) {
	return &ves_io_schema_ns.SetFastACLsForInternetVIPsResponse{}, nil

}
func (t *namespaceCustomAPIServer) GetFastACLsForInternetVIPs(ctx context.Context, req *ves_io_schema_ns.GetFastACLsForInternetVIPsRequest) (*ves_io_schema_ns.GetFastACLsForInternetVIPsResponse, error) {
	return &ves_io_schema_ns.GetFastACLsForInternetVIPsResponse{}, nil
}

func (t *namespaceCustomAPIServer) SetActiveServicePolicies(ctx context.Context, req *ves_io_schema_ns.SetActiveServicePoliciesRequest) (*ves_io_schema_ns.SetActiveServicePoliciesResponse, error) {
	return &ves_io_schema_ns.SetActiveServicePoliciesResponse{}, nil

}
func (t *namespaceCustomAPIServer) GetActiveServicePolicies(ctx context.Context, req *ves_io_schema_ns.GetActiveServicePoliciesRequest) (*ves_io_schema_ns.GetActiveServicePoliciesResponse, error) {
	return &ves_io_schema_ns.GetActiveServicePoliciesResponse{}, nil
}

func (t *namespaceCustomAPIServer) SetActiveNetworkPolicies(ctx context.Context, req *ves_io_schema_ns.SetActiveNetworkPoliciesRequest) (*ves_io_schema_ns.SetActiveNetworkPoliciesResponse, error) {
	return &ves_io_schema_ns.SetActiveNetworkPoliciesResponse{}, nil

}
func (t *namespaceCustomAPIServer) GetActiveNetworkPolicies(ctx context.Context, req *ves_io_schema_ns.GetActiveNetworkPoliciesRequest) (*ves_io_schema_ns.GetActiveNetworkPoliciesResponse, error) {
	return &ves_io_schema_ns.GetActiveNetworkPoliciesResponse{}, nil
}

func (t *namespaceCustomAPIServer) SetActiveAlertPolicies(ctx context.Context, req *ves_io_schema_ns.SetActiveAlertPoliciesRequest) (*ves_io_schema_ns.SetActiveAlertPoliciesResponse, error) {
	return &ves_io_schema_ns.SetActiveAlertPoliciesResponse{}, nil

}
func (t *namespaceCustomAPIServer) GetActiveAlertPolicies(ctx context.Context, req *ves_io_schema_ns.GetActiveAlertPoliciesRequest) (*ves_io_schema_ns.GetActiveAlertPoliciesResponse, error) {
	return &ves_io_schema_ns.GetActiveAlertPoliciesResponse{}, nil
}

func (t *namespaceCustomAPIServer) CascadeDelete(ctx context.Context, req *ves_io_schema_ns.CascadeDeleteRequest) (*ves_io_schema_ns.CascadeDeleteResponse, error) {
	return &ves_io_schema_ns.CascadeDeleteResponse{}, nil
}

func (t *namespaceCustomAPIServer) SuggestValues(ctx context.Context, req *ves_io_schema_ns.SuggestValuesReq) (*ves_io_schema_ns.SuggestValuesResp, error) {
	return &ves_io_schema_ns.SuggestValuesResp{}, nil
}

func (t *namespaceCustomAPIServer) ValidateRules(ctx context.Context, req *ves_io_schema_ns.ValidateRulesReq) (*ves_io_schema_ns.ValidateRulesResponse, error) {
	return &ves_io_schema_ns.ValidateRulesResponse{}, nil
}

func (n *namespaceCustomAPIServer) UpdateAllowAdvertiseOnPublic(context.Context,
	*ves_io_schema_ns.UpdateAllowAdvertiseOnPublicReq) (*ves_io_schema_ns.UpdateAllowAdvertiseOnPublicResp, error) {
	return &ves_io_schema_ns.UpdateAllowAdvertiseOnPublicResp{}, nil
}

func (n *namespaceCustomAPIServer) NetworkingInventory(context.Context,
	*ves_io_schema_ns.NetworkingInventoryRequest) (*ves_io_schema_ns.NetworkingInventoryResponse, error) {
	return &ves_io_schema_ns.NetworkingInventoryResponse{}, nil
}

func (n *namespaceCustomAPIServer) ApplicationInventory(context.Context,
	*ves_io_schema_ns.ApplicationInventoryRequest) (*ves_io_schema_ns.ApplicationInventoryResponse, error) {
	return &ves_io_schema_ns.ApplicationInventoryResponse{}, nil
}

func (n *namespaceCustomAPIServer) AllApplicationInventory(context.Context,
	*ves_io_schema_ns.AllApplicationInventoryRequest) (*ves_io_schema_ns.ApplicationInventoryResponse, error) {
	return &ves_io_schema_ns.ApplicationInventoryResponse{}, nil
}

func (n *namespaceCustomAPIServer) AllApplicationInventoryWaf(context.Context,
	*ves_io_schema_ns.AllApplicationInventoryWafFilterRequest) (*ves_io_schema_ns.AllApplicationInventoryWafFilterResponse, error) {
	return &ves_io_schema_ns.AllApplicationInventoryWafFilterResponse{}, nil
}

var _ ves_io_schema_ns.NamespaceCustomAPIServer = &namespaceCustomAPIServer{}

// ves.io.schema.views.gcp_vpc_site.CustomAPI handling - start
type gcpVpcCustomAPIServer struct {
	sf svcfw.Service
}

func newGcpVpcCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &gcpVpcCustomAPIServer{sf: sf}
}

func (t *gcpVpcCustomAPIServer) SetCloudSiteInfo(context.Context,
	*ves_io_schema_gcp_vpc_site.SetCloudSiteInfoRequest) (*ves_io_schema_gcp_vpc_site.SetCloudSiteInfoResponse, error) {
	return &ves_io_schema_gcp_vpc_site.SetCloudSiteInfoResponse{}, nil

}

func (t *gcpVpcCustomAPIServer) ValidateConfig(context.Context, *views.ValidateConfigRequest) (*views.ValidateConfigResponse, error) {
	return &views.ValidateConfigResponse{}, nil
}

var _ ves_io_schema_gcp_vpc_site.CustomAPIServer = &gcpVpcCustomAPIServer{}

// ves.io.schema.namespace.NamespaceCustomAPI handling - end

func getFixtureOpts(noTLS bool) []generic.ConfigOpt {
	fcOpts := append(generic.GetFixtureTLSOpts(noTLS, tdRoot),
		generic.WithCfgSTFConfigOpts(
			svcfw_test.WithCfgKindToTypeFn(func(refrType, refrUID, refrKind string) (string, error) {
				return "ves.io.schema." + refrKind, nil
			}),
		))
	return fcOpts
}

func createTestServer(t *testing.T, objectType string) (string, func()) {

	f, stop := makeTestServer(t, objectType)
	return fmt.Sprintf("https://localhost:%d", f.Svc.RestServerTLSPort()), stop
}

func makeTestServer(t *testing.T, objectType string) (*generic.Fixture, func()) {
	svcfw.SkipRestToRPCMap = append(svcfw.SkipRestToRPCMap,
		"ves.io.schema.implicit_label.CustomAPI.Get",
		"ves.io.schema.status_at_site.CustomAPI.GetStatus",
		"ves.io.schema.waf.WAFMonitoringAPI.SecurityEventQuery",
		"ves.io.schema.waf.WAFMonitoringAPI.SecurityEventScrollQuery",
	)

	fcOpts := append(getFixtureOpts(false),
		generic.WithCfgSTFConfigOpts(
			svcfw_test.WithCfgServedTypes(store.InMemory, objectType),
			svcfw_test.WithCfgPublicAPITypes(objectType),
			//svcfw_test.WithCfgNewModuleFn(generic.NewGenericModule),
		),
	)
	f := generic.NewFixture(t, ves_io_schema_combined.MDR, fcOpts...)
	stop := f.MustStart(f.MustNewConfig(t.Name()))
	return f, stop
}

func createTestCustomAPIServer(t *testing.T, objectTypes []string) (string, func(), *generic.Fixture) {

	f, stop := makeCustomTestServer(t, objectTypes)
	return fmt.Sprintf("https://localhost:%d", f.Svc.RestServerTLSPort()), stop, f
}

func makeCustomTestServer(t *testing.T, objectTypes []string) (*generic.Fixture, func()) {
	svcfw.SkipRestToRPCMap = append(svcfw.SkipRestToRPCMap,
		"ves.io.schema.implicit_label.CustomAPI.Get",
		"ves.io.schema.status_at_site.CustomAPI.GetStatus",
		"ves.io.schema.waf.WAFMonitoringAPI.SecurityEventQuery",
		"ves.io.schema.waf.WAFMonitoringAPI.SecurityEventScrollQuery",
	)

	customHandlers := map[string]func(svcfw.Service) server.APIHandler{
		// alphabetically sorted list of custom APIs to struct that handles the RPCs
		"ves.io.schema.api_credential.CustomAPI":                   newAPICredentialCustomAPIServer,
		"ves.io.schema.known_label.CustomAPI":                      newKnownLabelCustomAPIServer,
		"ves.io.schema.known_label_key.CustomAPI":                  newKnownLabelKeyCustomAPIServer,
		"ves.io.schema.namespace.CustomAPI":                        newNSCustomAPIServer,
		"ves.io.schema.virtual_host.CustomAPI":                     newVHCustomAPIServer,
		"ves.io.schema.site.CustomStateAPI":                        newSiteCustomAPIServer,
		"ves.io.schema.views.aws_tgw_site.CustomAPI":               newTGWCustomAPIServer,
		"ves.io.schema.views.aws_vpc_site.CustomAPI":               newVPCCustomAPIServer,
		"ves.io.schema.views.gcp_vpc_site.CustomAPI":               newGcpVpcCustomAPIServer,
		"ves.io.schema.views.azure_vnet_site.CustomAPI":            newVnetCustomAPIServer,
		"ves.io.schema.views.terraform_parameters.CustomAPI":       newTFCustomAPIServer,
		"ves.io.schema.views.terraform_parameters.CustomActionAPI": newTFCustomActionAPIServer,
		"ves.io.schema.namespace.NamespaceCustomAPI":               newNamespaceCustomAPIServer,
		"ves.io.schema.dns_zone.rrset.CustomAPI":                   newDnsZoneRecordCustomAPIServer,
	}

	// bail if there isn't a handler for every possible public custom API defined in schema repo
	missingHandlers := []string{}
	customSvrModFns := svcfw_test.ModuleFuncs{
		NameFn: func(sf svcfw.Service) string {
			return "all-customAPIs-server"
		},
		GetCustomAPIHandlersFn: func(sf svcfw.Service) map[string]server.APIHandler {
			// map of all customAPI protobuf service to struct implementing server
			allCustomHandlers := map[string]server.APIHandler{}
			for apiFQN := range sf.MDRegistry().PubCustomServiceRegistry.RestClientRegistry {
				handlerBuilder, exists := customHandlers[apiFQN]
				if !exists {
					missingHandlers = append(missingHandlers, apiFQN)
					continue
				}
				allCustomHandlers[apiFQN] = handlerBuilder(sf)
			}
			//if len(missingHandlers) != 0 {
			//	t.Fatalf("No handlers provided for custom APIs %s", missingHandlers)
			//}
			return allCustomHandlers
		},
	}
	customSvrMod := svcfw_test.NewModuleFactory(customSvrModFns)

	svcFWOpts := make([]svcfw_test.ConfigOpt, 0, 0)
	for _, obj := range objectTypes {
		svcFWOpts = append(svcFWOpts, svcfw_test.WithCfgPublicAPITypes(obj))
		svcFWOpts = append(svcFWOpts, svcfw_test.WithCfgServedTypes(store.InMemory, obj))
	}
	svcFWOpts = append(svcFWOpts,
		svcfw_test.WithCfgNewModuleFn(generic.NewGenericModule),
		svcfw_test.WithCfgNewModuleFn(customSvrMod),
	)
	svcFWOpts = append(svcFWOpts, svcfw_test.WithCfgKindToTypeFn(func(refrType, refrUID, refrKind string) (string, error) {
		return "ves.io.schema." + refrKind, nil
	}))

	configOpts := append(getFixtureOpts(false),
		generic.WithCfgSTFConfigOpts(svcFWOpts...))

	f := generic.NewFixture(t, ves_io_schema_combined.MDR, configOpts...)
	stop := f.MustStart(f.MustNewConfig(t.Name()))
	return f, stop
}

func getTestClientOpts(url, tenant string) ([]vesapi.ConfigOpt, error) {

	clKey := fmt.Sprintf("file:///%s/tls/client.examplesvc.ves.io.key", tdRoot)
	clCert := fmt.Sprintf("file:///%s/tls/client.examplesvc.ves.io.crt", tdRoot)
	clCACert := fmt.Sprintf("file:///%s/tls/client.ca.crt", tdRoot)

	remapperFn := uriRemapperFn(url, tenant)

	clOpts := []vesapi.ConfigOpt{
		vesapi.WithCfgKey(clKey),
		vesapi.WithCfgCert(clCert),
		vesapi.WithCfgCACert(clCACert),
		vesapi.WithCfgRoundTripperFn(remapperFn),
	}
	vesapi.ServerCN = "examplesvc.ves.io"

	return clOpts, nil
}

func generateResourceName() string {
	return acctest.RandStringFromCharSet(10, acctest.CharSetAlpha)
}

type TenantOption func(pbObj *ves_io_schema_tenant.Object)

func withAddonServices(addons ...string) TenantOption {
	return func(pbObj *ves_io_schema_tenant.Object) {
		pbObj.Spec.GcSpec.AddonServicesSubscribed = addons
	}
}

// mkDBObjTenant creates an instance of *ves_io_schema_tenant.DBObject
func mkDBObjTenant(name, uid string, opts ...TenantOption) *ves_io_schema_tenant.DBObject {
	pbObj := &ves_io_schema_tenant.Object{
		Metadata: &ves_io_schema.ObjectMetaType{
			Name:      name,
			Uid:       uid,
			Namespace: "system",
		},
		SystemMetadata: &ves_io_schema.SystemObjectMetaType{
			Tenant: "ves-io",
			Uid:    uid,
			Namespace: []*ves_io_schema.ObjectRefType{
				{
					Tenant: name,
					Kind:   "namespace",
					Name:   "system",
				},
			},
		},
		Spec: &ves_io_schema_tenant.SpecType{
			GcSpec: &ves_io_schema_tenant.GlobalSpecType{},
		},
	}
	for _, opt := range opts {
		opt(pbObj)
	}
	return ves_io_schema_tenant.NewDBObject(pbObj)
}

// ves.io.schema.known_label.CustomAPI handling - start
type dnsZoneRecordCustomAPIServer struct {
	sf svcfw.Service
}

func newDnsZoneRecordCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &dnsZoneRecordCustomAPIServer{sf: sf}
}

func (s *dnsZoneRecordCustomAPIServer) Create(ctx context.Context, req *ves_io_schema_dns_zone_rrset.CreateRequest) (*ves_io_schema_dns_zone_rrset.Response, error) {
	return &ves_io_schema_dns_zone_rrset.Response{}, nil
}

func (s *dnsZoneRecordCustomAPIServer) Get(ctx context.Context, req *ves_io_schema_dns_zone_rrset.GetRequest) (*ves_io_schema_dns_zone_rrset.Response, error) {
	return &ves_io_schema_dns_zone_rrset.Response{}, nil
}

func (s *dnsZoneRecordCustomAPIServer) Replace(ctx context.Context, req *ves_io_schema_dns_zone_rrset.ReplaceRequest) (*ves_io_schema_dns_zone_rrset.Response, error) {
	return &ves_io_schema_dns_zone_rrset.Response{}, nil
}

func (s *dnsZoneRecordCustomAPIServer) Delete(ctx context.Context, req *ves_io_schema_dns_zone_rrset.DeleteRequest) (*ves_io_schema_dns_zone_rrset.Response, error) {
	return &ves_io_schema_dns_zone_rrset.Response{}, nil
}

var _ ves_io_schema_dns_zone_rrset.CustomAPIServer = &dnsZoneRecordCustomAPIServer{}

// ves.io.schema.known_label.CustomAPI handling - start
type knownLabelCustomAPIServer struct {
	sf svcfw.Service
}

func newKnownLabelCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &knownLabelCustomAPIServer{sf: sf}
}

func (s *knownLabelCustomAPIServer) Create(ctx context.Context, req *ves_io_schema_known_label.CreateRequest) (*ves_io_schema_known_label.CreateResponse, error) {
	return &ves_io_schema_known_label.CreateResponse{}, nil
}

func (s *knownLabelCustomAPIServer) Get(ctx context.Context, req *ves_io_schema_known_label.GetRequest) (*ves_io_schema_known_label.GetResponse, error) {
	return &ves_io_schema_known_label.GetResponse{
		Label: []*ves_io_schema_known_label.LabelType{
			{
				Key:   req.Key,
				Value: req.Value,
			},
		},
	}, nil
}

func (s *knownLabelCustomAPIServer) Delete(ctx context.Context, req *ves_io_schema_known_label.DeleteRequest) (*ves_io_schema_known_label.DeleteResponse, error) {
	return &ves_io_schema_known_label.DeleteResponse{}, nil
}

var _ ves_io_schema_known_label.CustomAPIServer = &knownLabelCustomAPIServer{}

// ves.io.schema.known_label.CustomAPI handling - start
type knownLabelKeyCustomAPIServer struct {
	sf svcfw.Service
}

func newKnownLabelKeyCustomAPIServer(sf svcfw.Service) server.APIHandler {
	return &knownLabelKeyCustomAPIServer{sf: sf}
}

func (s *knownLabelKeyCustomAPIServer) Create(ctx context.Context, req *ves_io_schema_known_label_key.CreateRequest) (*ves_io_schema_known_label_key.CreateResponse, error) {
	return &ves_io_schema_known_label_key.CreateResponse{}, nil
}

func (s *knownLabelKeyCustomAPIServer) Get(ctx context.Context, req *ves_io_schema_known_label_key.GetRequest) (*ves_io_schema_known_label_key.GetResponse, error) {
	return &ves_io_schema_known_label_key.GetResponse{
		LabelKey: []*ves_io_schema_known_label_key.LabelKeyType{
			{
				Key: req.Key,
			},
		},
	}, nil
}

func (s *knownLabelKeyCustomAPIServer) Delete(ctx context.Context, req *ves_io_schema_known_label_key.DeleteRequest) (*ves_io_schema_known_label_key.DeleteResponse, error) {
	return &ves_io_schema_known_label_key.DeleteResponse{}, nil
}

var _ ves_io_schema_known_label_key.CustomAPIServer = &knownLabelKeyCustomAPIServer{}
