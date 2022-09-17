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

	_ "github.com/mattn/go-sqlite3"
)

/*
*  add by yjh 211125
*  for sqlserver
*
 */
type SqliteUtils struct {
	SqlUtils
	DbFile string
}

func BuildSqlite() *SqliteUtils {
	return &SqliteUtils{
		SqlUtils: SqlUtils{
			_error: nil,
		},
		DbFile: "",
	}
}

func CreateSqliteString(file string, config string) string {

	s := "file:" + file

	if len(config) > 0 {
		s += "?" + config
	}

	return s
}

func CreateSqliteStringWithPass(file string, user string, pass string, config string) string {
	s := "file:" + file + "?_auth&_auth_user=" + user + "&_auth_pass=" + pass

	if len(config) > 0 {
		s += "&" + config
	}
	return s

}

func (s *SqliteUtils) OpenFile(dbFile string) bool {

	s.DbFile = dbFile
	s.Db, s._error = sql.Open("sqlite3", s.DbFile)

	return s._error == nil

}

func (s *SqliteUtils) OpenFileWithPass(dbFile string, user string, pass string) bool {

	s.DbFile = dbFile
	s.Db, s._error = sql.Open("sqlite3", CreateSqliteStringWithPass(dbFile, user, pass, ""))

	return s._error == nil

}
