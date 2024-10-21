use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, PartialEq, Serialize, Deserialize, Eq)]
#[serde(rename_all = "camelCase")]
pub enum BoardingActivity {
    Alighting,
    NoAlighting,
    PassThru,
    Boarding,
    NoBoarding,
}