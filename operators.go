package precision

import (
  "math/big"
)
// Add adds two Decimals and returns the result.
func (d *Decimal) Add(o *Decimal) *Decimal {
  scale := d.Scale
  if scale < o.Scale {
    scale = o.Scale
  }
  value := new(big.Int)
  if scale == d.Scale {
	   value.Add(d.Value, o.Value.Mul(o.Value, big.NewInt(int64(10^(scale-o.Scale)))))
    } else {
	     value.Add(d.Value.Mul(d.Value, big.NewInt(int64(10^(scale-d.Scale)))), o.Value)
    }
    return &Decimal{
	     Value: value,
	     Scale: scale,
    }
}

// Sub subtracts two Decimals and returns the result.
func (d *Decimal) Sub(o *Decimal) *Decimal {
  scale := d.Scale
  if scale < o.Scale {
    scale = o.Scale
  }
  value := new(big.Int)
  if scale == d.Scale {
	   value.Sub(d.Value, o.Value.Mul(o.Value, big.NewInt(int64(10^(scale-o.Scale)))))
    } else {
	     value.Sub(d.Value.Mul(d.Value, big.NewInt(int64(10^(scale-d.Scale)))), o.Value)
    }
  return &Decimal{
    Value: value,
    Scale: scale,
  }
}

// Mul multiplies two Decimals and returns the result.
func (d *Decimal) Mul(o *Decimal) *Decimal {
  value := new(big.Int).Mul(d.Value, o.Value)
  scale := d.Scale + o.Scale
  return &Decimal{
	   Value: value,
     Scale: scale,
  }
}

// Div divides two Decimals and returns the result.
func (d *Decimal) Div(o *Decimal) *Decimal {
  value := new(big.Int).Quo(d.Value, o.Value)
  scale := d.Scale - o.Scale
  return &Decimal{
	   Value: value,
	   Scale: scale,
   }
}

/* CHECKERS */
/* Cmp compares two Decimals and returns:
    -1 if d < o
    0 if d == o
    +1 if d > o
*/
func (d *Decimal) Cmp(o *Decimal) int {
  scale := d.Scale
  if scale < o.Scale {
    scale = o.Scale
  }
  if scale == d.Scale {
	   return d.Value.Cmp(o.value.Mul(o.value, big.NewInt(int64(10^(scale-o.Scale)))))
   }
  return d.Value.Mul(d.Value, big.NewInt(int64(10^(scale-d.Scale)))).Cmp(o.value)
}

// IsPositive returns true if the Decimal is positive, false otherwise.
func (d *Decimal) IsPositive() bool {
  return d.Value.Sign() == 1
}

// IsNegative returns true if the Decimal is negative, false otherwise.
func (d *Decimal) IsNegative() bool {
  return d.Value.Sign() == -1
}

// IsZero returns true if the Decimal is zero, false otherwise.
func (d *Decimal) IsZero() bool {
  return d.Value.Sign() == 0
}

// IntPart returns the integer part of the Decimal.
func (d *Decimal) IntPart() *big.Int {
  if d.Scale == 0 {
    return d.Value
  }
  return new(big.Int).Quo(d.Value, big.NewInt(int64(10^d.Scale)))
}

// FractionPart returns the fractional part of the Decimal.
func (d *Decimal) FractionPart() *big.Int {
  if d.Scale == 0 {
    return big.NewInt(0)
  }
  return new(big.Int).Mod(d.Value, big.NewInt(int64(10^d.Scale)))
}

// Round rounds the Decimal to the specified precision and returns the result.
func (d *Decimal) Round(precision int) *Decimal {
  if d.Scale <= precision {
    return d
  }
  value := new(big.Int).Quo(d.Value, big.NewInt(int64(10^(d.Scale-precision))))
  value.Add(value, big.NewInt(int64(5)))
  value.Quo(value, big.NewInt(int64(10)))
  return &Decimal{
    Value: value,
	  Scale: precision,
  }
}

// Abs returns the absolute value of the Decimal.
func (d *Decimal) Abs() *Decimal {
  value := new(big.Int).Abs(d.Value)
  return &Decimal{
    Value: value,
	  Scale: d.Scale,
  }
}

// Neg returns the negated value of the Decimal.
func (d *Decimal) Neg() *Decimal {
  value := new(big.Int).Neg(d.Value)
  return &Decimal{
    Value: value,
	  Scale: d.Scale,
  }
}

// Pow raises the Decimal to the specified power and returns the result.
func (d *Decimal) Pow(power int) *Decimal {
  value := new(big.Int).Exp(d.Value, big.NewInt(int64(power)), nil)
  return &Decimal{
    Value: value,
    Scale: d.Scale * power,
  }
}

