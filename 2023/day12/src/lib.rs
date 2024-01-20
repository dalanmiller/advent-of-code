use itertools::zip_eq;
use std::collections::HashMap;

#[derive(PartialEq, Eq)]
pub enum Part {
    PartOne,
    PartTwo,
}

pub fn read_input(input: String, part: Part) -> (Vec<Vec<char>>, Vec<Vec<usize>>) {
    let mut springs: Vec<Vec<char>> = Vec::new();
    let mut arrangements: Vec<Vec<usize>> = Vec::new();

    for line in input.lines() {
        let split: Vec<&str> = line.split(" ").collect();

        match part {
            Part::PartOne => {
                springs.push(split[0].chars().collect());
                arrangements.push(
                    split[1]
                        .split(",")
                        .map(|n| n.parse::<usize>().unwrap())
                        .collect(),
                );
            }
            Part::PartTwo => {
                springs.push(
                    [split[0], split[0], split[0], split[0], split[0]]
                        .join("?")
                        .chars()
                        .collect(),
                );
                arrangements.push(
                    [split[1], split[1], split[1], split[1], split[1]]
                        .join(",")
                        .split(",")
                        .map(|n| n.parse::<usize>().unwrap())
                        .collect(),
                );
            }
        }
    }

    (springs, arrangements)
}

// const MEMO: HashMap<(&[char], &[usize]), usize> = HashMap::with_capacity(100000);

pub fn count(
    springs: &[char],
    arrangement: &[usize],
    memo: &mut HashMap<(Vec<char>, Vec<usize>), usize>,
) -> usize {
    // We've run out of springs and so either we've satisfied the requirement
    // . or we haven't and it's a deadend.
    if springs.is_empty() {
        return if arrangement.is_empty() { 1 } else { 0 };
    }

    // We've run out of the arrangement and so either we have found a
    // . valid set or there's still more spring in the string.
    if arrangement.is_empty() {
        return if springs.contains(&'#') { 0 } else { 1 };
    }

    let key = (springs.to_vec(), arrangement.to_vec());
    match memo.get(&key) {
        Some(v) => return *v,
        _ => {}
    };

    let mut result = 0;

    if ['.', '?'].contains(&springs[0]) {
        result += count(&springs[1..], arrangement, memo);
    }

    if ['#', '?'].contains(&springs[0]) {
        if arrangement[0] <= springs.len()
            && !springs[0..arrangement[0]].contains(&'.')
            && (arrangement[0] == springs.len() || springs[arrangement[0]] != '#')
        {
            if arrangement[0] == springs.len() {
                result += count(&springs[arrangement[0]..], &arrangement[1..], memo);
            } else {
                result += count(&springs[(arrangement[0] + 1)..], &arrangement[1..], memo);
            }
        }
    }

    memo.insert(key, result);
    result
}

pub fn part_one(springs: Vec<Vec<char>>, arrangements: Vec<Vec<usize>>) -> usize {
    zip_eq(springs, arrangements)
        .map(|(spring, arrangement)| {
            let mut memo = HashMap::with_capacity(spring.len() * spring.len());
            count(&spring, &arrangement, &mut memo)
        })
        .sum()
}
