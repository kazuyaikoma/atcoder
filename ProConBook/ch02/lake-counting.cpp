#include <iostream>
using namespace std;

void dfs(int x, int y);
const int N = 5;
const int M = 6;
// この場合だと、正しい答えは4になる
bool field[N][M] = {
  {true, false, false, false, true, true},
  {true, true, false, false, false, true},
  {false, false, false, false, false, false},
  {true, false, false, false, true, true},
  {true, true, false, false, true, false},
};

int main() {
  int count = 0;
  for (int x = 0; x < N; x++) {
    for (int y = 0; y < M; y++) {
      if (field[x][y]) {
        count++;
        dfs(x, y);
      }
    }
  }
  printf("%d\n", count);
  return 0;
}

void dfs(int x, int y) {
  // 現在地をfalseにする
  field[x][y] = false;
  // 8近傍を見る
  for (int dx = -1; dx <= 1; dx++) {
    for (int dy = -1; dy <= 1; dy++) {
      int nx = x + dx, ny = y + dy;
      if (0 <= nx && nx < N && 0 <= ny && ny < M && field[nx][ny]) {
        field[nx][ny] = false;
        dfs(nx, ny);
      }
    }
  }
}
