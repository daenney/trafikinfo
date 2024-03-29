// Package v1 contains the type definitions for MeasurementData20 v1.
package v1

import (
	"time"
)

// THIS FILE IS AUTOMATICALLY GENERATED FROM THE XSD. DO NOT EDIT.

type MeasurementData20 struct {
	// SV: Län.
	County *int `xml:"County,omitempty"`
	// SV: Huvudvägnummer.
	RoadMainNumber *int `xml:"RoadMainNumber,omitempty"`
	// SV: Undervägnummer. Kallas ibland även punktväg. Måste användas ihop med huvudvägnumret.
	RoadSubNumber *int       `xml:"RoadSubNumber,omitempty"`
	Direction     *Direction `xml:"Direction,omitempty"`
	// SV: Körfält. Räknas från höger sida och startar med körfält 10. Nästa är 20 och kan gå upp till 50.
	Lane *int `xml:"Lane,omitempty"`
	// SV: Start löpande längd. Anges i meter. Starten för aktuell data angivet i antal meter från vägnumrets start inom det länet, i aktuell riktning.
	StartContinuousLength *int `xml:"StartContinuousLength,omitempty"`
	// SV: Slut löpande längd. Anges i meter. Slutet för aktuell data angivet i antal meter från vägnumrets start inom det länet, i aktuell riktning.
	EndContinuousLength *int `xml:"EndContinuousLength,omitempty"`
	// SV: Längd i antal meter.
	Length *int `xml:"Length,omitempty"`
	// SV: Mätdatum. Ett gemensamt datum för en hel mätperiod. Är detta datum som visas i graferna i PMSV3 Analysera sträcka.
	MeasurementDate *time.Time `xml:"MeasurementDate,omitempty"`
	// SV: Mätdatum exakt. Det exakta mätdatumet för aktuellt mätdata.
	MeasurementDateSpecific *time.Time           `xml:"MeasurementDateSpecific,omitempty"`
	MeasurementDataType     *MeasurementDataType `xml:"MeasurementDataType,omitempty"`
	// SV: IRI vänster. Enhet mm/m. IRI (International Roughness Index) Beräknat mått baserat på uppmätt längsprofil i vänster hjulspår. Avser jämnhet i längsled.
	IRILeft *float64 `xml:"IRILeft,omitempty"`
	// SV: IRI höger. Enhet mm/m. IRI (International Roughness Index) Beräknat mått baserat på uppmätt längsprofil i höger hjulspår. Avser jämnhet i längsled.
	IRIRight *float64 `xml:"IRIRight,omitempty"`
	// SV: Mätbilens hastighet.
	MeasurementVehicleSpeed *float64 `xml:"MeasurementVehicleSpeed,omitempty"`
	// SV: Backighet. Anges i %. Backighet är medellutningen för 20m väg i mätriktningen. Uppför har positivt tecken och utför negativt.
	Hilliness *float64 `xml:"Hilliness,omitempty"`
	// SV: Kurvatur. Enhet 10000/r. Kurvatur beskriver medelvärdet av krökningsradien (r) över 20m.
	Curvature *float64 `xml:"Curvature,omitempty"`
	// SV: Kantdjup. Enhet mm. Kantdjup avser att detektera skador närmast vägrenen. Mäts inte direkt av mätbil utan beräknas i efterbearbetning baserat på mätta 20m data om tvärprofil. (se VTI Rapport 718)
	EdgeDepth *float64 `xml:"EdgeDepth,omitempty"`
	// SV: Megatextur höger. Enhet mm.
	MegatextureRight *float64 `xml:"MegatextureRight,omitempty"`
	// SV: Megatextur vänster. Enhet mm.
	MegatextureLeft *float64 `xml:"MegatextureLeft,omitempty"`
	// SV: MPD Makrotextur höger. Enhet mm. Makrotextur beräknad som MPD (Mean Profile Depth beräknat enligt ISO 13473-1), beskriver vägytans 'skrovlighet' i våglängdsintervallet 0,5-50 mm.
	MPDMacrotextureRight *float64 `xml:"MPDMacrotextureRight,omitempty"`
	// SV: MPD Makrotextur mellan. Enhet mm. Makrotextur beräknad som MPD (Mean Profile Depth beräknat enligt ISO 13473-1), beskriver vägytans 'skrovlighet' i våglängdsintervallet 0,5-50 mm.
	MPDMacrotextureMiddle *float64 `xml:"MPDMacrotextureMiddle,omitempty"`
	// SV: MPD Makrotextur vänster. Enhet mm. Makrotextur beräknad som MPD (Mean Profile Depth beräknat enligt ISO 13473-1), beskriver vägytans 'skrovlighet' i våglängdsintervallet 0,5-50 mm.
	MPDMacrotextureLeft *float64 `xml:"MPDMacrotextureLeft,omitempty"`
	// SV: Spårarea. Enhet dm2. Mått för deformation i tvärled, komplement till spårdjup max. Mäts inte direkt av mätbil utan beräknas i efterbearbetning baserat på mätta 20m data om tvärprofil. (se VTI Rapport 718)
	RutArea *float64 `xml:"RutArea,omitempty"`
	// SV: Spårbottenavstånd. Enhet mm. Avstånd mellan de djupaste spåren i vänster resp höger del av tvärprofilen. Mäts inte direkt av mätbil utan beräknas i efterbearbetning baserat på mätta 20m data om tvärprofil. (se VTI Rapport 718)
	RutBottomDistance *float64 `xml:"RutBottomDistance,omitempty"`
	// SV: Spårbredd höger. Enhet mm. Höger hjulspårs bredd. Mäts inte direkt av mätbil utan beräknas i efterbearbetning baserat på mätta 20m data om tvärprofil. (se VTI Rapport 718)
	RutWidthRight *float64 `xml:"RutWidthRight,omitempty"`
	// SV: Spårbredd vänster. Enhet mm. Vänster hjulspårs bredd. Mäts inte direkt av mätbil utan beräknas i efterbearbetning baserat på mätta 20m data om tvärprofil. (se VTI Rapport 718)
	RutWidthLeft *float64 `xml:"RutWidthLeft,omitempty"`
	// SV: Spårdjup höger15. Enhet mm. 15 lasrar.
	RutDepthRight15 *float64 `xml:"RutDepthRight15,omitempty"`
	// SV: Spårdjup höger17. Enhet mm. 17 lasrar.
	RutDepthRight17 *float64 `xml:"RutDepthRight17,omitempty"`
	// SV: Spårdjup max15. Maxvärde 15 lasrar.
	RutDepthMax15 *float64 `xml:"RutDepthMax15,omitempty"`
	// SV: Spårdjup max17. Maxvärde 17 lasrar.
	RutDepthMax17 *float64 `xml:"RutDepthMax17,omitempty"`
	// SV: Spårdjup vänster17. Enhet mm. 17 lasrar.
	RutDepthLeft17 *float64 `xml:"RutDepthLeft17,omitempty"`
	// SV: Spårbottentvärfall. Enhet %.
	CrossfallRutBottom *float64 `xml:"CrossfallRutBottom,omitempty"`
	// SV: Vattenarea. Enhet dm2. Mått för deformation i tvärled, komplement till spårdjup max. Mäts inte direkt av mätbil utan beräknas i efterbearbetning baserat på mätta 20m data om tvärprofil. (se VTI Rapport 718)
	WaterArea *float64 `xml:"WaterArea,omitempty"`
	// SV: Datum för när mätdata hämtades ut från källsystemen.
	TimeStamp *time.Time `xml:"TimeStamp,omitempty"`
	// SV: Tidpunkt då dataposten ändrades
	// EN: Time when the data item was changed
	ModifiedTime *time.Time `xml:"ModifiedTime,omitempty"`
	// SV: Anger att dataposten raderats
	// EN: Indicates that the data record has been deleted
	Deleted *bool `xml:"Deleted,omitempty"`
}

type Direction struct {
	// SV: Riktning. Värde.
	Code *int `xml:"Code,omitempty"`
	// SV: Riktning. Beskrivning.
	Value *string `xml:"Value,omitempty"`
}

type MeasurementDataType struct {
	// SV: Mätdatatyp. Värde.
	Code *int `xml:"Code,omitempty"`
	// SV: Mätdatatyp. Beskrivning.
	Value *string `xml:"Value,omitempty"`
}
