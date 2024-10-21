use serde::{Deserialize, Serialize};

#[derive(Debug, Clone, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
pub enum FirstOrLastJourneyEnum {
    FirstServiceOfDay,
    LastServiceOfDay,
    OtherService,
    Unspecified,
}