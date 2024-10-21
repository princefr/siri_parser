
use serde::{Deserialize, Serialize};


#[derive(Debug, Default, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct StopIdentity {
    stop_point_code: String, // StopPoint-Code
    // Add any additional fields as necessary
}