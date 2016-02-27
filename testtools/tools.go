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
	"github.com/nsqio/go-nsq"
	"encoding/json"
	"github.com/vil-coyote-acme/toto-build-common/message"
	"time"
)

type HandlerTest struct {
	Receip chan message.Report
}

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

func (h *HandlerTest) HandleMessage(mes *nsq.Message) (e error) {
	var report message.Report
	json.Unmarshal(mes.Body, &report)
	h.Receip <- report
	return e
}

func ConsumeStringChan(c chan string) string {
	var buffer bytes.Buffer
	for line := range c {
		buffer.WriteString(line)
	}
	return buffer.String()
}

func SetupListener(topic string, lookUpHttpAddrAndPort string) (chan message.Report, *nsq.Consumer) {
	duration, _ := time.ParseDuration("300ms")
	time.Sleep(duration)
	handler := new(HandlerTest)
	receip := make(chan message.Report, 1)
	handler.Receip = receip
	consumer, _ := nsq.NewConsumer(topic, "scheduler", nsq.NewConfig())
	consumer.AddHandler(handler)
	consumer.ConnectToNSQLookupds([]string{lookUpHttpAddrAndPort})
	return receip, consumer
}
