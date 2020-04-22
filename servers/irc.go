package servers

import (
	"fmt"
	"net"
	"sync"
	"time"
)

type Conn struct {
	wg           sync.WaitGroup
	conn         net.Conn
	close        chan struct{}
	Messages     chan *Message
	PrivMessages chan *PrivMsg
}

func New() *Conn {
	return &Conn{
		close:        make(chan struct{}),
		Messages:     make(chan *Message),
		PrivMessages: make(chan *PrivMsg),
	}
}

func (c *Conn) send(line string) {
	c.conn.Write([]byte(fmt.Sprintf("%s\r\n", line)))
	time.Sleep(1000 * time.Millisecond)
}

func (c *Conn) Connect(server string, port uint16, nick, realName string) error {
	var err error

	c.conn, err = net.Dial("tcp", fmt.Sprintf("%s:%d", server, port))
	if err != nil {
		return fmt.Errorf("Failed to connect to IRC server: %s", err)
	}

	c.send("USER " + nick + " 8 * :" + realName)
	c.send("NICK " + nick)
	return nil
}
