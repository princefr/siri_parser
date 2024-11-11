package siri_parser

type EstimatedTimetableDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	EstimatedJourneyVersionFrame EstimatedJourneyVersionFrame `json:"EstimatedJourneyVersionFrame,omitempty"`
}

type ConnectionMonitoringFeederDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	MonitoredFeederArrival MonitoredFeederArrival `json:"MonitoredFeederArrival,omitempty"`
	MonitoredFeederArrivalCancellation MonitoredFeederArrivalCancellation `json:"MonitoredFeederArrivalCancellation,omitempty"`
}

type FacilityMonitoringDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	FacilityCondition FacilityCondition `json:"FacilityCondition,omitempty"`
}

type GeneralMessageDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	InfoMessage *InfoMessage `json:"InfoMessage,omitempty"`
	InfoMessageCancellation *InfoMessageCancellation `json:"InfoMessageCancellation,omitempty"`
}

type ProductionTimetableDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	DatedTimetableVersionFrame DatedTimetableVersionFrame `json:"DatedTimetableVersionFrame,omitempty"`
}

type SituationExchangeDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	Situations []PtSituationElement `json:"Situations,omitempty"`
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

type ConnectionMonitoringDistributorDelivery struct {
	Leader XxxDelivery `json:"Leader,omitempty"`
	WaitProlongedDeparture *WaitProlongedDeparture `json:"WaitProlongedDeparture,omitempty"`
	StoppingPositionChangeDeparture *StoppingPositionChangeDeparture `json:"StoppingPositionChangeDeparture,omitempty"`
	DistributorDepartureCancellation *DistributorDepartureCancellation `json:"DistributorDepartureCancellation,omitempty"`
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


type Sensivity int

const (
	High Sensivity = 0
	Medium Sensivity = 1
	Low Sensivity = 2
)


type Audience int

const (
	Public Audience = 0
	EmergencyService Audience = 1
	Authorities Audience = 2
	TransportOperators Audience = 3
)


type ScopeType int

const (
	General ScopeType = 0
	Operator ScopeType = 1
	Network ScopeType = 2
	Route ScopeType = 3
	Line ScopeType = 4
	Place ScopeType = 5
	StopPlace ScopeType = 6
	StopPoint ScopeType = 7
	VehicleJourney ScopeType = 8
)


type VerificationStatus int

const (
	Unknown VerificationStatus = 0
	Unverified VerificationStatus = 1
	Verified VerificationStatus = 2
)


type ProgressStatus int

const (
	Open ProgressStatus = 0
	Published ProgressStatus = 1
	Closed ProgressStatus = 2
)


type QualityIndex int

const (
	Certain QualityIndex = 0
	VeryReliable QualityIndex = 1
	Reliable QualityIndex = 2
	Unreliable QualityIndex = 3
	Unconfirmed QualityIndex = 4
)


type SourceType int

const (
	DirectReport SourceType = 0
	Email SourceType = 1
	Phone SourceType = 2
	Post SourceType = 3
	Feed SourceType = 4
	Radio SourceType = 5
	TV SourceType = 6
	Web SourceType = 7
	Text SourceType = 8
	Other SourceType = 9
)


type EndTimeStatus int

const (
	Undefined EndTimeStatus = 0
	LongTerm EndTimeStatus = 1
	ShortTerm EndTimeStatus = 2
)


type Condition int

const (
	Unknown Condition = 0
	Altered Condition = 1
	Cancelled Condition = 2
	Delayed Condition = 3
	Diverted Condition = 4
	NoService Condition = 5
	Disrupted Condition = 6
	AdditionalService Condition = 7
	SpecialService Condition = 8
	OnTime Condition = 9
	NormalService Condition = 10
	IntermittentService Condition = 11
	ExtendedService Condition = 12
	SplittingTrain Condition = 13
	ReplacementTransport Condition = 14
	ArrivesEarly Condition = 15
	ShuttleService Condition = 16
	ReplacementService Condition = 17
	UndefinedServiceInformation Condition = 18
)


type Severity int

const (
	Unknown Severity = 0
	Slight Severity = 1
	Normal Severity = 2
	Severe Severity = 3
	NoImpact Severity = 4
	Undefined Severity = 5
)


type ActionStatus int

const (
	Unknown ActionStatus = 0
	Active ActionStatus = 1
	Inactive ActionStatus = 2
)


type AccessMode int

const (
	Foot AccessMode = 0
	Bicycle AccessMode = 1
	Car AccessMode = 2
	Taxi AccessMode = 3
	Shuttle AccessMode = 4
)


type VehicleFeature int

const (
	ShortTrain VehicleFeature = 0
	LongTrain VehicleFeature = 1
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
	ConnectionMonitoringFeederDelivery *ConnectionMonitoringFeederDelivery `json:"ConnectionMonitoringFeederDelivery,omitempty"`
	ConnectionMonitoringDistributorDelivery *ConnectionMonitoringDistributorDelivery `json:"ConnectionMonitoringDistributorDelivery,omitempty"`
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

type MonitoredFeederArrival struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	ItemIdentifier *string `json:"ItemIdentifier,omitempty"`
	InterchangeRef *string `json:"InterchangeRef,omitempty"`
	ConnectionLinkRef *string `json:"ConnectionLinkRef,omitempty"`
	StopPointRef *string `json:"StopPointRef,omitempty"`
	Order *uint32 `json:"Order,omitempty"`
	StropPointName *string `json:"StropPointName,omitempty"`
	ClearDownRef *string `json:"ClearDownRef,omitempty"`
	FeederJourney FeederJourney `json:"FeederJourney,omitempty"`
	VehicleAtStop *bool `json:"VehicleAtStop,omitempty"`
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
	ExpectedArrivalTime *string `json:"ExpectedArrivalTime,omitempty"`
	ArrivalPlatformTime *string `json:"ArrivalPlatformTime,omitempty"`
	NumberOfTransferPassengers *uint32 `json:"NumberOfTransferPassengers,omitempty"`
}

type MonitoredFeederArrivalCancellation struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	ItemRef *string `json:"ItemRef,omitempty"`
	InterchangeRef *string `json:"InterchangeRef,omitempty"`
	ConnectionLinkRef *string `json:"ConnectionLinkRef,omitempty"`
	StopPointRef *string `json:"StopPointRef,omitempty"`
	Order *uint32 `json:"Order,omitempty"`
	StropPointName *string `json:"StropPointName,omitempty"`
	JourneyInfo *JourneyInfo `json:"JourneyInfo,omitempty"`
	Reason *string `json:"Reason,omitempty"`
}

