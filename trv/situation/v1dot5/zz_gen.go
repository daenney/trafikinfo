// Package v1dot5 contains the type definitions for Situation v1.5.
//
// All types have accessor methods to access fields which can be chained on nils.
// This makes it possible to easily drill down into deeply nested data.
package v1dot5

import (
	"encoding/xml"
	"strings"
	"time"

	schema "code.dny.dev/trafikinfo/internal/trv/situation/v1dot5"
	"code.dny.dev/trafikinfo/trv"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

func ObjectType() trv.ObjectType {
	return trv.ObjectType{
		Kind:      "Situation",
		Version:   "1.5",
		Namespace: "",
	}
}

type Situation struct {
	data *schema.Situation
}

func (x *Situation) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Situation{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: <div class="toggleTitle">Landsbeteckning</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> "DE" - Tyskland<br /> "DK" - Danmark<br /> "NO" - Norge<br /> "SE" - Sverige </div>
func (x *Situation) CountryCode() *string {
	if x.data == nil {
		return nil
	}
	return x.data.CountryCode
}

// SV: Anger att dataposten raderats
func (x *Situation) Deleted() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.Deleted
}

func (x *Situation) Deviations() []Deviation {
	if len(x.data.Deviation) == 0 {
		return nil
	}
	data := []Deviation{}
	for _, mem := range x.data.Deviation {
		data = append(data, Deviation{data: &mem})
	}
	return data
}

// SV: Datapostens id
func (x *Situation) ID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

// SV: Tidpunkt då dataposten ändrades
func (x *Situation) ModifiedTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.ModifiedTime
}

// SV: Tidpunkt då dataposten publicerades
func (x *Situation) PublicationTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.PublicationTime
}

// SV: Aktuell versionstid för situationen
func (x *Situation) VersionTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.VersionTime
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

// SV: Påverkad riktning (Båda riktningarna, En riktning)
func (x *Deviation) AffectedDirection() *string {
	if x.data == nil {
		return nil
	}
	return x.data.AffectedDirection
}

// SV: <div class="toggleTitle">Påverkad riktningsvärde</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>BothDirections</td> <td>Båda riktningarna är påverkade</td> </tr> <tr> <td>OneDirection</td> <td>En rikting påverkad. I de fall Deviation.Geometry.Line är definierad så är dess koordinater ordnade efter riktingen.</td> </tr> </table> </div>
func (x *Deviation) AffectedDirectionValue() *string {
	if x.data == nil {
		return nil
	}
	return x.data.AffectedDirectionValue
}

// SV: <div class="toggleTitle">Länsnummer</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>0</td> <td>Alla län (kan förekomma för poster med Deviation.MessageType="Viktig trafikinformation" och meddelandet gäller då för hela Sverige)</td> </tr> <tr> <td>1</td> <td>Stockholms län</td> </tr> <tr> <td>2</td> <td> DEPRECATED<br /> Användes tidigare för Stockholms län </td> </tr> <tr> <td>3</td> <td>Uppsala län</td> </tr> <tr> <td>4</td> <td>Södermanlands län</td> </tr> <tr> <td>5</td> <td>Östergötlands län</td> </tr> <tr> <td>6</td> <td>Jönköpings län</td> </tr> <tr> <td>7</td> <td>Kronobergs län</td> </tr> <tr> <td>8</td> <td>Kalmar län</td> </tr> <tr> <td>9</td> <td>Gotlands län</td> </tr> <tr> <td>10</td> <td>Blekinge län</td> </tr> <tr> <td>12</td> <td>Skåne län</td> </tr> <tr> <td>13</td> <td>Hallands län</td> </tr> <tr> <td>14</td> <td>Västra Götalands län</td> </tr> <tr> <td>17</td> <td>Värmlands län</td> </tr> <tr> <td>18</td> <td>Örebro län</td> </tr> <tr> <td>19</td> <td>Västmanlands län</td> </tr> <tr> <td>20</td> <td>Dalarnas län</td> </tr> <tr> <td>21</td> <td>Gävleborgs län</td> </tr> <tr> <td>22</td> <td>Västernorrlands län</td> </tr> <tr> <td>23</td> <td>Jämtlands län</td> </tr> <tr> <td>24</td> <td>Västerbottens län</td> </tr> <tr> <td>25</td> <td>Norrbottens län</td> </tr> </table> </div>
func (x *Deviation) Counties() []int {
	if x.data == nil {
		return nil
	}
	return x.data.CountyNo
}

// SV: Källa till datat
func (x *Deviation) Creator() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Creator
}

