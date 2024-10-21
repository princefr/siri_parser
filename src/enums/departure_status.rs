
use serde::{Deserialize, Serialize};


#[derive(Debug, Clone, PartialEq, Serialize, Deserialize, Eq)]
#[serde(rename_all = "camelCase")]
pub enum DepartureStatus {
    OnTime,
    Early,
    Delayed,
    Cancelled,
    Arrived,
    Departed,
    NotExpected,
    NoReport,
}
