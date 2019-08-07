use std::str::FromStr;
use std::fmt::Debug;

fn read_line<T>() -> Vec<T>
where T: FromStr, <T as FromStr>::Err : Debug {
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).unwrap();
    s.trim().split_whitespace().map(|c| T::from_str(c).unwrap()).collect()
}

fn main() {
    let vec: Vec<i32> = read_line();
    let (a, b, c) = (vec[0], vec[1], vec[2]);
    let minus = a - b;
    let ans = if c - minus < 0 { 0 } else { c - minus };
    println!("{}", ans);
}