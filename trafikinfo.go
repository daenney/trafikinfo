// Package trafikinfo provides the necessary primitives to interact with the
// Trafikinfo API. It contains a query builder that can be used to build up a
// Request object. You can then xml.Marshal it and pass it on to your
// favourite HTTP client to retrieve it through a bytes.Buffer.
//
// The data returned by the Trafikinfo API is licensed under Creative Commons
// CC0:
// * https://api.trafikinfo.trafikverket.se/DynamicContent/ContentDetails/58e384810bb22118e8041667
// * https://creativecommons.org/publicdomain/zero/1.0/.
//
// The API endpoint is available through the Endpoint constant. The only
// supported endpoint is v2, v1.x is being deprecated by Trafikverket.
//
// The provided structs in the library have a somewhat odd design in that all
// fields are pointers. This happens because the structs are largely
// autogenerated from the JSON schema which defines all attributes as optional.
// This is a consequence of the Include and Exclude primitives which can result
// in only specific fields being present in the response, or specific fields be
// excluded.
package trafikinfo

// Endpoint is the current recommended endpoint for the Trafikinfo API
const Endpoint = "https://api.trafikinfo.trafikverket.se/v2/data.json"

// ObjectType is a type of object you can retrieve from the API
type ObjectType string

const (
	RailCrossing        ObjectType = "RailCrossing"
	ReasonCode          ObjectType = "ReasonCode"
	TrainAnnouncement   ObjectType = "TrainAnnouncement"
	TrainMessage        ObjectType = "TrainMessage"
	TrainStation        ObjectType = "TrainStation"
	TrainStationMessage ObjectType = "TrainStationMessage"
	TrainPosition       ObjectType = "TrainPosition"

	Camera                ObjectType = "Camera"
	FerryAnnouncement     ObjectType = "FerryAnnouncement"
	FerryRoute            ObjectType = "FerryRoute"
	Icon                  ObjectType = "Icon"
	Parking               ObjectType = "Parking"
	RoadCondition         ObjectType = "RoadCondition"
	RoadConditionOverview ObjectType = "RoadConditionOverview"
	Situation             ObjectType = "Situation"
	TrafficFlow           ObjectType = "TrafficFlow"
	TrafficSafetyCamera   ObjectType = "TrafficSafetyCamera"
	TravelTimeRoute       ObjectType = "TravelTimeRoute"
	WeatherMeasurePoint   ObjectType = "WeatherMeasurePoint"
	WeatherObservation    ObjectType = "WeatherObservation"
	WeatherStation        ObjectType = "WeatherStation"

	MeasurementData100 ObjectType = "MeasurementData100"
	MeasurementData20  ObjectType = "MeasurementData20"
	PavementData       ObjectType = "PavementData"
	RoadData           ObjectType = "RoadData"
	RoadGeometry       ObjectType = "RoadGeometry"
)

// APIError represents an error response from the API
type APIError struct {
	Response struct {
		Result []struct {
			Error struct {
				Message string `json:"MESSAGE"`
				Source  string `json:"SOURCE"`
			} `json:"ERROR"`
		} `json:"RESULT"`
	} `json:"RESPONSE"`
}
