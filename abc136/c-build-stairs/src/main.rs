use std::str::FromStr;
use std::fmt::Debug;

fn read_line<T>() -> Vec<T>
where T: FromStr, <T as FromStr>::Err : Debug {
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).unwrap();
    s.trim().split_whitespace().map(|c| T::from_str(c).unwrap()).collect()
}

fn main() {
    let _: Vec<i32> = read_line();
    let hs: Vec<i32> = read_line();
    let mut h: i32 = hs[0] - 1;
    let mut valid: bool = true;

    for i in 1..hs.len() {
        if h > hs[i] { valid = false; };
        if h == hs[i] { continue; };
        h = hs[i] - 1;
    };

    println!("{}", if valid { "Yes" } else { "No" });
}