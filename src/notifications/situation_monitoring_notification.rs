use serde::{Deserialize, Serialize};

use crate::deliveries::situation_monitoring_delivery::SituationMonitoringDelivery;



#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct SituationMonitoringNotification {
    pub situation_monitoring_delivery: SituationMonitoringDelivery
}