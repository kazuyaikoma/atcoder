fn main() {
    /// クロージャそれぞれが匿名の型(struct)を持つという説明

    let x = 4;
    // クロージャを定義し、xをそのクロージャに束縛。
    let addr = |n| n + x;
    assert_eq!(addr(2), 6);

    let mut state = false;

    // 別のクロージャを定義
    let mut flipflop = || {
        // stateが補足される
        state = !state;
        state
    };

    assert!(flipflop());    // true
    assert!(!flipflop());   // false
    assert!(flipflop());    // true

    // flipflopが補足してたstate自体も変わっていた
    assert!(state);         // true

    // つまり、クロージャを作るたびに別の匿名の型(struct)が作られる


    // 他の変数などを何も補足しないクロージャは関数ポインタになれる、という説明
    let mut f: fn(i32) -> i32 = |n| n * 3;
    assert_eq!(f(2), 6);

    let x = 4;
    /// ここで、
    /// f = |n| n * x;
    /// とするとエラー(mismatched types)になる
    /// なぜかというと、関数ポインタを期待してるのにクロージャが見つかったから。
}
