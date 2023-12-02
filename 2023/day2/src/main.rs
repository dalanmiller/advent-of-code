use std::cmp;
use std::fs::read_to_string;

struct Game {
    id: Option<i32>,
    max_blue: i32,
    max_green: i32,
    max_red: i32,
}

fn main() {
    part_one();
    part_two();
}

fn part_one() {
    let lines: Vec<String> = read_to_string("src/input")
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    let mut valid_games: Vec<Game> = Vec::new();

    for line in lines.iter() {
        let split: Vec<&str> = line.split(":").collect();
        let id: i32 = split[0]
            .to_string()
            .split(" ")
            .nth(1)
            .unwrap()
            .parse::<i32>()
            .unwrap();
        // println!("{}", id);
        let mut g = Game {
            id: Some(id.clone()),
            max_blue: 0,
            max_green: 0,
            max_red: 0,
            // sets: None,
        };

        let sets: Vec<String> = split[1]
            .to_string()
            .split(";")
            .map(|s| s.to_string())
            .collect();

        for set in &sets {
            let cubes: Vec<String> = set
                .trim_start()
                .split(", ")
                .map(|s| s.to_string())
                .collect();

            for cube in &cubes {
                let cube_details: Vec<String> = cube.split(" ").map(|s| s.to_string()).collect();
                let count = cube_details[0].parse::<i32>().unwrap();

                match cube_details[1].as_str() {
                    "red" => g.max_red = cmp::max(g.max_red, count),
                    "blue" => g.max_blue = cmp::max(g.max_blue, count),
                    "green" => g.max_green = cmp::max(g.max_green, count),
                    _ => println!("fail"),
                }
            }
        }

        let red_limit = 12;
        let green_limit = 13;
        let blue_limit = 14;

        if g.max_red <= red_limit && g.max_green <= green_limit && g.max_blue <= blue_limit {
            valid_games.push(g)
        }
    }

    let total: i32 = valid_games.iter().map(|g| g.id.unwrap()).sum();
    println!("total: {total}")
}

fn part_two() {
    let lines: Vec<String> = read_to_string("src/input")
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    let mut games: Vec<Game> = Vec::new();

    for line in lines.iter() {
        let split: Vec<&str> = line.split(":").collect();
        let mut g = Game {
            id: None,
            max_blue: 0,
            max_green: 0,
            max_red: 0,
            // sets: None,
        };

        let sets: Vec<String> = split[1]
            .to_string()
            .split(";")
            .map(|s| s.to_string())
            .collect();

        for set in &sets {
            let cubes: Vec<String> = set
                .trim_start()
                .split(", ")
                .map(|s| s.to_string())
                .collect();

            for cube in &cubes {
                let cube_details: Vec<String> = cube.split(" ").map(|s| s.to_string()).collect();
                let count = cube_details[0].parse::<i32>().unwrap();

                match cube_details[1].as_str() {
                    "red" => g.max_red = cmp::max(g.max_red, count),
                    "blue" => g.max_blue = cmp::max(g.max_blue, count),
                    "green" => g.max_green = cmp::max(g.max_green, count),
                    _ => println!("fail"),
                }
            }
        }

        games.push(g)
    }

    let total: i32 = games
        .iter()
        .map(|g| g.max_blue * g.max_green * g.max_red)
        .sum();
    println!("total: {total}")
}
