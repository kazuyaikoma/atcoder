#!/usr/bin/env python3
import sys
from typing import List


def solve(N: int, A: List[int]):
    # A.sort()
    # # 累積和の配列
    # cums = []
    # for i, num in enumerate(A):
    #     if i == len(A)-1:
    #         continue

    #     cums.append(A[i+1] - A[i])

    # cums.reverse()
    pass


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    N = int(next(tokens))  # type: int
    A = [int(next(tokens)) for _ in range(N)]  # type: "List[int]"
    solve(N, A)


if __name__ == '__main__':
    main()
