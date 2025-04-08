package pkg

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
	Cool  string
}

// Constant that remains the same
const StableConst = 100

// Function that will have signature changed in v2
func ChangingFunction(s string, b bool) int {
	return len(s)
}