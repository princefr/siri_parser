package generated

type Envelope struct {
	Body Body `json:"Body,omitempty"`
}

type Body struct {
	//NotifyProductionTimetable *ProductionTimetable `json:"ProductionTimetable,omitempty"`
	NotifyEstimatedTimetable *NotifyEstimatedTimetable `json:"EstimatedTimetable,omitempty"`
	// NotifyStopMonitoring *StopMonitoring `json:"NotifyStopMonitoring,omitempty"`
	// NotifyVechicleMonitoring *VechicleMonitoring `json:"NotifyVechicleMonitoring,omitempty"`
	// NotifyConnectionMonitoring *ConnectionMonitoring `json:"NotifyConnectionMonitoring,omitempty"`
	// NotifyGeneralMessage *GeneralMessage `json:"NotifyGeneralMessage,omitempty"`
	// NotifyFacilityMonitoring *FacilityMonitoring `json:"NotifyFacilityMonitoring,omitempty"`
	// NotifySituationExchange *SituationExchange `json:"NotifySituationExchange,omitempty"`
}

type EstimatedTimetableDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	Estimated_journey_version_frame EstimatedJourneyVersionFrame `json:"EstimatedJourneyVersionFrame,omitempty"`
}

type ConnectionMonitoringDelivery struct {
}

type FacilityMonitoringDelivery struct {
	leader XxxDelivery `json:"Leader,omitempty"`
	facility_condition FacilityCondition `json:"FacilityCondition,omitempty"`
}

type GeneralMessageDelivery struct {
	leader XxxDelivery `json:"Leader,omitempty"`
}

type ProductionTimetableDelivery struct {
}

type SituationExchangeDelivery struct {
	leader XxxDelivery `json:"Leader,omitempty"`
}

type StopMonitoringDelivery struct {
	leader XxxDelivery `json:"Leader,omitempty"`
	monitoring_ref *string `json:"MonitoringRef,omitempty"`
	monitored_stop_visit *[]MonitoredStopVisit `json:"MonitoredStopVisit,omitempty"`
	monitored_stop_visit_cancellation *[]MonitoredStopVisitCancellation `json:"MonitoredStopVisitCancellation,omitempty"`
}

type VehicleMonitoringDelivery struct {
	leader XxxDelivery `json:"Leader,omitempty"`
	vehicle_activity *[]VehicleActivity `json:"VehicleActivity,omitempty"`
	vehicle_activity_cancellation *[]VehicleActivityCancellation `json:"VehicleActivityCancellation,omitempty"`
}

type SituationMonitoringDelivery struct {
}

type ArrivalStatus int

const (
	OnTimeArrivalStatus ArrivalStatus = iota
	MissedArrivalStatus 
	ArrivedArrivalStatus 
	NotExpectedArrivalStatus 
	DelayedArrivalStatus 
	EarlyArrivalStatus 
	CancelledArrivalStatus  
	NoReportArrivalStatus 
)


type BoardingActivity int

const (
	Alighting BoardingActivity = 0
	NoAlighting BoardingActivity = 1
	PassThru BoardingActivity = 2
	Boarding BoardingActivity = 3
	NoBoarding BoardingActivity = 4
)


type DepartureStatus int

const (
	OnTimeDepartureStatus DepartureStatus = iota
	EarlyDepartureStatus 
	DelayedDepartureStatus  
	CancelledDepartureStatus 
	ArrivedDepartureStatus 
	DepartedDepartureStatus 
	NotExpectedDepartureStatus 
	NoReportDepartureStatus 
)


type Occupancy int

const (
	FullOccupancy Occupancy = iota
	SeatsAvailableOccupancy 
	StandingAvailableOccupancy
	UnknownOccupancy
	EmptyOccupancy
	ManySeatAvailableOccupancy 
	FewSeatAvailableOccupancy
	StandingRoomOnlyOccupancy
	CrushStandingRoomOnlyOccupancy
	NotAcceptingPassengersOccupancy
)


type VehicleMode int

const (
	Air VehicleMode = 0
	Bus VehicleMode = 1
	Coach VehicleMode = 2
	Ferry VehicleMode = 3
	Metro VehicleMode = 4
	Rail VehicleMode = 5
	Tram VehicleMode = 6
	Underground VehicleMode = 7
)


