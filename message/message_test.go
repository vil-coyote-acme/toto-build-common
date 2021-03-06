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
package message_test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"toto-build-common/message"
)

func Test_ToWork_String(t *testing.T) {
	assert.Equal(t, "JobId: 1, Cmd: 0, Package: myPackage, GoVersion: go1.6, RepoUrl: git@github.com:vil-coyote-acme/toto-build-agent.git",
		message.ToWork{int64(1), message.PACKAGE, "myPackage", "go1.6", "git@github.com:vil-coyote-acme/toto-build-agent.git"}.String())
}

func Test_Report_String(t *testing.T)  {
	assert.Equal(t, "JobId: 1, Status: 1, Logs: [my Log] ",
		message.Report{int64(1), message.WORKING, []string{"my Log"}}.String())
}

func Test_Error(t *testing.T) {
	assert.Equal(t, "myError", message.Error{"myError"}.Error())
}
