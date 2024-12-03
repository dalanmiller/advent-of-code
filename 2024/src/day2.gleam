import gleam/int
import gleam/io
import gleam/list
import gleam/string
import utils.{read_file}

pub fn main() {
  let input = case read_file("day2.input") {
    Ok(s) -> s
    _ -> "dunno"
  }
  // 999 too high
  io.debug(part1(input))
  io.debug(part2(input))
}

fn diff_is_within(a: Int, b: Int) -> Bool {
  let d = int.absolute_value(a - b)
  1 <= d && d <= 3
}

pub fn sublists_with_one_removed(l: List(a)) -> List(List(a)) {
  l
  |> list.index_map(fn(e, i) {
    let split_lists = list.split(l, i)
    list.flatten([split_lists.0, list.drop(split_lists.1, 1)])
  })
}

// fn level_problem_dampener_asc(l: List(Int)) -> Bool {
//   io.debug("asc")
//   {
//     // io.debug(pair)
//     l
//     |> list.zip(list.drop(l, 1))
//     |> list.map(fn(pair) {
//       // io.debug(pair)
//       case { pair.0 < pair.1 } && diff_is_within(pair.0, pair.1) {
//         True -> 0
//         False -> 1
//       }
//     })
//     |> io.debug
//     |> int.sum()
//   }
//   + {
//     // Conditional block that ensures that each pair is +/- 1-3 and accepts at most one pair that doesn't satisfy
//     l
//     |> list.zip(list.drop(l, 1))
//     |> list.map(fn(pair) {
//       // io.debug(pair)
//       case diff_is_within(pair.0, pair.1) {
//         True -> 0
//         False -> 1
//       }
//     })
//     |> io.debug
//     |> int.sum()
//   }
//   <= 1
// }

// fn level_problem_dampener_desc(l: List(Int)) -> Bool {
//   io.debug("desc")
//   {
//     l
//     |> list.zip(list.drop(l, 1))
//     |> list.map(fn(pair) {
//       case pair.0 > pair.1 {
//         True -> 0
//         False -> 1
//       }
//     })
//     |> io.debug
//     |> int.sum()
//   }
//   + {
//     // Conditional block that ensures that each pair is +/- 1-3 and accepts at most one pair that doesn't satisfy
//     l
//     |> list.zip(list.drop(l, 1))
//     |> list.map(fn(pair) {
//       case diff_is_within(pair.0, pair.1) {
//         True -> 0
//         False -> 1
//       }
//     })
//     |> io.debug
//     |> int.sum()
//   }
//   <= 1
// }

fn is_increasing(l: List(Int)) -> Bool {
  case l {
    [] -> True
    [_] -> True
    [a, b, ..rest] ->
      a < b && diff_is_within(a, b) && is_increasing([b, ..rest])
  }
}

fn is_decreasing(l: List(Int)) -> Bool {
  case l {
    [] -> True
    [_] -> True
    [a, b, ..rest] ->
      a > b && diff_is_within(a, b) && is_decreasing([b, ..rest])
  }
}

pub fn part1(input: String) -> Int {
  input
  |> string.split("\n")
  // Split lines -> List(String)
  |> list.map(fn(line) {
    // Each line, split on spaces, parse to int -> List(List(int))
    string.split(line, " ")
    |> list.filter_map(fn(n) {
      case int.parse(n) {
        Ok(n) -> Ok(n)
        Error(_) -> Error(Nil)
      }
    })
  })
  |> list.filter_map(fn(ints) -> Result(Int, Int) {
    case is_increasing(ints) || is_decreasing(ints) {
      True -> Ok(1)
      False -> Error(0)
    }
  })
  |> int.sum()
}

pub fn part2(input) -> Int {
  input
  |> string.split("\n")
  // Split lines -> List(String)
  |> list.map(fn(line) {
    // Each line, split on spaces, parse to int -> List(List(int))
    string.split(line, " ")
    |> list.filter_map(fn(n) {
      case int.parse(n) {
        Ok(n) -> Ok(n)
        Error(_) -> Error(Nil)
      }
    })
  })
  |> list.filter_map(fn(ints) {
    case ints
    |> sublists_with_one_removed
    |> list.any(fn(ints: List(Int)) -> Bool{
      is_increasing(ints) || is_decreasing(ints) 
    }) {
      True -> Ok(1)
      False -> Error(0)
    }
  })
  |> list.length
}
