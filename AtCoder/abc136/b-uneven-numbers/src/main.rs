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
    let n = vec[0];
    let mut cnt: i32 = 0;
    for i in 1..n+1 {
        if i.to_string().chars().count() % 2 != 0 {
            cnt += 1 
        };
    };
    println!("{}", cnt);
}