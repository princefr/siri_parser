use serde::{Deserialize, Serialize};

#[derive(Debug, Default, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct GroupReservation {
    name_of_group: String, // NLString
    number_of_seats: u32, // NumberOfPassengers
}
