package main

import (
	"fmt"
	"strings"
)

const (
	WaitingInLine = iota
	StowingLuggage
	WaitingForSeat
)

type Passenger struct {
	Row, Seat, Time, Status, Position int64
	Side                              bool
}

// 380. Посадка в самолет
func main() {
	var n int
	_, err := fmt.Scan(&n)
	if err != nil {
		panic(err)
	}
	passengers := make([]Passenger, n)
	plane := make([][]bool, 30)
	for i := range plane {
		plane[i] = make([]bool, 7)
	}
	for i := 0; i < n; i++ {
		var row, time int64
		var seat string
		_, err := fmt.Scan(&time, &row, &seat)
		if err != nil {
			panic(err)
		}
		seatInt := strings.Index("ABCDEF", seat)
		side := false
		if seatInt > 2 {
			side = true
			seatInt = 5 - seatInt
		}
		passengers[i] = Passenger{row - 1, int64(seatInt), time, WaitingInLine, int64(-i), side}
	}
	var time int64 = 0
	for len(passengers) > 0 {
		//println("Time:", time)
		for i := 0; i < len(passengers); i++ {
			//println("Position:", passengers[i].Position, "Row:", passengers[i].Row, "Seat:", passengers[i].Seat, "Ptime:", passengers[i].Time, "Status:", passengers[i].Status, "time:", time)
			seat := passengers[i].Seat
			neighbour := passengers[i].Seat + 1
			secondNeighbour := passengers[i].Seat + 2
			if passengers[i].Side {
				seat = 6 - seat
				neighbour = seat - 1
				secondNeighbour = seat - 2
			}
			switch passengers[i].Status {
			case WaitingInLine:
				if passengers[i].Position < passengers[i].Row && (passengers[i].Position+1 < int64(len(plane)) &&
					passengers[i].Position+1 > 0 &&
					!plane[passengers[i].Position+1][3]) {
					plane[passengers[i].Position][3] = false
					plane[passengers[i].Position+1][3] = true
					passengers[i].Position++
				} else if passengers[i].Position == passengers[i].Row {
					if (seat == 2 || seat == 4 || ((seat == 1 || seat == 5) && !plane[passengers[i].Row][neighbour]) || ((seat == 0 || seat == 6) && !plane[passengers[i].Row][neighbour] && !plane[passengers[i].Row][secondNeighbour])) && passengers[i].Time == 0 {
						plane[passengers[i].Row][seat] = true
						plane[passengers[i].Position][3] = false
						passengers = append(passengers[:i], passengers[i+1:]...)
						i--
					} else {
						plane[passengers[i].Position][3] = true
						passengers[i].Status = StowingLuggage
					}
				} else if passengers[i].Position+1 <= 0 && !plane[0][3] {
					passengers[i].Position++
					if i == len(passengers)-1 {
						plane[0][3] = true
					}
				}
			case StowingLuggage:
				passengers[i].Time--
				if passengers[i].Time == 0 {
					if (seat == 2 || seat == 4 || ((seat == 1 || seat == 5) && !plane[passengers[i].Row][neighbour]) || ((seat == 0 || seat == 6) && !plane[passengers[i].Row][neighbour] && !plane[passengers[i].Row][secondNeighbour])) && passengers[i].Time == 0 {
						plane[passengers[i].Row][seat] = true
						plane[passengers[i].Position][3] = false
						passengers = append(passengers[:i], passengers[i+1:]...)
						i--
					} else {
						passengers[i].Status = WaitingForSeat
						if seat == 0 || seat == 6 {
							if plane[passengers[i].Row][neighbour] {
								passengers[i].Time += 5
							}
							if plane[passengers[i].Row][secondNeighbour] {
								passengers[i].Time += 5
								if plane[passengers[i].Row][neighbour] {
									passengers[i].Time += 5
								}
							}
						} else if seat == 1 || seat == 5 {
							if plane[passengers[i].Row][neighbour] {
								passengers[i].Time += 5
							}
						}
					}
				} else {
					passengers[i].Status = WaitingForSeat
					if seat == 0 || seat == 6 {
						if plane[passengers[i].Row][neighbour] {
							passengers[i].Time += 5
						}
						if plane[passengers[i].Row][secondNeighbour] {
							passengers[i].Time += 5
							if plane[passengers[i].Row][neighbour] {
								passengers[i].Time += 5
							}
						}
					} else if seat == 1 || seat == 5 {
						if plane[passengers[i].Row][neighbour] {
							passengers[i].Time += 5
						}
					}
				}
			case WaitingForSeat:
				passengers[i].Time--
				if passengers[i].Time <= 0 {
					plane[passengers[i].Row][seat] = true
					plane[passengers[i].Position][3] = false
					passengers = append(passengers[:i], passengers[i+1:]...)
					i--
				}
			}
		}
		time++
	}
	fmt.Println(time)
}
