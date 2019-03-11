package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// HourGlass Sum method
// returns the Sum
// Algorithm used: first, make a matrix with length - 2 of original
// Use the leftmost index of HourGlass as starting point
// Iterate over original matrix excluding 2+ index up and down so
// we don't get out of bounds
// hardcode format of HourGlass that will be iterated
// Stored results for each sum in new matrix
// find max value in new matrix and return it
func hourglassSum(arr [][]int32) int32 {
    newArr := make([][]int32, len(arr)-2)
    for i := 0; i < len(arr[0])-2; i++ {
        newArr[i] = make([]int32, len(arr)-2)
    }
    for i := 0; i < len(arr[0])-2; i++ {
        for j := 0; j < len(arr[0])-2; j++{
            newArr[i][j] = arr[i][j]+arr[i][j+1]+arr[i][j+2]+arr[i+1][j+1]+arr[i+2][j]+arr[i+2][j+1]+arr[i+2][j+2]
        }
    }

    var max int32
    max = -200000
    for i := 0; i<len(newArr[0]); i++{
        for j:=0; j<len(newArr[0]); j++{
            if max < newArr[i][j] {
                max = newArr[i][j]
            }
        }
    }
    return max
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	// Write result to file
	outfile := os.Getenv("OUTPUT_PATH")+"/hourglassArr.txt"
    stdout, err := os.Create(outfile)
    checkError(err)

	// close only when main function is no longer running
    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	// Read input from command line and place in arr
    var arr [][]int32
    for i := 0; i < 6; i++ {
        arrRowTemp := strings.Split(readLine(reader), " ")

        var arrRow []int32
        for _, arrRowItem := range arrRowTemp {
            arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
            checkError(err)
            arrItem := int32(arrItemTemp)
            arrRow = append(arrRow, arrItem)
        }

        if len(arrRow) != int(6) {
            panic("Bad input")
        }

        arr = append(arr, arrRow)
    }

    result := hourglassSum(arr)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

// Helper to read by line and trim excess
func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

// Error Tester
func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
