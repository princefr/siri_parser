use serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "lowercase")]
pub enum VehicleMode {
    Air,
    Bus,
    Coach,
    Ferry,
    Metro,
    Rail,
    Tram,
    Underground,
    // Add more modes as needed
}