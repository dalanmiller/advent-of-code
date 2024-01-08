use std::collections::HashMap;

#[derive(PartialEq, Clone, Copy, Eq, Hash)]
pub struct Coordinate {
    pub x: isize,
    pub y: isize,
    pub c: Option<Pipe>,
}

impl Coordinate {
    fn get_coords_in_bounds(
        &self,
        g: &Vec<Vec<Coordinate>>,
    ) -> (
        Option<Coordinate>,
        Option<Coordinate>,
        Option<Coordinate>,
        Option<Coordinate>,
    ) {
        let mut top: Option<Coordinate> = None;
        let mut right: Option<Coordinate> = None;
        let mut down: Option<Coordinate> = None;
        let mut left: Option<Coordinate> = None;

        // wtf are we doing here...
        // . Ahh okay, first we just check if the coordinates we are going to check
        // . are actually in bounds of the grid. Okay cool, let's keep it.
        if (self.y - 1) >= 0 {
            top = Some(g[self.y as usize - 1][self.x as usize]);
        }

        if (self.x + 1) < g.first().unwrap().len() as isize {
            right = Some(g[self.y as usize][self.x as usize + 1]);
        }

        if (self.y + 1) < g.len() as isize {
            down = Some(g[self.y as usize + 1][self.x as usize]);
        }

        if (self.x - 1) >= 0 as isize {
            left = Some(g[self.y as usize][self.x as usize - 1]);
        }

        return (top, right, down, left);
    }

    pub fn get_connected_adjacent(&self, g: &Vec<Vec<Coordinate>>) -> HashMap<Direction, Coordinate> {
        // First let's get the possibility of adjacent Coordiantes by checking
        // . in bounds.
        let (top, right, down, left) = self.get_coords_in_bounds(g);

        // let mut possible_paths: HashMap<Direction, Option<Coordinate>> = HashMap::new();
        let mut possible_paths: HashMap<Direction, Coordinate> = HashMap::new();
        let mut direction_check: HashMap<Direction, bool> = HashMap::new();

        // Now we construct a hash map to indicate the possible directions we _can go_
        // . given the type of Pipe in question. Also reasonable I guess.
        match self.c {
            Some(Pipe::J) => {
                direction_check = HashMap::from([
                    (Direction::UP, true),
                    (Direction::LEFT, true),
                    (Direction::RIGHT, false),
                    (Direction::DOWN, false),
                ])
            }
            Some(Pipe::F) => {
                direction_check = HashMap::from([
                    (Direction::UP, false),
                    (Direction::LEFT, false),
                    (Direction::RIGHT, true),
                    (Direction::DOWN, true),
                ])
            }
            Some(Pipe::HYPHEN) => {
                direction_check = HashMap::from([
                    (Direction::UP, false),
                    (Direction::LEFT, true),
                    (Direction::RIGHT, true),
                    (Direction::DOWN, false),
                ])
            }
            Some(Pipe::PIPE) => {
                direction_check = HashMap::from([
                    (Direction::UP, true),
                    (Direction::LEFT, false),
                    (Direction::RIGHT, false),
                    (Direction::DOWN, true),
                ])
            }
            Some(Pipe::SEVEN) => {
                direction_check = HashMap::from([
                    (Direction::UP, false),
                    (Direction::LEFT, true),
                    (Direction::RIGHT, false),
                    (Direction::DOWN, true),
                ])
            }
            Some(Pipe::L) => {
                direction_check = HashMap::from([
                    (Direction::UP, true),
                    (Direction::LEFT, false),
                    (Direction::RIGHT, true),
                    (Direction::DOWN, false),
                ])
            }
            Some(Pipe::START) => {
                direction_check = HashMap::from([
                    (Direction::UP, true),
                    (Direction::LEFT, true),
                    (Direction::RIGHT, true),
                    (Direction::DOWN, true),
                ])
            }
            None | Some(Pipe::GROUND) => {}
        }

        // Now we get the intersection of the directions we can go to given our pipe,
        // . and the available coordinates that are in bounds.
        //
        // For each of those conditions we need to check if the adjacent pipe can connect to
        // . our current pipe.

        // Check top
        if direction_check[&Direction::UP] {
            match top {
                Some(top) => {
                    if [Pipe::PIPE, Pipe::F, Pipe::SEVEN, Pipe::START].contains(&top.c.unwrap()) {
                        possible_paths.insert(Direction::UP, top);
                    }
                }
                None => {}
            }
        }

        // Check right
        if direction_check[&Direction::RIGHT] {
            match right {
                Some(right) => {
                    if [Pipe::HYPHEN, Pipe::J, Pipe::SEVEN, Pipe::START].contains(&right.c.unwrap())
                    {
                        possible_paths.insert(Direction::RIGHT, right);
                    }
                }
                None => {}
            }
        }

        // Check down
        if direction_check[&Direction::DOWN] {
            match down {
                Some(down) => {
                    if [Pipe::PIPE, Pipe::J, Pipe::L, Pipe::START].contains(&down.c.unwrap()) {
                        possible_paths.insert(Direction::DOWN, down);
                    }
                }
                None => {}
            }
        }

        // Check left
        if direction_check[&Direction::LEFT] {
            match left {
                Some(left) => {
                    if [Pipe::HYPHEN, Pipe::F, Pipe::L, Pipe::START].contains(&left.c.unwrap()) {
                        possible_paths.insert(Direction::LEFT, left);
                    }
                }
                None => {}
            }
        }

        return possible_paths;
    }
    fn get_all_adjacent(&self, g: &Vec<Vec<Coordinate>>) -> HashMap<Direction, Option<Coordinate>> {
        let mut all_adjacent: HashMap<Direction, Option<Coordinate>> = HashMap::new();
        for (k, v) in self.get_connected_adjacent(g).iter() {
            all_adjacent.insert(*k, Some(*v));
        }

        // We then fill in the missing coordinates
        for d in [
            Direction::UP,
            Direction::RIGHT,
            Direction::DOWN,
            Direction::LEFT,
        ] {
            if !all_adjacent.contains_key(&d) {
                all_adjacent.insert(d, None);
            }
        }

        all_adjacent
    }
}

