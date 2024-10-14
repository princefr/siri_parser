use quote::ToTokens;
use serde::{Deserialize, Serialize};

use super::estimated_time_table_delivery::EstimatedTimetableDelivery;

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct Notification {
    #[serde(rename = "EstimatedTimetableDelivery")]
    pub estimated_timetable_delivery: EstimatedTimetableDelivery,
}

impl ToTokens for Notification {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let notification = &self.estimated_timetable_delivery;
        tokens.extend(notification.to_token_stream());
    }
}
