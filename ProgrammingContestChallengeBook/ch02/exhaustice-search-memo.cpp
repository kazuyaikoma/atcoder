#include <iostream>
using namespace std;

int fib(int n);
int memo[1000];

int main() {
  int n;
  cin >> n;
  cout << fib(n) << endl;
  return 0;
}

int fib(int n) {
  if (n <= 1) return n;
  // メモ化により演算をショートカットして高速化
  if (memo[n] != 0) return memo[n];
  return memo[n] = fib(n - 1) + fib(n - 2);
}