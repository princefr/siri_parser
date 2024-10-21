use serde::{Deserialize, Serialize};
use crate::{ models::service_delivery_info::ServiceDeliveryInfo, notifications::situation_exchange_notification::SituationExchangeNotification};


#[derive(Debug, Deserialize, Serialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct NotifySituationExchange {
    pub service_delivery_info: ServiceDeliveryInfo,
    pub notification: SituationExchangeNotification
}