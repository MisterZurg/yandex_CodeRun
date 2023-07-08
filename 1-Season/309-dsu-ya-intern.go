//package main
//
//import (
//	"bufio"
//	"fmt"
//	"os"
//	"strconv"
//	"strings"
//)
//
//type Vertex struct {
//	value int
//	group int
//}
//
//func main() {
//	reader := bufio.NewReader(os.Stdin)
//
//	parts, _ := reader.ReadString('\n')
//	parts = strings.TrimSuffix(parts, "\n")
//	splitParts := strings.Split(parts, " ")
//	n, _ := strconv.Atoi(splitParts[0])
//	m, _ := strconv.Atoi(splitParts[1])
//
//	graph := make(map[int][]int)
//	allNodes2 := make(map[int]bool)
//	allEdges := make(map[int][2]int)
//
//	mapVer := make(map[int]Vertex)
//	for i := 1; i <= n; i++ {
//		allNodes2[i] = true
//		graph[i] = []int{}
//		mapVer[i] = Vertex{i, 0}
//	}
//
//	for i := 0; i < m; i++ {
//		parts, _ := reader.ReadString('\n')
//		parts = strings.TrimSuffix(parts, "\n")
//		splitParts := strings.Split(parts, " ")
//		x, _ := strconv.Atoi(splitParts[0])
//		y, _ := strconv.Atoi(splitParts[1])
//		e := [2]int{x, y}
//		currentEdgeList := graph[x]
//		currentEdgeList = append(currentEdgeList, y)
//		graph[x] = currentEdgeList
//		if x != y {
//			currentEdgeList = graph[y]
//			currentEdgeList = append(currentEdgeList, x)
//			graph[y] = currentEdgeList
//		}
//		allEdges[i+1] = e
//	}
//
//	ar := 0
//	for _, e := range allEdges {
//		x := e[0]
//		y := e[1]
//		vx := mapVer[x]
//		vy := mapVer[y]
//
//		if _, ok := allNodes2[x]; ok && allNodes2[y] {
//			ar++
//			vx.group = ar
//			vy.group = ar
//		} else if vx.group != vy.group && vx.group != 0 && vy.group != 0 {
//			ar--
//			vx.group = vy.group
//			for k, v := range mapVer {
//				if v.group == vx.group {
//					v.group = vy.group
//					mapVer[k] = v
//				}
//			}
//		} else {
//			if _, ok := allNodes2[x]; ok {
//				vx.group = vy.group
//			}
//			if _, ok := allNodes2[y]; ok {
//				vy.group = vx.group
//			}
//		}
//		delete(allNodes2, x)
//		delete(allNodes2, y)
//	}
//
//	count := ar + len(allNodes2)
//
//	qStr, _ := reader.ReadString('\n')
//	qStr = strings.TrimSuffix(qStr, "\n")
//	q, _ := strconv.Atoi(qStr)
//
//	parts, _ = reader.ReadString('\n')
//	parts = strings.TrimSuffix(parts, "\n")
//	splitParts = strings.Split(parts, " ")
//
//	var sb strings.Builder
//	for j := 0; j < q; j++ {
//		i, _ := strconv.Atoi(splitParts[j])
//		ed := allEdges[i]
//		list := graph[ed[0]]
//		for k, v := range list {
//			if v == ed[1] {
//				list = append(list[:k], list[k+1:]...)
//				break
//			}
//		}
//		if ed[1] != ed[0] {
//			list = graph[ed[1]]
//			for k, v := range list {
//				if v == ed[0] {
//					list = append(list[:k], list[k+1:]...)
//					break
//				}
//			}
//		}
//		if !isConnectedBFS(graph, ed) {
//			count++
//		}
//		sb.WriteString(strconv.Itoa(count))
//		sb.WriteString(" ")
//	}
//	fmt.Println(strings.TrimSpace(sb.String()))
//}
//
//func isConnectedBFS(graph map[int][]int, ed [2]int) bool {
//	result := false
//	if ed[0] == ed[1] {
//		return true
//	}
//	nodeQueue := make([]int, 0)
//	nodeQueue = append(nodeQueue, ed[0])
//	found := false
//	doneVertex := make(map[int]bool)
//	for len(nodeQueue) > 0 && !found {
//		n := nodeQueue[0]
//		nodeQueue = nodeQueue[1:]
//		doneVertex[n] = true
//		currentEdgeList := graph[n]
//		for _, i := range currentEdgeList {
//			if i == ed[1] {
//				found = true
//				result = true
//				break
//			}
//			if !doneVertex[i] {
//				nodeQueue = append(nodeQueue, i)
//				doneVertex[i] = true
//			}
//		}
//	}
//	return result
//}

