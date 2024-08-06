package gateways

import (
	"katsan"
	"katsan/traceability/domain"

	"github.com/jmoiron/sqlx"
)

type TraceabilityQueryGateway struct {
	db *sqlx.DB
}

func NewTraceabilityQueryGateway() *TraceabilityQueryGateway {
	return &TraceabilityQueryGateway{
		db: katsan.NewDB(nil).GetDB(),
	}
}

func (g *TraceabilityQueryGateway) GetAllShipments() (domain.Shipment, error) {
	shipment := domain.Shipment{}

	rows, err := g.db.Queryx(`
		SELECT STK002_MalKodu, STK002_EvrakSeriNo, STK002_CariHesapKodu, STK002_Miktari ,STK004_KafileBuyuklugu
		FROM stk002
		INNER JOIN "STK004" ON STK004_MalKodu = STK002_MalKodu
		WHERE STK002_TeslimTarihi = $1 AND STK002_GC = 1 AND
			(STK002_MalKodu LIKE '6%' OR STK002_MalKodu LIKE '7%')
		ORDER BY STK002_CariHesapKodu
	`, "45510")

	for rows.Next() {
		var shipmentProduct domain.ShipmentProduct
		err = rows.Scan(
			&shipmentProduct.ProductCode,
			&shipment.ID,
			&shipment.AccountCode,
			&shipmentProduct.Quantity,
			&shipmentProduct.BatchSize,
		)
		if err != nil {
			return shipment, err
		}

		shipment.ShipmentProducts = append(shipment.ShipmentProducts, shipmentProduct)
	}

	if err != nil {
		return shipment, err
	}

	return shipment, nil
}
