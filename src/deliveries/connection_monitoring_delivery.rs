use serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, PartialEq, Eq)]
#[serde(rename_all = "PascalCase")]
pub struct ConnectionMonitoringDelivery {
    // pub connection_monitoring_feeder_delivery: ConnectionMonitoringFeederDelivery,
    // pub connection_distributor_delivery: ConnectionDistributorDelivery
}