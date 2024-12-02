import gleam/dict
import gleam/int
import gleam/io
import gleam/list
import gleam/string
import utils

pub fn main() {
  let input = case utils.read_file("day1.input") {
    Ok(s) -> s
    _ -> "dunno"
  }
  part1(input)
  part2(input)
}

pub fn part1(input) -> Int {
  let list_a =
    input
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
    input
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
  result
}

pub fn part2(input) -> Int {
  let counts =
    input
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
    input
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
        Ok(v) -> {
          acc + { a * v }
        }
        Error(_) -> acc
      }
    })

  // 1189304 too low 
  io.debug("part2: " <> int.to_string(result))
  result
}
