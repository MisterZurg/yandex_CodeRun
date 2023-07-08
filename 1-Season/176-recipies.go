package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Decimal string

func NewDecimal(value string) Decimal {
	return Decimal(value)
}

func (d Decimal) Add(other Decimal) Decimal {
	a, _ := DecimalToFloat64(d)
	b, _ := DecimalToFloat64(other)
	return NewDecimal(fmt.Sprintf("%.2f", a+b))
}

func (d Decimal) Mul(other Decimal) Decimal {
	a, _ := DecimalToFloat64(d)
	b, _ := DecimalToFloat64(other)
	return NewDecimal(fmt.Sprintf("%.2f", a*b))
}

func (d Decimal) Div(other Decimal) Decimal {
	a, _ := DecimalToFloat64(d)
	b, _ := DecimalToFloat64(other)
	return NewDecimal(fmt.Sprintf("%.2f", a/b))
}

func (d Decimal) LessThan(other Decimal) bool {
	a, _ := DecimalToFloat64(d)
	b, _ := DecimalToFloat64(other)
	return a < b
}

func DecimalToFloat64(d Decimal) (float64, error) {
	return strconv.ParseFloat(string(d), 64)
}

type Amount struct {
	Count Decimal
	Unit  string
}

func NewAmount(count Decimal, unit string) *Amount {
	return &Amount{
		Count: count,
		Unit:  unit,
	}
}

func (a *Amount) ConvertTo(unit string) *Amount {
	multipliersDict := map[string]Decimal{
		"g":    NewDecimal("1"),
		"kg":   NewDecimal("1000"),
		"ml":   NewDecimal("1"),
		"l":    NewDecimal("1000"),
		"cnt":  NewDecimal("1"),
		"tens": NewDecimal("10"),
	}

	if !isCompatible(a.Unit, unit) {
		panic("Incompatible units")
	}

	multiplier := a.Count.Mul(multipliersDict[a.Unit]).Div(multipliersDict[unit])

	return NewAmount(multiplier, unit)
}

func isCompatible(unit1, unit2 string) bool {
	unitClasses := [][]string{
		{"g", "kg"},
		{"ml", "l"},
		{"cnt", "tens"},
	}

	for _, unitClass := range unitClasses {
		if contains(unitClass, unit1) {
			return contains(unitClass, unit2)
		}
	}

	return false
}

func contains(units []string, unit string) bool {
	for _, u := range units {
		if u == unit {
			return true
		}
	}
	return false
}

type FoodInfo struct {
	Proteins      Decimal
	Fats          Decimal
	Carbohydrates Decimal
	FoodValue     Decimal
}

func NewFoodInfo(proteins, fats, carbohydrates, foodValue Decimal) *FoodInfo {
	return &FoodInfo{
		Proteins:      proteins,
		Fats:          fats,
		Carbohydrates: carbohydrates,
		FoodValue:     foodValue,
	}
}

func (fi *FoodInfo) Multiply(other Decimal) *FoodInfo {
	return NewFoodInfo(
		fi.Proteins.Mul(other),
		fi.Fats.Mul(other),
		fi.Carbohydrates.Mul(other),
		fi.FoodValue.Mul(other),
	)
}

func (fi *FoodInfo) Add(other *FoodInfo) *FoodInfo {
	return NewFoodInfo(
		fi.Proteins.Add(other.Proteins),
		fi.Fats.Add(other.Fats),
		fi.Carbohydrates.Add(other.Carbohydrates),
		fi.FoodValue.Add(other.FoodValue),
	)
}

type AmountFoodInfo struct {
	Amount   *Amount
	FoodInfo *FoodInfo
}

func NewAmountFoodInfo(amount *Amount, foodInfo *FoodInfo) *AmountFoodInfo {
	return &AmountFoodInfo{
		Amount:   amount,
		FoodInfo: foodInfo,
	}
}

type Ingredient struct {
	Name   string
	Amount *Amount
}

func NewIngredient(name string, amount *Amount) *Ingredient {
	return &Ingredient{
		Name:   name,
		Amount: amount,
	}
}

type CatalogIngredientInfo struct {
	Name   string
	Price  int
	Amount *Amount
}

func NewCatalogIngredientInfo(name string, price int, amount *Amount) *CatalogIngredientInfo {
	return &CatalogIngredientInfo{
		Name:   name,
		Price:  price,
		Amount: amount,
	}
}

type Dish struct {
	Name        string
	Count       int
	Ingredients []*Ingredient
}

func NewDish(name string, count int, ingredients []*Ingredient) *Dish {
	return &Dish{
		Name:        name,
		Count:       count,
		Ingredients: ingredients,
	}
}

