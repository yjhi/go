/*
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
*/
package jsql

import (
	"database/sql"
	"fmt"

	_ "github.com/denisenkom/go-mssqldb"
)

/*
*  add by yjh 211125
*  for sqlserver
*
 */
type SqlServerUtils struct {
	SqlUtils
	IsOpen bool
}

func BuildSqlServer() *SqlServerUtils {

	s := &SqlServerUtils{
		SqlUtils: SqlUtils{
			_error: nil,
		},
		IsOpen: false,
	}

	return s
}

func CreateSqlServerString(server string, user string, pass string, db string) string {
	constr := fmt.Sprintf("server=%s;user id=%s;password=%s;database=%s;encrypt=disable",
		server, user, pass, db)
	return constr
}

func (s *SqlServerUtils) ConnectSqlServerString(server string, user string, pass string, db string) bool {
	constr := CreateSqlServerString(server, user, pass, db)

	return s.ConnectSqlServer(constr)
}

func (s *SqlServerUtils) ConnectSqlServer(constr string) bool {

	s.Db, s._error = sql.Open("mssql", constr)

	if s._error == nil {
		s.IsOpen = true
	}

	return s.IsOpen
}

func (s *SqlServerUtils) CloseSqlServer() {
	defer s.Db.Close()
}
