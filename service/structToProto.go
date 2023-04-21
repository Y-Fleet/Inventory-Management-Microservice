package service

import (
	models "inventory/models"

	api "github.com/Y-Fleet/Grpc-Api/api"
)

func ItemsToProto(items []models.Item, err error) *api.GetItemResponse {
	if err != nil {
		return nil
	}
	response := &api.GetItemResponse{}

	for _, item := range items {
		protoItem := &api.Item{
			ID:   item.ID.Hex(),
			Name: item.Name,
			Desc: item.Desc,
			Kg:   item.Kg,
		}
		response.Items = append(response.Items, protoItem)
	}
	return response
}
