package ip

import (
	"fmt"
	"testing"
)

func TestGetLocalIP4(t *testing.T) {
	ip, err := GetLocalIP4()
	if err != nil {
		t.Error(err)
	}
	fmt.Println(ip)
}
