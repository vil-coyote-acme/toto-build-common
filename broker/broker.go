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
package broker

import (
	"github.com/nsqio/nsq/nsqd"
	"github.com/nsqio/nsq/nsqlookupd"
	"sync"
)

type Broker struct {
	lookup              *nsqlookupd.NSQLookupd
	broker              *nsqd.NSQD
				   //
	BrokerAddr          string // default value 0.0.0.0
	BrokerPort          string // default value 4150
	BrokerBroadCastAddr string // default 127.0.0.1 -- change it

	LookUpTcpAddrr      string
	LookUpTcpPort       string
	LookUpHttpAddrr     string
	LookUpHttpPort      string
}

func NewBroker() *Broker {
	b := new(Broker)
	b.BrokerAddr = "0.0.0.0"
	b.BrokerPort = "4150"
	b.BrokerBroadCastAddr = "127.0.0.1"
	b.LookUpTcpAddrr = "0.0.0.0"
	b.LookUpTcpPort = "4160"
	b.LookUpHttpAddrr = "0.0.0.0"
	b.LookUpHttpPort = "4161"
	return b
}

func (b *Broker) Start() {
	// start nsqlookup first
	b.StartLookUp()
	// then nsqd
	b.StartBroker()
}

func (b *Broker) StartLookUp() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		opt := nsqlookupd.NewOptions()
		opt.TCPAddress = b.LookUpTcpAddrr + ":" + b.LookUpTcpPort
		opt.HTTPAddress = b.LookUpHttpAddrr + ":" + b.LookUpHttpPort
		b.lookup = nsqlookupd.New(opt)
		b.lookup.Main()
		wg.Done()
	}()
	wg.Wait()
}

func (b *Broker) StartBroker() {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		opt := nsqd.NewOptions()
		opt.TCPAddress = b.BrokerAddr + ":" + b.BrokerPort
		opt.BroadcastAddress = b.BrokerBroadCastAddr
		opt.NSQLookupdTCPAddresses = []string{b.LookUpTcpAddrr + ":" + b.LookUpTcpPort}
		b.broker = nsqd.New(opt)
		b.broker.Main()
		wg.Done()
	}()
	wg.Wait()
}

func (b *Broker) Stop() {
	if b.broker != nil {
		b.broker.Exit()
	}
	if b.lookup != nil {
		b.lookup.Exit()
	}
}