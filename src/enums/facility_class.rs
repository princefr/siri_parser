use serde::{Serialize, Deserialize};


#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "camelCase")]
pub enum FacilityClass {
    FixedEquipment,
    MobileEquipment,
    SiteComponent,
    Site,
    ParkingBay,
    Vehicle,
}