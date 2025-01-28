package model

type Entries map[string][]string // repo, -pr1 -pr2

type Outputter interface {
	Write(entries Entries) error
}
