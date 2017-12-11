package main

import (
	"fmt"
)

func main() {
	limit := 277678

	vals := []int{1}

	sqSize := 1
	sqStart := 0
	sidePos := 0

	diff := 1

	for i := 1;; i++ {

		added := map[int]bool{}
		sum := 0

		sum += help(added, vals, i-1)

		vals = append(vals, 0)

		//new square
		if i == sqSize*sqSize {
			sqSize += 2
			sqStart = i
			sum += help(added, vals, i-(diff-1))
		} else {
			sidePos = (i - sqStart) % (sqSize - 1)

			if i-sqStart == 1 {
				sum += help(added, vals, i-2)

				sum += help(added, vals, i-(diff-1))
				sum += help(added, vals, i-(diff))
			} else if sidePos > 0 {
				sum += help(added, vals, i-(diff+1))
			}

			if sqSize*sqSize-2 == i {
				sum += help(added, vals, i-(diff-1))
				sum += help(added, vals, i-(diff))
				sum += help(added, vals, i-(diff+1))
			}
			if sqSize*sqSize-1 == i {
				sum += help(added, vals, i-(diff))
				sum += help(added, vals, i-(diff+1))
			}

			if sidePos == 0 {
				sum += help(added, vals, i-2)
			}

			if sidePos < sqSize-3 {
				sum += help(added, vals, i-(diff-1))
			}

			if sidePos < sqSize-2 {
				sum += help(added, vals, i-(diff))
			}



			if (i-sqStart)%(sqSize-1) == sqSize-2 {
				diff += 2
			}
		}

		fmt.Println(i, "i", "sidePos", sidePos, "diff", diff, "sqSize", sqSize, "sidepos", sidePos, "sum", sum)

		vals[i] = sum
		if sum > limit {
			fmt.Println("Found", sum)
			break
		}

	}

	fmt.Println(vals)

}

func help(added map[int]bool, vals []int, index int) int {
	if index >= 0 && index < len(vals) && !added[index] {
		added[index] = true
		return vals[index]
	} else {
		return 0
	}
}

//36 35	34	33	32	31	30
//37 16 15  14  13  12	29
//38 17  4   3   2  11	28
//39 18  5   0   1  10	27
//40 19  6   7   8  9	26
//41 20 21  22	23	24	25
//42 43	44	45	46	47	48

//1: -1     |
//2: -1     |        -2
//
//3: -1, -2 |        -3
//4: -1     |     -4
//
//5: -1, -2 |     -5
//6: -1     | -6
//
//7: -1, -2 |     -7, -6
//8: -1     | -8, -7, -6
//
//
//
//
//
//
//25: -1     |        *, -16
//26: -1, -2 |   *, -17, -16
//27: -1     | -18, -17, -16
//28: -1     | -18, -17, -16
//29: -1     | -18, -17
//30: -1     | -18
//
//31: -1, -2 |      -19, -18
//32: -1     | -20, -19, -18
//33: -1     | -20, -19, -18
//34: -1     | -20, -19, -18
//35: -1     | -20, -19
//36: -1     | -20
//
//37: -1, -2 |      -21, -20
//38: -1     | -22, -21, -20
//39: -1     | -22, -21, -20
//40: -1     | -22, -21, -20
//41: -1     | -22, -21
//42: -1     | -22
//
//43: -1, -2 |      -23, -22
//44: -1     | -24, -23, -22
//45: -1     | -24, -23, -22
//46: -1     | -24, -23, -22
//47: -1     | -24, -23, -22*
//48: -1     | -24, -23*

//10: -1,-2,-9,-8
//11: -1,-10,-9
//12: -1,-10
//
//13: -1,-2,-11
//14: -1, -12,-11,-10
//15: -1, -12,-11
//16: -1, -12
//
//17: -1, -2, -13,-12
//18: -1, -14,-13,-12
//19: -1, -14,-13
//20: -1, -14
//
//21: -1,-2,-15,-14
//22: -1, -16,-15,-14
//23: -1, -16,-15,-14
//24: -1,-16,-15
//
//25: -1,-16
