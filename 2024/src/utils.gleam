import filepath
import gleam/io
import gleave
import simplifile

pub fn read_file(path: String) -> Result(String, String) {
  io.debug("READING #{path}")
  io.debug(path)

  let codespaces_path = "/workspaces/advent-of-code/2024/src"
  let laptop_path = "/Users/dalan/repos/dalan_advent_of_code/2024/src"

  let pre_path = case simplifile.is_directory("/workspaces") {
    Ok(True) -> {
      codespaces_path
    }
    // This is a Github workspaces 
    Ok(False) | Error(_) -> {
      laptop_path
    }
    // This is maybe my home laptop 
  }

  let fullpath = filepath.join(pre_path, path)
  case simplifile.read(fullpath) {
    Ok(result) -> Ok(result)
    Error(_reason) -> {
      io.println_error("Failed to read file")
      io.println_error("path: " <> fullpath)
      gleave.exit(1)
      Error("failed")
    }
  }
}
