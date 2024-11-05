use parser::{Envelope, SiriServiceType, SIRI};

pub fn main(){
    let siri = SIRI::from_file::<Envelope>("src/fixtures/gm/gm_example.xml");
    match siri {
        Ok(siri) => {
            match siri.body.0 {
                SiriServiceType::EstimatedTimetable(siri) => println!("XML Estimated Timetable: {:?}", siri),
                SiriServiceType::StopMonitoring(siri) => println!("XML Stop Monitoring: {:?}", siri),
                SiriServiceType::VehicleMonitoring(siri) => println!("XML Vehicule monitoring: {:?}", siri),
                SiriServiceType::FacilityMonitoring(siri) => println!("XML Facility Monitoring: {:?}", siri),
                SiriServiceType::ConnectionMonitoring(siri) => println!("XML Connection Monitoring: {:?}", siri),
                SiriServiceType::GeneralMessage(siri) => println!("XML General Message: {:?}", siri),
                _ => println!("XML: {:?}", siri),
            }
        },
        Err(e) => println!("Error: {:?}", e),
    }
}