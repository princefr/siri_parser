use quote::ToTokens;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct RecordedCall {
    #[serde(alias = "StopPointRef")]
    pub stop_point_ref: Option<String>,

    #[serde(alias = "Order")]
    pub order: Option<u32>,

    #[serde(alias = "ExpectedDepartureTime")]
    pub expected_departure_time: Option<String>,

    #[serde(alias = "ExpectedArrivalTime")]
    pub expected_arrival_time: Option<String>,

    #[serde(alias = "Cancellation")]
    pub cancellation: Option<bool>,

    #[serde(alias = "ExtraCall")]
    pub extra_call: Option<bool>,
}

impl ToTokens for RecordedCall {   
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let stop_point_ref = &self.stop_point_ref;
        let order = &self.order;
        let expected_departure_time = &self.expected_departure_time;
        let expected_arrival_time = &self.expected_arrival_time;
        let cancellation = &self.cancellation;
        let extra_call = &self.extra_call;
        tokens.extend(stop_point_ref.to_token_stream());
        tokens.extend(order.to_token_stream());
        tokens.extend(expected_departure_time.to_token_stream());
        tokens.extend(expected_arrival_time.to_token_stream());
        tokens.extend(cancellation.to_token_stream());
        tokens.extend(extra_call.to_token_stream());
    }
}
