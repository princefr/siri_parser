
use serde::{Deserialize, Serialize};

use crate::SiriServiceType;


#[derive(Debug,  PartialEq, Deserialize, Serialize, Eq)]
pub struct Body(pub SiriServiceType);

