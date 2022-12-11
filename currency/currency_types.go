package currency

import (
  "github.com/alphamystic/precision"
)

type Currency struct {
  CValue *precision.Decimal
  Code: string
  Decimals int32
}
