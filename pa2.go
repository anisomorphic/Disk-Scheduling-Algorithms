// Michael Harris
// mi051467
// COP4600 pa2 - disk scheduling algorithms - 1.0

package main

//imports
import "os"
import "fmt"
import "math"
import "bufio"
import "strconv"

type req_a struct { //REQuest_Accessed, used in sstf
	distance int
	accessed bool
}

//constants
const MaxInt = int(^uint(0) >> 1)
const invalidModeError string = "ERROR: Valid modes: fcfs, sstf, scan, look, c-scan, c-look"
const error11 string = "ABORT(11):initial (%d) > upper (%d)\n"
const error12 string = "ABORT(12):initial (%d) < lower (%d)\n"
const error13 string = "ABORT(13):upper (%d) > lower (%d)\n"
const error15 string = "ERROR(15):Request out of bounds:	req (%d) > upper (%d) or < lower (%d)\n"


//functions

//floor, used in SCAN
func floor(array []int, curPos int) (int) {
	temp := 0

	for i := 0; i < len(array); i++ {
		if array[i] > curPos {
			return i
		}
	}

	return temp
}

//bubble sort an int array
func sort(array []int) {

	swapped := true;
	for swapped {
		swapped = false
		for i := 0; i < len(array) - 1; i++ {
			if array[i + 1] < array[i]{
				swap(array, i, i + 1)
				swapped = true
			}
		}
	}
}

//swap two array positions, used in bubble sorting methods [sort_(.*)]
func swap(array []int, i, j int) {
	tmp := array[j]
	array[j] = array[i]
	array[i] = tmp
}

// precondition: two arrays representing the same data set, one of which is []int
//							 representing absolute cylinder position, and the other []req_a,
//							 representing distance and accessed status, and the head's position.
// postcondition: update the distance component of the []req_a array based on head
func difference(a1 []int, a2 []req_a, head int) {
	for i := 0; i < len(a1); i++ {
		a2[i].distance = int(math.Abs(float64(head) - float64(a1[i])))
	}
}

// precondition: pass two signed integers
// postcondition: return the positive difference or 0 if both values are the same
func diff(a int, b int) (int) {
	if a > b {
		return a-b
	}
	if b > a {
		return b-a
	}
	return 0
}

// precondition: pass one of the 6 valid integer representation of a mode
// postcondition: return the uppercase string representation to display
func getMode(mode int) (string) {
	if mode == 1 {
		return "FCFS"
	}
	if mode == 2 {
		return "SSTF"
	}
	if mode == 3 {
		return "SCAN"
	}
	if mode == 4 {
		return "C-SCAN"
	}
	if mode == 5 {
		return "LOOK"
	}
	if mode == 6 {
		return "C-LOOK"
	}
	fmt.Println("\t" + invalidModeError + " (in getMode)")
	panic(mode)
}


