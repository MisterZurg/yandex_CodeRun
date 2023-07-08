package main

/*
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.*;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(reader.readLine());
        Pallet[] pallets = new Pallet[n];
        for (int i = 0; i < n; i++) {
            String[] parts = reader.readLine().split(" ");
            int w = Integer.parseInt(parts[0]);
            int h = Integer.parseInt(parts[1]);
            Pallet p = new Pallet(Math.max(w, h), Math.min(w, h));
            pallets[i] = p;
        }
        int result = solve(pallets, n);
        System.out.println(result);
        reader.close();
    }

    public static int solve(Pallet[] pallets, int n) {
        Arrays.sort(pallets);
        Map<Integer, Integer> highWidthMap = new HashMap<>(n);
        int maxWidth = pallets[0].getW();
        int maxHigh = pallets[0].getH();
        highWidthMap.put(maxHigh, maxWidth);
        List<Integer> hightsList = new ArrayList<>();
        hightsList.add(maxHigh);
        int result = 1;
        for (int i = 1; i < n; i++) {
            Pallet p = pallets[i];
            int w = p.getW();
            int h = p.getH();
            if (w == maxWidth || h >= maxHigh) {
                result++;
                if (h > maxHigh) {
                    maxHigh = h;
                    hightsList.add(maxHigh);
                }
            } else {
                boolean found = false;
                for (int j = hightsList.size() - 1; j >= 0; j--) {
                    Integer curH = hightsList.get(j);
                    if (curH <= h ) {
                        break;
                    }
                    Integer tempW = highWidthMap.get(curH);
                    if (w < tempW) {
                        found = true;
                        break;
                    }
                }
                if (!found) {
                    result++;
                }
            }
            Integer maxWForH = highWidthMap.get(h);
            if (maxWForH == null || w > maxWForH) {
                highWidthMap.put(h, w);
            }
        }
        return result;
    }
}

class Pallet implements Comparable<Pallet>{
    private int w;
    private int h;

    public Pallet(int w, int h) {
        this.w = w;
        this.h = h;
    }

    public int getW() {
        return w;
    }

    public int getH() {
        return h;
    }

    @Override
    public boolean equals(Object o) {
        if (this == o) return true;
        if (o == null || getClass() != o.getClass()) return false;
        Pallet poddon = (Pallet) o;
        return w == poddon.w && h == poddon.h;
    }

    @Override
    public int hashCode() {
        return Objects.hash(w, h);
    }

    @Override
    public int compareTo(Pallet o) {
        if (this.w == o.getW()) {
            return o.getH() - this.h;
        }

        return o.getW() - this.w;
    }
}
*/
