import gleam/int
import gleam/io
import gleam/list
import gleam/option.{None, Some}
import gleam/regexp
import gleam/string
import utils.{read_file}

pub fn main() {
  let input = case read_file("day3.input") {
    Ok(s) -> s
    _ -> "dunno"
  }

  io.debug(part1(input))
  io.debug(part2(input))
}

fn process(list: List(String), signal: Bool, acc: List(Int)) -> List(Int) {
  case list {
    [] -> acc
    [head, ..rest] -> {
      case int.parse(head) {
        Ok(v) -> {
          let new_acc = case signal {
            True -> {
              list.append(acc, [v])
            }
            False -> {
              acc
            }
          }
          process(rest, signal, new_acc)
        }
        Error(_) -> {
          case head {
            "do" -> process(rest, True, acc)
            "don't" -> process(rest, False, acc)
            _ -> process(rest, signal, acc)
          }
        }
      }
    }
  }
}

fn filter_mem(l: List(String)) -> List(Int) {
  process(l, True, [])
}

fn mult_pairs(l: List(Int)) -> List(Int) {
  case l {
    [a, b, ..rest] -> {
      list.append([a * b], mult_pairs(rest))
    }
    _ -> [0]
  }
}

pub fn part1(input: String) -> Int {
  let assert Ok(re) = regexp.from_string("mul\\((\\d+),(\\d+)\\)")

  input
  |> string.split("\n")
  |> list.map(fn(line: String) {
    regexp.scan(re, line)
    |> list.flat_map(fn(match: regexp.Match) { match.submatches })
  })
  |> list.flat_map(fn(sml) {
    list.map(sml, fn(sm) {
      case sm {
        Some(s) -> {
          case int.parse(s) {
            Ok(v) -> v
            Error(_) -> 1
          }
        }
        None -> 1
      }
    })
  })
  // |> io.debug
  |> mult_pairs()
  |> int.sum()
}

pub fn part2(input: String) -> Int {
  let assert Ok(re) = regexp.from_string("mul\\((\\d+),(\\d+)\\)|(do(?:n't)?)*")

  input
  |> string.split("\n")
  |> list.flat_map(fn(line: String) {
    regexp.scan(re, line)
    |> list.flat_map(fn(match: regexp.Match) { match.submatches })
  })
  |> option.values
  |> filter_mem()
  |> mult_pairs()
  |> int.sum()
}
