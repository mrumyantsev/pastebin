package errlib

import "fmt"

const (
	tmplWrap    = "%s: %w"
	tmplBlameDb = "error on %s method: %w"
	tmplAppend  = "(additional) " + tmplWrap
)

// The names for blaming database method which caused an error.
const (
	DbSelect   = "Select"
	DbGet      = "Get"
	DbQueryRow = "QueryRow"
	DbBegin    = "Begin"
	DbScan     = "Scan"
	DbExec     = "Exec"
	DbRollback = "Rollback"
	DbCommit   = "Commit"
)

func Wrap(err error, msg string) error {
	return fmt.Errorf(tmplWrap, msg, err)
}

func BlameDbMethod(err error, methodName string) error {
	return fmt.Errorf(tmplBlameDb, methodName, err)
}

func PassOrAppend(err error, add error, addmsg string) error {
	if add == nil {
		return err
	}

	return fmt.Errorf(tmplAppend, addmsg, add)
}
