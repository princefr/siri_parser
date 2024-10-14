use quote::ToTokens;
use serde::{Deserialize, Serialize};

use super::{notification::Notification, service_delivery_info::ServiceDeliveryInfo};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct NotifyEstimatedTimetable {
    #[serde(rename = "ServiceDeliveryInfo")]
    pub service_delivery_info: ServiceDeliveryInfo,

    #[serde(rename = "Notification")]
    pub notification: Notification,
}

impl ToTokens for NotifyEstimatedTimetable {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let service_delivery_info = &self.service_delivery_info;
        let notification = &self.notification;
        tokens.extend(service_delivery_info.to_token_stream());
        tokens.extend(notification.to_token_stream());
    }
}
