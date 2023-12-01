use std::collections::HashMap;
use std::fs::read_to_string;

fn main() {
    

    part_one();
    part_two();
}

fn part_one() {
    let inputs: Vec<String> = read_to_string("src/input")
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    let mut nums: Vec<i32> = Vec::new();
    for input in inputs.iter() {
        let mut a = String::from("");
        let mut b = String::from("");

        for char in input.chars() {
            if char.is_numeric() {
                if a == "" {
                    a.push(char);
                } else {
                    b = char.to_string();
                }
            }
        }

        if b == "" {
            match (a.clone() + &a).parse::<i32>() {
                Ok(number) => {
                    println!("{number}");
                    nums.push(number);
                }
                Err(_) => {
                    println!("Fail")
                }
            }
        } else {
            match (a.clone() + &b).parse::<i32>() {
                Ok(number) => {
                    println!("{number}");
                    nums.push(number)
                }
                Err(_) => {
                    println!("Fail")
                }
            }
        }
    }

    let total: i32 = nums.iter().sum();
    println!("total: {total}");
}

fn part_two() {
    let inputs: Vec<String> = read_to_string("src/input")
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    let mut num_words: HashMap<String, String> = HashMap::new();
    let words: Vec<String> = vec![String::from("one"), String::from("two"), String::from("three"), String::from("four"), String::from("five"), String::from("six"), String::from("seven"), String::from("eight"), String::from("nine")];

    for (i, word) in words.iter().enumerate() {
        let _ = num_words.insert(word.to_string(), (i + 1).to_string());
    }

    let mut nums: Vec<i32> = Vec::new();
    for input in inputs.iter() {
        let mut a = String::from("");
        let mut b = String::from("");
        for (i, char) in input.chars().enumerate() {
            // If the current char is a numeric then we assign
            if char.is_numeric() {
                if a == "" {
                    a = char.to_string()
                }
                b = char.to_string()
            }

            let (ok, str) = found_number_word(i, input.chars().collect(), &words, &num_words);
            if ok {
                if a == "" {
                    a = str.clone();
                }
                b = str;
            }
        }

        match (a.clone() + &b).parse::<i32>() {
            Ok(number) => nums.push(number),
            Err(_) => {}
        }
    }

    let total: i32 = nums.iter().sum();
    println!("total: {total}");
}


fn found_number_word(index: usize, input_chars: Vec<char>, words: &Vec<String>, num_words: &HashMap<String, String>) -> (bool, String) {

    for num_word in words.iter() {
        let first_char = num_word.chars().next().unwrap();
        if first_char != input_chars[index] {
            // Continue and skip if the first char doesn't match
            continue;
        }

        let end_index = index + num_word.len();

        // Can't jump over the edge
        if end_index > input_chars.len() {
            continue;
        }

        let slice: String = input_chars[index..end_index].iter().collect();
        if slice == num_word.to_string() {
            return (true, num_words.get(num_word).unwrap().clone());
        }
    }

    return (false, String::new());
}