// Sqrt calculates the square root of the Decimal and returns the result.
func (d *Decimal) Sqrt() *Decimal {
  if d.Value.Sign() < 0 {
    return nil
  }
  value := new(big.Int).Sqrt(d.Value)
  scale := d.Scale / 2
  return &Decimal{
    Value: value,
    Scale: scale,
  }
}

// Truncate truncates the Decimal to the specified precision and returns the result.
func (d *Decimal) Truncate(precision int) *Decimal {
  if d.Scale <= precision {
    return d
  }
  value := new(big.Int).Quo(d.Value, big.NewInt(int64(10^(d.Scale-precision))))
  return &Decimal{
    Value: value,
    Scale: precision,
  }
}

// Floor calculates the floor of the Decimal and returns the result.
func (d *Decimal) Floor() *big.Int {
  if d.Scale == 0 {
    return d.Value
  }
  return new(big.Int).Quo(d.Value, big.NewInt(int64(10^d.Scale)))
}

// Ceil calculates the ceiling of the Decimal and returns the result.
func (d *Decimal) Ceil() *big.Int {
  if d.Scale == 0 {
    return d.Value
  }
  value := new(big.Int).Quo(d.Value, big.NewInt(int64(10^d.Scale)))
  value.Add(value, big.NewInt(1))
  return value
}

// Int64 returns the int64 representation of the Decimal.
// If the Decimal has a non-zero fractional part, it will be truncated.
func (d *Decimal) Int64()int64 {
  return d.Truncate(0).value.Int64()
}

// Float64 returns the float64 representation of the Decimal.
func (d *Decimal) Float64() float64 {
  return float64(d.Value.Int64()) / math.Pow(10, float64(d.Scale))
}

// Scale returns the scale (number of decimal places) of the Decimal.
func (d *Decimal) Scale() int {
  return d.Scale
}

// DecimalSetScale sets the scale (number of decimal places) of the Decimal.
func (d *Decimal) DecimalSetScale(scale int) *Decimal {
  if d.Scale == scale {
    return d
  }
  value := new(big.Int)
  if scale > d.Scale {
	   value.Mul(d.Value, big.NewInt(int64(10^(scale-d.Scale))))
  } else {
	   value.Quo(d.Value, big.NewInt(int64(10^(d.Scale-scale))))
  }
  return &Decimal{
    Value: value,
	  Scale: scale,
  }
}

// Precision returns the precision (number of significant digits) of the Decimal.
func (d *Decimal) Precision() int {
  return len(d.Value.String())
}

// SetPrecision sets the precision (number of significant digits) of the Decimal.
func (d *Decimal) SetPrecision(precision int) *Decimal {
  if precision >= len(d.Value.String()) {
    return d
  }
  value := new(big.Int)
  if d.Scale == 0 {
	   value.SetString(d.Value.String()[:precision], 10)
  } else {
	   integer := d.IntPart().String()
     fraction := d.FractionPart().String()
	   if len(integer)+len(fraction) <= precision {
		     value.SetString(integer+fraction, 10)
	   } else if len(integer) > precision {
		      value.SetString(integer[:precision], 10)
	   } else {
		     value.SetString(integer+fraction[:precision-len(integer)], 10)
	      }
    }
    return &Decimal{
      Value: value,
      Scale: d.Scale,
    }
}

// Sign returns:
//
// -1 if d < 0
// 0 if d == 0
// +1 if d > 0
func (d *Decimal) Sign() int {
  return d.Value.Sign()
}

// IsInt returns true if the Decimal is an integer, false otherwise.
func (d *Decimal) IsInt() bool {
  return d.Scale == 0
}

// IsNaN returns true if the Decimal is not a number, false otherwise.
func (d *Decimal) IsNaN() bool {
  return false
}

// IsInf returns true if the Decimal is infinity, false otherwise.
func (d *Decimal) IsInf() bool {
  return false
}

// Int returns the *big.Int representation of the Decimal.
func (d *Decimal) Int() *big.Int {
  return d.Value
}

// Rat returns the *big.Rat representation of the Decimal.
func (d *Decimal) Rat() *big.Rat {
  return new(big.Rat).SetInt(d.Value).Quo(big.NewRat(1, int64(10^d.Scale)))
}

// Float returns the *big.Float representation of the Decimal.
func (d *Decimal) Float() *big.Float {
  f := new(big.Float).SetInt(d.Value)
  f.Quo(f, big.NewFloat(math.Pow(10, float64(d.Scale))))
  return f
}

