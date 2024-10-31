package models



type Body struct {
	EstimatedTimetable NotifyEstimatedTimetable
	StopMonitoring NotifyStopMonitoring
	VehicleMonitoring NotifyVehicleMonitoring
	FacilityMonitoring NotifyFacilityMonitoring
	ConnectionMonitoring NotifyConnectionMonitoring
}