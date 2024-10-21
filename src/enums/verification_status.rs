use serde::{Deserialize, Serialize};


#[derive(Debug, Serialize, Deserialize, PartialEq)]
pub enum VerificationStatus {
    Unknown,
    Unverified,
    Verified,
}