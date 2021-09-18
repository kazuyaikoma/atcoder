#!/usr/bin/env python3
from collections import deque

"""
BFSを使って解いている
解法の公式解説はこちら
https://img.atcoder.jp/past202005-open/editorial.pdf
"""


def solve(N, X, Y, walls):
    Q = deque()
    # (start_x, start_y, step)
    Q.append((0, 0, 0))

    visited = []
    for _ in range(403):
        visited.append([False] * 403)

    while 0 < len(Q):
        x, y, step = Q.popleft()
        if x == X and y == Y:
            print(step)
            return

        if visited[x][y] or (x, y) in walls:
            continue
        visited[x][y] = True

        next_xys = [
            (x+1, y+1),
            (x, y+1),
            (x-1, y+1),
            (x+1, y),
            (x-1, y),
            (x, y-1),
        ]
        for nx, ny in next_xys:
            if -201 <= nx <= 201 and -201 <= ny <= 201:
                if not visited[nx][ny] and (nx, ny) not in walls:
                    Q.append((nx, ny, step+1))
    print(-1)


def main():
    N, X, Y = map(int, input().split())
    walls = []
    for _ in range(N):
        walls.append(tuple(map(int, input().split())))

    solve(N, X, Y, walls)


if __name__ == '__main__':
    main()