type WaitProlongedDeparture struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	DistributorInfo *DistributorInfo `json:"DistributorInfo,omitempty"`
	ExpectedDepartureTime *string `json:"ExpectedDepartureTime,omitempty"`
}

type StoppingPositionChangeDeparture struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	DistributorInfo *DistributorInfo `json:"DistributorInfo,omitempty"`
	ChangeNote *string `json:"ChangeNote,omitempty"`
	NewLocation *LocationStructure `json:"NewLocation,omitempty"`
}

type DistributorDepartureCancellation struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	DistributorInfo *DistributorInfo `json:"DistributorInfo,omitempty"`
	Reason *string `json:"Reason,omitempty"`
}

type DistributorInfo struct {
	InterchangeRef string `json:"InterchangeRef,omitempty"`
	ConnectionLinkRef string `json:"ConnectionLinkRef,omitempty"`
	StopPointRef string `json:"StopPointRef,omitempty"`
	DistributorOrder *uint32 `json:"DistributorOrder,omitempty"`
	DistributorJourney *ConnectingJourney `json:"DistributorJourney,omitempty"`
	FeederVehicleJourneyRef *FramedVehicleJourneyRef `json:"FeederVehicleJourneyRef,omitempty"`
}

type ConnectingJourney struct {
	LineRef *string `json:"LineRef,omitempty"`
	FramedJourneyRef FramedVehicleJourneyRef `json:"FramedJourneyRef,omitempty"`
	JourneyPatternInfo *JourneyPatternInfoGroup `json:"JourneyPatternInfo,omitempty"`
	VehicleJourneyInfo *VehicleJourneyInfoGroup `json:"VehicleJourneyInfo,omitempty"`
	DistruptuinGroup *DisruptionGroup `json:"DistruptuinGroup,omitempty"`
	Monitored *bool `json:"Monitored,omitempty"`
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
}

