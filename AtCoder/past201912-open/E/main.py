#!/usr/bin/env python3
from typing import List


def print_graph(N: int, graph: List[int]):
    for i in range(N):
        ans = ""
        for j in range(N):
            if graph[i][j]:
                ans += 'Y'
            else:
                ans += 'N'
        print(ans)


def solve(N: int, Q: int, queries: List[int], graph: List[int]):
    for query in queries:
        mark = query[0]
        a = query[1] - 1
        if mark == 1:
            b = query[2] - 1
            graph[a][b] = True
        elif mark == 2:
            for i in range(N):
                if graph[i][a]:
                    graph[a][i] = True
        elif mark == 3:
            to_follow = []

            for i in range(N):
                if graph[a][i]:
                    for j in range(N):
                        if graph[i][j] and j != a:
                            to_follow.append(j)
            for f in to_follow:
                graph[a][f] = True

    print_graph(N, graph)


def main():
    N, Q = map(int, input().split())
    graph = []
    for _ in range(N):
        row = [False for _ in range(N)]
        graph.append(row)

    queries = []
    for _ in range(Q):
        query = list(map(int, input().split()))
        queries.append(query)

    solve(N, Q, queries, graph)


if __name__ == '__main__':
    main()
