package main

import (
	"bufio"
	"fmt"
	"math"
	"math/big"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	nStr, _ := reader.ReadString('\n')
	n, _ := strconv.Atoi(strings.TrimSpace(nStr))

	var a, b, c int64
	parts, _ := reader.ReadString('\n')
	parts = strings.TrimSpace(parts)
	values := strings.Split(parts, " ")
	a, _ = strconv.ParseInt(values[0], 10, 64)
	b, _ = strconv.ParseInt(values[1], 10, 64)
	c, _ = strconv.ParseInt(values[2], 10, 64)

	res := solveCards(n, a, b, c)
	sb := strings.Builder{}
	for i := 0; i < len(res); i++ {
		sb.WriteString(strconv.FormatInt(res[i], 10))
		sb.WriteString(" ")
	}
	fmt.Println(strings.TrimSpace(sb.String()))
}

func solveCards(n int, a, b, c int64) []int64 {
	sum := (int64(1+n) * int64(n) / 2) - a
	bigN := bigInt(int64(n))
	bigNplus1 := bigInt(int64(n + 1))
	nnplus1 := bigN.Mul(bigN, bigNplus1)
	bb := bigInt(int64(1 + 2*n))
	bigSum2 := nnplus1.Mul(nnplus1, bb)
	bigSum2 = bigSum2.Div(bigSum2, bigInt(int64(6)))

	sum2 := bigSum2.Sub(bigSum2, bigInt(b)).Int64()

	bigSum3 := bigN.Mul(bigN, bigN)
	bigSum3 = bigSum3.Mul(bigSum3, bigNplus1)
	bigSum3 = bigSum3.Mul(bigSum3, bigNplus1)
	bigSum3 = bigSum3.Div(bigSum3, bigInt(int64(4)))

	sum3 := bigSum3.Sub(bigSum3, bigInt(c)).Int64()
	x, y, z := int64(0), int64(0), int64(0)

	for i := 1; i <= n; i++ {
		tempYDouble := (math.Sqrt(2*float64(sum2)-float64(sum)*float64(sum)+2*float64(i)*float64(sum)-3*float64(i)*float64(i)) + float64(sum) - float64(i)) / 2
		tempY := int64(tempYDouble)
		tempZLong := sum - int64(i) - tempY
		tempZ := int64(tempZLong)
		if int64(i)+tempY+tempZ == sum && int64(i)*int64(i)+tempY*tempY+tempZ*tempZ == sum2 &&
			int64(i)*int64(i)*int64(i)+tempY*tempY*tempY+tempZ*tempZ*tempZ == sum3 {
			x = int64(i)
			y = tempY
			z = tempZ
		}
	}
	return []int64{x, y, z}
}

func bigInt(n int64) *big.Int {
	return new(big.Int).SetInt64(n)
}

/*
import java.io.BufferedReader;
import java.io.IOException;
import java.io.InputStreamReader;
import java.math.BigInteger;

public class Main {
    public static void main(String[] args) throws IOException {
        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        int n = Integer.parseInt(reader.readLine());
        long a, b, c;
        String[] parts = reader.readLine().split(" ");
        a = Long.parseLong(parts[0]);
        b = Long.parseLong(parts[1]);
        c = Long.parseLong(parts[2]);
        long[] res = solve(n, a, b, c);
        StringBuilder sb = new StringBuilder();
        for (int i = 0; i < res.length; i++) {
            sb.append(res[i]);
            sb.append(" ");
        }
        System.out.println(sb.toString().trim());
    }

    public static long[] solve(int n, long a, long b, long c) {
        long sum = ((long) (1 + n) * n / 2) - a;
        BigInteger bigN = BigInteger.valueOf(n);
        BigInteger bigNplus1 = BigInteger.valueOf(n + 1);
        BigInteger nnplus1 = bigN.multiply(bigNplus1);
        BigInteger bb = BigInteger.valueOf((1 + 2L * n));
        BigInteger bigSum2 = nnplus1.multiply(bb).divide(BigInteger.valueOf(6));
        long sum2 = bigSum2.longValue() - b;
        BigInteger bigSum3 = bigN.multiply(bigN).multiply(bigNplus1).multiply(bigNplus1).divide(BigInteger.valueOf(4));
        long sum3 = bigSum3.longValue() - c;
        long x = 0, y = 0, z = 0;

        for (int i = 1; i <= n; i++) {
            double tempYDouble = (Math.pow(2 * sum2 - sum * sum + 2 * i * sum - 3L * i * i, 0.5) + sum - i) / 2;
            int tempY = (int) tempYDouble;
            long tempZLong = sum - i - tempY;
            int tempZ = (int) tempZLong;
            if (i + tempY + tempZ == sum && (long) i * i + (long) tempY * tempY + (long) tempZ * tempZ == sum2
            && (long) i * i * i + (long) tempY * tempY  * tempY + (long) tempZ * tempZ * tempZ == sum3) {
                x = i;
                y = tempY;
                z = tempZ;
            }
        }
        return new long[]{x, y, z};
    }
}
*/
