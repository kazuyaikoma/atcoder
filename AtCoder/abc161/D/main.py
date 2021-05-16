#!/usr/bin/env python3
import sys
from collections import deque


def solve(K: int):
    cnt = 0

    Q = deque()
    for num in range(1, 10):
        num_str = str(num)
        Q.append(num_str)

    while cnt < K:
        num_str = Q.popleft()
        cnt += 1
        if cnt == K:
            print(num_str)
            return

        num_last = int(num_str[-1])
        next_ns = [num_last-1, num_last, num_last+1]
        if num_last == 0:
            next_ns.remove(-1)
        if num_last == 9:
            next_ns.remove(10)
        for next_n in next_ns:
            Q.append(num_str + str(next_n))


def main():
    def iterate_tokens():
        for line in sys.stdin:
            for word in line.split():
                yield word
    tokens = iterate_tokens()
    K = int(next(tokens))  # type: int
    solve(K)


if __name__ == '__main__':
    main()
