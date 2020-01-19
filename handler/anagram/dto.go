package anagram

import "gitlab.com/anagramms/storage"

type VersionRequest struct{}
type VersionResponse struct {
	Version string
	Message string
}

type LoadRequest struct {
	Anagrams []storage.TAnagram
}

type LoadResponse struct {
	Len     int
	Message string
}

type GetRequest struct {
	Word string
}

type GetResponse struct {
	Len      int
	Anagrams []storage.TAnagram
}
