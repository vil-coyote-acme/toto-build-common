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
	"github.com/nsqio/nsq/nsqd"
	"github.com/nsqio/nsq/nsqlookupd"
)

type Broker struct {
	exit chan bool
}

func NewBroker() *Broker {
	return new(Broker)
}

func (b *Broker) Start() {
	b.exit = make(chan bool, 1)
	// start nsqlookup first
	go func() {
		opt := nsqlookupd.NewOptions()
		opt.BroadcastAddress = "127.0.0.1"
		broker := nsqlookupd.New(opt)
		broker.Main()
		<-b.exit
		broker.Exit()
	}()
	// then nsqd
	go func() {
		opt := nsqd.NewOptions()
		opt.NSQLookupdTCPAddresses = []string{"127.0.0.1:4160"}
		n := nsqd.New(opt)
		n.Main()
		<-b.exit
		n.Exit()
	}()
}

func (b *Broker) Stop() {
	b.exit <- true
}
