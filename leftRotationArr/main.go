package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// Before we rotate the array around its axis several times,
// we take the mod of the asked rotation based on the length
// using the real rotation we get the slice of the array
// that will be put in the begining and the slice that will
// be in the end
// join these two parts and return the result
func rotLeft(a []int32, d int32) []int32 {
    realrot := int(d)%len(a)
    var rotLeft []int32 = a[:realrot]
    resultArr := make([]int32, len(a))
    j := realrot
    for ; j < len(a); j++{
        index := j-realrot
        resultArr[index] = a[j]
    }
    for i := 0; i < realrot; i++{
        valleft := rotLeft[i]
        resultArr[len(a)-realrot+i] = valleft
    }
    return resultArr
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	outfile := os.Getenv("OUTPUT_PATH")+"/leftRotationArr.txt"
    stdout, err := os.Create(outfile)

    checkError(err)

    defer stdout.Close()

    writer := bufio.NewWriterSize(stdout, 1024 * 1024)

    nd := strings.Split(readLine(reader), " ")

    nTemp, err := strconv.ParseInt(nd[0], 10, 64)
    checkError(err)
    n := int32(nTemp)

    dTemp, err := strconv.ParseInt(nd[1], 10, 64)
    checkError(err)
    d := int32(dTemp)

    aTemp := strings.Split(readLine(reader), " ")

    var a []int32

    for i := 0; i < int(n); i++ {
        aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
        checkError(err)
        aItem := int32(aItemTemp)
        a = append(a, aItem)
    }

    result := rotLeft(a, d)

    for i, resultItem := range result {
        fmt.Fprintf(writer, "%d", resultItem)

        if i != len(result) - 1 {
            fmt.Fprintf(writer, " ")
        }
    }

    fmt.Fprintf(writer, "\n")

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
