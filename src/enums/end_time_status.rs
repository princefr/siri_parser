use serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "camelCase")]
pub enum EndTimeStatus {
    Undefined,
    LongTerm,
    ShortTerm,
}