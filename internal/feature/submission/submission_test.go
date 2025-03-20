package submission

import (
	ekko "ekko/api"
	"encoding/json"
	"fmt"
	"testing"

	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func TestMain(t *testing.T) {
	req := &ekko.ListScenarioRequest{
		BmIds:          []uint64{},
		SearchContent:  new(string),
		SortMethods:    []*ekko.SortMethod{},
		PageIndex:      0,
		PageSize:       0,
		IsFavorite:     new(bool),
		IsFinished:     new(bool),
		FieldIds:       []uint64{},
		From:           timestamppb.Now(),
		To:             timestamppb.Now(),
		MinRating:      new(float32),
		MinParticipant: new(int32),
	}

	data, err := protojson.Marshal(req)
	if err != nil {
		t.Fatal("Lỗi khi serialize protobuf thành JSON:", err)
	}

	fmt.Println(string(data))

	// Nếu muốn in ra JSON format đẹp hơn:
	var prettyJSON map[string]interface{}
	json.Unmarshal(data, &prettyJSON)
	prettyData, _ := json.MarshalIndent(prettyJSON, "", "  ")
	fmt.Println(string(prettyData))
}
