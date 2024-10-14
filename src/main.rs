use parser::models::envelope::Envelope;

pub fn main(){
    let envelope_resut = Envelope::from_file("src/fixtures/siri_et_xml_tn/siri-destineo-et-sbx.xml");
    assert!(envelope_resut.is_ok());
    println!("Trip Add XMLEnvelope: {:?}", envelope_resut.unwrap());
}