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
    let hs: Vec<i32> = read_line();

    let mut dp: Vec<i32> = Vec::new();
    dp.push(0);
    dp.push((hs[1] - hs[0]).abs());
    for _ in 2..hs.len() { dp.push(0); }
    for i in 2..hs.len() {
        dp[i] = cmp::min(
                    dp[i - 2] + (hs[i] - hs[i - 2]).abs(),
                    dp[i - 1] + (hs[i] - hs[i - 1]).abs(),
                );
    }
    println!("{}", dp.last().unwrap());
}
