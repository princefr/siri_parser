use parser::models::body::Body as BodyParser;
use parser::services::connection_monitoring::NotifyConnectionMonitoring;
use parser::services::estimated_table::NotifyEstimatedTimetable;
use parser::services::facility_monitoring::NotifyFacilityMonitoring;
use parser::services::general_message::NotifyGeneralMessage;
use parser::services::production_timetable::NotifyProductionTimetable;
use parser::services::situation_exchange::NotifySituationExchange;
use parser::services::stop_monitoring::NotifyStopMonitoring;
use parser::services::vehicle_monitoring::NotifyVechicleMonitoring;
use parser::structures::accessibility_assesment::AccessibilityAssessment;
use parser::structures::action_data::ActionData;
use parser::structures::affect::Affect;
use parser::structures::affected_line::AffectedLine;
use parser::structures::affected_mode::AffectedMode;
use parser::structures::affected_network::AffectedNetwork;
use parser::structures::affected_operator::AffectedOperator;
use parser::structures::affected_place::AffectedPlace;
use parser::structures::affected_stop_point::AffectedStopPoint;
use parser::structures::affected_vehicle_journey::AffectedVehicleJourney;
use parser::structures::arrival::Arrival;
use parser::structures::arrival_info::ArrivalInfo;
use parser::structures::before_notice::BeforeNotice;
use parser::structures::blocking::Blocking;
use parser::structures::boarding::Boarding;
use parser::structures::calls::Calls;
use parser::structures::connecting_journey::ConnectingJourney;
use parser::structures::dated_call::DatedCall;
use parser::structures::dated_timetable_version_frame::DatedTimetableVersionFrame;
use parser::structures::dated_vehicle_journey::DatedVehicleJourney;
use parser::structures::dated_vehicle_journey_indirect_ref::DatedVehicleJourneyIndirectRef;
use parser::structures::departure::Departure;
use parser::structures::departure_info::DepartureInfo;
use parser::structures::direction::Direction;
use parser::structures::distribuor_info::DistributorInfo;
use parser::structures::facility::Facility;
use parser::structures::facility_status::FacilityStatus;
use parser::SiriServiceType;
use parser::{Envelope as EnvelopeParser, SIRI as SIRIParser};
use pyo3::prelude::*;

#[pyclass]
pub struct SIRI {}

#[pyclass]
pub struct Body(BodyParser);

#[pymethods]
impl Body {
    #[new]
    fn new(service_type: BodyParser) -> Self {
        Body(service_type)
    }

    /// Method to get string representation of the Body
    pub fn __str__(&self) -> String {
        format!("{:?}", self.0) // Assuming Body implements Display
    }

    /// Method to get NotifyProductionTimetable from Body in Python
    ///
    /// # Returns
    /// * `PyResult<Option<NotifyProductionTimetable>>`
    ///
    pub fn notify_production_timetable(&self) -> PyResult<Option<NotifyProductionTimetable>> {
        if let BodyParser(SiriServiceType::ProductionTimetable(ref production_timetable)) =
            self.0.clone()
        {
            Ok(Some(production_timetable.clone()))
        } else {
            Ok(None)
        }
    }

    /// Method to get NotifyEstimatedTimetable from Body in Python
    ///
    /// # Returns
    /// * `PyResult<Option<NotifyEstimatedTimetable>>`
    pub fn notify_estimated_timetable(&self) -> PyResult<Option<NotifyEstimatedTimetable>> {
        if let BodyParser(SiriServiceType::EstimatedTimetable(ref estimated_timetable)) =
            self.0.clone()
        {
            Ok(Some(estimated_timetable.clone()))
        } else {
            Ok(None)
        }
    }

    /// Method to get NotifyStopMonitoring from Body in Python
    ///
    /// # Returns
    /// * `PyResult<Option<NotifyStopMonitoring>>`
    pub fn notify_stop_monitoring(&self) -> PyResult<Option<NotifyStopMonitoring>> {
        if let BodyParser(SiriServiceType::StopMonitoring(ref stop_monitoring)) = self.0.clone() {
            Ok(Some(stop_monitoring.clone()))
        } else {
            Ok(None)
        }
    }

