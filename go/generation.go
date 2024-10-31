package siri_parser

type EstimatedTimetableDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	EstimatedJourneyVersionFrame EstimatedJourneyVersionFrame `json:"EstimatedJourneyVersionFrame,omitempty"`
}

type ConnectionMonitoringDelivery struct {
}

type FacilityMonitoringDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	FacilityCondition FacilityCondition `json:"FacilityCondition,omitempty"`
}

type GeneralMessageDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
}

type ProductionTimetableDelivery struct {
}

type SituationExchangeDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
}

type StopMonitoringDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	MonitoringRef *string `json:"MonitoringRef,omitempty"`
	MonitoredStopVisit *[]MonitoredStopVisit `json:"MonitoredStopVisit,omitempty"`
	MonitoredStopVisitCancellation *[]MonitoredStopVisitCancellation `json:"MonitoredStopVisitCancellation,omitempty"`
}

type VehicleMonitoringDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	VehicleActivity *[]VehicleActivity `json:"VehicleActivity,omitempty"`
	VehicleActivityCancellation *[]VehicleActivityCancellation `json:"VehicleActivityCancellation,omitempty"`
}

type SituationMonitoringDelivery struct {
}

type ArrivalStatus int

const (
	OnTime ArrivalStatus = 0
	Missed ArrivalStatus = 1
	Arrived ArrivalStatus = 2
	NotExpected ArrivalStatus = 3
	Delayed ArrivalStatus = 4
	Early ArrivalStatus = 5
	Cancelled ArrivalStatus = 6
	NoReport ArrivalStatus = 7
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
	OnTime DepartureStatus = 0
	Early DepartureStatus = 1
	Delayed DepartureStatus = 2
	Cancelled DepartureStatus = 3
	Arrived DepartureStatus = 4
	Departed DepartureStatus = 5
	NotExpected DepartureStatus = 6
	NoReport DepartureStatus = 7
)


type Occupancy int

const (
	Full Occupancy = 0
	SeatsAvailable Occupancy = 1
	StandingAvailable Occupancy = 2
	Unknown Occupancy = 3
	Empty Occupancy = 4
	ManySeatAvailable Occupancy = 5
	FewSeatAvailable Occupancy = 6
	StandingRoomOnly Occupancy = 7
	CrushStandingRoomOnly Occupancy = 8
	NotAcceptingPassengers Occupancy = 9
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


type Coordinates int

const (
	type LatLong struct {
	*float64
	*float64
}
	LatLong Coordinates = 0
	type GML struct {
	*string
}
	GML Coordinates = 1
)


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
	Calls *[]EstimatedCall `json:"Calls,omitempty"`
}

type EstimatedJourneyVersionFrame struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	VersionRef *string `json:"VersionRef,omitempty"`
	EstimatedVehicleJourney []EstimatedVehicleJourney `json:"EstimatedVehicleJourney,omitempty"`
}

