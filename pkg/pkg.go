package pkg

import "fmt"

type Seer interface {
	See(input string) string
}

type SeerImpl struct{}

func (s *SeerImpl) See(input string) string {
	return "Processed: " + input
}

type SeerImpl2 struct{}

func (s *SeerImpl2) See(input string) string {
	return "Processed 2: " + input
}

// --- Symbols expected to be REMOVED or UNEXPORTED in v2 ---

// Function that will be removed
func RemovedFunction(val int) string {
	return fmt.Sprintf("Removed v1: %d", val)
}

// Type that will be removed
type RemovedType struct {
	ID string
}

// Constant that will be removed
const RemovedConst = "v1_constant"

// Variable that will be removed
var RemovedVar = []string{"a", "b"}

// Function that will become unexported
func ExportedThenUnexportedFunc() {
	fmt.Println("I am exported in v1")
}

// Type that will become unexported
type ExportedThenUnexportedType struct {
	Name string
}

// --- Stable Symbols (Should NOT be reported) ---

// Function that remains the same
func StableFunction(input string) string {
	return "Stable v1: " + input
}

// Type that remains the same
type StableType struct {
	Value int
}

type AType struct {
	age   int
	Value int
	Place string
}

// Constant that remains the same
const StableConst = 100

// Function that will have signature changed in v2
func ChangingFunction(s string) int {
	return len(s)
}