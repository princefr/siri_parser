# SIRI Profile France Parser (Not Ready..yet)

![GitHub License](https://img.shields.io/github/license/princefr/siri_parser?style=for-the-badge)

[![PyPI](https://img.shields.io/pypi/v/maturin.svg?logo=python&style=for-the-badge)](https://pypi.org/project/siri-parser)
![Crates.io Version](https://img.shields.io/crates/v/siri-parser?style=for-the-badge)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/princefr/siri_parser/go?style=for-the-badge&label=Go%20V.1.0.0&link=https%3A%2F%2Fpkg.go.dev%2Fgithub.com%2Fprincefr%2Fsiri_parser%2Fgo)](https://pkg.go.dev/github.com/princefr/siri_parser/go)




![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/princefr/siri_parser/CI_golang.yml?style=for-the-badge&label=Golang%20build)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/princefr/siri_parser/python_workflow.yml?style=for-the-badge&label=Python%20build)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/princefr/siri_parser/CI_rust.yml?style=for-the-badge&label=Rust%20build)


A Rust library for parsing  SIRI (Service Interface for Real-time Information) messages according to the French national profile specification (SIRI Profile France).

## Overview
This library implements the French national adaptation of the SIRI standard, which is used for real-time public transport information exchange in France. It provides tools for parsing and working with SIRI messages that conform to the French profile specifications.


## Key Features

Full support for SIRI Profile France v2.0
XML parsing
Type-safe data structures
Support for all SIRI France service types:

- Stop Monitoring
- Estimated Timetable
- General Message
- Situation Exchange
- Production Timetable
- Facility Monitoring
- Connection Monitoring
- Vehicle Monitoring



## Siri interpretation notice et folders structures.
SIRI (Service Interface for Real-time Information) defines 
- Services (services available)
- Structures (data class/ structures for every object found in the specficication)
- Enums 

We decided to stricly follow the naming convention given by the specification and decided to arrange our parser folder this way

if you navigate to ``parser/src/``

- ``deliviries``, holds all different delivery types for the specification,  ex: ConnectionMonitoringDistributorDelivery
- ``enums``, holds all enums that can be found within the specification
- ``services``, holds all the services available data structures
- ``structures``, holds all the data structures the SIRI specification can offer


Sometimes you will find ``:::`` in the specification, ex: 

| LEADER | :::    | 1:1     | xxx-Delivery     | voir xxxDelivery  |
|--------|--------|---------|------------------|-------------------|
|        |        |         |                  |                   |


For us, it means  single or multiples members of xxx-Delivery/LEADER can be found flattened into the structures where xxx-Delivery/LEADER is implemented.

We have therefore decide to flatten XxxDelivery rust struct each time we found ``:::`` or ``LEADER`` defined into a response data structure and we build a custom deserializer for it.
We have done so everytime we found that a data structure needed to be flattened

for example: 

```rust
use go_generation_derive::GoGenerate;
use serde::{Deserialize, Serialize};

use crate::models::{
    estimated_journey_version_frame::EstimatedJourneyVersionFrame, xxx_delivery::XxxDelivery,
};



#[derive(Debug, Serialize, Clone, Deserialize, PartialEq, Eq, GoGenerate)]
#[serde(rename_all = "PascalCase")]
pub struct EstimatedTimetableDelivery {
    #[serde(flatten)]
    pub leader: XxxDelivery,
    pub estimated_journey_version_frame: EstimatedJourneyVersionFrame,
}

```


``EstimatedTimetableDelivery`` define a leader in his response, its then flatened.
When deserializing all the LEADER members are searched and deserialized within the custom deserializer.

Deserializer for the leader

```rust
impl<'de> Deserialize<'de> for XxxDelivery {
    fn deserialize<D>(deserializer: D) -> Result<Self, D::Error>
    where
        D: serde::Deserializer<'de>,
    {
        struct XxxDeliveryVisitor;

        impl<'de> serde::de::Visitor<'de> for XxxDeliveryVisitor {
            type Value = XxxDelivery;

            fn expecting(&self, formatter: &mut std::fmt::Formatter) -> std::fmt::Result {
                formatter.write_str("an XML element representing XxxDelivery")
            }

            fn visit_map<V>(self, mut map: V) -> Result<Self::Value, V::Error>
            where
                V: serde::de::MapAccess<'de>,
            {
                let mut response_timestamp = None;
                let mut request_message_ref = None;
                let mut subscriber_ref = None;
                let mut subscription_ref = None;
                let mut status = None;
                //let mut error_condition = None;
                let mut valid_until = None;
                let mut shortest_possible_cycle = None;

                while let Some(key) = map.next_key::<String>()? {
                    match key.as_str() {
                        key if key.starts_with("@xmlns:siri") => {
                            continue;
                        }
                        key if key.starts_with("@version") => {
                            continue;
                        }
                        "ResponseTimestamp" => {
                            let value: Value = map.next_value()?;
                            response_timestamp =
                                value.get("$text").and_then(Value::as_str).map(String::from);
                        }
                        "RequestMessageRef" => {
                            let value: Value = map.next_value()?;
                            request_message_ref =
                                value.get("$text").and_then(Value::as_str).map(String::from);
                        }
                        "SubscriberRef" => {
                            let value: Value = map.next_value()?;
                            subscriber_ref =
                                value.get("$text").and_then(Value::as_str).map(String::from);
                        }
                        "SubscriptionRef" => {
                            let value: Value = map.next_value()?;
                            subscription_ref =
                                value.get("$text").and_then(Value::as_str).map(String::from);
                        }
                        "Status" => {
                            let value: Value = map.next_value()?;
                            status = value.get("$text").and_then(Value::as_bool);
                        }
                        "ValidUntil" => {
                            let value: Value = map.next_value()?;
                            valid_until =
                                value.get("$text").and_then(Value::as_str).map(String::from);
                        }
                        "ShortestPossibleCycle" => {
                            let value: Value = map.next_value()?;
                            shortest_possible_cycle =
                                value.get("$text").and_then(Value::as_u64).map(|x| x as u32);
                        }

                        _ => {
                            return Err(serde::de::Error::unknown_field(
                                key.to_string().as_str(),
                                &[
                                    "ResponseTimestamp",
                                    "RequestMessageRef",
                                    "SubscriberRef",
                                    "Status",
                                    "ValidUntil",
                                    "ShortestPossibleCycle",
                                ],
                            ));
                        }
                    }
                }

                Ok(XxxDelivery {
                    response_timestamp: response_timestamp.unwrap(),
                    request_message_ref: request_message_ref,
                    subscriber_ref: subscriber_ref,
                    subscription_ref: subscription_ref,
                    status: status,
                    valid_until: valid_until,
                    shortest_possible_cycle: shortest_possible_cycle,
                })
            }
        }
        deserializer.deserialize_map(XxxDeliveryVisitor)
    }
}

```





## Examples

You can try the RUST parser this way:

```rust 

use parser::{Envelope, SiriServiceType, SIRI};

pub fn main(){
    let siri = SIRI::from_file::<Envelope>("src/fixtures/gm/gm_example.xml");
    match siri {
        Ok(siri) => {
            match siri.body.0 {
                SiriServiceType::EstimatedTimetable(siri) => println!("XML Estimated Timetable: {:?}", siri),
                SiriServiceType::StopMonitoring(siri) => println!("XML Stop Monitoring: {:?}", siri),
                SiriServiceType::VehicleMonitoring(siri) => println!("XML Vehicule monitoring: {:?}", siri),
                SiriServiceType::FacilityMonitoring(siri) => println!("XML Facility Monitoring: {:?}", siri),
                SiriServiceType::ConnectionMonitoring(siri) => println!("XML Connection Monitoring: {:?}", siri),
                SiriServiceType::GeneralMessage(siri) => println!("XML General Message: {:?}", siri),
                _ => println!("XML: {:?}", siri),
            }
        },
        Err(e) => println!("Error: {:?}", e),
    }
}

```


an example can be found under ``src/main.rs``


# Python 

We also provide a python package that can be installed using Pypi, it requires python to be at version >=3.7

``pip install siri-parser``

## Example
```python

import siri_parser as siri

siri_parser = siri.SIRI()


data = """SIRI XML"""

parsed_data = siri_parser.parse(data)

if parsed_data.body().notify_production_timetable():
    print("production_timetable")
elif parsed_data.body().notify_estimated_timetable():
    print("estimated_timetable")
else:
    print("other services")

```




## License
MIT license