type EstimatedVehicleJourney struct {
	LineRef string `json:"LineRef,omitempty"`
	PublishedLineName *string `json:"PublishedLineName,omitempty"`
	DirectionRef *string `json:"DirectionRef,omitempty"`
	JourneyIdentifier *JourneyIdentifier `json:"JourneyIdentifier,omitempty"`
	DatedVehiculeJourneyRef *string `json:"DatedVehiculeJourneyRef,omitempty"`
	Cancellation *string `json:"Cancellation,omitempty"`
	ExtraJourney *bool `json:"ExtraJourney,omitempty"`
	JourneyPatternName *string `json:"JourneyPatternName,omitempty"`
	JourneyPatternInfo *JourneyPatternInfo `json:"JourneyPatternInfo,omitempty"`
	VehicleMode *string `json:"VehicleMode,omitempty"`
	OriginRef *string `json:"OriginRef,omitempty"`
	OriginName *string `json:"OriginName,omitempty"`
	DestinationRef *string `json:"DestinationRef,omitempty"`
	DestinationName *string `json:"DestinationName,omitempty"`
	OperatorRef *string `json:"OperatorRef,omitempty"`
	ProductCategoryRef *string `json:"ProductCategoryRef,omitempty"`
	TrainNumbers *TrainNumbers `json:"TrainNumbers,omitempty"`
	VehiculeJourneyName *string `json:"VehiculeJourneyName,omitempty"`
	OriginAimedDepartureTime *string `json:"OriginAimedDepartureTime,omitempty"`
	DestinationAimedArrivalTime *string `json:"DestinationAimedArrivalTime,omitempty"`
	RecordedCalls *RecordedCalls `json:"RecordedCalls,omitempty"`
	EstimatedCalls *EstimatedCalls `json:"EstimatedCalls,omitempty"`
	FramedVehicleJourneyRef *FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef,omitempty"`
	DataSource *string `json:"DataSource,omitempty"`
	VehicleRef *string `json:"VehicleRef,omitempty"`
	AimedDepartureTime *string `json:"AimedDepartureTime,omitempty"`
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
	JourneyNote *string `json:"JourneyNote,omitempty"`
	HeadwayService *string `json:"HeadwayService,omitempty"`
	FirstOrLastJourney *string `json:"FirstOrLastJourney,omitempty"`
	DisruptionGroup *string `json:"DisruptionGroup,omitempty"`
	JourneyProgressInfo *string `json:"JourneyProgressInfo,omitempty"`
}

type FramedVehicleJourneyRef struct {
	DataFrameRef *string `json:"DataFrameRef,omitempty"`
	DatedVehicleJourneyRef *string `json:"DatedVehicleJourneyRef,omitempty"`
}

type RecordedCalls struct {
	Calls *[]RecordedCall `json:"Calls,omitempty"`
}

type ServiceDeliveryInfo struct {
	ResponseTimestamp *string `json:"ResponseTimestamp,omitempty"`
	ProducerRef string `json:"ProducerRef,omitempty"`
	RequestMessageIdentifier *string `json:"RequestMessageIdentifier,omitempty"`
	ResponseMessageRef *string `json:"ResponseMessageRef,omitempty"`
}

type TrainNumbers struct {
	TrainNumberRef *string `json:"TrainNumberRef,omitempty"`
}

type XxxDelivery struct {
	ResponseTimestamp string `json:"ResponseTimestamp,omitempty"`
	RequestMessageRef *string `json:"RequestMessageRef,omitempty"`
	SubscriberRef *string `json:"SubscriberRef,omitempty"`
	SubscriptionRef *string `json:"SubscriptionRef,omitempty"`
	Status *bool `json:"Status,omitempty"`
	ValidUntil *string `json:"ValidUntil,omitempty"`
	ShortestPossibleCycle *uint32 `json:"ShortestPossibleCycle,omitempty"`
}

type StopMonitoringNotification struct {
	StopMonitoringDelivery StopMonitoringDelivery `json:"StopMonitoringDelivery,omitempty"`
}

type ProductionTimetableNotification struct {
	ProductionTimetableDelivery ProductionTimetableDelivery `json:"ProductionTimetableDelivery,omitempty"`
}

type SituationMonitoringNotification struct {
	SituationMonitoringDelivery SituationMonitoringDelivery `json:"SituationMonitoringDelivery,omitempty"`
}

type GeneralMessageNotification struct {
	GeneralMessageDelivery GeneralMessageDelivery `json:"GeneralMessageDelivery,omitempty"`
}

type FacilityMonitoringNotification struct {
	FacilityMonitoringDelivery FacilityMonitoringDelivery `json:"FacilityMonitoringDelivery,omitempty"`
}

type ConnectionMonitoringNotification struct {
	ConnectionMonitoringDelivery ConnectionMonitoringDelivery `json:"ConnectionMonitoringDelivery,omitempty"`
}

type EstimatedTimetableNotification struct {
	EstimatedTimetableDelivery EstimatedTimetableDelivery `json:"EstimatedTimetableDelivery,omitempty"`
}

type SituationExchangeNotification struct {
	SituationExchangeDelivery SituationExchangeDelivery `json:"SituationExchangeDelivery,omitempty"`
}

