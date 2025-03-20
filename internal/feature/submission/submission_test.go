package submission

import (
	ekko "ekko/api"
	"fmt"
	"testing"

	"github.com/golang/protobuf/proto"
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

	data := proto.MarshalTextString(req)

	fmt.Println(string(data))
}
