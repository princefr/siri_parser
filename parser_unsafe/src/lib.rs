use std::ffi::{CStr, CString};
use std::os::raw::c_char;
use parser::{SIRI, Envelope};
use std::ptr;


#[no_mangle]
pub extern "C" fn parse_str(s: *const c_char) -> *mut c_char {
    if s.is_null() { return std::ptr::null_mut(); }

    let c_str = unsafe { CStr::from_ptr(s) };

    let _string = match c_str.to_str() {
        Ok(string) => string,
        Err(_) => return ptr::null_mut(),
    };

    let siri: Result<Envelope, Box<dyn std::error::Error>> = SIRI::from_str::<Envelope>(_string);
    let siri = match siri {
        Ok(envelope) => envelope,
        Err(_) => return ptr::null_mut(),
    };

    let json = match serde_json::to_string(&siri) {
        Ok(json_string) => json_string,
        Err(_) => return ptr::null_mut(),
    };

    let c_json = CString::new(json).unwrap();
    c_json.into_raw()
}


#[no_mangle]
pub extern "C" fn free_str(s: *mut c_char) {
    if s.is_null() { return; }
    unsafe { let _ = CString::from_raw(s); }; // Automatically frees memory
}
