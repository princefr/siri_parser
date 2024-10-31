package enums

type ServiceType string

const (
	EstimatedTimetable ServiceType = "EstimatedTimetable"
	StopMonitoring ServiceType = "StopMonitoring"
	VehicleMonitoring ServiceType = "VehicleMonitoring"
	FacilityMonitoring ServiceType = "FacilityMonitoring"
	ConnectionMonitoring ServiceType = "ConnectionMonitoring"
)