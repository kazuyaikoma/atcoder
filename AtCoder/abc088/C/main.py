#!/usr/bin/env python3
import sys
from typing import List

YES = "Yes"  # type: str
NO = "No"  # type: str


def solve(c: List[List[int]]):
    first_diff = c[0][1] - c[0][0]
    second_diff = c[0][2] - c[0][1]
    for c_elms in c:
        f_diff = c_elms[1] - c_elms[0]
        s_diff = c_elms[2] - c_elms[1]
        if f_diff != first_diff or s_diff != second_diff:
            print(NO)
            return

    print(YES)


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    c = [[int(next(tokens)) for _ in range(3)] for _ in range(3)]
    solve(c)


if __name__ == '__main__':
    main()
