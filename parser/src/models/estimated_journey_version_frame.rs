use quote::ToTokens;
use serde::{Deserialize, Serialize};

use super::estimated_vehicle_journey::EstimatedVehicleJourney;

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct EstimatedJourneyVersionFrame {
    #[serde(alias = "RecordedAtTime")]
    pub recorded_at_time: String,
    #[serde(alias = "EstimatedVehicleJourney")]
    pub estimated_vehicle_journeys: Vec<EstimatedVehicleJourney>,
}


impl ToTokens for EstimatedJourneyVersionFrame {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let recorded_at_time = &self.recorded_at_time;
        let estimated_vehicle_journeys = &self.estimated_vehicle_journeys;
        tokens.extend(recorded_at_time.to_token_stream());
        tokens.extend(estimated_vehicle_journeys.iter().map(|estimated_vehicle_journeys| estimated_vehicle_journeys.to_token_stream()));
    }
}