type MonitoringError int

const (
	GPS MonitoringError = 0
	GPRS MonitoringError = 1
	Radio MonitoringError = 2
)


// type Coordinates int

// const (
// 	type LatLong struct {
// 	*float64
// 	*float64
// }
// 	LatLong Coordinates = 0
// 	type GML struct {
// 	*string
// }
// 	GML Coordinates = 1
// )


type FacilityClass int

const (
	FixedEquipment FacilityClass = 0
	MobileEquipment FacilityClass = 1
	SiteComponent FacilityClass = 2
	Site FacilityClass = 3
	ParkingBay FacilityClass = 4
	Vehicle FacilityClass = 5
)


type EndTimePrecision int

const (
	Second EndTimePrecision = 0
	Minute EndTimePrecision = 1
	Hour EndTimePrecision = 2
	Day EndTimePrecision = 3
)


type FacilityAvailability int

const (
	Unknown FacilityAvailability = 0
	Available FacilityAvailability = 1
	NotAvailable FacilityAvailability = 2
	PartiallyAvailable FacilityAvailability = 3
	Added FacilityAvailability = 4
	Removed FacilityAvailability = 5
)


type EstimatedCalls struct {
	Calls []EstimatedCall `json:"EstimatedCall,omitempty"`
}

type EstimatedJourneyVersionFrame struct {
	Recorded_at_time string `json:"RecordedAtTime,omitempty"`
	Version_ref *string `json:"VersionRef,omitempty"`
	Estimated_vehicle_journey []EstimatedVehicleJourney `json:"EstimatedVehicleJourney,omitempty"`
}

type EstimatedVehicleJourney struct {
	Line_ref string `json:"LineRef,omitempty"`
	Published_line_name string `json:"PublishedLineName,omitempty"`
	Direction_ref string `json:"DirectionRef,omitempty"`
	//journey_identifier *JourneyIdentifier `json:"JourneyIdentifier,omitempty"`
	Dated_vehicule_journey_ref string `json:"DatedVehiculeJourneyRef,omitempty"`
	Cancellation string `json:"Cancellation,omitempty"`
	Extra_journey bool `json:"ExtraJourney,omitempty"`
	Journey_pattern_name string `json:"JourneyPatternName,omitempty"`
	Journey_pattern_info JourneyPatternInfo `json:"JourneyPatternInfo,omitempty"`
	Vehicle_mode string `json:"VehicleMode,omitempty"`
	Origin_ref string `json:"OriginRef,omitempty"`
	Origin_name string `json:"OriginName,omitempty"`
	Destination_ref string `json:"DestinationRef,omitempty"`
	Destination_name string `json:"DestinationName,omitempty"`
	Operator_ref string `json:"OperatorRef,omitempty"`
	Product_category_ref string `json:"ProductCategoryRef,omitempty"`
	Train_numbers TrainNumbers `json:"TrainNumbers,omitempty"`
	Vehicule_journey_name string `json:"VehiculeJourneyName,omitempty"`
	Origin_aimed_departure_time string `json:"OriginAimedDepartureTime,omitempty"`
	Destination_aimed_arrival_time string `json:"DestinationAimedArrivalTime,omitempty"`
	Recorded_calls RecordedCalls `json:"RecordedCalls,omitempty"`
	Estimated_calls EstimatedCalls `json:"EstimatedCalls,omitempty"`
	Framed_vehicle_journey_ref FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef,omitempty"`
	Data_source string `json:"DataSource,omitempty"`
	Vehicle_ref string `json:"VehicleRef,omitempty"`
	Aimed_departure_time string `json:"AimedDepartureTime,omitempty"`
	Aimed_arrival_time string `json:"AimedArrivalTime,omitempty"`
	Journey_note string `json:"JourneyNote,omitempty"`
	Headway_service string `json:"HeadwayService,omitempty"`
	First_or_last_journey string `json:"FirstOrLastJourney,omitempty"`
	Disruption_group string `json:"DisruptionGroup,omitempty"`
	Journey_progress_info string `json:"JourneyProgressInfo,omitempty"`
}

