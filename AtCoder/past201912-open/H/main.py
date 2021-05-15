#!/usr/bin/env python3
from typing import List


def solve(N: int, C: List[int], Q: int, S: List[List[int]]):
    origin = sum(C)

    for i in range(Q):
        query = S[i]
        if query[0] == 1:
            sell = query[2]
            sell_card_idx = query[1] - 1
            if sell <= C[sell_card_idx]:
                C[sell_card_idx] -= sell
        elif query[0] == 2:
            sell = query[1]
            ok = True
            for j in range(0, len(C), 2):
                if C[j] - sell < 0:
                    ok = False
            if ok:
                for j in range(0, len(C), 2):
                    C[j] -= sell
        elif query[0] == 3:
            sell = query[1]
            ok = True
            for j in range(len(C)):
                if C[j] - sell < 0:
                    ok = False
            if ok:
                for j in range(len(C)):
                    C[j] -= sell

    print(origin - sum(C))


def main():
    N = int(input())
    C = list(map(int, input().split()))
    Q = int(input())

    S = []
    for _ in range(Q):
        query = list(map(int, input().split()))
        S.append(query)
    solve(N, C, Q, S)


if __name__ == '__main__':
    main()
