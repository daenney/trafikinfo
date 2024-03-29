// Package v1dot4 contains the type definitions for TrainStation v1.4.
//
// All types have accessor methods to access fields which can be chained on nils.
// This makes it possible to easily drill down into deeply nested data.
package v1dot4

import (
	"encoding/xml"
	"strings"
	"time"

	schema "code.dny.dev/trafikinfo/internal/trv/trainstation/v1dot4"
	"code.dny.dev/trafikinfo/trv"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

func ObjectType() trv.ObjectType {
	return trv.ObjectType{
		Kind:      "TrainStation",
		Version:   "1.4",
		Namespace: "rail.infrastructure",
	}
}

type TrainStation struct {
	data *schema.TrainStation
}

func (x *TrainStation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.TrainStation{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Anger om stationen annonseras i tidtabell
// EN: Indicates if the station is advertised in the timetable
func (x *TrainStation) Advertised() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Advertised
}

// SV: Stationens namn
// EN: Stations name
func (x *TrainStation) AdvertisedLocationName() *string {
	if x.data == nil {
		return nil
	}
	return x.data.AdvertisedLocationName
}

// SV: Stationens namn i kort version
// EN: Stations name in short version
func (x *TrainStation) AdvertisedShortLocationName() *string {
	if x.data == nil {
		return nil
	}
	return x.data.AdvertisedShortLocationName
}

// SV: Beteckning för i vilket land stationen finns.<br /> 'DE' - Tyskland<br /> 'DK' - Danmark<br /> 'NO' - Norge<br /> 'SE' - Sverige
// EN: Designation of the country in which the station is located.<br /> 'DE' - Germany<br /> 'DK' - Denmark<br /> 'NO' - Norway<br /> 'SE' - Sweden
func (x *TrainStation) CountryCode() *string {
	if x.data == nil {
		return nil
	}
	return x.data.CountryCode
}

// SV: <div class="toggleTitle">Länsnummer</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>1</td> <td>Stockholms län</td> </tr> <tr> <td>2</td> <td> DEPRECATED<br /> Användes tidigare för Stockholms län </td> </tr> <tr> <td>3</td> <td>Uppsala län</td> </tr> <tr> <td>4</td> <td>Södermanlands län</td> </tr> <tr> <td>5</td> <td>Östergötlands län</td> </tr> <tr> <td>6</td> <td>Jönköpings län</td> </tr> <tr> <td>7</td> <td>Kronobergs län</td> </tr> <tr> <td>8</td> <td>Kalmar län</td> </tr> <tr> <td>9</td> <td>Gotlands län</td> </tr> <tr> <td>10</td> <td>Blekinge län</td> </tr> <tr> <td>12</td> <td>Skåne län</td> </tr> <tr> <td>13</td> <td>Hallands län</td> </tr> <tr> <td>14</td> <td>Västra Götalands län</td> </tr> <tr> <td>17</td> <td>Värmlands län</td> </tr> <tr> <td>18</td> <td>Örebro län</td> </tr> <tr> <td>19</td> <td>Västmanlands län</td> </tr> <tr> <td>20</td> <td>Dalarnas län</td> </tr> <tr> <td>21</td> <td>Gävleborgs län</td> </tr> <tr> <td>22</td> <td>Västernorrlands län</td> </tr> <tr> <td>23</td> <td>Jämtlands län</td> </tr> <tr> <td>24</td> <td>Västerbottens län</td> </tr> <tr> <td>25</td> <td>Norrbottens län</td> </tr> </table> </div>
func (x *TrainStation) Counties() []int {
	if x.data == nil {
		return nil
	}
	return x.data.CountyNo
}

func (x *TrainStation) Deleted() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Deleted
}

func (x *TrainStation) Geometry() *Geometry {
	if x.data == nil {
		return &Geometry{}
	}
	return &Geometry{data: x.data.Geometry}
}

// SV: Upplysningsinformation för stationen, ex. "SL-tåg omfattas ej.", "Ring 033-172444 för trafikinformation"
// EN: Disclosure information for the station, ex. "SL-train is excluded.", "Call 033-172444 for trafficinfomation"
func (x *TrainStation) LocationInformationText() *string {
	if x.data == nil {
		return nil
	}
	return x.data.LocationInformationText
}

// SV: Stationens unika signatur, ex. "Cst"
// EN: Stations unique signature, ex. "Cst"
func (x *TrainStation) LocationSignature() *string {
	if x.data == nil {
		return nil
	}
	return x.data.LocationSignature
}

// SV: Plattformens spår
// EN: Platform Tracks
func (x *TrainStation) PlatformLines() []string {
	if x.data == nil {
		return nil
	}
	return x.data.PlatformLine
}

// SV: Anger om stationen prognostiseras i tidtabell
// EN: Specifies if station forecasted in timetabell
func (x *TrainStation) Prognosticated() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Prognosticated
}

// SV: Det av Transportstyrelsen fastslagna officiella namnet på stationen
// EN: The official name of the station established by the Swedish Transport Agency
func (x *TrainStation) OfficialLocationName() *string {
	if x.data == nil {
		return nil
	}
	return x.data.OfficialLocationName
}

// EN: Specifies when the object is stored.
// SV: Anger när objektet är sparat.
func (x *TrainStation) ModifiedTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.ModifiedTime
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

// Response can be used to decode the response from the API.
//
// For example:
//
//	package main
//
//	import (
//		 "encoding/xml"
//		 api "code.dny.dev/trafikinfo/trv/trainstation/v1dot4"
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
		Error *trv.APIError  `xml:"ERROR"`
		Data  []TrainStation `xml:"TrainStation"`
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
