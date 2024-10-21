use serde::{Deserialize, Serialize};

use crate::models::xxx_delivery::XxxDelivery;


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct GeneralMessageDelivery{
    #[serde(flatten)]
    pub leader: XxxDelivery,
}