// Package v1dot1 contains the type definitions for TrainPosition v1.1.
package v1dot1

import (
	"time"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

type TrainPosition struct {
	// SV: Tåginformation
	// EN: Train information
	Train *Train `xml:"Train,omitempty"`
	// SV: Senast registrerade position för tåget
	// EN: Last known position of the train
	Position *Position `xml:"Position,omitempty"`
	// SV: Tiden då positionen uppmättes
	// EN: The time when the position was measured
	TimeStamp *time.Time `xml:"TimeStamp,omitempty"`
	// SV: Tågets aktuell status
	// EN: The train's current status
	Status *Status `xml:"Status,omitempty"`
	// SV: Tågets bäring i grader
	// EN: The train's bearing in degrees
	Bearing *int `xml:"Bearing,omitempty"`
	// SV: Tågets hastighet i kilometer per timme
	// EN: The train's speed in kilometers per hour
	Speed *int `xml:"Speed,omitempty"`
	// SV: Versionsnumret för ett tågs position
	// EN: The version number for a train's position
	VersionNumber *int64 `xml:"VersionNumber,omitempty"`
	// EN: Specifies when the object is stored.
	// SV: Anger när objektet är sparat.
	ModifiedTime *time.Time `xml:"ModifiedTime,omitempty"`
	Deleted      *bool      `xml:"Deleted,omitempty"`
}

type Train struct {
	// SV: Operativt tågnummer
	// EN: Operational train number
	OperationalTrainNumber *string `xml:"OperationalTrainNumber,omitempty"`
	// SV: Det operativa tågets utgångsdag
	// EN: The operational train's departure day
	OperationalTrainDepartureDate *time.Time `xml:"OperationalTrainDepartureDate,omitempty"`
	// SV: Tågets uppdragsnummer
	// EN: The train's daily journey plan number
	JourneyPlanNumber *string `xml:"JourneyPlanNumber,omitempty"`
	// SV: Tåguppdragets utgångsdag
	// EN: Daily journey plan departure day
	JourneyPlanDepartureDate *time.Time `xml:"JourneyPlanDepartureDate,omitempty"`
	// SV: Tågets annonserade tågnummer (tågnumret som står på biljetten)
	// EN: The train's advertised train number
	AdvertisedTrainNumber *string `xml:"AdvertisedTrainNumber,omitempty"`
}

type Position struct {
	// SV: Koordinater uttryckt i formatet sweref99tm
	// EN: Coordinates in sweref99tm format
	SWEREF99TM *string `xml:"SWEREF99TM,omitempty"`
	// SV: Koordinater uttryckt i formatet wgs84
	// EN: Coordinates in wgs84 format
	WGS84 *string `xml:"WGS84,omitempty"`
}

type Status struct {
	// SV: Flagga som visar om tåget är aktivt
	// EN: A flag that shows whether or not the train is active
	Active *bool `xml:"Active,omitempty"`
}