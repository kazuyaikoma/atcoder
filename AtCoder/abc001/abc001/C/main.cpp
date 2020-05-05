#include <bits/stdc++.h>
using namespace std;

const long long MOD = 5;

void solve(double deg, double dis){
  double ddeg = deg / 10;
  double ddis = round(dis/60*10)/10;

  string d1 = "N";
  int d2 = 0;
  if (11.25 <= ddeg && ddeg < 33.75) d1 = "NNE";
  else if (33.75 <= ddeg && ddeg < 56.25) d1 = "NE";
  else if (56.25 <= ddeg && ddeg < 78.75) d1 = "ENE";
  else if (78.75 <= ddeg && ddeg < 101.25) d1 = "E";
  else if (101.25 <= ddeg && ddeg < 123.75) d1 = "ESE";
  else if (123.75 <= ddeg && ddeg < 146.25) d1 = "SE";
  else if (146.25 <= ddeg && ddeg < 168.75) d1 = "SSE";
  else if (168.75 <= ddeg && ddeg < 191.25) d1 = "S";
  else if (191.25 <= ddeg && ddeg < 213.75) d1 = "SSW";
  else if (213.75 <= ddeg && ddeg < 236.25) d1 = "SW";
  else if (236.25 <= ddeg && ddeg < 258.75) d1 = "WSW";
  else if (258.75 <= ddeg && ddeg < 281.25) d1 = "W";
  else if (281.25 <= ddeg && ddeg < 303.75) d1 = "WNW";
  else if (303.75 <= ddeg && ddeg < 326.25) d1 = "NW";
  else if (326.25 <= ddeg && ddeg < 348.75) d1 = "NNW";

  if (0.3 <= ddis && ddis <= 1.5) d2 = 1;
  else if (1.6 <= ddis && ddis <= 3.3) d2 = 2;
  else if (3.4 <= ddis && ddis <= 5.4) d2 = 3;
  else if (5.5 <= ddis && ddis <= 7.9) d2 = 4;
  else if (8.0 <= ddis && ddis <= 10.7) d2 = 5;
  else if (10.8 <= ddis && ddis <= 13.8) d2 = 6;
  else if (13.9 <= ddis && ddis <= 17.1) d2 = 7;
  else if (17.2 <= ddis && ddis <= 20.7) d2 = 8;
  else if (20.8 <= ddis && ddis <= 24.4) d2 = 9;
  else if (24.5 <= ddis && ddis <= 28.4) d2 = 10;
  else if (28.5 <= ddis && ddis <= 32.6) d2 = 11;
  else if (32.7 <= ddis) d2 = 12;

  if (d2 == 0) d1 = "C";

  cout << d1 << " " << d2 << endl;
}

int main(){
    double deg;
    double dis;
    scanf("%lf",&deg);
    scanf("%lf",&dis);
    solve(deg, dis);
    return 0;
}

