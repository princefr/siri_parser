use serde::{Deserialize, Serialize};

#[derive(Debug, Serialize, Deserialize, PartialEq)]
#[serde(rename_all = "PascalCase")]
struct Facility {
    identify: String, // FacilityCode - Identifiant de la Facility
    description: Option<String>, // Description de la facility
    class: Option<FacilityClass>, // Classe de la facility
    features: Vec<FacilityFeature>, // Fonctionnalités du service
    temporal: Option<ValidityCondition>, // Période de validité de la facility
    location: Option<FacilityLocation>, // Localisation du service
    accessibility: Option<AccessibilityAssessment>, // Informations d’accessibilité
    extensions: Option<Extensions>, // Emplacement pour extension utilisateur
}