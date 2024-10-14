use quote::ToTokens;
use serde::{Deserialize, Serialize};

use super::notify_estimated_time_table::NotifyEstimatedTimetable;

#[derive(Debug, Deserialize, PartialEq, Serialize)]
pub struct Body {
    #[serde(rename = "NotifyEstimatedTimetable")]
    pub notify_estimated_timetable: NotifyEstimatedTimetable,
}

impl Default for Body {
    fn default() -> Self {
        todo!()
    }
}

impl ToTokens for Body {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let body = &self.notify_estimated_timetable;
        tokens.extend(body.to_token_stream());
    }
}
