package main

import (
    "fmt"
    "strings"
)

var SUBC = "spdfghijklmnoqrtuvwxyz"

type Subshell struct {
    N int
    C string
    MaxE int
}

func newSubshell(n int, c string) Subshell {
    var subshell Subshell
    
    subshell.N = n
    subshell.C = c
    subshell.MaxE = 2 * (2 * strings.Index(SUBC, c) + 1)
    
    return subshell
}

func (s *Subshell) Advance() {
    var nN int
    var nC string
    
    if s.C == "s" {
        nN = (s.N + 1) / 2 + 1;
        nC = string(SUBC[nN - (s.N % 2) - 1])
    } else {
        nN = s.N + 1
        nC = string(SUBC[strings.Index(SUBC, s.C) - 1])
    }
    
    s.N = nN
    s.C = nC
    s.MaxE = 2 * (2 * strings.Index(SUBC, nC) + 1)
}

func main() {
    var an int
    
    cSubshell := newSubshell(1, "s")
    
    fmt.Print("\nEnter the atomic number: ")
    fmt.Scan(&an)
    for an > 0 {
        fmt.Print(cSubshell.N, cSubshell.C)
        if an - cSubshell.MaxE > 0 {
            fmt.Print(cSubshell.MaxE)    
        } else {
            fmt.Print(an)
        }
        an -= cSubshell.MaxE
        cSubshell.Advance() 
        fmt.Print(" ")
    }
}
