use std::collections::HashSet;
use std::collections::HashMap;
use std::fs::read_to_string;

fn read_input() -> Vec<String> {
    let lines: Vec<String> = read_to_string("src/input")
        .unwrap()
        .lines()
        .map(String::from)
        .collect();
//     let lines: Vec<String> = r#"Card 1: 41 48 83 86 17 | 83 86  6 31 17  9 48 53
// Card 2: 13 32 20 16 61 | 61 30 68 82 17 32 24 19
// Card 3:  1 21 53 59 44 | 69 82 63 72 16 21 14  1
// Card 4: 41 92 73 84 69 | 59 84 76 51 58  5 54 83
// Card 5: 87 83 26 28 32 | 88 30 70 12 93 22 82 36
// Card 6: 31 18 13 56 72 | 74 77 10 23 35 67 36 11"#.to_string().lines().map(String::from).collect();

    lines
}

fn main() {
    let lines = read_input();
    part_one(lines.clone());
    part_two(lines.clone());
}

struct Card {
    id: i32,
    // winning_numbers: HashSet<i32>,
    // my_numbers: HashSet<i32>,
    matches: usize,
}

fn part_one(lines: Vec<String>) {
    let mut cards: Vec<Card> = Vec::new();
    for line in lines.iter() {
        let split: Vec<&str> = line.split(|c: char| c == '|' || c == ':').collect();

        let id = split[0]
            .trim_start_matches(|c: char| !c.is_numeric())
            .parse::<i32>()
            .unwrap();

        let winning_numbers: HashSet<i32> = split[1].trim()
            .split(char::is_whitespace)
            .filter(|c| !c.is_empty())
            .map(|n| n.trim().parse::<i32>().unwrap())
            .collect();

        let my_numbers: HashSet<i32> = split[2]
            .split(char::is_whitespace)
            .filter(|c| !c.is_empty())
            .map(|n| n.parse::<i32>().unwrap())
            .collect();

        let matches = winning_numbers.intersection(&my_numbers).collect::<HashSet<_>>().len();

        cards.push(Card {
            id: id,
            // winning_numbers: winning_numbers,
            // my_numbers: my_numbers,
            matches: matches,
        })
    }

    let base: i32 = 2;
    let total: i32 = cards
        .iter()
        .filter(|c| c.matches > 0 )
        .map(|c| base.pow(c.matches as u32 - 1))
        .sum();

    println!("total: {total}");
}

fn part_two(lines: Vec<String>) {
    let mut cards: HashMap<i32, Card> = HashMap::new();

    for line in lines.iter() {
        let split: Vec<&str> = line.split(|c: char| c == '|' || c == ':').collect();

        let id = split[0]
            .trim_start_matches(|c: char| !c.is_numeric())
            .parse::<i32>()
            .unwrap();

        let winning_numbers: HashSet<i32> = split[1].trim()
            .split(char::is_whitespace)
            .filter(|c| !c.is_empty())
            .map(|n| n.trim().parse::<i32>().unwrap())
            .collect();

        let my_numbers: HashSet<i32> = split[2]
            .split(char::is_whitespace)
            .filter(|c| !c.is_empty())
            .map(|n| n.parse::<i32>().unwrap())
            .collect();

        let matches = winning_numbers.intersection(&my_numbers).collect::<HashSet<_>>().len();

        cards.insert(id,
        Card {
            id: id,
            // winning_numbers: winning_numbers,
            // my_numbers: my_numbers,
            matches: matches,
        });
    };

    let mut cards_queue: Vec<&Card> = Vec::new();      
    let mut card: &Card;
    let mut total: HashMap<i32, i32> = HashMap::new();
    for card in cards.values() {
        cards_queue.push(card);
        total.insert(card.id, 0);
    }

    // Iterate through cards until none left
    while cards_queue.len() > 0 {

        // Pop latest card
        card = cards_queue.pop().unwrap();
        
        // Get reference to value and increment by oen
        total.entry(card.id).and_modify(|value| *value += 1);

        // Only need to add cards to queue if there are matches
        if card.matches > 0 {
            let card_range = (card.id + 1)..((1 + card.matches + card.id as usize) as i32);
            // println!("id: {}, adding: {}", card.id, card_range.len());
            card_range.for_each(|n| cards_queue.push(&cards[&n]))
        }
    }

    println!("total: {}", total.values().sum::<i32>())

    // not 1205, too low
    // not 3533816, too low
}
