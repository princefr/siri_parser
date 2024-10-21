
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct ServiceDeliveryInfo {
    pub response_timestamp: Option<String>,
    pub producer_ref: String,
    pub request_message_identifier: Option<String>,
    pub response_message_ref: Option<String>
}
