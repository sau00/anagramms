package handler

import "fmt"

const _errAnagramsNotFound = "Anagrams not found"

var ErrAnagramsNotFound = fmt.Errorf(_errAnagramsNotFound)
