use itertools::Itertools;
use std::fs::read_to_string;

const TEST_INPUT: &str = r#"Time:      7  15   30
Distance:  9  40  200"#;

struct Race {
    time: isize,
    record_distance: isize,
}

fn read_input(input: &str) -> Vec<Race> {
    let mut races: Vec<Race> = Vec::new();

    let lines: Vec<&str> = input.lines().collect();
    let times: Vec<&str> = lines[0].split_whitespace().collect();
    let distances: Vec<&str> = lines[1].split_whitespace().collect();

    for (t, d) in times.iter().zip(distances.iter()) {
        match (t.parse::<isize>(), d.parse::<isize>()) {
            (Ok(t), Ok(d)) => races.push(Race {
                time: t,
                record_distance: d,
            }),
            (Err(_), Err(_)) => continue,
            (Ok(_), Err(_)) => continue,
            (Err(_), Ok(_)) => continue,
        }
    }

    races
}

fn main() {
    let test_races: &Vec<Race> = &read_input(TEST_INPUT);
    assert_eq!(part_one(test_races), 288);
    assert_eq!(part_two(test_races), 71503);

    let input: String = read_to_string("src/input").unwrap();
    let races: &Vec<Race> = &read_input(input.as_str());
    assert_eq!(part_one(races), 4811940);
    assert_eq!(part_two(races), 30077773);
}

fn part_one(races: &Vec<Race>) -> isize {
    let mut result: isize = 1;
    for race in races.iter() {
        let mut wins = 0;
        for n in 1..race.time {
            let distance = n * (race.time - n);
            if distance > race.record_distance {
                wins += 1;
            }
        }

        result *= wins;
    }

    result
}

fn part_two(races: &Vec<Race>) -> isize {
    let time = races
        .iter()
        .map(|n| n.time.to_string())
        .join("")
        .parse::<isize>()
        .unwrap();
    let record_distance = races
        .iter()
        .map(|n| n.record_distance.to_string())
        .join("")
        .parse::<isize>()
        .unwrap();

    let mut result: isize = 1;
    let mut wins = 0;
    for n in 1..time {
        let distance = n * (time - n);
        if distance > record_distance {
            wins += 1;
        }
    }

    result *= wins;

    result
}
