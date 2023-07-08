package main

/*
import sys
import heapq

input = sys.stdin.readline
n, m, q = map(int, input().split())
db = [0] * n  # битовая маска включенных машин
da = [m] * n  # число включённыых машин
dr = [0] * n  # число перезапусков
dw = [0] * n  # метрика r*a

min_heap = [(0, i) for i in range(n)]
max_heap = [(0, i) for i in range(n)]

for _ in range(q):
    cmd, *a, = input().split()
    *a, = map(int, a)
    if cmd == "RESET":
        i = a[0] - 1
        db[i] = 0
        da[i] = m
        dr[i] += 1
        dw[i] = dr[i] * da[i]

        heapq.heappush(min_heap, (dw[i], i))
        heapq.heappush(max_heap, (-dw[i], i))
    elif cmd == "DISABLE":
        i, j = a
        i -= 1
        j -= 1
        t = 1 << j

        if db[i] & t:
            continue
        db[i] |= t
        da[i] -= 1
        dw[i] -= dr[i]

        heapq.heappush(min_heap, (dw[i], i))
        heapq.heappush(max_heap, (-dw[i], i))
    elif cmd == "GETMAX":
        while -max_heap[0][0] != dw[max_heap[0][1]]:
            heapq.heappop(max_heap)
        print(max_heap[0][1] + 1)
    elif cmd == "GETMIN":
        while min_heap[0][0] != dw[min_heap[0][1]]:
            heapq.heappop(min_heap)
        print(min_heap[0][1] + 1)
*/
