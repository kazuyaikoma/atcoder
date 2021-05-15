#!/usr/bin/env python3
import sys
from typing import List


def solve(N: int, S: "List[str]", T: "List[int]"):
    dic = {}
    for i in range(N):
        dic[S[i]] = T[i]

    max_val = max(T)
    T.remove(max_val)
    second_val = max(T)

    for k, v in dic.items():
        if v == second_val:
            print(k)
            return


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    N = int(next(tokens))  # type: int
    S = [str()] * (N)  # type: "List[str]"
    T = [int()] * (N)  # type: "List[int]"
    for i in range(N):
        S[i] = next(tokens)
        T[i] = int(next(tokens))
    solve(N, S, T)


if __name__ == '__main__':
    main()