#[derive(Copy, Clone, Eq, PartialEq, Hash)]
pub enum Direction {
    UP,
    RIGHT,
    DOWN,
    LEFT,
}

#[derive(PartialEq, Clone, Copy, Eq, Hash, Debug, Ord, PartialOrd)]
pub enum Pipe {
    F,
    GROUND,
    HYPHEN,
    J,
    L,
    PIPE,
    SEVEN,
    START,
}

impl Pipe {
    fn from_char(c: char) -> Option<Pipe> {
        match c {
            '|' => Some(Pipe::PIPE),
            '-' => Some(Pipe::HYPHEN),
            'J' => Some(Pipe::J),
            '7' => Some(Pipe::SEVEN),
            'L' => Some(Pipe::L),
            'F' => Some(Pipe::F),
            '.' => Some(Pipe::GROUND),
            'S' => Some(Pipe::START),
            _ => None,
        }
    }

    fn to_char(p: Pipe) -> char {
        match p {
            Pipe::PIPE => '|',
            Pipe::HYPHEN => '-',
            Pipe::J => 'J',
            Pipe::SEVEN => '7',
            Pipe::L => 'L',
            Pipe::F => 'F',
            Pipe::GROUND => '.',
            Pipe::START => 'S',
        }
    }
}

pub fn read_input(input: &str) -> (Coordinate, HashMap<Coordinate, bool>, Vec<Vec<Coordinate>>) {
    let mut start: Option<Coordinate> = None;
    let mut grid: Vec<Vec<Coordinate>> = Vec::new();
    for (y, line) in input.lines().enumerate() {
        grid.push(Vec::new());
        for (x, c) in line.chars().enumerate() {
            let pipe: Pipe = Pipe::from_char(c).unwrap();
            let coord: Coordinate = Coordinate {
                x: x as isize,
                y: y as isize,
                c: Some(pipe),
            };
            grid[y].push(coord);

            if pipe == Pipe::START {
                start = Some(coord);
            }
        }
    }

    let mut next_possible = start.unwrap().get_connected_adjacent(&grid);
    let (mut next_dir, mut current) = next_possible.drain().next().unwrap();
    let mut prev_dir = next_dir;
    let mut pipe: HashMap<Coordinate, bool> = HashMap::new();
    pipe.insert(current, true);

    while !pipe.contains_key(&start.unwrap()) {
        next_dir = next_direction(prev_dir, current.c.unwrap());
        current = next_coordinate(current, next_dir, &grid);
        prev_dir = next_dir;

        // Then this coordinate is inserted into the pipe hashmap
        if (current.x == 12 && current.y == 5) || (current.x == 2 && current.y == 2) {
            println!("wtf");
        }
        pipe.insert(current, true);
    }

    match start {
        None => panic!("start not found"),
        Some(start) => (start, pipe, grid),
    }
}

