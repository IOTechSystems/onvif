// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

package media

type AddAudioDecoderConfigurationFunction struct{}

func (_ *AddAudioDecoderConfigurationFunction) Request() interface{} {
	return &AddAudioDecoderConfiguration{}
}
func (_ *AddAudioDecoderConfigurationFunction) Response() interface{} {
	return &AddAudioDecoderConfigurationResponse{}
}

type AddAudioEncoderConfigurationFunction struct{}

func (_ *AddAudioEncoderConfigurationFunction) Request() interface{} {
	return &AddAudioEncoderConfiguration{}
}
func (_ *AddAudioEncoderConfigurationFunction) Response() interface{} {
	return &AddAudioEncoderConfigurationResponse{}
}

type AddAudioOutputConfigurationFunction struct{}

func (_ *AddAudioOutputConfigurationFunction) Request() interface{} {
	return &AddAudioOutputConfiguration{}
}
func (_ *AddAudioOutputConfigurationFunction) Response() interface{} {
	return &AddAudioOutputConfigurationResponse{}
}

type AddAudioSourceConfigurationFunction struct{}

func (_ *AddAudioSourceConfigurationFunction) Request() interface{} {
	return &AddAudioSourceConfiguration{}
}
func (_ *AddAudioSourceConfigurationFunction) Response() interface{} {
	return &AddAudioSourceConfigurationResponse{}
}

type AddMetadataConfigurationFunction struct{}

func (_ *AddMetadataConfigurationFunction) Request() interface{} {
	return &AddMetadataConfiguration{}
}
func (_ *AddMetadataConfigurationFunction) Response() interface{} {
	return &AddMetadataConfigurationResponse{}
}

type AddPTZConfigurationFunction struct{}

func (_ *AddPTZConfigurationFunction) Request() interface{} {
	return &AddPTZConfiguration{}
}
func (_ *AddPTZConfigurationFunction) Response() interface{} {
	return &AddPTZConfigurationResponse{}
}

type AddVideoAnalyticsConfigurationFunction struct{}

func (_ *AddVideoAnalyticsConfigurationFunction) Request() interface{} {
	return &AddVideoAnalyticsConfiguration{}
}
func (_ *AddVideoAnalyticsConfigurationFunction) Response() interface{} {
	return &AddVideoAnalyticsConfigurationResponse{}
}

type AddVideoEncoderConfigurationFunction struct{}

func (_ *AddVideoEncoderConfigurationFunction) Request() interface{} {
	return &AddVideoEncoderConfiguration{}
}
func (_ *AddVideoEncoderConfigurationFunction) Response() interface{} {
	return &AddVideoEncoderConfigurationResponse{}
}

type AddVideoSourceConfigurationFunction struct{}

func (_ *AddVideoSourceConfigurationFunction) Request() interface{} {
	return &AddVideoSourceConfiguration{}
}
func (_ *AddVideoSourceConfigurationFunction) Response() interface{} {
	return &AddVideoSourceConfigurationResponse{}
}

type CreateOSDFunction struct{}

func (_ *CreateOSDFunction) Request() interface{} {
	return &CreateOSD{}
}
func (_ *CreateOSDFunction) Response() interface{} {
	return &CreateOSDResponse{}
}

type CreateProfileFunction struct{}

func (_ *CreateProfileFunction) Request() interface{} {
	return &CreateProfile{}
}
func (_ *CreateProfileFunction) Response() interface{} {
	return &CreateProfileResponse{}
}

type DeleteOSDFunction struct{}

func (_ *DeleteOSDFunction) Request() interface{} {
	return &DeleteOSD{}
}
func (_ *DeleteOSDFunction) Response() interface{} {
	return &DeleteOSDResponse{}
}

type DeleteProfileFunction struct{}

func (_ *DeleteProfileFunction) Request() interface{} {
	return &DeleteProfile{}
}
func (_ *DeleteProfileFunction) Response() interface{} {
	return &DeleteProfileResponse{}
}

type GetAudioDecoderConfigurationFunction struct{}

func (_ *GetAudioDecoderConfigurationFunction) Request() interface{} {
	return &GetAudioDecoderConfiguration{}
}
func (_ *GetAudioDecoderConfigurationFunction) Response() interface{} {
	return &GetAudioDecoderConfigurationResponse{}
}

