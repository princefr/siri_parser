use serde::{Deserialize, Serialize};

use super::journey_end_names_group::JourneyEndNamesGroup;


#[derive(Debug, Default, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
pub struct VehicleJourneyInfoGroup {
    //pub service_info: Option<ServiceInfoGroup>, // Service-Info-Group
    pub journey_end_names: Option<JourneyEndNamesGroup>, // JourneyEndNamesGroup
    pub journey_info: Option<String>, // NLString (Vehicle-Journey-Name)
    pub journey_note: Option<String>, // NLString (additional text)
    pub headway_service: Option<bool>, // xsd:boolean
    pub origin_aimed_departure_time: Option<String>, // xsd:dateTime
    pub destination_aimed_arrival_time: Option<String>, // xsd:dateTime
    //first_or_last_journey: Option<FirstOrLastJourneyEnumeration> // FirstOrLast-Journey-Enumeration
}