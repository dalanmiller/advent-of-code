use day10;
use std::fs::read_to_string;

#[test]
fn test_map_parsing_and_pipe_generation() {
    let test_input = r#"..F7.
.FJ|.
SJ.L7
|F--J
LJ..."#;

    let (_, pipe, grid) = day10::read_input(test_input);
    let c = day10::Coordinate{
        x: 2,
        y: 1,
        c: Some(day10::Pipe::J),
    };

    assert!(pipe.get(&c).unwrap());

    // Count PIPE
    let pipe_count = pipe.iter().filter(|(k,_)| k.c.unwrap() == day10::Pipe::PIPE).count();
    assert_eq!(pipe_count, 2);

    // Count J
    let j_count = pipe.iter().filter(|(k,_)| k.c.unwrap() == day10::Pipe::J).count();
    assert_eq!(j_count, 4);

    // Count SEVEN
    assert_eq!(pipe.iter().filter(|(k,_)| k.c.unwrap() == day10::Pipe::SEVEN).count(), 2);

    // Count HYPHEN
    assert_eq!(pipe.iter().filter(|(k,_)| k.c.unwrap() == day10::Pipe::HYPHEN).count(), 2);

    // Count L
    assert_eq!(pipe.iter().filter(|(k,_)| k.c.unwrap() == day10::Pipe::L).count(), 2);

    // Count F
    assert_eq!(pipe.iter().filter(|(k,_)| k.c.unwrap() == day10::Pipe::F).count(), 3);


    assert_eq!(pipe.len(), 16);

    let next_possible = c.get_connected_adjacent(&grid);

    assert_eq!(next_possible.len(), 2);
}


#[test]
fn test_test_part_one() {
    let test_input = r#"..F7.
.FJ|.
SJ.L7
|F--J
LJ..."#;

    let (start, _, grid) = day10::read_input(test_input);
    let test_result = day10::part_one(start, grid);
    assert_eq!(test_result, 8);
}

#[test]
fn test_part_one() {
    let test_input = read_to_string("src/input").unwrap();
    let (start, _, grid) = day10::read_input(test_input.as_str());
    let test_result = day10::part_one(start, grid);
    assert_eq!(test_result, 6870);
}

#[test]
fn test_test_part_two() {
    let test_input = r#".F----7F7F7F7F-7....
.|F--7||||||||FJ....
.||.FJ||||||||L7....
FJL7L7LJLJ||LJ.L-7..
L--J.L7...LJS7F-7L7.
....F-J..F7FJ|L7L7L7
....L7.F7||L7|.L7L7|
.....|FJLJ|FJ|F7|.LJ
....FJL-7.||.||||...
....L---J.LJ.LJLJ..."#;

// 'I's here are corrected
// 
// OF----7F7F7F7F-7OOOO
// O|F--7||||||||FJOOOO
// O||OFJ||||||||L7OOOO
// FJL7L7LJLJ||LJIL-7OO
// L--JOL7IIILJS7F-7L7O
// OOOOF-JIIF7FJ|L7L7L7
// OOOOL7IF7||L7|IL7L7|
// OOOOO|FJLJ|FJ|F7|OLJ
// OOOOFJL-7O||O||||OOO
// OOOOL---JOLJOLJLJOOO


    let (_, pipe, grid) = day10::read_input(test_input);

    assert_eq!(grid[4][13].c.unwrap(), day10::Pipe::SEVEN);
    assert_eq!(grid[4][12].c.unwrap(), day10::Pipe::START);
    assert_ne!(grid[5][12].c.unwrap(), day10::Pipe::GROUND);

    let test_result = day10::part_two(pipe, grid);
    assert_eq!(test_result, 8);
}

#[test]
fn test_test_part_twopart_two() {
    let test_input = read_to_string("src/input").unwrap();
    let (_, pipe, grid) = day10::read_input(test_input.as_str());
    let test_result = day10::part_two(pipe, grid);
    assert_eq!(test_result, 287);
}