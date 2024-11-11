package main


import "unsafe"
import (
   "fmt"
   "encoding/json"
)
import "github.com/princefr/siri_parser/go/generated"



/*
#cgo LDFLAGS: -L${SRCDIR}/../parser_unsafe/target/release/ -lparser
#cgo LDFLAGS: -Wl,-rpath,${SRCDIR}/../parser_unsafe/target/release

char* parse_str(const char *s);

#include <stdlib.h>
*/
import "C"


// ParseStr parses the given C string and returns a C string
// containing the parse result.
// func ParseStr(s string) *string {
// 	cStr := C.CString(s)
// 	defer C.free(unsafe.Pointer(cStr)) // Free the C string after use

// 	// Call the Rust function
// 	result := C.parse_str(cStr)

// 	// transform the string to go string
// 	goString := C.GoString(result)
	
// 	// create a new string on the heap and return its pointer
// 	reutrnSr := new(string)
// 	*reutrnSr = goString
// 	return reutrnSr
// }


func ParseStr(s string) (*generated.Envelope, error) {
	cStr := C.CString(s)
	defer C.free(unsafe.Pointer(cStr)) // Free the C string after use

	// Call the Rust function
	result := C.parse_str(cStr)
	if result == nil {
		return nil, fmt.Errorf("failed to parse string")
	}

	// Transform the C string to Go string
	goString := C.GoString(result)

   fmt.Println(goString)

	// Create a variable to hold the unmarshaled data
	var body generated.Envelope

	// Unmarshal the JSON into the body variable
	err := json.Unmarshal([]byte(goString), &body)
	if err != nil {
		return nil, fmt.Errorf("failed to unmarshal JSON: %v", err)
	}

	return &body, nil
}

