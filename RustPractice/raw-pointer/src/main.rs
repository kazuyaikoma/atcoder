fn main() {
    // pointer <*const T>
    let c1 = 'A';
    let c1_ptr: *const char = &c1;
    assert_eq!(unsafe { *c1_ptr }, 'A');

    // pointer <*mut T>
    let mut n1 = 0;
    let n1_ptr: *mut i32 = &mut n1;
    assert_eq!(unsafe { *n1_ptr }, 0);

    // increment *n1_ptr by using unsafe block
    unsafe {
        *n1_ptr = 1;
        assert_eq!(*n1_ptr, 1);
    }
}
