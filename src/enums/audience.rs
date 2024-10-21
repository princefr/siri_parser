use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "camelCase")]
pub enum Audience {
    Public,
    EmergencyService,
    Authorities,
    TransportOperators,
}