fn next_coordinate(
    current: Coordinate,
    direction: Direction,
    grid: &Vec<Vec<Coordinate>>,
) -> Coordinate {
    let (x, y): (isize, isize) = match direction {
        Direction::UP => (current.x, current.y.checked_sub(1).unwrap()),
        Direction::RIGHT => (current.x.checked_add(1).unwrap(), current.y),
        Direction::DOWN => (current.x, current.y.checked_add(1).unwrap()),
        Direction::LEFT => (current.x.checked_sub(1).unwrap(), current.y),
    };

    grid[y as usize][x as usize]
}

fn next_direction(prev: Direction, pipe: Pipe) -> Direction {
    match pipe {
        Pipe::PIPE => {
            match prev {
                Direction::UP => Direction::UP,
                // RIGHT => Direction::UP,
                Direction::DOWN => Direction::DOWN,
                // LEFT => Direction::UP,
                _ => panic!("Impossible previous dir"),
            }
        },
        Pipe::HYPHEN => {
            match prev {
                // UP => Direction::UP,
                Direction::RIGHT => Direction::RIGHT,
                // DOWN => Direction::DOWN,
                Direction::LEFT => Direction::LEFT,
                _ => panic!("Impossible previous dir"),
            }
        },
        Pipe::F => {
            match prev {
                Direction::UP => Direction::RIGHT,
                // RIGHT => Direction::RIGHT,
                // DOWN => Direction::DOWN,
                Direction::LEFT => Direction::DOWN,
                _ => panic!("Impossible previous dir"),
            }
        },
        Pipe::J => {
            match prev {
                // UP => Direction::UP,
                Direction::RIGHT => Direction::UP,
                Direction::DOWN => Direction::LEFT,
                // LEFT => Direction::LEFT,
                _ => panic!("Impossible previous dir"),
            }
        },
        Pipe::L => {
            match prev {
                // UP => Direction::UP,
                // RIGHT => Direction::RIGHT,
                Direction::DOWN => Direction::RIGHT,
                Direction::LEFT => Direction::UP,
                _ => panic!("Impossible previous dir"),
            }
        },
        Pipe::SEVEN => {
            match prev {
                Direction::UP => Direction::LEFT,
                Direction::RIGHT => Direction::DOWN,
                // DOWN => Direction::DOWN,
                // LEFT => Direction::LEFT,
                _ => panic!("Impossible previous dir"),
            }
        },
        _ => panic!("Impossible"),
    }
}

pub fn part_one(start: Coordinate, grid: Vec<Vec<Coordinate>>) -> isize {
    let possible_adjacent = start.get_connected_adjacent(&grid);
    let mut dist_map: HashMap<Coordinate, usize> = HashMap::new();
    let mut distance: usize = 1;

    let mut possible_directions: Vec<_> = possible_adjacent.keys().collect();
    let mut possible_coords: Vec<Coordinate> = possible_directions
        .iter()
        .map(|d| possible_adjacent[&d])
        .collect();

    let mut prev_dir: Direction = *possible_directions.pop().unwrap();
    let mut current = possible_coords.pop().unwrap();

    // First we go around starting with the 'right path'
    while current != start {
        dist_map.insert(current, distance);
        let next_dir = next_direction(prev_dir, current.c.expect("Always true"));
        prev_dir = next_dir;
        current = next_coordinate(current, next_dir, &grid);
        distance += 1;
    }

    // Then we reset variables and go around from the 'left path'
    distance = 1;

    prev_dir = *possible_directions.pop().unwrap();
    current = possible_coords.pop().unwrap();
    while current != start {
        // Get the current value for the current Coordinate, update
        // . if the current value is greater than what our current
        // .  distance is, we expect this to be true until the halfway mark
        match dist_map.get(&current) {
            Some(current_val) => {
                if *current_val > distance {
                    dist_map.insert(current, distance);
                }
            }
            None => continue,
        }

        // Determine the next direction, save it for next iteration
        // then update the current Coordinate
        let next_dir = next_direction(prev_dir, current.c.expect("Always true"));
        prev_dir = next_dir;
        current = next_coordinate(current, next_dir, &grid);

        distance += 1;
    }

    *dist_map.values().max().unwrap() as isize
}

