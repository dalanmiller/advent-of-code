use std::fs::read_to_string;

#[test]
fn test_one() {
    let test_input = r#"???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1"#;

    let (springs, arr) = day12::read_input(test_input.to_string(), day12::Part::PartOne);

    let result = day12::part_one(springs, arr);

    assert_eq!(result, 21);
}

#[test]
fn test_part_one() {
    let test_input = read_to_string("src/input").unwrap();

    let (springs, arr) = day12::read_input(test_input.to_string(), day12::Part::PartOne);

    let result = day12::part_one(springs, arr);

    assert_eq!(result, 7705);
}

#[test]
fn test_two() {
    let test_input = r#"???.### 1,1,3
.??..??...?##. 1,1,3
?#?#?#?#?#?#?#? 1,3,1,6
????.#...#... 4,1,1
????.######..#####. 1,6,5
?###???????? 3,2,1"#;

    let (springs, arr) = day12::read_input(test_input.to_string(), day12::Part::PartTwo);

    let result = day12::part_one(springs, arr);

    assert_eq!(result, 525152);
}

#[test]
fn test_part_two() {
    let test_input = read_to_string("src/input").unwrap();

    let (springs, arr) = day12::read_input(test_input.to_string(), day12::Part::PartTwo);

    let result = day12::part_one(springs, arr);

    assert_eq!(result, 50338344809230);
}
