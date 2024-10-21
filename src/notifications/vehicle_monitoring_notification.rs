use serde::{Deserialize, Serialize};

use crate::deliveries::vehicle_monitoring_delivery::VehicleMonitoringDelivery;


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct VehicleMonitoringNotification {
    pub vehicle_monitoring_delivery: VehicleMonitoringDelivery
}