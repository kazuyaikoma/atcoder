fn main() {
    println!("{}", sort("test"));
}

// 文字列のソート
fn sort(s: &str) -> String {
    let mut chars: Vec<_> = s.chars().collect();
    chars.sort();
    chars.into_iter().collect()
}
