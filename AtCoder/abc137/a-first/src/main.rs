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
    let (a, b) = (vec[0], vec[1]);
    let plus = a + b;
    let minus = a - b;
    let by = a * b;
    let p_m = cmp::max(plus, minus);
    println!("{}", cmp::max(by, p_m));
}