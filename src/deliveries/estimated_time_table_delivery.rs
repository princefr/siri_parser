
use serde::{Deserialize, Serialize};

use crate::models::{estimated_journey_version_frame::EstimatedJourneyVersionFrame, xxx_delivery::XxxDelivery};


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
pub struct ResponseTimestamp(String);

#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct EstimatedTimetableDelivery {
    #[serde(flatten)]
    pub leader: XxxDelivery,
    pub estimated_journey_version_frame: EstimatedJourneyVersionFrame,
}





