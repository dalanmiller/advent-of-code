use std::fs::read_to_string;
use itertools::Itertools;

struct Rule {
    name: Option<String>, 
    source_start: isize, 
    source_end: isize, 
    dest_start: isize, 
    dest_end: isize,
}

impl Rule {
    fn convert(&self, n: isize) -> isize {
        // if n == 53 {
        //     println!("ohno")
        // }
        return self.dest_start + (n - self.source_start);
    }

    fn contains(&self, n: isize) -> bool {
        return self.source_start <= n && n <= self.source_end 
    }
}

fn read_input() -> (Vec<isize>, [Vec<Rule>; 7]) {

//     let lines: Vec<String> = r#"seeds: 79 14 55 13

// seed-to-soil map:
// 50 98 2
// 52 50 48

// soil-to-fertilizer map:
// 0 15 37
// 37 52 2
// 39 0 15

// fertilizer-to-water map:
// 49 53 8
// 0 11 42
// 42 0 7
// 57 7 4

// water-to-light map:
// 88 18 7
// 18 25 70

// light-to-temperature map:
// 45 77 23
// 81 45 19
// 68 64 13

// temperature-to-humidity map:
// 0 69 1
// 1 0 69

// humidity-to-location map:
// 60 56 37
// 56 93 4"#.lines().map(String::from).collect();

    let lines: Vec<String> = read_to_string("src/input")
        .unwrap()
        .lines()
        .map(String::from)
        .collect();

    // let split = lines.split("$\n");
    let seeds: Vec<isize> = lines[0][7..]
        .split(" ")
        .map(|n| n.parse::<isize>().unwrap())
        .collect();

    let (sts, stf, ftw, wtl, ltt, tth, htl): (
        Vec<Rule>,
        Vec<Rule>,
        Vec<Rule>,
        Vec<Rule>,
        Vec<Rule>,
        Vec<Rule>,
        Vec<Rule>,
    ) = (
        Vec::new(),
        Vec::new(),
        Vec::new(),
        Vec::new(),
        Vec::new(),
        Vec::new(),
        Vec::new(),
    );
    let mut rule_sets = [sts, stf, ftw, wtl, ltt, tth, htl];
    let mut i: usize = 0;
    let mut name: Option<String> = None; 
    for line in lines[1..].iter() {
        match line.as_str() {
            line if line.contains("map") => {
                i += 1;
                name = Some(line.to_string());
            },
            "\n" | "" => continue,
            _ => {
                let items: Vec<isize> = line
                    .split(" ")
                    .map(|n| n.parse::<isize>().unwrap())
                    .collect();
                let (dest, source, range) = (items[0], items[1], items[2]);
                rule_sets[i - 1].push(
                    Rule { name: name.clone(), source_start: source, source_end: (source+range -1), dest_start: (dest), dest_end: (dest + range -1)}
                );
            }
        }
    }

    (seeds, rule_sets)
}

fn main() {
    let (seeds, rule_sets) = read_input();
    part_one(&seeds, &rule_sets);
    part_two(&seeds, &rule_sets);
}

fn part_one(seeds: &Vec<isize>, rule_sets: &[Vec<Rule>; 7]) {
    let mut min = *seeds.iter().max().unwrap();
    for seed in seeds.iter() {
        let mut tmp = *seed;
        for rule_set in rule_sets.iter() {
            for rule in rule_set.iter() {
                if rule.contains(tmp) {
                    tmp = rule.convert(tmp);   
                    break
                }
            }
        }

        min = std::cmp::min(min, tmp);
    }

    println!("Answer: {:?}", min);
}

fn part_two(seeds: &Vec<isize>, rule_sets: &[Vec<Rule>; 7]) {

    let seed_ranges = seeds.chunks(2).map(|c| c[0]..(c[0]+c[1])).collect_vec();

    let mut index: usize = 0;
    let mut answer: Option<usize> = None;

    while answer.is_none() {
        // Our tmp variable that we are using to convert backwards
        // through all the maps 
        let mut tmp: usize = index;
        // println!("Starting with: {:?}", tmp);
        for rule_set in rule_sets.iter().rev() {
            for rule in rule_set {
                // println!("{:?}", rule.name);
                // We want to ensure that the current tmp variable is within 
                // the destination range 
                if rule.dest_start as usize <= tmp && tmp <= rule.dest_end as usize {
                    // This operations takes the output and reverses it into the input for the given
                    // rule
                    tmp = rule.source_start as usize + tmp - rule.dest_start as usize;
                    break;
                }
            }
        }
        // println!("Ending with: {:?}", tmp);

        for seed_range in seed_ranges.iter() {
            if seed_range.start as usize <= tmp && tmp <= seed_range.end as usize {
                // println!("Answer found: {:?}", index);
                answer = Some(index);
                break
            }   
        }
        // if seed_ranges.iter().any(|sr| sr.start as usize <= location && location <= sr.end as usize ) {
        //     answer = Some(index);
        // };
        index+=1;
    }
    
    println!("Answer: {:?}", answer.unwrap());

}