// SV: Tidpunkt då dataposten skapades
func (x *Deviation) CreationTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.CreationTime
}

// SV: Dataposten är giltig till och med
func (x *Deviation) EndTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.EndTime
}

func (x *Deviation) Geometry() *Geometry {
	if x.data == nil {
		return &Geometry{}
	}
	return &Geometry{data: x.data.Geometry}
}

// SV: Titel
func (x *Deviation) Header() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Header
}

// SV: Ikonid
func (x *Deviation) IconID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.IconId
}

// SV: Datapostens id
func (x *Deviation) ID() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Id
}

func (x *Deviation) Images() []Image {
	if len(x.data.Image) == 0 {
		return nil
	}
	data := []Image{}
	for _, mem := range x.data.Image {
		data = append(data, Image{data: &mem})
	}
	return data
}

// SV: Färjeled
func (x *Deviation) JourneyReference() *string {
	if x.data == nil {
		return nil
	}
	return x.data.JourneyReference
}

// SV: Objektet är orsaken till situationen
func (x *Deviation) ManagedCause() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.ManagedCause
}

// SV: Meddelandetext
func (x *Deviation) Message() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Message
}

// SV: Meddelandekod, ex. "Beläggningsarbete"
func (x *Deviation) MessageCode() *string {
	if x.data == nil {
		return nil
	}
	return x.data.MessageCode
}

// SV: Meddelandekodsvärde, ex. "resurfacingWork"
func (x *Deviation) MessageCodeValue() *string {
	if x.data == nil {
		return nil
	}
	return x.data.MessageCodeValue
}

// SV: <div class="toggleTitle">Meddelandetyp, ex: "Vägarbete"</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> "Viktig trafikinformation"<br /> "Färjor"<br /> "Hinder"<br /> "Olycka"<br /> "Restriktion"<br /> "Trafikmeddelande"<br /> "Vägarbete"<br /></div>
func (x *Deviation) MessageType() *string {
	if x.data == nil {
		return nil
	}
	return x.data.MessageType
}

// SV: Meddelandetypsvärde, ex: "MaintenanceWorks"
func (x *Deviation) MessageTypeValue() *string {
	if x.data == nil {
		return nil
	}
	return x.data.MessageTypeValue
}

// SV: Antal påverkade körfält
func (x *Deviation) NumberOfLanesRestricted() *int {
	if x.data == nil {
		return nil
	}
	return x.data.NumberOfLanesRestricted
}

// SV: Påverkad del, ex. "vägren"
func (x *Deviation) PositionalDescription() *string {
	if x.data == nil {
		return nil
	}
	return x.data.PositionalDescription
}

// SV: Vägnamn
func (x *Deviation) RoadName() *string {
	if x.data == nil {
		return nil
	}
	return x.data.RoadName
}

// SV: Vägnummer, ex. "Väg 73"
func (x *Deviation) RoadNumber() *string {
	if x.data == nil {
		return nil
	}
	return x.data.RoadNumber
}

// SV: Vägnummer som nummeriskt värde, ex: 73
func (x *Deviation) RoadNumberNumeric() *int {
	if x.data == nil {
		return nil
	}
	return x.data.RoadNumberNumeric
}

// SV: Indikerar att meddelandet är säkerhetsrelaterat i enlighet med Kommisionens Delegerade Förordning (EU) nr 886/2013 vad gäller data och förfaranden för kostnadsfritt tillhandahållande, när så är möjligt, av ett minimum av vägsäkerhetsrelaterad universell trafikinformation för användare.
func (x *Deviation) SafetyRelatedMessage() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.SafetyRelatedMessage
}

func (x *Deviation) Schedules() []Schedule {
	if len(x.data.Schedule) == 0 {
		return nil
	}
	data := []Schedule{}
	for _, mem := range x.data.Schedule {
		data = append(data, Schedule{data: &mem})
	}
	return data
}

// SV: Påverkansgrad, värden: 1, 2, 4, 5
func (x *Deviation) SeverityCode() *int {
	if x.data == nil {
		return nil
	}
	return x.data.SeverityCode
}

// SV: <div class="toggleTitle"> Påverkan: </div> <div class="toggle arrowR"> </div> <div class="toggleContent"> "Ingen påverkan"<br /> "Liten påverkan"<br /> "Stor påverkan"<br /> "Mycket stor påverkan" </div>
func (x *Deviation) SeverityText() *string {
	if x.data == nil {
		return nil
	}
	return x.data.SeverityText
}

// SV: Dataposten är giltig från och med
func (x *Deviation) StartTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.StartTime
}

