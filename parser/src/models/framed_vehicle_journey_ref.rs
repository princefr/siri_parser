use quote::ToTokens;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct FramedVehicleJourneyRef {
    #[serde(alias = "DataFrameRef")]
    pub data_frame_ref: String,

    #[serde(alias = "DatedVehicleJourneyRef")]
    pub dated_vehicle_journey_ref: String,
}


impl ToTokens for FramedVehicleJourneyRef {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let data_frame_ref = &self.data_frame_ref;
        let dated_vehicle_journey_ref = &self.dated_vehicle_journey_ref;
        tokens.extend(data_frame_ref.to_token_stream());
        tokens.extend(dated_vehicle_journey_ref.to_token_stream());
    }
}
