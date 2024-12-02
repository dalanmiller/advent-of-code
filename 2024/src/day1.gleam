import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/option
import gleam/string
import utils

pub fn main() {
  part1()
  part2()
}

pub fn part1() {
  // use file <- result.try(utils.read_file("day1.input"))

  let file = case utils.read_file("day1.input") {
    Ok(s) -> s
    _ -> "dunno"
  }

  io.debug("OK")

  let list_a =
    file
    |> string.split("\n")
    |> list.filter_map(fn(a) {
      a
      |> string.split("   ")
      |> list.first
    })
    |> list.filter_map(fn(n) {
      case int.parse(n) {
        Ok(n) -> Ok(n)
        Error(_) -> Error(Nil)
      }
    })
    |> list.sort(int.compare)

  let list_b =
    file
    |> string.split("\n")
    |> list.filter_map(fn(a) {
      a
      |> string.split("   ")
      |> list.last
    })
    |> list.filter_map(fn(n) {
      case int.parse(n) {
        Ok(n) -> Ok(n)
        Error(_) -> Error(Nil)
      }
    })
    |> list.sort(int.compare)

  let result =
    list_a
    |> list.zip(list_b)
    |> list.fold(0, fn(b, a) { b + int.absolute_value(a.0 - a.1) })

  io.debug("part1: " <> int.to_string(result))
}

pub fn part2() {
  let file = case utils.read_file("day1.input") {
    Ok(s) -> s
    _ -> "dunno"
  }

  let counts =
    file
    |> string.split("\n")
    |> list.filter_map(fn(a) {
      a
      |> string.split("   ")
      |> list.last
    })
    |> list.filter_map(fn(n) {
      case int.parse(n) {
        Ok(n) -> Ok(n)
        Error(_) -> Error(Nil)
      }
    })
    |> list.fold(dict.new(), fn(d, n) {
      case dict.get(d, n) {
        Ok(count) -> dict.insert(d, n, count + 1)
        Error(_) -> dict.insert(d, n, 1)
      }
    })

  let result =
    file
    |> string.split("\n")
    |> list.filter_map(fn(a) {
      a
      |> string.split("   ")
      |> list.first
    })
    |> list.filter_map(fn(n) {
      case int.parse(n) {
        Ok(n) -> Ok(n)
        Error(_) -> Error(Nil)
      }
    })
    |> list.fold(0, fn(acc, a) { 
      case dict.get(counts, a) {
        Ok(v) -> { acc + {a * v} } 
        Error(_) -> acc
      }
    })

  // 1189304 too low 
  io.debug("part2: " <> int.to_string(result))
}
