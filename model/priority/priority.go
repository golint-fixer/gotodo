package priority

// Priority represents a todo.txt priority (A-Z)
type Priority int

// MatchRegexp defines the regular expression matching priority strings
const MatchRegexp = `\([A-Z]\)`

//go:generate stringer -type=Priority
const (
	A    Priority = iota
	B    Priority = iota
	C    Priority = iota
	D    Priority = iota
	E    Priority = iota
	F    Priority = iota
	G    Priority = iota
	H    Priority = iota
	I    Priority = iota
	J    Priority = iota
	K    Priority = iota
	L    Priority = iota
	M    Priority = iota
	N    Priority = iota
	O    Priority = iota
	P    Priority = iota
	Q    Priority = iota
	R    Priority = iota
	S    Priority = iota
	T    Priority = iota
	U    Priority = iota
	V    Priority = iota
	W    Priority = iota
	X    Priority = iota
	Y    Priority = iota
	Z    Priority = iota
	None Priority = iota
)

var priorities = map[string]Priority{
	"(A)": A,
	"(B)": B,
	"(C)": C,
	"(D)": D,
	"(E)": E,
	"(F)": F,
	"(G)": G,
	"(H)": H,
	"(I)": I,
	"(J)": J,
	"(K)": K,
	"(L)": L,
	"(M)": M,
	"(N)": N,
	"(O)": O,
	"(P)": P,
	"(Q)": Q,
	"(R)": R,
	"(S)": S,
	"(T)": T,
	"(U)": U,
	"(V)": V,
	"(W)": W,
	"(X)": X,
	"(Y)": Y,
	"(Z)": Z,
}

var priorityStr = map[Priority]string{
	// None is handled separately in MarshalJSON()
	A: "A",
	B: "B",
	C: "C",
	D: "D",
	E: "E",
	F: "F",
	G: "G",
	H: "H",
	I: "I",
	J: "J",
	K: "K",
	L: "L",
	M: "M",
	N: "N",
	O: "O",
	P: "P",
	Q: "Q",
	R: "R",
	S: "S",
	T: "T",
	U: "U",
	V: "V",
	W: "W",
	X: "X",
	Y: "Y",
	Z: "Z",
}

// GetPriority maps input string(s) to their corresponding Priority/ies
func GetPriority(s string) Priority {
	for k, v := range priorities {
		if s == k {
			return v
		}
	}

	return None
}

// MarshalJSON implements Marshaler interface
func (p Priority) MarshalJSON() ([]byte, error) {
	if p == None {
		return []byte("null"), nil
	}
	return []byte(`"` + priorityStr[p] + `"`), nil
}
