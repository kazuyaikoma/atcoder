#include <iostream>
#include <queue>
using namespace std;

const int INF = 100000000;
typedef pair<int, int> Pos;

const int N = 10;
const int M = 10;
// true: 道(通れる), false: 石(通れない)。start座標は[0][1], goal座標は[9][8]
char maze[N][M] = {
  {false, true, false, false, false, false, false, false, true, false},
  {true, true, true, true, true, true, false, true, true, false},
  {true, false, true, false, false, true, false, false, true, false},
  {true, false, true, true, true, true, true, true, true, true},
  {false, false, true, false, false, true, false, false, false, false},
  {true, true, true, true, false, true, true, true, true, false},
  {true, false, false, false, false, false, false, false, true, false},
  {true, true, true, true, false, true, true, true, true, true},
  {true, false, false, false, false, true, false, false, false, true},
  {true, true, true, true, false, true, true, true, true, false},
};

int sx = 0, sy = 1;
int gx = 9, gy = 8;

int d[N][M];
// 移動のための上下左右ベクトル(インデックス指定は同じ添字で取るようにする)
int dx[4] = {1, 0, -1, 0}, dy[4] = {0, 1, 0, -1};

// interface
void bfs(void);
void dfs(int x, int y);

int main() {
  for (int i = 0; i < N; i++)
    for (int j = 0; j < M; j++) d[i][j] = INF;
  d[sx][sy] = 0;

  bfs();
  // dfs(sx, sy);

  printf("%d\n", d[gx][gy]);
  return 0;
}

// 幅優先探索(bfs)
void bfs() {
  queue<Pos> que;

  que.push(Pos(sx, sy));
  d[sx][sy] = 0;

  while (que.size()) {
    Pos p = que.front(); que.pop();
    if (p.first == gx && p.second == gy) break;

    for (int i = 0; i < 4; i++) {
      int nx = p.first + dx[i], ny = p.second + dy[i];

      if (0 <= nx && nx < N && 0 <= ny && ny < M && maze[nx][ny] && d[nx][ny] == INF) {
        que.push(Pos(nx, ny));
        d[nx][ny] = d[p.first][p.second] + 1;
      }
    }
  }
}

// 深さ優先探索(dfs)
void dfs(int x, int y) {
  for (int i = 0; i < 4; i++) {
    int nx = x + dx[i], ny = y + dy[i];
    if (0 <= nx && nx < N && 0 <= ny && ny < M && maze[nx][ny] && d[nx][ny] == INF) {
      d[nx][ny] = d[x][y] + 1;
      dfs(nx, ny);
    }
  }
}