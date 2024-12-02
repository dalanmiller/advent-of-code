import filepath
import gleam/io
import gleave
import simplifile

pub fn read_file(path: String) -> Result(String, String) {
  io.debug("READING #{path}")
  io.debug(path)
  let fullpath =
    filepath.join("/Users/dalan/repos/dalan_advent_of_code/2024/src", path)
  case simplifile.read(fullpath) {
    Ok(result) -> Ok(result)
    Error(_reason) -> {
      io.println_error("Failed to read file!")
      gleave.exit(1)
      Error("failed")
    }
  }
}
