use parser::{Envelope, SiriServiceType, SIRI};

pub fn main(){
    let siri = SIRI::from_file::<Envelope>("src/fixtures/fm/facility_monitoring_example.xml");
    match siri {
        Ok(siri) => {
            match siri.body.0 {
                SiriServiceType::EstimatedTimetable(siri) => println!("XML: {:?}", siri),
                SiriServiceType::StopMonitoring(siri) => println!("XML: {:?}", siri),
                SiriServiceType::VehicleMonitoring(siri) => println!("XML Vehicule monitoring: {:?}", siri),
                //SiriServiceType::SituationExchange(siri) => println!("XML Situation Exchange: {:?}", siri),
                SiriServiceType::FacilityMonitoring(siri) => println!("XML Facility Monitoring: {:?}", siri),
                SiriServiceType::ConnectionMonitoring(siri) => println!("XML Connection Monitoring: {:?}", siri),
                _ => println!("XML: {:?}", siri),
            }
        },
        Err(e) => println!("Error: {:?}", e),
    }
}