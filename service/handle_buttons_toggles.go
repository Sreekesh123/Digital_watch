package service

import "fmt"

type WatchMode int

const (
    LettersMode WatchMode = iota
    NumbersMode
)

var CurrentMode = LettersMode

func HandleButton1() {
    if CurrentMode == LettersMode {
        fmt.Println("A")
    } else {
        fmt.Println("1")
    }
}

func HandleButton2() {
    if CurrentMode == LettersMode {
        fmt.Println("B")
    } else {
        fmt.Println("2")
    }
}

func HandleButton3() {
    if CurrentMode == LettersMode {
        fmt.Println("C")
    } else {
        fmt.Println("3")
    }
}


func ToggleMode() {
    if CurrentMode == LettersMode {
        CurrentMode = NumbersMode
        fmt.Println("Switch to Number Mode")
    } else {
        CurrentMode = LettersMode
        fmt.Println("Switch to Letter Mode")
    }
}