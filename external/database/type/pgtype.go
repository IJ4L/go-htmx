package typ

import (
	"math/big"

	"github.com/jackc/pgx/v5/pgtype"
)

func Text(text string) pgtype.Text {
	return pgtype.Text{
		String: text,
		Valid:  len(text) != 0,
	}
}

func Numeric(data float64) pgtype.Numeric {
	return pgtype.Numeric{
		Int:   big.NewInt(int64(data)),
		Valid: data != 0,
	}
}
