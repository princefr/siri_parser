use quote::ToTokens;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct EstimatedCall {
    #[serde(alias = "StopPointRef")]
    pub stop_point_ref: Option<String>,

    #[serde(alias = "Order")]
    pub order: Option<u32>,

    #[serde(alias = "AimedArrivalTime")]
    pub aimed_arrival_time: Option<String>,

    #[serde(alias = "ExpectedArrivalTime")]
    pub expected_arrival_time: Option<String>,

    #[serde(alias = "AimedDepartureTime")]
    pub aimed_departure_time: Option<String>,

    #[serde(alias = "ExpectedDepartureTime")]
    pub expected_departure_time: Option<String>,
}

impl ToTokens for EstimatedCall {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let stop_point_ref = &self.stop_point_ref;
        let order = &self.order;
        let aimed_arrival_time = &self.aimed_arrival_time;
        let expected_arrival_time = &self.expected_arrival_time;
        let aimed_departure_time = &self.aimed_departure_time;
        let expected_departure_time = &self.expected_departure_time;
        tokens.extend(stop_point_ref.to_token_stream());
        tokens.extend(order.to_token_stream());
        tokens.extend(aimed_arrival_time.to_token_stream());
        tokens.extend(expected_arrival_time.to_token_stream());
        tokens.extend(aimed_departure_time.to_token_stream());
        tokens.extend(expected_departure_time.to_token_stream());
    }
}
