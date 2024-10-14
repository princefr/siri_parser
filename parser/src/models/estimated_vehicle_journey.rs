use quote::ToTokens;
use serde::{Deserialize, Serialize};

use super::{
    estimated_calls::EstimatedCalls, framed_vehicle_journey_ref::FramedVehicleJourneyRef, recorded_calls::RecordedCalls, train_numbers::TrainNumbers
};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct EstimatedVehicleJourney {
    #[serde(alias = "LineRef")]
    pub line_ref: String,

    #[serde(alias = "PublishedLineName")]
    pub published_line_name: Option<String>,

    #[serde(alias = "DirectionRef")]
    pub direction_ref: Option<String>,

    #[serde(alias = "DatedVehicleJourneyRef")]
    pub dated_vehicule_journey_ref: Option<String>,

    #[serde(alias = "Cancellation")]
    pub cancellation: Option<String>,

    #[serde(alias = "ExtraJourney")]
    pub extra_journey: Option<String>,

    #[serde(alias = "JourneyPatternName")]
    pub journey_pattern_name: Option<String>,

    #[serde(alias = "VehicleMode")]
    pub vehicle_mode: Option<String>,

    #[serde(alias = "OriginRef")]
    pub origin_ref: Option<String>,

    #[serde(alias = "OriginName")]
    pub origin_name: Option<String>,

    #[serde(alias = "DestinationRef")]
    pub destination_ref: Option<String>,

    #[serde(alias = "DestinationName")]
    pub destination_name: Option<String>,

    #[serde(alias = "OperatorRef")]
    pub operator_ref: Option<String>,

    #[serde(alias = "ProductCategoryRef")]
    pub product_category_ref: Option<String>,

    #[serde(alias = "TrainNumbers")]
    pub train_numbers: Option<TrainNumbers>,

    #[serde(alias = "VehicleJourneyName")]
    pub vehicule_journey_name: Option<String>,

    #[serde(alias = "OriginAimedDepartureTime")]
    pub origin_aimed_departure_time: Option<String>,

    #[serde(alias = "DestinationAimedArrivalTime")]
    pub destination_aimed_arrival_time: Option<String>,

    #[serde(alias = "RecordedCalls")]
    pub recorded_calls: Option<RecordedCalls>,

    #[serde(alias = "EstimatedCalls")]
    pub estimated_calls: Option<EstimatedCalls>,

    #[serde(alias = "FramedVehicleJourneyRef")]
    pub framed_vehicle_journey_ref: Option<FramedVehicleJourneyRef>,

    #[serde(alias = "DataSource")]
    pub data_source: Option<String>,

    #[serde(alias = "VehicleRef")]
    pub vehicle_ref: Option<String>,
}

impl ToTokens for EstimatedVehicleJourney {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let line_ref = &self.line_ref;
        let published_line_name = &self.published_line_name;
        let direction_ref = &self.direction_ref;
        let dated_vehicule_journey_ref = &self.dated_vehicule_journey_ref;
        let cancellation = &self.cancellation;
        let extra_journey = &self.extra_journey;
        let journey_pattern_name = &self.journey_pattern_name;
        let vehicle_mode = &self.vehicle_mode;
        let origin_ref = &self.origin_ref;
        let origin_name = &self.origin_name;
        let destination_ref = &self.destination_ref;
        let destination_name = &self.destination_name;
        let operator_ref = &self.operator_ref;
        let product_category_ref = &self.product_category_ref;
        let train_numbers = &self.train_numbers;
        let vehicule_journey_name = &self.vehicule_journey_name;
        let origin_aimed_departure_time = &self.origin_aimed_departure_time;
        let destination_aimed_arrival_time = &self.destination_aimed_arrival_time;
        let recorded_calls = &self.recorded_calls;
        let estimated_calls = &self.estimated_calls;
        let framed_vehicle_journey_ref = &self.framed_vehicle_journey_ref;
        let data_source = &self.data_source;
        let vehicle_ref = &self.vehicle_ref;
        tokens.extend(line_ref.to_token_stream());
        tokens.extend(published_line_name.to_token_stream());
        tokens.extend(direction_ref.to_token_stream());
        tokens.extend(dated_vehicule_journey_ref.to_token_stream());
        tokens.extend(cancellation.to_token_stream());
        tokens.extend(extra_journey.to_token_stream());
        tokens.extend(journey_pattern_name.to_token_stream());
        tokens.extend(vehicle_mode.to_token_stream());
        tokens.extend(origin_ref.to_token_stream());
        tokens.extend(origin_name.to_token_stream());
        tokens.extend(destination_ref.to_token_stream());
        tokens.extend(destination_name.to_token_stream());
        tokens.extend(operator_ref.to_token_stream());
        tokens.extend(product_category_ref.to_token_stream());
        tokens.extend(train_numbers.to_token_stream());
        tokens.extend(vehicule_journey_name.to_token_stream());
        tokens.extend(origin_aimed_departure_time.to_token_stream());
        tokens.extend(destination_aimed_arrival_time.to_token_stream());
        tokens.extend(recorded_calls.to_token_stream());
        tokens.extend(estimated_calls.to_token_stream());
        tokens.extend(framed_vehicle_journey_ref.to_token_stream());
        tokens.extend(data_source.to_token_stream());
        tokens.extend(vehicle_ref.to_token_stream());
    }
}
