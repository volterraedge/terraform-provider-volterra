// Copyright (c) 2023 F5 Inc. All rights reserved.
package volterra

import (
	"context"
	"fmt"
	"os"
	"testing"

	"github.com/gogo/protobuf/types"
	"github.com/stretchr/testify/assert"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/resource"
	ves_io_schema "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema"
	ves_io_schema_combined "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/combined"
	schema_registration "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/registration"
	schema_token "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/token"
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/notify"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
	svcfw_test "gopkg.volterra.us/stdlib/svcfw/test"
	"gopkg.volterra.us/stdlib/svcfw/test/generic"
	//ves_io_schema_ns "github.com/volterraedge/terraform-provider-volterra/pbgo/extschema/schema/namespace"
)

var rName = fmt.Sprintf("r-%s", rUID)
var rUID = "0376d64d-687c-4735-b2ab-36af79b49b6a"

var newReg = &schema_registration.Object{
	Metadata:       &ves_io_schema.ObjectMetaType{Name: rName, Namespace: svcfw.SystemNSVal, Uid: rUID},
	SystemMetadata: &ves_io_schema.SystemObjectMetaType{Uid: rUID, Tenant: "customer1", CreatorClass: "random"},
	Spec: &schema_registration.SpecType{
		GcSpec: &schema_registration.GlobalSpecType{
			Token: "token-uid",
			Passport: &schema_registration.Passport{
				ClusterName: "site-1",
				ClusterType: "ce",
				ClusterSize: 1,
			},
			Infra: &schema_registration.Infra{
				Hostname: "master-0",
			},
		},
	},
	Status: &schema_registration.StatusType{CurrentState: schema_registration.PENDING},
}

func getTLSOpts() []svcfw_test.ConfigOpt {
	clKey := "client.examplesvc.ves.io.key"
	clCert := "client.examplesvc.ves.io.crt"
	clCACert := "client.ca.crt"
	svrKey := "server.examplesvc.ves.io.key"
	svrCert := "server.examplesvc.ves.io.crt"
	svrCACert := "server.ca.crt"
	return []svcfw_test.ConfigOpt{
		svcfw_test.WithCfgTDRoot("../testdata/"),
		svcfw_test.WithCfgClientKey(clKey),
		svcfw_test.WithCfgClientCert(clCert),
		svcfw_test.WithCfgClientCACert(clCACert),
		svcfw_test.WithCfgServerKey(svrKey),
		svcfw_test.WithCfgServerCert(svrCert),
		svcfw_test.WithCfgServerCACert(svrCACert),
	}
}

var (
	etcdURL  string
	etcdStop func()
)

func TestMain(m *testing.M) {
	// The proxy config should be set in environment for all tests since otherwise golang http will cache read-in
	// env var value and does not care for changing over tests
	_, etcdURL, etcdStop = svcfw_test.NewEmbedEtcd(&testing.T{})

	code := m.Run()

	// os.Exit() does not respect defer statements
	etcdStop()
	os.Exit(code)
}

func TestRegistrationApproval(t *testing.T) {

	//set noTLS
	unSetenv := svcfw_test.Setenv(server.EnvVarServiceOnNonTLS, "true")
	defer unSetenv()

	// start maurice server
	fcOptsMaurice := GetMauriceFixtureOpts(etcdURL)
	srvF := generic.NewFixture(t, ves_io_schema_combined.MDR, fcOptsMaurice...)
	srvStop := srvF.MustStart(srvF.MustNewConfig(t.Name()))
	defer srvStop()

	srvF.MustAddEntry(newReg.ToEntry())

	name := rName
	resourceName := approvalResource + "." + name
	os.Setenv("VOLT_API_TEST", "true")
	os.Setenv("VOLT_API_URL", fmt.Sprintf("https://localhost:%d", srvF.Svc.RestServerTLSPort()))
	os.Setenv("TF_ACC", "true")
	os.Setenv("VOLT_VESENV", "true")
	os.Setenv("VOLT_TENANT", "customer1")
	resource.Test(t, resource.TestCase{
		Providers: testAccProviders,
		PreCheck:  func() { testAccPreCheck() },
		Steps: []resource.TestStep{
			{
				Config: testRegApproval(name),
				Check: resource.ComposeTestCheckFunc(
					resource.TestCheckResourceAttr(resourceName, "registration_name", name),
					resource.TestCheckResourceAttr(resourceName, "cluster_name", "site-1"),
				),
			},
		},
	})

	e := srvF.MustFetchEntry(newReg.Metadata.Uid, schema_registration.ObjectType)
	obj := e.(*schema_registration.DBObject)

	assert.Equal(t, schema_registration.RETIRED, obj.Status.CurrentState)
}

func testRegApproval(name string) string {
	return fmt.Sprintf(`
		resource "%s" "%s" {
		  cluster_name = "site-1"
		  cluster_size = 1
		  latitude = 50
		  longitude = 50
		  hostname = "master-0"
		  tunnel_type = "SITE_TO_SITE_TUNNEL_IPSEC"
		}
		`, approvalResource, name)
}

