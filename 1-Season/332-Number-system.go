package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"regexp"
	"strconv"
	"strings"
)

var (
	alphabet    = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	alphabetAll = alphabet + "+-="
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	lines := strings.Split(sc.Text(), "\n")
	for _, phrase := range lines {
		phrase = stripChars(phrase, alphabetAll)

		minBase := getMaxBase(phrase) + 1
		maxBase := 36

		leftPart, rightPart := splitEquation(phrase)
		for base := minBase; base <= maxBase; base++ {
			left := getValue(leftPart, base)
			right := getValue(rightPart, base)
			if left == right {
				fmt.Println(base)
				break
			}
		}
		fmt.Println(-1)
	}
}

func stripChars(phrase, validChars string) string {
	var result strings.Builder
	for _, c := range phrase {
		if strings.ContainsRune(validChars, c) {
			result.WriteRune(c)
		}
	}
	return result.String()
}

func getMaxBase(phrase string) int {
	maxBase := -1
	for _, c := range phrase {
		if strings.ContainsRune(alphabet, c) {
			index := strings.IndexRune(alphabet, c)
			if index > maxBase {
				maxBase = index
			}
		}
	}
	return maxBase
}

func splitEquation(phrase string) (string, string) {
	parts := strings.SplitN(phrase, "=", 2)
	if len(parts) != 2 {
		return "", ""
	}
	return parts[0], parts[1]
}

func getValue(phrase string, base int) float64 {
	result := 0.0
	sign := 1.0
	re := regexp.MustCompile(`([+-])`)
	for _, part := range re.Split(phrase, -1) {
		if part == "+" {
			sign = 1.0
		} else if part == "-" {
			sign = -1.0
		} else {
			value, _ := strconv.ParseFloat(part, base)
			result += math.Copysign(value, sign)
		}
	}
	return result
}

/*// TODO : HT 115/250 Wrong answer fuck u
import java.io.BufferedReader;
import java.math.BigDecimal;
import java.io.IOException;
import java.io.InputStreamReader;
import java.util.*;


public class Main {
    public static String result = "no_data";

    public static void main(String[] args) {
        try (BufferedReader reader = new BufferedReader(new InputStreamReader(System.in))) {

            // BufferedReader reader = new BufferedReader(new InputStreamReader(System.in));
            // int n = Integer.parseInt(reader.readLine());
            // Число связей между серверами.
            String line = reader.readLine().toUpperCase();
            String[] lines = line.split("=");
            int max_base = 26 + 10;
            int base = 2;
            for (char chr : line.toCharArray()) {
                if (chr >= '0' && chr <= '9') {
                    int temp = chr - 47;
                    if (base < temp)
                        base = temp;
                } else if (chr >= 'A' && chr <= 'Z') {
                    int temp = chr - 54;
                    if (base < temp)
                        base = temp;
                }
            }
            for (; base <= max_base; base++) {
                BigDecimal big_base = new BigDecimal(base);
                BigDecimal res1 = get_result(lines[0], big_base);
                BigDecimal res2 = get_result(lines[1], big_base);
                if (res1.equals(res2))
                    break;
            }
            if (base <= 36) result = String.valueOf(base);
            else result = "-1";

            System.out.println(result);
        } catch (Exception e) {
            System.out.println(e.getMessage());
        }
    }

    private static BigDecimal get_result(String line, BigDecimal base) {
        BigDecimal res =BigDecimal.ZERO;
        BigDecimal val = BigDecimal.ZERO;
        BigDecimal m = BigDecimal.ONE;
        for (char chr : line.toCharArray()) {
            if (Character.isDigit(chr)) {
                val = val.multiply(base).add(new BigDecimal(chr - 48));
            } else if (Character.isAlphabetic(chr)) {
                val = val.multiply(base).add(new BigDecimal(chr - 55));
            } else if (chr == '+') {
                res = res.add(val.multiply(m));
                val = BigDecimal.ZERO;
                m = BigDecimal.ONE;
            } else if (chr == '-') {
                res = res.add(val.multiply(m));
                val = BigDecimal.ZERO;
                m = new BigDecimal(-1);
            }
        }
        res = res.add(val.multiply(m));

        return res;
    }
}
*/

/*
// TODO : HT 10 Wrong answer
import re
from math import copysign

alphabet = '0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ'
alphabet_all = alphabet + '+-='

with open('input.txt', 'r') as f:
    for phrase in f.readlines():
        phrase = ''.join(c for c in phrase.strip() if c in alphabet_all)

        min_base = max([
            alphabet.index(c) if c in alphabet else -1
            for c in phrase
        ]) + 1
        max_base = 36

        def get_value(phrase, base):
            result = 0
            sign = 1
            for part in re.split(r'([+-])', phrase):
                if part in ['+', '-']:
                    sign = [-1, 1][part == '+']
                else:
                    value = int(part, base)
                    result += copysign(value, sign)
            return result


        left_part, right_part = phrase.split('=')
        for base in range(min_base, max_base + 1):
            left_ = get_value(left_part, base)
            right_ = get_value(right_part, base)
            if left_ == right_:
                print(base)
                break
        else:
            print(-1)
*/
