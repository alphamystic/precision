package decimal

import (
  "errors"
  "strings"
  "math/big"
)

func NewDecimal(s string) (*Decimal,error){
  parts := strings.Split(s,".")
  if len(parts)> 2 {
    return nil,errors.New("Invalid decimal format")
  }
  var integer,fraction string
  var scane int
  if len(parts) == 2 {
    integer = parts[0]
    fraction = parts[1]
    scale = len(fraction)
  } else {
    integer = parts[0]
    fraction = ""
    scale = 0
  }
  value, ok := new(big.Int).SetString(integer+fraction,10)
  if !ok {
    return nil,errors.New("Invalid decimal format")
  }
  return &Decimal{
    Value: value,
    Scale: scale,
  },nil
}

// String returns the string representation of the Decimal.
func (d *Decimal) DecimalToString() string {
  if d.Scale == 0 {
    return d.Value.String()
  }
  str := d.Value.String()
  if len(str) <= d.Scale {
	   return "0." + strings.Repeat("0", d.Scale-len(str)) + str
   }
   return str[:len(str)-d.Scale] + "." + str[len(str)-d.Scale:]
}


// String returns the string representation of the Decimal.
func (d *Decimal) IntToString() string {
  return d.IntString() + "." + d.FractionString()
}
