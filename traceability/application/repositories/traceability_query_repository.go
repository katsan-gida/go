package repositories

import "katsan/traceability/domain"

type TraceabilityQueryRepository interface {
	GetAllShipments() (domain.Shipment, error)
}
