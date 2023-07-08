package main

import "fmt"

func main() {
	var V, E uint
	fmt.Scan(&V, &E)

	matrix := make([][]int, V)
	for i := range matrix {
		matrix[i] = make([]int, V)
	}

	var a, b uint
	for i := uint(0); i < E; i++ {
		fmt.Scan(&a, &b)
		matrix[a-1][b-1] = 0
		matrix[b-1][a-1] = 0
	}

	partition := make([]int, V)
	for i := range partition {
		partition[i] = -1
	}

	cnt := 0
	divisible := true

	for curr := uint(0); curr < V; curr++ {
		if partition[curr] != -1 {
			continue
		}

		partition[curr] = 1
		cnt++

		// queue
		var nexts []uint
		nexts = append(nexts, curr)

		for len(nexts) > 0 && divisible {
			curr = nexts[0]
			// pop
			nexts = nexts[1:]
			for next := uint(0); next < V; next++ {
				if matrix[curr][next] == 1 && next != curr {
					if partition[next] == partition[curr] {
						divisible = false
						break
					} else if partition[next] == -1 {
						switch partition[curr] {
						case 0:
							partition[next] = 1
						case 1:
							partition[next] = 0
						}

						if partition[next] == 1 {
							cnt++
						}
						nexts = append(nexts, next)
					}
				}
			}
		}
		if !divisible {
			break
		}
	}
	if divisible {
		if uint(cnt) == V {
			fmt.Println("1")
			fmt.Println("1")
			for i := uint(1); i < V; i++ {
				fmt.Printf("%d ", i+1)
			}
		} else {
			fmt.Println(cnt)
			for i := uint(0); i < V; i++ {
				if partition[i] == 1 {
					fmt.Printf("%d ", i+1)
				}
			}
			fmt.Println()

			for i := uint(0); i < V; i++ {
				if partition[i] == 0 {
					fmt.Printf("%d ", i+1)
				}
			}
		}
	} else {
		fmt.Println(-1)
	}
}

/*
#include <vector>
#include <iostream>
#include <queue>

int bipartition()
{
    unsigned V, E;
    std::cin >> V >> E;
    std::vector<std::vector<int>> matrix(V, std::vector<int>(V, 1));
    for (unsigned i = 0; i < E; ++i) {
        unsigned a, b;
        std::cin >> a >> b;
        matrix[a-1][b-1] = matrix[b-1][a-1] = 0;
    }

    std::vector<int> partition(V, -1);
    unsigned counter = 0;
    bool divisible = true;
    for (unsigned current = 0; current < V; ++current) {
        if (partition[current] != -1)
            continue;
        partition[current] = 1;
        ++counter;

        std::queue<unsigned> nexts;
        nexts.push(current);
        while (!nexts.empty() && divisible) {
            current = nexts.front();
            nexts.pop();
            for (unsigned next = 0; next < V; ++next) {
                if (matrix[current][next] == 1 && next != current)
                    if (partition[next] == partition[current]) {
                        divisible = false;
                        break;
                    }
                    else if (partition[next] == -1) {
                        partition[next] = !partition[current];
                        if (partition[next] == 1)
                            ++counter;
                        nexts.push(next);
                    }
            }
        }
        if (!divisible)
            break;
    }
    if (divisible) {
        if (counter == V) {
            std::cout << 1 << std::endl;
            std::cout << 1 << std::endl;
            for (unsigned i = 1; i < V; ++i) {
                    std::cout << i + 1 << " ";
            }
        }
        else {
            std::cout << counter;
            std::cout << std::endl;
            for (unsigned i = 0; i < V; ++i) {
                if (partition[i])
                    std::cout << i + 1 << " ";
            }
            std::cout << std::endl;
            for (unsigned i = 0; i < V; ++i) {
                if (!partition[i])
                    std::cout << i + 1 << " ";
            }
        }
    }
    else
        std::cout << -1;

    return 0;
}

int main()
{
    bipartition();
    return 0;
}
*/
