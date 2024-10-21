use serde::{Deserialize, Serialize};

#[derive(Debug, Default, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct JourneyPartInfo {
    journey_part_ref: String, // JourneyPart-Code
    train_number_ref: Option<String>, // TrainNumber
}