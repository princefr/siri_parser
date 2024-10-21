

#[derive(Debug)]
struct FacilityStatus {
    status: FacilityAvailability, // Etat d’une Facility
    description: Option<String>, // Description associée à l’état d’une Facility
    special_needs: Vec<AccessibilityAssessment>, // État de l’accessibilité pour différents types de besoins spéciaux
}