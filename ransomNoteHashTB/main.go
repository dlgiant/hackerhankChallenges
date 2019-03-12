package main

import (
    "bufio"
    "fmt"
    "io"
    "os"
    "strconv"
    "strings"
)

// checkMagazine takes a magazine string array and a note string array
// first we declare a map to count the occurrence of the words in
// magazine
// Populate the map for each value in magazine
// Iterate over the note and check if the value is in the map or if
// the occurrences subtracted to zero, in that case, print "No" and
// leave function
// Otherwise, reduce the counter in the mao
// If the entire iteration of the note went through the checks
// print "Yes"
func checkMagazine(magazine []string, note []string) {
    dict := make(map[string]int, len(magazine))
    for _, v := range magazine {
        dict[v] += 1
    }
    for _, v := range note {
        _, ok := dict[v]
        if !ok || dict[v] <= 0{
            fmt.Println("No")
            return
        } else {
            dict[v]--
        }
    }
    fmt.Println("Yes")
}

func main() {
    reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

    mn := strings.Split(readLine(reader), " ")

    mTemp, err := strconv.ParseInt(mn[0], 10, 64)
    checkError(err)
    m := int32(mTemp)

    nTemp, err := strconv.ParseInt(mn[1], 10, 64)
    checkError(err)
    n := int32(nTemp)

    magazineTemp := strings.Split(readLine(reader), " ")

    var magazine []string

    for i := 0; i < int(m); i++ {
        magazineItem := magazineTemp[i]
        magazine = append(magazine, magazineItem)
    }

    noteTemp := strings.Split(readLine(reader), " ")

    var note []string

    for i := 0; i < int(n); i++ {
        noteItem := noteTemp[i]
        note = append(note, noteItem)
    }

    checkMagazine(magazine, note)
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