type FramedVehicleJourneyRef struct {
	Data_frame_ref *string `json:"DataFrameRef,omitempty"`
	Dated_vehicle_journey_ref *string `json:"DatedVehicleJourneyRef,omitempty"`
}

type RecordedCalls struct {
	Calls []RecordedCall `json:"*[]RecordedCall,omitempty"`
}

type ServiceDeliveryInfo struct {
	Response_timestamp string `json:"ResponseTimestamp,omitempty"`
	Producer_ref string `json:"ProducerRef,omitempty"`
	Request_message_identifier *string `json:"RequestMessageIdentifier,omitempty"`
	Response_message_ref *string `json:"ResponseMessageRef,omitempty"`
}

type TrainNumbers struct {
	train_number_ref *string `json:"TrainNumberRef,omitempty"`
}

type XxxDelivery struct {
	response_timestamp string `json:"ResponseTimestamp,omitempty"`
	request_message_ref *string `json:"RequestMessageRef,omitempty"`
	subscriber_ref *string `json:"SubscriberRef,omitempty"`
	subscription_ref *string `json:"SubscriptionRef,omitempty"`
	status *bool `json:"Status,omitempty"`
	valid_until *string `json:"ValidUntil,omitempty"`
	shortest_possible_cycle *uint32 `json:"ShortestPossibleCycle,omitempty"`
}

type StopMonitoringNotification struct {
	stop_monitoring_delivery StopMonitoringDelivery `json:"StopMonitoringDelivery,omitempty"`
}

type ProductionTimetableNotification struct {
	production_timetable_delivery ProductionTimetableDelivery `json:"ProductionTimetableDelivery,omitempty"`
}

type SituationMonitoringNotification struct {
	situation_monitoring_delivery SituationMonitoringDelivery `json:"SituationMonitoringDelivery,omitempty"`
}

type GeneralMessageNotification struct {
	general_message_delivery GeneralMessageDelivery `json:"GeneralMessageDelivery,omitempty"`
}

type FacilityMonitoringNotification struct {
	facility_monitoring_delivery FacilityMonitoringDelivery `json:"FacilityMonitoringDelivery,omitempty"`
}

type ConnectionMonitoringNotification struct {
	connection_monitoring_delivery ConnectionMonitoringDelivery `json:"ConnectionMonitoringDelivery,omitempty"`
}

type EstimatedTimetableNotification struct {
	Estimated_timetable_delivery EstimatedTimetableDelivery `json:"EstimatedTimetableDelivery,omitempty"`
}

type SituationExchangeNotification struct {
	situation_exchange_delivery SituationExchangeDelivery `json:"SituationExchangeDelivery,omitempty"`
}

type VehicleMonitoringNotification struct {
	vehicle_monitoring_delivery VehicleMonitoringDelivery `json:"VehicleMonitoringDelivery,omitempty"`
}

type NotifyEstimatedTimetable struct {
	Service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification EstimatedTimetableNotification `json:"Notification,omitempty"`
}

type NotifyStopMonitoring struct {
	service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	notification StopMonitoringNotification `json:"Notification,omitempty"`
}

type NotifyConnectionMonitoring struct {
	service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	notification ConnectionMonitoringNotification `json:"Notification,omitempty"`
}

type NotifyFacilityMonitoring struct {
	service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	notification FacilityMonitoringNotification `json:"Notification,omitempty"`
}

type NotifyVechicleMonitoring struct {
	service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	notification VehicleMonitoringNotification `json:"Notification,omitempty"`
}

type NotifySituationExchange struct {
	service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	notification SituationExchangeNotification `json:"Notification,omitempty"`
}

type NotifyGeneralMessage struct {
	service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	notification GeneralMessageNotification `json:"Notification,omitempty"`
}

type NotifyProductionTimetable struct {
	service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	notification ProductionTimetableNotification `json:"Notification,omitempty"`
}

type JourneyEndNames struct {
}

// type JourneyIdentifier int

// const (
// 	type DatedVehicleJourneyRef struct {
// 	string
// }
// 	DatedVehicleJourneyRef JourneyIdentifier = 0
// 	type EstimatedVehicleJourneyCode struct {
// 	string
// }
// 	EstimatedVehicleJourneyCode JourneyIdentifier = 1
// 	type DatedVehicleJourneyIndirectRef struct {
// 	DatedVehicleJourneyIndirectRef
// }
// 	DatedVehicleJourneyIndirectRef JourneyIdentifier = 2
// )


