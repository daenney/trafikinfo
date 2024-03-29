// Package v1 contains the type definitions for MeasurementData100 v1.
//
// All types have accessor methods to access fields which can be chained on nils.
// This makes it possible to easily drill down into deeply nested data.
package v1

import (
	"encoding/xml"
	"strings"
	"time"

	schema "code.dny.dev/trafikinfo/internal/trv/measurementdata100/v1"
	"code.dny.dev/trafikinfo/trv"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

func ObjectType() trv.ObjectType {
	return trv.ObjectType{
		Kind:      "MeasurementData100",
		Version:   "1",
		Namespace: "",
	}
}

type MeasurementData100 struct {
	data *schema.MeasurementData100
}

func (x *MeasurementData100) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.MeasurementData100{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Län
func (x *MeasurementData100) County() *int {
	if x.data == nil {
		return nil
	}
	return x.data.County
}

// SV: Huvudvägnummer
func (x *MeasurementData100) RoadMainNumber() *int {
	if x.data == nil {
		return nil
	}
	return x.data.RoadMainNumber
}

// SV: Undervägnummer. Kallas ibland även punktväg. Måste användas ihop med huvudvägnumret.
func (x *MeasurementData100) RoadSubNumber() *int {
	if x.data == nil {
		return nil
	}
	return x.data.RoadSubNumber
}

func (x *MeasurementData100) Direction() *Direction {
	if x.data == nil {
		return &Direction{}
	}
	return &Direction{data: x.data.Direction}
}

// SV: Körfält. Räknas från höger sida och startar med körfält 10. Nästa är 20 och kan gå upp till 50.
func (x *MeasurementData100) Lane() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Lane
}

// SV: Start löpande längd. Anges i meter. Starten för aktuell åtgärdsinformation angivet i antal meter från vägnumrets start inom det länet, i aktuell riktning.
func (x *MeasurementData100) StartContinuousLength() *int {
	if x.data == nil {
		return nil
	}
	return x.data.StartContinuousLength
}

// SV: Slut löpnande längd. Anges i meter.Slutet för aktuell åtgärdsinformation angivet i antal meter från vägnumrets start inom det länet, i aktuell riktning.
func (x *MeasurementData100) EndContinuousLength() *int {
	if x.data == nil {
		return nil
	}
	return x.data.EndContinuousLength
}

// SV: Längd i antal meter
func (x *MeasurementData100) Length() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Length
}

// SV: Mätdatum. Ett gemensamt datum för en hel mätperiod. Är detta datum som visas i graferna i PMSV3 Analysera sträcka.
func (x *MeasurementData100) MeasurementDate() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.MeasurementDate
}

// SV: Mätdatum exakt. Det exakta mätdatumet för aktuellt mätdata.
func (x *MeasurementData100) MeasurementDateSpecific() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.MeasurementDateSpecific
}

// SV: Mätdatatyp
func (x *MeasurementData100) MeasurementDataType() *MeasurementDataType {
	if x.data == nil {
		return &MeasurementDataType{}
	}
	return &MeasurementDataType{data: x.data.MeasurementDataType}
}

// SV: IRI höger medelvärde. Avser jämnhet i längsled. IRI (International Roughness Index). Beräknat mått baserat på uppmätt längsprofil i höger hjulspår.
func (x *MeasurementData100) IRIRightAverageValue() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.IRIRightAverageValue
}

// SV: Kantdjup medelvärde. Kantdjup avser att detektera skador närmast vägrenen. Mäts inte direkt av mätbil utan beräknas i efterbearbetning baserat på mätta 20m data om tvärprofil.
func (x *MeasurementData100) EdgeDepthAverageValue() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.EdgeDepthAverageValue
}

// SV: Spårdjup max15 medelvärde. 15 lasrar.
func (x *MeasurementData100) RutDepthMax15AverageValue() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.RutDepthMax15AverageValue
}

// SV: Spårdjup max17 medelvärde. 17 lasrar. Medelvärdet av de största spårdjupen beräknade enligt 'trådprincipen' för 200 profiler inom 20m sträckan. Se TRV metodbeskrivningar för vägytemätning.
func (x *MeasurementData100) RutDepthMax17AverageValue() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.RutDepthMax17AverageValue
}

// SV: MPD Makrotextur medelvärde. beskriver vägytans 'skrovlighet' i våglängdsintervallet 0,5-50 mm. Makrotextur beräknad som MPD. (Mean Profile Depth beräknat enligt ISO 13473-1)
func (x *MeasurementData100) MPDMacrotextureAverageValue() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.MPDMacrotextureAverageValue
}

// SV: Datum för när mätdata hämtades ut från källsystemen och 100m-data skapades.
func (x *MeasurementData100) TimeStamp() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.TimeStamp
}

// SV: Tidpunkt då dataposten ändrades
// EN: Time when the data item was changed
func (x *MeasurementData100) ModifiedTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.ModifiedTime
}

// SV: Anger att dataposten raderats
// EN: Indicates that the data record has been deleted
func (x *MeasurementData100) Deleted() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Deleted
}

type Direction struct {
	data *schema.Direction
}

func (x *Direction) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Direction{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Riktning. Värde.
func (x *Direction) Code() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Code
}

// SV: Riktning. Beskrivning.
func (x *Direction) Value() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Value
}

type MeasurementDataType struct {
	data *schema.MeasurementDataType
}

func (x *MeasurementDataType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.MeasurementDataType{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Mätdatatyp. Värde.
func (x *MeasurementDataType) Code() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Code
}

// SV: Mätdatatyp. Beskrivning.
func (x *MeasurementDataType) Value() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Value
}

// Response can be used to decode the response from the API.
//
// For example:
//
//	package main
//
//	import (
//		 "encoding/xml"
//		 api "code.dny.dev/trafikinfo/trv/measurementdata100/v1"
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
		Error *trv.APIError        `xml:"ERROR"`
		Data  []MeasurementData100 `xml:"MeasurementData100"`
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
