use  serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "camelCase")]
enum FacilityAvailability {
    Unknown,
    Available,
    NotAvailable,
    PartiallyAvailable,
    Added,
    Removed,
}