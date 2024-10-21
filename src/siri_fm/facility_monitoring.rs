use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
pub struct FacilityMonitoring {
    facility_monitoring_delivery: Option<FacilityMonitoringDelivery>,
    extensions: Option<Extensions>,
}