package main

import (
	"fmt"
	"sort"
	"strings"
	"unsafe"
)

type Note struct {
	warehouseID   byte
	comma1        byte
	dateFirst     [10]byte
	space         byte
	dateLast      [10]byte
	comma2        byte
	commodityType [3]byte
}

func findAnswer(v []string) string {
	if len(v) == 0 {
		return ""
	}

	replaceLater := make(map[string]string)
	compareNotes := func(s1, s2 string) bool {
		if s1[0] < s2[0] {
			return true
		} else if s1[0] > s2[0] {
			return false
		}

		note1 := (*Note)(unsafe.Pointer(&([]byte(s1)[0])))
		note2 := (*Note)(unsafe.Pointer(&([]byte(s2)[0])))

		const (
			KGT   = 'K'
			COLD  = 'C'
			OTHER = 'O'
		)

		// KGT < COLD < OTHER
		if note2.commodityType[0] == KGT && note1.commodityType[0] != KGT {
			return false
		}
		if note2.commodityType[0] == COLD && note1.commodityType[0] == OTHER {
			return false
		}
		if note2.commodityType[0] == OTHER && note1.commodityType[0] != OTHER {
			return true
		}
		if note2.commodityType[0] == COLD && note1.commodityType[0] == KGT {
			return true
		}

		// If there is no intersection of time intervals:
		if strings.Compare(string(note1.dateFirst[:]), string(note2.dateLast[:])) > 0 { // first1 > last2
			return false
		}
		if strings.Compare(string(note1.dateLast[:]), string(note2.dateFirst[:])) < 0 { // last1 < first2
			return true
		}

		// If there is an intersection of time intervals:
		s3 := s1
		note3 := (*Note)(unsafe.Pointer(&([]byte(s3)[0])))

		if strings.Compare(string(note1.dateFirst[:]), string(note2.dateFirst[:])) < 0 { // first1 < first2
			copy(note3.dateFirst[:], note1.dateFirst[:])
		} else {
			copy(note3.dateFirst[:], note2.dateFirst[:])
		}

		if strings.Compare(string(note1.dateLast[:]), string(note2.dateLast[:])) > 0 { // last1 > last2
			copy(note3.dateLast[:], note1.dateLast[:])
		} else {
			copy(note3.dateLast[:], note2.dateLast[:])
		}

		replaceLater[s1] = s3
		replaceLater[s2] = ""

		return true
	}

	sort.Slice(v, func(i, j int) bool {
		return compareNotes(v[i], v[j])
	})

	// Remove extra elements
	for from, to := range replaceLater {
		for i := 0; i < len(v); i++ {
			v[i] = strings.Replace(v[i], from, to, -1)
		}
	}
	v = removeEmptyStrings(v)

	// Prepare results
	var out strings.Builder
	for _, s := range v {
		out.WriteString(s + "\n")
	}
	return out.String()
}

func removeEmptyStrings(s []string) []string {
	result := make([]string, 0, len(s))
	for _, str := range s {
		if str != "" {
			result = append(result, str)
		}
	}
	return result
}

func appendElements(v *[]string, s string) {
	if len(s) < 4 {
		return
	}

	nullString := "NULL"
	if !strings.HasSuffix(s, nullString) {
		*v = append(*v, s)
		return
	}

	subString := s[:len(s)-4]
	*v = append(*v, subString+"KGT", subString+"COLD", subString+"OTHER")
}

func main() {
	in := make([]string, 0)
	var buf string
	for {
		_, err := fmt.Scanln(&buf)
		if err != nil {
			break
		}

		for buf[len(buf)-1] == '\r' || buf[len(buf)-1] == '\n' {
			buf = buf[:len(buf)-1]
		}
		if len(buf) < 27 {
			continue
		}

		appendElements(&in, buf)
	}

	fmt.Println(findAnswer(in))
}

/*
#include <algorithm>
#include <cstring>
#include <string>
#include <iostream>
#include <vector>
#include <unordered_map>

using namespace std;

struct Note {
    char warehouse_id       = '0';
    char comma1             = ',';
    char date_first[10]     = {'0','0','0','0','-','0','0','-','0','0'};
    char space              = ' ';
    char date_last[10]      = {'9','9','9','9','-','9','9','-','9','9'};
    char comma2             = ',';
    char commodity_type[3]  = {'K','G','T'};
}__attribute__((packed));

string findAnswer(vector<string>& v)
{
    if (v.empty())
        return string();

    unordered_map<string,string> replaceLater;
    auto compare_notes = [&replaceLater](const string &s1, const string &s2) {
        if (s1.front() < s2.front())
            return true;
        else if (s1.front() > s2.front())
            return false;

        const Note *note1 = reinterpret_cast<const Note*>(s1.data());
        const Note *note2 = reinterpret_cast<const Note*>(s2.data());

        static const char KGT   = 'K';
        static const char COLD  = 'C';
        static const char OTHER = 'O';

        // KGT < COLD < OTHER
        if (note2->commodity_type[0] == KGT   && note1->commodity_type[0] != KGT)
            return false;
        if (note2->commodity_type[0] == COLD  && note1->commodity_type[0] == OTHER)
            return false;
        if (note2->commodity_type[0] == OTHER && note1->commodity_type[0] != OTHER)
            return true;
        if (note2->commodity_type[0] == COLD  && note1->commodity_type[0] == KGT)
            return true;

        // If there is no intersection of time intervals:
        if (strncmp(note1->date_first, note2->date_last,  10) > 0) // first1 > last2
            return false;
        if (strncmp(note1->date_last,  note2->date_first, 10) < 0) // last1  < first2
            return true;

        // If there is intersection of time intervals:
        string s3 = s1;
        Note *note3 = reinterpret_cast<Note*>(s3.data());

        if (strncmp(note1->date_first, note2->date_first, 10) < 0) // first1 < first2
            strncpy(note3->date_first, note1->date_first, 10);
        else
            strncpy(note3->date_first, note2->date_first, 10);

        if (strncmp(note1->date_last,  note2->date_last,  10) > 0) // last1  > first2
            strncpy(note3->date_last, note1->date_last, 10);
        else
            strncpy(note3->date_last, note2->date_last, 10);

        replaceLater[s1] = s3;
        replaceLater[s2] = string();

        return true;
    };

    // sort notes
    sort(begin(v), end(v), compare_notes);

    // remove extra elements
    for (auto&& [from, to] : replaceLater) {
        replace(begin(v), end(v), from, to);
    }
    (void)remove(begin(v), end(v), string());
    while (v.back() == string()) {
        v.resize(v.size()-1);
    }

    // prepare results
    string out;
    for (const string &s : v) {
        out.append(s + "\n");
    }
    return out;
}

void appendElements(vector<string>& v, const string& s)
{
    if (s.size() < 4)
        return;

    static const string nullString = "NULL";
    if (s.substr(s.size()-4) != nullString) {
        v.push_back(s);
        return;
    }

    const string subString = s.substr(0, s.size()-4);
    v.push_back(subString + "KGT");
    v.push_back(subString + "COLD");
    v.push_back(subString + "OTHER");
}

int main()
{
    vector<string> in;
    string buf;
    while (!cin.eof()) {
        getline(cin, buf);

        while (buf.back() == '\r' || buf.back() == '\n')
            buf.resize(buf.size()-1);
        if (buf.size() < 27)
            continue;

        appendElements(in, buf);
    }

    cout << findAnswer(in) << endl;

    findAnswer(in);

    return 0;
}
*/
