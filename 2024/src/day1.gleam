import gleam/list
import gleam/io
import gleam/result
import gleam/string
import simplifile
import utils

pub fn part1() {
  // use file <- result.try(utils.read_file("day1.input"))

  let file = case utils.read_file("day1.input") {
    Ok(s) -> s
    _ -> "dunno"
  }
  let lines = string.split(file, "\n")

}

pub fn part2() {
  let file = utils.read_file("day2.input")
  let lines = string.split(file, "\n")
}
