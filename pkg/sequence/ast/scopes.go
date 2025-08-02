package ast

type ScopeDefining interface {
	AppendStmt(Stmt)
}

func (sd *AltBlock) AppendStmt(stmt Stmt)            { sd.Stmts = append(sd.Stmts, stmt) }
func (sd *Break) AppendStmt(stmt Stmt)               { sd.Stmts = append(sd.Stmts, stmt) }
func (sd *CriticalRegionBlock) AppendStmt(stmt Stmt) { sd.Stmts = append(sd.Stmts, stmt) }
func (sd *Diagram) AppendStmt(stmt Stmt)             { sd.Stmts = append(sd.Stmts, stmt) }
func (sd *Loop) AppendStmt(stmt Stmt)                { sd.Stmts = append(sd.Stmts, stmt) }
func (sd *ParallelBlock) AppendStmt(stmt Stmt)       { sd.Stmts = append(sd.Stmts, stmt) }