// IntString returns the string representation of the integer part of the Decimal.
func (d *Decimal) IntString() string {
  if d.Scale == 0 {
    return d.Value.String()
  }
  return d.IntPart().String()
}

// FractionString returns the string representation of the fractional part of the Decimal.
func (d *Decimal) FractionString() string {
  if d.Scale == 0 {
    return "0"
  }
  str := d.FractionPart().String()
  if len(str) < d.Scale {
	   return strings.Repeat("0", d.Scale-len(str)) + str
   }
   return str
}

// Exp returns the Decimal e^d, where e is the base of the natural logarithm.
func (d *Decimal) Exp() *Decimal {
  f := new(big.Float).SetInt(d.Value)
  f.Quo(f, big.NewFloat(math.Pow(10, float64(d.Scale))))
  f.Exp(f)
  value, _ := new(big.Int).SetString(f.Text('f', d.Scale), 10)
  return &Decimal{
    Value: value,
    Scale: d.Scale,
  }
}

// Log returns the natural logarithm of the Decimal.
func (d *Decimal) Log() *Decimal {
  if d.Value.Sign() <= 0 {
    return nil
  }
  f := new(big.Float).SetInt(d.Value)
  f.Quo(f, big.NewFloat(math.Pow(10, float64(d.Scale))))
  f.Log(f)
  value, _ := new(big.Int).SetString(f.Text('f', d.Scale), 10)
  return &Decimal{
    Value: value,
    Scale: d.Scale,
  }
}

// Log10 returns the base-10 logarithm of the Decimal.
func (d *Decimal) Log10() *Decimal {
  if d.Value.Sign() <= 0 {
    return nil
  }
  f := new(big.Float).SetInt(d.Value)
  f.Quo(f, big.NewFloat(math.Pow(10, float64(d.Scale))))
  f.Log10
  value, _ := new(big.Int).SetString(f.Text('f', d.Scale), 10)
  return &Decimal{
    Value: value,
    Scale: d.Scale,
  }
}

// Pow10 raises 10 to the specified power and returns the result as a Decimal.
func Pow10(power int) *Decimal {
  value := new(big.Int).Exp(big.NewInt(10), big.NewInt(int64(power)), nil)
  return &Decimal{
    Value: value,
    Scale: power,
  }
}

// One returns the Decimal representation of 1.
func One() *Decimal {
  return &Decimal{
    Value: big.NewInt(1),
    Scale: 0,
  }
}

// Zero returns the Decimal representation of 0.
func Zero() *Decimal {
  return &Decimal{
    Value: big.NewInt(0),
    Scale: 0,
  }
}

// Pi returns the Decimal representation of pi.
func Pi() *Decimal {
  value, _ := new(big.Int).SetString("3.141592653589793238462643383279502884197169399375105820974944592307816406286208998628034825342117067982148086513282306647093844609550582231725359408128481117450284102701938521105559644622948954930381964428810975665933446128475648233785303861", 10)
  return &Decimal{
    Value: value,
    Scale: 100,
  }
}

// E returns the Decimal representation of e.
func E() *Decimal {
  value, _ := new(big.Int).SetString("2.7182818284590452353602874713526624977572470936999595749669676277240766303535475945713821785251664274274663919320030599218174135966290435729003342952605956307381323286279434907632338298807531952510190115738341879307021540891499348841675092447614606680", 10)
  return &Decimal{
    Value: value,
    Scale: 100,
  }
}


