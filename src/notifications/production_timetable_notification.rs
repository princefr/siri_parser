use serde::{Deserialize, Serialize};

use crate::deliveries::production_timetable_delivery::ProductionTimetableDelivery;


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct ProductionTimetableNotification {
    pub production_timetable_delivery: ProductionTimetableDelivery
}