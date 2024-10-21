use serde::{Deserialize, Serialize};
use super::dated_vehicle_journey_indirect_ref::DatedVehicleJourneyIndirectRef;



#[derive(Debug, Clone, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "camelCase")]
pub enum JourneyIdentifier {
    DatedVehicleJourneyRef(String),
    EstimatedVehicleJourneyCode(String),
    DatedVehicleJourneyIndirectRef(DatedVehicleJourneyIndirectRef),
}