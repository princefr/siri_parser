use serde::{Deserialize, Serialize};
use crate::{ models::service_delivery_info::ServiceDeliveryInfo, notifications::production_timetable_notification::ProductionTimetableNotification};




#[derive(Debug, Deserialize, Serialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct NotifyProductionTimetable {
    pub service_delivery_info: ServiceDeliveryInfo,
    pub notification : ProductionTimetableNotification
}