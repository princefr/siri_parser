use serde::{Deserialize, Serialize};

use crate::deliveries::facility_monitoring_delivery::FacilityMonitoringDelivery;


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct FacilityMonitoringNotification {
    pub facility_monitoring_delivery: FacilityMonitoringDelivery
}