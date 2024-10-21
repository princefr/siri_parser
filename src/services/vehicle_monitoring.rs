use serde::{Deserialize, Serialize};
use crate::{models::service_delivery_info::ServiceDeliveryInfo, notifications::vehicle_monitoring_notification::VehicleMonitoringNotification};



#[derive(Debug, Deserialize, Serialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct NotifyVechicleMonitoring {
    pub service_delivery_info: ServiceDeliveryInfo,
    pub notification: VehicleMonitoringNotification
}