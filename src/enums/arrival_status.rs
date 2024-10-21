use serde::{Deserialize, Serialize};


#[derive(Debug, Clone, PartialEq, Serialize, Deserialize, Eq)]
#[serde(rename_all = "camelCase")]
pub enum ArrivalStatus {
    OnTime,
    Missed,
    Arrived,
    NotExpected,
    Delayed,
    Early,
    Cancelled,
    NoReport,
}