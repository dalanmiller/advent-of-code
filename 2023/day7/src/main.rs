use std::collections::HashMap;
use std::fs::read_to_string;

#[derive(PartialEq, Eq, PartialOrd, Ord, Hash, Debug, Copy, Clone)]
enum Card {
    A,
    K,
    Q,
    J,
    Ten,
    Nine,
    Eight,
    Seven,
    Six,
    Five,
    Four,
    Three,
    Two,
    None,
}

#[derive(PartialEq, Eq, PartialOrd, Ord, Hash, Debug, Copy, Clone)]
enum Card2 {
    A,
    K,
    Q,
    Ten,
    Nine,
    Eight,
    Seven,
    Six,
    Five,
    Four,
    Three,
    Two,
    J,
    None,
}

#[derive(PartialEq, Eq, PartialOrd, Ord, Copy, Clone)]
enum Type {
    FiveKind,
    FourKind,
    FullHouse,
    ThreeKind,
    TwoPair,
    OnePair,
    HighCard,
}

#[derive(Clone, Copy)]
struct Hand {
    bid: isize,
    hand: [SuperCard; 5],
    hand_type: Type,
}


impl Hand {
    fn cmp(&self, other: Hand) -> std::cmp::Ordering {
        if self.hand_type != other.hand_type {
            return self.hand_type.cmp(&other.hand_type);
        } else {
            return self.hand.cmp(&other.hand);
        }
    }
}

#[derive(PartialEq, Eq, PartialOrd, Ord, Hash, Debug, Copy, Clone)]
enum SuperCard {
    PartOne(Card),
    PartTwo(Card2),
}

impl From<SuperCard> for Card2 {
    fn from(card: SuperCard) -> Self {
        match card {
            SuperCard::PartOne(card) => {
                match card {
                    Card::A => Card2::A,
                    Card::K => Card2::K,
                    Card::Q => Card2::Q,
                    Card::J => Card2::J,
                    Card::Ten => Card2::Ten,
                    Card::Nine  => Card2::Nine,
                    Card::Eight => Card2::Eight, 
                    Card::Seven => Card2::Seven,
                    Card::Six => Card2::Six,
                    Card::Five => Card2::Five,
                    Card::Four => Card2::Four,
                    Card::Three => Card2::Three,
                    Card::Two => Card2::Two,
                    Card::None => Card2::None,
                }
            }
            _ => Card2::None,
        }
    }
}

fn get_card(c: char, part: i8) -> Option<SuperCard> {
        let card = match c {
            'A' => SuperCard::PartOne(Card::A),
            'K' => SuperCard::PartOne(Card::K),
            'Q' => SuperCard::PartOne(Card::Q),
            'J' => SuperCard::PartOne(Card::J),
            'T' => SuperCard::PartOne(Card::Ten),
            '9' => SuperCard::PartOne(Card::Nine),
            '8' => SuperCard::PartOne(Card::Eight),
            '7' => SuperCard::PartOne(Card::Seven),
            '6' => SuperCard::PartOne(Card::Six),
            '5' => SuperCard::PartOne(Card::Five),
            '4' => SuperCard::PartOne(Card::Four),
            '3' => SuperCard::PartOne(Card::Three),
            '2' => SuperCard::PartOne(Card::Two),
            _ => SuperCard::PartOne(Card::None),
        };
    
        if part == 2 {
            return Some(SuperCard::PartTwo(Card2::from(card)))
        }
        Some(card)
}