type GetAudioDecoderConfigurationOptionsFunction struct{}

func (_ *GetAudioDecoderConfigurationOptionsFunction) Request() interface{} {
	return &GetAudioDecoderConfigurationOptions{}
}
func (_ *GetAudioDecoderConfigurationOptionsFunction) Response() interface{} {
	return &GetAudioDecoderConfigurationOptionsResponse{}
}

type GetAudioDecoderConfigurationsFunction struct{}

func (_ *GetAudioDecoderConfigurationsFunction) Request() interface{} {
	return &GetAudioDecoderConfigurations{}
}
func (_ *GetAudioDecoderConfigurationsFunction) Response() interface{} {
	return &GetAudioDecoderConfigurationsResponse{}
}

type GetAudioEncoderConfigurationFunction struct{}

func (_ *GetAudioEncoderConfigurationFunction) Request() interface{} {
	return &GetAudioEncoderConfiguration{}
}
func (_ *GetAudioEncoderConfigurationFunction) Response() interface{} {
	return &GetAudioEncoderConfigurationResponse{}
}

type GetAudioEncoderConfigurationOptionsFunction struct{}

func (_ *GetAudioEncoderConfigurationOptionsFunction) Request() interface{} {
	return &GetAudioEncoderConfigurationOptions{}
}
func (_ *GetAudioEncoderConfigurationOptionsFunction) Response() interface{} {
	return &GetAudioEncoderConfigurationOptionsResponse{}
}

type GetAudioEncoderConfigurationsFunction struct{}

func (_ *GetAudioEncoderConfigurationsFunction) Request() interface{} {
	return &GetAudioEncoderConfigurations{}
}
func (_ *GetAudioEncoderConfigurationsFunction) Response() interface{} {
	return &GetAudioEncoderConfigurationsResponse{}
}

type GetAudioOutputConfigurationFunction struct{}

func (_ *GetAudioOutputConfigurationFunction) Request() interface{} {
	return &GetAudioOutputConfiguration{}
}
func (_ *GetAudioOutputConfigurationFunction) Response() interface{} {
	return &GetAudioOutputConfigurationResponse{}
}

type GetAudioOutputConfigurationOptionsFunction struct{}

func (_ *GetAudioOutputConfigurationOptionsFunction) Request() interface{} {
	return &GetAudioOutputConfigurationOptions{}
}
func (_ *GetAudioOutputConfigurationOptionsFunction) Response() interface{} {
	return &GetAudioOutputConfigurationOptionsResponse{}
}

type GetAudioOutputConfigurationsFunction struct{}

func (_ *GetAudioOutputConfigurationsFunction) Request() interface{} {
	return &GetAudioOutputConfigurations{}
}
func (_ *GetAudioOutputConfigurationsFunction) Response() interface{} {
	return &GetAudioOutputConfigurationsResponse{}
}

type GetAudioOutputsFunction struct{}

func (_ *GetAudioOutputsFunction) Request() interface{} {
	return &GetAudioOutputs{}
}
func (_ *GetAudioOutputsFunction) Response() interface{} {
	return &GetAudioOutputsResponse{}
}

type GetAudioSourceConfigurationFunction struct{}

func (_ *GetAudioSourceConfigurationFunction) Request() interface{} {
	return &GetAudioSourceConfiguration{}
}
func (_ *GetAudioSourceConfigurationFunction) Response() interface{} {
	return &GetAudioSourceConfigurationResponse{}
}

type GetAudioSourceConfigurationOptionsFunction struct{}

func (_ *GetAudioSourceConfigurationOptionsFunction) Request() interface{} {
	return &GetAudioSourceConfigurationOptions{}
}
func (_ *GetAudioSourceConfigurationOptionsFunction) Response() interface{} {
	return &GetAudioSourceConfigurationOptionsResponse{}
}

type GetAudioSourceConfigurationsFunction struct{}

func (_ *GetAudioSourceConfigurationsFunction) Request() interface{} {
	return &GetAudioSourceConfigurations{}
}
func (_ *GetAudioSourceConfigurationsFunction) Response() interface{} {
	return &GetAudioSourceConfigurationsResponse{}
}

type GetAudioSourcesFunction struct{}

func (_ *GetAudioSourcesFunction) Request() interface{} {
	return &GetAudioSources{}
}
func (_ *GetAudioSourcesFunction) Response() interface{} {
	return &GetAudioSourcesResponse{}
}

