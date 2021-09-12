#!/usr/bin/env python3
from collections import deque

"""
計算量(O)
    BFSの計算量は探索ノード数 x ノードあたり探索先(辺)の数
    ノード数 = N, 辺数 = Mとすると O(NM)
    今回で言えば、
        N = (R-2)(C-2)
        M <= 2N
    なので、O(RC) と見積もる
    今回は 0 <= R, C < 50 のため、およそ2500 x (定数)程度の計算回数と見積もる
"""


def solve(R, C, sy, sx, gy, gx, c):
    Q = deque()
    visited = []
    for _ in range(50):
        visited.append([False] * 50)
    # Qには(y, x, ゴールまでたどり着くステップ数の累計)を入れていく
    Q.append((sy-1, sx-1, 0))

    while 0 < len(Q):
        y, x, step = Q.popleft()
        if visited[y][x]:
            continue

        visited[y][x] = True
        if y == gy-1 and x == gx-1:
            print(step)
            break

        for dy, dx in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
            args = (y + dy, x + dx, step + 1)
            is_path = c[args[0]][args[1]] == '.'
            if 0 <= args[0] < R and 0 <= args[1] < C and is_path:
                Q.append(args)


def main():
    R, C = map(int, input().split())
    sy, sx = map(int, input().split())
    gy, gx = map(int, input().split())
    c = []
    for _ in range(R):
        text = str(input())
        c.append(text)
    solve(R, C, sy, sx, gy, gx, c)


if __name__ == '__main__':
    main()
