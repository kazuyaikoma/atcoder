#!/usr/bin/env python3
import sys
"""
計算量(O)
    BFSと同じく、ノード総数(N) x 辺の数(M)としてO(N+M)
"""

sys.setrecursionlimit(1000000)
YES = "Yes"  # type: str
NO = "No"  # type: str


def dfs(H, W, visited, maze, y, x):
    if visited[y][x]:
        return False
    visited[y][x] = True

    outputs = []
    for dy, dx in [(0, 1), (0, -1), (1, 0), (-1, 0)]:
        next_y, next_x = y + dy, x + dx
        if 0 <= next_y < H and 0 <= next_x < W:
            if maze[next_y][next_x] == '.':
                dfs_output = dfs(H, W, visited, maze, next_y, next_x)
                outputs.append(dfs_output)
            elif maze[next_y][next_x] == 'g':
                return True

    return any(outputs)


def solve(H, W, start, maze):
    visited = []
    for _ in range(H):
        visited.append([False] * W)

    takahashi_flag = dfs(H, W, visited, maze, start[0], start[1])
    if takahashi_flag:
        print(YES)
    else:
        print(NO)


def main():
    H, W = map(int, input().split())
    maze = []
    for _ in range(H):
        maze.append(str(input()))

    start = (0, 0)
    for i in range(H):
        for j in range(W):
            if maze[i][j] == 's':
                start = (i, j)

    solve(H, W, start, maze)


if __name__ == '__main__':
    main()
