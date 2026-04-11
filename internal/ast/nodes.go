package ast

// Span identifies a source location (best-effort).
type Span struct {
	StartLine int
	StartCol  int
	EndLine   int
	EndCol    int
}

// ----- Program -----

type Program struct {
	Stmts []Stmt
	Span  Span
}

// ----- Statements -----

type Stmt interface {
	isStmt()
	GetSpan() Span
}

type AssignStmt struct {
	Name  string
	Value Expr
	Span  Span
}

func (AssignStmt) isStmt()         {}
func (s AssignStmt) GetSpan() Span { return s.Span }

type IncDecOp string

const (
	Inc IncDecOp = "++"
	Dec IncDecOp = "--"
)

type IncDecStmt struct {
	Name string
	Op   IncDecOp
	Span Span
}

func (IncDecStmt) isStmt()         {}
func (s IncDecStmt) GetSpan() Span { return s.Span }

type PrintStmt struct {
	Args []Expr
	Span Span
}

func (PrintStmt) isStmt()         {}
func (s PrintStmt) GetSpan() Span { return s.Span }

type ReadStmt struct {
	Name string
	Span Span
}

func (ReadStmt) isStmt()         {}
func (s ReadStmt) GetSpan() Span { return s.Span }

type ReadIntStmt struct {
	Name string
	Span Span
}

func (ReadIntStmt) isStmt()         {}
func (s ReadIntStmt) GetSpan() Span { return s.Span }

type ElseIfClause struct {
	Cond Expr
	Body []Stmt
	Span Span
}

type IfStmt struct {
	Cond   Expr
	Then   []Stmt
	ElseIf []ElseIfClause
	Else   []Stmt // nil if absent
	Span   Span
}

func (IfStmt) isStmt()         {}
func (s IfStmt) GetSpan() Span { return s.Span }

type WhileStmt struct {
	Cond Expr
	Body []Stmt
	Span Span
}

func (WhileStmt) isStmt()         {}
func (s WhileStmt) GetSpan() Span { return s.Span }

type DoWhileStmt struct {
	Body []Stmt
	Cond Expr
	Span Span
}

func (DoWhileStmt) isStmt()         {}
func (s DoWhileStmt) GetSpan() Span { return s.Span }

type ForStmt struct {
	Init Stmt // nil or AssignStmt
	Cond Expr // nil allowed by grammar
	Post Stmt // nil or AssignStmt/IncDecStmt
	Body []Stmt
	Span Span
}

func (ForStmt) isStmt()         {}
func (s ForStmt) GetSpan() Span { return s.Span }

type ForEachStmt struct {
	Var  string
	Iter Expr
	Body []Stmt
	Span Span
}

func (ForEachStmt) isStmt()         {}
func (s ForEachStmt) GetSpan() Span { return s.Span }

// ----- Expressions -----

type Expr interface {
	isExpr()
	GetSpan() Span
}

type IdentExpr struct {
	Name string
	Span Span
}

func (IdentExpr) isExpr()         {}
func (e IdentExpr) GetSpan() Span { return e.Span }

type LiteralKind string

const (
	LitInt    LiteralKind = "int"
	LitFloat  LiteralKind = "float"
	LitString LiteralKind = "string"
)

type LiteralExpr struct {
	Kind LiteralKind
	Text string // Keep original token text (e.g. "1_000.0_5")
	Span Span
}

func (LiteralExpr) isExpr()         {}
func (e LiteralExpr) GetSpan() Span { return e.Span }

type ListExpr struct {
	Elems []Expr
	Span  Span
}

func (ListExpr) isExpr()         {}
func (e ListExpr) GetSpan() Span { return e.Span }

type UnaryOp string

const (
	OpPos UnaryOp = "+"
	OpNeg UnaryOp = "-"
)

type UnaryExpr struct {
	Op   UnaryOp
	X    Expr
	Span Span
}

func (UnaryExpr) isExpr()         {}
func (e UnaryExpr) GetSpan() Span { return e.Span }

type BinaryOp string

const (
	OpPow BinaryOp = "**"

	OpMul BinaryOp = "*"
	OpDiv BinaryOp = "/"
	OpMod BinaryOp = "%"

	OpAdd BinaryOp = "+"
	OpSub BinaryOp = "-"

	OpLT  BinaryOp = "<"
	OpLTE BinaryOp = "<="
	OpGT  BinaryOp = ">"
	OpGTE BinaryOp = ">="

	OpEQ  BinaryOp = "=="
	OpNEQ BinaryOp = "!="
)

type BinaryExpr struct {
	Op   BinaryOp
	LHS  Expr
	RHS  Expr
	Span Span
}

func (BinaryExpr) isExpr()         {}
func (e BinaryExpr) GetSpan() Span { return e.Span }
