use serde::{Deserialize, Serialize};

use crate::deliveries::estimated_time_table_delivery::EstimatedTimetableDelivery;



#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct EstimatedTimetableNotification {
    #[serde(alias = "EstimatedTimetableDelivery", alias = "siri1:EstimatedTimetableDelivery")]
    pub estimated_timetable_delivery: EstimatedTimetableDelivery,
}