// SV: Beskrivning av position
func (x *Deviation) LocationDescriptor() *string {
	if x.data == nil {
		return nil
	}
	return x.data.LocationDescriptor
}

// SV: Tillfälliga begränsningar, ex. "bruttovikt 8 ton"
func (x *Deviation) TemporaryLimit() *string {
	if x.data == nil {
		return nil
	}
	return x.data.TemporaryLimit
}

// SV: Trafikrestriktion, ex. "körfält blockerat"
func (x *Deviation) TrafficRestrictionType() *string {
	if x.data == nil {
		return nil
	}
	return x.data.TrafficRestrictionType
}

// SV: Dataposten gäller på obestämd framtid
func (x *Deviation) ValidUntilFurtherNotice() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.ValidUntilFurtherNotice
}

// SV: Url till mer information
func (x *Deviation) WebLink() *string {
	if x.data == nil {
		return nil
	}
	return x.data.WebLink
}

// SV: Aktuell versionstid för störningen
func (x *Deviation) VersionTime() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.VersionTime
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

func (x *Geometry) Point() *GeoPoint {
	if x.data == nil {
		return &GeoPoint{}
	}
	return &GeoPoint{data: x.data.Point}
}

func (x *Geometry) Line() *GeoLine {
	if x.data == nil {
		return &GeoLine{}
	}
	return &GeoLine{data: x.data.Line}
}

type Image struct {
	data *schema.Image
}

func (x *Image) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.Image{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Gäller för meddelandetyp 'Viktig trafikinformation': anger om det finns ett högupplöst foto
func (x *Image) HasFullSizePhoto() *bool {
	if x.data == nil {
		return nil
	}
	return x.data.HasFullSizePhoto
}

// SV: Gäller för meddelandetyp 'Viktig trafikinformation': bildens URL
func (x *Image) URL() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Url
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

// SV: Objektet är aktivt till och med period
func (x *Schedule) EndOfPeriod() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.EndOfPeriod
}

func (x *Schedule) RecurringTimePeriodOfDays() []RecurringTimePeriodOfDay {
	if len(x.data.RecurringTimePeriodOfDay) == 0 {
		return nil
	}
	data := []RecurringTimePeriodOfDay{}
	for _, mem := range x.data.RecurringTimePeriodOfDay {
		data = append(data, RecurringTimePeriodOfDay{data: &mem})
	}
	return data
}

// SV: Objektet är aktivt från och med period
func (x *Schedule) StartOfPeriod() *time.Time {
	if x.data == nil {
		return nil
	}
	return x.data.StartOfPeriod
}

type GeoPoint struct {
	data *schema.GeoPoint
}

func (x *GeoPoint) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.GeoPoint{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Geometrisk punkt i koordinatsystem SWEREF99TM
func (x *GeoPoint) SWEREF99TM() *string {
	if x.data == nil {
		return nil
	}
	return x.data.SWEREF99TM
}

// SV: Geometrisk punkt i koordinatsystem WGS84
func (x *GeoPoint) WGS84() *string {
	if x.data == nil {
		return nil
	}
	return x.data.WGS84
}

type GeoLine struct {
	data *schema.GeoLine
}

func (x *GeoLine) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.GeoLine{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Geometrisk linje i koordinatsystem SWEREF99TM
func (x *GeoLine) SWEREF99TM() *string {
	if x.data == nil {
		return nil
	}
	return x.data.SWEREF99TM
}

// SV: Geometrisk linje i koordinatsystem WGS84
func (x *GeoLine) WGS84() *string {
	if x.data == nil {
		return nil
	}
	return x.data.WGS84
}

type RecurringTimePeriodOfDay struct {
	data *schema.RecurringTimePeriodOfDay
}

func (x *RecurringTimePeriodOfDay) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	res := &schema.RecurringTimePeriodOfDay{}
	if err := d.DecodeElement(res, &start); err != nil {
		return err
	}
	x.data = res
	return nil
}

// SV: Per dag återkommande aktiv till och med tidpunkt, ex: "16:00"
func (x *RecurringTimePeriodOfDay) End() *string {
	if x.data == nil {
		return nil
	}
	return x.data.End
}

// SV: Per dag återkommande aktiv från och med tidpunkt, ex. "07:00"
func (x *RecurringTimePeriodOfDay) Start() *string {
	if x.data == nil {
		return nil
	}
	return x.data.Start
}

// Response can be used to decode the response from the API.
//
// For example:
//
//	package main
//
//	import (
//		 "encoding/xml"
//		 api "code.dny.dev/trafikinfo/trv/situation/v1dot5"
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
		Data  []Situation   `xml:"Situation"`
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