type FeederJourney struct {
	LineRef *string `json:"LineRef,omitempty"`
	DirectionRef *string `json:"DirectionRef,omitempty"`
	FramedJourneyRef FramedVehicleJourneyRef `json:"FramedJourneyRef,omitempty"`
	Monitored *bool `json:"Monitored,omitempty"`
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
}

type JourneyInfo struct {
	LineRef *string `json:"LineRef,omitempty"`
	DirectionRef *string `json:"DirectionRef,omitempty"`
	JourneyPatternRef *string `json:"JourneyPatternRef,omitempty"`
	JourneyPatternName *string `json:"JourneyPatternName,omitempty"`
	VehicleMode *VehicleMode `json:"VehicleMode,omitempty"`
	RouteRef *string `json:"RouteRef,omitempty"`
	PublishedLineName *string `json:"PublishedLineName,omitempty"`
	GroupOfLinesRef *string `json:"GroupOfLinesRef,omitempty"`
	DirectionName *string `json:"DirectionName,omitempty"`
	Reason *string `json:"Reason,omitempty"`
}

type PtSituationElement struct {
	CreationTime string `json:"CreationTime,omitempty"`
	SituationBasedIdentityGroup SituationBasedIdentityGroup `json:"SituationBasedIdentityGroup,omitempty"`
	Source *SituationSource `json:"Source,omitempty"`
	VersionedAtTime *string `json:"VersionedAtTime,omitempty"`
	Verification *PtSituationBodyGroup `json:"Verification,omitempty"`
	ValidityPeriod *ValidityPeriod `json:"ValidityPeriod,omitempty"`
	ReasonName *string `json:"ReasonName,omitempty"`
	Severity *string `json:"Severity,omitempty"`
	Priority *uint32 `json:"Priority,omitempty"`
	Sensivity *Sensivity `json:"Sensivity,omitempty"`
	Audience *Audience `json:"Audience,omitempty"`
	ScopeType *ScopeType `json:"ScopeType,omitempty"`
	Planned *bool `json:"Planned,omitempty"`
	Keywords *string `json:"Keywords,omitempty"`
	Summary *string `json:"Summary,omitempty"`
	Description *string `json:"Description,omitempty"`
	Detail *string `json:"Detail,omitempty"`
	Advice *string `json:"Advice,omitempty"`
	Internal *string `json:"Internal,omitempty"`
	Affects []Affect `json:"Affects,omitempty"`
	Consequences []PtConsequence `json:"Consequences,omitempty"`
	PublishingActions *PublishingActions `json:"PublishingActions,omitempty"`
}

type SituationBasedIdentityGroup struct {
	CountryRef string `json:"CountryRef,omitempty"`
	ParticipantRef string `json:"ParticipantRef,omitempty"`
	SituationNumber string `json:"SituationNumber,omitempty"`
	Version *string `json:"Version,omitempty"`
}

type SituationSource struct {
	Name *string `json:"Name,omitempty"`
	SourceType *SourceType `json:"SourceType,omitempty"`
}

type PtSituationBodyGroup struct {
	Verification *VerificationStatus `json:"Verification,omitempty"`
	Progress *ProgressStatus `json:"Progress,omitempty"`
	QualityIndex *QualityIndex `json:"QualityIndex,omitempty"`
	Publication *string `json:"Publication,omitempty"`
}

