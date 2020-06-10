#include <bits/stdc++.h>
using namespace std;


void solve(int r, int D, int x){
  for (int i=0; i<10; i++) {
    x = r * x - D;
    cout << x << endl;
  }
}

int main(){
  int r;
  scanf("%d",&r);
  int D;
  scanf("%d",&D);
  int x;
  scanf("%d",&x);
  solve(r, D, x);
  return 0;
}
