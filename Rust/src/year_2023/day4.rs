use std::fmt::Debug;
use std::fs::File;
use std::io::{self, Read};

/*
   https://adventofcode.com/2023/day/4
   Bogdan Bernovici
*/

struct Card {
    id: u32,
    win_nums: Vec<u32>,
    nums: Vec<u32>,
    score: u32,
}

impl Debug for Card {
    fn fmt(&self, f: &mut std::fmt::Formatter<'_>) -> std::fmt::Result {
        f.debug_struct("Card")
            .field("id", &self.id)
            .field("win_nums", &self.win_nums)
            .field("nums", &self.nums)
            .field("score", &self.score)
            .finish()
    }
}

fn from_str_arr_to_u32_arr<'a>(str: &'a str) -> Vec<u32> {
    str.split(|c| c == ' ')
        .filter(|c| !c.is_empty())
        .collect::<Vec<&str>>()
        .iter()
        .flat_map(|n| n.parse::<u32>())
        .collect::<Vec<u32>>()
}

pub fn execute() -> io::Result<()> {
    let file_path = "../Inputs/2023/Day4_.txt";
    let mut file = File::open(file_path)?;
    let mut contents = String::new();
    file.read_to_string(&mut contents)?;

    //Debug
    println!("File contents:\n{}", contents);

    let lines: Vec<&str> = contents.lines().collect();

    let numerics: Vec<String> = lines
        .iter()
        .map(|l| l.chars())
        .map(|chars| chars.filter(|&c| !c.is_alphabetic()).collect::<String>())
        .collect::<Vec<String>>();

    let cards: Vec<Card> = numerics
        .iter()
        .map(|s| s.split(|c| c == ':' || c == '|').collect::<Vec<&str>>())
        .map(|v| {
            let id = from_str_arr_to_u32_arr(v[0]);
            let win_nums = from_str_arr_to_u32_arr(v[1]);
            let nums = from_str_arr_to_u32_arr(v[2]);
            let count = nums.iter().filter(|n| win_nums.contains(n)).count();
            let mut score = 0;
            if count != 0 {
                score = (2 as u32).pow((count - 1).try_into().unwrap())
            }
            let card = Card {
                id: *id.first().unwrap(),
                win_nums,
                nums,
                score,
            };
            card
        }).collect();

    // Debug
    println!("\nCards:");
    cards.iter().for_each(|card| println!("{:?}", card));

    let sum: u32 = cards.iter().map(|c| c.score).sum();
    println!("\nPart 1: {}", sum);

    let counter = cards.iter().fold(0, |accumulator, card| {
        accumulator + go_through_cards(card, &cards)
    });
    println!("\nPart 2: {}", counter);

    Ok(())
}

fn go_through_cards(c: &Card, cards: &[Card]) -> u32 {
    let matching_numbers = ((c.score as f64).log2() + 1.0) as u32;
    let mut count = 1 as u32;
    for i in c.id..c.id+matching_numbers {
        count += go_through_cards(&cards[i as usize], cards)
    }
    return count;
}
