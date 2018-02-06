package rpc

type RoutingEngineRpc struct {
	Information struct {
		RouteEngine RouteEngine `xml:"route-engine"`
	} `xml:"route-engine-information"`
}

type RouteEngine struct {
	Temperature        RouteEngineTemperature `xml:"temperature"`
	CPUTemperature     RouteEngineTemperature `xml:"cpu-temperature"`
	CPUUser            float64                `xml:"cpu-user"`
	CPUBackground      float64                `xml:"cpu-background"`
	CPUSystem          float64                `xml:"cpu-system"`
	CPUInterrupt       float64                `xml:"cpu-interrupt"`
	CPUIdle            float64                `xml:"cpu-idle"`
	LoadAverageOne     float64                `xml:"load-average-one"`
	LoadAAverageFive   float64                `xml:"load-average-five"`
	LoadAverageFifteen float64                `xml:"load-average-fifteen"`
}

type RouteEngineTemperature struct {
	Value float64 `xml:"celsius,attr"`
}
