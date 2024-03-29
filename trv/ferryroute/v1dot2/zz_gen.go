// Package v1dot2 contains the type definitions for FerryRoute v1.2.
//
// All types have accessor methods to access fields which can be chained on nils.
// This makes it possible to easily drill down into deeply nested data.
package v1dot2

import (
	"encoding/xml"
	"strings"
	"time"

	schema "code.dny.dev/trafikinfo/internal/trv/ferryroute/v1dot2"
	"code.dny.dev/trafikinfo/trv"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

func ObjectType() trv.ObjectType {
	return trv.ObjectType{
		Kind:      "FerryRoute",
		Version:   "1.2",
		Namespace: "Ferry.TrafficInfo",
	}
}

type FerryRoute struct {
	data *schema.FerryRoute
}

func (x *FerryRoute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.FerryRoute{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

func (x *FerryRoute) Deleted() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Deleted
}

// SV: Referens till Deviation.Id i objektet Situation
// EN: Reference to Deviation.Id in the Situation object
func (x *FerryRoute) DeviationID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.DeviationId
}

func (x *FerryRoute) Geometry() *Geometry {
	if x.data == nil {
		return &Geometry{}
	}
	return &Geometry{data: x.data.Geometry}
}

// SV: Ledens id
// EN: Trail ID
func (x *FerryRoute) ID() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Ledens namn
// EN: Name of the trail
func (x *FerryRoute) Name() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Name
}

// SV: Förkortning av ledens namn
// EN: Abbreviation of the name of the trail
func (x *FerryRoute) Shortname() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Shortname
}

func (x *FerryRoute) Type() *Type {
	if x.data == nil {
		return &Type{}
	}
	return &Type{data: x.data.Type}
}

func (x *FerryRoute) Harbors() []Harbor {
	if len(x.data.Harbor) == 0 {
		return nil
	}
	data := []Harbor{}
	for _, mem := range x.data.Harbor {
		data = append(data, Harbor{data: &mem})
	}
	return data
}

func (x *FerryRoute) Timetables() []Timetable {
	if len(x.data.Timetable) == 0 {
		return nil
	}
	data := []Timetable{}
	for _, mem := range x.data.Timetable {
		data = append(data, Timetable{data: &mem})
	}
	return data
}

// EN: Specifies when the object is stored.
// SV: Anger när objektet är sparat.
func (x *FerryRoute) ModifiedTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.ModifiedTime
}

func (x *FerryRoute) CheckoutID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Checkoutid
}

func (x *FerryRoute) Checkouts() *int64 {
	if x.data == nil {
		return nil
	}
	return x.data.Checkouts
}

func (x *FerryRoute) LeasedUntil() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.Leaseduntil
}

func (x *FerryRoute) Acknowledged() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.Acknowledged
}

func (x *FerryRoute) StatusCode() *uint8 {
	if x.data == nil {
		return nil
	}
	return x.data.Statuscode
}

func (x *FerryRoute) ItemsAcknowledged() *int64 {
	if x.data == nil {
		return nil
	}
	return x.data.Itemsacknowledged
}

func (x *FerryRoute) ItemsLeft() *int64 {
	if x.data == nil {
		return nil
	}
	return x.data.Itemsleft
}

type Geometry struct {
	data *schema.Geometry
}

