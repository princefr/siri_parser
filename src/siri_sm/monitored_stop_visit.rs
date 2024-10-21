use serde::{Deserialize, Serialize};

use super::monitored_vehicle_journey::MonitoredVehicleJourney;

#[derive(Debug,  Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct MonitoredStopVisit {
    recorded_at_time: String, // xsd:dateTime
    item_identifier: String, // ItemIdentifier
    monitoring_ref: String, // Monitoring-Code
    monitored_vehicle_journey: MonitoredVehicleJourney, // Monitored-Vehicle-Journey-Structure
    // extensions: Option<Extensions>, // user-defined extensions
}