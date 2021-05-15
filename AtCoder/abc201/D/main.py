#!/usr/bin/env python3

def calc(H, W, A, visited, dp, i, j):
    if visited[i][j]:
        return dp[i][j]
    visited[i][j] = True

    takahashi_turn = (i+j) % 2 == 0
    if i == H - 1 and j == W - 1:
        dp[i][j] = 0
        return 0

    if takahashi_turn:
        # Takahashi
        dp[i][j] = -(2000 * 2000) - 1
        if i < H - 1:
            if A[i+1][j] == '+':
                val = 1
            else:
                val = -1
            calc_val = calc(H, W, A, visited, dp, i+1, j) + val
            dp[i][j] = max(dp[i][j], calc_val)
        if j < W - 1:
            if A[i][j+1] == '+':
                val = 1
            else:
                val = -1
            calc_val = calc(H, W, A, visited, dp, i, j+1) + val
            dp[i][j] = max(dp[i][j], calc_val)
    else:
        # Aoki
        dp[i][j] = (2000 * 2000) + 1
        if i < H - 1:
            if A[i+1][j] == '+':
                val = 1
            else:
                val = -1
            calc_val = calc(H, W, A, visited, dp, i+1, j) + val
            dp[i][j] = min(dp[i][j], calc_val)
        if j < W - 1:
            if A[i][j+1] == '+':
                val = 1
            else:
                val = -1
            calc_val = calc(H, W, A, visited, dp, i, j+1) + val
            dp[i][j] = min(dp[i][j], calc_val)

    return dp[i][j]


def solve(H, W, A):
    visited = []
    dp = []
    for _ in range(H):
        visited.append([False] * W)
        dp.append([0] * W)
    val = calc(H, W, A, visited, dp, 0, 0)
    if val < 0:
        print("Aoki")
    elif val == 0:
        print("Draw")
    else:
        print("Takahashi")


def main():
    H, W = map(int, input().split())
    A = []
    for i in range(H):
        A.append(str(input()))
    solve(H, W, A)


if __name__ == '__main__':
    main()
