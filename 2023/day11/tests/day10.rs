use std::fs::read_to_string;

// ...#......
// .......#..
// #.........
// ..........
// ......#...
// .#........
// .........#
// ..........
// .......#..
// #...#.....

#[test]
fn test_read_input() {
    let test_input = r#"...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#....."#;

    let (galaxies, mut row_exp, mut col_exp, grid) = day11::read_input(test_input.to_string());

    assert_eq!(row_exp.len(), 2);
    row_exp.sort();
    assert_eq!(row_exp, vec![3, 7]);

    assert_eq!(col_exp.len(), 3);
    col_exp.sort();
    assert_eq!(col_exp, vec![2, 5, 8]);

    assert_eq!(galaxies.len(), 9);
}

#[test]
fn test_part_one_example_one() {
    let test_input = r#"...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#....."#;

    let (galaxies, row_exp, col_exp, _) = day11::read_input(test_input.to_string());
    let sum = day11::part_one_two(2, galaxies, row_exp, col_exp);

    assert_eq!(sum, 374);
}

#[test]
fn test_part_one() {
    let input = read_to_string("src/input").unwrap();
    let (galaxies, row_exp, col_exp, _) = day11::read_input(input);
    let sum = day11::part_one_two(2, galaxies, row_exp, col_exp);

    assert_eq!(sum, 9556896);
}

#[test]
fn test_part_two_example() {
    let test_input = r#"...#......
.......#..
#.........
..........
......#...
.#........
.........#
..........
.......#..
#...#....."#;

    let (galaxies, row_exp, col_exp, _) = day11::read_input(test_input.to_string());

    let sum = day11::part_one_two(10, galaxies.clone(), row_exp.clone(), col_exp.clone());

    assert_eq!(sum, 1030);

    let sum = day11::part_one_two(100, galaxies, row_exp, col_exp);

    assert_eq!(sum, 8410);
}

#[test]
fn test_part_two() {
    let input = read_to_string("src/input").unwrap();
    let (galaxies, row_exp, col_exp, _) = day11::read_input(input);
    let sum = day11::part_one_two(1000000, galaxies, row_exp, col_exp);

    assert_eq!(sum, 685038186836);

    // too high > 685038871866
    //            685038186836
}