type ReasonGroup struct {
	ReasonName *string `json:"ReasonName,omitempty"`
	Severity *string `json:"Severity,omitempty"`
	Priority *uint32 `json:"Priority,omitempty"`
	Sensivity *Sensivity `json:"Sensivity,omitempty"`
	Audience *Audience `json:"Audience,omitempty"`
	ScopeType *ScopeType `json:"ScopeType,omitempty"`
	Planned *bool `json:"Planned,omitempty"`
	Keywords *string `json:"Keywords,omitempty"`
}

type PtAdvice struct {
	AdviceRef *string `json:"AdviceRef,omitempty"`
	Details string `json:"Details,omitempty"`
}

type PtConsequence struct {
	Condition Condition `json:"Condition,omitempty"`
	Severity Severity `json:"Severity,omitempty"`
	Advice *PtAdvice `json:"Advice,omitempty"`
	Blocking *Blocking `json:"Blocking,omitempty"`
	Boarding Boarding `json:"Boarding,omitempty"`
	Delays *uint32 `json:"Delays,omitempty"`
}

type Blocking struct {
	JourneyPlanner bool `json:"JourneyPlanner,omitempty"`
}

type Boarding struct {
	ArrivalBoardingActivity *BoardingActivity `json:"ArrivalBoardingActivity,omitempty"`
	DepartureBoardingActivity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
}

type PublishingActions struct {
	PublishToWebAction *PublishToWebAction `json:"PublishToWebAction,omitempty"`
	PublishToMobileAction *PublishToMobileAction `json:"PublishToMobileAction,omitempty"`
	PublishToDisplayAction *PublishToDisplayAction `json:"PublishToDisplayAction,omitempty"`
	NotifyByEmailAction *NotifyByEmailAction `json:"NotifyByEmailAction,omitempty"`
	NotifyBySmsAction *NotifyBySmsAction `json:"NotifyBySmsAction,omitempty"`
}

type NotifyBySmsAction struct {
	ParameterizedAction *ParameterisedAction `json:"ParameterizedAction,omitempty"`
	BeforeNotices []BeforeNotice `json:"BeforeNotices,omitempty"`
	ClearNotice *bool `json:"ClearNotice,omitempty"`
	Phone *string `json:"Phone,omitempty"`
	Premium *bool `json:"Premium,omitempty"`
}

type NotifyByEmailAction struct {
	ParameterizedAction *ParameterisedAction `json:"ParameterizedAction,omitempty"`
	BeforeNotices []BeforeNotice `json:"BeforeNotices,omitempty"`
	ClearNotice *bool `json:"ClearNotice,omitempty"`
	Email *string `json:"Email,omitempty"`
}

type PublishToWebAction struct {
	ParameterizedAction *ParameterisedAction `json:"ParameterizedAction,omitempty"`
	Incidents *bool `json:"Incidents,omitempty"`
	HomePage *bool `json:"HomePage,omitempty"`
	Ticker *bool `json:"Ticker,omitempty"`
	SocialNetwork []string `json:"SocialNetwork,omitempty"`
}

type PublishToMobileAction struct {
	ParameterizedAction *ParameterisedAction `json:"ParameterizedAction,omitempty"`
	Incidents *bool `json:"Incidents,omitempty"`
	HomePage *bool `json:"HomePage,omitempty"`
}

type PublishToDisplayAction struct {
	ParameterizedAction *ParameterisedAction `json:"ParameterizedAction,omitempty"`
	OnPlace *bool `json:"OnPlace,omitempty"`
	Onboard *bool `json:"Onboard,omitempty"`
}

type ParameterisedAction struct {
	ActionStatus *ActionStatus `json:"ActionStatus,omitempty"`
	Description *string `json:"Description,omitempty"`
	ActionData []ActionData `json:"ActionData,omitempty"`
}

