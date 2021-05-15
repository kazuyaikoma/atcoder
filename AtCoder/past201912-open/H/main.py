#!/usr/bin/env python3
from typing import List


def solve(N: int, C: List[int], Q: int, S: List[List[int]]):
    sell = 0
    # 全種類販売で売った1種類あたりの枚数
    set_cnt = 0
    # セット販売で売った1種類あたりの枚数
    all_cnt = 0
    # セット販売対象のCの最小値を記録
    min_C_set = 1000000001
    # セット販売対象ではない(全ての)Cの最小値を記録
    min_C_all = min(C)

    for i in range(N):
        if i % 2 == 0:
            min_C_set = min(min_C_set, C[i])

    for query in S:
        if query[0] == 1:
            x = query[1] - 1
            a = query[2]

            # 値を考慮する場合はset_cnt(セット販売)も考慮に入れる
            delta = 0
            if x % 2 == 0:
                delta = set_cnt
            if a <= C[x] - all_cnt - delta:
                C[x] -= a
                sell += a
                if x % 2 == 0:
                    min_C_set = min(min_C_set, C[x])
                min_C_all = min(min_C_all, C[x])
        elif query[0] == 2:
            a = query[1]
            if a <= min_C_set - all_cnt - set_cnt:
                set_cnt += a
        elif query[0] == 3:
            a = query[1]
            if a <= min(min_C_set - all_cnt - set_cnt, min_C_all - all_cnt):
                all_cnt += a

    # セット販売した枚数を合算
    for i in range(N):
        if i % 2 == 0:
            sell += set_cnt

    # 全種類販売した枚数を合算
    sell += N * all_cnt

    print(sell)


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