type VehicleMonitoringNotification struct {
	VehicleMonitoringDelivery VehicleMonitoringDelivery `json:"VehicleMonitoringDelivery,omitempty"`
}

type NotifyEstimatedTimetable struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification EstimatedTimetableNotification `json:"Notification,omitempty"`
}

type NotifyStopMonitoring struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification StopMonitoringNotification `json:"Notification,omitempty"`
}

type NotifyConnectionMonitoring struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification ConnectionMonitoringNotification `json:"Notification,omitempty"`
}

type NotifyFacilityMonitoring struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification FacilityMonitoringNotification `json:"Notification,omitempty"`
}

type NotifyVechicleMonitoring struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification VehicleMonitoringNotification `json:"Notification,omitempty"`
}

type NotifySituationExchange struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification SituationExchangeNotification `json:"Notification,omitempty"`
}

type NotifyGeneralMessage struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification GeneralMessageNotification `json:"Notification,omitempty"`
}

type NotifyProductionTimetable struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
	Notification ProductionTimetableNotification `json:"Notification,omitempty"`
}

type JourneyEndNames struct {
}

type JourneyIdentifier int

const (
	type DatedVehicleJourneyRef struct {
	string
}
	DatedVehicleJourneyRef JourneyIdentifier = 0
	type EstimatedVehicleJourneyCode struct {
	string
}
	EstimatedVehicleJourneyCode JourneyIdentifier = 1
	type DatedVehicleJourneyIndirectRef struct {
	DatedVehicleJourneyIndirectRef
}
	DatedVehicleJourneyIndirectRef JourneyIdentifier = 2
)


type JourneyPatternInfo struct {
	JourneyPatternRef *string `json:"JourneyPatternRef,omitempty"`
	JourneyPatternName *string `json:"JourneyPatternName,omitempty"`
	VehicleMode *VehicleMode `json:"VehicleMode,omitempty"`
	RouteRef *string `json:"RouteRef,omitempty"`
	PublishedLineName string `json:"PublishedLineName,omitempty"`
	DirectionName *string `json:"DirectionName,omitempty"`
}

type ServiceInfo struct {
}

type VehicleJourneyInfo struct {
}

type EstimatedInfo struct {
	HeadwayService *bool `json:"HeadwayService,omitempty"`
	OriginAimedDepartureTime *string `json:"OriginAimedDepartureTime,omitempty"`
	DestinationAimedArrivalTime *string `json:"DestinationAimedArrivalTime,omitempty"`
	FirstOrLastJourney *FirstOrLastJourneyEnum `json:"FirstOrLastJourney,omitempty"`
}

type JourneyProgressInfo struct {
}

type OperationalInfo struct {
	TrainNumbers []TrainNumber `json:"TrainNumbers,omitempty"`
	JourneyParts []JourneyPart `json:"JourneyParts,omitempty"`
}

type DatedVehicleJourneyIndirectRef struct {
	OriginRef string `json:"OriginRef,omitempty"`
	AimedDepartureTime string `json:"AimedDepartureTime,omitempty"`
	DestinationRef string `json:"DestinationRef,omitempty"`
	AimedArrivalTime string `json:"AimedArrivalTime,omitempty"`
}

type EstimatedVehicleJourney struct {
	LineRef string `json:"LineRef,omitempty"`
	DirectionRef string `json:"DirectionRef,omitempty"`
	JourneyIdentifier JourneyIdentifier `json:"JourneyIdentifier,omitempty"`
	ExtraJourney *bool `json:"ExtraJourney,omitempty"`
	Cancellation *bool `json:"Cancellation,omitempty"`
	JourneyPatternInfo *JourneyPatternInfo `json:"JourneyPatternInfo,omitempty"`
	JourneyEndNames *JourneyEndNames `json:"JourneyEndNames,omitempty"`
	VehicleJourneyInfo *VehicleJourneyInfo `json:"VehicleJourneyInfo,omitempty"`
	ServiceInfo *ServiceInfo `json:"ServiceInfo,omitempty"`
	EstimatedInfo EstimatedInfo `json:"EstimatedInfo,omitempty"`
	DisruptionGroup *DisruptionGroup `json:"DisruptionGroup,omitempty"`
	JourneyProgressInfo *JourneyProgressInfo `json:"JourneyProgressInfo,omitempty"`
	OperationalInfo *OperationalInfo `json:"OperationalInfo,omitempty"`
	Calls Calls `json:"Calls,omitempty"`
}