// Scan scans a value from a string.
func (d *Decimal) Scan(s fmt.ScanState, ch rune) error {
  var intString, fractionString string
  var scale int
  for {
    ch, _, err := s.ReadRune()
	  if err != nil {
		    return err
	     }
    if ch == '.' {
      ch, _, err = s.ReadRune()
		    if err != nil {
          return err
		      }
		    for ch != ' ' && ch != '\t' && ch != '\n' {
          fractionString += string(ch)
          scale++
          ch, _, err = s.ReadRune()
			    if err != nil {
				     return err
           }
         }
		     break
       }
	     intString += string(ch)
       ch, _, err = s.ReadRune()
       if err != nil {
         return err
       }
     }
  value, ok := new(big.Int).SetString(intString+fractionString, 10)
  if !ok {
	   return fmt.Errorf("invalid decimal value: %s", intString+fractionString)
   }
  *d = Decimal{
    Value: value,
    Scale: scale,
  }
  return nil
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (d *Decimal) UnmarshalJSON(data []byte) error {
  var str string
  if err := json.Unmarshal(data, &str); err != nil {
    return err
  }
  var intString, fractionString string
  var scale int
  for i, ch := range str {
    if ch == '.' {
      intString = str[:i]
      fractionString = str[i+1:]
		  scale = len(fractionString)
		  break
	   }
   }
   if intString == "" {
     intString = str
   }
   value, ok := new(big.Int).SetString(intString+fractionString, 10)
   if !ok {
	    return fmt.Errorf("invalid decimal value: %s", intString+fractionString)
    }
  *d = Decimal{
      Value: value,
      Scale: scale,
    }
    return nil
}

// MarshalJSON implements the json.Marshaler interface.
func (d *Decimal) MarshalJSON() ([]byte, error) {
  return json.Marshal(d.String())
}

// Quo returns the quotient x / y.
func (x *Decimal) Quo(y *Decimal) *Decimal {
  scale := x.scale + y.scale
  value := new(big.Int).Mul(x.value, big.NewInt(int64(10^scale)))
  value.Quo(value, y.value)
  return &Decimal{
    Value: value,
    Scale: scale,
  }
}

// Round returns the Decimal rounded to the nearest integer.
//
// If the fractional part is greater than or equal to 0.5, the result is rounded up.
// Otherwise, it is rounded down.
func (d *Decimal) Round(scale int) *Decimal {
  if d.Scale == scale {
    return d
  }
  if d.Scale < scale {
	   return d.SetScale(scale)
   }
  value := new(big.Int)
  if d.Value.Sign() >= 0 {
	   value.Quo(d.Value, big.NewInt(int64(10^(d.Scale-scale))))
	   if d.Value.Mod(d.Value, big.NewInt(int64(10^(d.Scale-scale)))).Cmp(big.NewInt(int64(5*10^(d.Scale-scale-1)))) >= 0 {
		     value.Add(value, big.NewInt(1))
	      }
     } else {
	      value.Div(d.Value, big.NewInt(int64(10^(d.Scale-scale))))
	      if d.Value.Mod(d.Value, big.NewInt(int64(10^(d.Scale-scale)))).Cmp(big.NewInt(int64(-5*10^(d.Scale-scale-1)))) <= 0 {
		        value.Sub(value, big.NewInt(1))
	       }
    }
  return &Decimal{
    Value: value,
    Scale: scale,
  }
}

// SetScale sets the scale of the Decimal and returns the result.
func (d *Decimal) SetScale(scale int) *Decimal {
  if d.Scale == scale {
    return d
  }
  if d.Scale < scale {
	   return &Decimal{
       Value: new(big.Int).Mul(d.Value, big.NewInt(int64(10^(scale-d.Scale)))),
       Scale: scale,
	    }
  }
  return &Decimal{
    Value: new(big.Int).Quo(d.Value, big.NewInt(int64(10^(d.Scale-scale)))),
    Scale: scale,
  }
}

// Scale returns the scale of the Decimal.
func (d *Decimal) Scale() int {
  return d.Scale
}

// Prec returns the precision of the Decimal.
func (d *Decimal) Prec() int {
  if d.Value.Sign() == 0 {
    return 1
  }
  prec, _ := new(big.Int).SetString(d.Value.String(), 10)
}
return prec.BitLen()

// IntString returns the integer part of the Decimal as a string.
func (d *Decimal) IntString() string {
  return d.IntPart().String()
}


// Int64 returns the Decimal as an int64.
func (d *Decimal) ToInt64() int64 {
  value := d.Truncate(0)
  return value.Value.Int64()
}

// Int32 returns the Decimal as an int32.
func (d *Decimal) Int32() int32 {
  value := d.Truncate(0)
  return int32(value.Value.Int64())
}

// Int16 returns the Decimal as an int16.
func (d *Decimal) Int16() int16 {
  value := d.Truncate(0)
  return int16(value.Value.Int64())
}

// Int8 returns the Decimal as an int8.
func (d *Decimal) Int8() int8 {
  value := d.Truncate(0)
  return int8(value.Value.Int64())
}

// Int returns the Decimal as an int.
func (d *Decimal) Int() int {
  value := d.Truncate(0)
  return int(value.Value.Int64())
}

// Uint64 returns the Decimal as a uint64.
func (d *Decimal) Uint64() uint64 {
  value := d.Truncate(0)
  return value.Value.Uint64()
}

// Uint32 returns the Decimal as a uint32.
func (d *Decimal) Uint32() uint32 {
  value := d.Truncate(0)
  return uint32(value.Value.Uint64())
}

// Uint16 returns the Decimal as a uint16.
func (d *Decimal) Uint16() uint16 {
  value := d.Truncate(0)
  return uint16(value.Value.Uint64())
}

// Uint8 returns the Decimal as a uint8.
func (d *Decimal) Uint8() uint8 {
  value := d.Truncate(0)
  return uint8(value.Value.Uint64())
}

// Uint returns the Decimal as a uint.
func (d *Decimal) Uint() uint {
  value := d.Truncate(0)
  return uint(value.Value.Uint64())
}

// Float64 returns the Decimal as a float64.
func (d *Decimal) Float64() float64 {
  return d.Value.Float64() * math.Pow10(-d.Scale)
}

// Float32 returns the Decimal as a float32.
func (d *Decimal) Float32() float32 {
  return float32(d.Value.Float64() * math.Pow10(-d.Scale))
}


// IsInt returns true if the Decimal has a scale of 0, and false otherwise.
func (d *Decimal) IsInt() bool {
  return d.Scale == 0
}

// IsZero returns true if the Decimal is equal to 0, and false otherwise.
func (d *Decimal) IsZero() bool {
  return d.Value.Sign() == 0
}


// NewFromInt returns a new Decimal with the given int value and scale.
func NewFromInt(value int, scale int) *Decimal {
  return New(big.NewInt(int64(value)), scale)
}

// NewFromInt64 returns a new Decimal with the given int64 value and scale.
func NewFromInt64(value int64, scale int) *Decimal {
  return New(big.NewInt(value), scale)
}

// NewFromInt32 returns a new Decimal with the given int32 value and scale.
func NewFromInt32(value int32, scale int) *Decimal {
  return New(big.NewInt(int64(value)), scale)
}

// NewFromInt16 returns a new Decimal with the given int16 value and scale.
func NewFromInt16(value int16, scale int) *Decimal {
  return New(big.NewInt(int64(value)), scale)
}

// NewFromInt8 returns a new Decimal with the given int8 value and scale.
func NewFromInt8(value int8, scale int) *Decimal {
  return New(big.NewInt(int64(value)), scale)
}

// NewFromUint returns a new Decimal with the given uint value and scale.
func NewFromUint(value uint, scale int) *Decimal {
  return New(new(big.Int).SetUint64(uint64(value)), scale)
}

// NewFromUint64 returns a new Decimal with the given uint64 value and scale.
func NewFromUint64(value uint64, scale int) *Decimal {
  return New(new(big.Int).SetUint64(value), scale)
}

// NewFromUint32 returns a new Decimal with the given uint32 value and scale.
func NewFromUint32(value uint32, scale int) *Decimal {
  return New(new(big.Int).SetUint64(uint64(value)), scale)
}

// NewFromUint16 returns a new Decimal with the given uint16 value and scale.
func NewFromUint16(value uint16, scale int) *Decimal {
  return New(new(big.Int).SetUint64(uint64(value)), scale)
}

// NewFromUint8 returns a new Decimal with the given uint8 value and scale.
func NewFromUint8(value uint8, scale int) *Decimal {
  return New(new(big.Int).SetUint64(uint64(value)), scale)
}

// NewFromFloat64 returns a new Decimal with the given float64 value and scale.
func NewFromFloat64(value float64, scale int) *Decimal {
  valueStr := strconv.FormatFloat(value, 'f', -1, 64)
  decimalValue, decimalScale, err := big.ParseFloat(valueStr, 10, -1, big.ToNearestEven)
  if err != nil {
	   return New(big.NewInt(0), 0)
  }
  decimalScale -= int(math.Log10(math.Abs(value)))
  return New(decimalValue, decimalScale+scale)
}

// NewFromFloat32 returns a new Decimal with the given float32 value and scale.
func NewFromFloat32(value float32, scale int) *Decimal {
  return NewFromFloat64(float64(value), scale)
}

// NewFromString returns a new Decimal with the given string value and scale.
func NewFromString(value string, scale int) *Decimal {
  decimalValue, decimalScale, err := big.ParseFloat(value, 10, -1, big.ToNearestEven)
  if err != nil {
    return New(big.NewInt(0), 0)
  }
  return New(decimalValue, decimalScale+scale)
}

// ParseDecimal parses a decimal value from a string and returns a new Decimal.
func ParseDecimal(s string) (*Decimal, error) {
  s = strings.TrimSpace(s)
  if s == "" {
	   return nil, errors.New("decimal: string is empty")
  }
  valueStr := s
  scale := 0
  if strings.Contains(s, ".") {
	   parts := strings.Split(s, ".")
	   valueStr = parts[0] + parts[1]
	   scale = len(parts[1])
  }
  value, success := new(big.Int).SetString(valueStr, 10)
  if !success {
	   return nil, errors.New("decimal: unable to parse value from string")
  }
  return New(value, scale), nil
}
