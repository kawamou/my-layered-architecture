package domain

import (

)

type Memos []Memo

type Memo struct {
	ID int
	Title string
	Text string
	Date string
}