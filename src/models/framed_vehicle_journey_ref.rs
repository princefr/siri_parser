use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct FramedVehicleJourneyRef {
    pub data_frame_ref: String,
    pub dated_vehicle_journey_ref: String,
}
