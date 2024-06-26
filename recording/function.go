// -*- Mode: Go; indent-tabs-mode: t -*-
//
// Copyright (C) 2022 Intel Corporation
//
// SPDX-License-Identifier: Apache-2.0

// Code generated by gen_commands.py DO NOT EDIT.

package recording

type CreateRecordingFunction struct{}

func (_ *CreateRecordingFunction) Request() interface{} {
	return &CreateRecording{}
}
func (_ *CreateRecordingFunction) Response() interface{} {
	return &CreateRecordingResponse{}
}

type CreateRecordingJobFunction struct{}

func (_ *CreateRecordingJobFunction) Request() interface{} {
	return &CreateRecordingJob{}
}
func (_ *CreateRecordingJobFunction) Response() interface{} {
	return &CreateRecordingJobResponse{}
}

type CreateTrackFunction struct{}

func (_ *CreateTrackFunction) Request() interface{} {
	return &CreateTrack{}
}
func (_ *CreateTrackFunction) Response() interface{} {
	return &CreateTrackResponse{}
}

type DeleteRecordingFunction struct{}

func (_ *DeleteRecordingFunction) Request() interface{} {
	return &DeleteRecording{}
}
func (_ *DeleteRecordingFunction) Response() interface{} {
	return &DeleteRecordingResponse{}
}

type DeleteRecordingJobFunction struct{}

func (_ *DeleteRecordingJobFunction) Request() interface{} {
	return &DeleteRecordingJob{}
}
func (_ *DeleteRecordingJobFunction) Response() interface{} {
	return &DeleteRecordingJobResponse{}
}

type DeleteTrackFunction struct{}

func (_ *DeleteTrackFunction) Request() interface{} {
	return &DeleteTrack{}
}
func (_ *DeleteTrackFunction) Response() interface{} {
	return &DeleteTrackResponse{}
}

type ExportRecordedDataFunction struct{}

func (_ *ExportRecordedDataFunction) Request() interface{} {
	return &ExportRecordedData{}
}
func (_ *ExportRecordedDataFunction) Response() interface{} {
	return &ExportRecordedDataResponse{}
}

type GetExportRecordedDataStateFunction struct{}

func (_ *GetExportRecordedDataStateFunction) Request() interface{} {
	return &GetExportRecordedDataState{}
}
func (_ *GetExportRecordedDataStateFunction) Response() interface{} {
	return &GetExportRecordedDataStateResponse{}
}

type GetRecordingConfigurationFunction struct{}

func (_ *GetRecordingConfigurationFunction) Request() interface{} {
	return &GetRecordingConfiguration{}
}
func (_ *GetRecordingConfigurationFunction) Response() interface{} {
	return &GetRecordingConfigurationResponse{}
}

type GetRecordingJobConfigurationFunction struct{}

func (_ *GetRecordingJobConfigurationFunction) Request() interface{} {
	return &GetRecordingJobConfiguration{}
}
func (_ *GetRecordingJobConfigurationFunction) Response() interface{} {
	return &GetRecordingJobConfigurationResponse{}
}

type GetRecordingJobStateFunction struct{}

func (_ *GetRecordingJobStateFunction) Request() interface{} {
	return &GetRecordingJobState{}
}
func (_ *GetRecordingJobStateFunction) Response() interface{} {
	return &GetRecordingJobStateResponse{}
}

type GetRecordingJobsFunction struct{}

func (_ *GetRecordingJobsFunction) Request() interface{} {
	return &GetRecordingJobs{}
}
func (_ *GetRecordingJobsFunction) Response() interface{} {
	return &GetRecordingJobsResponse{}
}

type GetRecordingOptionsFunction struct{}

func (_ *GetRecordingOptionsFunction) Request() interface{} {
	return &GetRecordingOptions{}
}
func (_ *GetRecordingOptionsFunction) Response() interface{} {
	return &GetRecordingOptionsResponse{}
}

type GetRecordingsFunction struct{}

func (_ *GetRecordingsFunction) Request() interface{} {
	return &GetRecordings{}
}
func (_ *GetRecordingsFunction) Response() interface{} {
	return &GetRecordingsResponse{}
}

type GetServiceCapabilitiesFunction struct{}

func (_ *GetServiceCapabilitiesFunction) Request() interface{} {
	return &GetServiceCapabilities{}
}
func (_ *GetServiceCapabilitiesFunction) Response() interface{} {
	return &GetServiceCapabilitiesResponse{}
}

type GetTrackConfigurationFunction struct{}

func (_ *GetTrackConfigurationFunction) Request() interface{} {
	return &GetTrackConfiguration{}
}
func (_ *GetTrackConfigurationFunction) Response() interface{} {
	return &GetTrackConfigurationResponse{}
}

type SetRecordingConfigurationFunction struct{}

func (_ *SetRecordingConfigurationFunction) Request() interface{} {
	return &SetRecordingConfiguration{}
}
func (_ *SetRecordingConfigurationFunction) Response() interface{} {
	return &SetRecordingConfigurationResponse{}
}

type SetRecordingJobConfigurationFunction struct{}

func (_ *SetRecordingJobConfigurationFunction) Request() interface{} {
	return &SetRecordingJobConfiguration{}
}
func (_ *SetRecordingJobConfigurationFunction) Response() interface{} {
	return &SetRecordingJobConfigurationResponse{}
}

type SetRecordingJobModeFunction struct{}

func (_ *SetRecordingJobModeFunction) Request() interface{} {
	return &SetRecordingJobMode{}
}
func (_ *SetRecordingJobModeFunction) Response() interface{} {
	return &SetRecordingJobModeResponse{}
}

type SetTrackConfigurationFunction struct{}

func (_ *SetTrackConfigurationFunction) Request() interface{} {
	return &SetTrackConfiguration{}
}
func (_ *SetTrackConfigurationFunction) Response() interface{} {
	return &SetTrackConfigurationResponse{}
}

type StopExportRecordedDataFunction struct{}

func (_ *StopExportRecordedDataFunction) Request() interface{} {
	return &StopExportRecordedData{}
}
func (_ *StopExportRecordedDataFunction) Response() interface{} {
	return &StopExportRecordedDataResponse{}
}
