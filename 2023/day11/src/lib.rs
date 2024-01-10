use itertools::iproduct;
use itertools::Itertools;
use std::collections::HashSet;

pub fn read_input(input: String) -> (Vec<(usize, usize)>, Vec<usize>, Vec<usize>, Vec<Vec<char>>) {
    let mut grid: Vec<Vec<char>> = Vec::with_capacity(input.lines().count());
    let mut galaxies: Vec<(usize, usize)> = Vec::new();
    let mut col_expansions = Vec::new();
    let mut row_expansions = Vec::new();

    for (y, line) in input.lines().enumerate() {
        grid.push(Vec::with_capacity(line.len()));

        if line.chars().all(|x| x == '.') {
            row_expansions.push(y)
        }

        for (x, ch) in line.chars().enumerate() {
            if ch == '#' {
                galaxies.push((x, y));
            }
            grid[y].push(ch);
        }
    }

    for x in 0..grid.len() {
        if grid.iter().map(|row| row[x]).all(|c| c == '.') {
            col_expansions.push(x);
        }
    }

    (galaxies, row_expansions, col_expansions, grid)
}

pub fn part_one_two(
    increase: usize,
    galaxies: Vec<(usize, usize)>,
    row_exp: Vec<usize>,
    col_exp: Vec<usize>,
) -> isize {
    let mut true_galaxies: Vec<(isize, isize)> = Vec::with_capacity(galaxies.len());
    for (x, y) in galaxies.iter() {
        let tx = x + ((increase - 1) * col_exp.iter().filter(|dx| dx < &x).count());
        let ty = y + ((increase - 1) * row_exp.iter().filter(|dy| dy < &y).count());

        true_galaxies.push((tx as isize, ty as isize));
    }

    let mut sum_dist: isize = 0;

    for (i, (a_x, a_y)) in true_galaxies.iter().enumerate() {
        for (b_x, b_y) in true_galaxies.iter().skip(i) {
            if a_x == b_x && a_y == b_y {
                continue;
            }

            // println!(
            //     "{:?}, {:?}, dist: {}",
            //     (a_x, a_y),
            //     (b_x, b_y),
            //     ((a_x - b_x).abs() + (a_y - b_y).abs())
            // );

            sum_dist += ((a_x - b_x).abs() + (a_y - b_y).abs());
        }
    }

    sum_dist
}