type GetCompatibleAudioDecoderConfigurationsFunction struct{}

func (_ *GetCompatibleAudioDecoderConfigurationsFunction) Request() interface{} {
	return &GetCompatibleAudioDecoderConfigurations{}
}
func (_ *GetCompatibleAudioDecoderConfigurationsFunction) Response() interface{} {
	return &GetCompatibleAudioDecoderConfigurationsResponse{}
}

type GetCompatibleAudioEncoderConfigurationsFunction struct{}

func (_ *GetCompatibleAudioEncoderConfigurationsFunction) Request() interface{} {
	return &GetCompatibleAudioEncoderConfigurations{}
}
func (_ *GetCompatibleAudioEncoderConfigurationsFunction) Response() interface{} {
	return &GetCompatibleAudioEncoderConfigurationsResponse{}
}

type GetCompatibleAudioOutputConfigurationsFunction struct{}

func (_ *GetCompatibleAudioOutputConfigurationsFunction) Request() interface{} {
	return &GetCompatibleAudioOutputConfigurations{}
}
func (_ *GetCompatibleAudioOutputConfigurationsFunction) Response() interface{} {
	return &GetCompatibleAudioOutputConfigurationsResponse{}
}

type GetCompatibleAudioSourceConfigurationsFunction struct{}

func (_ *GetCompatibleAudioSourceConfigurationsFunction) Request() interface{} {
	return &GetCompatibleAudioSourceConfigurations{}
}
func (_ *GetCompatibleAudioSourceConfigurationsFunction) Response() interface{} {
	return &GetCompatibleAudioSourceConfigurationsResponse{}
}

type GetCompatibleMetadataConfigurationsFunction struct{}

func (_ *GetCompatibleMetadataConfigurationsFunction) Request() interface{} {
	return &GetCompatibleMetadataConfigurations{}
}
func (_ *GetCompatibleMetadataConfigurationsFunction) Response() interface{} {
	return &GetCompatibleMetadataConfigurationsResponse{}
}

type GetCompatibleVideoAnalyticsConfigurationsFunction struct{}

func (_ *GetCompatibleVideoAnalyticsConfigurationsFunction) Request() interface{} {
	return &GetCompatibleVideoAnalyticsConfigurations{}
}
func (_ *GetCompatibleVideoAnalyticsConfigurationsFunction) Response() interface{} {
	return &GetCompatibleVideoAnalyticsConfigurationsResponse{}
}

type GetCompatibleVideoEncoderConfigurationsFunction struct{}

func (_ *GetCompatibleVideoEncoderConfigurationsFunction) Request() interface{} {
	return &GetCompatibleVideoEncoderConfigurations{}
}
func (_ *GetCompatibleVideoEncoderConfigurationsFunction) Response() interface{} {
	return &GetCompatibleVideoEncoderConfigurationsResponse{}
}

type GetCompatibleVideoSourceConfigurationsFunction struct{}

func (_ *GetCompatibleVideoSourceConfigurationsFunction) Request() interface{} {
	return &GetCompatibleVideoSourceConfigurations{}
}
func (_ *GetCompatibleVideoSourceConfigurationsFunction) Response() interface{} {
	return &GetCompatibleVideoSourceConfigurationsResponse{}
}

type GetGuaranteedNumberOfVideoEncoderInstancesFunction struct{}

func (_ *GetGuaranteedNumberOfVideoEncoderInstancesFunction) Request() interface{} {
	return &GetGuaranteedNumberOfVideoEncoderInstances{}
}
func (_ *GetGuaranteedNumberOfVideoEncoderInstancesFunction) Response() interface{} {
	return &GetGuaranteedNumberOfVideoEncoderInstancesResponse{}
}

type GetMetadataConfigurationFunction struct{}

func (_ *GetMetadataConfigurationFunction) Request() interface{} {
	return &GetMetadataConfiguration{}
}
func (_ *GetMetadataConfigurationFunction) Response() interface{} {
	return &GetMetadataConfigurationResponse{}
}

type GetMetadataConfigurationOptionsFunction struct{}

func (_ *GetMetadataConfigurationOptionsFunction) Request() interface{} {
	return &GetMetadataConfigurationOptions{}
}
func (_ *GetMetadataConfigurationOptionsFunction) Response() interface{} {
	return &GetMetadataConfigurationOptionsResponse{}
}

