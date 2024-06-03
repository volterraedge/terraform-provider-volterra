// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package topology

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/server"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.topology.ListCloudNetworkTagKeysRequest"] = ListCloudNetworkTagKeysRequestValidator()
	vr["ves.io.schema.topology.ListCloudNetworkTagKeysResponse"] = ListCloudNetworkTagKeysResponseValidator()
	vr["ves.io.schema.topology.ListCloudNetworkTagValuesRequest"] = ListCloudNetworkTagValuesRequestValidator()
	vr["ves.io.schema.topology.ListCloudNetworkTagValuesResponse"] = ListCloudNetworkTagValuesResponseValidator()
	vr["ves.io.schema.topology.ListCloudSubnetTagKeysRequest"] = ListCloudSubnetTagKeysRequestValidator()
	vr["ves.io.schema.topology.ListCloudSubnetTagKeysResponse"] = ListCloudSubnetTagKeysResponseValidator()
	vr["ves.io.schema.topology.ListCloudSubnetTagValuesRequest"] = ListCloudSubnetTagValuesRequestValidator()
	vr["ves.io.schema.topology.ListCloudSubnetTagValuesResponse"] = ListCloudSubnetTagValuesResponseValidator()

	vr["ves.io.schema.topology.DCClusterGroupSummaryInfo"] = DCClusterGroupSummaryInfoValidator()
	vr["ves.io.schema.topology.DCClusterGroupsSummaryRequest"] = DCClusterGroupsSummaryRequestValidator()
	vr["ves.io.schema.topology.DCClusterTopologyRequest"] = DCClusterTopologyRequestValidator()
	vr["ves.io.schema.topology.Edge"] = EdgeValidator()
	vr["ves.io.schema.topology.EdgeInfoSummary"] = EdgeInfoSummaryValidator()
	vr["ves.io.schema.topology.LinkInfo"] = LinkInfoValidator()
	vr["ves.io.schema.topology.LinkInfoSummary"] = LinkInfoSummaryValidator()
	vr["ves.io.schema.topology.LinkTypeData"] = LinkTypeDataValidator()
	vr["ves.io.schema.topology.MetricData"] = MetricDataValidator()
	vr["ves.io.schema.topology.MetricSelector"] = MetricSelectorValidator()
	vr["ves.io.schema.topology.MetricTypeData"] = MetricTypeDataValidator()
	vr["ves.io.schema.topology.NetworkRouteTableData"] = NetworkRouteTableDataValidator()
	vr["ves.io.schema.topology.NetworkRouteTableMetaData"] = NetworkRouteTableMetaDataValidator()
	vr["ves.io.schema.topology.NetworkRouteTablesRequest"] = NetworkRouteTablesRequestValidator()
	vr["ves.io.schema.topology.NetworkRouteTablesResponse"] = NetworkRouteTablesResponseValidator()
	vr["ves.io.schema.topology.NetworkRoutesData"] = NetworkRoutesDataValidator()
	vr["ves.io.schema.topology.NetworkRoutesMetaData"] = NetworkRoutesMetaDataValidator()
	vr["ves.io.schema.topology.NetworkSummaryInfo"] = NetworkSummaryInfoValidator()
	vr["ves.io.schema.topology.Node"] = NodeValidator()
	vr["ves.io.schema.topology.NodeMetaData"] = NodeMetaDataValidator()
	vr["ves.io.schema.topology.NodeTypeDCClusterGroup"] = NodeTypeDCClusterGroupValidator()
	vr["ves.io.schema.topology.NodeTypeInstance"] = NodeTypeInstanceValidator()
	vr["ves.io.schema.topology.NodeTypeNetwork"] = NodeTypeNetworkValidator()
	vr["ves.io.schema.topology.NodeTypeSite"] = NodeTypeSiteValidator()
	vr["ves.io.schema.topology.NodeTypeSiteMeshGroup"] = NodeTypeSiteMeshGroupValidator()
	vr["ves.io.schema.topology.NodeTypeSubnet"] = NodeTypeSubnetValidator()
	vr["ves.io.schema.topology.NodeTypeTransitGateway"] = NodeTypeTransitGatewayValidator()
	vr["ves.io.schema.topology.RouteTableData"] = RouteTableDataValidator()
	vr["ves.io.schema.topology.RouteTableMetaData"] = RouteTableMetaDataValidator()
	vr["ves.io.schema.topology.RouteTableRequest"] = RouteTableRequestValidator()
	vr["ves.io.schema.topology.RouteTableResponse"] = RouteTableResponseValidator()
	vr["ves.io.schema.topology.SiteMeshGroupSummaryInfo"] = SiteMeshGroupSummaryInfoValidator()
	vr["ves.io.schema.topology.SiteMeshGroupsSummaryRequest"] = SiteMeshGroupsSummaryRequestValidator()
	vr["ves.io.schema.topology.SiteMeshTopologyRequest"] = SiteMeshTopologyRequestValidator()
	vr["ves.io.schema.topology.SiteNetworksRequest"] = SiteNetworksRequestValidator()
	vr["ves.io.schema.topology.SiteNetworksResponse"] = SiteNetworksResponseValidator()
	vr["ves.io.schema.topology.SiteSummaryInfo"] = SiteSummaryInfoValidator()
	vr["ves.io.schema.topology.SiteTopologyRequest"] = SiteTopologyRequestValidator()
	vr["ves.io.schema.topology.SubnetData"] = SubnetDataValidator()
	vr["ves.io.schema.topology.SubnetMetaData"] = SubnetMetaDataValidator()
	vr["ves.io.schema.topology.SubnetSummaryInfo"] = SubnetSummaryInfoValidator()
	vr["ves.io.schema.topology.TopologyResponse"] = TopologyResponseValidator()

	vr["ves.io.schema.topology.AWSPolicyType"] = AWSPolicyTypeValidator()
	vr["ves.io.schema.topology.AWSRouteAttributes"] = AWSRouteAttributesValidator()
	vr["ves.io.schema.topology.AddressInfoType"] = AddressInfoTypeValidator()
	vr["ves.io.schema.topology.AzureResourceGroupInfo"] = AzureResourceGroupInfoValidator()
	vr["ves.io.schema.topology.AzureVnetPeer"] = AzureVnetPeerValidator()
	vr["ves.io.schema.topology.CloudPolicyType"] = CloudPolicyTypeValidator()
	vr["ves.io.schema.topology.DCClusterGroupType"] = DCClusterGroupTypeValidator()
	vr["ves.io.schema.topology.GCPRouteAttributes"] = GCPRouteAttributesValidator()
	vr["ves.io.schema.topology.InstanceType"] = InstanceTypeValidator()
	vr["ves.io.schema.topology.LoadBalancer"] = LoadBalancerValidator()
	vr["ves.io.schema.topology.MetaType"] = MetaTypeValidator()
	vr["ves.io.schema.topology.NetworkInterfaceType"] = NetworkInterfaceTypeValidator()
	vr["ves.io.schema.topology.NetworkPeerType"] = NetworkPeerTypeValidator()
	vr["ves.io.schema.topology.NetworkType"] = NetworkTypeValidator()
	vr["ves.io.schema.topology.ProviderInfo"] = ProviderInfoValidator()
	vr["ves.io.schema.topology.RouteTableType"] = RouteTableTypeValidator()
	vr["ves.io.schema.topology.RouteType"] = RouteTypeValidator()
	vr["ves.io.schema.topology.SiteMeshGroupType"] = SiteMeshGroupTypeValidator()
	vr["ves.io.schema.topology.SiteType"] = SiteTypeValidator()
	vr["ves.io.schema.topology.SubnetType"] = SubnetTypeValidator()
	vr["ves.io.schema.topology.TransitGatewayType"] = TransitGatewayTypeValidator()
	vr["ves.io.schema.topology.TunnelSetType"] = TunnelSetTypeValidator()
	vr["ves.io.schema.topology.TunnelType"] = TunnelTypeValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {
	sm["ves.io.schema.topology.CustomDataAPI"] = "data"

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

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

		customCSR.SwaggerRegistry["ves.io.schema.topology.Object"] = PrivateCustomAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.topology.PrivateCustomAPI"] = NewPrivateCustomAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.topology.PrivateCustomAPI"] = NewPrivateCustomAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.topology.PrivateCustomAPI"] = RegisterPrivateCustomAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.topology.PrivateCustomAPI"] = RegisterGwPrivateCustomAPIHandler
		customCSR.ServerRegistry["ves.io.schema.topology.PrivateCustomAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewPrivateCustomAPIServer(svc)
		}

	}()

	customCSR = mdr.PubCustomServiceRegistry

	func() {
		// set swagger jsons for our and external schemas

		customCSR.SwaggerRegistry["ves.io.schema.topology.Object"] = CustomDataAPISwaggerJSON

		customCSR.GrpcClientRegistry["ves.io.schema.topology.CustomDataAPI"] = NewCustomDataAPIGrpcClient
		customCSR.RestClientRegistry["ves.io.schema.topology.CustomDataAPI"] = NewCustomDataAPIRestClient
		if isExternal {
			return
		}
		mdr.SvcRegisterHandlers["ves.io.schema.topology.CustomDataAPI"] = RegisterCustomDataAPIServer
		mdr.SvcGwRegisterHandlers["ves.io.schema.topology.CustomDataAPI"] = RegisterGwCustomDataAPIHandler
		customCSR.ServerRegistry["ves.io.schema.topology.CustomDataAPI"] = func(svc svcfw.Service) server.APIHandler {
			return NewCustomDataAPIServer(svc)
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
