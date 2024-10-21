use serde::{Deserialize, Serialize};



#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "camelCase")]
enum Severity {
    Normal,
    // Add additional severity levels as needed
}