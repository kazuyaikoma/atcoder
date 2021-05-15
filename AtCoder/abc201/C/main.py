#!/usr/bin/env python3
import sys


def solve(S: str):
    dic = {}
    for k, v in enumerate(S):
        dic[k] = v

    ans = 0
    for num in range(10000):
        s = str(num).rjust(4, '0')
        flag = True
        for k, v in dic.items():
            if dic[k] == 'o':
                if str(k) not in s:
                    flag = False
            elif dic[k] == 'x':
                if str(k) in s:
                    flag = False
        if flag:
            ans += 1

    print(ans)


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    S = next(tokens)  # type: str
    solve(S)


if __name__ == '__main__':
    main()