//main
func main() {

	//make sure we have command line arguments
	if len(os.Args) < 2 {
		fmt.Println("Please provide a valid input file as the first parameter.")
		return
	}

	//file management and produce errors if needed
	file, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println("File is invalid: ", os.Args[1])
		fmt.Print("\n\t"); panic(err)
	  return
	}
	defer file.Close()

	//scan input file by word
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	//variables for main
	cmpStr := ""
	mode := -1
	upperCYL  := -1
	lowerCYL := -1
	initCYL := -1
	traversal := 0
	var cylreqs []int

	//start handling input file
	for scanner.Scan() {
		inStr := scanner.Text()
		if inStr == "end" {
			break;
		}

		if cmpStr == "use" && mode == -1 {
			if inStr == "fcfs" {
				mode = 1
			} else if inStr == "sstf" {
				mode = 2
			} else if inStr == "scan" {
				mode = 3
			} else if inStr == "c-scan" {
				mode = 4
			} else if inStr == "look" {
				mode = 5
			} else if inStr == "c-look" {
				mode = 6
			} else {
				fmt.Println("\t" + invalidModeError + " (in main)")
				panic(inStr)
			}
		}
		if cmpStr == "lowerCYL" && lowerCYL == -1 {
			lowerCYL, _ = strconv.Atoi(inStr)
		}
		if cmpStr == "upperCYL" && upperCYL == -1 {
			upperCYL, _ = strconv.Atoi(inStr)
			if upperCYL < lowerCYL {
				fmt.Printf(error13, upperCYL, lowerCYL)
				os.Exit(0)
			}
		}
		if cmpStr == "initCYL" && initCYL == -1 {
			initCYL, _ = strconv.Atoi(inStr)
			if initCYL > upperCYL {
				fmt.Printf(error11, initCYL, upperCYL)
				os.Exit(0)
			}
			if initCYL < lowerCYL {
				fmt.Printf(error12, initCYL, lowerCYL)
				os.Exit(0)
			}
		}
		if cmpStr == "cylreq" {
			tempVal, _ := strconv.Atoi(scanner.Text())

			//error checking
			if ((tempVal > upperCYL) || (tempVal < lowerCYL)) {
				fmt.Printf(error15, tempVal, upperCYL, lowerCYL)
			} else { //only add a cylinder if it's valid
				cylreqs = append(cylreqs, tempVal)
			}


		}

		//storing previous word in order to parse
		cmpStr = scanner.Text()
	} //stop processing input


	//safety checks (abort) ~ moved to input processing per error.base files



	// show input info to screen
	fmt.Print("Seek algorithm: " + getMode(mode) + "\n")
	fmt.Printf("\tLower cylinder: %5d\n", lowerCYL)
	fmt.Printf("\tUpper cylinder: %5d\n", upperCYL)
	fmt.Printf("\tInit cylinder:  %5d\n", initCYL)
	fmt.Println("\tCylinder requests:")
	for i := 0; i < len(cylreqs); i++ {
		fmt.Printf("\t\tCylinder %5d\n", cylreqs[i])
	}


	// run algorithm
	if mode == 1 { //FCFS
		curPos := initCYL

		// loop through each request
		for i := 0; i < len(cylreqs); i++ {
			req := cylreqs[i]
			fmt.Printf("Servicing %5d\n", req)
			traversal += diff(req, curPos)
			curPos = req
		} //loop block
	} //end mode1

	if mode == 2 { //SSTF
		curPos := initCYL

		// initialize and prepare an array to hold the absolute distance between current head
		var requests []req_a
		for i := 0; i < len(cylreqs); i++ {
			var tempReq req_a
			tempReq.distance = int(math.Abs(float64(curPos) - float64(cylreqs[i])))
			tempReq.accessed = false
			requests = append(requests, tempReq)
		}

		// main execution loop
		for i := 0; i < len(requests); i++ {
			indexToService := -1
			trackClosest := MaxInt

			// update distances based on current head position
			difference(cylreqs, requests, curPos)

			// find closest cylinder index
			for j := 0; j < len(requests); j++ {
				if (trackClosest > requests[j].distance) && (!requests[j].accessed) {
					indexToService = j
					trackClosest = requests[j].distance
				}
			} //j

			// service the sector
			fmt.Printf("Servicing %5d\n", cylreqs[indexToService])
			traversal += requests[indexToService].distance
			requests[indexToService].accessed = true
			curPos = cylreqs[indexToService]
		} //i
	} //end mode2

	if mode == 3 { //SCAN
		curPos := initCYL

		sort(cylreqs)
		startLocation := floor(cylreqs, curPos)

		for i := startLocation; i < len(cylreqs); i++ {
			fmt.Printf("Servicing %5d\n", cylreqs[i])
			traversal += diff(curPos, cylreqs[i])
			curPos = cylreqs[i]
		}

		if startLocation != 0 {
			//fmt.Println("need to do more")

			traversal += upperCYL - curPos
			traversal += upperCYL - initCYL
			curPos = initCYL

			for i := startLocation - 1; i > -1; i-- {
				fmt.Printf("Servicing %5d\n", cylreqs[i])
				traversal += diff(curPos, cylreqs[i])
				curPos = cylreqs[i]
			}
		}
	} //end mode3

	if mode == 4 { //C-SCAN
		curPos := initCYL

		sort(cylreqs)
		startLocation := floor(cylreqs, curPos)

		for i := startLocation; i < len(cylreqs); i++ {
			fmt.Printf("Servicing %5d\n", cylreqs[i])
			traversal += diff(curPos, cylreqs[i])
			curPos = cylreqs[i]
		}

		if startLocation != 0 {
			//fmt.Println("need to do more")

			traversal += upperCYL - curPos
			traversal += upperCYL - lowerCYL
			curPos = lowerCYL

			for i := 0; i < startLocation; i++ {
				fmt.Printf("Servicing %5d\n", cylreqs[i])
				traversal += diff(curPos, cylreqs[i])
				curPos = cylreqs[i]
			}
		}
	} //end mode4

	if mode == 5 { //LOOK
		curPos := initCYL

		sort(cylreqs)
		startLocation := floor(cylreqs, curPos)

		for i := startLocation; i < len(cylreqs); i++ {
			fmt.Printf("Servicing %5d\n", cylreqs[i])
			traversal += diff(curPos, cylreqs[i])
			curPos = cylreqs[i]
		}

		if startLocation != 0 {
			//fmt.Println("need to do more")

			traversal += curPos - initCYL
			curPos = initCYL

			//move down from initCYL
			for i := startLocation - 1; i > -1; i-- {
				fmt.Printf("Servicing %5d\n", cylreqs[i])
				traversal += diff(curPos, cylreqs[i])
				curPos = cylreqs[i]
			}
		}
	} //end mode5

	if mode == 6 { //C-LOOK
		curPos := initCYL

		sort(cylreqs)
		startLocation := floor(cylreqs, curPos)

		for i := startLocation; i < len(cylreqs); i++ {
			fmt.Printf("Servicing %5d\n", cylreqs[i])
			traversal += diff(curPos, cylreqs[i])
			curPos = cylreqs[i]
		}

		if startLocation != 0 {
			//fmt.Println("need to do more")

			traversal += curPos - cylreqs[0]
			curPos = cylreqs[0]

			for i := 0; i < startLocation; i++ {
				fmt.Printf("Servicing %5d\n", cylreqs[i])
				traversal += diff(curPos, cylreqs[i])
				curPos = cylreqs[i]
			}
		}
	} //end mode6

	// show output info to screen
	fmt.Printf(getMode(mode) + " traversal count = %5d\n", traversal)

} //end main
