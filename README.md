# Simple Int64 Calculator

This is a simple command line-calculator program written in Go. It takes 3 arguments: 2 numeric values and an mathematical operator.

## Usage

To use the calculator, run the program with the following command-line arguments.

    go run . <value1> <operator> <value2>

* `<value1>`: The first numeric value
* `<operator>`: The mathematical operator (`+, -, /, *, %`). The "*" should always be wrapped in double quotes.
* `<value2>`: The second numeric value

## Example

    go run . 5 + 3

The output will be 8.

## Error Handling

* If the number of arguments is not exactly 4, the program exits.
* If the operator is not one of the supported operators (`+, -, /, *, %`), the program exits.
* If either of the numeric values cannot be converted to an integer, the program exits.
* Division by zero and modulo by zero errors are handled.

## Limitations

* The calculator only works with integer values.
* The maximum and minimum values for operands are limited to 9223372036854775807 and -9223372036854775807 respectively.

## Note

This calculator does not support floating-point operations.



















