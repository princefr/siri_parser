use serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub enum MonitoringError {
    GPS,
    GPRS,
    Radio,
}