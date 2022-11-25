package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "strconv"
    "sort"
)

var square_feet int

func main() {
    file_path := os.Args[1]
    readFile, err := os.Open(file_path)
    if err != nil {
        fmt.Println("Error opening file!!!")
    }

    fileScanner := bufio.NewScanner(readFile)
    fileScanner.Split(bufio.ScanLines)
    var fileLines []string

    for fileScanner.Scan() {
        fileLines = append(fileLines, fileScanner.Text())
    }

    readFile.Close()

    for _, line := range fileLines {
        // fmt.Println(line)
	// total, smallest_side := process_line(line)
	total := process_line(line)
	square_feet += total
	// fmt.Println(line, total)
    }
    fmt.Println("Square Feet", square_feet)
}

func process_line(full_line string) int {
    line_split := strings.Split(full_line, "x")

    h, _ := strconv.Atoi(line_split[0])
    w, _ := strconv.Atoi(line_split[1])
    l, _ := strconv.Atoi(line_split[2])

    side1 := h * w
    side2 := w * l
    side3 := l * h

    smallest_side := find_smallest(side1, side2, side3)

    return (2 * side1) + (2 * side2) + (2 * side3) + smallest_side
}

func find_smallest(side1 int, side2 int, side3 int) int {
    fmt.Println(side1, side2, side3)
    //var smallest int
    //smallest = -999

    intSlice := []int{side1, side2, side3}
    sort.Ints(intSlice)
    /*    if side1 <= side2 {
              if side2 <= side3 {
                  fmt.Println("Beep")
                  smallest = side1
      	} else if side2 >= side3 {
                  fmt.Println("Beep")
                  smallest = side2
              }
          } else if side1 >= side2 && side2 >= side3 {
              fmt.Println("Boop")
              smallest = side3
          } else if side1 >= side2 && side2 <= side3 {
              fmt.Println("Baap")
              smallest = side2
        }*/
    return intSlice[0]
}

