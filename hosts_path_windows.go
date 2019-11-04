// +build windows

package hostsparser

import (
	"os"
)

var HostsPath = os.Getenv("SystemRoot") + `\System32\drivers\etc\hosts`
