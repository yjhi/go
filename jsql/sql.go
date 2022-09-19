/****************************************************************************
MIT License

Copyright (c) 2022 yjhi

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
*****************************************************************************/
package jsql

import (
	"database/sql"
)

/*
* add by yjh 211109
* define base type
*
 */
type SqlUtils struct {
	_error error
	Db     *sql.DB
	Result sql.Result
	Tx     *sql.Tx
	Stmt   *sql.Stmt
}

func (s *SqlUtils) SetError(err error) {
	s._error = err
}

func (s *SqlUtils) Error() string {
	if s._error == nil {
		return ""
	} else {
		return s._error.Error()
	}
}

func (s *SqlUtils) ExecCmd(cmd string) bool {

	if s.Db == nil {
		return false
	}

	var stmt *sql.Stmt
	stmt, s._error = s.Db.Prepare(cmd)

	if s._error != nil {
		return false
	}
	s.Result, s._error = stmt.Exec()

	return s._error == nil
}

func (s *SqlUtils) ExecCmdWithArgs(cmd string, args ...interface{}) bool {

	if s.Db == nil {
		return false
	}

	var stmt *sql.Stmt
	stmt, s._error = s.Db.Prepare(cmd)

	if s._error != nil {
		return false
	}

	s.Result, s._error = stmt.Exec(args...)

	return s._error == nil
}

func (s *SqlUtils) GetInt(cmd string) (int64, bool) {

	if s.Db == nil {
		return 0, false
	}

	var count int64

	s._error = s.Db.QueryRow(cmd).Scan(&count)

	if s._error != nil {
		return 0, false
	}
	return count, true
}

func (s *SqlUtils) GetString(cmd string) (string, bool) {

	if s.Db == nil {
		return "", false
	}

	var str string

	s._error = s.Db.QueryRow(cmd).Scan(&str)

	if s._error != nil {
		return "", false
	}

	return str, true
}

func (s *SqlUtils) Begin() bool {

	if s.Db == nil {
		return false
	}

	s.Tx, s._error = s.Db.Begin()

	return s._error == nil

}

func (s *SqlUtils) End() bool {
	if s.Db == nil {
		return false
	}

	if s.Tx == nil {
		return false
	}

	s._error = s.Tx.Commit()

	if s._error != nil {
		s.Tx.Rollback()

		return false
	}

	return true

}

func (s *SqlUtils) Prepare(cmd string) bool {

	if s.Db == nil {
		return false
	}

	s.Stmt, s._error = s.Db.Prepare(cmd)

	return s._error == nil
}

func (s *SqlUtils) ExecBatch(args ...interface{}) bool {
	if s.Db == nil {
		return false
	}

	if s.Stmt == nil {
		return false
	}

	s.Result, s._error = s.Stmt.Exec(args...)

	return s._error == nil

}
