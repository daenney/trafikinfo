// Package v1dot4 contains the type definitions for TrainStation v1.4.
package v1dot4

import (
	"time"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

type TrainStation struct {
	// SV: Anger om stationen annonseras i tidtabell
	// EN: Indicates if the station is advertised in the timetable
	Advertised *bool `xml:"Advertised,omitempty"`
	// SV: Stationens namn
	// EN: Stations name
	AdvertisedLocationName *string `xml:"AdvertisedLocationName,omitempty"`
	// SV: Stationens namn i kort version
	// EN: Stations name in short version
	AdvertisedShortLocationName *string `xml:"AdvertisedShortLocationName,omitempty"`
	// SV: Beteckning för i vilket land stationen finns.<br /> 'DE' - Tyskland<br /> 'DK' - Danmark<br /> 'NO' - Norge<br /> 'SE' - Sverige
	// EN: Designation of the country in which the station is located.<br /> 'DE' - Germany<br /> 'DK' - Denmark<br /> 'NO' - Norway<br /> 'SE' - Sweden
	CountryCode *string `xml:"CountryCode,omitempty"`
	// SV: <div class="toggleTitle">Länsnummer</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>1</td> <td>Stockholms län</td> </tr> <tr> <td>2</td> <td> DEPRECATED<br /> Användes tidigare för Stockholms län </td> </tr> <tr> <td>3</td> <td>Uppsala län</td> </tr> <tr> <td>4</td> <td>Södermanlands län</td> </tr> <tr> <td>5</td> <td>Östergötlands län</td> </tr> <tr> <td>6</td> <td>Jönköpings län</td> </tr> <tr> <td>7</td> <td>Kronobergs län</td> </tr> <tr> <td>8</td> <td>Kalmar län</td> </tr> <tr> <td>9</td> <td>Gotlands län</td> </tr> <tr> <td>10</td> <td>Blekinge län</td> </tr> <tr> <td>12</td> <td>Skåne län</td> </tr> <tr> <td>13</td> <td>Hallands län</td> </tr> <tr> <td>14</td> <td>Västra Götalands län</td> </tr> <tr> <td>17</td> <td>Värmlands län</td> </tr> <tr> <td>18</td> <td>Örebro län</td> </tr> <tr> <td>19</td> <td>Västmanlands län</td> </tr> <tr> <td>20</td> <td>Dalarnas län</td> </tr> <tr> <td>21</td> <td>Gävleborgs län</td> </tr> <tr> <td>22</td> <td>Västernorrlands län</td> </tr> <tr> <td>23</td> <td>Jämtlands län</td> </tr> <tr> <td>24</td> <td>Västerbottens län</td> </tr> <tr> <td>25</td> <td>Norrbottens län</td> </tr> </table> </div>
	CountyNo []int     `xml:"CountyNo,omitempty"`
	Deleted  *bool     `xml:"Deleted,omitempty"`
	Geometry *Geometry `xml:"Geometry,omitempty"`
	// SV: Upplysningsinformation för stationen, ex. "SL-tåg omfattas ej.", "Ring 033-172444 för trafikinformation"
	// EN: Disclosure information for the station, ex. "SL-train is excluded.", "Call 033-172444 for trafficinfomation"
	LocationInformationText *string `xml:"LocationInformationText,omitempty"`
	// SV: Stationens unika signatur, ex. "Cst"
	// EN: Stations unique signature, ex. "Cst"
	LocationSignature *string `xml:"LocationSignature,omitempty"`
	// SV: Plattformens spår
	// EN: Platform Tracks
	PlatformLine []string `xml:"PlatformLine,omitempty"`
	// SV: Anger om stationen prognostiseras i tidtabell
	// EN: Specifies if station forecasted in timetabell
	Prognosticated *bool `xml:"Prognosticated,omitempty"`
	// SV: Det av Transportstyrelsen fastslagna officiella namnet på stationen
	// EN: The official name of the station established by the Swedish Transport Agency
	OfficialLocationName *string `xml:"OfficialLocationName,omitempty"`
	// EN: Specifies when the object is stored.
	// SV: Anger när objektet är sparat.
	ModifiedTime *time.Time `xml:"ModifiedTime,omitempty"`
}

type Geometry struct {
	// SV: Geometrisk punkt i koordinatsystem SWEREF99TM
	// EN: Geometric point in coordinate system SWEREF99TM
	SWEREF99TM *string `xml:"SWEREF99TM,omitempty"`
	// SV: Geometrisk punkt i koordinatsystem WGS84
	// EN: Geometric point in coordinate system WGS84
	WGS84 *string `xml:"WGS84,omitempty"`
}
