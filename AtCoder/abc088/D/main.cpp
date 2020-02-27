#include <bits/stdc++.h>
using namespace std;

int W, H;
vector<vector<char>> s;

// 最短経路[y][x]
vector<vector<int>> cost;

// 最短経路を求めるためのキューを使った幅優先探索
int bfs() {
    // 上下左右
    int dx[] = {0, 1, 0, -1}, dy[] = {1, 0, -1, 0};

    // 座標キュー
    queue<pair<int, int>> que;
    // スタート: (0, 0)
    que.push(make_pair(0, 0));
    cost[0][0] = 0;

    while (!que.empty()) {
        // 現在地
        pair<int, int> p = que.front();
        que.pop();

        if(p == make_pair(H-1, W-1)) return cost[p.first][p.second];
        for (int i = 0; i < 4; i++) {
            // 現在地から上下左右、次にどこに行くか
            int n_y = p.first + dy[i], n_x = p.second + dx[i];
            // 壁より外側ならcontinue
            if (n_x < 0 || n_y < 0 || n_x >= W || n_y >= H) continue;
            // 元々黒ならそこは通れないのでcontinue
            if (s[n_y][n_x] == '#') continue;
            // すでに探索済みならcontinue
            if (cost[n_y][n_x] != -1) continue;

            cost[n_y][n_x] = cost[p.first][p.second] + 1;
            que.push(make_pair(n_y, n_x));
        }
    }
    return -1;
}


// Generated by 1.1.6 https://github.com/kyuridenamida/atcoder-tools  (tips: You use the default template now. You can remove this line by using your custom template)
int main(){
    cin >> H >> W;
    s = vector<vector<char>>(H, vector<char>(W));
    int white = 0;
    for (int i = 0; i < H; i++) {
        for (int j = 0; j < W; j++) {
            cin >> s[i][j];
            if (s[i][j] == '.') white++;
        }
    }

    // 全てのマスを-1で初期化
    cost = vector<vector<int>>(H, vector<int>(W, -1));
    int min_cost = bfs();
    if (min_cost == -1) {
        cout << -1 << endl;
    } else {
        // ゴール分の白カウントは入ってるが、スタート分の白カウントがmin_costに入ってないので+1する
        cout << white - (min_cost + 1) << endl;
    }
    return 0;
}