type FirstOrLastJourneyEnum int

const (
	FirstServiceOfDay FirstOrLastJourneyEnum = 0
	LastServiceOfDay FirstOrLastJourneyEnum = 1
	OtherService FirstOrLastJourneyEnum = 2
	Unspecified FirstOrLastJourneyEnum = 3
)


type Calls struct {
	RecordedCalls *[]RecordedCall `json:"RecordedCalls,omitempty"`
	EstimatedCalls *[]EstimatedCall `json:"EstimatedCalls,omitempty"`
	IsCompleteStopSequence *bool `json:"IsCompleteStopSequence,omitempty"`
}

type RecordedCall struct {
	StopPointRef string `json:"StopPointRef,omitempty"`
	Order uint32 `json:"Order,omitempty"`
	StopPointName *string `json:"StopPointName,omitempty"`
	ExtraCall *bool `json:"ExtraCall,omitempty"`
	Cancellation *bool `json:"Cancellation,omitempty"`
	Occupancy *Occupancy `json:"Occupancy,omitempty"`
	PlatformTraversal *bool `json:"PlatformTraversal,omitempty"`
	DisruptionGroup *DisruptionGroup `json:"DisruptionGroup,omitempty"`
	Arrival *ArrivalInfo `json:"Arrival,omitempty"`
	Departure *DepartureInfo `json:"Departure,omitempty"`
	ExpectedDepartureOccupancy *ExpectedDepartureOccupancy `json:"ExpectedDepartureOccupancy,omitempty"`
	ExpectedDepartureCapacity *ExpectedDepartureCapacity `json:"ExpectedDepartureCapacity,omitempty"`
}

type EstimatedCall struct {
	StopPointRef *string `json:"StopPointRef,omitempty"`
	Order *uint32 `json:"Order,omitempty"`
	StopPointName *string `json:"StopPointName,omitempty"`
	ExtraCall *bool `json:"ExtraCall,omitempty"`
	Cancellation *bool `json:"Cancellation,omitempty"`
	Occupancy *Occupancy `json:"Occupancy,omitempty"`
	PlatformTraversal *bool `json:"PlatformTraversal,omitempty"`
	DestinationDisplay *string `json:"DestinationDisplay,omitempty"`
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
	ExpectedArrivalTime *string `json:"ExpectedArrivalTime,omitempty"`
	ArrivalStatus *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	ArrivalProximityText *[]string `json:"ArrivalProximityText,omitempty"`
	ArrivalPlatformName *string `json:"ArrivalPlatformName,omitempty"`
	ArrivalStopAssignment *string `json:"ArrivalStopAssignment,omitempty"`
	AimedQuayName *string `json:"AimedQuayName,omitempty"`
	AimedDepartureTime *string `json:"AimedDepartureTime,omitempty"`
	ExpectedDepartureTime *string `json:"ExpectedDepartureTime,omitempty"`
	DepartureStatus *DepartureStatus `json:"DepartureStatus,omitempty"`
	DeparturePlatformName *string `json:"DeparturePlatformName,omitempty"`
	DepartureBoardingActivity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
}

type DisruptionGroup struct {
}

type TrainNumber struct {
	TrainNumberRef string `json:"TrainNumberRef,omitempty"`
}

type JourneyPart struct {
	JourneyPartRef *string `json:"JourneyPartRef,omitempty"`
	TrainNumberRef *string `json:"TrainNumberRef,omitempty"`
}

type StopAssignment struct {
	AimedQuayRef *string `json:"AimedQuayRef,omitempty"`
	AimedQuayName *string `json:"AimedQuayName,omitempty"`
	ExpectedQuayRef *string `json:"ExpectedQuayRef,omitempty"`
}

