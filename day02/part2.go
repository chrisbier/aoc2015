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
var ribbon int

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
	    total, ribbon_temp := process_line(line)
	    square_feet += total
	    ribbon += ribbon_temp
	    fmt.Println(line, total, ribbon)
    }
    fmt.Println("Square Feet", square_feet, "Ribbon", ribbon)
}

func process_line(full_line string) (int, int)  {
    line_split := strings.Split(full_line, "x")

    h, _ := strconv.Atoi(line_split[0])
    w, _ := strconv.Atoi(line_split[1])
    l, _ := strconv.Atoi(line_split[2])

    side1 := h * w
    side2 := w * l
    side3 := l * h

    sorted_list := find_smallest(side1, side2, side3)

    ribbon_len := find_ribbon([]int{h, w, l})

    return (2 * side1) + (2 * side2) + (2 * side3) + sorted_list[0], ribbon_len
}

func find_smallest(side1 int, side2 int, side3 int) []int {
    // fmt.Println(side1, side2, side3)

    intSlice := []int{side1, side2, side3}
    sort.Ints(intSlice)
    return intSlice
}

func find_ribbon(sorted_list[]int) int {
    sort.Ints(sorted_list)
    ribbon_len := (sorted_list[0] * 2) + (sorted_list[1] * 2)
    bow_len := sorted_list[0] * sorted_list[1] * sorted_list[2]

    fmt.Println(sorted_list, ribbon_len, bow_len)
    return ribbon_len + bow_len
}

