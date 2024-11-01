use go_generation_derive::GoGenerate;
use serde::{Deserialize, Serialize};

use crate::models::xxx_delivery::XxxDelivery;

#[derive(Debug, Serialize, Clone, Deserialize, PartialEq, Eq, GoGenerate)]
#[serde(rename_all = "PascalCase")]
pub struct GeneralMessageDelivery {
    #[serde(flatten)]
    pub leader: XxxDelivery,
}
