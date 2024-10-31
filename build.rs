use std::fs;
use std::fs::OpenOptions;
use std::io;
use std::io::Write;
use std::path::Path;

// specify the pa
const GO_FILE_PATH: &str = "./go/generation.go";

// Version that checks permissions first
pub fn safe_remove_if_exists<P: AsRef<Path>>(path: P) -> io::Result<()> {
    let path = path.as_ref();
    if path.exists() {
        // Check if we have permission to remove the file
        let metadata = fs::metadata(path)?;
        if metadata.permissions().readonly() {
            return Err(io::Error::new(
                io::ErrorKind::PermissionDenied,
                "File is read-only", 
            ));
        }
        fs::remove_file(path)?;
    }
    Ok(())
}

fn main() {
    //Remove file if exist
    let _ = safe_remove_if_exists(GO_FILE_PATH);
    // This is where you'd generate your enum based on your Rust code.
    // For demonstration, we'll just append a simple enum representation.
    let enum_code = "package siri_parser\n";

    // Append the generated code to the existing Go file
    let path = Path::new(GO_FILE_PATH);
    let mut file = OpenOptions::new()
        .write(true)
        .append(true)
        .create(true) // Create the file if it does not exist
        .open(&path)
        .expect("Unable to open file");

    // Write the generated code to the file
    if let Err(e) = writeln!(file, "{}", enum_code) {
        eprintln!("Failed to write to file: {}", e);
    }
}
