package mof

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"testing"
)

func TestUnmarshal(t *testing.T) {
	b, err := ioutil.ReadFile("test.mof")
	if err != nil {
		t.Fatal(err)
	}
	var v interface{}
	if err := Unmarshal(b, &v); err != nil {
		t.Fatal(err)
	}
	var s MSFT_DSCConfigurationStatus
	if err := Unmarshal(b, &s); err != nil {
		t.Fatal(err)
	}
	vb, _ := json.MarshalIndent(&v, "", "  ")
	sb, _ := json.MarshalIndent(&s, "", "  ")
	if string(vb) != string(sb) {
		fmt.Println("interface:", string(vb))
		fmt.Println("struct:", string(sb))
		t.Fatalf("interface != struct")
	}
}

type MSFT_DSCConfigurationStatus struct {
	MSFT_DSCConfigurationStatus []struct {
		DurationInSeconds int
		Error             string
		HostName          string
		IPV4Addresses     []string
		IPV6Addresses     []string
		JobID             string
		LCMVersion        string
		Locale            string
		MACAddresses      []string
		MetaConfiguration struct {
			AllowModuleOverwrite           bool
			ConfigurationDownloadManagers  []interface{}
			ConfigurationMode              string
			ConfigurationModeFrequencyMins int
			DebugMode                      []string
			LCMCompatibleVersions          []string
			LCMState                       string
			LCMVersion                     string
			RebootNodeIfNeeded             bool
			RefreshFrequencyMins           int
			RefreshMode                    string
			ReportManagers                 []interface{}
			ResourceModuleManagers         []interface{}
			StatusRetentionTimeInDays      int
		}
		Mode            string
		RebootRequested bool
		StartDate       string
		Status          string
		Type            string
	}
	ReturnValue int
}
