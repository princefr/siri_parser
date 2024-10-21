use serde::{Deserialize, Serialize};



#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
struct FacilityCondition {
    facility: Facility, // Description générale d'une facility
    facility_ref: String, // Identifiant d'une Facility
    status: FacilityStatus, // Description de l'état d'une Facility
    counting: Option<MonitoredCounting>, // Mise à jour du compteur associé à la facility
    position: Option<FacilityUpdatedPosition>, // Mise à jour de la position de la facility
    situation_ref: Option<SituationCode>, // Identifiant d'une situation associée
    timing_information: Option<ValidityPeriod>, // Période de validité de la condition
    extensions: Option<Extensions>, // Emplacement pour extension utilisateur
}