package service

import (
	models "inventory/models"

	pb "github.com/Yfleet/shared_proto/api"
)

func protoToStruct(rq *pb.AddItemRequest) models.Item {
	return models.Item{
		Name: rq.GetName(),
		Desc: rq.GetDesc(),
		Kg:   rq.GetKg(),
	}
}
