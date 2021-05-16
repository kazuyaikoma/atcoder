#!/usr/bin/env python3
import sys
from typing import List

"""
ランタイムをPyhton3にて実行した場合2050ms程度でギリギリ実行時間の制約をオーバーするため、
PyPy3を選択するとACする(300~400ms程度)
"""


def solve(N: int, K: int, h: "List[int]"):
    dp = [0] * N
    dp[1] = abs(h[1] - h[0])

    for i in range(2, N):
        hs = []
        min_j = i - K
        if min_j < 0:
            min_j = 0
        for j in range(min_j, i):
            h_diff = abs(h[i] - h[j])
            hs.append(dp[j] + h_diff)
        dp[i] = min(hs)

    print(dp[N-1])


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    N = int(next(tokens))  # type: int
    K = int(next(tokens))  # type: int
    h = [int(next(tokens)) for _ in range(N)]  # type: "List[int]"
    solve(N, K, h)


if __name__ == '__main__':
    main()
