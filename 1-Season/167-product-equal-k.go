package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Point struct {
	index int
	value int
}

type ByValueDesc []Point

func (a ByValueDesc) Len() int           { return len(a) }
func (a ByValueDesc) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a ByValueDesc) Less(i, j int) bool { return a[i].value > a[j].value }

func main() {
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')
	input = strings.TrimSpace(input)
	inputArr := strings.Split(input, " ")
	n, _ := strconv.Atoi(inputArr[0]) // size of array
	m, _ := strconv.Atoi(inputArr[1]) // multiplication
	k, _ := strconv.Atoi(inputArr[2]) // count of elements

	numbersStr, _ := reader.ReadString('\n')
	numbersStr = strings.TrimSpace(numbersStr)
	numbersArr := strings.Split(numbersStr, " ")
	array := make([]Point, n)
	arrayCorrect := make([]Point, 0)
	for i := 0; i < n; i++ {
		cur, _ := strconv.Atoi(numbersArr[i])
		p := Point{i, cur}
		array[i] = p
		if cur == 0 && m != 0 {
			continue
		}
		if cur == 0 {
			if m == 0 {
				arrayCorrect = append(arrayCorrect, p)
			}
		} else if m%cur == 0 {
			arrayCorrect = append(arrayCorrect, p)
		}
	}

	res := solve(array, arrayCorrect, k, m)
	fmt.Println(res)
}

func solve(array, arrayCorrect []Point, k, m int) string {
	sort.Sort(ByValueDesc(arrayCorrect))
	var result strings.Builder
	if m == 0 {
		for i := 0; i < len(array); i++ {
			if array[i].value == 0 {
				result.WriteString(strconv.Itoa(i + 1))
				result.WriteString(" ")
				k--
				for j := 0; j < len(array); j++ {
					if i == j {
						continue
					}
					if k == 0 {
						break
					}
					result.WriteString(strconv.Itoa(j + 1))
					result.WriteString(" ")
					k--
				}
				return strings.TrimSpace(result.String())
			}
		}
	}
	if m == 1 {
		for i := 0; i < len(array); i++ {
			if array[i].value == 1 {
				result.WriteString(strconv.Itoa(i + 1))
				result.WriteString(" ")
				k--
				if k <= 0 {
					break
				}
			}
		}
		return strings.TrimSpace(result.String())
	}

	var indexStack []int
	var st []Point
	lastBad := 0
	end := len(arrayCorrect) - 1
	rest := m
	indexOfNextValue := make(map[int]int, end)
	indexOfOne := -1
	amountOfOne := 0
	prev := -1
	for i := 0; i <= end; i++ {
		p := arrayCorrect[i]
		cur := p.value
		if cur != prev && prev != -1 {
			indexOfNextValue[prev] = i
		}
		if cur == 1 && indexOfOne == -1 {
			indexOfOne = i
			amountOfOne = end - i + 1
		}
		prev = cur
	}

	for i := 0; i <= end; i++ {
		p := arrayCorrect[i]
		cur := p.value
		if rest == 1 {
			if indexOfOne == -1 {
				rest = rest * st[len(st)-1].value
				lastBad = st[len(st)-1].value
				i = indexStack[len(indexStack)-1]
				continue
			} else {
				if i < (indexOfOne - 1) {
					need := k - len(st)
					if need <= amountOfOne {
						for j := indexOfOne; j < indexOfOne+need; j++ {
							pp := arrayCorrect[j]
							st = append(st, pp)
						}
						break
					} else {
						rest = rest * st[len(st)-1].value
						lastBad = st[len(st)-1].value
						i = indexStack[len(indexStack)-1]
					}
					continue
				}
			}
		}
		if lastBad == cur {
			in, ok := indexOfNextValue[cur]
			if !ok {
				rest = rest * st[len(st)-1].value
				lastBad = st[len(st)-1].value
				i = indexStack[len(indexStack)-1]
			} else {
				i = in - 1
			}
			continue
		}
		if cur > rest || (rest%cur != 0) {
			lastBad = cur
			if i == end {
				for i >= (end - 1) {
					lastBad = st[len(st)-1].value
					rest *= lastBad
					st = st[:len(st)-1]
					i = indexStack[len(indexStack)-1]
					indexStack = indexStack[:len(indexStack)-1]
				}
			}
			continue
		}
		rest = rest / cur
		st = append(st, p)
		indexStack = append(indexStack, i)
		if len(st) == k && rest == 1 {
			break
		}
		if i == end {
			for i >= (end - 1) {
				lastBad = st[len(st)-1].value
				rest *= lastBad
				st = st[:len(st)-1]
				i = indexStack[len(indexStack)-1]
				indexStack = indexStack[:len(indexStack)-1]
			}
		}
	}
	for _, point := range st {
		result.WriteString(strconv.Itoa(point.index + 1))
		result.WriteString(" ")
	}
	return strings.TrimSpace(result.String())
}