func readDishes(scanner *bufio.Scanner) []*Dish {
	dishCountStr := scanner.Text()
	dishCount, err := strconv.Atoi(dishCountStr)
	if err != nil {
		panic(err)
	}

	dishes := make([]*Dish, 0, dishCount)
	for i := 0; i < dishCount; i++ {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		dishName := parts[0]
		dishCountStr := parts[1]
		ingredientCountStr := parts[2]

		ingredientCount, err := strconv.Atoi(ingredientCountStr)
		if err != nil {
			panic(err)
		}

		ingredients := make([]*Ingredient, 0, ingredientCount)
		for j := 0; j < ingredientCount; j++ {
			ingredientLine := scanner.Text()
			ingredientParts := strings.Split(ingredientLine, " ")
			ingredientName := ingredientParts[0]
			amountStr := ingredientParts[1]
			unit := ingredientParts[2]

			amount := NewDecimal(amountStr)

			ingredients = append(ingredients, NewIngredient(ingredientName, NewAmount(amount, unit)))
		}

		dishCount, err := strconv.Atoi(dishCountStr)
		if err != nil {
			panic(err)
		}

		dishes = append(dishes, NewDish(dishName, dishCount, ingredients))
	}

	return dishes
}

func readCatalog(scanner *bufio.Scanner) map[string]*CatalogIngredientInfo {
	ingredientCountStr := scanner.Text()
	ingredientCount, err := strconv.Atoi(ingredientCountStr)
	if err != nil {
		panic(err)
	}

	catalog := make(map[string]*CatalogIngredientInfo)
	for i := 0; i < ingredientCount; i++ {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		name := parts[0]
		priceStr := parts[1]
		amountStr := parts[2]
		unit := parts[3]

		price, err := strconv.Atoi(priceStr)
		if err != nil {
			panic(err)
		}

		amount := NewDecimal(amountStr)

		catalog[name] = NewCatalogIngredientInfo(name, price, NewAmount(amount, unit))
	}

	return catalog
}

func readFoodInfo(scanner *bufio.Scanner) map[string]*AmountFoodInfo {
	infoCountStr := scanner.Text()
	infoCount, err := strconv.Atoi(infoCountStr)
	if err != nil {
		panic(err)
	}

	foodInfo := make(map[string]*AmountFoodInfo)
	for i := 0; i < infoCount; i++ {
		line := scanner.Text()
		parts := strings.Split(line, " ")
		name := parts[0]
		amountStr := parts[1]
		unit := parts[2]
		proteinsStr := parts[3]
		fatsStr := parts[4]
		carbohydratesStr := parts[5]
		foodValueStr := parts[6]

		amount := NewDecimal(amountStr)
		proteins := NewDecimal(proteinsStr)
		fats := NewDecimal(fatsStr)
		carbohydrates := NewDecimal(carbohydratesStr)
		foodValue := NewDecimal(foodValueStr)

		foodInfo[name] = NewAmountFoodInfo(
			NewAmount(amount, unit),
			NewFoodInfo(proteins, fats, carbohydrates, foodValue),
		)
	}

	return foodInfo
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	dishes := readDishes(scanner)
	catalog := readCatalog(scanner)
	foodInfo := readFoodInfo(scanner)

	needIngredients := make(map[string]Decimal)
	needIngredientsCount := make(map[string]int)
	dishInfo := make(map[string]*FoodInfo)

	for _, dish := range dishes {
		for _, ingredient := range dish.Ingredients {
			convertedAmount := ingredient.Amount.ConvertTo(catalog[ingredient.Name].Amount.Unit)
			convertedFoodInfoAmount := foodInfo[ingredient.Name].Amount.ConvertTo(catalog[ingredient.Name].Amount.Unit)

			count := convertedAmount.Count.Mul(NewDecimal(strconv.Itoa(dish.Count)))
			needIngredients[ingredient.Name] = needIngredients[ingredient.Name].Add(count)

			ratio := count.Div(convertedFoodInfoAmount.Count)
			dishInfo[dish.Name] = dishInfo[dish.Name].Add(foodInfo[ingredient.Name].FoodInfo.Multiply(ratio))
		}
	}

	totalPrice := NewDecimal("0")
	for ingredientName, needIngredient := range needIngredients {
		needCount := needIngredient.Div(catalog[ingredientName].Amount.Count)
		if needCount.Mul(catalog[ingredientName].Amount.Count).LessThan(needIngredient) {
			needCount = needCount.Add(NewDecimal("1"))
		}

		dtflt64, _ := DecimalToFloat64(needCount)
		needIngredientsCount[ingredientName] = int(dtflt64)
		totalPrice = totalPrice.Add(NewDecimal(strconv.Itoa(catalog[ingredientName].Price)).Mul(needCount))
	}

	fmt.Println(totalPrice)
	for name, count := range needIngredientsCount {
		fmt.Println(name, count)
	}

	for name, info := range dishInfo {
		fmt.Println(name, info.Proteins, info.Fats, info.Carbohydrates, info.FoodValue)
	}
}

