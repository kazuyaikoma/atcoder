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
    let vec: Vec<i64> = read_line();
    let n = vec[0];
    let mut num = 0;

    for i in 1..n { num += i; }
    if n == 1 {
        println!("{}", 0);
    } else {
        println!("{}", num);
    }
}