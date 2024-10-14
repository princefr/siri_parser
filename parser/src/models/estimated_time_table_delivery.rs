use quote::ToTokens;
use serde::{Deserialize, Serialize};

use super::estimated_journey_version_frame::EstimatedJourneyVersionFrame;

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct EstimatedTimetableDelivery {
    #[serde(alias = "ResponseTimestamp", alias = "response_timestamp")]
    pub response_timestamp: String,

    #[serde(alias = "SubscriberRef")]
    pub subscriber_ref: Option<String>,

    #[serde(alias = "EstimatedJourneyVersionFrame")]
    pub estimated_journey_version_frame: EstimatedJourneyVersionFrame,
}

impl ToTokens for EstimatedTimetableDelivery {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let response_timestamp = &self.response_timestamp;
        let subscriber_ref = &self.subscriber_ref;
        let estimated_journey_version_frame = &self.estimated_journey_version_frame;
        tokens.extend(response_timestamp.to_token_stream());
        tokens.extend(subscriber_ref.to_token_stream());
        tokens.extend(estimated_journey_version_frame.to_token_stream());
    }
}
