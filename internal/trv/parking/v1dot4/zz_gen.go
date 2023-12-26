// Package v1dot4 contains the type definitions for Parking v1.4.
package v1dot4

import (
	"time"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

type Parking struct {
	// SV: <div class="toggleTitle">Länsnummer</div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>1</td> <td>Stockholms län</td> </tr> <tr> <td>2</td> <td> DEPRECATED<br /> Användes tidigare för Stockholms län </td> </tr> <tr> <td>3</td> <td>Uppsala län</td> </tr> <tr> <td>4</td> <td>Södermanlands län</td> </tr> <tr> <td>5</td> <td>Östergötlands län</td> </tr> <tr> <td>6</td> <td>Jönköpings län</td> </tr> <tr> <td>7</td> <td>Kronobergs län</td> </tr> <tr> <td>8</td> <td>Kalmar län</td> </tr> <tr> <td>9</td> <td>Gotlands län</td> </tr> <tr> <td>10</td> <td>Blekinge län</td> </tr> <tr> <td>12</td> <td>Skåne län</td> </tr> <tr> <td>13</td> <td>Hallands län</td> </tr> <tr> <td>14</td> <td>Västra Götalands län</td> </tr> <tr> <td>17</td> <td>Värmlands län</td> </tr> <tr> <td>18</td> <td>Örebro län</td> </tr> <tr> <td>19</td> <td>Västmanlands län</td> </tr> <tr> <td>20</td> <td>Dalarnas län</td> </tr> <tr> <td>21</td> <td>Gävleborgs län</td> </tr> <tr> <td>22</td> <td>Västernorrlands län</td> </tr> <tr> <td>23</td> <td>Jämtlands län</td> </tr> <tr> <td>24</td> <td>Västerbottens län</td> </tr> <tr> <td>25</td> <td>Norrbottens län</td> </tr> </table> </div>
	CountyNo []int `xml:"CountyNo"`
	// SV: Anger att dataposten raderats
	Deleted   *bool       `xml:"Deleted,omitempty"`
	Equipment []Equipment `xml:"Equipment"`
	// SV: Avstånd till närmaste stad
	DistanceToNearestCity *string `xml:"DistanceToNearestCity,omitempty"`
	// SV: Lägesbeskrivning
	LocationDescription *string `xml:"LocationDescription,omitempty"`
	// SV: Beskrivning av parkeringen
	Description *string    `xml:"Description,omitempty"`
	Facility    []Facility `xml:"Facility"`
	// SV: Ikonid
	IconId *string `xml:"IconId,omitempty"`
	// SV: Parkeringens id
	Id *string `xml:"Id,omitempty"`
	// SV: Parkeringens position
	Geometry *Geometry `xml:"Geometry,omitempty"`
	// SV: Tidpunkt då dataposten ändrades
	ModifiedTime *time.Time `xml:"ModifiedTime,omitempty"`
	// SV: Parkeringens namn
	Name *string `xml:"Name,omitempty"`
	// SV: Anger om rastplatsen är öppen eller stängd (open, closed)
	OpenStatus *string `xml:"OpenStatus,omitempty"`
	// SV: Anger om det finns några driftstörningar på rastplatsen (limitedOperation) eller om allt fungerar (inOperation)
	OperationStatus   *string            `xml:"OperationStatus,omitempty"`
	Operator          *Operator          `xml:"Operator,omitempty"`
	ParkingAccess     []ParkingAccess    `xml:"ParkingAccess"`
	Photo             []Photo            `xml:"Photo"`
	TariffsAndPayment *TariffsAndPayment `xml:"TariffsAndPayment,omitempty"`
	// SV: Anger användningsområde
	UsageSenario           []string                 `xml:"UsageSenario"`
	VehicleCharacteristics []VehicleCharacteristics `xml:"VehicleCharacteristics"`
}

type Equipment struct {
	// SV: <div class="toggleTitle"> Typ av utrustning </div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>toilet</td> </tr> <tr> <td>shower</td> </tr> <tr> <td>informationPoint</td> </tr> <tr> <td>informatonStele</td> </tr> <tr> <td>internetTerminal</td> </tr> <tr> <td>internetWireless</td> </tr> <tr> <td>payDesk</td> </tr> <tr> <td>paymentMachine</td> </tr> <tr> <td>cashMachine</td> </tr> <tr> <td>vendingMachine</td> </tr> <tr> <td>faxMachineOrService</td> </tr> <tr> <td>copyMachineOrService</td> </tr> <tr> <td>safeDeposit</td> </tr> <tr> <td>luggageLocker</td> </tr> <tr> <td>publicPhone</td> </tr> <tr> <td>publicCoinPhone</td> </tr> <tr> <td>publicCardPhone</td> </tr> <tr> <td>elevator</td> </tr> <tr> <td>picnicFacilities</td> </tr> <tr> <td>dumpingStation</td> </tr> <tr> <td>freshWater</td> </tr> <tr> <td>wasteDisposal</td> </tr> <tr> <td>refuseBin</td> </tr> <tr> <td>iceFreeScaffold</td> </tr> <tr> <td>playground</td> </tr> <tr> <td>electricChargingStation</td> </tr> <tr> <td>bikeParking</td> </tr> <tr> <td>tollTerminal</td> </tr> <tr> <td>defibrillator</td> </tr> <tr> <td>firstAidEquipment</td> </tr> <tr> <td>fireHose</td> </tr> <tr> <td>fireExtingiusher</td> </tr> <tr> <td>fireHydrant</td> </tr> <tr> <td>none</td> </tr> <tr> <td>unknown</td> </tr> <tr> <td>other</td> </tr> </table> </div>
	Type *string `xml:"Type,omitempty"`
	// SV: <div class="toggleTitle"> Utrustningens tillgänglighet </div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>barrierFreeAccessible</td> </tr> <tr> <td>handicappedAccessible</td> </tr> <tr> <td>wheelChairAccessible</td> </tr> <tr> <td>handicappedEasements</td> </tr> <tr> <td>orientationSystemForBlindPeople</td> </tr> <tr> <td>handicappedMarked</td> </tr> <tr> <td>none</td> </tr> <tr> <td>unknown</td> </tr> <tr> <td>other</td> </tr> </table> </div>
	Accessibility *string `xml:"Accessibility,omitempty"`
}

type Facility struct {
	// SV: <div class="toggleTitle"> Serviceanläggningar vid parkeringen </div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>hotel</td> </tr> <tr> <td>motel</td> </tr> <tr> <td>overnightAccommodation</td> </tr> <tr> <td>shop</td> </tr> <tr> <td>kiosk</td> </tr> <tr> <td>foodShopping</td> </tr> <tr> <td>cafe</td> </tr> <tr> <td>restaurant</td> </tr> <tr> <td>restaurantSelfService</td> </tr> <tr> <td>motorwayRestaurant</td> </tr> <tr> <td>motorwayRestaurantSmall</td> </tr> <tr> <td>sparePartsShopping</td> </tr> <tr> <td>petrolStation</td> </tr> <tr> <td>vehicleMaintenance</td> </tr> <tr> <td>tyreRepair</td> </tr> <tr> <td>truckRepair</td> </tr> <tr> <td>truckWash</td> </tr> <tr> <td>carWash</td> </tr> <tr> <td>pharmacy</td> </tr> <tr> <td>medicalFacility</td> </tr> <tr> <td>police</td> </tr> <tr> <td>touristInformation</td> </tr> <tr> <td>bikeSharing</td> </tr> <tr> <td>docstop</td> </tr> <tr> <td>laundry</td> </tr> <tr> <td>leisureActivities</td> </tr> <tr> <td>unknown</td> </tr> <tr> <td>other</td> </tr> </table> </div>
	Type *string `xml:"Type,omitempty"`
	// SV: <div class="toggleTitle"> Serviceanläggningens tillgänglighet </div> <div class="toggle arrowR"> </div> <div class="toggleContent"> <table class="table table-condensed"> <tr> <td>barrierFreeAccessible</td> </tr> <tr> <td>handicappedAccessible</td> </tr> <tr> <td>wheelChairAccessible</td> </tr> <tr> <td>handicappedEasements</td> </tr> <tr> <td>orientationSystemForBlindPeople</td> </tr> <tr> <td>handicappedMarked</td> </tr> <tr> <td>none</td> </tr> <tr> <td>unknown</td> </tr> <tr> <td>other</td> </tr> </table> </div>
	Accessibility *string `xml:"Accessibility,omitempty"`
}

type Geometry struct {
	// SV: Parkeringens position som en geometrisk punkt i koordinatsystemet SWEREF99TM
	SWEREF99TM *string `xml:"SWEREF99TM,omitempty"`
	// SV: Parkeringens position som en geometrisk punkt i koordinatsystemet WGS84
	WGS84 *string `xml:"WGS84,omitempty"`
}

type Operator struct {
	// SV: Namn på operatörens kontakt
	Contact *string `xml:"Contact,omitempty"`
	// SV: Operatörens kontaktmail
	ContactEmail *string `xml:"ContactEmail,omitempty"`
	// SV: Operatörens kontakttelefon
	ContactTelephoneNumber *string `xml:"ContactTelephoneNumber,omitempty"`
	// SV: Operatörens namn
	Name *string `xml:"Name,omitempty"`
}

type ParkingAccess struct {
	// SV: Parkeringens anslutning som en geometrisk punkt i koordinatsystemet SWEREF99TM
	SWEREF99TM *string `xml:"SWEREF99TM,omitempty"`
	// SV: Parkeringens anslutning som en geometrisk punkt i koordinatsystemet WGS84
	WGS84 *string `xml:"WGS84,omitempty"`
}

type Photo struct {
	// SV: Namn på foto
	Title *string `xml:"Title,omitempty"`
	// SV: Url till foto
	Url *string `xml:"Url,omitempty"`
}

type TariffsAndPayment struct {
	// SV: Anger om parkeringen är gratis att använda
	FreeOfCharge *bool `xml:"FreeOfCharge,omitempty"`
	// SV: Anger parkeringens avgift
	Tariff *string `xml:"Tariff,omitempty"`
}

type VehicleCharacteristics struct {
	// SV: Fordonstyp parkeringen är avsedd för
	VehicleType *string `xml:"VehicleType,omitempty"`
	// SV: Antal platser för fordonstypen inkl släp
	NumberOfSpaces *uint8 `xml:"NumberOfSpaces,omitempty"`
	// SV: Typ av last parkeringen är avsedd för, exempelvis fryst gods (refrigeratedGoods)
	LoadType *string `xml:"LoadType,omitempty"`
}
