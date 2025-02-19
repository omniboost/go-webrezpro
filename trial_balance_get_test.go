package webrezpro_test

import (
	"encoding/json"
	"fmt"
	"testing"
	"time"

	webrezpro "github.com/omniboost/go-webrezpro"
)

func TestGetTrialBalanceRequest(t *testing.T) {
	req := client.NewGetTrialBalanceRequest()
	req.QueryParams().Date = webrezpro.Date{time.Date(2024, 10, 24, 0, 0, 0, 0, time.UTC)}
	resp, err := req.Do()
	if err != nil {
		t.Error(err)
	}

	b, _ := json.MarshalIndent(resp, "", "  ")
	fmt.Println(string(b))
}
