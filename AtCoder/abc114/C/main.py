#!/usr/bin/env python3
import sys

shichi = ['7', '5', '3']
global ans


def check_shichi(str_num):
    flag = True
    for s in shichi:
        if s not in str_num:
            flag = False

    return flag


def dfs(N, str_num):
    global ans

    num = int(str_num)
    if N < num:
        return

    if 357 <= int(str_num) <= N and check_shichi(str_num):
        ans += 1
    for s in shichi:
        next_num = str_num + s
        dfs(N, next_num)


def solve(N: int):
    global ans
    ans = 0

    dfs(N, '7')
    dfs(N, '5')
    dfs(N, '3')

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
