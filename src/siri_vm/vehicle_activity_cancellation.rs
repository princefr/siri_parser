use crate::siri_sm::journey_pattern_info_group::JourneyPatternInfoGroup;
use serde::{Deserialize, Serialize};

#[derive(Debug, Default, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
struct VehicleActivityCancellation {
    recorded_at_time: DateTime<Utc>, // Heure à laquelle l’annulation a été signalée
    event_identity: Option<ItemIdentifier>, // Identifiant de l’objet annulé
    vehicle_monitoring_ref: Option<VehicleMonitoringCode>, // Identifiant du véhicule
    framed_vehicle_journey_ref: Option<FramedVehicleJourney>, // Description de la course annulée
    line_ref: Option<LineCode>, // Identifiant de la ligne
    journey_pattern_info: Option<JourneyPatternInfoGroup>, // Informations sur le modèle de trajet
    reasons: Vec<String>, // Description textuelle de la cause de l’annulation
    extensions: Option<Extensions>, // Emplacement pour extension utilisateur
}