type JourneyPatternInfo struct {
	Journey_pattern_ref string `json:"JourneyPatternRef,omitempty"`
	Journey_pattern_name string `json:"JourneyPatternName,omitempty"`
	Vehicle_mode VehicleMode `json:"VehicleMode,omitempty"`
	Route_ref string `json:"RouteRef,omitempty"`
	Published_line_name string `json:"PublishedLineName,omitempty"`
	Direction_name *string `json:"DirectionName,omitempty"`
}

type ServiceInfo struct {
}

type VehicleJourneyInfo struct {
}

type EstimatedInfo struct {
	headway_service *bool `json:"HeadwayService,omitempty"`
	origin_aimed_departure_time *string `json:"OriginAimedDepartureTime,omitempty"`
	destination_aimed_arrival_time *string `json:"DestinationAimedArrivalTime,omitempty"`
	first_or_last_journey *FirstOrLastJourneyEnum `json:"FirstOrLastJourney,omitempty"`
}

type JourneyProgressInfo struct {
}

type OperationalInfo struct {
	train_numbers []TrainNumber `json:"TrainNumbers,omitempty"`
	journey_parts []JourneyPart `json:"JourneyParts,omitempty"`
}

type DatedVehicleJourneyIndirectRef struct {
	origin_ref string `json:"OriginRef,omitempty"`
	aimed_departure_time string `json:"AimedDepartureTime,omitempty"`
	destination_ref string `json:"DestinationRef,omitempty"`
	aimed_arrival_time string `json:"AimedArrivalTime,omitempty"`
}

type FirstOrLastJourneyEnum int

const (
	FirstServiceOfDay FirstOrLastJourneyEnum = 0
	LastServiceOfDay FirstOrLastJourneyEnum = 1
	OtherService FirstOrLastJourneyEnum = 2
	Unspecified FirstOrLastJourneyEnum = 3
)


type Calls struct {
	recorded_calls *[]RecordedCall `json:"RecordedCalls,omitempty"`
	estimated_calls *[]EstimatedCall `json:"EstimatedCalls,omitempty"`
	is_complete_stop_sequence *bool `json:"IsCompleteStopSequence,omitempty"`
}

type RecordedCall struct {
	Stop_point_ref string `json:"StopPointRef,omitempty"`
	Order uint32 `json:"Order,omitempty"`
	Stop_point_name string `json:"StopPointName,omitempty"`
	Extra_call bool `json:"ExtraCall,omitempty"`
	Cancellation bool `json:"Cancellation,omitempty"`
	Occupancy Occupancy `json:"Occupancy,omitempty"`
	Platform_traversal bool `json:"PlatformTraversal,omitempty"`
	Disruption_group DisruptionGroup `json:"DisruptionGroup,omitempty"`
	Arrival ArrivalInfo `json:"Arrival,omitempty"`
	Departure DepartureInfo `json:"Departure,omitempty"`
	Expected_departure_occupancy ExpectedDepartureOccupancy `json:"ExpectedDepartureOccupancy,omitempty"`
	Expected_departure_capacity ExpectedDepartureCapacity `json:"ExpectedDepartureCapacity,omitempty"`
}

type EstimatedCall struct {
	Stop_point_ref string `json:"StopPointRef,omitempty"`
	Srder uint32 `json:"Order,omitempty"`
	Stop_point_name string `json:"StopPointName,omitempty"`
	Sxtra_call bool `json:"ExtraCall,omitempty"`
	Sancellation bool `json:"Cancellation,omitempty"`
	Occupancy Occupancy `json:"Occupancy,omitempty"`
	Platform_traversal bool `json:"PlatformTraversal,omitempty"`
	Destination_display string `json:"DestinationDisplay,omitempty"`
	Aimed_arrival_time string `json:"AimedArrivalTime,omitempty"`
	Expected_arrival_time string `json:"ExpectedArrivalTime,omitempty"`
	Arrival_status ArrivalStatus `json:"ArrivalStatus,omitempty"`
	Arrival_proximity_text []string `json:"ArrivalProximityText,omitempty"`
	Arrival_platform_name string `json:"ArrivalPlatformName,omitempty"`
	Arrival_stop_assignment string `json:"ArrivalStopAssignment,omitempty"`
	Aimed_quay_name string `json:"AimedQuayName,omitempty"`
	Aimed_departure_time string `json:"AimedDepartureTime,omitempty"`
	Expected_departure_time string `json:"ExpectedDepartureTime,omitempty"`
	Departure_status DepartureStatus `json:"DepartureStatus,omitempty"`
	Departure_platform_name string `json:"DeparturePlatformName,omitempty"`
	Departure_boarding_activity BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
}