/*
#include <stdio.h>
#include <iostream>
#include <stdlib.h>
#include <string>
#include <string.h>
#include <vector>
#include <stack>
#include <queue>
#include <deque>
#include <set>
#include <map>
#include <assert.h>
#include <algorithm>
#include <iomanip>
#include <time.h>
#include <math.h>
#include <bitset>
#include <unordered_map>

#pragma comment(linker, "/STACK:256000000")

using namespace std;

typedef long long int ll;
typedef long double ld;

const int INF = 1000 * 1000 * 1000 + 21;
const ll LLINF = (1ll << 60) + 5;
const int MOD = 1000 * 1000 * 1000 + 7;

const int MAX_N = 10 * 1000 + 227;
const int MAX_LEN = 35;

struct ig_info {
    ll cost = 0;
    ll cnt = 0;
    ll buy_cnt = 0;
    ll eg_cnt = 0;
    ld arr[4] = {0.0, 0.0, 0.0, 0.0};
};

struct rs {
    string name;
    ll mult = 0;
    vector<pair<string, ll>> igs;
    ld arr[4] = {0.0, 0.0, 0.0, 0.0};
};

int n, k, m;
char buf[MAX_LEN];
char buf_cnt[MAX_LEN];
rs arr[MAX_N];
unordered_map<string, ig_info> info;

ll get_cnt() {
    ll cnt;
    scanf("%lld %s", &cnt, buf_cnt);
    if (!strcmp(buf_cnt, "kg") || !strcmp(buf_cnt, "l")) {
        return 1000 * cnt;
    }
    if (!strcmp(buf_cnt, "tens")) {
        return 10 * cnt;
    }
    return cnt;
}

int main() {
    scanf("%d", &n);
    for (int i = 0; i < n; ++i) {
        int sz;
        scanf("%s%lld%d", buf, &arr[i].mult, &sz);
        arr[i].name = buf;
        arr[i].igs.resize(sz);
        for (int j = 0; j < sz; ++j) {
            scanf("%s", buf);
            arr[i].igs[j].first = buf;
            arr[i].igs[j].second = get_cnt();
        }
    }

    scanf("%d", &k);
    for (int i = 0; i < k; ++i) {
        ll cost;
        scanf("%s%lld", buf, &cost);
        ig_info& cur_info = info[buf];
        cur_info.cost = cost;
        cur_info.cnt = get_cnt();
    }

    scanf("%d", &m);
    for (int i = 0; i < m; ++i) {
        scanf("%s", buf);
        ig_info& cur_info = info[buf];
        cur_info.eg_cnt = get_cnt();
        for (int j = 0; j < 4; ++j) {
            double x;
            scanf("%lf", &x);
            cur_info.arr[j] = x;
        }
        if (cur_info.cnt == 0) {
            info.erase(buf);
        }
    }

    for (int i = 0; i < n; ++i) {
        for (int j = 0; j < (int)arr[i].igs.size(); ++j) {
            ig_info& cur_info = info[arr[i].igs[j].first];
            cur_info.buy_cnt += arr[i].mult * arr[i].igs[j].second;
            for (int k = 0; k < 4; ++k) {
                arr[i].arr[k] += ((double)arr[i].igs[j].second / (double)cur_info.eg_cnt) * cur_info.arr[k];
            }
        }
    }

    ll ans = 0;
    for (auto& cur_info : info) {
        cur_info.second.buy_cnt = (cur_info.second.buy_cnt + cur_info.second.cnt - 1) / cur_info.second.cnt;
        ans += cur_info.second.buy_cnt * cur_info.second.cost;
    }

    printf("%lld\n", ans);

    for (const auto& cur_info : info) {
        printf("%s %lld\n", cur_info.first.c_str(), cur_info.second.buy_cnt);
    }

    for (int i = 0; i < n; ++i) {
        printf("%s ", arr[i].name.c_str());
        for (int j = 0; j < 4; ++j) {
            printf("%.20lf%c", (double)arr[i].arr[j], " \n"[j == 3]);
        }
    }

    return 0;
}
*/
