#!/usr/bin/env python3
from typing import List


def solve(cards: List[int], Q: int, S: List[List[int]]):
    sell = 0

    # 全種類販売で売った1種類あたりの枚数
    all_cnt = 0
    # セット販売で売った1種類あたりの枚数
    set_cnt = 0
    # セット販売対象のcardsの最小値を記録
    min_cards_set = 1000000001
    # セット販売対象ではない(全ての)cardsの最小値を記録
    min_cards_all = min(cards)

    for i in range(len(cards)):
        if i % 2 == 0:
            min_cards_set = min(min_cards_set, cards[i])

    for query in S:
        if query[0] == 1:
            x = query[1] - 1
            a = query[2]

            # 値を考慮する場合はset_cnt(セット販売)も考慮に入れる
            delta = 0
            if x % 2 == 0:
                delta = set_cnt
            if a <= cards[x] - all_cnt - delta:
                cards[x] -= a
                sell += a
                if x % 2 == 0:
                    min_cards_set = min(min_cards_set, cards[x])
                min_cards_all = min(min_cards_all, cards[x])
        elif query[0] == 2:
            a = query[1]
            if a <= min_cards_set - all_cnt - set_cnt:
                set_cnt += a
        elif query[0] == 3:
            a = query[1]
            if a <= min(
                    min_cards_set - all_cnt - set_cnt,
                    min_cards_all - all_cnt
            ):
                all_cnt += a

    # セット販売した枚数を合算
    for i in range(len(cards)):
        if i % 2 == 0:
            sell += set_cnt

    # 全種類販売した枚数を合算
    sell += len(cards) * all_cnt

    print(sell)


def main():
    _ = int(input())
    cards = list(map(int, input().split()))
    Q = int(input())

    S = []
    for _ in range(Q):
        query = list(map(int, input().split()))
        S.append(query)
    solve(cards, Q, S)


if __name__ == '__main__':
    main()
