use std::str::FromStr;
use std::fmt::Debug;
use std::collections::BinaryHeap;

fn read_line<T>() -> Vec<T>
where T: FromStr, <T as FromStr>::Err : Debug {
    let mut s = String::new();
    std::io::stdin().read_line(&mut s).unwrap();
    s.trim().split_whitespace().map(|c| T::from_str(c).unwrap()).collect()
}

fn main() {
    let vec: Vec<i32> = read_line();
    let arr: Vec<i64> = read_line();
    let (_n, m) = (vec[0], vec[1]);

    let mut heap: BinaryHeap<i64> = BinaryHeap::new();

    for i in 0..arr.len() { heap.push(arr[i]); }

    for _ in 0..m {
        let n: i64 = heap.pop().unwrap();
        heap.push(n / 2);
    }
    let ans: i64 = heap.iter().sum();
    println!("{}", ans);
}