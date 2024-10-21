use serde::{Deserialize, Serialize};
use crate::{models::service_delivery_info::ServiceDeliveryInfo, notifications::general_message_notification::GeneralMessageNotification};


#[derive(Debug, Deserialize, Serialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct NotifyGeneralMessage {
    pub service_delivery_info: ServiceDeliveryInfo,
    pub notification: GeneralMessageNotification
}