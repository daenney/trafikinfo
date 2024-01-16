// Package v1 contains the type definitions for Camera v1.
package v1

import (
	"time"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

type Camera struct {
	// SV: Anger om kameran är aktiv eller ej
	Active *bool `xml:"Active,omitempty"`
	// SV: Id för kameragrupp som kameran tillhör
	CameraGroup *string `xml:"CameraGroup,omitempty"`
	// SV: Filändelse
	ContentType *string `xml:"ContentType,omitempty"`
	// SV: <div class="toggleTitle">Länsnummer</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>1</td> <td>Stockholms län</td> </tr> <tr> <td>2</td> <td> DEPRECATED<br /> Användes tidigare för Stockholms län </td> </tr> <tr> <td>3</td> <td>Uppsala län</td> </tr> <tr> <td>4</td> <td>Södermanlands län</td> </tr> <tr> <td>5</td> <td>Östergötlands län</td> </tr> <tr> <td>6</td> <td>Jönköpings län</td> </tr> <tr> <td>7</td> <td>Kronobergs län</td> </tr> <tr> <td>8</td> <td>Kalmar län</td> </tr> <tr> <td>9</td> <td>Gotlands län</td> </tr> <tr> <td>10</td> <td>Blekinge län</td> </tr> <tr> <td>12</td> <td>Skåne län</td> </tr> <tr> <td>13</td> <td>Hallands län</td> </tr> <tr> <td>14</td> <td>Västra Götalands län</td> </tr> <tr> <td>17</td> <td>Värmlands län</td> </tr> <tr> <td>18</td> <td>Örebro län</td> </tr> <tr> <td>19</td> <td>Västmanlands län</td> </tr> <tr> <td>20</td> <td>Dalarnas län</td> </tr> <tr> <td>21</td> <td>Gävleborgs län</td> </tr> <tr> <td>22</td> <td>Västernorrlands län</td> </tr> <tr> <td>23</td> <td>Jämtlands län</td> </tr> <tr> <td>24</td> <td>Västerbottens län</td> </tr> <tr> <td>25</td> <td>Norrbottens län</td> </tr> </table> </div>
	CountyNo []int `xml:"CountyNo,omitempty"`
	// SV: Anger att dataposten raderats
	Deleted *bool `xml:"Deleted,omitempty"`
	// SV: Beskrivning
	Description *string `xml:"Description,omitempty"`
	// SV: Anger i grader vilket håll kameran är riktad åt
	Direction *int      `xml:"Direction,omitempty"`
	Geometry  *Geometry `xml:"Geometry,omitempty"`
	// SV: Anger om det finns ett högupplöst foto. Finns det en högupplöst version av bilden fås den genom att ange queryparametern type=fullsize efter PhotoUrl
	HasFullSizePhoto *bool `xml:"HasFullSizePhoto,omitempty"`
	// SV: Anger om det finns skiss över kamerans position och riktning. Finns det en skiss fås den genom att ange queryparametern type=sketch efter PhotoUrl
	HasSketchImage *bool `xml:"HasSketchImage,omitempty"`
	// SV: Anger ikonid för kameratypen
	IconId *string `xml:"IconId,omitempty"`
	// SV: Unikt id för kameran
	Id *string `xml:"Id,omitempty"`
	// SV: Fritext som anger var kameran är placerad
	Location *string `xml:"Location,omitempty"`
	// SV: Tidpunkt då dataposten ändrades
	ModifiedTime *time.Time `xml:"ModifiedTime,omitempty"`
	// SV: Namn på kameran
	Name *string `xml:"Name,omitempty"`
	// SV: Typ av kamera, "Väglagskamera" eller "Trafikflödeskamera"
	Type *string `xml:"Type,omitempty"`
	// SV: Tidpunkt då bilden är tagen.
	PhotoTime *time.Time `xml:"PhotoTime,omitempty"`
	// SV: <div class="toggleTitle">Url till bild 385px*290px, ytterligare funktionalitet</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> Följande queryparametrar kan användas efter urlen till bilden. <table class="table table-condensed"><tr><td>maxage</td><td>Ange hur gamla bilder i minuter du tillåter. Är bilden äldre än det du anger returneras istället en bild med texten en "aktuell bild saknas". Ex: maxage=15 visar endast bilden om den är nyare än 15 minuter</td></tr><tr><td>type=fullsize</td><td>Om bilden har en högupplöst bild (anges i HasFullSizePhoto) kan man få den genom att ange type=fullsize</td></tr><tr><td>type=sketch</td><td>Om bilden har en översiktsbild (anges i HasSketchImage) kan man få den genom att ange type=sketch</td></tr><tr><td>type=thumbnail</td><td>Vill man ha en mindre version (180px*135px) av bilden anges type=thumbnail</td></tr></table> Ex: <a href="https://api.trafikinfo.trafikverket.se/v1.3/Images/TrafficFlowCamera_39635270.Jpeg?type=fullsize&amp;maxage=15">https://api.trafikinfo.trafikverket.se/v1.3/Images/​TrafficFlowCamera_39635270.Jpeg?type=fullsize&amp;maxage=15</a></div>
	PhotoUrl *string `xml:"PhotoUrl,omitempty"`
	// SV: Anger en statustext för kameran.
	Status *string `xml:"Status,omitempty"`
}

type Geometry struct {
	// SV: Geometrisk punkt i koordinatsystem SWEREF99TM
	SWEREF99TM *string `xml:"SWEREF99TM,omitempty"`
	// SV: Geometrisk punkt i koordinatsystem WGS84
	WGS84 *string `xml:"WGS84,omitempty"`
}
