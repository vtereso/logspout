package main

import (
	_ "github.com/vtereso/logspout/adapters/syslog"
	_ "github.com/vtereso/logspout/transports/tcp"
	_ "github.com/vtereso/logspout/transports/tls"
	_ "github.com/vtereso/logspout/transports/udp"
	_ "github.com/looplab/logspout-logstash"
)
