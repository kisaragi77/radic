package test

import (
	"fmt"
	"testing"

	"github.com/Orisun/radic/v2/demo"
)

func TestGetClassBits(t *testing.T) {
	fmt.Printf("%064b\n", demo.GetClassBits([]string{"五月天", "北京", "资讯", "热点"}))
}

// go test -v ./demo/test -run=^TestGetClassBits$ -count=1
