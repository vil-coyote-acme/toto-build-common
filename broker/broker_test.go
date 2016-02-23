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
package broker_test

import (
	"github.com/stretchr/testify/assert"
	"github.com/nsqio/go-nsq"
	"toto-build-common/broker"
	"testing"
)

func Test_Embedded_Broker_On_Publish(t *testing.T) {
	// given
	b := broker.NewBroker()
	b.Start()
	defer b.Stop()
	// and
	config := nsq.NewConfig()
	p, errP := nsq.NewProducer("127.0.0.1:4150", config)
	// when
	errPub := p.Publish("myTopic", make([]byte, 256))
	//then
	assert.Nil(t, errP)
	assert.Nil(t, errPub)
	assert.NotNil(t, p)
}
