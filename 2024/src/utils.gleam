import gleam/io
import gleave
import simplifile

pub fn read_file(path: String) -> Result(String, String) {
  case simplifile.read(path) {
    Ok(result) -> Ok(result)
    Error(_reason) -> {
      io.println_error("Failed to read file!")
      gleave.exit(1)
      Error("failed")
    }
  }
}
