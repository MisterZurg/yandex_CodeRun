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
//func library() {
//	sc := bufio.NewScanner(os.Stdin)
//	sc.Scan()
//	kmc := sc.Text()
//	vars := strings.Split(kmc, " ")
//	k, _ := strconv.Atoi(vars[0])
//	m, _ := strconv.Atoi(vars[1])
//	d, _ := strconv.Atoi(vars[2])
//	// fmt.Scan(&k, &m, &d)
//	count := 1
//	firstWeekPass := true
//	for d != 1 {
//		if d < 6 {
//			m += k
//		}
//		if m < count {
//			firstWeekPass = false
//			break
//		}
//		m -= count
//		count++
//		d = (d % 7) + 1
//	}
//	if firstWeekPass {
//		for m+5*k >= 7*count+21 {
//			m = m + 5*k - (7*count + 21)
//			count += 7
//		}
//		for {
//			if d < 6 {
//				m += k
//			}
//			if m < count {
//				break
//			}
//			m -= count
//			count++
//			d = (d % 7) + 1
//		}
//	}
//	fmt.Println(count - 1)
//}
//
//func main() {
//	library()
//}

/*
import java.io.BufferedReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.math.BigDecimal;
import java.math.MathContext;
import java.math.RoundingMode;

public class Main {

    public static void main(String[] args) throws IOException {

        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        String[] parts = reader.readLine().split(" ");
        long k = Long.parseLong(parts[0]);
        long m = Long.parseLong(parts[1]);
        int d = Integer.parseInt(parts[2]);

        System.out.println(solveFast(k, m, d));

       reader.close();
    }

    public static long getRandomNumber(long min, long max) {
        return (long) ((Math.random() * (max - min)) + min);
    }
    public static int getRandomNumberInt(int min, int max) {
        return (int) ((Math.random() * (max - min)) + min);
    }


    private static long solveFast(long k, long m, int d) {
        if (d > 5 && m < 3) {
            if (m == 0) return 0;
            if (d != 7) return solveSlow(k, m, d);
        }
        double a = -24.5;
        double b = 5 * k - 3.5;
        double c = m;
        double[] s = solveQuadraticEquation(a, b, c);
        double q = Math.max(s[0], s[1]);
        long numberOfWeek = (long) q;
        long result = numberOfWeek * 7 + 1;
        long perDay;
        double countBeginD = -24.5 * numberOfWeek * numberOfWeek + numberOfWeek * (5*k-3.5) + m;
        long countBegin = (long) countBeginD;
        for (int i = 1; i <= 7; i++) {
            perDay = (d > 5 ? 0 : k);
            countBegin = countBegin + perDay - result;
            if (countBegin < 0) {
                break;
            }
            result++;
            d = d % 7 + 1;
        }
        return result - 1;
    }

    private static double[] solveQuadraticEquation(double a, double b, double c) {
        double[] result = new double[2];
        if (b == 0 && c == 0) {
            result[0] = 0;
            return result;
        }
        if (b == 0 && a != 0 && c != 0) {
            double n = -c/a;
            if (n < 0 ) return null;
            result[0] = Math.sqrt(n);
            result[1] = -result[0];
            return result;
        }
        if (a == 0 && b != 0) {
            result[0] = -c / b;
            return result;
        }
        double d = b * b - 4 * a * c;
        if (d < 0) {
            return null;
        }
        if (d == 0) {
            result[0] = -b / (2 * a);
            return result;
        }
        result[0] = (-b - Math.sqrt(d)) / (2 * a);
        result[1] = (-b + Math.sqrt(d)) / (2 * a);
        return result;
    }

    private static double[] solveQuadraticEquationBigDecimal(BigDecimal a, BigDecimal b, BigDecimal c) {
        double[] result = new double[2];
        if (b.compareTo(BigDecimal.ZERO) == 0 && c.compareTo(BigDecimal.ZERO) == 0) {
            result[0] = 0;
            return result;
        }
        if (b.compareTo(BigDecimal.ZERO) == 0 && a.compareTo(BigDecimal.ZERO) != 0 && c.compareTo(BigDecimal.ZERO) != 0) {
            BigDecimal n = c.multiply(new BigDecimal(-1)).divide(a);

            if (n.compareTo(BigDecimal.ZERO) == -1) {
				return null;
			}
            result[0] = n.sqrt(new MathContext(10)).doubleValue();
            result[1] = -result[0];
            return result;
        }
        if (a.compareTo(BigDecimal.ZERO) == 0 && b.compareTo(BigDecimal.ZERO) != 0) {
            result[0] = c.multiply(new BigDecimal(-1)).divide(b).doubleValue();
            return result;
        }
        BigDecimal d = b.multiply(b).subtract(a.multiply(c).multiply(new BigDecimal(4)));

        if (d.compareTo(BigDecimal.ZERO) == -1) {
            return null;
        }
        if (d.compareTo(BigDecimal.ZERO) == 0) {
            b.multiply(new BigDecimal(-1)).divide(a.multiply(new BigDecimal(2))).doubleValue();
            return result;
        }

        BigDecimal top = b.multiply(new BigDecimal(-1)).subtract(d.sqrt(new MathContext(2)));
        BigDecimal top2 = b.multiply(new BigDecimal(-1)).add(d.sqrt(new MathContext(2)));
        BigDecimal bottom = a.multiply(new BigDecimal(2));

        result[0] = top.divide(bottom, 2, RoundingMode.HALF_UP).doubleValue();
        result[1] = top2.divide(bottom, 2, RoundingMode.HALF_UP).doubleValue();

        return result;
    }

}
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"math/rand"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')
	parts := splitString(line)
	k, _ := strconv.ParseInt(parts[0], 10, 64)
	m, _ := strconv.ParseInt(parts[1], 10, 64)
	d, _ := strconv.ParseInt(parts[2], 10, 32)

	fmt.Println(solveFast(k, m, int(d)))
}

func splitString(str string) []string {
	str = strings.TrimSuffix(str, "\n")
	return strings.Split(str, " ")
}

func getRandomNumber(min, max int64) int64 {
	return min + int64(math.Floor(float64(max-min)*rand.Float64()))
}

func getRandomNumberInt(min, max int) int {
	return min + int(math.Floor(float64(max-min)*rand.Float64()))
}

func solveFast(k, m int64, d int) int64 {
	if d > 5 && m < 3 {
		if m == 0 {
			return 0
		}
		if d != 7 {
			return solveSlow(k, m, d)
		}
	}
	a := -24.5
	b := 5*float64(k) - 3.5
	c := float64(m)
	s := solveQuadraticEquation(a, b, c)
	q := math.Max(s[0], s[1])
	numberOfWeek := int64(q)
	result := numberOfWeek*7 + 1
	var perDay int64
	countBeginD := -24.5*float64(numberOfWeek)*float64(numberOfWeek) + float64(numberOfWeek)*(5*float64(k)-3.5) + float64(m)
	countBegin := int64(countBeginD)
	for i := 1; i <= 7; i++ {
		perDay = int64(d % 7)
		countBegin = countBegin + perDay - result
		if countBegin < 0 {
			break
		}
		result++
		d = (d % 7) + 1
	}
	return result - 1
}

func solveQuadraticEquation(a, b, c float64) []float64 {
	result := make([]float64, 2)
	if b == 0 && c == 0 {
		result[0] = 0
		return result
	}
	if b == 0 && a != 0 && c != 0 {
		n := -c / a
		if n < 0 {
			return nil
		}
		result[0] = math.Sqrt(n)
		result[1] = -result[0]
		return result
	}
	if a == 0 && b != 0 {
		result[0] = -c / b
		return result
	}
	d := b*b - 4*a*c
	if d < 0 {
		return nil
	}
	if d == 0 {
		result[0] = -b / (2 * a)
		return result
	}
	result[0] = (-b - math.Sqrt(d)) / (2 * a)
	result[1] = (-b + math.Sqrt(d)) / (2 * a)
	return result
}

func solveSlow(k, m int64, d int) int64 {
	count := int64(1)
	firstWeekPass := true
	for d != 1 {
		if d < 6 {
			m += k
		}
		if m < count {
			firstWeekPass = false
			break
		}
		m -= count
		count++
		d = (d % 7) + 1
	}
	if firstWeekPass {
		for m+5*k >= 7*count+21 {
			m = m + 5*k - (7*count + 21)
			count += 7
		}
		for {
			if d < 6 {
				m += k
			}
			if m < count {
				break
			}
			m -= count
			count++
			d = (d % 7) + 1
		}
	}
	// fmt.Println(count - 1)
	return count - 1
}

/*
import java.io.BufferedReader;
import java.io.FileWriter;
import java.io.IOException;
import java.io.InputStreamReader;
import java.math.BigDecimal;
import java.math.MathContext;
import java.math.RoundingMode;

public class Main {

    public static void main(String[] args) throws IOException {

        BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
        String[] parts = reader.readLine().split(" ");
        long k = Long.parseLong(parts[0]);
        long m = Long.parseLong(parts[1]);
        int d = Integer.parseInt(parts[2]);


        System.out.println(solveFast(k, m, d));
        System.out.println(solveSlow(k, m, d));

       reader.close();
    }

    public static long getRandomNumber(long min, long max) {
        return (long) ((Math.random() * (max - min)) + min);
    }
    public static int getRandomNumberInt(int min, int max) {
        return (int) ((Math.random() * (max - min)) + min);
    }


    private static long solveFast(long k, long m, int d) {
        if (d > 5 && m < 3) {
            if (m == 0) return 0;
            if (d != 7) return solveSlow(k, m, d);
        }
        double a = -24.5;
        double b = 5 * k - 3.5;
        double c = m;
        double[] s = solveQuadraticEquation(a, b, c);
        double q = Math.max(s[0], s[1]);
        long numberOfWeek = (long) q;
        long result = numberOfWeek * 7 + 1;
        long perDay;
        double countBeginD = -24.5 * numberOfWeek * numberOfWeek + numberOfWeek * (5*k-3.5) + m;
        long countBegin = (long) countBeginD;
        for (int i = 1; i <= 7; i++) {
            perDay = (d > 5 ? 0 : k);
            countBegin = countBegin + perDay - result;
            if (countBegin < 0) {
                break;
            }
            result++;
            d = d % 7 + 1;
        }
        return result - 1;
    }

    private static double[] solveQuadraticEquation(double a, double b, double c) {
        double[] result = new double[2];
        if (b == 0 && c == 0) {
            result[0] = 0;
            return result;
        }
        if (b == 0 && a != 0 && c != 0) {
            double n = -c/a;
            if (n < 0 ) return null;
            result[0] = Math.sqrt(n);
            result[1] = -result[0];
            return result;
        }
        if (a == 0 && b != 0) {
            result[0] = -c / b;
            return result;
        }
        double d = b * b - 4 * a * c;
        if (d < 0) {
            return null;
        }
        if (d == 0) {
            result[0] = -b / (2 * a);
            return result;
        }
        result[0] = (-b - Math.sqrt(d)) / (2 * a);
        result[1] = (-b + Math.sqrt(d)) / (2 * a);
        return result;
    }

    private static double[] solveQuadraticEquationBigDecimal(BigDecimal a, BigDecimal b, BigDecimal c) {
        double[] result = new double[2];
        if (b.compareTo(BigDecimal.ZERO) == 0 && c.compareTo(BigDecimal.ZERO) == 0) {
            result[0] = 0;
            return result;
        }
        if (b.compareTo(BigDecimal.ZERO) == 0 && a.compareTo(BigDecimal.ZERO) != 0 && c.compareTo(BigDecimal.ZERO) != 0) {
            BigDecimal n = c.multiply(new BigDecimal(-1)).divide(a);

            if (n.compareTo(BigDecimal.ZERO) == -1) return null;

            result[0] = n.sqrt(new MathContext(10)).doubleValue();
            result[1] = -result[0];
            return result;
        }
        if (a.compareTo(BigDecimal.ZERO) == 0 && b.compareTo(BigDecimal.ZERO) != 0) {
            result[0] = c.multiply(new BigDecimal(-1)).divide(b).doubleValue();
            return result;
        }
        BigDecimal d = b.multiply(b).subtract(a.multiply(c).multiply(new BigDecimal(4)));

        if (d.compareTo(BigDecimal.ZERO) == -1) {
            return null;
        }
        if (d.compareTo(BigDecimal.ZERO) == 0) {
            b.multiply(new BigDecimal(-1)).divide(a.multiply(new BigDecimal(2))).doubleValue();
            return result;
        }

        BigDecimal top = b.multiply(new BigDecimal(-1)).subtract(d.sqrt(new MathContext(2)));
        BigDecimal top2 = b.multiply(new BigDecimal(-1)).add(d.sqrt(new MathContext(2)));
        BigDecimal bottom = a.multiply(new BigDecimal(2));

        return result;
    }

    private static int solveSlow(long k, long m, int d) {
        int count = 0;
        long rest = m;
        long perDay;
        while (rest >= 0) {
            perDay = (d > 5 ? 0 : k);
            count++;
            rest = rest + perDay - count;
            d = d % 7 + 1;
        }
        return count - 1;
    }
}
*/
