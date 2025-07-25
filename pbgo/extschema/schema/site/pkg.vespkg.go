// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package site

import (
	"reflect"

	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/store"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.site.SpecType"] = SpecTypeValidator()

	vr["ves.io.schema.site.Object"] = ObjectValidator()
	vr["ves.io.schema.site.StatusObject"] = StatusObjectValidator()

	vr["ves.io.schema.site.GetRequest"] = GetRequestValidator()
	vr["ves.io.schema.site.GetResponse"] = GetResponseValidator()
	vr["ves.io.schema.site.ListRequest"] = ListRequestValidator()
	vr["ves.io.schema.site.ListResponse"] = ListResponseValidator()
	vr["ves.io.schema.site.ListResponseItem"] = ListResponseItemValidator()
	vr["ves.io.schema.site.ReplaceRequest"] = ReplaceRequestValidator()
	vr["ves.io.schema.site.ReplaceResponse"] = ReplaceResponseValidator()

	vr["ves.io.schema.site.ConfigMapListRequest"] = ConfigMapListRequestValidator()
	vr["ves.io.schema.site.CronJobListRequest"] = CronJobListRequestValidator()
	vr["ves.io.schema.site.DaemonSetListRequest"] = DaemonSetListRequestValidator()
	vr["ves.io.schema.site.DeploymentListRequest"] = DeploymentListRequestValidator()
	vr["ves.io.schema.site.EndpointsListRequest"] = EndpointsListRequestValidator()
	vr["ves.io.schema.site.JobListRequest"] = JobListRequestValidator()
	vr["ves.io.schema.site.NamespaceListRequest"] = NamespaceListRequestValidator()
	vr["ves.io.schema.site.NodeListRequest"] = NodeListRequestValidator()
	vr["ves.io.schema.site.PersistentVolumeClaimListRequest"] = PersistentVolumeClaimListRequestValidator()
	vr["ves.io.schema.site.PersistentVolumeListRequest"] = PersistentVolumeListRequestValidator()
	vr["ves.io.schema.site.PodListRequest"] = PodListRequestValidator()
	vr["ves.io.schema.site.PodsMetricData"] = PodsMetricDataValidator()
	vr["ves.io.schema.site.PodsMetricTypeData"] = PodsMetricTypeDataValidator()
	vr["ves.io.schema.site.PodsMetricsRequest"] = PodsMetricsRequestValidator()
	vr["ves.io.schema.site.PodsMetricsResponse"] = PodsMetricsResponseValidator()
	vr["ves.io.schema.site.ReplicaSetListRequest"] = ReplicaSetListRequestValidator()
	vr["ves.io.schema.site.SecretListRequest"] = SecretListRequestValidator()
	vr["ves.io.schema.site.ServiceListRequest"] = ServiceListRequestValidator()
	vr["ves.io.schema.site.StatefulSetListRequest"] = StatefulSetListRequestValidator()
	vr["ves.io.schema.site.VirtualMachineInstancesMetricData"] = VirtualMachineInstancesMetricDataValidator()
	vr["ves.io.schema.site.VirtualMachineInstancesMetricTypeData"] = VirtualMachineInstancesMetricTypeDataValidator()
	vr["ves.io.schema.site.VirtualMachineInstancesMetricsRequest"] = VirtualMachineInstancesMetricsRequestValidator()
	vr["ves.io.schema.site.VirtualMachineInstancesMetricsResponse"] = VirtualMachineInstancesMetricsResponseValidator()

	vr["ves.io.schema.site.GlobalNetworkItem"] = GlobalNetworkItemValidator()
	vr["ves.io.schema.site.GlobalNetworkListRequest"] = GlobalNetworkListRequestValidator()
	vr["ves.io.schema.site.GlobalNetworkListResp"] = GlobalNetworkListRespValidator()
	vr["ves.io.schema.site.SegmentItem"] = SegmentItemValidator()
	vr["ves.io.schema.site.SegmentListRequest"] = SegmentListRequestValidator()
	vr["ves.io.schema.site.SegmentListResp"] = SegmentListRespValidator()

	vr["ves.io.schema.site.SetStateReq"] = SetStateReqValidator()
	vr["ves.io.schema.site.SetStateResp"] = SetStateRespValidator()

	vr["ves.io.schema.site.UpgradeOSRequest"] = UpgradeOSRequestValidator()
	vr["ves.io.schema.site.UpgradeOSResponse"] = UpgradeOSResponseValidator()
	vr["ves.io.schema.site.UpgradeSWRequest"] = UpgradeSWRequestValidator()
	vr["ves.io.schema.site.UpgradeSWResponse"] = UpgradeSWResponseValidator()

	vr["ves.io.schema.site.CheckSiteExistRequest"] = CheckSiteExistRequestValidator()
	vr["ves.io.schema.site.CheckSiteExistResponse"] = CheckSiteExistResponseValidator()
	vr["ves.io.schema.site.SiteStatusMetricsRequest"] = SiteStatusMetricsRequestValidator()
	vr["ves.io.schema.site.SiteStatusMetricsResponse"] = SiteStatusMetricsResponseValidator()

	vr["ves.io.schema.site.AresConnectionStatus"] = AresConnectionStatusValidator()
	vr["ves.io.schema.site.AzureExpressRouteCircuitStatusType"] = AzureExpressRouteCircuitStatusTypeValidator()
	vr["ves.io.schema.site.AzureExpressRouteStatusType"] = AzureExpressRouteStatusTypeValidator()
	vr["ves.io.schema.site.AzureHubSpokeVnetPeeringStatusInfo"] = AzureHubSpokeVnetPeeringStatusInfoValidator()
	vr["ves.io.schema.site.AzureVNetToVnetPeeringStatus"] = AzureVNetToVnetPeeringStatusValidator()
	vr["ves.io.schema.site.AzureVnetPeeringStateType"] = AzureVnetPeeringStateTypeValidator()
	vr["ves.io.schema.site.Bios"] = BiosValidator()
	vr["ves.io.schema.site.Board"] = BoardValidator()
	vr["ves.io.schema.site.BondMembersType"] = BondMembersTypeValidator()
	vr["ves.io.schema.site.Chassis"] = ChassisValidator()
	vr["ves.io.schema.site.Coordinates"] = CoordinatesValidator()
	vr["ves.io.schema.site.Cpu"] = CpuValidator()
	vr["ves.io.schema.site.CreateGlobalKubeConfigReq"] = CreateGlobalKubeConfigReqValidator()
	vr["ves.io.schema.site.CreateKubeConfigReq"] = CreateKubeConfigReqValidator()
	vr["ves.io.schema.site.CreateSpecType"] = CreateSpecTypeValidator()
	vr["ves.io.schema.site.DefaultUnderlayNetworkType"] = DefaultUnderlayNetworkTypeValidator()
	vr["ves.io.schema.site.DeploymentState"] = DeploymentStateValidator()
	vr["ves.io.schema.site.DirectConnectBGPPeerInfo"] = DirectConnectBGPPeerInfoValidator()
	vr["ves.io.schema.site.DirectConnectStatusInfo"] = DirectConnectStatusInfoValidator()
	vr["ves.io.schema.site.DirectConnectVIFStateInfo"] = DirectConnectVIFStateInfoValidator()
	vr["ves.io.schema.site.ExpressRoutePeeringStatusType"] = ExpressRoutePeeringStatusTypeValidator()
	vr["ves.io.schema.site.FleetCondition"] = FleetConditionValidator()
	vr["ves.io.schema.site.FleetDeploymentState"] = FleetDeploymentStateValidator()
	vr["ves.io.schema.site.FleetStatus"] = FleetStatusValidator()
	vr["ves.io.schema.site.GPU"] = GPUValidator()
	vr["ves.io.schema.site.GPUDevice"] = GPUDeviceValidator()
	vr["ves.io.schema.site.GetSpecType"] = GetSpecTypeValidator()
	vr["ves.io.schema.site.GlobalAccessCheckRequest"] = GlobalAccessCheckRequestValidator()
	vr["ves.io.schema.site.GlobalAccessCheckResponse"] = GlobalAccessCheckResponseValidator()
	vr["ves.io.schema.site.GlobalSpecType"] = GlobalSpecTypeValidator()
	vr["ves.io.schema.site.InterfaceStatus"] = InterfaceStatusValidator()
	vr["ves.io.schema.site.K8SApiServerParameters"] = K8SApiServerParametersValidator()
	vr["ves.io.schema.site.Kernel"] = KernelValidator()
	vr["ves.io.schema.site.KubeConfigStatusRsp"] = KubeConfigStatusRspValidator()
	vr["ves.io.schema.site.ListGlobalKubeConfigReq"] = ListGlobalKubeConfigReqValidator()
	vr["ves.io.schema.site.ListKubeConfigReq"] = ListKubeConfigReqValidator()
	vr["ves.io.schema.site.ListKubeConfigRsp"] = ListKubeConfigRspValidator()
	vr["ves.io.schema.site.ListKubeConfigRspItem"] = ListKubeConfigRspItemValidator()
	vr["ves.io.schema.site.Memory"] = MemoryValidator()
	vr["ves.io.schema.site.NetworkDevice"] = NetworkDeviceValidator()
	vr["ves.io.schema.site.Node"] = NodeValidator()
	vr["ves.io.schema.site.NodeInfo"] = NodeInfoValidator()
	vr["ves.io.schema.site.OS"] = OSValidator()
	vr["ves.io.schema.site.OfflineSurvivabilityStatus"] = OfflineSurvivabilityStatusValidator()
	vr["ves.io.schema.site.OperatingSystemStatus"] = OperatingSystemStatusValidator()
	vr["ves.io.schema.site.OsInfo"] = OsInfoValidator()
	vr["ves.io.schema.site.Product"] = ProductValidator()
	vr["ves.io.schema.site.PublishVIPParamsPerAz"] = PublishVIPParamsPerAzValidator()
	vr["ves.io.schema.site.ReMeshGroup"] = ReMeshGroupValidator()
	vr["ves.io.schema.site.ReplaceSpecType"] = ReplaceSpecTypeValidator()
	vr["ves.io.schema.site.RevokeKubeConfigReq"] = RevokeKubeConfigReqValidator()
	vr["ves.io.schema.site.RouteServerPeeringStatusType"] = RouteServerPeeringStatusTypeValidator()
	vr["ves.io.schema.site.RouteServerStatusType"] = RouteServerStatusTypeValidator()
	vr["ves.io.schema.site.ScalingStatus"] = ScalingStatusValidator()
	vr["ves.io.schema.site.SiteReachabilityStatus"] = SiteReachabilityStatusValidator()
	vr["ves.io.schema.site.SiteStatusMetricsData"] = SiteStatusMetricsDataValidator()
	vr["ves.io.schema.site.SiteStatusMetricsFieldData"] = SiteStatusMetricsFieldDataValidator()
	vr["ves.io.schema.site.SiteToSiteTunnelConnectivity"] = SiteToSiteTunnelConnectivityValidator()
	vr["ves.io.schema.site.StorageDevice"] = StorageDeviceValidator()
	vr["ves.io.schema.site.TunnelConnectionStatus"] = TunnelConnectionStatusValidator()
	vr["ves.io.schema.site.TunnelFlapReason"] = TunnelFlapReasonValidator()
	vr["ves.io.schema.site.USBDevice"] = USBDeviceValidator()
	vr["ves.io.schema.site.VerInstanceRunningStateStatusType"] = VerInstanceRunningStateStatusTypeValidator()
	vr["ves.io.schema.site.VerMasterStatusType"] = VerMasterStatusTypeValidator()
	vr["ves.io.schema.site.VerStatusType"] = VerStatusTypeValidator()
	vr["ves.io.schema.site.VnetGatewayStatusType"] = VnetGatewayStatusTypeValidator()
	vr["ves.io.schema.site.VolterraSoftwareStatus"] = VolterraSoftwareStatusValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {
	mdr.EntryFactory["ves.io.schema.site.Object"] = NewEntryObject
	mdr.EntryStoreMap["ves.io.schema.site.Object"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.site.Object"] = reflect.TypeOf(&DBObject{})
	mdr.EntryIndexers["ves.io.schema.site.Object"] = GetObjectIndexers
	mdr.EntryFactory["ves.io.schema.site.StatusObject"] = NewEntryStatusObject
	mdr.EntryStoreMap["ves.io.schema.site.StatusObject"] = store.InMemory
	mdr.EntryRegistry["ves.io.schema.site.StatusObject"] = reflect.TypeOf(&DBStatusObject{})
	mdr.EntryIndexers["ves.io.schema.site.StatusObject"] = GetStatusObjectIndexers

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.site.API.Get"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "replace_form.spec.bgp_peer_address_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.inside_vip_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.outside_vip_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "replace_form.spec.vip_selection",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.bgp_peer_address_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.inside_vip_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.outside_vip_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.vip_selection",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCAvailableInResFieldRegistry["ves.io.schema.site.API.List"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "items.#.get_spec.bgp_peer_address_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.inside_vip_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.outside_vip_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "items.#.get_spec.vip_selection",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

	mdr.RPCSubscriptionFieldsRegistry["ves.io.schema.site.API.Replace"] = []svcfw.SubscriptionField{
		{
			FieldPath:     "ves.io.schema.site.ReplaceRequest.spec.bgp_peer_address_v6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.site.ReplaceRequest.spec.inside_nameserver_v6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.site.ReplaceRequest.spec.inside_vip_v6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.site.ReplaceRequest.spec.outside_nameserver_v6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
		{
			FieldPath:     "ves.io.schema.site.ReplaceRequest.spec.outside_vip_v6",
			AddonServices: []string{"f5xc-ipv6-standard"},
		},
	}

	mdr.RPCAvailableInReqFieldRegistry["ves.io.schema.site.API.Replace"] = []svcfw.EnvironmentField{
		{
			FieldPath:           "spec.bgp_peer_address_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.inside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.inside_vip_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.outside_nameserver_v6",
			AllowedEnvironments: []string{"crt", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.outside_vip_v6",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
		{
			FieldPath:           "spec.vip_selection",
			AllowedEnvironments: []string{"crt", "demo1", "prod", "softbank_mec", "staging", "test"},
		},
	}

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.site.ConfigKubeConfigAPI"] = "config"
	sm["ves.io.schema.site.API"] = "config"
	sm["ves.io.schema.site.CustomDataK8SAPI"] = "data"
	sm["ves.io.schema.site.CustomVirtualNetworkListAPI"] = "config"
	sm["ves.io.schema.site.CustomStateAPI"] = "register"
	sm["ves.io.schema.site.UamKubeConfigAPI"] = "web"
	sm["ves.io.schema.site.UpgradeAPI"] = "config"
	sm["ves.io.schema.site.CustomSiteStatusAPI"] = "data"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

	sm["config"] = svcfw.P0PolicyInfo{
		Name:            "ves-io-allow-config",
		ServiceSelector: "akar\\.gc.*\\",
	}

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

	customCSR = mdr.PvtCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.site.Object"] = PrivateConfigKubeConfigAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.site.PrivateConfigKubeConfigAPI"] = NewPrivateConfigKubeConfigAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.site.PrivateConfigKubeConfigAPI"] = NewPrivateConfigKubeConfigAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.site.PrivateConfigKubeConfigAPI"] = RegisterPrivateConfigKubeConfigAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.PrivateConfigKubeConfigAPI"] = RegisterGwPrivateConfigKubeConfigAPIHandler
		customCSR.ServerRegistry["ves.io.schema.site.PrivateConfigKubeConfigAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewPrivateConfigKubeConfigAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.site.Object"] = ConfigKubeConfigAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.site.ConfigKubeConfigAPI"] = NewConfigKubeConfigAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.site.ConfigKubeConfigAPI"] = NewConfigKubeConfigAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.site.ConfigKubeConfigAPI"] = RegisterConfigKubeConfigAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.ConfigKubeConfigAPI"] = RegisterGwConfigKubeConfigAPIHandler
		customCSR.ServerRegistry["ves.io.schema.site.ConfigKubeConfigAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewConfigKubeConfigAPIServer(svc)
		}

	}()

	csr = mdr.PubCRUDServiceRegistry

	func() {
		// set swagger jsons for our and external schemas
		csr.CRUDSwaggerRegistry["ves.io.schema.site.Object"] = APISwaggerJSON
		csr.CRUDGrpcClientRegistry["ves.io.schema.site.Object"] = NewCRUDAPIGrpcClient
		csr.CRUDRestClientRegistry["ves.io.schema.site.Object"] = NewCRUDAPIRestClient
		csr.CRUDInprocClientRegistry["ves.io.schema.site.Object"] = NewCRUDAPIInprocClient
		if isExternal {
			return
		}
		// registration of api handlers if our own schema
		mdr.SvcRegisterHandlers["ves.io.schema.site.API"] = RegisterAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.API"] = RegisterGwAPIHandler
		csr.CRUDServerRegistry["ves.io.schema.site.Object"] = NewCRUDAPIServer

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.site.Object"] = CustomDataK8SAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.site.CustomDataK8SAPI"] = NewCustomDataK8SAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.site.CustomDataK8SAPI"] = NewCustomDataK8SAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.site.CustomDataK8SAPI"] = RegisterCustomDataK8SAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.CustomDataK8SAPI"] = RegisterGwCustomDataK8SAPIHandler
		customCSR.ServerRegistry["ves.io.schema.site.CustomDataK8SAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomDataK8SAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.site.Object"] = CustomVirtualNetworkListAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.site.CustomVirtualNetworkListAPI"] = NewCustomVirtualNetworkListAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.site.CustomVirtualNetworkListAPI"] = NewCustomVirtualNetworkListAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.site.CustomVirtualNetworkListAPI"] = RegisterCustomVirtualNetworkListAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.CustomVirtualNetworkListAPI"] = RegisterGwCustomVirtualNetworkListAPIHandler
		customCSR.ServerRegistry["ves.io.schema.site.CustomVirtualNetworkListAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomVirtualNetworkListAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.site.Object"] = CustomStateAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.site.CustomStateAPI"] = NewCustomStateAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.site.CustomStateAPI"] = NewCustomStateAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.site.CustomStateAPI"] = RegisterCustomStateAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.CustomStateAPI"] = RegisterGwCustomStateAPIHandler
		customCSR.ServerRegistry["ves.io.schema.site.CustomStateAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomStateAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.site.Object"] = UamKubeConfigAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.site.UamKubeConfigAPI"] = NewUamKubeConfigAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.site.UamKubeConfigAPI"] = NewUamKubeConfigAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.site.UamKubeConfigAPI"] = RegisterUamKubeConfigAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.UamKubeConfigAPI"] = RegisterGwUamKubeConfigAPIHandler
		customCSR.ServerRegistry["ves.io.schema.site.UamKubeConfigAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewUamKubeConfigAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.site.Object"] = UpgradeAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.site.UpgradeAPI"] = NewUpgradeAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.site.UpgradeAPI"] = NewUpgradeAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.site.UpgradeAPI"] = RegisterUpgradeAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.UpgradeAPI"] = RegisterGwUpgradeAPIHandler
		customCSR.ServerRegistry["ves.io.schema.site.UpgradeAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewUpgradeAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.site.Object"] = CustomSiteStatusAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.site.CustomSiteStatusAPI"] = NewCustomSiteStatusAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.site.CustomSiteStatusAPI"] = NewCustomSiteStatusAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.site.CustomSiteStatusAPI"] = RegisterCustomSiteStatusAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.site.CustomSiteStatusAPI"] = RegisterGwCustomSiteStatusAPIHandler
		customCSR.ServerRegistry["ves.io.schema.site.CustomSiteStatusAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomSiteStatusAPIServer(svc)
		}

	}()

}

func InitializeMDRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	initializeEntryRegistry(mdr)
	initializeValidatorRegistry(mdr.ValidatorRegistry)

	initializeCRUDServiceRegistry(mdr, isExternal)
	if isExternal {
		return
	}

	initializeRPCRegistry(mdr)
	initializeAPIGwServiceSlugsRegistry(mdr.APIGwServiceSlugs)
	initializeP0PolicyRegistry(mdr.P0PolicyRegistry)

}
