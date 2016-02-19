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
	"time"
	"sync"
)

type Broker struct {
	lookup *nsqlookupd.NSQLookupd
	broker *nsqd.NSQD
}

func NewBroker() *Broker {
	return new(Broker)
}

func (b *Broker) Start() {
	// start nsqlookup first
	b.StartLookUp()
	// then nsqd
	b.StartBroker()
}

func (b *Broker) StartLookUp()  {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		opt := nsqlookupd.NewOptions()
		b.lookup = nsqlookupd.New(opt)
		b.lookup.Main()
		wg.Done()
	}()
	wg.Wait()
}

func (b *Broker) StartBroker(){
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		opt := nsqd.NewOptions()
		opt.BroadcastAddress = "127.0.0.1"
		opt.NSQLookupdTCPAddresses = []string{"127.0.0.1:4160"}
		b.broker = nsqd.New(opt)
		b.broker.Main()
		wg.Done()
	}()
	wg.Wait()
}

func (b *Broker) Stop() {
	b.broker.Exit()
	b.lookup.Exit()
}