/*
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.time.Duration;
import java.time.LocalDateTime;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));

        String[] parts = reader.readLine().split(" ");
        int n = Integer.parseInt(parts[0]);
        int m = Integer.parseInt(parts[1]);


        Map<Integer, List<Integer>> graph = new HashMap<>(n);
        Set<Integer> allNodes2 = new HashSet<>(n);
        Map<Integer, Integer[]> allEdges = new HashMap();

        Map <Integer, Vertex> mapVer = new HashMap<>();
        for (Integer i = 1; i <= n; i++) {
            allNodes2.add(i);
            graph.put(i, new ArrayList<>());
            mapVer.put(i, new Vertex(i, 0));
        }

        for (int i = 0; i < m; i++) {
            parts = reader.readLine().split(" ");
            int x = Integer.parseInt(parts[0]);
            int y = Integer.parseInt(parts[1]);
            Integer[] e = new Integer[] {x, y};
            List<Integer> currentEdgeList = graph.get(x);
            currentEdgeList.add(y);
            if (x != y) {
                currentEdgeList = graph.get(y);
                currentEdgeList.add(x);
            }
            allEdges.put(i + 1, e);
        }

        int ar = 0;
        for (Integer key: allEdges.keySet()) {
            Integer[] e = allEdges.get(key);
            Integer x = e[0];
            Integer y = e[1];
            Vertex vx = mapVer.get(x);
            Vertex vy = mapVer.get(y);

            if (allNodes2.contains(x) && allNodes2.contains(y)) {
                ar++;
                vx.group = ar;
                vy.group = ar;
            } else if (!Objects.equals(vx.group, vy.group) && vx.group != 0 && vy.group != 0) {
                ar--;
                vx.group = vy.group;
                for (Integer k : mapVer.keySet()) {
                    Vertex v = mapVer.get(k);
                    if (Objects.equals(v.group, vx.group)) {
                        v.group = vy.group;
                    }
                }
            } else {
                if (allNodes2.contains(x)) {
                    vx.group = vy.group;
                }
                if (allNodes2.contains(y)) {
                    vy.group = vx.group;
                }
            }
            allNodes2.remove(x);
            allNodes2.remove(y);
        }

        int count = ar + allNodes2.size();

        int q = Integer.parseInt(reader.readLine());
        parts = reader.readLine().split(" ");

        StringBuilder sb = new StringBuilder();
        for (int j = 0; j < q; j++) {
            int i = Integer.parseInt(parts[j]);
            Integer[] ed = allEdges.get(i);
            List<Integer> list = graph.get(ed[0]);
            list.remove(ed[1]);
            if (ed[1] != ed[0]) {
                list = graph.get(ed[1]);
                list.remove(ed[0]);
            }
            if (!isConnectedBFS(graph, ed)) {
                count++;
            }
            sb.append(count);
            sb.append(" ");
        }
        System.out.println(sb.toString().trim());
        reader.close();

    }
    private static boolean isConnectedBFS(Map<Integer, List<Integer>> graph, Integer[] ed) {
        boolean result = false;
        if (ed[0] == ed[1]) {
            return true;
        }
        Queue<Integer> nodeQueue = new LinkedList<>();
        nodeQueue.offer(ed[0]);
        boolean found = false;
        Set<Integer> doneVertex = new HashSet<>();
        while (!nodeQueue.isEmpty() && !found) {
            Integer n = nodeQueue.poll();
            doneVertex.add(n);
            List<Integer> currentEdgeList = graph.get(n);
            for (Integer i: currentEdgeList) {
                if (i == ed[1]) {
                    found = true;
                    result = true;
                    break;
                }
                if (!doneVertex.contains(i)) {
                    nodeQueue.offer(i);
                    doneVertex.add(i);
                }
            }
        }
        return result;
    }

    private static int getAreas(Map<Integer, List<Integer>> graph, Set<Integer> allNodes) {
        Stack<Integer> nodeStack = new Stack<>();
        List<Set<Integer>> teams = new ArrayList();
        while (!allNodes.isEmpty()) {
            Set<Integer> tempTeam = new HashSet<>();
            Set<Integer> team = new HashSet<>();
            Iterator<Integer> it = allNodes.iterator();
            Integer startNode = it.next();
            nodeStack.push(startNode);
            allNodes.remove(startNode);
            while (!nodeStack.empty()) {
                boolean end = false;
                List<Integer> currentEdgeList = graph.get(nodeStack.peek());
                tempTeam.add(nodeStack.peek());
                if (currentEdgeList.size() > 0) {
                    boolean hasItem = false;
                    Iterator<Integer> iterator = currentEdgeList.iterator();
                    while (iterator.hasNext()) {
                        Integer n = iterator.next();
                        if (allNodes.contains(n)) {
                            nodeStack.push(n);
                            allNodes.remove(n);
                            hasItem = true;
                            break;
                        }
                    }
                    end = !hasItem;
                } else {
                    end = true;
                }
                if (end) {
                    team.add(nodeStack.pop());
                }
            }
            teams.add((HashSet<Integer>) team);
        }
        return teams.size();
    }

    private static int getAreas2(Map<Integer, List<Integer>> graph, Set<Integer> allNodes) {
        Queue<Integer> nodeQueue = new LinkedList<>();
        int result = 0;
        while (!allNodes.isEmpty()) {
            Iterator<Integer> it = allNodes.iterator();
            Integer startNode = it.next();
            nodeQueue.offer(startNode);
            allNodes.remove(startNode);
            while (!nodeQueue.isEmpty()) {
                Integer n = nodeQueue.poll();
                allNodes.remove(n);
                List<Integer> currentEdgeList = graph.get(n);
                for (Integer i: currentEdgeList) {
                    if (allNodes.contains(i)) {
                        nodeQueue.offer(i);
                    }
                }
            }
            result++;
        }
        return result;
    }

    private static boolean isConnectedDFS(Map<Integer, List<Integer>> graph, Integer[] ed) {
        boolean result = false;
        if (ed[0] == ed[1]) {
            return true;
        }
        Stack<Integer> nodeStack = new Stack<>();
        nodeStack.add(ed[0]);
        boolean found = false;
        Set<Integer> doneVertex = new HashSet<>();
        while (!nodeStack.isEmpty() && !found) {
            doneVertex.add(nodeStack.peek());
            List<Integer> currentEdgeList = graph.get(nodeStack.peek());
            if (currentEdgeList.size() > 0) {
                Iterator<Integer> iterator = currentEdgeList.iterator();
                boolean hasItem = false;
                while (iterator.hasNext()) {
                    Integer next = iterator.next();
                    if (next == ed[1]) {
                        found = true;
                        result = true;
                        break;
                    }
                    if (!doneVertex.contains(next)) {
                        nodeStack.add(next);
                        hasItem = true;
                        break;
                    }
                }
                if (!hasItem) {
                    nodeStack.pop();
                }
            } else {
                nodeStack.pop();
            }
        }
        return result;
    }
}

class Vertex {
    public Integer value;
    public Integer group;

    public Vertex(Integer value, Integer group) {
        this.value = value;
        this.group = group;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Vertex edge = (Vertex) o;
        return value == edge.value && group == edge.group;
    }

    @Override
    public int hashCode() {
        return Objects.hash(value, group);
    }

}
 */