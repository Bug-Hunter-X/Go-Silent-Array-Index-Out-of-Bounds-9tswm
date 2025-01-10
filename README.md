# Go Silent Array Index Out of Bounds

This repository demonstrates a subtle bug in Go where array index out of bounds access does not result in a runtime panic, but instead silently corrupts data.  This can lead to difficult-to-debug issues.

## Bug Description
Go arrays have a fixed size.  Accessing an index beyond this size should ideally result in a runtime error.  However, Go doesn't explicitly check for this in all scenarios, leading to unexpected behavior.

## How to Reproduce
1. Clone this repository.
2. Navigate to the directory in your terminal.
3. Run the code using `go run bug.go`.  Observe the output; the program will not crash.

## Solution
The solution involves being explicit in your code to prevent this issue using explicit bounds checks.

## Additional Notes
This issue often surfaces in situations where you have a loop that iterates beyond the bounds of the array, or where you have a potential for invalid user input feeding into array index calculations.