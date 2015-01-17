package bit2int

import (
	"log"
	"testing"
)

func Test_ToInt(t *testing.T) {
	val := BToI(101)

	log.Println(val)
}
