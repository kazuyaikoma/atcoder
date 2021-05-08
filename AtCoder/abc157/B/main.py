#!/usr/bin/env python3
import sys
from typing import List

YES = "Yes"  # type: str
NO = "No"  # type: str


def solve(A: List[List[int]], N: int, b: List[int]):
    row_bingos = [
        [A[0][0], A[0][1], A[0][2]],
        [A[1][0], A[1][1], A[1][2]],
        [A[2][0], A[2][1], A[2][2]],
    ]
    column_bingos = [
        [A[0][0], A[1][0], A[2][0]],
        [A[0][1], A[1][1], A[2][1]],
        [A[0][2], A[1][2], A[2][2]],
    ]
    slash_bingos = [
        [A[0][0], A[1][1], A[2][2]],
        [A[0][2], A[1][1], A[2][0]],
    ]

    for bingo in row_bingos + column_bingos + slash_bingos:
        is_bingo = True
        for elm in bingo:
            if elm not in b:
                is_bingo = False
                break
        if is_bingo:
            print(YES)
            return

    print(NO)


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    A = [[int(next(tokens)) for _ in range(3)] for _ in range(3)]
    N = int(next(tokens))
    b = [int(next(tokens)) for _ in range(N)]
    solve(A, N, b)


if __name__ == '__main__':
    main()
