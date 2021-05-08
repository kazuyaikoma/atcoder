#!/usr/bin/env python3
import sys


def solve(N: int, K: int):
    for _ in range(K):
        if N % 200 == 0:
            N = int(N / 200)
        else:
            num = str(N) + '200'
            N = int(num)
    print(N)


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    N = int(next(tokens))  # type: int
    K = int(next(tokens))  # type: int
    solve(N, K)


if __name__ == '__main__':
    main()
