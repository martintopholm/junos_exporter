package routingengine

import (
	"regexp"
)

type RoutingEngineRpc struct {
	Information struct {
		RouteEngine []RouteEngine `xml:"route-engine"`
	} `xml:"route-engine-information"`
}

type RouteEngine struct {
	Slot              string                 `xml:"slot"`
	Status            string                 `xml:"status"`
	Temperature       RouteEngineTemperature `xml:"temperature"`
	MemoryUtilization float64                `xml:"memory-buffer-utilization"`
	CPUTemperature    RouteEngineTemperature `xml:"cpu-temperature"`
	CPUUser           float64                `xml:"cpu-user"`
	CPUBackground     float64                `xml:"cpu-background"`
	CPUSystem         float64                `xml:"cpu-system"`
	CPUInterrupt      float64                `xml:"cpu-interrupt"`
	CPUIdle           float64                `xml:"cpu-idle"`
	UpTime            struct {
		Seconds uint64 `xml:"seconds,attr"`
	} `xml:"up-time"`
	LoadAverageOne     float64 `xml:"load-average-one"`
	LoadAverageFive    float64 `xml:"load-average-five"`
	LoadAverageFifteen float64 `xml:"load-average-fifteen"`
}

type RouteEngineTemperature struct {
	Value float64 `xml:"celsius,attr"`
}

type VersionRpc struct {
	MultiEngineResults struct {
		Items []struct {
			Slot         string      `xml:"re-name"`
			SoftwareInfo VersionInfo `xml:"software-information"`
		} `xml:"multi-routing-engine-item"`
	} `xml:"multi-routing-engine-results"`
	SoftwareInfo VersionInfo `xml:"software-information"`
}

type VersionInfo struct {
	Model        string `xml:"product-model"`
	JunosVersion string `xml:"junos-version"`
	PackageInfo  []struct {
		Name    string `xml:"name"`
		Comment string `xml:"comment"`
	} `xml:"package-information"`
}

var reVersionComment = regexp.MustCompile(`[^[]*\[([^]]+)\]$`)

func (v VersionInfo) Version() string {
	if v.JunosVersion != "" {
		return v.JunosVersion
	}
	for _, x := range v.PackageInfo {
		if x.Name != "jbase" {
			continue
		}
		parts := reVersionComment.FindStringSubmatch(x.Comment)
		if len(parts) > 1 {
			return parts[1]
		}
		break
	}
	return ""
}