type DisruptionGroup struct {
}

type TrainNumber struct {
	train_number_ref string `json:"TrainNumberRef,omitempty"`
}

type JourneyPart struct {
	journey_part_ref *string `json:"JourneyPartRef,omitempty"`
	train_number_ref *string `json:"TrainNumberRef,omitempty"`
}

type StopAssignment struct {
	aimed_quay_ref *string `json:"AimedQuayRef,omitempty"`
	aimed_quay_name *string `json:"AimedQuayName,omitempty"`
	expected_quay_ref *string `json:"ExpectedQuayRef,omitempty"`
}

type ExpectedDepartureCapacity struct {
	train_formation_reference_group *string `json:"TrainFormationReferenceGroup,omitempty"`
	passenger_category *string `json:"PassengerCategory,omitempty"`
	total_capacity *uint32 `json:"TotalCapacity,omitempty"`
	seating_capacity *uint32 `json:"SeatingCapacity,omitempty"`
	standing_capacity *uint32 `json:"StandingCapacity,omitempty"`
	pushchair_capacity *uint32 `json:"PushchairCapacity,omitempty"`
	wheelchair_place_capacity *uint32 `json:"WheelchairPlaceCapacity,omitempty"`
	pram_place_capacity *uint32 `json:"PramPlaceCapacity,omitempty"`
	bicycle_rack_capacity *uint32 `json:"BicycleRackCapacity,omitempty"`
}

type ExpectedDepartureOccupancy struct {
	passenger_category *string `json:"PassengerCategory,omitempty"`
	occupancy_level *Occupancy `json:"OccupancyLevel,omitempty"`
	occupancy_percentage *uint32 `json:"OccupancyPercentage,omitempty"`
	alighting_count *uint32 `json:"AlightingCount,omitempty"`
	boarding_count *uint32 `json:"BoardingCount,omitempty"`
	onboard_count *uint32 `json:"OnboardCount,omitempty"`
	group_reservations []GroupReservation `json:"GroupReservations,omitempty"`
}

type DepartureInfo struct {
	aimed_departure_time *string `json:"AimedDepartureTime,omitempty"`
	actual_departure_time *string `json:"ActualDepartureTime,omitempty"`
	expected_departure_time *string `json:"ExpectedDepartureTime,omitempty"`
	departure_status *DepartureStatus `json:"DepartureStatus,omitempty"`
	departure_platform_name *string `json:"DeparturePlatformName,omitempty"`
	departure_boarding_activity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
	expected_quay_ref *string `json:"ExpectedQuayRef,omitempty"`
}

type ArrivalInfo struct {
	aimed_arrival_time *string `json:"AimedArrivalTime,omitempty"`
	actual_arrival_time *string `json:"ActualArrivalTime,omitempty"`
	expected_arrival_time *string `json:"ExpectedArrivalTime,omitempty"`
	arrival_status *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	arrival_proximity_text []string `json:"ArrivalProximityText,omitempty"`
	arrival_platform_name *string `json:"ArrivalPlatformName,omitempty"`
	arrival_boarding_activity *BoardingActivity `json:"ArrivalBoardingActivity,omitempty"`
	arrival_stop_assignment *StopAssignment `json:"ArrivalStopAssignment,omitempty"`
}

