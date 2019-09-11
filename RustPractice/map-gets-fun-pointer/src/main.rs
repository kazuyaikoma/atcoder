fn main() {
    let v = vec!["I", "am", "NixieSquid"]
        .into_iter()
        // .map(|s| s.len()) を関数ポインタで書き換えた
        .map(str::len)
        .collect::<Vec<_>>();
    assert_eq!(v, vec![1, 2, 10]);
}
