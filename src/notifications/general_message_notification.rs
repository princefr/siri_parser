use serde::{Deserialize, Serialize};

use crate::deliveries::general_message_delivery::GeneralMessageDelivery;

#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct GeneralMessageNotification {
    pub general_message_delivery: GeneralMessageDelivery
}