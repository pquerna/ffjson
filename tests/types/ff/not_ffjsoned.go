package ff

// ffjson: skip
type NoFF struct {
    Everything
    Baz int
}

func NewNoFF(n *NoFF) {
    NewEverything(&n.Everything)
    n.Baz = 32
}