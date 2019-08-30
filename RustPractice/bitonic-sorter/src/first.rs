pub fn sort(x: &mut [u32], up: bool) {
  if x.len() > 1 {
    let mid_point = x.len() / 2;
    sort(&mut x[..mid_point], true);
    sort(&mut x[mid_point..], false);
    sub_sort(x, up);
  }
}

fn sub_sort(x: &mut [u32], up: bool) {
  if x.len() > 1 {
    compare_and_swap(x, up);
    let mid_point = x.len() / 2;
    sub_sort(&mut x[..mid_point], up);
    sub_sort(&mut x[mid_point..], up);
  }
}

fn compare_and_swap(x: &mut [u32], up: bool) {
  let mid_point = x.len() / 2;
  for i in 0..mid_point {
    if (x[i] > x[mid_point + i]) == up {
      x.swap(i, mid_point + i);
    }
  }
}

#[cfg(test)]
mod tests {
  use super::sort;

  #[test]
  fn sort_u32_ascending() {
    let mut x = vec![8, 4, 2, 6, 3, 1, 5, 7];
    sort(&mut x, true);

    assert_eq!(x, vec![1, 2, 3, 4, 5, 6, 7, 8]);
  }

  #[test]
  fn sort_u32_descending() {
    let mut x = vec![8, 4, 2, 6, 3, 1, 5, 7];
    sort(&mut x, false);

    assert_eq!(x, vec![8, 7, 6, 5, 4, 3, 2, 1]);
  }
}