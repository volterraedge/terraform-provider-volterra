// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package schema

import (
	"gopkg.volterra.us/stdlib/db"
	"gopkg.volterra.us/stdlib/svcfw"
)

func initializeValidatorRegistry(vr map[string]db.Validator) {
	vr["ves.io.schema.CertInfoType"] = CertInfoTypeValidator()
	vr["ves.io.schema.CertificateParamsType"] = CertificateParamsTypeValidator()
	vr["ves.io.schema.DownstreamTlsParamsType"] = DownstreamTlsParamsTypeValidator()
	vr["ves.io.schema.HostAccessInfoType"] = HostAccessInfoTypeValidator()
	vr["ves.io.schema.TLSCoalescingOptions"] = TLSCoalescingOptionsValidator()
	vr["ves.io.schema.TlsParamsType"] = TlsParamsTypeValidator()
	vr["ves.io.schema.UpstreamCertificateParamsType"] = UpstreamCertificateParamsTypeValidator()
	vr["ves.io.schema.UpstreamTlsParamsType"] = UpstreamTlsParamsTypeValidator()

	vr["ves.io.schema.ErrorType"] = ErrorTypeValidator()

	vr["ves.io.schema.AwsSubnetList"] = AwsSubnetListValidator()
	vr["ves.io.schema.AwsVpcList"] = AwsVpcListValidator()
	vr["ves.io.schema.IpAddressType"] = IpAddressTypeValidator()
	vr["ves.io.schema.IpSubnetType"] = IpSubnetTypeValidator()
	vr["ves.io.schema.Ipv4AddressType"] = Ipv4AddressTypeValidator()
	vr["ves.io.schema.Ipv4SubnetType"] = Ipv4SubnetTypeValidator()
	vr["ves.io.schema.Ipv6AddressType"] = Ipv6AddressTypeValidator()
	vr["ves.io.schema.Ipv6SubnetType"] = Ipv6SubnetTypeValidator()
	vr["ves.io.schema.MacAddressType"] = MacAddressTypeValidator()
	vr["ves.io.schema.PrefixListType"] = PrefixListTypeValidator()

	vr["ves.io.schema.Dependencies"] = DependenciesValidator()
	vr["ves.io.schema.EnumInfo"] = EnumInfoValidator()
	vr["ves.io.schema.Key"] = KeyValidator()
	vr["ves.io.schema.Keys"] = KeysValidator()
	vr["ves.io.schema.LogField"] = LogFieldValidator()
	vr["ves.io.schema.MetricDef"] = MetricDefValidator()
	vr["ves.io.schema.On"] = OnValidator()
	vr["ves.io.schema.ThreatLevelInfo"] = ThreatLevelInfoValidator()
	vr["ves.io.schema.Via"] = ViaValidator()

	vr["ves.io.schema.DaemonEnvironmentType"] = DaemonEnvironmentTypeValidator()
	vr["ves.io.schema.DaemonTLSParamsType"] = DaemonTLSParamsTypeValidator()
	vr["ves.io.schema.DaemonTlsCertificateType"] = DaemonTlsCertificateTypeValidator()
	vr["ves.io.schema.DaemonTlsParametersType"] = DaemonTlsParametersTypeValidator()
	vr["ves.io.schema.DcClusterGroupStatus"] = DcClusterGroupStatusValidator()
	vr["ves.io.schema.OperMetaType"] = OperMetaTypeValidator()
	vr["ves.io.schema.ServiceParameters"] = ServiceParametersValidator()
	vr["ves.io.schema.SiteMeshGroupStatus"] = SiteMeshGroupStatusValidator()
	vr["ves.io.schema.SuggestValuesResp"] = SuggestValuesRespValidator()
	vr["ves.io.schema.SuggestedItem"] = SuggestedItemValidator()
	vr["ves.io.schema.SyncServerParamsType"] = SyncServerParamsTypeValidator()
	vr["ves.io.schema.UseragentType"] = UseragentTypeValidator()
	vr["ves.io.schema.VTRPGracefulRestartParamsType"] = VTRPGracefulRestartParamsTypeValidator()

	vr["ves.io.schema.Action"] = ActionValidator()
	vr["ves.io.schema.AppFirewallRefType"] = AppFirewallRefTypeValidator()
	vr["ves.io.schema.AppRoleAuthInfoType"] = AppRoleAuthInfoTypeValidator()
	vr["ves.io.schema.AuthnTypeBasicAuth"] = AuthnTypeBasicAuthValidator()
	vr["ves.io.schema.AuthnTypeHeaders"] = AuthnTypeHeadersValidator()
	vr["ves.io.schema.AuthnTypeQueryParams"] = AuthnTypeQueryParamsValidator()
	vr["ves.io.schema.BlindfoldSecretInfoType"] = BlindfoldSecretInfoTypeValidator()
	vr["ves.io.schema.BotDefenseFlowLabelAccountManagementChoiceType"] = BotDefenseFlowLabelAccountManagementChoiceTypeValidator()
	vr["ves.io.schema.BotDefenseFlowLabelAuthenticationChoiceType"] = BotDefenseFlowLabelAuthenticationChoiceTypeValidator()
	vr["ves.io.schema.BotDefenseFlowLabelCategoriesChoiceType"] = BotDefenseFlowLabelCategoriesChoiceTypeValidator()
	vr["ves.io.schema.BotDefenseFlowLabelFinancialServicesChoiceType"] = BotDefenseFlowLabelFinancialServicesChoiceTypeValidator()
	vr["ves.io.schema.BotDefenseFlowLabelFlightChoiceType"] = BotDefenseFlowLabelFlightChoiceTypeValidator()
	vr["ves.io.schema.BotDefenseFlowLabelProfileManagementChoiceType"] = BotDefenseFlowLabelProfileManagementChoiceTypeValidator()
	vr["ves.io.schema.BotDefenseFlowLabelSearchChoiceType"] = BotDefenseFlowLabelSearchChoiceTypeValidator()
	vr["ves.io.schema.BotDefenseFlowLabelShoppingGiftCardsChoiceType"] = BotDefenseFlowLabelShoppingGiftCardsChoiceTypeValidator()
	vr["ves.io.schema.BotDefenseTransactionResult"] = BotDefenseTransactionResultValidator()
	vr["ves.io.schema.BotDefenseTransactionResultCondition"] = BotDefenseTransactionResultConditionValidator()
	vr["ves.io.schema.BotDefenseTransactionResultType"] = BotDefenseTransactionResultTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelAccountManagementChoiceType"] = BotPolicyFlowLabelAccountManagementChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelAuthenticationChoiceType"] = BotPolicyFlowLabelAuthenticationChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelCategoriesChoiceType"] = BotPolicyFlowLabelCategoriesChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelCreditCardChoiceType"] = BotPolicyFlowLabelCreditCardChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelDeliveryServicesChoiceType"] = BotPolicyFlowLabelDeliveryServicesChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelFinancialServicesChoiceType"] = BotPolicyFlowLabelFinancialServicesChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelFlightChoiceType"] = BotPolicyFlowLabelFlightChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelGuestSessionChoiceType"] = BotPolicyFlowLabelGuestSessionChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelLoyaltyChoiceType"] = BotPolicyFlowLabelLoyaltyChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelMailingListChoiceType"] = BotPolicyFlowLabelMailingListChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelMediaChoiceType"] = BotPolicyFlowLabelMediaChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelMiscellaneousChoiceType"] = BotPolicyFlowLabelMiscellaneousChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelProfileManagementChoiceType"] = BotPolicyFlowLabelProfileManagementChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelQuotesChoiceType"] = BotPolicyFlowLabelQuotesChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelSearchChoiceType"] = BotPolicyFlowLabelSearchChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelShoppingGiftCardsChoiceType"] = BotPolicyFlowLabelShoppingGiftCardsChoiceTypeValidator()
	vr["ves.io.schema.BotPolicyFlowLabelSocialsChoiceType"] = BotPolicyFlowLabelSocialsChoiceTypeValidator()
	vr["ves.io.schema.BufferConfigType"] = BufferConfigTypeValidator()
	vr["ves.io.schema.CRMInfo"] = CRMInfoValidator()
	vr["ves.io.schema.ClearSecretInfoType"] = ClearSecretInfoTypeValidator()
	vr["ves.io.schema.CloudConnectRefType"] = CloudConnectRefTypeValidator()
	vr["ves.io.schema.CloudElasticIpRefListType"] = CloudElasticIpRefListTypeValidator()
	vr["ves.io.schema.ConditionType"] = ConditionTypeValidator()
	vr["ves.io.schema.CookieManipulationOptionType"] = CookieManipulationOptionTypeValidator()
	vr["ves.io.schema.CookieValueOption"] = CookieValueOptionValidator()
	vr["ves.io.schema.CorsPolicy"] = CorsPolicyValidator()
	vr["ves.io.schema.CsrfPolicy"] = CsrfPolicyValidator()
	vr["ves.io.schema.DateRange"] = DateRangeValidator()
	vr["ves.io.schema.DomainMatcherType"] = DomainMatcherTypeValidator()
	vr["ves.io.schema.DomainNameList"] = DomainNameListValidator()
	vr["ves.io.schema.DomainType"] = DomainTypeValidator()
	vr["ves.io.schema.Empty"] = EmptyValidator()
	vr["ves.io.schema.File"] = FileValidator()
	vr["ves.io.schema.ForwardProxyConfigType"] = ForwardProxyConfigTypeValidator()
	vr["ves.io.schema.FractionalPercent"] = FractionalPercentValidator()
	vr["ves.io.schema.HashAlgorithms"] = HashAlgorithmsValidator()
	vr["ves.io.schema.HeaderManipulationOptionType"] = HeaderManipulationOptionTypeValidator()
	vr["ves.io.schema.HeaderMatcherType"] = HeaderMatcherTypeValidator()
	vr["ves.io.schema.HeaderTransformationType"] = HeaderTransformationTypeValidator()
	vr["ves.io.schema.HostIdentifier"] = HostIdentifierValidator()
	vr["ves.io.schema.IPv4AddressRange"] = IPv4AddressRangeValidator()
	vr["ves.io.schema.IPv6AddressRange"] = IPv6AddressRangeValidator()
	vr["ves.io.schema.InitializerType"] = InitializerTypeValidator()
	vr["ves.io.schema.InitializersType"] = InitializersTypeValidator()
	vr["ves.io.schema.InterfaceIdentifier"] = InterfaceIdentifierValidator()
	vr["ves.io.schema.InterfaceOrNetwork"] = InterfaceOrNetworkValidator()
	vr["ves.io.schema.IpPrefixSetRefType"] = IpPrefixSetRefTypeValidator()
	vr["ves.io.schema.KubeRefType"] = KubeRefTypeValidator()
	vr["ves.io.schema.L4DestType"] = L4DestTypeValidator()
	vr["ves.io.schema.LabelMatcherType"] = LabelMatcherTypeValidator()
	vr["ves.io.schema.LabelSelectorType"] = LabelSelectorTypeValidator()
	vr["ves.io.schema.ListMetaType"] = ListMetaTypeValidator()
	vr["ves.io.schema.MessageMetaType"] = MessageMetaTypeValidator()
	vr["ves.io.schema.MetricTypeData"] = MetricTypeDataValidator()
	vr["ves.io.schema.MetricValue"] = MetricValueValidator()
	vr["ves.io.schema.NamespaceAccessType"] = NamespaceAccessTypeValidator()
	vr["ves.io.schema.NamespaceRoleType"] = NamespaceRoleTypeValidator()
	vr["ves.io.schema.NetworkInterfaceRefType"] = NetworkInterfaceRefTypeValidator()
	vr["ves.io.schema.NetworkRefType"] = NetworkRefTypeValidator()
	vr["ves.io.schema.NetworkSiteRefSelector"] = NetworkSiteRefSelectorValidator()
	vr["ves.io.schema.NetworkingStackType"] = NetworkingStackTypeValidator()
	vr["ves.io.schema.NextHopType"] = NextHopTypeValidator()
	vr["ves.io.schema.NodeInterfaceInfo"] = NodeInterfaceInfoValidator()
	vr["ves.io.schema.NodeInterfaceType"] = NodeInterfaceTypeValidator()
	vr["ves.io.schema.ObjectCreateMetaType"] = ObjectCreateMetaTypeValidator()
	vr["ves.io.schema.ObjectGetMetaType"] = ObjectGetMetaTypeValidator()
	vr["ves.io.schema.ObjectMetaType"] = ObjectMetaTypeValidator()
	vr["ves.io.schema.ObjectRefType"] = ObjectRefTypeValidator()
	vr["ves.io.schema.ObjectReplaceMetaType"] = ObjectReplaceMetaTypeValidator()
	vr["ves.io.schema.PathMatcherType"] = PathMatcherTypeValidator()
	vr["ves.io.schema.PolicerRefType"] = PolicerRefTypeValidator()
	vr["ves.io.schema.PortMatcherType"] = PortMatcherTypeValidator()
	vr["ves.io.schema.PortRangesType"] = PortRangesTypeValidator()
	vr["ves.io.schema.PortValueType"] = PortValueTypeValidator()
	vr["ves.io.schema.ProtocolPolicerRefType"] = ProtocolPolicerRefTypeValidator()
	vr["ves.io.schema.QueryParameterMatcherType"] = QueryParameterMatcherTypeValidator()
	vr["ves.io.schema.RegexMatchRewrite"] = RegexMatchRewriteValidator()
	vr["ves.io.schema.ResponseMeta"] = ResponseMetaValidator()
	vr["ves.io.schema.RestAuthInfoType"] = RestAuthInfoTypeValidator()
	vr["ves.io.schema.RetryBackOff"] = RetryBackOffValidator()
	vr["ves.io.schema.RetryPolicyType"] = RetryPolicyTypeValidator()
	vr["ves.io.schema.RoleListType"] = RoleListTypeValidator()
	vr["ves.io.schema.RouteMatch"] = RouteMatchValidator()
	vr["ves.io.schema.RouteTarget"] = RouteTargetValidator()
	vr["ves.io.schema.RouteTarget2ByteAsn"] = RouteTarget2ByteAsnValidator()
	vr["ves.io.schema.RouteTarget4ByteAsn"] = RouteTarget4ByteAsnValidator()
	vr["ves.io.schema.RouteTargetIPv4Addr"] = RouteTargetIPv4AddrValidator()
	vr["ves.io.schema.SecretType"] = SecretTypeValidator()
	vr["ves.io.schema.SegmentRefType"] = SegmentRefTypeValidator()
	vr["ves.io.schema.SetCookieValueOption"] = SetCookieValueOptionValidator()
	vr["ves.io.schema.SiteInfo"] = SiteInfoValidator()
	vr["ves.io.schema.SiteRefType"] = SiteRefTypeValidator()
	vr["ves.io.schema.SiteReferenceListType"] = SiteReferenceListTypeValidator()
	vr["ves.io.schema.SiteReferenceType"] = SiteReferenceTypeValidator()
	vr["ves.io.schema.SiteVirtualSiteRefSelector"] = SiteVirtualSiteRefSelectorValidator()
	vr["ves.io.schema.StaticRouteType"] = StaticRouteTypeValidator()
	vr["ves.io.schema.StatusMetaType"] = StatusMetaTypeValidator()
	vr["ves.io.schema.StatusType"] = StatusTypeValidator()
	vr["ves.io.schema.SystemObjectGetMetaType"] = SystemObjectGetMetaTypeValidator()
	vr["ves.io.schema.SystemObjectMetaType"] = SystemObjectMetaTypeValidator()
	vr["ves.io.schema.TlsCertificateType"] = TlsCertificateTypeValidator()
	vr["ves.io.schema.TlsInterceptionPolicy"] = TlsInterceptionPolicyValidator()
	vr["ves.io.schema.TlsInterceptionRule"] = TlsInterceptionRuleValidator()
	vr["ves.io.schema.TlsInterceptionType"] = TlsInterceptionTypeValidator()
	vr["ves.io.schema.TlsValidationParamsType"] = TlsValidationParamsTypeValidator()
	vr["ves.io.schema.TrendValue"] = TrendValueValidator()
	vr["ves.io.schema.TrustedCAList"] = TrustedCAListValidator()
	vr["ves.io.schema.UpstreamConnPoolReuseType"] = UpstreamConnPoolReuseTypeValidator()
	vr["ves.io.schema.VSiteRefType"] = VSiteRefTypeValidator()
	vr["ves.io.schema.VaultAccessInfoType"] = VaultAccessInfoTypeValidator()
	vr["ves.io.schema.VaultAuthInfoType"] = VaultAuthInfoTypeValidator()
	vr["ves.io.schema.VaultSecretInfoType"] = VaultSecretInfoTypeValidator()
	vr["ves.io.schema.VaultSecretType"] = VaultSecretTypeValidator()
	vr["ves.io.schema.ViewRefType"] = ViewRefTypeValidator()
	vr["ves.io.schema.VirtualNetworkReferenceType"] = VirtualNetworkReferenceTypeValidator()
	vr["ves.io.schema.VirtualNetworkSelectorType"] = VirtualNetworkSelectorTypeValidator()
	vr["ves.io.schema.VolterraSecretType"] = VolterraSecretTypeValidator()
	vr["ves.io.schema.WafType"] = WafTypeValidator()
	vr["ves.io.schema.WingmanSecretInfoType"] = WingmanSecretInfoTypeValidator()

	vr["ves.io.schema.AnyRules"] = AnyRulesValidator()
	vr["ves.io.schema.BoolRules"] = BoolRulesValidator()
	vr["ves.io.schema.BytesRules"] = BytesRulesValidator()
	vr["ves.io.schema.DateRangeRules"] = DateRangeRulesValidator()
	vr["ves.io.schema.DoubleRules"] = DoubleRulesValidator()
	vr["ves.io.schema.DurationRules"] = DurationRulesValidator()
	vr["ves.io.schema.EnumRules"] = EnumRulesValidator()
	vr["ves.io.schema.FieldRules"] = FieldRulesValidator()
	vr["ves.io.schema.Fixed32Rules"] = Fixed32RulesValidator()
	vr["ves.io.schema.Fixed64Rules"] = Fixed64RulesValidator()
	vr["ves.io.schema.FloatRules"] = FloatRulesValidator()
	vr["ves.io.schema.Int32Rules"] = Int32RulesValidator()
	vr["ves.io.schema.Int64Rules"] = Int64RulesValidator()
	vr["ves.io.schema.MapRules"] = MapRulesValidator()
	vr["ves.io.schema.MessageRules"] = MessageRulesValidator()
	vr["ves.io.schema.RepeatedRules"] = RepeatedRulesValidator()
	vr["ves.io.schema.SFixed32Rules"] = SFixed32RulesValidator()
	vr["ves.io.schema.SFixed64Rules"] = SFixed64RulesValidator()
	vr["ves.io.schema.SInt32Rules"] = SInt32RulesValidator()
	vr["ves.io.schema.SInt64Rules"] = SInt64RulesValidator()
	vr["ves.io.schema.StringRules"] = StringRulesValidator()
	vr["ves.io.schema.TimestampRules"] = TimestampRulesValidator()
	vr["ves.io.schema.UInt32Rules"] = UInt32RulesValidator()
	vr["ves.io.schema.UInt64Rules"] = UInt64RulesValidator()

	vr["ves.io.schema.ChoiceItem"] = ChoiceItemValidator()
	vr["ves.io.schema.ChoiceItemList"] = ChoiceItemListValidator()
	vr["ves.io.schema.Choices"] = ChoicesValidator()
	vr["ves.io.schema.Column"] = ColumnValidator()
	vr["ves.io.schema.Columns"] = ColumnsValidator()
	vr["ves.io.schema.ConstraintLength"] = ConstraintLengthValidator()
	vr["ves.io.schema.DateRangeOptions"] = DateRangeOptionsValidator()
	vr["ves.io.schema.DisplayElements"] = DisplayElementsValidator()
	vr["ves.io.schema.DisplayExistsNotExists"] = DisplayExistsNotExistsValidator()
	vr["ves.io.schema.DisplayKVItem"] = DisplayKVItemValidator()
	vr["ves.io.schema.DisplayKVItemList"] = DisplayKVItemListValidator()
	vr["ves.io.schema.DisplayLength"] = DisplayLengthValidator()
	vr["ves.io.schema.DisplayMapElements"] = DisplayMapElementsValidator()
	vr["ves.io.schema.DisplayOneValue"] = DisplayOneValueValidator()
	vr["ves.io.schema.DisplayOneof"] = DisplayOneofValidator()
	vr["ves.io.schema.DisplayOneofItem"] = DisplayOneofItemValidator()
	vr["ves.io.schema.DisplayOneofItemList"] = DisplayOneofItemListValidator()
	vr["ves.io.schema.DisplayValue"] = DisplayValueValidator()
	vr["ves.io.schema.FieldViewOptions"] = FieldViewOptionsValidator()
	vr["ves.io.schema.HiddenConditions"] = HiddenConditionsValidator()
	vr["ves.io.schema.HintWithAction"] = HintWithActionValidator()
	vr["ves.io.schema.HintWithLink"] = HintWithLinkValidator()
	vr["ves.io.schema.LabelKeyClassList"] = LabelKeyClassListValidator()
	vr["ves.io.schema.LabelSelectorOperatorList"] = LabelSelectorOperatorListValidator()
	vr["ves.io.schema.MapOptions"] = MapOptionsValidator()
	vr["ves.io.schema.RepeatedOptions"] = RepeatedOptionsValidator()
	vr["ves.io.schema.StoredObjectURL"] = StoredObjectURLValidator()
	vr["ves.io.schema.SuffixText"] = SuffixTextValidator()
	vr["ves.io.schema.SuggestedValues"] = SuggestedValuesValidator()
	vr["ves.io.schema.Tile"] = TileValidator()
	vr["ves.io.schema.Tiles"] = TilesValidator()

}

func initializeEntryRegistry(mdr *svcfw.MDRegistry) {

}

func initializeRPCRegistry(mdr *svcfw.MDRegistry) {

}

func initializeAPIGwServiceSlugsRegistry(sm map[string]string) {

}

func initializeP0PolicyRegistry(sm map[string]svcfw.P0PolicyInfo) {

}

func initializeCRUDServiceRegistry(mdr *svcfw.MDRegistry, isExternal bool) {
	var (
		csr       *svcfw.CRUDServiceRegistry
		customCSR *svcfw.CustomServiceRegistry
	)
	_, _ = csr, customCSR

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
