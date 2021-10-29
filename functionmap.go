package onvif

import (
	"github.com/IOTechSystems/onvif/device"
	"github.com/IOTechSystems/onvif/media"
)

var DeviceFunctionMap = map[string]Function{
	// Network Configuration
	GetHostname:              &device.GetHostnameFunction{},
	SetHostname:              &device.SetHostnameFunction{},
	GetDNS:                   &device.GetDNSFunction{},
	SetDNS:                   &device.SetDNSFunction{},
	GetNetworkInterfaces:     &device.GetNetworkInterfacesFunction{},
	SetNetworkInterfaces:     &device.SetNetworkInterfacesFunction{},
	GetNetworkProtocols:      &device.GetNetworkProtocolsFunction{},
	SetNetworkProtocols:      &device.SetNetworkProtocolsFunction{},
	GetNetworkDefaultGateway: &device.GetNetworkDefaultGatewayFunction{},
	SetNetworkDefaultGateway: &device.SetNetworkDefaultGatewayFunction{},
	// System Function
	GetDeviceInformation:    &device.GetDeviceInformationFunction{},
	GetSystemDateAndTime:    &device.GetSystemDateAndTimeFunction{},
	SetSystemDateAndTime:    &device.SetSystemDateAndTimeFunction{},
	SetSystemFactoryDefault: &device.SetSystemFactoryDefaultFunction{},
	SystemReboot:            &device.SystemRebootFunction{},
}

var MediaFunctionMap = map[string]Function{
	// Metadata Configuration
	GetMetadataConfiguration:            &media.GetMetadataConfigurationFunction{},
	GetMetadataConfigurations:           &media.GetMetadataConfigurationsFunction{},
	AddMetadataConfiguration:            &media.AddMetadataConfigurationFunction{},
	RemoveMetadataConfiguration:         &media.RemoveMetadataConfigurationFunction{},
	SetMetadataConfiguration:            &media.SetMetadataConfigurationFunction{},
	GetCompatibleMetadataConfigurations: &media.GetCompatibleMetadataConfigurationsFunction{},
	GetMetadataConfigurationOptions:     &media.GetMetadataConfigurationOptionsFunction{},
	GetProfiles:                         &media.GetProfilesFunction{},
	GetStreamUri:                        &media.GetStreamUriFunction{},
	GetVideoEncoderConfiguration:        &media.GetVideoEncoderConfigurationFunction{},
	SetVideoEncoderConfiguration:        &media.SetVideoEncoderConfigurationFunction{},
	GetVideoEncoderConfigurationOptions: &media.GetVideoEncoderConfigurationOptionsFunction{},
}