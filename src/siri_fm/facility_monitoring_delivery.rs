use serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
pub struct FacilityMonitoringDelivery {
    facility_conditions: Option<FacilityConditions>,
    extensions: Option<Extensions>,
}