package converter

import (
	ekko "ekko/api"
	"ekko/package/ent"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func ConvertField(ent *ent.ScenarioField) *ekko.Field {
	return &ekko.Field{
		Id:   ent.ID,
		Name: ent.Name,
		BaseData: &ekko.BaseData{
			CreatedAt: timestamppb.New(ent.CreatedAt),
			UpdatedAt: timestamppb.New(ent.UpdatedAt),
		},
	}
}
