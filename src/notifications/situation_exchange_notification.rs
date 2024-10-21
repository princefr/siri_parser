use serde::{Deserialize, Serialize};

use crate::deliveries::situation_exchange_delivery::SituationExchangeDelivery;


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct SituationExchangeNotification {
    pub situation_exchange_delivery: SituationExchangeDelivery
}