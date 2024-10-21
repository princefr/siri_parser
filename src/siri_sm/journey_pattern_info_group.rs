use crate::enums::vehicle_mode::VehicleMode;
use serde::{Deserialize, Serialize};



#[derive(Debug, Default, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct JourneyPatternInfoGroup {
    journey_pattern_ref: Option<String>, // JourneyPatternCode for the mission
    journey_pattern_name: Option<String>, // Public name or number of the journey
    vehicle_mode: Option<VehicleMode>, // Mode of transport (defaults to bus)
    route_ref: Option<String>, // RouteCode for the followed route
    published_line_name: String, // Name of the line (mandatory)
    direction_name: Option<String>, // Name of the direction (optional)
}