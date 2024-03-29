// Package v2 contains the type definitions for WeatherMeasurepoint v2.
package v2

import (
	"time"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

type WeatherMeasurepoint struct {
	// SV: Unik identitet för en mätpunkt
	Id *string `xml:"Id,omitempty"`
	// SV: Mätpunktens namn
	Name        *string      `xml:"Name,omitempty"`
	Geometry    *Geometry    `xml:"Geometry,omitempty"`
	Observation *Observation `xml:"Observation,omitempty"`
	// SV: Tidpunkt då dataposten ändrades i cachen
	// EN: Time when the data item was changed in the cache
	ModifiedTime *time.Time `xml:"ModifiedTime,omitempty"`
	// SV: Anger att dataposten raderats
	// EN: Indicates that the data record has been deleted
	Deleted *bool `xml:"Deleted,omitempty"`
}

type Geometry struct {
	// SV: Geometrisk punkt som observationen avser i koordinatsystem SWEREF99TM
	SWEREF99TM *string `xml:"SWEREF99TM,omitempty"`
	// SV: Geometrisk punkt som observationen avser i koordinatsystem WGS84
	WGS84 *string `xml:"WGS84,omitempty"`
}

type Observation struct {
	// SV: Tidpunkt som observationen avser, inklusive tidzon för att hantera sommartid och normaltid.
	Sample              *time.Time           `xml:"Sample,omitempty"`
	Weather             *Weather             `xml:"Weather,omitempty"`
	Surface             *SurfaceCondition    `xml:"Surface,omitempty"`
	Air                 *AirCondition        `xml:"Air,omitempty"`
	Wind                []WindCondition      `xml:"Wind,omitempty"`
	DeicingChemical     *DeicingChemical     `xml:"DeicingChemical,omitempty"`
	Subsurface          *SubsurfaceCondition `xml:"Subsurface,omitempty"`
	Aggregated5minutes  *Aggregated          `xml:"Aggregated5minutes,omitempty"`
	Aggregated10minutes *Aggregated          `xml:"Aggregated10minutes,omitempty"`
	Aggregated30minutes *Aggregated          `xml:"Aggregated30minutes,omitempty"`
	// SV: Unik publiceringsidentitet för observationen.
	Id *string `xml:"Id,omitempty"`
}

type Weather struct {
	// SV: Vilken typ av nederbörd som detekterats
	Precipitation *string `xml:"Precipitation,omitempty"`
}

type SurfaceCondition struct {
	// SV: Vägytans temperatur. Value [C]
	Temperature *Celsius `xml:"Temperature,omitempty"`
	// SV: Förekomst av vatten på vägytan.
	Water *bool `xml:"Water,omitempty"`
	// SV: Förekomst av is på vägytan.
	Ice *bool `xml:"Ice,omitempty"`
	// SV: Förekomst av snö på vägytan.
	Snow *bool `xml:"Snow,omitempty"`
	// SV: Value [0-1]
	Grip *Grip `xml:"Grip,omitempty"`
	// SV: Vattendjup på vägytan. Value [mm]
	WaterDepth *Millimeters `xml:"WaterDepth,omitempty"`
	// SV: Isdjup på vägytan. Value [mm]
	IceDepth *Millimeters `xml:"IceDepth,omitempty"`
	// SV: Snödjup på vägytan.
	SnowDepth *Snow `xml:"SnowDepth,omitempty"`
}

type AirCondition struct {
	// SV: Lufttemperatur. Value [C]
	Temperature *Celsius `xml:"Temperature,omitempty"`
	// SV: Daggpunkt, den temperatur där vatten kondenserar. Value [C]
	Dewpoint *Celsius `xml:"Dewpoint,omitempty"`
	// SV: Relativ luftfuktighet. Andel av den fukt som luften kan bära. Vid 100% är luften mättad. Value [%]
	RelativeHumidity *Percentage `xml:"RelativeHumidity,omitempty"`
	// SV: Den sträcka det finns sikt. Value [m]
	VisibleDistance *Meters `xml:"VisibleDistance,omitempty"`
	// SV: Förekomst av tillräckligt klart väder för att vara av vikt för vägytans temperatur.
	CloudFree *bool `xml:"CloudFree,omitempty"`
}

type WindCondition struct {
	Height *int `xml:"Height,omitempty"`
	// SV: Mått på vindhastighet vid en viss tidpunkt. Medelvärde över tiominutersperiod t.o.m. tidpunkten. Value [m/s]
	Speed *MetersPerSecond `xml:"Speed,omitempty"`
	// SV: Mått på vindriktning vid en viss tidpunkt. Medelvärde över tiominutersperiod t.o.m. tidpunkten. Value [grader]
	Direction *Degrees `xml:"Direction,omitempty"`
}

type DeicingChemical struct {
	// SV: Mängd salt vid mätpunkten. Value [g/kvm]
	Amount *GramPerSquareMeter `xml:"Amount,omitempty"`
}

type SubsurfaceCondition struct {
	Ground []GroundCondition `xml:"Ground,omitempty"`
}

type Celsius struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string  `xml:"SensorNames,omitempty"`
	Value       *float64 `xml:"Value,omitempty"`
}

