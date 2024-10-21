use serde::{Deserialize, Serialize};

use crate::deliveries::connection_monitoring_delivery::ConnectionMonitoringDelivery;


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct ConnectionMonitoringNotification {
    pub connection_monitoring_delivery: ConnectionMonitoringDelivery
}