type GetMetadataConfigurationsFunction struct{}

func (_ *GetMetadataConfigurationsFunction) Request() interface{} {
	return &GetMetadataConfigurations{}
}
func (_ *GetMetadataConfigurationsFunction) Response() interface{} {
	return &GetMetadataConfigurationsResponse{}
}

type GetOSDFunction struct{}

func (_ *GetOSDFunction) Request() interface{} {
	return &GetOSD{}
}
func (_ *GetOSDFunction) Response() interface{} {
	return &GetOSDResponse{}
}

type GetOSDOptionsFunction struct{}

func (_ *GetOSDOptionsFunction) Request() interface{} {
	return &GetOSDOptions{}
}
func (_ *GetOSDOptionsFunction) Response() interface{} {
	return &GetOSDOptionsResponse{}
}

type GetOSDsFunction struct{}

func (_ *GetOSDsFunction) Request() interface{} {
	return &GetOSDs{}
}
func (_ *GetOSDsFunction) Response() interface{} {
	return &GetOSDsResponse{}
}

type GetProfileFunction struct{}

func (_ *GetProfileFunction) Request() interface{} {
	return &GetProfile{}
}
func (_ *GetProfileFunction) Response() interface{} {
	return &GetProfileResponse{}
}

type GetProfilesFunction struct{}

func (_ *GetProfilesFunction) Request() interface{} {
	return &GetProfiles{}
}
func (_ *GetProfilesFunction) Response() interface{} {
	return &GetProfilesResponse{}
}

type GetServiceCapabilitiesFunction struct{}

func (_ *GetServiceCapabilitiesFunction) Request() interface{} {
	return &GetServiceCapabilities{}
}
func (_ *GetServiceCapabilitiesFunction) Response() interface{} {
	return &GetServiceCapabilitiesResponse{}
}

type GetSnapshotUriFunction struct{}

func (_ *GetSnapshotUriFunction) Request() interface{} {
	return &GetSnapshotUri{}
}
func (_ *GetSnapshotUriFunction) Response() interface{} {
	return &GetSnapshotUriResponse{}
}

type GetStreamUriFunction struct{}

func (_ *GetStreamUriFunction) Request() interface{} {
	return &GetStreamUri{}
}
func (_ *GetStreamUriFunction) Response() interface{} {
	return &GetStreamUriResponse{}
}

type GetVideoAnalyticsConfigurationFunction struct{}

func (_ *GetVideoAnalyticsConfigurationFunction) Request() interface{} {
	return &GetVideoAnalyticsConfiguration{}
}
func (_ *GetVideoAnalyticsConfigurationFunction) Response() interface{} {
	return &GetVideoAnalyticsConfigurationResponse{}
}

type GetVideoAnalyticsConfigurationsFunction struct{}

func (_ *GetVideoAnalyticsConfigurationsFunction) Request() interface{} {
	return &GetVideoAnalyticsConfigurations{}
}
func (_ *GetVideoAnalyticsConfigurationsFunction) Response() interface{} {
	return &GetVideoAnalyticsConfigurationsResponse{}
}

type GetVideoEncoderConfigurationFunction struct{}

func (_ *GetVideoEncoderConfigurationFunction) Request() interface{} {
	return &GetVideoEncoderConfiguration{}
}
func (_ *GetVideoEncoderConfigurationFunction) Response() interface{} {
	return &GetVideoEncoderConfigurationResponse{}
}

type GetVideoEncoderConfigurationOptionsFunction struct{}

func (_ *GetVideoEncoderConfigurationOptionsFunction) Request() interface{} {
	return &GetVideoEncoderConfigurationOptions{}
}
func (_ *GetVideoEncoderConfigurationOptionsFunction) Response() interface{} {
	return &GetVideoEncoderConfigurationOptionsResponse{}
}

type GetVideoEncoderConfigurationsFunction struct{}

func (_ *GetVideoEncoderConfigurationsFunction) Request() interface{} {
	return &GetVideoEncoderConfigurations{}
}
func (_ *GetVideoEncoderConfigurationsFunction) Response() interface{} {
	return &GetVideoEncoderConfigurationsResponse{}
}

type GetVideoSourceConfigurationFunction struct{}

