#!/usr/bin/env python3
from typing import List


def print_graph(N: int, graph: List[int]):
    for i in range(N):
        ans = ""
        for j in range(N):
            if j in graph[i]:
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
            graph[a].append(b)
        elif mark == 2:
            followers = get_followers(graph, a)
            graph[a] += followers
        elif mark == 3:
            follow_follows = get_follow_follows(graph, a)
            graph[a] += follow_follows

    print_graph(N, graph)


# aをフォローするユーザー群を取得
def get_followers(graph: List[int], a: int) -> List[int]:
    followers = []
    for i, follows in enumerate(graph):
        if i == a:
            continue
        if a in follows:
            followers.append(i)

    return list(set(followers))


# aがフォローするユーザー群(x)のさらにフォローしてるユーザー群を取得
def get_follow_follows(graph: List[int], a: int) -> List[int]:
    users = []
    for x in graph[a]:
        users += graph[x]
    if a in users:
        users.remove(a)

    return list(set(users))


def main():
    N, Q = map(int, input().split())

    queries = []
    for _ in range(Q):
        query = list(map(int, input().split()))
        queries.append(query)

    # aがbをフォローしていることを示す隣接リスト
    graph = [[] for _ in range(N)]
    solve(N, Q, queries, graph)


if __name__ == '__main__':
    main()
