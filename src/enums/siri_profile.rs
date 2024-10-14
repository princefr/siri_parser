use derives::SiriXMLType;
use serde::{Deserialize, Serialize};
use parser::models::envelope::Envelope;


#[derive(SiriXMLType, Debug, Deserialize, Serialize)]
pub enum SiriProfile {
    #[siri(file = "src/fixtures/siri_et_xml_tn/trip_add.xml")]
    TripAdd(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/trip_removal.xml")]
    TripRemoval(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/trip_with_deleted_stop_not_in_base_VERN.xml")]
    TripWithDeletedStopNotInBaseVern(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/back_to_normal.xml")]
    BackToNormal(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/back_to_normal_trip_with_skipped_stop_in_base_VERN.xml")]
    BackToNormalTripWithSkippedStopInBaseVern(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/delayed_stop_times.xml")]
    DelayedStopTimes(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/add_first_stop_and_delayed_stop_times.xml")]
    AddFirstStopAndDelayedStopTimes(Envelope),

    #[siri(file = "src/fixtures/siri_et_xml_tn/siri-destineo-et-sbx.xml")]
    Destineo(Envelope),
}


impl SiriProfile {
    pub fn get_envelope(&self) -> &Envelope {
        match self {
            SiriProfile::TripAdd(envelope) => envelope,
            SiriProfile::TripRemoval(envelope) => envelope,
            SiriProfile::TripWithDeletedStopNotInBaseVern(envelope) => envelope,
            SiriProfile::BackToNormal(envelope) => envelope,
            SiriProfile::BackToNormalTripWithSkippedStopInBaseVern(envelope) => envelope,
            SiriProfile::DelayedStopTimes(envelope) => envelope,
            SiriProfile::AddFirstStopAndDelayedStopTimes(envelope) => envelope,
            SiriProfile::Destineo(envelope) => envelope,
        }
    }

    // pub fn get_envelope_profile(envelope: &Envelope) -> &SiriProfile {
    //     let all_envelopes: Vec<SiriProfile> = SiriProfile::get_profiles();

    // }
        
}





