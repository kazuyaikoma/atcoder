#!/usr/bin/env python3
from typing import List


def solve(N: int, S: List[List[str]]):
    for i in range(N-2, -1, -1):
        for j in range(1, 2*N-1):
            if S[i][j] == '#':
                c1 = S[i+1][j-1] == 'X'
                c2 = S[i+1][j] == 'X'
                c3 = S[i+1][j+1] == 'X'
                if c1 or c2 or c3:
                    S[i][j] = 'X'

    for row in S:
        print(''.join(row))


def main():
    N = int(input())
    S = []
    for _ in range(N):
        row = list(str(input()))
        S.append(row)
    solve(N, S)


if __name__ == '__main__':
    main()