type Affect struct {
	AreaOfInterest *string `json:"AreaOfInterest,omitempty"`
	Operators []AffectedOperatorType `json:"Operators,omitempty"`
	Networks []Network `json:"Networks,omitempty"`
	StopPoints []AffectedStopPoint `json:"StopPoints,omitempty"`
	StopPlaces []AffectedPlace `json:"StopPlaces,omitempty"`
	Places []AffectedPlace `json:"Places,omitempty"`
	VehicleJourneys []AffectedVehicleJourney `json:"VehicleJourneys,omitempty"`
}

type ActionData struct {
	Name string `json:"Name,omitempty"`
	Prompt *string `json:"Prompt,omitempty"`
	PublishAtScope *PublishAtScope `json:"PublishAtScope,omitempty"`
}

type PublishAtScope struct {
	ScopeType *ScopeType `json:"ScopeType,omitempty"`
	Affects []Affect `json:"Affects,omitempty"`
}

type BeforeNotice struct {
	Interval string `json:"Interval,omitempty"`
}

type Network struct {
	NetworkRef *string `json:"NetworkRef,omitempty"`
	NetworkName *string `json:"NetworkName,omitempty"`
	RoutesAffected *string `json:"RoutesAffected,omitempty"`
}

type AffectedNetwork struct {
	Operators []AffectedOperator `json:"Operators,omitempty"`
	Network Network `json:"Network,omitempty"`
	Mode AffectedMode `json:"Mode,omitempty"`
	Lines []LineAffected `json:"Lines,omitempty"`
}

type LineAffected int

const (
	AllLines LineAffected = 0
	SelectedRoutes LineAffected = 1
	type AffectedLine struct {
	AffectedLine
}
	AffectedLine LineAffected = 2
)


type Zone struct {
	PlaceRef string `json:"PlaceRef,omitempty"`
	PlaceName string `json:"PlaceName,omitempty"`
	StopCondition string `json:"StopCondition,omitempty"`
}

type AffectedStopPoint struct {
	StopPointRef *string `json:"StopPointRef,omitempty"`
	AffectedModes *string `json:"AffectedModes,omitempty"`
	Zone *Zone `json:"Zone,omitempty"`
}

type Direction struct {
	DirectionRef string `json:"DirectionRef,omitempty"`
	DirectionName string `json:"DirectionName,omitempty"`
}

type AffectedLine struct {
	LineRef string `json:"LineRef,omitempty"`
	Destinations []AffectedStopPoint `json:"Destinations,omitempty"`
}

type AffectedOperator struct {
	OperatorRef string `json:"OperatorRef,omitempty"`
}

type AffectedOperatorType int

const (
	AllOperators AffectedOperatorType = 0
	type AffectedOperator struct {
	AffectedOperator
}
	AffectedOperator AffectedOperatorType = 1
)


type AffectedVehicleJourney struct {
	VehicleJourneyRef string `json:"VehicleJourneyRef,omitempty"`
}

type AffectedPlace struct {
	PlaceRef string `json:"PlaceRef,omitempty"`
	PlaceName string `json:"PlaceName,omitempty"`
}

type AffectedMode struct {
	AccessMode AccessMode `json:"AccessMode,omitempty"`
	Severity Severity `json:"Severity,omitempty"`
}

type DatedTimetableVersionFrame struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	LineRef string `json:"LineRef,omitempty"`
	DirectionRef string `json:"DirectionRef,omitempty"`
	DatedVehicleJourney DatedVehicleJourney `json:"DatedVehicleJourney,omitempty"`
}

type DatedVehicleJourney struct {
	VehicleJourneyName *string `json:"VehicleJourneyName,omitempty"`
	DestinationDisplay *string `json:"DestinationDisplay,omitempty"`
	HeadwayService *bool `json:"HeadwayService,omitempty"`
	Monitored *bool `json:"Monitored,omitempty"`
	DatedCalls []DatedCall `json:"DatedCalls,omitempty"`
	ServiceInfoGroup *ServiceInfoGroup `json:"ServiceInfoGroup,omitempty"`
	JourneyPatternInfo *JourneyPatternInfoGroup `json:"JourneyPatternInfo,omitempty"`
}