type MauriceCustomAPI struct {
	sf svcfw.Service
}

var _ schema_registration.CustomAPIServer = &MauriceCustomAPI{}

// volterra.MauriceCustomAPI does not satisfy interface ves_io_schema_registration.CustomAPIServer
func GetMauriceFixtureOpts(etcdURL string) []generic.ConfigOpt {

	// custom modules
	customModFns := svcfw_test.ModuleFuncs{
		NameFn: func(sf svcfw.Service) string {
			return "customAPIHandler"
		},

		GetCRUDAPIsFn: func(sf svcfw.Service) map[string][]*server.CRUDServerCfg {
			return map[string][]*server.CRUDServerCfg{
				schema_registration.ObjectType: {
					server.NewCRUDServerCfg(
						db.DefaultTableName(schema_registration.ObjectType),
						server.CRUDWithPublic,
					),
					server.NewCRUDServerCfg(
						db.DefaultTableName(schema_registration.ObjectType),
					),
				},
				schema_token.ObjectType: {
					server.NewCRUDServerCfg(
						db.DefaultTableName(schema_token.ObjectType),
						server.CRUDWithPublic,
					),
				},
			}
		},

		GetCustomAPIHandlersFn: func(sf svcfw.Service) map[string]server.APIHandler {
			return map[string]server.APIHandler{
				"ves.io.schema.registration.CustomAPI": &MauriceCustomAPI{
					sf,
				},
			}
		},
		// GetNotifHandlers
		GetNotifHandlersFn: func(sf svcfw.Service) map[string]notify.Handler {
			return map[string]notify.Handler{}
		},
	}
	customMod := svcfw_test.NewModuleFactory(customModFns)

	oTypes := []string{}
	for ot := range ves_io_schema_combined.MDR.EntryStoreMap {
		oTypes = append(oTypes, ot)
	}
	fcOpts := []generic.ConfigOpt{}
	fcOpts = append(fcOpts, generic.WithCfgSTFConfigOpts(getTLSOpts()...))
	fcOpts = append(fcOpts, generic.WithCfgSTFConfigOpts(
		svcfw_test.WithCfgEtcdURLs(etcdURL),
		svcfw_test.WithCfgServedTypes(store.Etcd, oTypes...),
		svcfw_test.WithCfgNewModuleFn(customMod),
		svcfw_test.WithCfgPublicAPITypes(schema_registration.ObjectType, schema_token.ObjectType),
		svcfw_test.WithCfgKindToTypeFn(func(tp, uid, kind string) (string, error) { return "ves.io.schema." + kind, nil }),
		//svcfw_test.WithCfgEtcdKeyPfx("/akar"),
	))
	return fcOpts
}

func (s *MauriceCustomAPI) Get(ctx context.Context, r *schema_registration.GetRequest) (*schema_registration.GetResponse, error) {
	return &schema_registration.GetResponse{}, nil
}

func (s *MauriceCustomAPI) List(ctx context.Context, r *schema_registration.ListRequest) (*schema_registration.ListResponse, error) {
	return nil, nil
}

func (s *MauriceCustomAPI) Delete(ctx context.Context, r *schema_registration.DeleteRequest) (*types.Empty, error) {
	return nil, nil
}

func (s *MauriceCustomAPI) RegistrationApprove(ctx context.Context, r *schema_registration.ApprovalReq) (*schema_registration.ObjectChangeResp, error) {
	tenant := server.TenantFromContext(ctx)
	if tenant == "" {
		return nil, fmt.Errorf("Tenant missing in request")
	}
	if r.State == schema_registration.PENDING || r.State == schema_registration.ADMITTED {
		return nil, fmt.Errorf("This state (%s) can't be set", r.State)
	}

	e, ok, err := s.sf.FindEntry(ctx, schema_registration.ObjectDefTblName, r.Name[2:])
	if err != nil {
		return nil, fmt.Errorf("cannot find registration entry %s, err: %s", r.Name[2:], err)
	}
	if !ok {
		return nil, fmt.Errorf("Test Registration %s does not exist", r.Name)
	}
	obj := e.(*schema_registration.DBObject)

	// return APPROVED in case of approval
	if r.State == schema_registration.APPROVED {
		if r.Passport == nil {
			return nil, fmt.Errorf("Passport is required")
		}
		if r.Passport.ClusterName == "" {
			return nil, fmt.Errorf("ClusterName is missing")
		}
		obj.Spec.GcSpec.Passport.ClusterName = r.Passport.ClusterName
		obj.Status.CurrentState = schema_registration.APPROVED
		if _, err := s.sf.ChgEntry(ctx, obj.ToEntry()); err != nil {
			return nil, fmt.Errorf("Could not change registration entry %s", err)
		}

		return &schema_registration.ObjectChangeResp{Obj: obj.Object}, nil
	}

	if r.State == schema_registration.RETIRED {
		obj.Status.CurrentState = schema_registration.RETIRED
		if _, err := s.sf.ChgEntry(ctx, obj.ToEntry()); err != nil {
			return nil, fmt.Errorf("Could not change registration entry %s", err)
		}
		return &schema_registration.ObjectChangeResp{Obj: obj.Object}, nil
	}

	return nil, fmt.Errorf("RegistrationApprove API supports only APPROVE and RETIRE for now")
}

