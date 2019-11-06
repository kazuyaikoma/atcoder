use std::str::FromStr;
use std::fmt::Debug;
use std::cmp;

// 標準入力
fn read_line<T>() -> Vec<T>
where T: FromStr, <T as FromStr>::Err : Debug {
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).unwrap();
    s.trim().split_whitespace().map(|c| T::from_str(c).unwrap()).collect()
}

fn main() {
    let vec: Vec<usize> = read_line();
    let n = vec[0];
    let mut abcs: Vec<Vec<i32>> = Vec::new();
    let mut dp: Vec<Vec<i32>> = Vec::new();
    for _ in 0..n {
        let abc: Vec<i32> = read_line();
        abcs.push(abc);
        dp.push(Vec::new());
    }

    // calc
    dp[0] = abcs[0].clone();
    for i in 1..n {
        for j in 0..3 {
            let mut indices: Vec<usize> = vec![0, 1, 2];
            indices.remove(j);
            let score = cmp::max(
                dp[i - 1][indices[0]] + abcs[i][j],
                dp[i - 1][indices[1]] + abcs[i][j]
            );
            dp[i].push(score);
        }
    }

    let last = dp.last().unwrap();
    println!("{}", last.into_iter().max().unwrap());
}
