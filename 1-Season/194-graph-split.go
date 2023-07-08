package main

/*
// Official by Yandex
#include <algorithm>
#include <string>
#include <iostream>
#include <vector>
#include <utility>

struct Solution {
    int n, m;
    std::vector<std::vector<std::pair<int, int>>> graph;
    std::vector<int> colors;

    bool dfs(int v, int color, int bound) {
        colors[v] = color;
        for (const auto &edge : graph[v]) {
            if (edge.second >= bound) {
                continue;
            }
            if (colors[edge.first] == color) {
                return false;
            }
            if (colors[edge.first] == -1) {
                if (!dfs(edge.first, 1 - color, bound)) {
                    return false;
                }
            }
        }
        return true;
    }

    bool check(int bound) {
        for (int i = 0; i < n; i++) {
            if (colors[i] == -1) {
                if (!dfs(i, 0, bound)) {
                    return false;
                }
            }
        }
        return true;
    }

    void run(std::istream &in, std::ostream &out) {
        in >> n >> m;
        graph.assign(n, std::vector<std::pair<int, int>>());
        for (int i = 0; i < m; i++) {
            int u, v, w;
            in >> u >> v >> w;
            u--;
            v--;
            graph[u].emplace_back(v, w);
            graph[v].emplace_back(u, w);
        }
        int l = 0;
        int r = 1000000001;
        while (r - l > 1) {
            int m = (l + r) / 2;
            colors.assign(n, -1);
            if (!check(m)) {
                r = m;
            } else {
                l = m;
            }
        }
        out << l << "\n";
    }
};

int main() {
    std::cin.sync_with_stdio(false);
    std::cin.tie(nullptr);
    Solution().run(std::cin, std::cout);
    return 0;
}
*/
