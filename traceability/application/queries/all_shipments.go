package queries

import (
	"katsan/traceability/application/repositories"
	"katsan/traceability/infrastructure/gateways"
)

type GetAllShipmentsQuery struct{}

type GetAllShipmentsResponse struct {
	ID          string
	AccountCode string
	ProductCode string
	Quantity    float64
	BatchSize   float64
}

type GetAllShipmentsQueryHandler interface {
	Handle(GetAllShipmentsQuery) ([]GetAllShipmentsResponse, error)
}

type getAllShipmentsQueryHandler struct {
	repo repositories.TraceabilityQueryRepository
}

func (h *getAllShipmentsQueryHandler) Handle(GetAllShipmentsQuery) ([]GetAllShipmentsResponse, error) {
	response := []GetAllShipmentsResponse{}

	shipment, err := h.repo.GetAllShipments()
	if err != nil {
		return response, err
	}

	for _, shipmentProduct := range shipment.ShipmentProducts {
		response = append(response, GetAllShipmentsResponse{
			ID:          shipment.ID,
			AccountCode: shipment.AccountCode,
			ProductCode: shipmentProduct.ProductCode,
			Quantity:    shipmentProduct.Quantity,
			BatchSize:   shipmentProduct.BatchSize,
		})
	}

	return response, nil
}

func NewGetAllShipmentsQueryHandler() GetAllShipmentsQueryHandler {
	return &getAllShipmentsQueryHandler{
		repo: gateways.NewTraceabilityQueryGateway(),
	}
}
