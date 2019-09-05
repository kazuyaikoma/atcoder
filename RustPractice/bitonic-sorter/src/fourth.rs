use super::SortOrder;
use rayon;
use std::cmp::Ordering;

const PARALLEL_THRESHOLD: usize = 4096;

pub fn sort_by<T, F>(x: &mut [T], comparator: &F) -> Result<(), String>
  where T: Send,
        F: Sync + Fn(&T, &T) -> Ordering
{
  if x.len().is_power_of_two() {
    do_sort(x, true, comparator);
    Ok(())
  } else {
    Err(format!("The length of x is not a power of two. (x.len(): {})", x.len()))
  }
}

pub fn sort<T: Ord + Send>(x: &mut [T], order: &SortOrder) -> Result<(), String> {
  match *order {
    SortOrder::Ascending => sort_by(x, &|a, b| a.cmp(b)),
    SortOrder::Descending => sort_by(x, &|a, b| b.cmp(a)),
  }
}

fn do_sort<T, F>(x: &mut [T], forward: bool, comparator: &F)
  where T: Send,
        F: Sync + Fn(&T, &T) -> Ordering
{
  if x.len() > 1 {
    let mid_point = x.len() / 2;
    // divide x into two for parallel processing
    let (first, second) = x.split_at_mut(mid_point);
    if mid_point >= PARALLEL_THRESHOLD {
      rayon::join(
        || do_sort(first, true, comparator),
        || do_sort(second, false, comparator)
      );
    } else {
      do_sort(first, true, comparator);
      do_sort(second, false, comparator);
    }
    sub_sort(x, forward, comparator);
  }
}

fn sub_sort<T, F>(x: &mut [T], forward: bool, comparator: &F)
  where F: Fn(&T, &T) -> Ordering
{
  if x.len() > 1 {
    compare_and_swap(x, forward, comparator);
    let mid_point = x.len() / 2;
    sub_sort(&mut x[..mid_point], forward, comparator);
    sub_sort(&mut x[mid_point..], forward, comparator);
  }
}

fn compare_and_swap<T, F>(x: &mut [T], forward: bool, comparator: &F)
  where F: Fn(&T, &T) -> Ordering
{
  let swap_condition = if forward { Ordering::Greater } else { Ordering::Less };
  let mid_point = x.len() / 2;

  for i in 0..mid_point {
    if comparator(&x[i], &x[mid_point + i]) == swap_condition {
      x.swap(i, mid_point + i);
    }
  }
}

#[cfg(test)]
mod tests {
  use super::{sort_by, sort};
  use super::SortOrder::*;
  use crate::utils::{new_u32_vec, is_sorted};

  #[derive(Debug, PartialEq)]
  struct Student {
    first_name: String,
    last_name: String,
    age: u8,
  }

  impl Student {
    fn new(first_name: &str, last_name: &str, age: u8) -> Self {
      Self {
        first_name: first_name.to_string(),
        last_name: last_name.to_string(),
        age,
      }
    }
  }

  #[test]
  fn sort_students_by_age_ascending() {
    // 4 students test data
    let taro = Student::new("Taro", "Yamada", 16);
    let hanako = Student::new("Hanako", "Yamada", 14);
    let kyoko = Student::new("Kyoko", "Ito", 15);
    let ryosuke = Student::new("Ryosuke", "Hayashi", 17);

    let mut x = vec![&taro, &hanako, &kyoko, &ryosuke];
    let expected = vec![&hanako, &kyoko, &taro, &ryosuke];
    
    assert_eq!(
      sort_by(&mut x, &|a, b| a.age.cmp(&b.age)),
      Ok(())
    );

    assert_eq!(x, expected);
  }

  #[test]
  fn sort_students_by_name_ascending() {
    let taro = Student::new("Taro", "Yamada", 16);
    let hanako = Student::new("Hanako", "Yamada", 14);
    let kyoko = Student::new("Kyoko", "Ito", 15);
    let ryosuke = Student::new("Ryosuke", "Hayashi", 17);

    let mut x = vec![&taro, &hanako, &kyoko, &ryosuke];
    let expected = vec![&ryosuke, &kyoko, &hanako, &taro];

    assert_eq!(
      sort_by(&mut x,
        &|a, b| a.last_name.cmp(&b.last_name)
          .then_with(|| a.first_name.cmp(&b.first_name))
      ),
      Ok(())
    );

    assert_eq!(x, expected);
  }

  #[test]
  fn sort_u32_large() {
    {
      // generate randn of 2 to the 16th power (= 65536)
      let mut x = new_u32_vec(65536);
      // ascending sort
      assert_eq!(sort(&mut x, &Ascending), Ok(()));
      // confirm that the result is collect
      assert!(is_sorted(&x, &Ascending));
    }
    {
      // generate randn of 2 to the 16th power (= 65536)
      let mut x = new_u32_vec(65536);
      // ascending sort
      assert_eq!(sort(&mut x, &Descending), Ok(()));
      // confirm that the result is collect
      assert!(is_sorted(&x, &Descending));
    }
  }
}