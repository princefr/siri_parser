
use serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
pub struct VehicleMonitoringDelivery{
    vehicle_activitiy: Vec<VehicleActivity>,
    vehicle_activitiy_cancellation: Vec<VehicleActivityCancellation>,
    extensions: Option<Extensions>

}