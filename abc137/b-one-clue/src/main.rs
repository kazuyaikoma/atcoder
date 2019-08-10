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
    let (k, x) = (vec[0], vec[1]);
    let min = x - (k - 1);
    let max = x + (k - 1);

    for i in min..max+1 { print!("{} ", i); }
    println!();
}