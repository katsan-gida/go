package gateways

import "github.com/midir99/sqload"

var queries = sqload.MustLoadFromString[struct {
	LoadActiveOrders string `query:"LoadActiveOrders"`
}](`
-- query: LoadActiveOrders
SELECT
    car002_row_id,
    car002_unvan1,
    car002_unvan2
FROM
    car002
`)
