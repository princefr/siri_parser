use serde::{Deserialize, Serialize};

use crate::deliveries::stop_monitoring_delivery::StopMonitoringDelivery;




#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct StopMonitoringNotification {
    pub stop_monitoring_delivery: StopMonitoringDelivery
}