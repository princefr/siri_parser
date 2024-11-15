use std::fmt;

use go_generation_derive::GoGenerate;
use quick_xml::de;
use serde::{de::{MapAccess, SeqAccess, Visitor}, ser, Deserialize, Deserializer, Serialize};
use serde_json::Value;

use crate::structures::estimated_call::EstimatedCall;



#[derive(Debug, Serialize, Clone,  PartialEq, Eq, GoGenerate)]
#[serde(rename_all = "PascalCase")]
#[serde(transparent)]
pub struct EstimatedCalls(pub Vec<EstimatedCall>);



impl<'de> Deserialize<'de> for EstimatedCalls {
    fn deserialize<D>(deserializer: D) -> Result<EstimatedCalls, D::Error>
    where
        D: Deserializer<'de>,
    {
        struct EstimatedCallsVisitor;

        impl<'de> Visitor<'de> for EstimatedCallsVisitor {
            type Value = EstimatedCalls;

            fn expecting(&self, formatter: &mut fmt::Formatter) -> fmt::Result {
                formatter.write_str("struct EstimatedCalls")
            }

            fn visit_seq<A>(self, mut seq: A) -> Result<EstimatedCalls, A::Error>
            where
                A: SeqAccess<'de>,
            {
                let mut estimated_calls = Vec::new();

                while let Some(value) = seq.next_element()? {
                    println!("EstimatedCall_life: {:?}", value);
                    estimated_calls.push(value);
                }

                Ok(EstimatedCalls(estimated_calls))
            }

            fn visit_map<A>(self, mut map: A) -> Result<EstimatedCalls, A::Error>
            where
                A: MapAccess<'de>,
            {
                // Handle case where we are deserializing a single "EstimatedCall" object instead of an array.
                let mut estimated_calls = Vec::new();

                while let Some((key, value)) = map.next_entry()? {
                    println!("EstimatedCall_life2: {:?}", value);
                    match key {
                        "EstimatedCall" => {
                            println!("EstimatedCall_life: {:?}", value);
                            estimated_calls.push(value);
                        }
                        _ => {
                            return Err(serde::de::Error::unknown_field(key, &["EstimatedCall"]));
                        }
                    }
                }

                Ok(EstimatedCalls(estimated_calls))
            }
        }

        deserializer.deserialize_seq(EstimatedCallsVisitor)
    }
}
