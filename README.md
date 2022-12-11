# PRECISION

Precision is an implementation of a decimal package in Go.
It includes functions for performing arithmetic operations on decimal values, such as addition, subtraction, multiplication, and division.
It also includes functions for converting decimal values to and from various types, such as int64 and float64.
Additionally, it includes functions for rounding, setting the scale, and obtaining the precision and sign of a decimal value.


When using this package, a value must be turned into the precision.Decimal{} data type structure
say for a value a * b
change both a and b to Decimal then call a.Mul(b)

*Whenever you are receiving a value(amount), convert it into a new decimal then proceed to do your operations*
*When converting

## CURRENCY
To implement a currency conversion functionality in the precision package, we can create a new Currency type and add a Convert() method to it.
The Convert() method would take in a conversion rate and return a new Currency value with the converted value.


## ML
ML is a minimallistic machine learning library implementing the basic machine learning functions that can be imported into an ml project and used on the fly.
We'll probably implement our own matrix package and add sorting/data structures libraries
The whole idea is to have an ai that has it's own weight adjustment techniques.
