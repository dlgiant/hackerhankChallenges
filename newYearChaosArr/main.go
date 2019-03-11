package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// In beginning of first for loop if index of item+3
// is less than item value, there is no way for that value
// to achieve its position in less then 2 moves, thus, we
// return "Too chaotic"
// otherwise, if the current value-2 is greater than 0
// use that as a starting index, otherwise use 0
// Until the second index reaches the first iterate over
// array and compare with current value, if value iterated
// is greater, add to moves
// Return number of moves
func minimumBribes(q []int32) {
    var moves int
    for i, v := range q {
        if i + 3 < int(v) {
            fmt.Println("Too chaotic")
            return
        } else {
            var j int
            if 0 > v-2 {
                j = 0
            } else {
                j = int(v)-2
            }
            for ; j < i; j++ {
                if q[j] > v {
                    moves++
                }
            }
        }
    }
    fmt.Println(moves)
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
    checkError(err)
    t := int32(tTemp)

    for tItr := 0; tItr < int(t); tItr++ {
        nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
        checkError(err)
        n := int32(nTemp)

        qTemp := strings.Split(readLine(reader), " ")

        var q []int32

        for i := 0; i < int(n); i++ {
            qItemTemp, err := strconv.ParseInt(qTemp[i], 10, 64)
            checkError(err)
            qItem := int32(qItemTemp)
            q = append(q, qItem)
        }

        minimumBribes(q)
    }
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
