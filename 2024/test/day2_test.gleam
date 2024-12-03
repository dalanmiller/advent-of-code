import day2
import gleeunit
import gleeunit/should

pub fn main() {
  gleeunit.main()
}

// gleeunit test functions end in `_test`
pub fn day_1_part_1_test() {
  "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"
  |> day2.part1
  |> should.equal(2)
}

pub fn sublist_test() {
  [1, 2, 3]
  |> day2.sublists_with_one_removed
  |> should.equal([[2, 3], [1, 3], [1, 2]])
}

pub fn ascending_test() {
  "1 2 3 4"
  |> day2.part1
  |> should.equal(1)
}

pub fn descending_test() {
  "4 3 2 1"
  |> day2.part1
  |> should.equal(1)
}

pub fn desc_row_with_equal_test() {
  "4 2 2 1"
  |> day2.part1
  |> should.equal(0)
}

pub fn asc_row_with_equal_test() {
  "1 2 2 4"
  |> day2.part1
  |> should.equal(0)
}

pub fn extra_test() {
  "1 1 1 1"
  |> day2.part1
  |> should.equal(0)

  "1 2 3 4"
  |> day2.part1
  |> should.equal(1)

  "4 3 2 1"
  |> day2.part1
  |> should.equal(1)
}

pub fn day_2_part_2_test() {
  "7 6 4 2 1
1 2 7 8 9
9 7 6 2 1
1 3 2 4 5
8 6 4 4 1
1 3 6 7 9"
  |> day2.part2
  |> should.equal(4)
}

pub fn unsafe_test() {
  "1 2 7 8 9
9 7 6 2 1"
  |> day2.part2
  |> should.equal(0)
}

pub fn day_2_part_2_ascending_test() {
  "1 3 2 4 5"
  |> day2.part2
  |> should.equal(1)

  "1 2 3 4 5"
  |> day2.part2
  |> should.equal(1)

  "1 2 3 4 3"
  |> day2.part2
  |> should.equal(1)
}

pub fn day_2_part_2_desc_test() {
  "5 4 3 2 1"
  |> day2.part2
  |> should.equal(1)

  "5 6 4 3 2"
  |> day2.part2
  |> should.equal(1)

  "5 4 3 2 5"
  |> day2.part2
  |> should.equal(1)
}