func (_ *GetVideoSourceConfigurationFunction) Request() interface{} {
	return &GetVideoSourceConfiguration{}
}
func (_ *GetVideoSourceConfigurationFunction) Response() interface{} {
	return &GetVideoSourceConfigurationResponse{}
}

type GetVideoSourceConfigurationOptionsFunction struct{}

func (_ *GetVideoSourceConfigurationOptionsFunction) Request() interface{} {
	return &GetVideoSourceConfigurationOptions{}
}
func (_ *GetVideoSourceConfigurationOptionsFunction) Response() interface{} {
	return &GetVideoSourceConfigurationOptionsResponse{}
}

type GetVideoSourceConfigurationsFunction struct{}

func (_ *GetVideoSourceConfigurationsFunction) Request() interface{} {
	return &GetVideoSourceConfigurations{}
}
func (_ *GetVideoSourceConfigurationsFunction) Response() interface{} {
	return &GetVideoSourceConfigurationsResponse{}
}

type GetVideoSourceModesFunction struct{}

func (_ *GetVideoSourceModesFunction) Request() interface{} {
	return &GetVideoSourceModes{}
}
func (_ *GetVideoSourceModesFunction) Response() interface{} {
	return &GetVideoSourceModesResponse{}
}

type GetVideoSourcesFunction struct{}

func (_ *GetVideoSourcesFunction) Request() interface{} {
	return &GetVideoSources{}
}
func (_ *GetVideoSourcesFunction) Response() interface{} {
	return &GetVideoSourcesResponse{}
}

type RemoveAudioDecoderConfigurationFunction struct{}

func (_ *RemoveAudioDecoderConfigurationFunction) Request() interface{} {
	return &RemoveAudioDecoderConfiguration{}
}
func (_ *RemoveAudioDecoderConfigurationFunction) Response() interface{} {
	return &RemoveAudioDecoderConfigurationResponse{}
}

type RemoveAudioEncoderConfigurationFunction struct{}

func (_ *RemoveAudioEncoderConfigurationFunction) Request() interface{} {
	return &RemoveAudioEncoderConfiguration{}
}
func (_ *RemoveAudioEncoderConfigurationFunction) Response() interface{} {
	return &RemoveAudioEncoderConfigurationResponse{}
}

type RemoveAudioOutputConfigurationFunction struct{}

func (_ *RemoveAudioOutputConfigurationFunction) Request() interface{} {
	return &RemoveAudioOutputConfiguration{}
}
func (_ *RemoveAudioOutputConfigurationFunction) Response() interface{} {
	return &RemoveAudioOutputConfigurationResponse{}
}

type RemoveAudioSourceConfigurationFunction struct{}

func (_ *RemoveAudioSourceConfigurationFunction) Request() interface{} {
	return &RemoveAudioSourceConfiguration{}
}
func (_ *RemoveAudioSourceConfigurationFunction) Response() interface{} {
	return &RemoveAudioSourceConfigurationResponse{}
}

type RemoveMetadataConfigurationFunction struct{}

func (_ *RemoveMetadataConfigurationFunction) Request() interface{} {
	return &RemoveMetadataConfiguration{}
}
func (_ *RemoveMetadataConfigurationFunction) Response() interface{} {
	return &RemoveMetadataConfigurationResponse{}
}

type RemovePTZConfigurationFunction struct{}

func (_ *RemovePTZConfigurationFunction) Request() interface{} {
	return &RemovePTZConfiguration{}
}
func (_ *RemovePTZConfigurationFunction) Response() interface{} {
	return &RemovePTZConfigurationResponse{}
}

type RemoveVideoAnalyticsConfigurationFunction struct{}

func (_ *RemoveVideoAnalyticsConfigurationFunction) Request() interface{} {
	return &RemoveVideoAnalyticsConfiguration{}
}
func (_ *RemoveVideoAnalyticsConfigurationFunction) Response() interface{} {
	return &RemoveVideoAnalyticsConfigurationResponse{}
}

type RemoveVideoEncoderConfigurationFunction struct{}

func (_ *RemoveVideoEncoderConfigurationFunction) Request() interface{} {
	return &RemoveVideoEncoderConfiguration{}
}
func (_ *RemoveVideoEncoderConfigurationFunction) Response() interface{} {
	return &RemoveVideoEncoderConfigurationResponse{}
}

type RemoveVideoSourceConfigurationFunction struct{}

