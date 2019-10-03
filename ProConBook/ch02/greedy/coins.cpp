#include <iostream>
using namespace std;

const int NUM = 6;
const int Yen[NUM] = {500, 100, 50, 10, 5, 1};
const int Nums[NUM] = {2, 0, 3, 1, 2, 3};
int sum = 620;

int main() {
  int cnt = 0;
  for (int i = 0; i < NUM; i++) {
    int new_cnt = min(sum / Yen[i], Nums[i]);
    cnt += new_cnt;
    sum -= Yen[i] * new_cnt;
  }
  printf("%d\n", cnt);
  return 0;
}