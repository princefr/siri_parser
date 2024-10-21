use serde::{Deserialize, Serialize};
use crate::{models::service_delivery_info::ServiceDeliveryInfo, notifications::facility_monitoring_notification::FacilityMonitoringNotification};


#[derive(Debug, Deserialize, Serialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct NotifyFacilityMonitoring {
    pub service_delivery_info: ServiceDeliveryInfo,
    pub notification: FacilityMonitoringNotification
}