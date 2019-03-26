package main

import (
	_ "github.com/vtereso/logspout/healthcheck"
	_ "github.com/vtereso/logspout/adapters/raw"
	_ "github.com/vtereso/logspout/adapters/syslog"
	_ "github.com/vtereso/logspout/adapters/multiline"
	_ "github.com/vtereso/logspout/httpstream"
	_ "github.com/vtereso/logspout/routesapi"
	_ "github.com/vtereso/logspout/transports/tcp"
	_ "github.com/vtereso/logspout/transports/udp"
	_ "github.com/vtereso/logspout/transports/tls"
)