type Arrival struct {
	aimed_arrival_time *string `json:"AimedArrivalTime,omitempty"`
	expected_arrival_time *string `json:"ExpectedArrivalTime,omitempty"`
	arrival_status *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	arrival_proximity_text *[]string `json:"ArrivalProximityText,omitempty"`
	arrival_platform_name *string `json:"ArrivalPlatformName,omitempty"`
	arrival_stop_assignment *string `json:"ArrivalStopAssignment,omitempty"`
	aimed_quay_name *string `json:"AimedQuayName,omitempty"`
}

type Departure struct {
	aimed_departure_time *string `json:"AimedDepartureTime,omitempty"`
	expected_departure_time *string `json:"ExpectedDepartureTime,omitempty"`
	departure_status *DepartureStatus `json:"DepartureStatus,omitempty"`
	departure_platform_name *string `json:"DeparturePlatformName,omitempty"`
	departure_boarding_activity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
}

type ExpectedCapacity struct {
}

type ExpectedOccupancy struct {
}

type GroupReservation struct {
	name_of_group string `json:"NameOfGroup,omitempty"`
	number_of_seats uint32 `json:"NumberOfSeats,omitempty"`
}

type MonitoredStopVisit struct {
	recorded_at_time string `json:"RecordedAtTime,omitempty"`
	item_identifier string `json:"ItemIdentifier,omitempty"`
	monitoring_ref string `json:"MonitoringRef,omitempty"`
	monitored_vehicle_journey MonitoredVehicleJourney `json:"MonitoredVehicleJourney,omitempty"`
}

type MonitoredStopVisitCancellation struct {
	recorded_at_time string `json:"RecordedAtTime,omitempty"`
	event_identity *string `json:"EventIdentity,omitempty"`
	monitoring_ref *string `json:"MonitoringRef,omitempty"`
	line_ref *string `json:"LineRef,omitempty"`
	vehicle_journey_ref *FramedVehicleJourneyRef `json:"VehicleJourneyRef,omitempty"`
	journey_pattern_info *JourneyPatternInfoGroup `json:"JourneyPatternInfo,omitempty"`
	message_reason *string `json:"MessageReason,omitempty"`
}

type MonitoredVehicleJourney struct {
	line_ref string `json:"LineRef,omitempty"`
	framed_vehicle_journey_ref FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef,omitempty"`
}

type JourneyPartInfo struct {
	journey_part_ref string `json:"JourneyPartRef,omitempty"`
	train_number_ref *string `json:"TrainNumberRef,omitempty"`
}

type MonitoredCall struct {
	stop_identity StopIdentity `json:"StopIdentity,omitempty"`
	order *uint32 `json:"Order,omitempty"`
	stop_point_name string `json:"StopPointName,omitempty"`
	vehicle_at_stop *bool `json:"VehicleAtStop,omitempty"`
	platform_traversal *bool `json:"PlatformTraversal,omitempty"`
	destination_display *string `json:"DestinationDisplay,omitempty"`
	disruption_group *DisruptionGroup `json:"DisruptionGroup,omitempty"`
	aimed_arrival_time *string `json:"AimedArrivalTime,omitempty"`
	actual_arrival_time *string `json:"ActualArrivalTime,omitempty"`
	expected_arrival_time *string `json:"ExpectedArrivalTime,omitempty"`
	arrival_status *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	arrival_proximity_text []string `json:"ArrivalProximityText,omitempty"`
	arrival_platform_name *string `json:"ArrivalPlatformName,omitempty"`
	aimed_quay_name *string `json:"AimedQuayName,omitempty"`
	aimed_departure_time *string `json:"AimedDepartureTime,omitempty"`
	actual_departure_time *string `json:"ActualDepartureTime,omitempty"`
	expected_departure_time *string `json:"ExpectedDepartureTime,omitempty"`
	departure_status *DepartureStatus `json:"DepartureStatus,omitempty"`
	departure_platform_name *string `json:"DeparturePlatformName,omitempty"`
	departure_boarding_activity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
	expected_departure_occupancy []ExpectedDepartureOccupancy `json:"ExpectedDepartureOccupancy,omitempty"`
	expected_departure_capacity []ExpectedDepartureCapacity `json:"ExpectedDepartureCapacity,omitempty"`
	aimed_headway_interval *uint32 `json:"AimedHeadwayInterval,omitempty"`
	expected_headway_interval *uint32 `json:"ExpectedHeadwayInterval,omitempty"`
	distance_from_stop *uint32 `json:"DistanceFromStop,omitempty"`
	number_of_stops_away *uint32 `json:"NumberOfStopsAway,omitempty"`
}

