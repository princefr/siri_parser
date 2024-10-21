use serde::{Deserialize, Serialize};
use super::fist_or_last_journey_enum::FirstOrLastJourneyEnum;


#[derive(Debug, Clone, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
pub struct EstimatedInfo {
    pub headway_service: Option<bool>,
    pub origin_aimed_departure_time: Option<String>,
    pub destination_aimed_arrival_time: Option<String>,
    pub first_or_last_journey: Option<FirstOrLastJourneyEnum>,
}