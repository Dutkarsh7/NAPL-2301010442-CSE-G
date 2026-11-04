# Experiment 7: Error Handling and Unit Testing

## Objective
To understand and implement error handling patterns in Go using the built-in `error` type and `fmt.Errorf()`, and write comprehensive unit tests using Go's built-in `testing` package to verify both valid inputs and error conditions.

## Concepts Used

1. **Go Error Handling Model**:
   - Go does not use exceptions (like `try-catch`). Instead, functions return an `(result, error)` pair as their last return value.
   - The caller explicitly checks `if err != nil` to handle errors gracefully without crashing the program.
2. **Dynamic Error Creation (`fmt.Errorf()`)**:
   - Formats descriptive, contextual error messages using verb specifiers (e.g., `%.2f`, `%s`).
3. **Unit Testing in Go (`testing` package)**:
   - Test files are named with the suffix `_test.go` and placed alongside the code being tested.
   - Test functions follow the signature `func TestXxx(t *testing.T)`.
   - Subtests using `t.Run(name, func(t *testing.T))` organize test cases hierarchically.
   - Failure methods:
     - `t.Errorf()`: Logs a test failure and continues test execution.
     - `t.Fatalf()`: Logs a fatal test failure and terminates that test immediately.

## Program Functions

1. **`Divide(a, b float64) (float64, error)`**:
   - Checks if divisor `b == 0`.
   - Returns descriptive error: `"division by zero is undefined: cannot divide %.2f by 0"`.
2. **`SquareRoot(x float64) (float64, error)`**:
   - Checks if `x < 0`.
   - Returns descriptive error: `"math error: square root of negative number %.2f is not defined in real numbers"`.
3. **`ValidateInput(input string) (string, error)`**:
   - Checks if `input == ""`.
   - Returns descriptive error: `"validation error: input string cannot be empty"`.

## Unit Test Suite (`main_test.go`)

- **`TestDivide`**:
  - `Valid division`: Asserts that `Divide(10.0, 2.0)` yields `5.00` without errors.
  - `Division by zero`: Asserts that `Divide(10.0, 0.0)` returns an error (`err != nil`).
- **`TestSquareRoot`**:
  - `Valid positive number`: Asserts that `SquareRoot(16.0)` yields `4.00` without errors.
  - `Square root of zero`: Asserts that `SquareRoot(0.0)` yields `0.00` without errors.
  - `Negative number error`: Asserts that `SquareRoot(-9.0)` returns an error (`err != nil`).
- **`TestValidateInput`**:
  - `Valid non-empty string`: Asserts that non-empty string passes without errors.
  - `Empty string error`: Asserts that `""` returns an error (`err != nil`).

## Commands to Run

### Run the Main Program

```bash
cd lab-04/experiment-07-error-handling-testing
go run main.go
```

### Run Unit Tests (Verbose)

```bash
go test -v
```

### Save Test Output to File

```bash
go test -v | Out-File -Encoding utf8 test_output.txt
```

### Format Code

```bash
gofmt -w main.go main_test.go
```

## Program Output (`go run main.go`)

```text
================================================================
  Lab 4 - Experiment 7: Error Handling & Unit Testing
================================================================

--- Scenario 1: Division Operation ---
Valid Case   : Divide(10.0, 2.0) = 5.00
Invalid Case : Divide(10.0, 0.0) -> Handled Error: division by zero is undefined: cannot divide 10.00 by 0

--- Scenario 2: Square Root Operation ---
Valid Case   : SquareRoot(25.0) = 5.00
Valid Case   : SquareRoot(0.0)  = 0.00
Invalid Case : SquareRoot(-16.0) -> Handled Error: math error: square root of negative number -16.00 is not defined in real numbers

--- Scenario 3: Input String Validation ---
Valid Case   : ValidateInput("Go Language Practical") -> Validated: "Go Language Practical"
Invalid Case : ValidateInput("") -> Handled Error: validation error: input string cannot be empty

================================================================
  All error scenarios handled gracefully without program crash!
================================================================
```

## Test Output (`go test -v`)

```text
=== RUN   TestDivide
=== RUN   TestDivide/Valid_division
=== RUN   TestDivide/Division_by_zero
--- PASS: TestDivide (0.00s)
    --- PASS: TestDivide/Valid_division (0.00s)
    --- PASS: TestDivide/Division_by_zero (0.00s)
=== RUN   TestSquareRoot
=== RUN   TestSquareRoot/Valid_positive_number
=== RUN   TestSquareRoot/Square_root_of_zero
=== RUN   TestSquareRoot/Negative_number_error
--- PASS: TestSquareRoot (0.00s)
    --- PASS: TestSquareRoot/Valid_positive_number (0.00s)
    --- PASS: TestSquareRoot/Square_root_of_zero (0.00s)
    --- PASS: TestSquareRoot/Negative_number_error (0.00s)
=== RUN   TestValidateInput
=== RUN   TestValidateInput/Valid_non-empty_string
=== RUN   TestValidateInput/Empty_string_error
--- PASS: TestValidateInput (0.00s)
    --- PASS: TestValidateInput/Valid_non-empty_string (0.00s)
    --- PASS: TestValidateInput/Empty_string_error (0.00s)
PASS
ok  	experiment7_error_handling_testing	0.739s
```
