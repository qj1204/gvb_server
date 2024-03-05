package random

import (
	"math/rand"
	"time"
)

func ListID(IDList []uint) uint {
	rand.Seed(time.Now().UnixNano())
	return IDList[rand.Intn(len(IDList))]
}
