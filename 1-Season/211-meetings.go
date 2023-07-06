package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Time struct {
	Day           uint
	TimeStr       string
	Duration      uint
	Start         uint
	AbsoluteStart uint
	AbsoluteEnd   uint
}

func NewTime(day uint, timeStr string, duration uint) Time {
	hours, _ := strconv.Atoi(timeStr[0:2])
	minutes, _ := strconv.Atoi(timeStr[3:5])
	start := uint(hours*60 + minutes)
	absoluteStart := day*24*60 + start
	absoluteEnd := absoluteStart + duration

	return Time{
		Day:           day,
		TimeStr:       timeStr,
		Duration:      duration,
		Start:         start,
		AbsoluteStart: absoluteStart,
		AbsoluteEnd:   absoluteEnd,
	}
}

func (t Time) TimeStamp() string {
	return t.TimeStr + " " + strconv.Itoa(int(t.Duration))
}

type Person struct {
	Meetings map[Time][]string
}

func NewPerson() Person {
	return Person{
		Meetings: make(map[Time][]string),
	}
}

func (p Person) IsBusy(time Time) bool {
	for t := range p.Meetings {
		if t.Day == time.Day && ((t.Start >= time.Start && t.Start < time.Start+time.Duration) || (time.Start >= t.Start && time.Start < t.Start+t.Duration)) {
			return true
		}
	}
	return false
}

func (p Person) Appoint(time Time, members []string) {
	if !p.IsBusy(time) {
		p.Meetings[time] = members
	}
}

func (p Person) GetMeetings(day uint) []Meeting {
	var meetingsThatDay []Meeting
	for t, members := range p.Meetings {
		if t.Day == day {
			meetingsThatDay = append(meetingsThatDay, Meeting{Time: t, Members: members})
		}
	}
	sort.SliceStable(meetingsThatDay, func(i, j int) bool {
		return meetingsThatDay[i].Time.Start < meetingsThatDay[j].Time.Start
	})
	return meetingsThatDay
}

type Meeting struct {
	Time    Time
	Members []string
}

// Made with ChatCBD
// CBD - Here's the equivalent code in Go:
// Den  - Please check code again thereis assigment into nil map
// Den  - Also thereis no TimeStamp() method
// CBD  - Apologies for the oversight. Here's the corrected code with the TimeStamp method added to the Time struct:
//
// Den  -You don't call NewPerson()
// CBD -  ...
// Den  - Program worls, however it fails on testcase
// CBD -  ...
// Den - Program works however the output sometimes is unordered because of the map, can you fix it?
// CBD - Certainly! To ensure that the output is ordered, you can modify the GetMeetings method of the Person struct to return the meetings in a sorted order based on the start time of each meeting. Here's the updated code:

func main() {
	var inputLength uint
	fmt.Scan(&inputLength)
	var output strings.Builder

	people := make(map[string]Person)

	for i := uint(0); i < inputLength; i++ {
		var command string
		fmt.Scan(&command)

		if command == "APPOINT" {
			var day uint
			var timeStr string
			var duration uint
			fmt.Scan(&day, &timeStr, &duration)
			time := NewTime(day, timeStr, duration)

			var peopleCount uint
			fmt.Scan(&peopleCount)
			var members []string
			var fails []string
			for j := uint(0); j < peopleCount; j++ {
				var member string
				fmt.Scan(&member)

				p, found := people[member]
				if found && p.IsBusy(time) {
					fails = append(fails, member)
				}

				members = append(members, member)
			}

			if len(fails) == 0 {
				output.WriteString("OK\n")
				for _, member := range members {
					p, found := people[member]
					if !found {
						p = NewPerson()
					}
					p.Appoint(time, members)
					people[member] = p
				}
			} else {
				output.WriteString("FAIL\n")
				output.WriteString(strings.Join(fails, " "))
				output.WriteString("\n")
			}

		} else if command == "PRINT" {
			var day uint
			var member string
			fmt.Scan(&day, &member)

			var meetings []Meeting
			p, found := people[member]
			if found {
				meetings = p.GetMeetings(day)
			}

			for _, meeting := range meetings {
				output.WriteString(meeting.Time.TimeStamp())
				for _, person := range meeting.Members {
					output.WriteString(" ")
					output.WriteString(person)
				}
				output.WriteString("\n")
			}
		}
	}

	fmt.Print(output.String())
}
