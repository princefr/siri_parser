use go_generation_derive::GoGenerate;
use serde::{Deserialize, Serialize};

use crate::models::xxx_delivery::XxxDelivery;

#[derive(Debug, Serialize, Clone, Deserialize, PartialEq, Eq, GoGenerate)]
#[serde(rename_all = "PascalCase")]
pub struct SituationExchangeDelivery {
    #[serde(flatten)]
    pub leader: XxxDelivery,
    //pub situations: Vec<PtSituationElement>
}
