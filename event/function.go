// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package event

type CreatePullPointSubscriptionFunction struct{}

func (_ *CreatePullPointSubscriptionFunction) Request() interface{} {
	return &CreatePullPointSubscription{}
}
func (_ *CreatePullPointSubscriptionFunction) Response() interface{} {
	return &CreatePullPointSubscriptionResponse{}
}

type GetEventPropertiesFunction struct{}

func (_ *GetEventPropertiesFunction) Request() interface{} {
	return &GetEventProperties{}
}
func (_ *GetEventPropertiesFunction) Response() interface{} {
	return &GetEventPropertiesResponse{}
}

type GetServiceCapabilitiesFunction struct{}

func (_ *GetServiceCapabilitiesFunction) Request() interface{} {
	return &GetServiceCapabilities{}
}
func (_ *GetServiceCapabilitiesFunction) Response() interface{} {
	return &GetServiceCapabilitiesResponse{}
}

type PullMessagesFunction struct{}

func (_ *PullMessagesFunction) Request() interface{} {
	return &PullMessages{}
}
func (_ *PullMessagesFunction) Response() interface{} {
	return &PullMessagesResponse{}
}

type RenewFunction struct{}

func (_ *RenewFunction) Request() interface{} {
	return &Renew{}
}
func (_ *RenewFunction) Response() interface{} {
	return &RenewResponse{}
}

type SeekFunction struct{}

func (_ *SeekFunction) Request() interface{} {
	return &Seek{}
}
func (_ *SeekFunction) Response() interface{} {
	return &SeekResponse{}
}

type SetSynchronizationPointFunction struct{}

func (_ *SetSynchronizationPointFunction) Request() interface{} {
	return &SetSynchronizationPoint{}
}
func (_ *SetSynchronizationPointFunction) Response() interface{} {
	return &SetSynchronizationPointResponse{}
}

type SubscribeFunction struct{}

func (_ *SubscribeFunction) Request() interface{} {
	return &Subscribe{}
}
func (_ *SubscribeFunction) Response() interface{} {
	return &SubscribeResponse{}
}

type SubscriptionReferenceFunction struct{}

func (_ *SubscriptionReferenceFunction) Request() interface{} {
	return &SubscriptionReference{}
}
func (_ *SubscriptionReferenceFunction) Response() interface{} {
	return &SubscriptionReferenceResponse{}
}

type UnsubscribeFunction struct{}

func (_ *UnsubscribeFunction) Request() interface{} {
	return &Unsubscribe{}
}
func (_ *UnsubscribeFunction) Response() interface{} {
	return &UnsubscribeResponse{}
}
