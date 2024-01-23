package im

import (
	"context"
	"fmt"
	"testing"
)

func TestGenAndVerify(t *testing.T) {
	userSig, err := GetClientByName(context.Background(), "1")
	if err != nil {
		fmt.Printf(userSig.Sign)
	}
}
