package webrezpro_test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func TestGetHotelDetailsRequest(t *testing.T) {
	req := client.NewGetHotelDetailsRequest()
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