pub fn part_two(pipe: HashMap<Coordinate, bool>, grid: Vec<Vec<Coordinate>>) -> isize {
    // First let's clean up the original grid of all the useless pipes and create a
    // a clean copy
    let mut pipe = pipe.clone();
    let mut clean_grid: Vec<Vec<Coordinate>> = Vec::new();

    for (y, row) in grid.iter().enumerate() {
        clean_grid.push(Vec::new());
        for c in row.iter() {
            if c.c.unwrap() == Pipe::START {
                let adj_map = c.get_all_adjacent(&grid);

                // UP, RIGHT, DOWN, LEFT
                let adj_coordinates: (Option<Pipe>, Option<Pipe>, Option<Pipe>, Option<Pipe>) = (
                    adj_map[&Direction::UP].and_then(|v| v.c),
                    adj_map[&Direction::RIGHT].and_then(|v| v.c),
                    adj_map[&Direction::DOWN].and_then(|v| v.c),
                    adj_map[&Direction::LEFT].and_then(|v| v.c),
                );

                let whoami: Option<Pipe> = match adj_coordinates {
                    // Hyphen
                    (_, Some(Pipe::J), _, Some(Pipe::L)) // L - J
                    | (_, Some(Pipe::SEVEN), _, Some(Pipe::F)) // F - 7
                    | (_, Some(Pipe::HYPHEN), _, Some(Pipe::HYPHEN)) // - - -
                    | (_, Some(Pipe::J), _, Some(Pipe::HYPHEN)) // - - J
                    | (_, Some(Pipe::SEVEN), _, Some(Pipe::HYPHEN)) // - - 7
                    | (_,Some(Pipe::HYPHEN), _, Some(Pipe::F)) // F - -
                    | (_, Some(Pipe::HYPHEN), _, Some(Pipe::L)) // L - -
                    => Some(Pipe::HYPHEN),

                    // Pipe
                    (Some(Pipe::F), _, Some(Pipe::L), _)
                    | (Some(Pipe::F), _, Some(Pipe::J), _)
                    | (Some(Pipe::SEVEN), _, Some(Pipe::L), _)
                    | (Some(Pipe::SEVEN), _, Some(Pipe::J), _)
                    | (Some(Pipe::PIPE), _, Some(Pipe::PIPE), _)
                    | (Some(Pipe::PIPE), _, Some(Pipe::J), _)
                    | (Some(Pipe::PIPE), _, Some(Pipe::L), _)
                    | (Some(Pipe::F), _, Some(Pipe::PIPE), _)
                    | (Some(Pipe::SEVEN), _, Some(Pipe::PIPE), _)
                    => Some(Pipe::PIPE),

                    // F
                    (_, Some(Pipe::HYPHEN), Some(Pipe::L), _)
                    | (_, Some(Pipe::HYPHEN), Some(Pipe::J), _)
                    | (_, Some(Pipe::SEVEN), Some(Pipe::L), _)
                    | (_, Some(Pipe::SEVEN), Some(Pipe::J), _)
                    | (_, Some(Pipe::HYPHEN), Some(Pipe::PIPE), _)
                    | (_, Some(Pipe::SEVEN), Some(Pipe::PIPE), _)
                    | (_, Some(Pipe::J), Some(Pipe::L), _)
                    | (_, Some(Pipe::J), Some(Pipe::J), _)
                    | (_, Some(Pipe::J), Some(Pipe::PIPE), _)
                    => Some(Pipe::F),

                    // J
                    (Some(Pipe::PIPE),_, _, Some(Pipe::F))
                    | (Some(Pipe::PIPE),_, _, Some(Pipe::SEVEN))
                    | (Some(Pipe::PIPE),_, _, Some(Pipe::HYPHEN))
                    | (Some(Pipe::F),_, _, Some(Pipe::F))
                    | (Some(Pipe::F),_, _, Some(Pipe::SEVEN))
                    | (Some(Pipe::F),_, _, Some(Pipe::HYPHEN))
                    | (Some(Pipe::SEVEN),_, _, Some(Pipe::F))
                    | (Some(Pipe::SEVEN),_, _, Some(Pipe::SEVEN))
                    | (Some(Pipe::SEVEN),_, _, Some(Pipe::HYPHEN))
                    => Some(Pipe::J),

                    // SEVEN
                    (_, _, Some(Pipe::PIPE), Some(Pipe::HYPHEN))
                    | (_, _, Some(Pipe::PIPE), Some(Pipe::F))
                    | (_, _, Some(Pipe::PIPE), Some(Pipe::L))
                    | (_, _, Some(Pipe::J), Some(Pipe::HYPHEN))
                    | (_, _, Some(Pipe::J), Some(Pipe::F))
                    | (_, _, Some(Pipe::J), Some(Pipe::L))
                    | (_, _, Some(Pipe::L), Some(Pipe::HYPHEN))
                    | (_, _, Some(Pipe::L), Some(Pipe::F))
                    | (_, _, Some(Pipe::L), Some(Pipe::L))
                    => Some(Pipe::J),

                    // L
                    (Some(Pipe::PIPE), Some(Pipe::HYPHEN), _, _)
                    | (Some(Pipe::PIPE), Some(Pipe::J), _, _)
                    | (Some(Pipe::PIPE), Some(Pipe::SEVEN), _, _)
                    | (Some(Pipe::SEVEN), Some(Pipe::HYPHEN), _, _)
                    | (Some(Pipe::SEVEN), Some(Pipe::J), _, _)
                    | (Some(Pipe::SEVEN), Some(Pipe::SEVEN), _, _)
                    | (Some(Pipe::F), Some(Pipe::HYPHEN), _, _)
                    | (Some(Pipe::F), Some(Pipe::J), _, _)
                    | (Some(Pipe::F), Some(Pipe::SEVEN), _, _)
                    => Some(Pipe::L),
                    _ => panic!("Help!")
                };

                match whoami {
                    Some(_whoami) => {
                        // Push identified Pipe for 'S' into the grid.
                        let s = Coordinate {
                            x: c.x,
                            y: c.y,
                            c: Some(_whoami),
                        };
                        clean_grid[y].push(s);

                        // Now we insert into the pipe to ensure that
                        // . we don't include this as an insider later on.
                        pipe.insert(s, true);
                    }
                    None => panic!("Help!"),
                }
            } else if pipe.contains_key(c) || c.c.unwrap() == Pipe::GROUND {
                clean_grid[y].push(*c);
            } else {
                clean_grid[y].push(Coordinate {
                    x: c.x,
                    y: c.y,
                    c: Some(Pipe::GROUND),
                })
            }
        }
    }

    for row in clean_grid.iter() {
        for c in row.iter() {
            print!("{}", Pipe::to_char(c.c.unwrap()))
        }
        println!();
    }

    // ..........
    // .S------7.
    // .|F----7|.
    // .||OOOO||.
    // .||OOOO||.
    // .|L-7F-J|.
    // .|II||II|.
    // .L--JL--J.
    // ..........

    let mut count = 0;

    for (y, row) in clean_grid.iter().enumerate() {
        // Reset for each row
        let mut p: Option<Pipe> = None;
        let mut inside: bool = false;

        for (x, c) in row.iter().enumerate() {
            // Taken from:
            // https://github.com/UncleScientist/aoclib-rs/blob/main/crates/aoc2023/src/aoc2023_10.rs#L127-L153

            // After watching a few videos of people solving Part 2, it seems like,
            // . one only has to check the row and not the column. (n=2)
            // . Hoping that this is the case for me as well.

            if pipe.contains_key(c) {
                match c.c {
                    Some(Pipe::F) => p = Some(Pipe::F),
                    Some(Pipe::GROUND) => {}
                    Some(Pipe::HYPHEN) => {}
                    Some(Pipe::J) => {
                        if p == Some(Pipe::F) {
                            inside = !inside
                        }
                    }
                    Some(Pipe::L) => p = Some(Pipe::L),
                    Some(Pipe::PIPE) => inside = !inside,
                    Some(Pipe::SEVEN) => {
                        if p == Some(Pipe::L) {
                            inside = !inside
                        }
                    }
                    Some(Pipe::START) => {}
                    None => panic!("Not possible"),
                }

                // println!(
                //     "PIPE | x: {}, y: {}, c: {:?}, inside: {:?}, pipe: {:?}",
                //     x, y, c.c.unwrap(), inside, p
                // )
            } else if inside {
                // println!("INSIDE | x: {}, y: {}, c: {:?}", x, y, c.c.unwrap());
                count += 1
            }
        }
    }

    count
}
