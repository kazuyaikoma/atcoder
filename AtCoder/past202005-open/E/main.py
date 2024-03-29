#!/usr/bin/env python3

# 隣接行列を使った場合
# def solve(N, M, Q, graph, colors, queries):
#     for query in queries:
#         x = query[1]
#         print(colors[x])
#         if query[0] == 1:
#             # スプリンクラー起動
#             for i in range(N):
#                 if graph[x][i]:
#                     colors[i] = colors[x]
#         elif query[0] == 2:
#             colors[x] = query[2]

# 隣接リストを使った場合
def solve(N, M, Q, graph, colors, queries):
    for query in queries:
        x = query[1]
        print(colors[x])
        if query[0] == 1:
            # スプリンクラー起動
            for target in graph[x]:
                colors[target] = colors[x]
        elif query[0] == 2:
            colors[x] = query[2]


def main():
    N, M, Q = map(int, input().split())

    # 隣接行列
    # graph = []
    # for _ in range(N):
    #     row = [False for _ in range(N)]
    #     graph.append(row)

    # for _ in range(M):
    #     u, v = map(int, input().split())
    #     # 頂点idxは扱いやすいよう-1している
    #     graph[u-1][v-1] = True
    #     graph[v-1][u-1] = True

    # 隣接リスト
    graph = [[] for _ in range(N)]
    for _ in range(M):
        u, v = map(int, input().split())
        # 頂点idxは扱いやすいよう-1している
        graph[u-1].append(v-1)
        graph[v-1].append(u-1)

    colors = list(map(int, input().split()))

    queries = []
    for _ in range(Q):
        query = list(map(int, input().split()))
        # 頂点idxは扱いやすいよう-1している
        query[1] -= 1
        queries.append(query)

    solve(N, M, Q, graph, colors, queries)


if __name__ == '__main__':
    main()
