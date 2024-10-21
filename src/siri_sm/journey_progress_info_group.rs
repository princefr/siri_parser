
use serde::{Deserialize, Serialize};
use crate::enums::{monitoring_error::MonitoringError, occupancy::Occupancy};
use super::location_structure::LocationStructure;


#[derive(Debug, Default, Serialize, Deserialize, PartialEq)]
struct JourneyProgressInfoGroup {
    monitored: Option<bool>, // Indicates if the vehicle is still located
    monitoring_error: Option<MonitoringError>, // Cause of delocalization, if applicable
    in_congestion: Option<bool>, // Indicates if the vehicle is in congestion
    in_panic: Option<bool>, // Indicates if the vehicle alarm is activated
    vehicle_location: Option<LocationStructure>, // Position of the vehicle
    bearing: Option<f64>, // Orientation of the vehicle in degrees (0-360)
    occupancy: Option<Occupancy>, // Level of vehicle occupancy
    delay: Option<i32>, // Level of delay (negative value indicates early)
}