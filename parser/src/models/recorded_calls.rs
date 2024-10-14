use quote::{ToTokens, quote};
use serde::{Deserialize, Serialize};

use super::recorded_call::RecordedCall;

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct RecordedCalls {
    #[serde(rename = "RecordedCall")]
    pub calls: Option<Vec<RecordedCall>>,
}

impl ToTokens for RecordedCalls {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        if let Some(ref calls) = self.calls {
            // Iterate over each EstimatedCall and convert it to tokens
            let call_tokens: Vec<proc_macro2::TokenStream> = calls.iter().map(|call| {
                let mut ts = proc_macro2::TokenStream::new();
                call.to_tokens(&mut ts); // Assuming EstimatedCall also implements ToTokens
                ts
            }).collect();

            // Generate the final token stream
            tokens.extend(quote! {
                RecordedCalls {
                    calls: Some(vec![ #(#call_tokens),* ]),
                }
            });
        } else {
            tokens.extend(quote! {
                RecordedCalls {
                    calls: None,
                }
            });
        }
    }
}