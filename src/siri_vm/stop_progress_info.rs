


#[derive(Debug)]
struct StopProgressInfo {
    progress_between_stops: Location, // Position entre arrêts
    link_distance: Option<f64>, // Distance totale entre les deux arrêts
    percentage: Option<f64>, // Pourcentage de distance couverte
}