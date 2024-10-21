use serde::{Deserialize, Serialize};

// Structure for vehicle location
#[derive(Debug, Default, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
pub struct LocationStructure {
    // Define fields for the vehicle location as needed
}