func (_ *RemoveVideoSourceConfigurationFunction) Request() interface{} {
	return &RemoveVideoSourceConfiguration{}
}
func (_ *RemoveVideoSourceConfigurationFunction) Response() interface{} {
	return &RemoveVideoSourceConfigurationResponse{}
}

type SetAudioDecoderConfigurationFunction struct{}

func (_ *SetAudioDecoderConfigurationFunction) Request() interface{} {
	return &SetAudioDecoderConfiguration{}
}
func (_ *SetAudioDecoderConfigurationFunction) Response() interface{} {
	return &SetAudioDecoderConfigurationResponse{}
}

type SetAudioEncoderConfigurationFunction struct{}

func (_ *SetAudioEncoderConfigurationFunction) Request() interface{} {
	return &SetAudioEncoderConfiguration{}
}
func (_ *SetAudioEncoderConfigurationFunction) Response() interface{} {
	return &SetAudioEncoderConfigurationResponse{}
}

type SetAudioOutputConfigurationFunction struct{}

func (_ *SetAudioOutputConfigurationFunction) Request() interface{} {
	return &SetAudioOutputConfiguration{}
}
func (_ *SetAudioOutputConfigurationFunction) Response() interface{} {
	return &SetAudioOutputConfigurationResponse{}
}

type SetAudioSourceConfigurationFunction struct{}

func (_ *SetAudioSourceConfigurationFunction) Request() interface{} {
	return &SetAudioSourceConfiguration{}
}
func (_ *SetAudioSourceConfigurationFunction) Response() interface{} {
	return &SetAudioSourceConfigurationResponse{}
}

type SetMetadataConfigurationFunction struct{}

func (_ *SetMetadataConfigurationFunction) Request() interface{} {
	return &SetMetadataConfiguration{}
}
func (_ *SetMetadataConfigurationFunction) Response() interface{} {
	return &SetMetadataConfigurationResponse{}
}

type SetOSDFunction struct{}

func (_ *SetOSDFunction) Request() interface{} {
	return &SetOSD{}
}
func (_ *SetOSDFunction) Response() interface{} {
	return &SetOSDResponse{}
}

type SetSynchronizationPointFunction struct{}

func (_ *SetSynchronizationPointFunction) Request() interface{} {
	return &SetSynchronizationPoint{}
}
func (_ *SetSynchronizationPointFunction) Response() interface{} {
	return &SetSynchronizationPointResponse{}
}

type SetVideoAnalyticsConfigurationFunction struct{}

func (_ *SetVideoAnalyticsConfigurationFunction) Request() interface{} {
	return &SetVideoAnalyticsConfiguration{}
}
func (_ *SetVideoAnalyticsConfigurationFunction) Response() interface{} {
	return &SetVideoAnalyticsConfigurationResponse{}
}

type SetVideoEncoderConfigurationFunction struct{}

func (_ *SetVideoEncoderConfigurationFunction) Request() interface{} {
	return &SetVideoEncoderConfiguration{}
}
func (_ *SetVideoEncoderConfigurationFunction) Response() interface{} {
	return &SetVideoEncoderConfigurationResponse{}
}

type SetVideoSourceConfigurationFunction struct{}

func (_ *SetVideoSourceConfigurationFunction) Request() interface{} {
	return &SetVideoSourceConfiguration{}
}
func (_ *SetVideoSourceConfigurationFunction) Response() interface{} {
	return &SetVideoSourceConfigurationResponse{}
}

type SetVideoSourceModeFunction struct{}

func (_ *SetVideoSourceModeFunction) Request() interface{} {
	return &SetVideoSourceMode{}
}
func (_ *SetVideoSourceModeFunction) Response() interface{} {
	return &SetVideoSourceModeResponse{}
}

type StartMulticastStreamingFunction struct{}

func (_ *StartMulticastStreamingFunction) Request() interface{} {
	return &StartMulticastStreaming{}
}
func (_ *StartMulticastStreamingFunction) Response() interface{} {
	return &StartMulticastStreamingResponse{}
}

type StopMulticastStreamingFunction struct{}

func (_ *StopMulticastStreamingFunction) Request() interface{} {
	return &StopMulticastStreaming{}
}
func (_ *StopMulticastStreamingFunction) Response() interface{} {
	return &StopMulticastStreamingResponse{}
}