type ExpectedDepartureCapacity struct {
	TrainFormationReferenceGroup *string `json:"TrainFormationReferenceGroup,omitempty"`
	PassengerCategory *string `json:"PassengerCategory,omitempty"`
	TotalCapacity *uint32 `json:"TotalCapacity,omitempty"`
	SeatingCapacity *uint32 `json:"SeatingCapacity,omitempty"`
	StandingCapacity *uint32 `json:"StandingCapacity,omitempty"`
	PushchairCapacity *uint32 `json:"PushchairCapacity,omitempty"`
	WheelchairPlaceCapacity *uint32 `json:"WheelchairPlaceCapacity,omitempty"`
	PramPlaceCapacity *uint32 `json:"PramPlaceCapacity,omitempty"`
	BicycleRackCapacity *uint32 `json:"BicycleRackCapacity,omitempty"`
}

type ExpectedDepartureOccupancy struct {
	PassengerCategory *string `json:"PassengerCategory,omitempty"`
	OccupancyLevel *Occupancy `json:"OccupancyLevel,omitempty"`
	OccupancyPercentage *uint32 `json:"OccupancyPercentage,omitempty"`
	AlightingCount *uint32 `json:"AlightingCount,omitempty"`
	BoardingCount *uint32 `json:"BoardingCount,omitempty"`
	OnboardCount *uint32 `json:"OnboardCount,omitempty"`
	GroupReservations []GroupReservation `json:"GroupReservations,omitempty"`
}

type DepartureInfo struct {
	AimedDepartureTime *string `json:"AimedDepartureTime,omitempty"`
	ActualDepartureTime *string `json:"ActualDepartureTime,omitempty"`
	ExpectedDepartureTime *string `json:"ExpectedDepartureTime,omitempty"`
	DepartureStatus *DepartureStatus `json:"DepartureStatus,omitempty"`
	DeparturePlatformName *string `json:"DeparturePlatformName,omitempty"`
	DepartureBoardingActivity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
	ExpectedQuayRef *string `json:"ExpectedQuayRef,omitempty"`
}

type ArrivalInfo struct {
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
	ActualArrivalTime *string `json:"ActualArrivalTime,omitempty"`
	ExpectedArrivalTime *string `json:"ExpectedArrivalTime,omitempty"`
	ArrivalStatus *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	ArrivalProximityText []string `json:"ArrivalProximityText,omitempty"`
	ArrivalPlatformName *string `json:"ArrivalPlatformName,omitempty"`
	ArrivalBoardingActivity *BoardingActivity `json:"ArrivalBoardingActivity,omitempty"`
	ArrivalStopAssignment *StopAssignment `json:"ArrivalStopAssignment,omitempty"`
}

type Arrival struct {
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
	ExpectedArrivalTime *string `json:"ExpectedArrivalTime,omitempty"`
	ArrivalStatus *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	ArrivalProximityText *[]string `json:"ArrivalProximityText,omitempty"`
	ArrivalPlatformName *string `json:"ArrivalPlatformName,omitempty"`
	ArrivalStopAssignment *string `json:"ArrivalStopAssignment,omitempty"`
	AimedQuayName *string `json:"AimedQuayName,omitempty"`
}

type Departure struct {
	AimedDepartureTime *string `json:"AimedDepartureTime,omitempty"`
	ExpectedDepartureTime *string `json:"ExpectedDepartureTime,omitempty"`
	DepartureStatus *DepartureStatus `json:"DepartureStatus,omitempty"`
	DeparturePlatformName *string `json:"DeparturePlatformName,omitempty"`
	DepartureBoardingActivity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
}

type ExpectedCapacity struct {
}

type ExpectedOccupancy struct {
}

type GroupReservation struct {
	NameOfGroup string `json:"NameOfGroup,omitempty"`
	NumberOfSeats uint32 `json:"NumberOfSeats,omitempty"`
}

