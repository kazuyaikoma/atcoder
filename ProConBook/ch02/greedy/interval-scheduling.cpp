#include <iostream>
#include <algorithm>
using namespace std;

const int N = 5;
int starts[N] = {1, 2, 4, 6, 8};
int _ends[N] = {3, 5, 7, 9, 10};

pair<int, int> work_intervals[N];

int main() {
  // 先にwork_intervalsへstarts, endsを格納
  for (int i = 0; i < N; i++) {
    work_intervals[i].first = starts[i];
    work_intervals[i].second = _ends[i];
  }
  sort(work_intervals, work_intervals + N);

  int prev_endtime = 1;
  int cnt = 0;
  for (int i = 0; i < N; i++) {
    pair<int, int> work = work_intervals[i];
    if (work.first >= prev_endtime) {
      cnt += 1;
      prev_endtime = work.second;
    }
  }

  printf("%d\n", cnt);
  return 0;
}