fn get_type(hand: &[SuperCard; 5], part: i8) -> Type {

    if hand.iter().all(|c| c == &hand[0]) {
        return Type::FiveKind;
    }

    let mut result: Option<Type> = None;
    let mut hm: HashMap<SuperCard, u8> = HashMap::new();
    hand.iter().for_each(|c| *hm.entry(*c).or_insert(0) += 1);

    if hm.values().any(|v| *v == 4) {
        result = Some(Type::FourKind);
    }

    let mut values: Vec<u8> = hm.values().cloned().collect::<Vec<_>>();
    values.sort();
    values.reverse();

    if values
        .iter()
        .take(2)
        .cloned()
        .collect::<Vec<_>>()
        .cmp(&[3, 2].to_vec())
        == std::cmp::Ordering::Equal
    {
        result = Some(Type::FullHouse);
    } else if hm.values().any(|x| *x == 3) {
        result = Some(Type::ThreeKind);
    } else if values
        .iter()
        .take(2)
        .cloned()
        .collect::<Vec<_>>()
        .cmp(&[2, 2].to_vec())
        == std::cmp::Ordering::Equal
    {
        result = Some(Type::TwoPair);
    } else if values
        .iter()
        .take(1)
        .cloned()
        .collect::<Vec<_>>()
        .cmp(&[2].to_vec())
        == std::cmp::Ordering::Equal
    {
        result = Some(Type::OnePair);
    }

    if part == 1 {
        match result {
            None => return Type::HighCard,
            Some(result) => return result,
        }
    }

    let n_jokers = hand.iter().filter(|c| [SuperCard::PartOne(Card::J), SuperCard::PartTwo(Card2::J)].contains(*c) ).count();

    let mut others_map : HashMap<SuperCard, i8> = HashMap::new();
    hand.iter().filter(|c| ![SuperCard::PartOne(Card::J), SuperCard::PartTwo(Card2::J)].contains(*c)).for_each(|c| *others_map.entry(*c).or_insert(0) += 1);
    let mut n_others: Vec<&i8> = others_map.values().collect::<Vec<_>>();
    n_others.sort();
    n_others.reverse();

    match n_jokers {
        5 => return Type::FiveKind,
        4 => return Type::FiveKind,
        3 => {
            // Need to count index one because 0 index would represent 
            // the three jokers
            match values[1] {
                2 => return Type::FiveKind,
                1 => return Type::FourKind,
                _ => println!("PANIC"),
            }
        },
        2 => {

            match n_others.as_slice() {
                [3] => return Type::FiveKind,
                [2,1] => return Type::FourKind,
                [1,1,1] => return Type:: ThreeKind,
                _ => println!("PANIC"),
            }
        },
        1 => {
            match n_others.as_slice() {
                [4] => return Type::FiveKind,
                [3, 1] => return Type::FourKind,
                [2, 2] => return Type::FullHouse,
                [2,1, 1] => return Type::ThreeKind,
                [1,1,1, 1] => return Type::OnePair,
                _ => println!("PANIC"),
            }
        }        
        _ => println!("PANIC"),   
    }

    match result {
        None => return Type::HighCard,
        Some(result) => return result,
    }
}

fn read_input(input: &str, part: i8) -> Vec<Hand> {
    let mut hands: Vec<Hand> = Vec::new();
    for line in input.lines() {
        let split: Vec<&str> = line.split(" ").collect();
        let hand: [SuperCard; 5] = split[0]
            .chars()
            .map(|c: char| get_card(c, part).unwrap())
            .collect::<Vec<_>>()
            .try_into()
            .unwrap();
        let bid: isize = split[1].parse::<isize>().unwrap();
        let hand_type = get_type(&hand, part);
        hands.push(Hand {
            bid,
            hand,
            hand_type,
        })
    }

    hands
}

fn main() {
    let test = r#"32T3K 765
T55J5 684
KK677 28
KTJJT 220
QQQJA 483"#;
    let mut test_input: Vec<Hand> = read_input(test, 1);
    assert_eq!(part_one(&mut test_input), 6440);
    test_input = read_input(test, 2);
    assert_eq!(part_two(&mut test_input), 5905);

    let mut input: Vec<Hand> = read_input(read_to_string("src/input").unwrap().as_str(), 1);
    assert_eq!(part_one(&mut input), 247961593);
    input = read_input(read_to_string("src/input").unwrap().as_str(), 2);
    assert_eq!(part_two(&mut input), 248750699);
}

fn part_one(hands: &mut Vec<Hand>) -> isize {
    hands.sort_by(|a, b| b.cmp(*a));

    hands
        .iter()
        .enumerate()
        .map(|(i, h)| ((i + 1) * h.bid as usize) as isize)
        .sum()
}

fn part_two(hands: &mut Vec<Hand>) -> isize {
    hands.sort_by(|a, b| b.cmp(*a));

    hands
        .iter()
        .enumerate()
        .map(|(i, h)| ((i + 1) * h.bid as usize) as isize)
        .sum()
}
