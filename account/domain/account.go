package domain

type Account struct {
	ID    string `db:"car002_row_id"`
	Code  string `db:"car002_hesapkodu"`
	Name1 string `db:"car002_unvan1"`
	Name2 string `db:"car002_unvan2"`
}
