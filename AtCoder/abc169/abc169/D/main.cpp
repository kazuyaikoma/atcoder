#include <bits/stdc++.h>
using namespace std;

bool isPrime(long long num, vector<long long>& primes) {
  if (num == 2) {
    primes.push_back(num);
    return true;
  }
  else if (num % 2 == 0) return false;

  double sqrtNum = sqrt(num);
  for (int i=3; i<=sqrtNum; i+=2) if (num % i == 0) return false;
  primes.push_back(num);
  return true;
}

bool isPrimes(long long z, vector<long long>& primes) {
  for (long long p : primes) {
    if (z % p == 0) return true;
  }
  return false;
}

void solve(long long N){
  if (N == 1) {
    cout << 0 << endl;
    return;
  }

  vector<long long> primes = {};
  long long cnt = 0;
  long long z = 2;
  while (N != 0) {
    while (N % z != 0 || (!isPrime(z, primes) && !isPrimes(z, primes))) {
       z = z + 1;
      cout << z << endl;
    }
    N /= z;
    cnt++;
  }

  cout << cnt << endl;
}

int main(){
    long long N;
    scanf("%lld",&N);
    solve(N);
    return 0;
}
