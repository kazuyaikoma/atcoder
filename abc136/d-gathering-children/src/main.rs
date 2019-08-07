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
    let vec: Vec<String> = read_line();
    let s = &vec[0];
    let cnt = s.len();
    let mut ans: Vec<i32> = Vec::new();
    for _ in 0..cnt { ans.push(0); };

    let mut r_cnt: i32 = 0;
    let mut l_cnt: i32 = 0;
    let mut r_last_idx: usize = 0;

    let mut chars: Vec<char> = s.chars().collect();
    // add invalid 'L' for correct loop execution
    chars.push('L');

    for i in 0..cnt {
        if chars[i] == 'R' && chars[i+1] == 'L' {
            r_cnt += 1;
            r_last_idx = i;
        } else if chars[i] == 'L' && chars[i+1] == 'R' || i == cnt-1 {
            l_cnt += 1;

            let sum = r_cnt + l_cnt;
            if sum % 2 == 0 {
                ans[r_last_idx] = sum / 2;
                ans[r_last_idx+1] = sum / 2;
            } else {
                ans[r_last_idx] = (sum - 1) / 2;
                ans[r_last_idx+1] = (sum - 1) / 2;
                if cmp::max(r_cnt, l_cnt) % 2 == 0 {
                    if r_cnt < l_cnt { ans[r_last_idx] += 1; } else { ans[r_last_idx+1] += 1; };
                } else {
                    if r_cnt > l_cnt { ans[r_last_idx] += 1; } else { ans[r_last_idx+1] += 1; };
                }
            };

            // reset R, L count variables
            r_cnt = 0;
            l_cnt = 0;
        } else {
            if chars[i] == 'R' { r_cnt += 1; } else { l_cnt += 1; };
        }
    }

    for a in ans.iter() { print!("{} ", a); };
    println!();
}