type StopIdentity struct {
	stop_point_code string `json:"StopPointCode,omitempty"`
}

type OnwardCall struct {
	stop_identity StopIdentity `json:"StopIdentity,omitempty"`
	order *uint32 `json:"Order,omitempty"`
	stop_point_name string `json:"StopPointName,omitempty"`
	vehicle_at_stop *bool `json:"VehicleAtStop,omitempty"`
	aimed_arrival_time *string `json:"AimedArrivalTime,omitempty"`
	expected_arrival_time *string `json:"ExpectedArrivalTime,omitempty"`
	arrival_status *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	arrival_proximity_text []string `json:"ArrivalProximityText,omitempty"`
	arrival_platform_name *string `json:"ArrivalPlatformName,omitempty"`
	aimed_departure_time *string `json:"AimedDepartureTime,omitempty"`
	expected_departure_time *string `json:"ExpectedDepartureTime,omitempty"`
	departure_status *DepartureStatus `json:"DepartureStatus,omitempty"`
	departure_platform_name *string `json:"DeparturePlatformName,omitempty"`
	departure_boarding_activity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
	aimed_headway_interval *uint32 `json:"AimedHeadwayInterval,omitempty"`
	expected_headway_interval *uint32 `json:"ExpectedHeadwayInterval,omitempty"`
	distance_from_stop *uint32 `json:"DistanceFromStop,omitempty"`
	number_of_stops_away *uint32 `json:"NumberOfStopsAway,omitempty"`
}

type VehicleJourneyInfoGroup struct {
	journey_end_names *JourneyEndNamesGroup `json:"JourneyEndNames,omitempty"`
	journey_info *string `json:"JourneyInfo,omitempty"`
	journey_note *string `json:"JourneyNote,omitempty"`
	headway_service *bool `json:"HeadwayService,omitempty"`
	origin_aimed_departure_time *string `json:"OriginAimedDepartureTime,omitempty"`
	destination_aimed_arrival_time *string `json:"DestinationAimedArrivalTime,omitempty"`
}

type JourneyEndNamesGroup struct {
	origin_ref *string `json:"OriginRef,omitempty"`
	origin_name *string `json:"OriginName,omitempty"`
	via *Via `json:"Via,omitempty"`
	destination_ref string `json:"DestinationRef,omitempty"`
	destination_name string `json:"DestinationName,omitempty"`
}

type JourneyPatternInfoGroup struct {
	journey_pattern_ref *string `json:"JourneyPatternRef,omitempty"`
	journey_pattern_name *string `json:"JourneyPatternName,omitempty"`
	vehicle_mode *VehicleMode `json:"VehicleMode,omitempty"`
	route_ref *string `json:"RouteRef,omitempty"`
	published_line_name string `json:"PublishedLineName,omitempty"`
	direction_name *string `json:"DirectionName,omitempty"`
}

type Via struct {
	place_ref *string `json:"PlaceRef,omitempty"`
	place_name *string `json:"PlaceName,omitempty"`
}

type JourneyProgressInfoGroup struct {
	monitored *bool `json:"Monitored,omitempty"`
	monitoring_error *MonitoringError `json:"MonitoringError,omitempty"`
	in_congestion *bool `json:"InCongestion,omitempty"`
	in_panic *bool `json:"InPanic,omitempty"`
	vehicle_location *LocationStructure `json:"VehicleLocation,omitempty"`
	bearing *float64 `json:"Bearing,omitempty"`
	occupancy *Occupancy `json:"Occupancy,omitempty"`
	delay *int32 `json:"Delay,omitempty"`
}

type LocationStructure struct {
}

type NotifyMonitoring struct {
	service_delivery_info ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
}

