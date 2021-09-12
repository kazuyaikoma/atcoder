#!/usr/bin/env python3
from typing import List
import math


def solve(cards: List[int], Q: int, queries: List[List[int]]):
    sell = 0

    # セット販売で売った1種類あたりの枚数
    set_sell = 0
    # 全種類販売で売った1種類あたりの枚数
    all_sell = 0

    # セット販売対象のcardsの最小値を記録
    min_cards_set = max(cards)
    # セット販売対象以外のcardsの最小値を記録
    min_cards_others = max(cards)

    # セット販売の最小値を算出
    for i in range(len(cards)):
        if is_set_card(i):
            min_cards_set = min(min_cards_set, cards[i])

    for query in queries:
        if query[0] == 1:
            x = query[1] - 1
            a = query[2]

            # 値を考慮する場合はset_sell(セット販売)も考慮に入れる
            delta = 0
            if is_set_card(x):
                delta = set_sell

            if a <= cards[x] - all_sell - delta:
                cards[x] -= a
                sell += a

                if is_set_card(x):
                    min_cards_set = min(min_cards_set, cards[x])
                min_cards_others = min(min_cards_others, cards[x])
        elif query[0] == 2:
            a = query[1]
            if a <= min_cards_set - all_sell - set_sell:
                set_sell += a
        elif query[0] == 3:
            a = query[1]
            if a <= min(
                    min_cards_set - all_sell - set_sell,
                    min_cards_others - all_sell
            ):
                all_sell += a

    # セット販売した枚数を合算
    sell += set_sell * math.ceil(len(cards)/2)

    # 全種類販売した枚数を合算
    sell += all_sell * len(cards)

    print(sell)


def is_set_card(x: int) -> bool:
    return x % 2 == 0


def main():
    _ = int(input())
    cards = list(map(int, input().split()))
    Q = int(input())

    queries = []
    for _ in range(Q):
        query = list(map(int, input().split()))
        queries.append(query)
    solve(cards, Q, queries)


if __name__ == '__main__':
    main()
