import day1
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

// gleeunit test functions end in `_test`
pub fn day_1_part_1_test() {
  "3   4
4   3
2   5
1   3
3   9
3   3
"
  |> day1.part1
  |> should.equal(11)
}

pub fn day_1_part_2_test() {
  "3   4
4   3
2   5
1   3
3   9
3   3
"
  |> day1.part2
  |> should.equal(31)
}