type MonitoredStopVisit struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	ItemIdentifier string `json:"ItemIdentifier,omitempty"`
	MonitoringRef string `json:"MonitoringRef,omitempty"`
	MonitoredVehicleJourney MonitoredVehicleJourney `json:"MonitoredVehicleJourney,omitempty"`
}

type MonitoredStopVisitCancellation struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	EventIdentity *string `json:"EventIdentity,omitempty"`
	MonitoringRef *string `json:"MonitoringRef,omitempty"`
	LineRef *string `json:"LineRef,omitempty"`
	VehicleJourneyRef *FramedVehicleJourneyRef `json:"VehicleJourneyRef,omitempty"`
	JourneyPatternInfo *JourneyPatternInfoGroup `json:"JourneyPatternInfo,omitempty"`
	MessageReason *string `json:"MessageReason,omitempty"`
}

type MonitoredVehicleJourney struct {
	LineRef string `json:"LineRef,omitempty"`
	FramedVehicleJourneyRef FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef,omitempty"`
}

type JourneyPartInfo struct {
	JourneyPartRef string `json:"JourneyPartRef,omitempty"`
	TrainNumberRef *string `json:"TrainNumberRef,omitempty"`
}

type MonitoredCall struct {
	StopIdentity StopIdentity `json:"StopIdentity,omitempty"`
	Order *uint32 `json:"Order,omitempty"`
	StopPointName string `json:"StopPointName,omitempty"`
	VehicleAtStop *bool `json:"VehicleAtStop,omitempty"`
	PlatformTraversal *bool `json:"PlatformTraversal,omitempty"`
	DestinationDisplay *string `json:"DestinationDisplay,omitempty"`
	DisruptionGroup *DisruptionGroup `json:"DisruptionGroup,omitempty"`
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
	ActualArrivalTime *string `json:"ActualArrivalTime,omitempty"`
	ExpectedArrivalTime *string `json:"ExpectedArrivalTime,omitempty"`
	ArrivalStatus *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	ArrivalProximityText []string `json:"ArrivalProximityText,omitempty"`
	ArrivalPlatformName *string `json:"ArrivalPlatformName,omitempty"`
	AimedQuayName *string `json:"AimedQuayName,omitempty"`
	AimedDepartureTime *string `json:"AimedDepartureTime,omitempty"`
	ActualDepartureTime *string `json:"ActualDepartureTime,omitempty"`
	ExpectedDepartureTime *string `json:"ExpectedDepartureTime,omitempty"`
	DepartureStatus *DepartureStatus `json:"DepartureStatus,omitempty"`
	DeparturePlatformName *string `json:"DeparturePlatformName,omitempty"`
	DepartureBoardingActivity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
	ExpectedDepartureOccupancy []ExpectedDepartureOccupancy `json:"ExpectedDepartureOccupancy,omitempty"`
	ExpectedDepartureCapacity []ExpectedDepartureCapacity `json:"ExpectedDepartureCapacity,omitempty"`
	AimedHeadwayInterval *uint32 `json:"AimedHeadwayInterval,omitempty"`
	ExpectedHeadwayInterval *uint32 `json:"ExpectedHeadwayInterval,omitempty"`
	DistanceFromStop *uint32 `json:"DistanceFromStop,omitempty"`
	NumberOfStopsAway *uint32 `json:"NumberOfStopsAway,omitempty"`
}

type StopIdentity struct {
	StopPointCode string `json:"StopPointCode,omitempty"`
}

type OnwardCall struct {
	StopIdentity StopIdentity `json:"StopIdentity,omitempty"`
	Order *uint32 `json:"Order,omitempty"`
	StopPointName string `json:"StopPointName,omitempty"`
	VehicleAtStop *bool `json:"VehicleAtStop,omitempty"`
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
	ExpectedArrivalTime *string `json:"ExpectedArrivalTime,omitempty"`
	ArrivalStatus *ArrivalStatus `json:"ArrivalStatus,omitempty"`
	ArrivalProximityText []string `json:"ArrivalProximityText,omitempty"`
	ArrivalPlatformName *string `json:"ArrivalPlatformName,omitempty"`
	AimedDepartureTime *string `json:"AimedDepartureTime,omitempty"`
	ExpectedDepartureTime *string `json:"ExpectedDepartureTime,omitempty"`
	DepartureStatus *DepartureStatus `json:"DepartureStatus,omitempty"`
	DeparturePlatformName *string `json:"DeparturePlatformName,omitempty"`
	DepartureBoardingActivity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
	AimedHeadwayInterval *uint32 `json:"AimedHeadwayInterval,omitempty"`
	ExpectedHeadwayInterval *uint32 `json:"ExpectedHeadwayInterval,omitempty"`
	DistanceFromStop *uint32 `json:"DistanceFromStop,omitempty"`
	NumberOfStopsAway *uint32 `json:"NumberOfStopsAway,omitempty"`
}

