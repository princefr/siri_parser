use std::fs;

use quote::ToTokens;
use serde::{Deserialize, Serialize};
use serde_xml_rs::from_str;

use super::body::Body;

#[derive(Debug, Deserialize, Serialize, PartialEq)]
pub struct Envelope {
    #[serde(rename = "Body", default)]
    pub body: Body,
}


impl Envelope {
    pub fn new(body: Body) -> Self {
        Self { body }
    }

    pub fn from_file(file_path: &str) -> Result<Self, Box<dyn std::error::Error>> {
        let content = fs::read_to_string(file_path)?;
        let envelope = from_str(&content)?;
        Ok(envelope)
    }


    pub fn from_str(xml_str: &str) -> Result<Self, Box<dyn std::error::Error>> {
        let envelope = from_str(xml_str)?;
        Ok(envelope)
    }
}


impl ToTokens for Envelope {
    fn to_tokens(&self, tokens: &mut proc_macro2::TokenStream) {
        let body = &self.body;
        tokens.extend(body.to_token_stream());
    }
}
