// +build windows

package hosts_parser

import (
	"os"
)

var HostsPath = os.Getenv("SystemRoot") + `\System32\drivers\etc\hosts`