/* Java
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.time.Duration;
import java.time.LocalDateTime;
import java.util.*;

public class Main {

    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        String[] input = reader.readLine().split(" ");
        int n = Integer.parseInt(input[0]); // size of array
        int m = Integer.parseInt(input[1]); // multiplication
        int k = Integer.parseInt(input[2]); // count of elements

        String[] numbersStr = reader.readLine().split(" ");
        List<Point> array = new ArrayList<>(n);
        List<Point> arrayCorrect = new ArrayList<>(n);
        for (int i = 0; i < n; i++) {
            Integer cur = Integer.parseInt(numbersStr[i]);
            Point p = new Point(i, cur);
            array.add(p);
            if (cur == 0 && m != 0) {
                continue;
            }
            if (cur == 0) {
                if (m == 0) {
                    arrayCorrect.add(p);
                }
            } else if (m % cur == 0) {
                arrayCorrect.add(p);
            }
        }
        reader.close();

        String res = solve(array, arrayCorrect, k, m);
        System.out.println(res);
    }

    private static String solve(List<Point> array, List<Point> arrayCorrect, int k, int m) {
        Collections.sort(arrayCorrect, Collections.reverseOrder());
        StringBuilder result = new StringBuilder(k);
        if (m == 0) {
            for (int i = 0; i < array.size(); i++) {
                if (array.get(i).value == 0) {
                    result.append(i + 1);
                    result.append(" ");
                    k--;
                    for (int j = 0; j < array.size(); j++) {
                        if (i == j) continue;
                        if (k == 0) break;
                        result.append(j + 1);
                        result.append(" ");
                        k--;
                    }
                    return result.toString().trim();
                }
            }
        }
        if (m == 1) {
            for (int i = 0; i < array.size(); i++) {
                if (array.get(i).value == 1) {
                    result.append(i + 1);
                    result.append(" ");
                    k--;
                    if (k <= 0) break;
                }
            }
            return result.toString().trim();
        }

        Deque<Integer> indexStack = new ArrayDeque<>();
        Deque<Point> st = new ArrayDeque<>();
        int lastBad = 0;
        int end = arrayCorrect.size() - 1;
        int rest = m;
        Map<Integer, Integer> indexOfNextValue = new HashMap<>(end);
        Integer indexOfOne = -1;
        int amountOfOne = 0;
        Integer prev = -1;
        for (int i = 0; i <= end; ++i) {
            Point p = arrayCorrect.get(i);
            Integer cur = p.value;
            if (cur != prev && prev != -1) {
                indexOfNextValue.put(prev, i);
            }
            if (cur == 1 && indexOfOne == -1) {
                indexOfOne = i;
                amountOfOne = end - i + 1;
            }
            prev = cur;
        }

        for (int i = 0; i <= end; ++i) {
            Point p = arrayCorrect.get(i);
            Integer cur = p.value;
            if (rest == 1) {
                if (indexOfOne == -1) {
                    rest = rest * st.peekLast().value;
                    lastBad = st.pollLast().value;
                    i = indexStack.pollLast();
                    continue;
                } else {
                    if (i < (indexOfOne - 1)) {
                        int need = k - st.size();
                        if (need <= amountOfOne) {
                            for (int j = indexOfOne; j < indexOfOne + need; j++) {
                                Point pp = arrayCorrect.get(j);
                                st.offerLast(pp);
                            }
                            break;
                        } else {
                            rest = rest * st.peekLast().value;
                            lastBad = st.pollLast().value;
                            i = indexStack.pollLast();
                        }
                        continue;
                    }
                }
            }
            if (lastBad == cur) {
                Integer in = indexOfNextValue.get(cur);
                if (in == null) {
                    rest = rest * st.peekLast().value;
                    lastBad = st.pollLast().value;
                    i = indexStack.pollLast();
                } else {
                    i = in - 1;
                }
                continue;
            }
            if (cur > rest || (rest % cur != 0)) {
                lastBad = cur;
                if (i == end) {
                    while (i >= (end - 1)) {
                        lastBad = st.pollLast().value;
                        rest *= lastBad;
                        i = indexStack.pollLast();
                    }
                }
                continue;
            }
            rest = rest / cur;
            st.offerLast(p);
            indexStack.offerLast(i);
            if (st.size() == k && rest == 1) {
                break;
            }
            if (i == end) {
                while (i >= (end - 1)) {
                    lastBad = st.pollLast().value;
                    rest *= lastBad;
                    i = indexStack.pollLast();
                }
            }
        }
        for (Point point : st) {
            result.append(point.index + 1);
            result.append(" ");
        }
        return result.toString().trim();
    }
}

class Point implements Comparable<Point> {
    public Integer index;
    public Integer value;

    public Point(Integer index, Integer value) {
        this.index = index;
        this.value = value;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Point point = (Point) o;
        return index == point.index && value == point.value;
    }

    @Override
    public int hashCode() {
        return Objects.hash(index, value);
    }

    @Override
    public int compareTo(Point o) {
        return value.compareTo(o.value);
    }
}
*/