func (s *MauriceCustomAPI) RegistrationCreate(ctx context.Context, r *schema_registration.RegistrationCreateRequest) (*schema_registration.Object, error) {
	return nil, nil
}

func (s *MauriceCustomAPI) ListRegistrationsByState(ctx context.Context, r *schema_registration.ListStateReq) (*schema_registration.ListResponse, error) {
	if r.Namespace != svcfw.SystemNSVal {
		return nil, fmt.Errorf("Validation namespace is not system, but : %s", r.Namespace)
	}
	tenant := server.TenantFromContext(ctx)
	if tenant == "" {
		return nil, fmt.Errorf("Tenant missing in request")
	}
	if r.State != schema_registration.PENDING {
		return nil, fmt.Errorf("ListRegistrationsByState API supports only registration state PENDING for now")
	}
	e, ok, err := s.sf.FindEntry(ctx, schema_registration.ObjectDefTblName, newReg.SystemMetadata.Uid)
	if err != nil {
		return nil, fmt.Errorf("err in listing registrations %s", err)
	}
	if !ok {
		return nil, fmt.Errorf("Reg %s does not exists", newReg.SystemMetadata.Uid)
	}
	obj := e.(*schema_registration.DBObject)

	ri := &schema_registration.ListResponseItem{
		Tenant:         tenant,
		Namespace:      svcfw.SystemNSVal,
		Name:           obj.Metadata.Name,
		Uid:            rUID,
		Description:    obj.GetMetadata().GetDescription(),
		Disabled:       obj.GetMetadata().GetDisable(),
		Labels:         obj.GetMetadata().GetLabels(),
		Annotations:    obj.GetMetadata().GetAnnotations(),
		Metadata:       &ves_io_schema.ObjectGetMetaType{},
		SystemMetadata: &ves_io_schema.SystemObjectGetMetaType{},
		GetSpec:        &schema_registration.GetSpecType{},
	}
	ri.GetSpec.FromGlobalSpecTypeWithoutDeepCopy(obj.GetSpec().GetGcSpec())
	ri.Metadata.FromObjectMetaTypeWithoutDeepCopy(obj.GetMetadata())
	ri.SystemMetadata.FromSystemObjectMetaTypeWithoutDeepCopy(obj.GetSystemMetadata())
	// return single list item with PENDING state
	return &schema_registration.ListResponse{Items: []*schema_registration.ListResponseItem{ri}}, nil
}

func (s *MauriceCustomAPI) RegistrationConfig(ctx context.Context, r *schema_registration.ConfigReq) (*schema_registration.ConfigResp, error) {
	return nil, nil
}

func (s *MauriceCustomAPI) ListRegistrationsBySite(ctx context.Context, r *schema_registration.ListBySiteReq) (*schema_registration.ListResponse, error) {
	if r.Namespace != svcfw.SystemNSVal {
		return nil, fmt.Errorf("Validation namespace is not system, but : %s", r.Namespace)
	}
	tenant := server.TenantFromContext(ctx)
	if tenant == "" {
		return nil, fmt.Errorf("Tenant missing in request")
	}
	e, ok, err := s.sf.FindEntry(ctx, schema_registration.ObjectDefTblName, newReg.SystemMetadata.Uid)
	if err != nil {
		return nil, fmt.Errorf("err in listing registrations %s", err)
	}
	if !ok {
		return nil, fmt.Errorf("Reg %s does not exists", newReg.SystemMetadata.Uid)
	}
	obj := e.(*schema_registration.DBObject)
	ri := &schema_registration.ListResponseItem{
		Tenant:         tenant,
		Namespace:      svcfw.SystemNSVal,
		Name:           obj.Metadata.Name,
		Uid:            rUID,
		Description:    obj.GetMetadata().GetDescription(),
		Disabled:       obj.GetMetadata().GetDisable(),
		Labels:         obj.GetMetadata().GetLabels(),
		Annotations:    obj.GetMetadata().GetAnnotations(),
		Metadata:       &ves_io_schema.ObjectGetMetaType{},
		SystemMetadata: &ves_io_schema.SystemObjectGetMetaType{},
		GetSpec:        &schema_registration.GetSpecType{},
	}
	ri.GetSpec.FromGlobalSpecTypeWithoutDeepCopy(obj.GetSpec().GetGcSpec())
	ri.Metadata.FromObjectMetaTypeWithoutDeepCopy(obj.GetMetadata())
	ri.SystemMetadata.FromSystemObjectMetaTypeWithoutDeepCopy(obj.GetSystemMetadata())

	// return single list item with PENDING state
	return &schema_registration.ListResponse{Items: []*schema_registration.ListResponseItem{ri}}, nil
}