func (x *Geometry) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Geometry{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Geometrisk punkt i koordinatsystem SWEREF99TM
// EN: Geometric point in coordinate system SWEREF99TM
func (x *Geometry) SWEREF99TM() *string {
	if x.data == nil {
		return nil
	}
	return x.data.SWEREF99TM
}

// SV: Geometrisk punkt i koordinatsystem WGS84
// EN: Geometric point in coordinate system WGS84
func (x *Geometry) WGS84() *string {
	if x.data == nil {
		return nil
	}
	return x.data.WGS84
}

type Type struct {
	data *schema.Type
}

func (x *Type) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Type{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Ledtypens id
// EN: Led type ID
func (x *Type) ID() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Ledtypens namn
// EN: Name of the joint type
func (x *Type) Name() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Name
}

type Harbor struct {
	data *schema.Harbor
}

func (x *Harbor) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Harbor{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Hamnens id
// EN: Port ID
func (x *Harbor) ID() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Hamnens namn
// EN: Name of the port
func (x *Harbor) Name() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Name
}

// SV: Sorteringsordning
// EN: Sort order
func (x *Harbor) SortOrder() *int {
	if x.data == nil {
		return nil
	}
	return x.data.SortOrder
}

func (x *Harbor) StopType() *StopType {
	if x.data == nil {
		return &StopType{}
	}
	return &StopType{data: x.data.StopType}
}

type Timetable struct {
	data *schema.Timetable
}

func (x *Timetable) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Timetable{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Tidtabellens beskrivning
// EN: Description of the timetable
func (x *Timetable) Description() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Description
}

// SV: Tidtabellens prioritet, det kan finnas flera tidtabeller som är giltiga samma datum, den med högst prioritet gäller.
// EN: The priority of the timetable, there may be several timetables valid on the same date, the one with the highest priority applies.
func (x *Timetable) Priority() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Priority
}

func (x *Timetable) Valid() []Valid {
	if len(x.data.Valid) == 0 {
		return nil
	}
	data := []Valid{}
	for _, mem := range x.data.Valid {
		data = append(data, Valid{data: &mem})
	}
	return data
}

func (x *Timetable) Periods() []Period {
	if len(x.data.Period) == 0 {
		return nil
	}
	data := []Period{}
	for _, mem := range x.data.Period {
		data = append(data, Period{data: &mem})
	}
	return data
}

type StopType struct {
	data *schema.StopType
}

func (x *StopType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.StopType{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Stopptypens id
// EN: Stop type ID
func (x *StopType) ID() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Namn på stopptypen. Det finns tre olika "Avg", "Ank/Avg" &amp; "Ank"
// EN: Name of the stop type. There are three different "Avg", "Ank/Avg" &amp; "Ank"
func (x *StopType) Name() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Name
}

// SV: De med värdet true är synliga i tidtabellen
// EN: Those with a value of true are visible in the timetable
func (x *StopType) Visible() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Visible
}

type Valid struct {
	data *schema.Valid
}

func (x *Valid) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Valid{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Datumet då tidtabellen börjar gälla
// EN: The date on which the timetable takes effect
func (x *Valid) From() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.From
}

// SV: Datumet då tidtabellen slutar gälla
// EN: The date the timetable expires
func (x *Valid) To() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.To
}

type Period struct {
	data *schema.Period
}

func (x *Period) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Period{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Periodens namn
// EN: Name of the period
func (x *Period) Name() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Name
}

// SV: Periodens sorteringsordning
// EN: Sort order of the period
func (x *Period) SortOrder() *int {
	if x.data == nil {
		return nil
	}
	return x.data.SortOrder
}

func (x *Period) Weekdays() []Weekday {
	if len(x.data.Weekday) == 0 {
		return nil
	}
	data := []Weekday{}
	for _, mem := range x.data.Weekday {
		data = append(data, Weekday{data: &mem})
	}
	return data
}

func (x *Period) Schedules() []Schedule {
	if len(x.data.Schedule) == 0 {
		return nil
	}
	data := []Schedule{}
	for _, mem := range x.data.Schedule {
		data = append(data, Schedule{data: &mem})
	}
	return data
}

type Weekday struct {
	data *schema.Weekday
}

func (x *Weekday) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Weekday{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Veckodagar som ingår i perioden (dagens namn)
// EN: Days of the week included in the period (day name)
func (x *Weekday) Day() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Day
}

// SV: Veckodagar som ingår i perioden (dagens id)
// EN: Days of the week included in the period (day id)
func (x *Weekday) ID() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

type Schedule struct {
	data *schema.Schedule
}

func (x *Schedule) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Schedule{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

func (x *Schedule) Deviations() []Deviation {
	if len(x.data.Deviation) == 0 {
		return nil
	}
	data := []Deviation{}
	for _, mem := range x.data.Deviation {
		data = append(data, Deviation{data: &mem})
	}
	return data
}

// SV: Tidpunkt för händelse
// EN: Time of event
func (x *Schedule) Time() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Time
}