/*
//enter n, k and m: 7
//2
//27
//
//enter 7 numbers: 9
//1
//1
//27
//3
//27
//3
//
//9 * 3 = 27

#include <stdio.h>
#include <stdlib.h>

int found(int n, const int set[], int flags[], int i, int k, int m) {
    if (k <= 0)
        return m == 1;
    if (i >= n)
        return 0;
    if (set[i] == 0) {  // must special case 0
        if (m == 0) {
            if (i + k <= n) {
                while (k-- > 0)
                    flags[i++] = 1;
                return 1;
            }
            return 0;
        }
    } else if (m % set[i] == 0) {
        if (found(n, set, flags, i + 1, k - 1, m / set[i])) {
            flags[i] = 1;
            return 1;
        }
    }
    return found(n, set, flags, i + 1, k, m);
}

// get a number from the command line or from stdin
int getnum(int *index, char *argv[]) {
    int value;
    if (argv[*index]) {
        value = strtol(argv[(*index)++], NULL, 0);
        printf("%d ", value);
        return value;
    }
    if (scanf("%d", &value) != 1) {
        printf("invalid input\n");
        exit(1);
    }
    return value;
}

int main(int argc, char *argv[]) {
    int n, m, k, arg_i = 1;

    printf("enter n, k and m: ");
    n = getnum(&arg_i, argv);
    k = getnum(&arg_i, argv);
    m = getnum(&arg_i, argv);
    printf("\n");
    int set[n];
    int flags[n];
    printf("enter %d numbers: ", n);
    for (int i = 0; i < n; i++) {
        set[i] = getnum(&arg_i, argv);
        flags[i] = 0;
    }
    printf("\n");
    if (found(n, set, flags, 0, k, m)) {
        for (int i = 0, j = k; i < n; i++) {
            if (flags[i])
                printf("%d %c ", set[i], --j > 0 ? '*' : '=');
        }
        printf("%d\n", m);
    } else {
        printf("no solution\n");
    }
    return 0;
}
*/
