#!/usr/bin/env python3
from typing import List


def solve(N: int, C: List[int], Q: int, S: List[List[int]]):
    sell = 0

    # 全種類販売で売った1種類あたりの枚数
    z = 0
    # セット販売で売った1種類あたりの枚数
    s = 0

    # セット販売対象のCの最小値を記録
    min_s_C = 1000000000000
    # セット販売対象ではないCの最小値を記録
    min_z_C = 1000000000000

    for i in range(N):
        if i % 2 == 0:
            min_s_C = min(min_s_C, C[i])
        else:
            min_z_C = min(min_z_C, C[i])

    for i in range(Q):
        query = S[i]
        if query[0] == 1:
            x = query[1] - 1
            a = query[2]

            if x % 2 == 0:
                card_x = C[x] - z - s
            else:
                card_x = C[x] - z

            if a <= card_x:
                C[x] -= a
                sell += a

                if x % 2 == 0:
                    min_s_C = min(min_s_C, C[x])
                else:
                    min_z_C = min(min_z_C, C[x])
        elif query[0] == 2:
            a = query[1]
            if a <= min_s_C - s - z:
                s += a
        elif query[0] == 3:
            a = query[1]
            if a <= min(min_s_C - s - z, min_z_C - z):
                z += a

    # セット販売した枚数を合算
    # sell += (N // 2) * s
    for i in range(N):
        if i % 2 == 0:
            sell += s
    # 全種類販売した枚数を合算
    sell += N * z

    print(int(sell))


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
