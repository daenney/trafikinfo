// Package v1dot5 contains the type definitions for TravelTimeRoute v1.5.
//
// All types have accessor methods to access fields which can be chained on nils.
// This makes it possible to easily drill down into deeply nested data.
package v1dot5

import (
	"encoding/xml"
	"strings"
	"time"

	schema "code.dny.dev/trafikinfo/internal/trv/traveltimeroute/v1dot5"
	"code.dny.dev/trafikinfo/trv"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

func ObjectType() trv.ObjectType {
	return trv.ObjectType{
		Kind:      "TravelTimeRoute",
		Version:   "1.5",
		Namespace: "",
	}
}

type TravelTimeRoute struct {
	data *schema.TravelTimeRoute
}

func (x *TravelTimeRoute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.TravelTimeRoute{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Genomsnittlig funktionell vägklass för sträckan.<br /> En klassificering 0-9 baserad på hur viktig en väg är för det totala vägnätets förbindelsemöjligheter.<br /> 0 - De viktigaste vägarna<br /> 9 - De minst viktiga vägarna
func (x *TravelTimeRoute) AverageFunctionalRoadClass() *int {
	if x.data == nil {
		return nil
	}
	return x.data.AverageFunctionalRoadClass
}

// SV: <div class="toggleTitle">Länsnummer</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>1</td> <td>Stockholms län</td> </tr> <tr> <td>2</td> <td> DEPRECATED<br /> Användes tidigare för Stockholms län </td> </tr> <tr> <td>3</td> <td>Uppsala län</td> </tr> <tr> <td>4</td> <td>Södermanlands län</td> </tr> <tr> <td>5</td> <td>Östergötlands län</td> </tr> <tr> <td>6</td> <td>Jönköpings län</td> </tr> <tr> <td>7</td> <td>Kronobergs län</td> </tr> <tr> <td>8</td> <td>Kalmar län</td> </tr> <tr> <td>9</td> <td>Gotlands län</td> </tr> <tr> <td>10</td> <td>Blekinge län</td> </tr> <tr> <td>12</td> <td>Skåne län</td> </tr> <tr> <td>13</td> <td>Hallands län</td> </tr> <tr> <td>14</td> <td>Västra Götalands län</td> </tr> <tr> <td>17</td> <td>Värmlands län</td> </tr> <tr> <td>18</td> <td>Örebro län</td> </tr> <tr> <td>19</td> <td>Västmanlands län</td> </tr> <tr> <td>20</td> <td>Dalarnas län</td> </tr> <tr> <td>21</td> <td>Gävleborgs län</td> </tr> <tr> <td>22</td> <td>Västernorrlands län</td> </tr> <tr> <td>23</td> <td>Jämtlands län</td> </tr> <tr> <td>24</td> <td>Västerbottens län</td> </tr> <tr> <td>25</td> <td>Norrbottens län</td> </tr> </table> </div>
func (x *TravelTimeRoute) Counties() []int {
	if x.data == nil {
		return nil
	}
	return x.data.CountyNo
}

// SV: Landskod
func (x *TravelTimeRoute) CountryCode() *string {
	if x.data == nil {
		return nil
	}
	return x.data.CountryCode
}

// SV: Anger att dataposten raderats
func (x *TravelTimeRoute) Deleted() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Deleted
}

// SV: Anger normal restid i sekunder för sträckan vid skyltad hastighet baserat på historiskt data
func (x *TravelTimeRoute) ExpectedFreeFlowTravelTime() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.ExpectedFreeFlowTravelTime
}

// SV: Anger normal restid i sekunder för sträckan vid skyltad hastighet
func (x *TravelTimeRoute) FreeFlowTravelTime() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.FreeFlowTravelTime
}

func (x *TravelTimeRoute) Geometry() *Geometry {
	if x.data == nil {
		return &Geometry{}
	}
	return &Geometry{data: x.data.Geometry}
}

// SV: Anger id för sträckan
func (x *TravelTimeRoute) ID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Anger längden i meter för sträckan
func (x *TravelTimeRoute) Length() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.Length
}

// SV: Tidpunkt då mätningen av restiden gjordes. Det är slutet på tidsperioden som skickas från källsystemet
func (x *TravelTimeRoute) MeasureTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.MeasureTime
}

// SV: Tidpunkt då dataposten ändrades
func (x *TravelTimeRoute) ModifiedTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.ModifiedTime
}

// SV: Anger namnet på sträckan
func (x *TravelTimeRoute) Name() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Name
}

// SV: <div class="toggleTitle">Ruttägare</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>1</td> <td>Trafiken.nu</td> </tr> <tr> <td>2</td> <td> Ums </td> </tr> <tr> <td>3</td> <td>Pendlingsstråk</td> </tr> <tr> <td>4</td> <td>Trafikverket</td> </tr> </table> </div>
func (x *TravelTimeRoute) RouteOwner() *int {
	if x.data == nil {
		return nil
	}
	return x.data.RouteOwner
}

// SV: Anger hastighet i km/h
func (x *TravelTimeRoute) Speed() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.Speed
}

// SV: <div class="toggleTitle">Anger restidsstatus för sträckan</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> freeflow - fri framkomlighet<br /> heavy - svår framkomlighet<br /> congested - mycket svår framkomlighet<br /> impossible - framkomlighet omöjlig<br /> unknown - framkomlighet okänd </div>
func (x *TravelTimeRoute) TrafficStatus() *string {
	if x.data == nil {
		return nil
	}
	return x.data.TrafficStatus
}

// SV: Anger aktuell restid i sekunder för sträckan
func (x *TravelTimeRoute) TravelTime() *float64 {
	if x.data == nil {
		return nil
	}
	return x.data.TravelTime
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

// SV: Geometrisk linje i koordinatsystem SWEREF99TM
func (x *Geometry) SWEREF99TM() *string {
	if x.data == nil {
		return nil
	}
	return x.data.SWEREF99TM
}

// SV: Geometrisk linje i koordinatsystem WGS84
func (x *Geometry) WGS84() *string {
	if x.data == nil {
		return nil
	}
	return x.data.WGS84
}

// SV: Tidpunkt då elementet ändrades
func (x *Geometry) ModifiedTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.ModifiedTime
}

// Response can be used to decode the response from the API.
//
// For example:
//
//	package main
//
//	import (
//		 "encoding/xml"
//		 api "code.dny.dev/trafikinfo/trv/traveltimeroute/v1dot5"
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
		Error *trv.APIError     `xml:"ERROR"`
		Data  []TravelTimeRoute `xml:"TravelTimeRoute"`
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