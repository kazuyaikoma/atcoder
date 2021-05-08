#!/usr/bin/env python3
import sys


# 全探索の場合
def brute_force_solve(N: int):
    num = 1
    cnt = 0

    while True:
        if len(set(list(str(num)))) == 1:
            cnt += 1
        if cnt == N:
            print(num)
            return

        num += 1


def solve(N: int):
    digit = N // 9
    num = N % 9

    arr = []
    for _ in range(digit+1):
        arr.append(str(num))
    ans = int(''.join(arr))
    print(ans)


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    N = int(next(tokens))  # type: int
    solve(N)


if __name__ == '__main__':
    main()
