// Package v1 contains the type definitions for Camera v1.
//
// All types have accessor methods to access fields which can be chained on nils.
// This makes it possible to easily drill down into deeply nested data.
package v1

import (
	"encoding/xml"
	"strings"
	"time"

	schema "code.dny.dev/trafikinfo/internal/trv/camera/v1"
	"code.dny.dev/trafikinfo/trv"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

func ObjectType() trv.ObjectType {
	return trv.ObjectType{
		Kind:      "Camera",
		Version:   "1",
		Namespace: "",
	}
}

type Camera struct {
	data *schema.Camera
}

func (x *Camera) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Camera{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Anger om kameran är aktiv eller ej
func (x *Camera) Active() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Active
}

// SV: Id för kameragrupp som kameran tillhör
func (x *Camera) CameraGroup() *string {
	if x.data == nil {
		return nil
	}
	return x.data.CameraGroup
}

// SV: Filändelse
func (x *Camera) ContentType() *string {
	if x.data == nil {
		return nil
	}
	return x.data.ContentType
}

// SV: <div class="toggleTitle">Länsnummer</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>1</td> <td>Stockholms län</td> </tr> <tr> <td>2</td> <td> DEPRECATED<br /> Användes tidigare för Stockholms län </td> </tr> <tr> <td>3</td> <td>Uppsala län</td> </tr> <tr> <td>4</td> <td>Södermanlands län</td> </tr> <tr> <td>5</td> <td>Östergötlands län</td> </tr> <tr> <td>6</td> <td>Jönköpings län</td> </tr> <tr> <td>7</td> <td>Kronobergs län</td> </tr> <tr> <td>8</td> <td>Kalmar län</td> </tr> <tr> <td>9</td> <td>Gotlands län</td> </tr> <tr> <td>10</td> <td>Blekinge län</td> </tr> <tr> <td>12</td> <td>Skåne län</td> </tr> <tr> <td>13</td> <td>Hallands län</td> </tr> <tr> <td>14</td> <td>Västra Götalands län</td> </tr> <tr> <td>17</td> <td>Värmlands län</td> </tr> <tr> <td>18</td> <td>Örebro län</td> </tr> <tr> <td>19</td> <td>Västmanlands län</td> </tr> <tr> <td>20</td> <td>Dalarnas län</td> </tr> <tr> <td>21</td> <td>Gävleborgs län</td> </tr> <tr> <td>22</td> <td>Västernorrlands län</td> </tr> <tr> <td>23</td> <td>Jämtlands län</td> </tr> <tr> <td>24</td> <td>Västerbottens län</td> </tr> <tr> <td>25</td> <td>Norrbottens län</td> </tr> </table> </div>
func (x *Camera) Counties() []int {
	if x.data == nil {
		return nil
	}
	return x.data.CountyNo
}

// SV: Anger att dataposten raderats
func (x *Camera) Deleted() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Deleted
}

// SV: Beskrivning
func (x *Camera) Description() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Description
}

// SV: Anger i grader vilket håll kameran är riktad åt
func (x *Camera) Direction() *int {
	if x.data == nil {
		return nil
	}
	return x.data.Direction
}

func (x *Camera) Geometry() *Geometry {
	if x.data == nil {
		return &Geometry{}
	}
	return &Geometry{data: x.data.Geometry}
}

// SV: Anger om det finns ett högupplöst foto. Finns det en högupplöst version av bilden fås den genom att ange queryparametern type=fullsize efter PhotoUrl
func (x *Camera) HasFullSizePhoto() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.HasFullSizePhoto
}

// SV: Anger om det finns skiss över kamerans position och riktning. Finns det en skiss fås den genom att ange queryparametern type=sketch efter PhotoUrl
func (x *Camera) HasSketchImage() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.HasSketchImage
}

// SV: Anger ikonid för kameratypen
func (x *Camera) IconID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.IconId
}

// SV: Unikt id för kameran
func (x *Camera) ID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Fritext som anger var kameran är placerad
func (x *Camera) Location() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Location
}

// SV: Tidpunkt då dataposten ändrades
func (x *Camera) ModifiedTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.ModifiedTime
}

// SV: Namn på kameran
func (x *Camera) Name() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Name
}

// SV: Typ av kamera, "Väglagskamera" eller "Trafikflödeskamera"
func (x *Camera) Type() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Type
}

// SV: Tidpunkt då bilden är tagen.
func (x *Camera) PhotoTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.PhotoTime
}

// SV: <div class="toggleTitle">Url till bild 385px*290px, ytterligare funktionalitet</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> Följande queryparametrar kan användas efter urlen till bilden. <table class="table table-condensed"><tr><td>maxage</td><td>Ange hur gamla bilder i minuter du tillåter. Är bilden äldre än det du anger returneras istället en bild med texten en "aktuell bild saknas". Ex: maxage=15 visar endast bilden om den är nyare än 15 minuter</td></tr><tr><td>type=fullsize</td><td>Om bilden har en högupplöst bild (anges i HasFullSizePhoto) kan man få den genom att ange type=fullsize</td></tr><tr><td>type=sketch</td><td>Om bilden har en översiktsbild (anges i HasSketchImage) kan man få den genom att ange type=sketch</td></tr><tr><td>type=thumbnail</td><td>Vill man ha en mindre version (180px*135px) av bilden anges type=thumbnail</td></tr></table> Ex: <a href="https://api.trafikinfo.trafikverket.se/v1.3/Images/TrafficFlowCamera_39635270.Jpeg?type=fullsize&amp;maxage=15">https://api.trafikinfo.trafikverket.se/v1.3/Images/​TrafficFlowCamera_39635270.Jpeg?type=fullsize&amp;maxage=15</a></div>
func (x *Camera) PhotoURL() *string {
	if x.data == nil {
		return nil
	}
	return x.data.PhotoUrl
}

// SV: Anger en statustext för kameran.
func (x *Camera) Status() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Status
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
func (x *Geometry) SWEREF99TM() *string {
	if x.data == nil {
		return nil
	}
	return x.data.SWEREF99TM
}

// SV: Geometrisk punkt i koordinatsystem WGS84
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
//		 api "code.dny.dev/trafikinfo/trv/camera/v1"
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
		Data  []Camera      `xml:"Camera"`
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
