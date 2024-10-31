use go_generation_derive::GoGenerate;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq, Eq, GoGenerate)]
#[serde(rename_all = "camelCase")]
pub enum EndTimePrecision {
    Second,
    Minute,
    Hour,
    Day,
}
