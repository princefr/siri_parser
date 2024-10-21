use serde::{Deserialize, Serialize};


#[derive(Debug, Clone, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct DatedVehicleJourneyIndirectRef {
    pub origin_ref: String,
    pub aimed_departure_time: String,
    pub destination_ref: String,
    pub aimed_arrival_time: String,
}