use std::fs::read_to_string;
use itertools::Itertools;

type Sequence = Vec<isize>;

fn read_input(input: String) -> Vec<Sequence> {
    let mut numbers: Vec<Sequence> = Vec::new();
    for line in input.lines() {
        let split: Vec<&str> = line.split_whitespace().collect();
        let seq: Sequence = split.iter().map(|n| n.parse::<isize>().unwrap()).collect();
        numbers.push(seq);
    }

    numbers
}

fn main() {
    let test_data = r#"0 3 6 9 12 15
1 3 6 10 15 21
10 13 16 21 30 45"#;

    let test_input = read_input(test_data.to_string());
    let test_result_one = part_one(test_input);
    assert_eq!(test_result_one, 114);

    let input_one = read_input(read_to_string("src/input").unwrap());
    let result_one = part_one(input_one);
    assert_eq!(result_one, 1743490457);

    //           1743490457
    // Too high, 1743490461
    // Too high, 1759382353

    let test_input_two = read_input(test_data.to_string());
    let test_result_two = part_two(test_input_two);
    assert_eq!(test_result_two, 2);

    let input_two = read_input(read_to_string("src/input").unwrap());
    let result_two = part_two(input_two);
    assert_eq!(result_two, 0);

    
}

fn part_one(sequences: Vec<Sequence>) -> isize {
    let mut values: isize = 0;
    
    for sequence in sequences.iter() {
        let mut temp: Vec<isize> = Vec::new();
        let mut temp_stack: Vec<Sequence> = vec![sequence.clone()];
        
        // Keep pushing until our last Sequence is 0s
        while !temp_stack.last().unwrap().iter().all(|n| *n == 0) {
            let s: Vec<isize> = temp_stack.last().unwrap().clone();

            for window in s.windows(2) {
                if let [a, b] = *window {
                    temp.push(b - a);
                }
            }

            temp_stack.push(temp);
            temp = Vec::new();
        }

        // Now we can fill in upwards, we start by popping off the 0-filled Sequence
        temp_stack.pop();

        let mut current: isize = 0;
        let mut current_seq: Sequence;
        while temp_stack.len() > 0 {
            current_seq = temp_stack.pop().unwrap();
            current += current_seq.last().unwrap();
        }

        values += current;
    }

    values
}

fn part_two(sequences: Vec<Sequence>) -> isize {
    let mut values: isize = 0;
    
    for sequence in sequences.iter() {
        let mut temp: Vec<isize> = Vec::new();
        let mut temp_stack: Vec<Sequence> = vec![sequence.clone()];
        
        // Keep pushing until our last Sequence is 0s
        while !temp_stack.last().unwrap().iter().all(|n| *n == 0) {
            let s: Vec<isize> = temp_stack.last().unwrap().clone();

            for window in s.windows(2) {
                if let [a, b] = *window {
                    temp.push(b - a);
                }
            }

            temp_stack.push(temp);
            temp = Vec::new();
        }

        // Now we can fill in upwards, we start by popping off the 0-filled Sequence
        temp_stack.pop();

        // This time we start with the 2nd to last vec
        let mut current: isize = *temp_stack.pop().unwrap().first().unwrap();
        let mut current_seq: Sequence;
        while temp_stack.len() > 0 {
            current_seq = temp_stack.pop().unwrap();
            current = current_seq.first().unwrap() - current;
        }

        values += current;
    }

    values
}
