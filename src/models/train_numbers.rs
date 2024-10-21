
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
pub struct TrainNumbers {
    #[serde(rename = "TrainNumberRef")]
    pub train_number_ref: Option<String>,
}
