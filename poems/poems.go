package main

import (
    "fmt"
    "os"
)

const FILENAME = "verse"

func main() {
    verse := readFile(FILENAME)
    sl := readLine(verse)
    printSlice(sl)
}

func createFile(nameFile string) *os.File {
    file, err := os.Create(nameFile)
    if err != nil {
        panic(err)
    }
    return file
}

func readFile(nameFile string) string {
    f, err := os.ReadFile(nameFile)
    if err != nil {
        panic(err)
        return "err"
    }
    return string(f)
}

func writeToFile(f *os.File, nameFile string, text string) {
    f, err := os.OpenFile(nameFile, os.O_APPEND|os.O_WRONLY, 0600)
    if err != nil {
        panic(err)
    }
    defer f.Close()
    if _, err = f.WriteString(text); err != nil {
        panic(err)
    }
}

func readLine(text string) []string {
    s := []string{}
    for i := 0; i < len(text); i++ {
        s = append(s, string(text[i]))
    }
    return s
}

func printSlice(s []string) {
    for i := 0; i < len(s); i++ {
        fmt.Println(s[i])
    }
}