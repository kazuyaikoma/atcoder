#!/usr/bin/env python3
import sys


def is_match(S: str, T: str):
    if len(S) == 0 or len(T) == 0:
        return False

    for i in range(len(S)-len(T)+1):
        if T[0] == '.' or T[0] == S[i]:
            if len(T) == 1:
                return True
            if T[1] == '.' or T[1] == S[i+1]:
                if len(T) == 2:
                    return True
                if T[2] == '.' or T[2] == S[i+2]:
                    return True
    return False


def solve(S: str):
    ans = 0
    strs = []
    alphabets = "abcdefghijklmnopqrstuvwxyz."
    for a1 in alphabets:
        strs.append(a1)
        for a2 in alphabets:
            strs.append(a1 + a2)
            for a3 in alphabets:
                strs.append(a1 + a2 + a3)
    for s in strs:
        if is_match(S, s):
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
