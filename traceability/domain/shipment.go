package domain

type ShipmentProduct struct {
	ProductCode string  `db:"STK002_MalKodu"`
	Quantity    float64 `db:"STK002_Miktari"`
	BatchSize   float64 `db:"STK004_KafileBuyuklugu"`
}

type Shipment struct {
	ID               string `db:"STK002_EvrakSeriNo"`
	ShipmentProducts []ShipmentProduct
	AccountCode      string `db:"STK002_CariHesapKodu"`
}

//STK002_MalKodu, STK002_EvrakSeriNo, STK002_CariHesapKodu,  ,
