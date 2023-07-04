package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type NumberFormat struct {
	countryCode  string // take meee home
	operatorCode string
	number       string
	name         string
}

func parseFormat(rawFormat string) NumberFormat {
	// +875 (29) 1XXXXX - Atlantis MythCell
	parts := strings.Split(rawFormat, " ")
	return NumberFormat{
		countryCode:  parts[0][1:],                  // skip +
		operatorCode: parts[1][1 : len(parts[1])-1], // skip ()
		number:       parts[2],
		name:         parts[4] + " " + parts[5],
	}
}

func preporocess(rawNumber string) string {
	if len(rawNumber) == 0 {
		return rawNumber
	}
	output := ""
	for i := 0; i < len(rawNumber); i++ {
		// symb := rawNumber[i]
		if rawNumber[i] > 47 && rawNumber[i] < 58 {
			output += string(rawNumber[i])
		}
	}
	return output
}

func (nf *NumberFormat) match(number string) bool {
	cnt := len(nf.countryCode) + len(nf.operatorCode) + len(nf.number)

	if cnt != len(number) { // check number and format length
		return false
	}
	// Compare numbers country code
	if strings.Compare(number[0:len(nf.countryCode)], nf.countryCode) != 0 {
		return false
	}

	// Compare numbers operator code
	if strings.Compare(number[len(nf.countryCode):len(nf.countryCode)+len(nf.operatorCode)], nf.operatorCode) != 0 {
		return false
	}

	for i := 0; i < len(nf.number); i++ {
		if nf.number[i] == 'X' {
			continue
		} else if nf.number[i] != number[len(nf.countryCode)+len(nf.operatorCode)+i] {
			return false
		}
	}
	return true
}

func (nf *NumberFormat) getFormattedNumber(number string) string {
	// +875 (29) 1XXXXX - Atlantis MythCell
	return fmt.Sprintf("+%s (%s) %s - %s", nf.countryCode, nf.operatorCode, number[len(nf.countryCode)+len(nf.operatorCode):], nf.name)
}

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Scan()

	n, _ := strconv.Atoi(sc.Text())
	numbers := make([]string, n)
	for i := range numbers {
		sc.Scan()
		raw := sc.Text()
		numbers[i] = preporocess(raw)
	}

	sc.Scan()

	// количество шаблонов
	m, _ := strconv.Atoi(sc.Text())
	patterns := make([]NumberFormat, m)

	for i := range patterns {
		sc.Scan()
		raw := sc.Text()

		patterns[i] = parseFormat(raw)
	}

	// fmt.Println(patterns)

	var formattedNumbers []string
	for _, num := range numbers { // for each number
		for _, pt := range patterns { // for each format
			if pt.match(num) { // if number fits format, then get formatted number and push it in vector
				formattedNumbers = append(formattedNumbers, pt.getFormattedNumber(num))
			}
		}
	}

	for _, fn := range formattedNumbers {
		fmt.Println(fn)
	}
}

/*
#include <iostream>
#include <string>
#include <vector>

std::string split(std::string input)
{
    if (input.empty())
        return input;
    std::string output(std::move(input));

    for (int i = 0; i < output.size(); i++)
    {
        char symb = (char) output[i];
        if (!(symb > 47 && symb < 58))
        {
            output.erase(i, 1);
            i--;
        }
    }
    return output;
}

class NumberFormat
{
protected:
    std::string country_code;
    std::string operator_code;
    std::string personal_number;
    std::string name;
public:
    explicit NumberFormat(const std::string& format)
    {
        read_format(format);
    }

    void read_format(const std::string& format)
    {
        country_code.clear();
        operator_code.clear();
        personal_number.clear();
        name.clear();

        if (format.empty())
            return;

        auto it = format.begin();

        /// country code
        ++it;   // skip plus sign
        for (; *it != ' '; ++it)
            country_code.push_back(*it);

        /// operator code
        it += 2; // skip " ("
        for (; *it != ')'; ++it)
            operator_code.push_back(*it);

        /// personal number
        it += 2; // skip ") "
        for (; *it != ' '; ++it)
            personal_number.push_back(*it);

        /// name
        it += 3; // skip " - "
        for (; it != format.end(); ++it)
            name.push_back(*it);
    }

    bool match(const std::string& number) const
    {
        int count = country_code.size()
                    + operator_code.size()
                    + personal_number.size();

        if (count != number.size()) // check number and format length
            return false;

        if (number.compare(0, country_code.size(), country_code) != 0) // compare numbers country code
            return false;

        if (number.compare(country_code.size(), operator_code.size(), operator_code) != 0) // compare numbers operator code
            return false;

        for (size_t i = 0; i < personal_number.size(); i++)
        {
            if (personal_number[i] == 'X')
                continue;
            else if (personal_number[i] != number[country_code.size() + operator_code.size() + i])
                    return false;
        }
        return true;
    }

    std::string get_formatted_number(const std::string& number)
    {
        std::string formatted;
        formatted.append('+' + country_code + " (" + operator_code + ") "
                         + number.substr(operator_code.size() + country_code.size()) + " - " + name);
        return formatted;
    }
};

int main()
{
    size_t n = 0;                       // for storing total number of numbers/formats
    std::vector<std::string> numbers;   // vector for storing numbers
    std::vector<std::string> formatted; // vector for storing formatted numbers
    std::vector<NumberFormat> formats;  // vector of format types
    std::string tmp;                    // temporary string for storing input

    // get total amount of phone numbers
    std::cin >> n;
    std::getchar();

    for (size_t i = 0; i < n; i++)
    {
        std::getline(std::cin, tmp);
        numbers.push_back(split(tmp)); // get rid of unnecessary symbols and push in the vector
    }

    // get total amount of formats
    std::cin >> n;
    std::getchar();

    for (size_t i = 0; i < n; i++)
    {
        std::getline(std::cin, tmp);
        formats.emplace_back(NumberFormat(tmp)); // push format in vector
    }

    for (const auto& num : numbers)     // for each number
        for (auto& form : formats)      // for each format
            if (form.match(num))        // if number fits format, then get formatted number and push it in vector
                formatted.emplace_back(form.get_formatted_number(num));


    std::cout << std::endl;
    for (const auto& p  :formatted)
        std::cout << p << std::endl;

    return 0;
}
*/
