// Package v1dot2 contains the type definitions for FerryRoute v1.2.
package v1dot2

import (
	"time"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

type FerryRoute struct {
	Deleted *bool `xml:"Deleted,omitempty"`
	// SV: Referens till Deviation.Id i objektet Situation
	// EN: Reference to Deviation.Id in the Situation object
	DeviationId *string   `xml:"DeviationId,omitempty"`
	Geometry    *Geometry `xml:"Geometry,omitempty"`
	// SV: Ledens id
	// EN: Trail ID
	Id *int `xml:"Id,omitempty"`
	// SV: Ledens namn
	// EN: Name of the trail
	Name *string `xml:"Name,omitempty"`
	// SV: Förkortning av ledens namn
	// EN: Abbreviation of the name of the trail
	Shortname *string     `xml:"Shortname,omitempty"`
	Type      *Type       `xml:"Type,omitempty"`
	Harbor    []Harbor    `xml:"Harbor"`
	Timetable []Timetable `xml:"Timetable"`
	// EN: Specifies when the object is stored.
	// SV: Anger när objektet är sparat.
	ModifiedTime      *time.Time `xml:"ModifiedTime,omitempty"`
	Checkoutid        *string    `xml:"checkoutid,attr,omitempty"`
	Checkouts         *int64     `xml:"checkouts,attr,omitempty"`
	Leaseduntil       *time.Time `xml:"leaseduntil,attr,omitempty"`
	Acknowledged      *time.Time `xml:"acknowledged,attr,omitempty"`
	Statuscode        *uint8     `xml:"statuscode,attr,omitempty"`
	Itemsacknowledged *int64     `xml:"itemsacknowledged,attr,omitempty"`
	Itemsleft         *int64     `xml:"itemsleft,attr,omitempty"`
}

type Geometry struct {
	// SV: Geometrisk punkt i koordinatsystem SWEREF99TM
	// EN: Geometric point in coordinate system SWEREF99TM
	SWEREF99TM *string `xml:"SWEREF99TM,omitempty"`
	// SV: Geometrisk punkt i koordinatsystem WGS84
	// EN: Geometric point in coordinate system WGS84
	WGS84 *string `xml:"WGS84,omitempty"`
}

type Type struct {
	// SV: Ledtypens id
	// EN: Led type ID
	Id *int `xml:"Id,omitempty"`
	// SV: Ledtypens namn
	// EN: Name of the joint type
	Name *string `xml:"Name,omitempty"`
}

type Harbor struct {
	// SV: Hamnens id
	// EN: Port ID
	Id *int `xml:"Id,omitempty"`
	// SV: Hamnens namn
	// EN: Name of the port
	Name *string `xml:"Name,omitempty"`
	// SV: Sorteringsordning
	// EN: Sort order
	SortOrder *int      `xml:"SortOrder,omitempty"`
	StopType  *StopType `xml:"StopType,omitempty"`
}

type Timetable struct {
	// SV: Tidtabellens beskrivning
	// EN: Description of the timetable
	Description *string `xml:"Description,omitempty"`
	// SV: Tidtabellens prioritet, det kan finnas flera tidtabeller som är giltiga samma datum, den med högst prioritet gäller.
	// EN: The priority of the timetable, there may be several timetables valid on the same date, the one with the highest priority applies.
	Priority *int     `xml:"Priority,omitempty"`
	Valid    []Valid  `xml:"Valid"`
	Period   []Period `xml:"Period"`
}

type StopType struct {
	// SV: Stopptypens id
	// EN: Stop type ID
	Id *int `xml:"Id,omitempty"`
	// SV: Namn på stopptypen. Det finns tre olika "Avg", "Ank/Avg" &amp; "Ank"
	// EN: Name of the stop type. There are three different "Avg", "Ank/Avg" &amp; "Ank"
	Name *string `xml:"Name,omitempty"`
	// SV: De med värdet true är synliga i tidtabellen
	// EN: Those with a value of true are visible in the timetable
	Visible *bool `xml:"Visible,omitempty"`
}

type Valid struct {
	// SV: Datumet då tidtabellen börjar gälla
	// EN: The date on which the timetable takes effect
	From *time.Time `xml:"From,omitempty"`
	// SV: Datumet då tidtabellen slutar gälla
	// EN: The date the timetable expires
	To *time.Time `xml:"To,omitempty"`
}

type Period struct {
	// SV: Periodens namn
	// EN: Name of the period
	Name *string `xml:"Name,omitempty"`
	// SV: Periodens sorteringsordning
	// EN: Sort order of the period
	SortOrder *int       `xml:"SortOrder,omitempty"`
	Weekday   []Weekday  `xml:"Weekday"`
	Schedule  []Schedule `xml:"Schedule"`
}

type Weekday struct {
	// SV: Veckodagar som ingår i perioden (dagens namn)
	// EN: Days of the week included in the period (day name)
	Day *string `xml:"Day,omitempty"`
	// SV: Veckodagar som ingår i perioden (dagens id)
	// EN: Days of the week included in the period (day id)
	Id *int `xml:"Id,omitempty"`
}

type Schedule struct {
	Deviation []Deviation `xml:"Deviation"`
	// SV: Tidpunkt för händelse
	// EN: Time of event
	Time   *string `xml:"Time,omitempty"`
	Harbor *Harbor `xml:"Harbor,omitempty"`
	// SV: Sorteringsordning
	// EN: Sort order
	SortOrder *int      `xml:"SortOrder,omitempty"`
	StopType  *StopType `xml:"StopType,omitempty"`
}

type Deviation struct {
	// SV: Beskrivning av avvikelsen
	// EN: Description of the deviation
	Description *string        `xml:"Description,omitempty"`
	Type        *DeviationType `xml:"Type,omitempty"`
	// SV: Avvikelsens id
	// EN: The id of the deviation
	Id *string `xml:"Id,omitempty"`
	// SV: Om avvikelsen gäller under en period så finns den en sträng fråndatum med formatet "mmdd"
	// EN: If the deviation applies for a period, there is a string from date with the format "mmdd"
	FromDate *string `xml:"FromDate,omitempty"`
	// SV: Om avvikelsen gäller under en period så finns den en sträng slutdatum med formatet "mmdd"
	// EN: If the deviation applies for a period, there is a string end date with the format "mmdd"
	ToDate *string `xml:"ToDate,omitempty"`
	// SV: Om avvikelsen gäller specifika datum läggs de till en sträng med formatet "mmdd, mmdd"
	// EN: If the discrepancy applies to specific dates, they are added to a string with the format "mmdd, mmdd
	SpecDate *string `xml:"SpecDate,omitempty"`
}

type DeviationType struct {
	// SV: Avvikelsetypens id
	// EN: Anomaly type ID
	Id *string `xml:"Id,omitempty"`
	// SV: Avvikelsetypens namn, det finns fyra: "Meddelande", "Kallelse", "Går" &amp; "Går ej"
	// EN: The name of the deviation type, there are four: "Meddelande", "Kallelse", "Går" &amp; "Går ej"
	Name *string `xml:"Name,omitempty"`
}
