fn double(n: i32) -> i32 {
    n + n
}

fn abs(n: i32) -> i32 {
    if n >= 0 { n } else { -n }
}

fn main() {
    // get function pointer
    let mut f: fn(i32) -> i32 = double;
    assert_eq!(f(-42), -84);

    // change function pointer's value to abs
    f = abs;
    assert_eq!(f(-42), 42);
}
