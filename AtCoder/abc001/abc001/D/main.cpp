#include <bits/stdc++.h>
using namespace std;


string fixed(int t) {
  // それぞれ時間・分に直す
  int h = t/12;
  int m = (t%12)*5;
  return ( h<10 ? "0" : "" ) + to_string(h) + (m<10 ? "0" : "" ) + to_string(m);
}

int main(){
  long long N;
  scanf("%lld",&N);
  vector<string> A(N);
  // imos法のための配列
  int rain[12*24+1] = {};
  for(int i = 0 ; i < N ; i++){
    string str;
    cin >> str;

    int sh, sm, eh, em;
    sh = stoi(str.substr(0, 2));
    sm = stoi(str.substr(2, 2));
    eh = stoi(str.substr(5, 2));
    em = stoi(str.substr(7, 2));

    rain[sh*12+sm/5]++;
    // imos法では、区間の1つ先をデクリメントするため(5-1)をemに足している
    rain[eh*12+(em+(5-1))/5]--;
  }

  for (int i=1; i<12*24+1; i++) rain[i] += rain[i-1];

  for (int i=0; i<12*24+1; i++) {
    if (i == 0) {
      if(rain[i] > 0) cout << fixed(i) << "-";
    } else if (i == 288) {
      if (rain[i-1] > 0) cout << fixed(i) << endl;
    } else {
      if (rain[i] > 0) {
       if (rain[i-1] == 0) cout << fixed(i) << "-";
      } else if (rain[i-1] > 0) cout << fixed(i) << endl;
    }
  }
  return 0;
}