func (x *Schedule) Harbor() *Harbor {
	if x.data == nil {
		return &Harbor{}
	}
	return &Harbor{data: x.data.Harbor}
}

// SV: Sorteringsordning
// EN: Sort order
func (x *Schedule) SortOrder() *int {
	if x.data == nil {
		return nil
	}
	return x.data.SortOrder
}

func (x *Schedule) StopType() *StopType {
	if x.data == nil {
		return &StopType{}
	}
	return &StopType{data: x.data.StopType}
}

type Deviation struct {
	data *schema.Deviation
}

func (x *Deviation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Deviation{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Beskrivning av avvikelsen
// EN: Description of the deviation
func (x *Deviation) Description() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Description
}

func (x *Deviation) Type() *DeviationType {
	if x.data == nil {
		return &DeviationType{}
	}
	return &DeviationType{data: x.data.Type}
}

// SV: Avvikelsens id
// EN: The id of the deviation
func (x *Deviation) ID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Om avvikelsen gäller under en period så finns den en sträng fråndatum med formatet "mmdd"
// EN: If the deviation applies for a period, there is a string from date with the format "mmdd"
func (x *Deviation) FromDate() *string {
	if x.data == nil {
		return nil
	}
	return x.data.FromDate
}

// SV: Om avvikelsen gäller under en period så finns den en sträng slutdatum med formatet "mmdd"
// EN: If the deviation applies for a period, there is a string end date with the format "mmdd"
func (x *Deviation) ToDate() *string {
	if x.data == nil {
		return nil
	}
	return x.data.ToDate
}

// SV: Om avvikelsen gäller specifika datum läggs de till en sträng med formatet "mmdd, mmdd"
// EN: If the discrepancy applies to specific dates, they are added to a string with the format "mmdd, mmdd
func (x *Deviation) SpecDate() *string {
	if x.data == nil {
		return nil
	}
	return x.data.SpecDate
}

type DeviationType struct {
	data *schema.DeviationType
}

func (x *DeviationType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.DeviationType{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Avvikelsetypens id
// EN: Anomaly type ID
func (x *DeviationType) ID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Avvikelsetypens namn, det finns fyra: "Meddelande", "Kallelse", "Går" &amp; "Går ej"
// EN: The name of the deviation type, there are four: "Meddelande", "Kallelse", "Går" &amp; "Går ej"
func (x *DeviationType) Name() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Name
}

// Response can be used to decode the response from the API.
//
// For example:
//
//	package main
//
//	import (
//		 "encoding/xml"
//		 api "code.dny.dev/trafikinfo/trv/ferryroute/v1dot2"
//	)
//
//	func main() {
//		var res api.Response
//		err := xml.Unmarshal(data, &res)
//	}
type Response struct {
	XMLName xml.Name `xml:"RESPONSE"`
	Results []struct {
		Info struct {
			LastModified trv.LastModified `xml:"LASTMODIFIED"`
			LastChangeID string           `xml:"LASTCHANGEID"`
			EvalResult   []any            `xml:"EVALRESULT"`
			SSEURL       string           `xml:"SSEURL"`
		} `xml:"INFO"`
		Error *trv.APIError `xml:"ERROR"`
		Data  []FerryRoute  `xml:"FerryRoute"`
	} `xml:"RESULT"`
}

// HasErrors returns whether any of the results
// includes an error.
func (r Response) HasErrors() bool {
	if len(r.Results) == 0 {
		return false
	}

	res := false
	for _, rr := range r.Results {
		if rr.Error != nil {
			res = true
			break
		}
	}
	return res
}

// Errors returns a slice of [trv.Error], if any error was
// included in the response.
func (r Response) ErrorMsg() string {
	if !r.HasErrors() {
		return ""
	}

	res := []string{}
	for _, rr := range r.Results {
		if rr.Error != nil {
			res = append(res, rr.Error.Error())
		}
	}
	return strings.Join(res, ", ")
}