type VehicleActivityCancellation struct {
	recorded_at_time *string `json:"RecordedAtTime,omitempty"`
	event_identity *string `json:"EventIdentity,omitempty"`
	vehicle_monitoring_ref *string `json:"VehicleMonitoringRef,omitempty"`
	framed_vehicle_journey_ref *FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef,omitempty"`
	line_ref *string `json:"LineRef,omitempty"`
	journey_pattern_info *JourneyPatternInfoGroup `json:"JourneyPatternInfo,omitempty"`
	reasons []string `json:"Reasons,omitempty"`
}

type VehicleActivity struct {
	recorded_at_time *string `json:"RecordedAtTime,omitempty"`
	valid_until_time *string `json:"ValidUntilTime,omitempty"`
	item_identifier *string `json:"ItemIdentifier,omitempty"`
	vehicle_monitoring_ref *string `json:"VehicleMonitoringRef,omitempty"`
	progress_between_stops *ProgressBetweenStops `json:"ProgressBetweenStops,omitempty"`
	monitored_vehicle_journey MonitoredVehicleJourney `json:"MonitoredVehicleJourney,omitempty"`
	vehicle_activity_note *string `json:"VehicleActivityNote,omitempty"`
}

type ProgressBetweenStops struct {
	link_distance *int32 `json:"LinkDistance,omitempty"`
	percentage *int32 `json:"Percentage,omitempty"`
}

// type FramedVehicleJourneyRef struct {
// 	data_frame_ref *string `json:"DataFrameRef,omitempty"`
// 	dated_vehicle_journey_ref *string `json:"DatedVehicleJourneyRef,omitempty"`
// }

type FacilityCondition struct {
	facility Facility `json:"Facility,omitempty"`
	facility_ref *string `json:"FacilityRef,omitempty"`
	facility_status *FacilityStatus `json:"FacilityStatus,omitempty"`
	validity_period *ValidityPeriod `json:"ValidityPeriod,omitempty"`
}

type Facility struct {
	facility_code *string `json:"FacilityCode,omitempty"`
	description *string `json:"Description,omitempty"`
	facility_class *FacilityClass `json:"FacilityClass,omitempty"`
	validity_condition *ValidityCondition `json:"ValidityCondition,omitempty"`
	facility_location *FacilityLocation `json:"FacilityLocation,omitempty"`
	accessibility_assesment *AccessibilityAssessment `json:"AccessibilityAssesment,omitempty"`
	limitations *Limitations `json:"Limitations,omitempty"`
	suitabilities *Suitabilities `json:"Suitabilities,omitempty"`
}

type Limitations struct {
	wheelchair_access *bool `json:"WheelchairAccess,omitempty"`
	step_free_access *bool `json:"StepFreeAccess,omitempty"`
	lift_free_access *bool `json:"LiftFreeAccess,omitempty"`
}

type Suitabilities struct {
	suitability *[]Suitability `json:"Suitability,omitempty"`
}

type Suitability struct {
	suitable string `json:"Suitable,omitempty"`
	user_need *UserNeed `json:"UserNeed,omitempty"`
}

type UserNeed struct {
	mobility_need *string `json:"MobilityNeed,omitempty"`
}

type FacilityLocation struct {
	line_ref *string `json:"LineRef,omitempty"`
	stop_point_ref *string `json:"StopPointRef,omitempty"`
	vehicle_ref *string `json:"VehicleRef,omitempty"`
	stop_place_ref *string `json:"StopPlaceRef,omitempty"`
	stop_place_component_id *string `json:"StopPlaceComponentId,omitempty"`
	operator_ref *string `json:"OperatorRef,omitempty"`
}

type FacilityStatus struct {
	status *FacilityAvailability `json:"Status,omitempty"`
	description *string `json:"Description,omitempty"`
	accessibility_assesment []AccessibilityAssessment `json:"AccessibilityAssesment,omitempty"`
}

type AccessibilityAssessment struct {
}

type ValidityPeriod struct {
	start_time string `json:"StartTime,omitempty"`
	end_time string `json:"EndTime,omitempty"`
	end_time_precision *EndTimePrecision `json:"EndTimePrecision,omitempty"`
}

type Line struct {
	line_direction string `json:"LineDirection,omitempty"`
	line_ref string `json:"LineRef,omitempty"`
}

type ValidityCondition struct {
	period []ValidityPeriod `json:"Period,omitempty"`
}

type SiriServiceType int

const (
)