func main() {
	input := `
	<?xml version="1.0" encoding="UTF-8"?>
<soapenv:Envelope xmlns:soapenv="http://schemas.xmlsoap.org/soap/envelope/" xmlns:siri="http://wsdl.siri.org.uk" xmlns:siri1="http://www.siri.org.uk/siri" xmlns:acsb="http://www.ifopt.org.uk/acsb" xmlns:ifop="http://www.ifopt.org.uk/ifopt">
   <soapenv:Header/>
   <soapenv:Body>
      <siri:NotifyEstimatedTimetable>
         <ServiceDeliveryInfo>
            <siri1:ResponseTimestamp>2022-01-17T09:30:46+02:00</siri1:ResponseTimestamp>
            <siri1:ProducerRef>CBOX</siri1:ProducerRef>
         </ServiceDeliveryInfo>
         <Notification>
            <siri1:EstimatedTimetableDelivery version="2.0">
               <siri1:ResponseTimestamp>2022-01-17T09:30:46+02:00</siri1:ResponseTimestamp>
               <siri1:SubscriberRef>Navitia</siri1:SubscriberRef>
               <siri1:EstimatedJourneyVersionFrame>
                  <siri1:RecordedAtTime>2022-01-17T09:30:00+02:00</siri1:RecordedAtTime>
                  <siri1:EstimatedVehicleJourney>
                     <siri1:LineRef>STIF:Line::C01727:</siri1:LineRef>
                     <siri1:PublishedLineName xml:lang="fr">Mon libellé de ligne</siri1:PublishedLineName>
                     <siri1:DirectionRef>Outbound</siri1:DirectionRef> <!-- valeur en dur -->
                     <siri1:EstimatedVehicleJourneyCode>SNCF_ACCES_CLOUD:VehicleJourney::ef7c2657-c9eb-4fd2-8571-f5831ea4d59b:LOC</siri1:EstimatedVehicleJourneyCode>
                     <siri1:Cancellation>false</siri1:Cancellation>
                     <siri1:ExtraJourney>true</siri1:ExtraJourney>
                     <siri1:JourneyPatternName>PIKO</siri1:JourneyPatternName>
                     <siri1:VehicleMode>rail</siri1:VehicleMode> <!-- bus ou rail ou tram -->
                     <siri1:OriginRef>?</siri1:OriginRef>
                     <siri1:OriginName xml:lang="?">?</siri1:OriginName>
                     <siri1:DestinationRef>?</siri1:DestinationRef>
                     <siri1:DestinationName xml:lang="?">?</siri1:DestinationName>
                     <siri1:OperatorRef>800</siri1:OperatorRef>
                     <siri1:ProductCategoryRef>suburbanRailway</siri1:ProductCategoryRef> <!--
                        mode avancé : bus, railShuttle (CDGVal et OrlyVal), suburbanRailway (trains TN),
                        regionalRail (TER), interregionalRail (intercités), local (RER), tram -->
                     <siri1:TrainNumbers>
                        <siri1:TrainNumberRef>155840-155841</siri1:TrainNumberRef>
                     </siri1:TrainNumbers>
                     <siri1:VehicleJourneyName>148836-148837</siri1:VehicleJourneyName>
                     <siri1:OriginAimedDepartureTime>2021-11-22T14:06:10.000Z</siri1:OriginAimedDepartureTime>
                     <siri1:DestinationAimedArrivalTime>2021-11-22T14:06:10.000Z</siri1:DestinationAimedArrivalTime>
                     <siri1:RecordedCalls>
                        <siri1:RecordedCall>
                          <siri1:StopPointRef>STIF:StopPoint:SP:47095:</siri1:StopPointRef>
                          <siri1:Order>1</siri1:Order>
                          <siri1:ExpectedDepartureTime>2021-11-22T14:06:10.000Z</siri1:ExpectedDepartureTime>
                          <siri1:ExpectedArrivalTime>2021-11-22T14:06:10.000Z</siri1:ExpectedArrivalTime>
                          <siri1:Cancellation>false</siri1:Cancellation>
                          <siri1:ExtraCall>true</siri1:ExtraCall>
                        </siri1:RecordedCall>
                        <siri1:RecordedCall>
                          <siri1:StopPointRef>STIF:StopPoint:SP:43202:</siri1:StopPointRef>
                          <siri1:Order>3</siri1:Order>
                          <siri1:ExpectedDepartureTime>2021-11-22T14:44:50.000Z</siri1:ExpectedDepartureTime>
                          <siri1:ExpectedArrivalTime>2021-11-22T14:44:10.000Z</siri1:ExpectedArrivalTime>
                          <siri1:Cancellation>false</siri1:Cancellation>
                          <siri1:ExtraCall>true</siri1:ExtraCall>
                        </siri1:RecordedCall>
                     </siri1:RecordedCalls>
                     <siri1:EstimatedCalls>
                        <siri1:EstimatedCall>
                          <siri1:StopPointRef>STIF:StopPoint:SP:45301:</siri1:StopPointRef>
                          <siri1:Order>4</siri1:Order>
                          <siri1:ExpectedDepartureTime>2021-11-22T15:11:40.000Z</siri1:ExpectedDepartureTime>
                          <siri1:ExpectedArrivalTime>2021-11-22T15:10:40.000Z</siri1:ExpectedArrivalTime>
                          <siri1:Cancellation>false</siri1:Cancellation>
                          <siri1:ExtraCall>true</siri1:ExtraCall>
                        </siri1:EstimatedCall>
                        <siri1:EstimatedCall>
                          <siri1:StopPointRef>STIF:StopPoint:SP:43072:</siri1:StopPointRef>
                          <siri1:Order>5</siri1:Order>
                          <siri1:ExpectedDepartureTime>2021-11-22T15:15:10.000Z</siri1:ExpectedDepartureTime>
                          <siri1:ExpectedArrivalTime>2021-11-22T15:14:20.000Z</siri1:ExpectedArrivalTime>
                          <siri1:Cancellation>false</siri1:Cancellation>
                          <siri1:ExtraCall>true</siri1:ExtraCall>
                        </siri1:EstimatedCall>
                        <siri1:EstimatedCall>
                          <siri1:StopPointRef>STIF:StopPoint:Q:58757:</siri1:StopPointRef>
                          <siri1:Order>6</siri1:Order>
                          <siri1:ExpectedDepartureTime>2021-11-22T15:30:10.000Z</siri1:ExpectedDepartureTime>
                          <siri1:ExpectedArrivalTime>2021-11-22T15:29:30.000Z</siri1:ExpectedArrivalTime>
                          <siri1:Cancellation>false</siri1:Cancellation>
                          <siri1:ExtraCall>true</siri1:ExtraCall>
                        </siri1:EstimatedCall>
                        <siri1:EstimatedCall>
                          <siri1:StopPointRef>STIF:StopPoint:BP:43219:</siri1:StopPointRef>
                          <siri1:Order>7</siri1:Order>
                          <siri1:ExpectedDepartureTime>2021-11-22T15:57:40.000Z</siri1:ExpectedDepartureTime>
                          <siri1:ExpectedArrivalTime>2021-11-22T15:56:40.000Z</siri1:ExpectedArrivalTime>
                          <siri1:Cancellation>false</siri1:Cancellation>
                          <siri1:ExtraCall>true</siri1:ExtraCall>
                        </siri1:EstimatedCall>
                        <siri1:EstimatedCall>
                          <siri1:StopPointRef>STIF:StopPoint:SP:461504:</siri1:StopPointRef>
                          <siri1:Order>8</siri1:Order>
                          <siri1:ExpectedDepartureTime>2021-11-22T16:06:20.000Z</siri1:ExpectedDepartureTime>
                          <siri1:ExpectedArrivalTime>2021-11-22T16:06:20.000Z</siri1:ExpectedArrivalTime>
                          <siri1:Cancellation>false</siri1:Cancellation>
                          <siri1:ExtraCall>true</siri1:ExtraCall>
                        </siri1:EstimatedCall>
                     </siri1:EstimatedCalls>
                  </siri1:EstimatedVehicleJourney>
               </siri1:EstimatedJourneyVersionFrame>
            </siri1:EstimatedTimetableDelivery>
         </Notification>
      </siri:NotifyEstimatedTimetable>
   </soapenv:Body>
</soapenv:Envelope>
`
	data, err := ParseStr(input)
   if err != nil {
      fmt.Printf("Parse error: %s\n", err.Error())
   }
   fmt.Printf("%+v\n", data.Body.EstimatedTimetable())
}