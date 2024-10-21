


#[derive(Debug)]
struct VehicleActivity {
    recorded_at_time: DateTime<Utc>, // Heure de mise à jour de la position
    valid_until_time: DateTime<Utc>,  // Heure jusqu'à laquelle l'information est valable
    identity: Option<ItemIdentifier>,   // Identifiant pour annulation
    vehicle_monitoring_ref: Option<VehicleMonitoringIdentifier>, // Identifiant du véhicule
    stop_progress_info: Option<StopProgressInfo>, // Position du véhicule entre les arrêts
    journey_info: MonitoredVehicleJourney, // Détails de la course effectuée par le véhicule
    messages: Vec<String>, // Informations textuelles concernant le véhicule
    extensions: Option<Extensions>, // Emplacement pour extension utilisateur
}