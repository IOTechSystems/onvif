// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package media2

type AddConfigurationFunction struct{}

func (_ *AddConfigurationFunction) Request() interface{} {
	return &AddConfiguration{}
}
func (_ *AddConfigurationFunction) Response() interface{} {
	return &AddConfigurationResponse{}
}

type GetAnalyticsConfigurationsFunction struct{}

func (_ *GetAnalyticsConfigurationsFunction) Request() interface{} {
	return &GetAnalyticsConfigurations{}
}
func (_ *GetAnalyticsConfigurationsFunction) Response() interface{} {
	return &GetAnalyticsConfigurationsResponse{}
}

type GetProfilesFunction struct{}

func (_ *GetProfilesFunction) Request() interface{} {
	return &GetProfiles{}
}
func (_ *GetProfilesFunction) Response() interface{} {
	return &GetProfilesResponse{}
}

type RemoveConfigurationFunction struct{}

func (_ *RemoveConfigurationFunction) Request() interface{} {
	return &RemoveConfiguration{}
}
func (_ *RemoveConfigurationFunction) Response() interface{} {
	return &RemoveConfigurationResponse{}
}