type DatedCall struct {
	StopPointRef *string `json:"StopPointRef,omitempty"`
	Order *uint32 `json:"Order,omitempty"`
	StopPointName *string `json:"StopPointName,omitempty"`
	DestinationDisplay *string `json:"DestinationDisplay,omitempty"`
	AimedArrivalTime *string `json:"AimedArrivalTime,omitempty"`
	ArrivalPlatform *string `json:"ArrivalPlatform,omitempty"`
	AimedQuayName *string `json:"AimedQuayName,omitempty"`
	AimedDepartureTime *string `json:"AimedDepartureTime,omitempty"`
	DeparturePlatformName *string `json:"DeparturePlatformName,omitempty"`
	DepartureBoardingActivity *BoardingActivity `json:"DepartureBoardingActivity,omitempty"`
	AimedHeadwayInterval *uint32 `json:"AimedHeadwayInterval,omitempty"`
	TargetedInterchange *TargetedInterchange `json:"TargetedInterchange,omitempty"`
}

type TargetedInterchange struct {
	InterchangeCode string `json:"InterchangeCode,omitempty"`
	DistributorVehicleJourneyRef string `json:"DistributorVehicleJourneyRef,omitempty"`
	DistributorConnectionLink DistributorConnectionLink `json:"DistributorConnectionLink,omitempty"`
	StaySeated *bool `json:"StaySeated,omitempty"`
	Guaranteed *bool `json:"Guaranteed,omitempty"`
	MaximumWaitTime *string `json:"MaximumWaitTime,omitempty"`
}

type DistributorConnectionLink struct {
	ConnectionCode *string `json:"ConnectionCode,omitempty"`
	StopPointRef *string `json:"StopPointRef,omitempty"`
	InterchangeDuration *uint32 `json:"InterchangeDuration,omitempty"`
	FrequentTravellerDuration *uint32 `json:"FrequentTravellerDuration,omitempty"`
	OccasionalTravellerDuration *uint32 `json:"OccasionalTravellerDuration,omitempty"`
	ImpairedAccessDuration *uint32 `json:"ImpairedAccessDuration,omitempty"`
}

type ServiceInfoGroup struct {
	OperatorRef *string `json:"OperatorRef,omitempty"`
	ProductCategoryRef *string `json:"ProductCategoryRef,omitempty"`
	ServiceFeatureRef []string `json:"ServiceFeatureRef,omitempty"`
	VehicleFeatureRef []VehicleFeature `json:"VehicleFeatureRef,omitempty"`
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
	EndTimeStatus *EndTimeStatus `json:"EndTimeStatus,omitempty"`
	EndTimePrecision *EndTimePrecision `json:"EndTimePrecision,omitempty"`
}

type Line struct {
	LineDirection string `json:"LineDirection,omitempty"`
	LineRef string `json:"LineRef,omitempty"`
}

type ValidityCondition struct {
	Period []ValidityPeriod `json:"Period,omitempty"`
}

type InfoMessage struct {
	FormatRef *string `json:"FormatRef,omitempty"`
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	ItemIdentifier *string `json:"ItemIdentifier,omitempty"`
	InfoMessageIdentifier string `json:"InfoMessageIdentifier,omitempty"`
	InfoMessageVersion *uint32 `json:"InfoMessageVersion,omitempty"`
	InfoChannelRef string `json:"InfoChannelRef,omitempty"`
	ValidUntilTime string `json:"ValidUntilTime,omitempty"`
	SituationRef *string `json:"SituationRef,omitempty"`
	Content *string `json:"Content,omitempty"`
}

type InfoMessageCancellation struct {
	RecordedAtTime string `json:"RecordedAtTime,omitempty"`
	InfoMessageIdentifier string `json:"InfoMessageIdentifier,omitempty"`
	InfoChannelRef *string `json:"InfoChannelRef,omitempty"`
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

