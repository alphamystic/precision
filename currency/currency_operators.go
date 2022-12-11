package currency

import (
  "fmt"
  "github.com/alphamystic/precision"
)

func NewCurrency(value *precision.Decimal,code string,decimals int32) *Currency{
  return &Currency {
    CValue: value,
    Code: code,
    Decimals: decimals,
  }
}

// Takes ina rate (float64) that
func (c *Currency) Convert(conversionRate float64,newCode string,decimalPoints int) *Currency {
	// Calculate the converted value using the conversion rate
	convertedValue := c.CValue.Mul(conversionRate)
	// Round the converted value to the appropriate number of decimal places
	convertedValue = convertedValue.Truncate(decimalPoints)
	// Return a new Currency value with the converted value
  c.Code = newCode
	return NewCurrency(convertedValue, c.Code, c.Decimals)
}

func (c *Currency) CurrencyAdd(other *Currency) (*Currency, error) {
	// Check if the other currency has the same code and number of decimal places as this one
	if c.Code != other.Code || c.Decimals != other.Decimals {
		// If not, return an error
		return nil, fmt.Errorf("Cannot add currencies with different codes or decimal places")
	}
	// Add the two values together
	sum := c.CValue.Add(other.CValue)
	// Return a new Currency value with the sum
	return NewCurrency(sum, c.Code, c.Decimals), nil
}

func (c *Currency) CurrencySub(other *Currency) (*Currency, error) {
	// Check if the other currency has the same code and number of decimal places as this one
	if c.Code != other.Code || c.Decimals != other.Decimals {
		// If not, return an error
		return nil, fmt.Errorf("Cannot subtract currencies with different codes or decimal places")
	}
	// Subtract the other value from this one
	difference := c.CValue.Sub(other.CValue)
	// Return a new Currency value with the difference
	return NewCurrency(difference, c.Code, c.Decimals), nil
}

func (c *Currency) CurrencyMul(other *Currency) (*Currency, error) {
	// Check if the other currency has the same code and number of decimal places as this one
	if c.Code != other.Code || c.Decimals != other.Decimals {
		// If not, return an error
		return nil, fmt.Errorf("Cannot multiply currencies with different codes or decimal places")
	}
	// Multiply the two values together
	product := c.CValue.Mul(other.CValue)
	// Return a new Currency value with the product
	return NewCurrency(product, c.Code, c.Decimals), nil
}

func (c *Currency) CurrencyDiv(other *Currency) (*Currency, error) {
	// Check if the other currency has the same code and number of decimal places as this one
  if c.Code != other.Code || c.Decimals != other.Decimals {
		// If not, return an error
		return nil, fmt.Errorf("Cannot multiply currencies with different codes or decimal places")
	}
  // Divide the two values together
	divided := c.CValue.Div(other.CValue)
	// Return a new Currency value with the product
	return NewCurrency(divided, c.Code, c.Decimals), nil
}
