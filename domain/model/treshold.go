package model

import "time"

type Treshold struct {
	DepositHistory []Wallet    `json:"wallet"`
	TimeHistory    []time.Time `json:"time"`
}
