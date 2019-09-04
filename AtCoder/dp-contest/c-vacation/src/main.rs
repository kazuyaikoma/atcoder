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
    for _ in 0..n {
        let abc: Vec<i32> = read_line();
        abcs.push(abc);
    }

    // dp用の二重ベクタを作る
    let mut dp: Vec<Vec<i32>> = Vec::new();
    dp.push(abcs[0].clone());
    for _ in 1..n {
        let dp_abc: Vec<i32> = Vec::new();
        dp.push(dp_abc);
    }

    // 日数分、各ABCごとにdp計算
    'days: for i in 1..n {
        'dp_of_abc: for j in 0..3 {
            let mut indices: Vec<usize> = vec![0, 1, 2];
            // 前のdp値を見る際、ABCが被らないようにindicesからjを除外
            indices.remove(j);
            let v1 = dp[i - 1][indices[0]] + abcs[i][j];
            let v2 = dp[i - 1][indices[1]] + abcs[i][j];
            dp[i].push(cmp::max(v1, v2));
        }
    }

    let ans = [
        dp[n - 1][0],
        dp[n - 1][1],
        dp[n - 1][2],
    ];
    println!("{}", ans.iter().max().unwrap());
}
