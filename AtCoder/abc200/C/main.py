#!/usr/bin/env python3
import sys
from typing import List


def solve(N: int, A: List[int]):
    for i, num in enumerate(A):
        A[i] = num % 200

    A.sort()
    cums = [0] * 200
    for a in A:
        cums[a] += 1

    sum = 0
    for cum in cums:
        if cum == 0 or cum == 1:
            continue
        else:
            sum += cum * (cum-1) / 2

    print(int(sum))


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
