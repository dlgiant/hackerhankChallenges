package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Iterate over array using index and current value
// if current index+1 is not the same as the current value
// the value is in incorrect position
// Iterate again to find the value that should be in that
// position and hold the index, break from loop when find it
// Swap elements and count up by 1 the swaps
func minimumSwaps(arr []int32) int32 {
    var swaps int32
    swaps = 0
    for i, v := range arr {
        if i+1 != int(v) {
            var indexv int
            for j := i+1; j < len(arr); j++ {
                if int(arr[j]) == i+1 {
                    indexv = j
                    break
                }
            }
            var val int32
            val = arr[i]
            arr[i] = arr[indexv]
            arr[indexv] = val
            swaps++
        }
    }
    return swaps
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	outfile := os.Getenv("OUTPUT_PATH")+"/minSwapsArr.txt"
    stdout, err := os.Create(outfile)
    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    n := int32(nTemp)

    arrTemp := strings.Split(readLine(reader), " ")

    var arr []int32

    for i := 0; i < int(n); i++ {
        arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
        checkError(err)
        arrItem := int32(arrItemTemp)
        arr = append(arr, arrItem)
    }

    res := minimumSwaps(arr)

    fmt.Fprintf(writer, "%d\n", res)

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