    /// Method to get NotifySituationExchange from Body in Python
    ///
    /// # Returns
    /// * `PyResult<Option<NotifySituationExchange>>`
    pub fn notify_vehicle_monitoring(&self) -> PyResult<Option<NotifyVechicleMonitoring>> {
        if let BodyParser(SiriServiceType::VehicleMonitoring(ref vehicle_monitoring)) =
            self.0.clone()
        {
            Ok(Some(vehicle_monitoring.clone()))
        } else {
            Ok(None)
        }
    }

    /// Method to get NotifyConnectionMonitoring from Body in Python
    ///
    /// # Returns
    /// * `PyResult<Option<NotifyConnectionMonitoring>>`
    pub fn notify_connection_monitoring(&self) -> PyResult<Option<NotifyConnectionMonitoring>> {
        if let BodyParser(SiriServiceType::ConnectionMonitoring(ref connection_monitoring)) =
            self.0.clone()
        {
            Ok(Some(connection_monitoring.clone()))
        } else {
            Ok(None)
        }
    }

    /// Method to get NotifyGeneralMessage from Body in Python
    ///
    /// # Returns
    /// * `PyResult<Option<NotifyGeneralMessage>>`
    pub fn notify_general_message(&self) -> PyResult<Option<NotifyGeneralMessage>> {
        if let BodyParser(SiriServiceType::GeneralMessage(ref general_message)) = self.0.clone() {
            Ok(Some(general_message.clone()))
        } else {
            Ok(None)
        }
    }

    /// Method to get NotifyFacilityMonitoring from Body in Python
    ///
    /// # Returns
    /// * `PyResult<Option<NotifyFacilityMonitoring>>`
    pub fn notify_facility_monitoring(&self) -> PyResult<Option<NotifyFacilityMonitoring>> {
        if let BodyParser(SiriServiceType::FacilityMonitoring(ref facility_monitoring)) =
            self.0.clone()
        {
            Ok(Some(facility_monitoring.clone()))
        } else {
            Ok(None)
        }
    }
}

#[pyclass]
pub struct Envelope(EnvelopeParser); // Wrapper for the Envelope type

#[pymethods]
impl Envelope {
    // Method to get string representation
    pub fn __str__(&self) -> String {
        format!("{:?}", self.0) // Assuming Envelope implements Display
    }

    pub fn body(&self) -> PyResult<Body> {
        return Ok(Body(self.0.body.clone()));
    }
}

#[pymethods]
impl SIRI {
    #[new]
    pub fn new() -> Self {
        SIRI {}
    }

    // Method to parse SIRI XML string to Envelope
    pub fn parse(&self, s: &str) -> PyResult<Envelope> {
        Python::with_gil(|_py| {
            match SIRIParser::from_str::<EnvelopeParser>(s) {
                Ok(envelope) => Ok(Envelope(envelope)), // Wrap Envelope
                Err(e) => Err(PyErr::new::<pyo3::exceptions::PyValueError, _>(
                    e.to_string(),
                )),
            }
        })
    }
}

#[pymodule]
fn siri_parser(_py: Python, m: &PyModule) -> PyResult<()> {
    m.add_class::<SIRI>()?;
    m.add_class::<Envelope>()?; // Add PyEnvelope class
    m.add_class::<Body>()?;
    m.add_class::<NotifyProductionTimetable>()?;
    m.add_class::<NotifyStopMonitoring>()?;
    m.add_class::<NotifySituationExchange>()?;
    m.add_class::<NotifyGeneralMessage>()?;
    m.add_class::<NotifyFacilityMonitoring>()?;
    m.add_class::<NotifyVechicleMonitoring>()?;
    m.add_class::<NotifyConnectionMonitoring>()?;
    m.add_class::<NotifyEstimatedTimetable>()?;
    Ok(())
}
