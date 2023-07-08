package main

// TODO : 6 Wrong answer
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	nStr, _ := reader.ReadString('\n')
//	nStr = strings.TrimSpace(nStr)
//	n, _ := strconv.Atoi(nStr)
//
//	deliveryStr, _ := reader.ReadString('\n')
//	deliveryStr = strings.TrimSpace(deliveryStr)
//	deliveryParts := strings.Split(deliveryStr, " ")
//	delivery := make([]int, n)
//	for i := 0; i < n; i++ {
//		delivery[i], _ = strconv.Atoi(deliveryParts[i])
//	}
//
//	parentStr, _ := reader.ReadString('\n')
//	parentStr = strings.TrimSpace(parentStr)
//	parentParts := strings.Split(parentStr, " ")
//	parent := make([]int, n)
//	pallets := make(map[int]bool)
//	for i := 0; i < n; i++ {
//		parent[i], _ = strconv.Atoi(parentParts[i])
//		if parent[i] == 0 {
//			pallets[i+1] = true
//		}
//	}
//
//	kStr, _ := reader.ReadString('\n')
//	kStr = strings.TrimSpace(kStr)
//	k, _ := strconv.Atoi(kStr)
//
//	notDelivery := make(map[int]bool)
//	if k != 0 {
//		notDeliveryStr, _ := reader.ReadString('\n')
//		notDeliveryStr = strings.TrimSpace(notDeliveryStr)
//		notDeliveryParts := strings.Split(notDeliveryStr, " ")
//		for i := 0; i < k; i++ {
//			notDeliveryNum, _ := strconv.Atoi(notDeliveryParts[i])
//			notDelivery[notDeliveryNum] = true
//		}
//	}
//
//	if n == k {
//		fmt.Println(0)
//	} else {
//		solve(notDelivery, delivery, parent, n, pallets)
//	}
//}
//
//func solve(notDelivery map[int]bool, delivery []int, parent []int, n int, pallets map[int]bool) {
//	badBoxAndPallet := make(map[int]bool)
//	for i := 0; i < n; i++ {
//		num := i + 1
//		if notDelivery[delivery[i]] {
//			if badBoxAndPallet[num] {
//				continue
//			}
//			p := parent[i]
//			if p == 0 {
//				delete(pallets, num)
//			} else {
//				needRemove := true
//				for {
//					if p-1 >= n || parent[p-1] == 0 {
//						break
//					} else {
//						p = parent[p-1]
//						if badBoxAndPallet[p] {
//							needRemove = false
//							break
//						} else {
//							badBoxAndPallet[p] = true
//						}
//					}
//				}
//				if needRemove {
//					delete(pallets, p)
//				}
//			}
//		}
//	}
//
//	fmt.Println(len(pallets))
//	for p := range pallets {
//		fmt.Println(p)
//	}
//}

/*
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(reader.readLine());
        String[] parts = reader.readLine().split(" ");
        int[] delivery = new int[n];
        for (int i = 0; i < n; i++) {
            delivery[i] = Integer.parseInt(parts[i]);
        }
        parts = reader.readLine().split(" ");
        int[] parent = new int[n];
        TreeSet<Integer> pallets = new TreeSet<>();
        for (int i = 0; i < n; i++) {
            parent[i] = Integer.parseInt(parts[i]);
            if (parent[i] == 0) {
                pallets.add(i + 1);
            }
        }
        int k = Integer.parseInt(reader.readLine());
        HashSet<Integer> notDelivery = new HashSet<>(k);
        if (k != 0) {
            parts = reader.readLine().split(" ");
            for (int i = 0; i < k; i++) {
                notDelivery.add(Integer.parseInt(parts[i]));
            }
        }
        if (n == k) {
            System.out.println(0);
        } else {
            solve(notDelivery, delivery, parent, n, pallets);
        }
        reader.close();
    }

    private static void solve(HashSet<Integer> notDelivery, int[] delivery, int[] parent, int n, TreeSet<Integer> pallets) {
        Set<Integer> badBoxAndPallet = new HashSet<>();
        for (int i = 0; i < n; i++) {
            int num = i + 1;
            if (notDelivery.contains(delivery[i])) {
                if (badBoxAndPallet.contains(num)) {
                    continue;
                }
                int p = parent[i];
                if (p == 0) {
                    pallets.remove(num);
                } else {
                    boolean needRemove = true;
                    while (true) {
                        if (p - 1 >= n) {
                            break;
                        }
                        if (parent[p - 1] == 0) {
                            break;
                        } else {
                            p = parent[p - 1];
                            if (badBoxAndPallet.contains(p)) {
                                needRemove = false;
                                break;
                            } else {
                                badBoxAndPallet.add(p);
                            }
                        }
                    }
                    if (needRemove) {
                        pallets.remove(p);
                    }
                }
            }
        }

        StringBuilder sb = new StringBuilder();
        sb.append(pallets.size());
        sb.append("\n");
        for (Integer p : pallets) {
            sb.append(p);
            sb.append("\n");
        }
        System.out.println(sb);
    }
}
*/