type VehicleJourneyInfoGroup struct {
	JourneyEndNames *JourneyEndNamesGroup `json:"JourneyEndNames,omitempty"`
	JourneyInfo *string `json:"JourneyInfo,omitempty"`
	JourneyNote *string `json:"JourneyNote,omitempty"`
	HeadwayService *bool `json:"HeadwayService,omitempty"`
	OriginAimedDepartureTime *string `json:"OriginAimedDepartureTime,omitempty"`
	DestinationAimedArrivalTime *string `json:"DestinationAimedArrivalTime,omitempty"`
}

type JourneyEndNamesGroup struct {
	OriginRef *string `json:"OriginRef,omitempty"`
	OriginName *string `json:"OriginName,omitempty"`
	Via *Via `json:"Via,omitempty"`
	DestinationRef string `json:"DestinationRef,omitempty"`
	DestinationName string `json:"DestinationName,omitempty"`
}

type JourneyPatternInfoGroup struct {
	JourneyPatternRef *string `json:"JourneyPatternRef,omitempty"`
	JourneyPatternName *string `json:"JourneyPatternName,omitempty"`
	VehicleMode *VehicleMode `json:"VehicleMode,omitempty"`
	RouteRef *string `json:"RouteRef,omitempty"`
	PublishedLineName string `json:"PublishedLineName,omitempty"`
	DirectionName *string `json:"DirectionName,omitempty"`
}

type Via struct {
	PlaceRef *string `json:"PlaceRef,omitempty"`
	PlaceName *string `json:"PlaceName,omitempty"`
}

type JourneyProgressInfoGroup struct {
	Monitored *bool `json:"Monitored,omitempty"`
	MonitoringError *MonitoringError `json:"MonitoringError,omitempty"`
	InCongestion *bool `json:"InCongestion,omitempty"`
	InPanic *bool `json:"InPanic,omitempty"`
	VehicleLocation *LocationStructure `json:"VehicleLocation,omitempty"`
	Bearing *float64 `json:"Bearing,omitempty"`
	Occupancy *Occupancy `json:"Occupancy,omitempty"`
	Delay *int32 `json:"Delay,omitempty"`
}

type LocationStructure struct {
}

type NotifyMonitoring struct {
	ServiceDeliveryInfo ServiceDeliveryInfo `json:"ServiceDeliveryInfo,omitempty"`
}

type VehicleActivityCancellation struct {
	RecordedAtTime *string `json:"RecordedAtTime,omitempty"`
	EventIdentity *string `json:"EventIdentity,omitempty"`
	VehicleMonitoringRef *string `json:"VehicleMonitoringRef,omitempty"`
	FramedVehicleJourneyRef *FramedVehicleJourneyRef `json:"FramedVehicleJourneyRef,omitempty"`
	LineRef *string `json:"LineRef,omitempty"`
	JourneyPatternInfo *JourneyPatternInfoGroup `json:"JourneyPatternInfo,omitempty"`
	Reasons []string `json:"Reasons,omitempty"`
}

type VehicleActivity struct {
	RecordedAtTime *string `json:"RecordedAtTime,omitempty"`
	ValidUntilTime *string `json:"ValidUntilTime,omitempty"`
	ItemIdentifier *string `json:"ItemIdentifier,omitempty"`
	VehicleMonitoringRef *string `json:"VehicleMonitoringRef,omitempty"`
	ProgressBetweenStops *ProgressBetweenStops `json:"ProgressBetweenStops,omitempty"`
	MonitoredVehicleJourney MonitoredVehicleJourney `json:"MonitoredVehicleJourney,omitempty"`
	VehicleActivityNote *string `json:"VehicleActivityNote,omitempty"`
}

