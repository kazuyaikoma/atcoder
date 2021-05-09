#!/usr/bin/env python3
import sys
from typing import List


def solve(N: int, A: List[int]):
    remainders = [val % 200 for val in A]
    remainders.sort()

    ans = 0
    cums = [0] * 200
    for val in remainders:
        cums[val] += 1

    for cum in cums:
        if 2 <= cum:
            ans += cum * (cum-1) / 2
    print(int(ans))


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
