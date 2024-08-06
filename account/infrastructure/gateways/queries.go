package gateways

import (
	"github.com/midir99/sqload"
)

var queries = sqload.MustLoadFromString[struct {
	LoadAllActiveAccounts string `query:"LoadAllActiveAccounts"`
}](`
-- query: LoadAllActiveAccounts
SELECT
    car002_row_id,
    car002_hesapkodu,
    car002_unvan1,
    car002_unvan2
FROM
    car002
WHERE
    car002_aktifflag = 1
    AND car002_grupkodu = '120'
    AND car002_kod3 <> ''
LIMIT $1
OFFSET $2
`)
