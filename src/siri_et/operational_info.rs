use serde::{Deserialize, Serialize};
use super::{journey_part::JourneyPart, train_number::TrainNumber};


#[derive(Debug, Clone, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
pub struct OperationalInfo {
    pub train_numbers: Vec<TrainNumber>,
    pub journey_parts: Vec<JourneyPart>,
}