type Grip struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string  `xml:"SensorNames,omitempty"`
	Value       *float64 `xml:"Value,omitempty"`
}

type Millimeters struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string  `xml:"SensorNames,omitempty"`
	Value       *float64 `xml:"Value,omitempty"`
}

type Snow struct {
	// SV: Mängd snö. Value [mm]
	Solid *Millimeters `xml:"Solid,omitempty"`
	// SV: Mängd vatten som snön motsvarar i smält form. Value [mm]
	WaterEquivalent *Millimeters `xml:"WaterEquivalent,omitempty"`
}

type Percentage struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string  `xml:"SensorNames,omitempty"`
	Value       *float64 `xml:"Value,omitempty"`
}

type Meters struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string  `xml:"SensorNames,omitempty"`
	Value       *float64 `xml:"Value,omitempty"`
}

type MetersPerSecond struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string  `xml:"SensorNames,omitempty"`
	Value       *float64 `xml:"Value,omitempty"`
}

type Degrees struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string `xml:"SensorNames,omitempty"`
	Value       *int    `xml:"Value,omitempty"`
}

type GramPerSquareMeter struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string  `xml:"SensorNames,omitempty"`
	Value       *float64 `xml:"Value,omitempty"`
}

type GroundCondition struct {
	// SV: Det djup som observationen avser och är negativt för att återspegla att det avser under markytan. Value [cm]
	Depth *IntegerCentimeters `xml:"Depth,omitempty"`
	// SV: Temperatur (grader celsius) vid ett visst djup i marken. Value [C]
	Temperature *Celsius `xml:"Temperature,omitempty"`
}

type Aggregated struct {
	Wind          *WindConditionAggregated          `xml:"Wind,omitempty"`
	Precipitation *PrecipitationConditionAggregated `xml:"Precipitation,omitempty"`
}

type IntegerCentimeters struct {
	// SV: Mätvärdets ursprung
	Origin *string `xml:"Origin,omitempty"`
	// SV: Sensorernas beteckning
	SensorNames *string `xml:"SensorNames,omitempty"`
	Value       *int    `xml:"Value,omitempty"`
}

type WindConditionAggregated struct {
	Height *int `xml:"Height,omitempty"`
	// SV: Högst uppmätt 3-sekundersmedelvärde under perioden. Value [m/s]
	SpeedMax *MetersPerSecond `xml:"SpeedMax,omitempty"`
	// SV: Value [m/s]
	SpeedAverage *MetersPerSecond `xml:"SpeedAverage,omitempty"`
}

type PrecipitationConditionAggregated struct {
	// SV: Förekomst av regn.
	Rain *bool `xml:"Rain,omitempty"`
	// SV: Förekomst av snö.
	Snow *bool `xml:"Snow,omitempty"`
	// SV: Mängd regn under perioden. Value [mm]
	RainSum *Millimeters `xml:"RainSum,omitempty"`
	// SV: Mängd snö under perioden.
	SnowSum *Snow `xml:"SnowSum,omitempty"`
	// SV: Mängd vatten som nederbörden under perioden motsvarar. Value [mm]
	TotalWaterEquivalent *Millimeters `xml:"TotalWaterEquivalent,omitempty"`
}
