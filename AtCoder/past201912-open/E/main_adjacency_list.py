#!/usr/bin/env python3

def print_graph(N, graph):
    for i in range(N):
        ans = ""
        for j in range(N):
            if j in graph[i]:
                ans += 'Y'
            else:
                ans += 'N'
        print(ans)


def solve(N, Q, graph, queries):
    for query in queries:
        a = query[1]-1
        if query[0] == 1:
            b = query[2]-1
            if b not in graph[a]:
                graph[a].append(b)
        elif query[0] == 2:
            for i in range(N):
                if a in graph[i] and i not in graph[a]:
                    graph[a].append(i)
        elif query[0] == 3:
            follows = graph[a]
            ffs = []
            for f in follows:
                ffs += graph[f]
            ffs = list(set(ffs))
            for f in ffs:
                if f not in graph[a]:
                    graph[a].append(f)


def main():
    N, Q = map(int, input().split())
    graph = [[] for _ in range(N)]
    queries = []
    for _ in range(Q):
        query = list(map(int, input().split()))
        queries.append(query)

    solve(N, Q, graph, queries)
    print_graph(N, graph)


if __name__ == '__main__':
    main()