type ProgressBetweenStops struct {
	LinkDistance *int32 `json:"LinkDistance,omitempty"`
	Percentage *int32 `json:"Percentage,omitempty"`
}

type FramedVehicleJourneyRef struct {
	DataFrameRef *string `json:"DataFrameRef,omitempty"`
	DatedVehicleJourneyRef *string `json:"DatedVehicleJourneyRef,omitempty"`
}

type FacilityCondition struct {
	Facility Facility `json:"Facility,omitempty"`
	FacilityRef *string `json:"FacilityRef,omitempty"`
	FacilityStatus *FacilityStatus `json:"FacilityStatus,omitempty"`
	ValidityPeriod *ValidityPeriod `json:"ValidityPeriod,omitempty"`
}

type Facility struct {
	FacilityCode *string `json:"FacilityCode,omitempty"`
	Description *string `json:"Description,omitempty"`
	FacilityClass *FacilityClass `json:"FacilityClass,omitempty"`
	ValidityCondition *ValidityCondition `json:"ValidityCondition,omitempty"`
	FacilityLocation *FacilityLocation `json:"FacilityLocation,omitempty"`
	AccessibilityAssesment *AccessibilityAssessment `json:"AccessibilityAssesment,omitempty"`
	Limitations *Limitations `json:"Limitations,omitempty"`
	Suitabilities *Suitabilities `json:"Suitabilities,omitempty"`
}

type Limitations struct {
	WheelchairAccess *bool `json:"WheelchairAccess,omitempty"`
	StepFreeAccess *bool `json:"StepFreeAccess,omitempty"`
	LiftFreeAccess *bool `json:"LiftFreeAccess,omitempty"`
}

type Suitabilities struct {
	Suitability *[]Suitability `json:"Suitability,omitempty"`
}

type Suitability struct {
	Suitable string `json:"Suitable,omitempty"`
	UserNeed *UserNeed `json:"UserNeed,omitempty"`
}

type UserNeed struct {
	MobilityNeed *string `json:"MobilityNeed,omitempty"`
}

type FacilityLocation struct {
	LineRef *string `json:"LineRef,omitempty"`
	StopPointRef *string `json:"StopPointRef,omitempty"`
	VehicleRef *string `json:"VehicleRef,omitempty"`
	StopPlaceRef *string `json:"StopPlaceRef,omitempty"`
	StopPlaceComponentId *string `json:"StopPlaceComponentId,omitempty"`
	OperatorRef *string `json:"OperatorRef,omitempty"`
}

type FacilityStatus struct {
	Status *FacilityAvailability `json:"Status,omitempty"`
	Description *string `json:"Description,omitempty"`
	AccessibilityAssesment []AccessibilityAssessment `json:"AccessibilityAssesment,omitempty"`
}

type AccessibilityAssessment struct {
}

type ValidityPeriod struct {
	StartTime string `json:"StartTime,omitempty"`
	EndTime string `json:"EndTime,omitempty"`
	EndTimePrecision *EndTimePrecision `json:"EndTimePrecision,omitempty"`
}

type Line struct {
	LineDirection string `json:"LineDirection,omitempty"`
	LineRef string `json:"LineRef,omitempty"`
}

type ValidityCondition struct {
	Period []ValidityPeriod `json:"Period,omitempty"`
}

type SiriServiceType int

const (
)

type NotifyProductionTimetable Body 
type NotifyEstimatedTimetable Body 
type NotifyStopMonitoring Body 
type NotifyVechicleMonitoring Body 
type NotifyConnectionMonitoring Body 
type NotifyGeneralMessage Body 
type NotifyFacilityMonitoring Body 
type NotifySituationExchange Body 

type Envelope struct {
	Body Body `json:"Body,omitempty"`
}

