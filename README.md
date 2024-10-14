# SIRI Profile France Parser (Rust) - Proof of concept

## Overview
The SIRI (Service Interface for Real-time Information) Profile France Parser is a Rust project aimed at creating a universal parser for the SIRI spécification, specifically tailored to the French public transport context. This parser is designed to be the "One parser to rule them all," capable of handling various aspects of the SIRI standard efficiently and comprehensively.


## Purpose
The main goal of this parser is to provide a robust, efficient, and flexible solution for processing SIRI data in the French public transport ecosystem. By creating a unified parser, it aims to simplify the integration of real-time public transport information systems across different operators and regions in France.



## Key Features
- Implements parsing capabilities for cases below:
    - Trip adding
    - Trip removal
    - Trip add first stop and delayed stop times
    - Back to normal
    - Delayed stop times
    - Destineo-et-SBX
    - Trip with deleted stop not in base (VERN)
    - Back to normal trip with skipped stop in base (VERN)

Examples :

You can get the envelope of a xml file or a xml string by respectively calling from_file or from_str functions of the Envelope struct

from_file example: 

```rust

use parser::models::envelope::Envelope;

pub fn main(){
    let envelope_resut = Envelope::from_file("src/fixtures/siri_et_xml_tn/trip_add.xml");
    assert!(envelope_resut.is_ok());
    println!("Trip Add XMLEnvelope: {:?}", envelope_resut.unwrap());
}

```

Then you can get any value you want by calling any of the model

``` rust
pub struct Body {
    #[serde(rename = "NotifyEstimatedTimetable")]
    pub notify_estimated_timetable: NotifyEstimatedTimetable,
}

pub struct NotifyEstimatedTimetable {
    #[serde(rename = "ServiceDeliveryInfo")]
    pub service_delivery_info: ServiceDeliveryInfo,

    #[serde(rename = "Notification")]
    pub notification: Notification,
}

```
 
All models available can be found under  ``parser/src/models``




### 2. French Profile Specialization
- This repo is tailored to handle France-specific extensions and modifications to the SIRI standard
- Supports French-specific data formats and conventions


### 3. High Performance
- Utilizes Rust's performance benefits for fast and efficient parsing
- Designed with concurrency in mind to handle large volumes of real-time data


## SIRI Profile Handling

We use Rust Enum and macros derives abstraction to allow the specification of siri profiles more easily.

To add a siri profil just add a new ``SiriProfile`` enum with a file corresponding to the profile.

```rust

use derives::SiriXMLType;
use serde::{Deserialize, Serialize};
use parser::models::envelope::Envelope;


#[derive(SiriXMLType, Debug, Deserialize, Serialize)]
pub enum SiriProfile {
    #[siri(file = "src/fixtures/siri_et_xml_tn/trip_add.xml")]
    TripAdd(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/trip_removal.xml")]
    TripRemoval(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/trip_with_deleted_stop_not_in_base_VERN.xml")]
    TripWithDeletedStopNotInBaseVern(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/back_to_normal.xml")]
    BackToNormal(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/back_to_normal_trip_with_skipped_stop_in_base_VERN.xml")]
    BackToNormalTripWithSkippedStopInBaseVern(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/delayed_stop_times.xml")]
    DelayedStopTimes(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/add_first_stop_and_delayed_stop_times.xml")]
    AddFirstStopAndDelayedStopTimes(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/siri-destineo-et-sbx.xml")]
    Destineo(Envelope),
}

```

We adopted this modular architecture to allow easy updates and extentions as the SIRI standard evoles.
