use regex::Regex;
use std::collections::HashMap;
use std::fs::read_to_string;

#[derive(Clone)]
struct Node {
    pub name: String,
    pub left: String,
    pub right: String,
}

enum Part {
    One,
    Two,
}

fn read_input(input: &str, part: Part) -> (Vec<char>, Vec<Node>, HashMap<String, Node>) {
    let lines: Vec<&str> = input.lines().collect();

    let directions: Vec<char> = lines[0].chars().collect();

    let mut heads: Vec<Node> = Vec::new();
    let mut nodes: HashMap<String, Node> = HashMap::new();

    let rex = Regex::new(r"(\w{3}) = \((\w{3}), (\w{3})\)").unwrap();
    for line in lines[2..].iter() {
        // let (cue, left, right) = rex.captures(line);
        let caps = rex.captures(line).unwrap();
        let name = caps.get(1).unwrap().as_str().to_string();

        let left = caps.get(2).unwrap().as_str().to_string();
        let right = caps.get(3).unwrap().as_str().to_string();
        let n = Node {
            name: name.clone(),
            left: left,
            right: right,
        };

        match part {
            Part::One => {
                if n.name == "AAA".to_string() {
                    heads.push(n.clone());
                }
            }
            Part::Two => {
                if n.name.ends_with("A") {
                    heads.push(n.clone());
                }
            }
        }

        nodes.insert(name, n);
    }

    return (directions, heads, nodes);
}

fn main() {
    let test_input = r#"LLR

AAA = (BBB, BBB)
BBB = (AAA, ZZZ)
ZZZ = (ZZZ, ZZZ)
"#;

    let test_input_two = r#"LR

11A = (11B, XXX)
11B = (XXX, 11Z)
11Z = (11B, XXX)
22A = (22B, XXX)
22B = (22C, 22C)
22C = (22Z, 22Z)
22Z = (22B, 22B)
XXX = (XXX, XXX)"#;

    let (mut dirs, mut heads, mut nodes) = read_input(test_input, Part::One);
    let test_result_one = part_one(dirs.clone(), heads.first().unwrap().clone(), nodes.clone());
    assert_eq!(test_result_one, 6);

    (dirs, heads, nodes) = read_input(test_input_two, Part::Two);
    let test_result_two = part_two(dirs.clone(), heads, nodes.clone());
    assert_eq!(test_result_two, 6);

    let input = read_to_string("src/input").unwrap();
    (dirs, heads, nodes) = read_input(input.as_str(), Part::One);
    let result_one = part_one(dirs.clone(), heads.first().unwrap().clone(), nodes.clone());
    assert_eq!(result_one, 12599);

    (dirs, heads, nodes) = read_input(input.as_str(), Part::Two);
    let result_two = part_two(dirs.clone(), heads, nodes.clone());
    assert_eq!(result_two, 8245452805243);
}

fn part_one(dirs: Vec<char>, head: Node, nodes: HashMap<String, Node>) -> isize {
    let mut head: &Node = &head;
    let mut steps = 0;

    let mut cycle_iter = dirs.iter().cycle();

    while let Some(&dir) = cycle_iter.next() {
        match dir {
            'L' => head = nodes.get(&head.left).unwrap(),
            'R' => head = nodes.get(&head.right).unwrap(),
            _ => continue,
        }

        steps += 1;

        if head.name == "ZZZ" {
            break;
        }
    }

    steps
}

fn gcd(a: isize, b: isize) -> isize {
    if b == 0 {
        a.abs()
    } else {
        gcd(b, a % b)
    }
}

fn lcm(a: isize, b: isize) -> isize {
    a.abs() / gcd(a, b) * b.abs()
}

fn lcm_of_vector(numbers: &[isize]) -> isize {
    numbers.iter().fold(1, |acc, &num| lcm(acc, num))
}

fn part_two(dirs: Vec<char>, mut heads: Vec<Node>, nodes: HashMap<String, Node>) -> isize {
    let mut steps = 0;
    let mut cycle_steps: Vec<isize> = vec![0; heads.len()];
    let mut cycle_iter = dirs.iter().cycle();

    while let Some(&dir) = cycle_iter.next() {
        heads = heads
            .iter()
            .map(|n| {
                if dir == 'L' {
                    n.left.clone()
                } else {
                    n.right.clone()
                }
            })
            .map(|name| nodes.get(&name).unwrap().clone())
            .collect();

        steps += 1;

        for (i, head) in heads.iter().enumerate() {
            if head.name.ends_with("Z") {
                // println!("Found Z for index: {i}, steps: {steps}");
                cycle_steps[i] = steps;
            }
        }

        if cycle_steps
            .iter()
            .all(|n| n.cmp(&0) != std::cmp::Ordering::Equal)
        {
            return lcm_of_vector(&cycle_steps);
        }
    }

    steps // Won't ever reach here
}
