use serde::{Serialize, Deserialize};



#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "camelCase")]
enum Trend {
    Decreasing, // La valeur est actuellement en baisse.
    Increasing, // La valeur est actuellement en hausse.
    Stable,     // La valeur est actuellement stable.
    Unstable,   // La valeur est actuellement instable sans tendance claire.
    Unknown,    // La tendance est inconnue.
}