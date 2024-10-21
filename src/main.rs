use siri_et_parser::{Envelope, SiriServiceType, SIRI};

pub fn main(){
    let siri = SIRI::from_file::<Envelope>("src/fixtures/siri_sm/siri-destineo-sm-cus-2-2024-04-04-13-02-25.xml");
    match siri {
        Ok(siri) => {
            match siri.body.0 {
                SiriServiceType::EstimatedTimetable(siri) => println!("XML: {:?}", siri),
                SiriServiceType::StopMonitoring(siri) => println!("XML: {:?}", siri),
                _ => println!("XML: {:?}", siri),
            }
        },
        Err(e) => println!("Error: {:?}", e),
    }
}