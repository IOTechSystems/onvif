// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package ptz

type AbsoluteMoveFunction struct{}

func (_ *AbsoluteMoveFunction) Request() interface{} {
	return &AbsoluteMove{}
}
func (_ *AbsoluteMoveFunction) Response() interface{} {
	return &AbsoluteMoveResponse{}
}

type ContinuousMoveFunction struct{}

func (_ *ContinuousMoveFunction) Request() interface{} {
	return &ContinuousMove{}
}
func (_ *ContinuousMoveFunction) Response() interface{} {
	return &ContinuousMoveResponse{}
}

type CreatePresetTourFunction struct{}

func (_ *CreatePresetTourFunction) Request() interface{} {
	return &CreatePresetTour{}
}
func (_ *CreatePresetTourFunction) Response() interface{} {
	return &CreatePresetTourResponse{}
}

type GeoMoveFunction struct{}

func (_ *GeoMoveFunction) Request() interface{} {
	return &GeoMove{}
}
func (_ *GeoMoveFunction) Response() interface{} {
	return &GeoMoveResponse{}
}

type GetCompatibleConfigurationsFunction struct{}

func (_ *GetCompatibleConfigurationsFunction) Request() interface{} {
	return &GetCompatibleConfigurations{}
}
func (_ *GetCompatibleConfigurationsFunction) Response() interface{} {
	return &GetCompatibleConfigurationsResponse{}
}

type GetConfigurationFunction struct{}

func (_ *GetConfigurationFunction) Request() interface{} {
	return &GetConfiguration{}
}
func (_ *GetConfigurationFunction) Response() interface{} {
	return &GetConfigurationResponse{}
}

type GetConfigurationOptionsFunction struct{}

func (_ *GetConfigurationOptionsFunction) Request() interface{} {
	return &GetConfigurationOptions{}
}
func (_ *GetConfigurationOptionsFunction) Response() interface{} {
	return &GetConfigurationOptionsResponse{}
}

type GetConfigurationsFunction struct{}

func (_ *GetConfigurationsFunction) Request() interface{} {
	return &GetConfigurations{}
}
func (_ *GetConfigurationsFunction) Response() interface{} {
	return &GetConfigurationsResponse{}
}

type GetNodeFunction struct{}

func (_ *GetNodeFunction) Request() interface{} {
	return &GetNode{}
}
func (_ *GetNodeFunction) Response() interface{} {
	return &GetNodeResponse{}
}

type GetNodesFunction struct{}

func (_ *GetNodesFunction) Request() interface{} {
	return &GetNodes{}
}
func (_ *GetNodesFunction) Response() interface{} {
	return &GetNodesResponse{}
}

type GetPresetTourFunction struct{}

func (_ *GetPresetTourFunction) Request() interface{} {
	return &GetPresetTour{}
}
func (_ *GetPresetTourFunction) Response() interface{} {
	return &GetPresetTourResponse{}
}

type GetPresetTourOptionsFunction struct{}

func (_ *GetPresetTourOptionsFunction) Request() interface{} {
	return &GetPresetTourOptions{}
}
func (_ *GetPresetTourOptionsFunction) Response() interface{} {
	return &GetPresetTourOptionsResponse{}
}

type GetPresetToursFunction struct{}

func (_ *GetPresetToursFunction) Request() interface{} {
	return &GetPresetTours{}
}
func (_ *GetPresetToursFunction) Response() interface{} {
	return &GetPresetToursResponse{}
}

type GetPresetsFunction struct{}

func (_ *GetPresetsFunction) Request() interface{} {
	return &GetPresets{}
}
func (_ *GetPresetsFunction) Response() interface{} {
	return &GetPresetsResponse{}
}

type GetServiceCapabilitiesFunction struct{}

func (_ *GetServiceCapabilitiesFunction) Request() interface{} {
	return &GetServiceCapabilities{}
}
func (_ *GetServiceCapabilitiesFunction) Response() interface{} {
	return &GetServiceCapabilitiesResponse{}
}

type GetStatusFunction struct{}

func (_ *GetStatusFunction) Request() interface{} {
	return &GetStatus{}
}
func (_ *GetStatusFunction) Response() interface{} {
	return &GetStatusResponse{}
}

type GotoHomePositionFunction struct{}

func (_ *GotoHomePositionFunction) Request() interface{} {
	return &GotoHomePosition{}
}
func (_ *GotoHomePositionFunction) Response() interface{} {
	return &GotoHomePositionResponse{}
}

type GotoPresetFunction struct{}

func (_ *GotoPresetFunction) Request() interface{} {
	return &GotoPreset{}
}
func (_ *GotoPresetFunction) Response() interface{} {
	return &GotoPresetResponse{}
}

type ModifyPresetTourFunction struct{}

func (_ *ModifyPresetTourFunction) Request() interface{} {
	return &ModifyPresetTour{}
}
func (_ *ModifyPresetTourFunction) Response() interface{} {
	return &ModifyPresetTourResponse{}
}

type OperatePresetTourFunction struct{}

func (_ *OperatePresetTourFunction) Request() interface{} {
	return &OperatePresetTour{}
}
func (_ *OperatePresetTourFunction) Response() interface{} {
	return &OperatePresetTourResponse{}
}

type RelativeMoveFunction struct{}

func (_ *RelativeMoveFunction) Request() interface{} {
	return &RelativeMove{}
}
func (_ *RelativeMoveFunction) Response() interface{} {
	return &RelativeMoveResponse{}
}

type RemovePresetFunction struct{}

func (_ *RemovePresetFunction) Request() interface{} {
	return &RemovePreset{}
}
func (_ *RemovePresetFunction) Response() interface{} {
	return &RemovePresetResponse{}
}

type RemovePresetTourFunction struct{}

func (_ *RemovePresetTourFunction) Request() interface{} {
	return &RemovePresetTour{}
}
func (_ *RemovePresetTourFunction) Response() interface{} {
	return &RemovePresetTourResponse{}
}

type SendAuxiliaryCommandFunction struct{}

func (_ *SendAuxiliaryCommandFunction) Request() interface{} {
	return &SendAuxiliaryCommand{}
}
func (_ *SendAuxiliaryCommandFunction) Response() interface{} {
	return &SendAuxiliaryCommandResponse{}
}

type SetConfigurationFunction struct{}

func (_ *SetConfigurationFunction) Request() interface{} {
	return &SetConfiguration{}
}
func (_ *SetConfigurationFunction) Response() interface{} {
	return &SetConfigurationResponse{}
}

type SetHomePositionFunction struct{}

func (_ *SetHomePositionFunction) Request() interface{} {
	return &SetHomePosition{}
}
func (_ *SetHomePositionFunction) Response() interface{} {
	return &SetHomePositionResponse{}
}

type SetPresetFunction struct{}

func (_ *SetPresetFunction) Request() interface{} {
	return &SetPreset{}
}
func (_ *SetPresetFunction) Response() interface{} {
	return &SetPresetResponse{}
}

type StopFunction struct{}

func (_ *StopFunction) Request() interface{} {
	return &Stop{}
}
func (_ *StopFunction) Response() interface{} {
	return &StopResponse{}
}
