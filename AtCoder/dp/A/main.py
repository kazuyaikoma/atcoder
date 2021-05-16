#!/usr/bin/env python3
import sys
from typing import List


def solve(N: int, h: "List[int]"):
    dp = [0] * N
    dp[0] = 0
    dp[1] = abs(h[1] - h[0])
    for i in range(2, len(h)):
        h1 = abs(h[i] - h[i-1])
        h2 = abs(h[i] - h[i-2])
        dp[i] = min(dp[i-1] + h1, dp[i-2] + h2)
    print(dp[N-1])


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    N = int(next(tokens))  # type: int
    h = [int(next(tokens)) for _ in range(N)]
    solve(N, h)


if __name__ == '__main__':
    main()
