package nats

import (
	"log"

	nats "github.com/nats-io/nats.go"

	"github.com/amonaco/goapi/libs/config"
)

var _nc *nats.Conn

func Connect() *nats.Conn {
	conf := config.Get()
	nc, err := nats.Connect(conf.Nats)
	_nc = nc
	if err != nil {
		log.Fatal(err)
	}

	return nc
}

func Publish(topic string, msg []byte) error {
	return _nc.Publish(topic, msg)
}

func Subscribe(topic string) (chan *nats.Msg, *nats.Subscription, error) {
	ch := make(chan *nats.Msg)
	sub, err := _nc.ChanSubscribe(topic, ch)
	return ch, sub, err
}
