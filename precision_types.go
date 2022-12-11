package precision

// A Decimal iss anumber with an arbitraryprecision
type Decimal struct {
  Value *big.Int
  Scale int
}
