use std::fs::read_to_string;
use std::collections::HashSet;

fn gather_input() -> Vec<String> {
    let lines: Vec<String> = read_to_string("src/input")
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    lines
}

fn main() {
    let input = gather_input();
//     let input: Vec<String> = r#"467..114..
// ...*......
// ..35..633.
// ......#...
// 617*......
// .....+.58.
// ..592.....
// ......755.
// ...$.*....
// .664.598.."#.to_string().lines().map(String::from).collect();
    part_one(input.clone());
    part_two(input.clone());
}

struct Grid(Vec<Vec<char>>);

impl Grid {
    // fn new(vec: Vec<Vec<char>>) -> Self {
    //     Grid(vec)
    // }
    // fn len(&self) -> usize {
    //     self.0.len()
    // }

    fn valid(&self, x: isize, y: isize) -> bool {
        x >= 0 && y >= 0 && x < (*self).0.len() as isize && y < (*self).0.len() as isize
    }

    fn push(&mut self, e: Vec<char>) {
        self.0.push(e);
    }

    fn iter(&self) -> std::slice::Iter<'_, Vec<char>> {
        self.0.iter()
    }

    fn adjacent(&self, x: isize, y: isize) -> Vec<(isize, isize)> {
        let mut adj: Vec<(isize, isize)> = Vec::new();
        for dx in -1..2 {
            for dy in -1..2 {
                if dx == 0 && dy == 0 { continue }

                if self.valid(x + dx, y + dy) {
                    adj.push((x + dx, y + dy));
                }
            }
        }
        adj
    }
}

impl IntoIterator for Grid {
    type Item = Vec<char>;
    type IntoIter = std::vec::IntoIter<Self::Item>;

    fn into_iter(self) -> Self::IntoIter {
        self.0.into_iter()
    }
}

fn make_grid(lines: Vec<String>) -> Grid {
    let mut grid: Grid = Grid(vec![]);
    for line in lines.iter() {
        grid.push(line.chars().collect());
    }

    grid
}

#[derive(Hash, Eq, PartialEq, Debug, Copy, Clone)]
struct Number {
    xs: isize,
    xe: isize,
    y: isize,
    n: isize,
    adj_symbol: bool,
}

fn part_one(lines: Vec<String>) {
    let grid = make_grid(lines.clone());

    let mut numbers: Vec<Number> = Vec::new();
    for (y, line) in lines.clone().iter().enumerate() {
        let mut snumber: String = "".to_string();
        let mut start: isize = 0;
        for (i, char) in line.chars().enumerate() {
            match char {
                char if char.is_numeric() => {
                    if snumber == "".to_string() {
                        start = i as isize;
                    }
                    snumber.push(char);
                },
                '#' | '$'| '%' | '&'| '*'| '+'| '-' | '/' | '=' | '@' | '.' => {
                    if snumber != "".to_string() {
                        numbers.push(Number {
                            xs: start as isize,
                            xe: (i - 1) as isize,
                            y: y as isize,
                            n: snumber.parse::<isize>().unwrap(),
                            adj_symbol: false, 
                        });
                        snumber = "".to_string();
                    }
                    
                },
                _ => println!("fail"),
            }    
        }
        if snumber != "".to_string() {
            numbers.push(Number {
                xs: start as isize,
                xe: 139 as isize,
                y: y as isize,
                n: snumber.parse::<isize>().unwrap(),
                adj_symbol: false, 
            });
        }
    }

    for (y, col) in grid.iter().enumerate() {
        for (x, chr) in col.iter().enumerate() {
            let adj: Vec<(isize, isize)> = match chr {
                '#' | '$'| '%' | '&'| '*'| '+'| '-' | '/' | '=' | '@' => grid.adjacent(x as isize, y as isize),
                _ => vec![],
            };

            if adj.len() == 0 {
                continue;
            }

            // Loop through the adjacent coordinates around the symbol
            for (x, y) in adj.iter() {
                // Then we loop through all the numbers we found and
                // see if there is intersection between the coords around
                // those numbers and the current symbol
                for n in numbers.iter_mut() {
                    // If the below is true then we can mark the number as
                    // being adjacent to a symbol
                    // if n.n == 89 && *chr == '*' && n.xs > 119 {
                        // println!("here we go")
                    // }
                    if ((n.xs <= *x) && (*x <= n.xe)) && (n.y == *y) {
                        n.adj_symbol = true;
                    }
                }
            }
        }
    }

    for number in numbers.iter() {
        if number.adj_symbol == true {
            println!("{}", number.n);
        }
    }

    let total: isize = numbers.iter().filter(|n| n.adj_symbol).map(|n| n.n).sum();
    println!("total: {total}")

    // not 408271 too low 
    // not 578552 too high
    // not 497433 too ??? 
    // not 497027   ``
}



fn part_two(lines: Vec<String>) {

    let grid = make_grid(lines.clone());

    let mut numbers: Vec<Number> = Vec::new();
    for (y, line) in lines.clone().iter().enumerate() {
        let mut snumber: String = "".to_string();
        let mut start: isize = 0;
        for (i, char) in line.chars().enumerate() {
            match char {
                char if char.is_numeric() => {
                    if snumber == "".to_string() {
                        start = i as isize;
                    }
                    snumber.push(char);
                },
                '#' | '$'| '%' | '&'| '*'| '+'| '-' | '/' | '=' | '@' | '.' => {
                    if snumber != "".to_string() {
                        numbers.push(Number {
                            xs: start as isize,
                            xe: (i - 1) as isize,
                            y: y as isize,
                            n: snumber.parse::<isize>().unwrap(),
                            adj_symbol: false, 
                        });
                        snumber = "".to_string();
                    }
                    
                },
                _ => println!("fail"),
            }    
        }
        if snumber != "".to_string() {
            numbers.push(Number {
                xs: start as isize,
                xe: 139 as isize,
                y: y as isize,
                n: snumber.parse::<isize>().unwrap(),
                adj_symbol: false, 
            });
        }
    }

    let mut sum = 0;
    for (y, col) in grid.iter().enumerate() {
        for (x, chr) in col.iter().enumerate() {
            let adj: Vec<(isize, isize)> = match chr {
                '*' => grid.adjacent(x as isize, y as isize),
                _ => continue,
            };

            let mut part_numbers: HashSet<Number> = HashSet::new();
            // Loop through the adjacent coordinates around the symbol
            'outer: for (x, y) in adj.iter() {
                // Then we loop through all the numbers we found and
                // see if there is intersection between the coords around
                // those numbers and the current symbol
                for n in numbers.iter().filter(|n| n.y <= y + 1 && n.y >= y-1 ) {
                    // If the below is true then we can mark the number as
                    // being adjacent to a symbol
                    // if n.n == 89 && *chr == '*' && n.xs > 119 {
                        // println!("here we go")
                    // }
                    if ((n.xs <= *x) && (*x <= n.xe)) && (n.y == *y) {
                        let new_n = n.clone();
                        part_numbers.insert(new_n);
                    }

                    if part_numbers.len() > 2 {
                        break 'outer; 
                    }
                }
            }

            if part_numbers.len() == 2 {
                sum += part_numbers.iter().map(|n| n.n).product::<isize>();
            }
        }
    }

    println!("{}", sum);

    // not 29458644, too low
    // not 290151, too low
}
