// Copyright (c) 2022 F5, Inc. All rights reserved.
// Code generated by ves-gen-schema-go. DO NOT EDIT.

package saas_config

func (o *DBObject) GetSubscribedAddonServices() []string {
	return o.GetSpec().GetGcSpec().GetAddonServicesSubscribed()
}
