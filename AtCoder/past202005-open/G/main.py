#!/usr/bin/env python3
import sys
from typing import List
from collections import deque

"""
BFSを使って解いている
解法の公式解説はこちら
https://img.atcoder.jp/past202005-open/editorial.pdf
"""


def solve(N: int, X: int, Y: int, xs: "List[int]", ys: "List[int]"):
    visited = []
    for _ in range(403):
        visited.append([False] * 403)

    walls = []
    for i in range(N):
        walls.append((xs[i], ys[i]))

    Q = deque()
    # スタート時の(x, y, step)をdequeに追加
    Q.append((0, 0, 0))

    while 0 < len(Q):
        x, y, step = Q.popleft()
        if visited[x][y] or ((x, y) in walls):
            continue
        visited[x][y] = True

        if x == X and y == Y:
            print(step)
            return

        next_xys = [
            (x+1, y+1),
            (x, y+1),
            (x-1, y+1),
            (x+1, y),
            (x-1, y),
            (x, y-1),
        ]

        for next_x, next_y in next_xys:
            if -201 <= next_x <= 201 and -201 <= next_y <= 201:
                if not visited[next_x][next_y] and ((next_x, next_y) not in walls):
                    Q.append((next_x, next_y, step + 1))

    print(-1)


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    N = int(next(tokens))  # type: int
    X = int(next(tokens))  # type: int
    Y = int(next(tokens))  # type: int
    x = [int()] * (N)  # type: "List[int]"
    y = [int()] * (N)  # type: "List[int]"
    for i in range(N):
        x[i] = int(next(tokens))
        y[i] = int(next(tokens))
    solve(N, X, Y, x, y)


if __name__ == '__main__':
    main()
