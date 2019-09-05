use rand::{Rng, SeedableRng};
use rand::distributions::Standard;
use rand_pcg::Pcg64Mcg;
use super::SortOrder;

pub fn new_u32_vec(n: usize) -> Vec<u32> {
  // init RNG. use seed in order to reproducibility
  let mut rng = Pcg64Mcg::from_seed([0; 16]);
  rng.sample_iter(&Standard).take(n).collect()
}

pub fn is_sorted<T: Ord>(x: &[T], order: &SortOrder) -> bool {
  match *order {
    SortOrder::Ascending => x.windows(2).all(|pair| pair[0] <= pair[1]),
    SortOrder::Descending => x.windows(2).all(|pair| pair[0] >= pair[1]),
  }
}
