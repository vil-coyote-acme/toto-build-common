/*
Toto-build, the stupid Go continuous build server.

Toto-build is free software; you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation; either version 3 of the License, or
(at your option) any later version.

Toto-build is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with this program; if not, write to the Free Software Foundation,
Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301  USA
*/
package testtools

import (
	"bytes"
)

type TestErr struct {
	message string
}

func (err TestErr) Error() string {
	return err.message
}

func NewTestErr(mess string) TestErr {
	err := new(TestErr)
	err.message = mess
	return *err
}

func ConsumeStringChan(c chan string) string {
	var buffer bytes.Buffer
	for line := range c {
		buffer.WriteString(line)
	}
	return buffer.String()
}
