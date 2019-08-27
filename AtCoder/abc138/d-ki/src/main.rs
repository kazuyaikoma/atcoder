// This is NOT COMPLETED (not got AC) code
use std::str::FromStr;
use std::fmt::Debug;
use std::cmp;

fn read_line<T>() -> Vec<T>
where T: FromStr, <T as FromStr>::Err : Debug {
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).unwrap();
    s.trim().split_whitespace().map(|c| T::from_str(c).unwrap()).collect()
}

fn main() {
    let vec: Vec<i32> = read_line();
    let (n, q) = (vec[0], vec[1]);

    let mut nodes = Vec<i32>::new();
    for i in (1..n+1) { nodes.append(i) }
    for i in (0..n) { 
        let input: Vec<i32> = read_line();
        let idx = nodes.iter().position(|&e| e == input[0]) + 1;
        nodes.remove(input[1]);
        nodes.insert(idx, input[1]);
    }
    
    let mut values = Vec<i32>::new();
    for i in (0..n){ values.push(0); }
    for i in (0..q) {
        let input: Vec<i32> = read_line();
        let idx = nodes.iter().position(|&e| e == input[0]);
        let slice = &nodes[idx .. n - 1];
        
        for i in (0..slice.size) { values[i] += input[1]; }
     }

    for i in (0..n-1) { print!("{} ", values[i]); }
    println!();
}