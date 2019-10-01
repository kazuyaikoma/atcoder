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
        dfs(x, y);
        count++;
      }
    }
  }
  printf("%d\n", count);

  return 0;
}

void dfs(int x, int y) {
  // 現在地をfalseにする
  field[x][y] = false;

  for (int dx = -1; dx <= 1; dx++) {
    for (int dy = -1; dy <= 1; dy++) {
      int nx = x + dx, ny = y + dy;
      bool isLake = field[nx][ny];
      if (0 <= nx && nx < N && 0 <= ny && ny < M && isLake) {
        dfs(nx, ny);
      }
    }
  }
  return;
}

