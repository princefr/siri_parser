use quote::{quote, ToTokens};
use serde::{Deserialize, Serialize};

use super::estimated_call::EstimatedCall;

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct EstimatedCalls {
    #[serde(alias = "EstimatedCall")]
    pub calls: Option<Vec<EstimatedCall>>,
}

impl ToTokens for EstimatedCalls {
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
                EstimatedCalls {
                    calls: Some(vec![ #(#call_tokens),* ]),
                }
            });
        } else {
            tokens.extend(quote! {
                EstimatedCalls {
                    calls: None,
                }
            });
        }
    }
}
