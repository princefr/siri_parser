use serde::{Deserialize, Serialize};

use crate::{models::service_delivery_info::ServiceDeliveryInfo, notifications::estimated_timetable_notification::EstimatedTimetableNotification};


#[derive(Deserialize, Serialize, PartialEq, Debug, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct NotifyEstimatedTimetable {
    pub service_delivery_info: ServiceDeliveryInfo,
    pub notification: EstimatedTimetableNotification,
}

