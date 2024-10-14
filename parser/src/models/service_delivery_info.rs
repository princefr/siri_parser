use quote::ToTokens;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct ServiceDeliveryInfo {
    #[serde(alias = "ResponseTimestamp")]
    pub response_timestamp: Option<String>,
    #[serde(alias = "ProducerRef")]
    pub producer_ref: Option<String>,
}


impl ToTokens for ServiceDeliveryInfo {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let response_timestamp = match &self.response_timestamp {
            Some(ts) => quote::quote! { Some(#ts.to_string()) },
            None => quote::quote! { None },
        };

        let producer_ref = match &self.producer_ref {
            Some(ref pr) => quote::quote! { Some(#pr.to_string()) },
            None => quote::quote! { None },
        };

        tokens.extend(quote::quote! {
            ServiceDeliveryInfo {
                response_timestamp: #response_timestamp,
                producer_ref: #producer_ref,
            }
        });
    }
}