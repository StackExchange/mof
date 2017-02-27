package mof

import (
	"testing"
)

func TestNil(t *testing.T) {
	target := &MSFT_DSCConfigurationStatus_Nil{}
	err := Unmarshal([]byte(snippetWithNil), target)
	if err != nil {
		t.Fatal(err)
	}
	if target.ReturnValue != 32 {
		t.Errorf("Return value %d != 32", target.ReturnValue)
	}
	rids := target.MSFT_DSCConfigurationStatus[0].ResourcesInDesiredState[0]
	if rids.ConfigurationName != "" {
		t.Fatal("Expected NULL to be empty string")
	}
}

type MSFT_DSCConfigurationStatus_Nil struct {
	MSFT_DSCConfigurationStatus []struct {
		DurationInSeconds       int
		Mode                    string
		NumberOfResources       int
		RebootRequested         bool
		ResourcesInDesiredState []struct {
			ConfigurationName string
			DurationInSeconds float64
		}
		StartDate string
		Status    string
		Type      string
	}
	ReturnValue int
}

const snippetWithNil = `class __PARAMETERS
{
	[Stream: DisableOverride ToSubClass, EmbeddedInstance("MSFT_DSCConfigurationStatus"): ToSubClass, Out, ID(1): DisableOverride ToInstance] MSFT_DSCConfigurationStatus configurationStatus[] = {
instance of MSFT_DSCConfigurationStatus
{
	DurationInSeconds = 261;
	HostName = "NY-WINBUILD03";
	IPV4Addresses = {"10.7.0.108", "127.0.0.1"};
	IPV6Addresses = {"fe80::24e3:b6d2:c6c2:2050%12", "::2000:0:0:0", "::1", "::2000:0:0:0"};
	JobID = "{C1B263FB-FD3A-11E6-80D0-0050569D6777}";
	LCMVersion = "2.0";
	Locale = "en-US";
	MACAddresses = {"00-50-56-9D-67-77", "00-00-00-00-00-00-00-E0"};
	MetaData = "Author: MSFT_DSCLocalConfigurationManager; Name: MSFT_DSCLocalConfigurationManager; Version: 2.0.0; GenerationDate: 11/23/2016 18:20:33; GenerationHost: GBRAY03;";
	Mode = "Pull";
	NumberOfResources = 23;
	RebootRequested = FALSE;
	ResourcesInDesiredState = {
instance of MSFT_ResourceInDesiredState
{
	ConfigurationName = NULL;
	DurationInSeconds = 0;
	InDesiredState = TRUE;
	InstanceName = "MonitorAllTheThings::[SEC_Base_Server]JustTheBasics";
	ModuleName = NULL;
	ModuleVersion = NULL;
	RebootRequested = FALSE;
	ResourceId = "[scollector]MonitorAllTheThings::[SEC_Base_Server]JustTheBasics";
	ResourceName = "scollector";
	SourceInfo = NULL;
	StartDate = "00000000000000.000000+000";
}};
	StartDate = "20170227221910.721000+000";
	Status = "Success";
	Type = "Consistency";
}};
	[out] uint32 ReturnValue = 32;
};`
