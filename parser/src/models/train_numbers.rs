use quote::ToTokens;
use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub struct TrainNumbers {
    #[serde(rename = "TrainNumberRef")]
    pub train_number_ref: Option<String>,
}


impl ToTokens for TrainNumbers {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let train_number_ref = &self.train_number_ref;
        tokens.extend(train_number_ref.to_token_stream());
    }
}
