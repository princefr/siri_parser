use serde::{Deserialize, Serialize};

use crate::{models::framed_vehicle_journey_ref::FramedVehicleJourneyRef, siri_et::{distribution_group::DisruptionGroup, journey_pattern_info::JourneyPatternInfo, journey_progress_info::JourneyProgressInfo, train_number::TrainNumber, vehicle_journey_info::VehicleJourneyInfo}};

use super::{journey_part_info::JourneyPartInfo, monitored_call::MonitoredCall, onward_call::OnwardCall};



#[derive(Debug,  Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct MonitoredVehicleJourney {
    line_ref: String, // LineCode
    framed_vehicle_journey_ref: FramedVehicleJourneyRef, // Framed-Vehicle-JourneyRef-Structure
    // journey_pattern_info: Option<JourneyPatternInfo>, // Journey-Pattern-Info-Group
    // vehicle_journey_info: Option<VehicleJourneyInfo>, // Vehicle-JourneyInfo-Group
    // disruption_group: Option<DisruptionGroup>, // Disruption-Group
    // journey_progress_info: Option<JourneyProgressInfo>, // Journey-Progress-Info-Group
    // operational_info: Vec<TrainNumber>, // sequence
    // journey_parts: Vec<JourneyPartInfo>, // List of journey parts
    // calling_pattern: Option<MonitoredCall>, // Monitored-Call
    // onward_calls: Option<OnwardCall>, // Onward-Calls
}