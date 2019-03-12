package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

func arrayManipulation(a int, b int, k int32, arrSlope *[]int64) {
		slice := *arrSlope
		slice[a-1] += int64(k)
		if b < len(slice) {
			slice[b] -= int64(k)
		}
}

func sumToMax(arr []int64) int64 {
	var sum, maxV int64
	sum = 0
	maxV = 0
	for i := 0; i < len(arr); i++ {
		sum += arr[i]
		maxV = findMax(maxV, sum)
	}
	return maxV
}

func findMax(a, b int64) int64 {
	if a > b {
		return a
	} else {
		return b
	}
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	outfile := os.Getenv("OUTPUT_PATH")+"/arrayManipulation.txt"
    stdout, err := os.Create(outfile)
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nm := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nm[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    mTemp, err := strconv.ParseInt(nm[1], 10, 64)
    checkError(err)
    m := int32(mTemp)

	arrSlope := make([]int64, n)
    for i := 0; i < int(m); i++ {
        queriesRowTemp := strings.Split(readLine(reader), " ")

        var queriesRow []int32
        for _, queriesRowItem := range queriesRowTemp {
            queriesItemTemp, err := strconv.ParseInt(queriesRowItem, 10, 64)
            checkError(err)
            queriesItem := int32(queriesItemTemp)
            queriesRow = append(queriesRow, queriesItem)
        }

        if len(queriesRow) != int(3) {
            panic("Bad input")
        }

		arrayManipulation(int(queriesRow[0]), int(queriesRow[1]), queriesRow[2], &arrSlope)
    }

    result := sumToMax(arrSlope)

    fmt.Fprintf(writer, "%d\n", result)

    writer.Flush()
}

func readLine(reader *bufio.Reader) string {
    str, _, err := reader.ReadLine()
    if err == io.EOF {
        return ""
    }

    return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
    if err != nil {
        panic(err)
    }
}
