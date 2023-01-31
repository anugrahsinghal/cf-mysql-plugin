package cfmysql

import (
	"code.cloudfoundry.org/cli/plugin"
	"fmt"
	"strconv"
)

//go:generate counterfeiter . SshRunner
type SshRunner interface {
	OpenSshTunnel(cliConnection plugin.CliConnection, toService MysqlService, throughApp string, localPort int)
}

func NewSshRunner() SshRunner {
	return new(sshRunner)
}

type sshRunner struct{}

func (self *sshRunner) OpenSshTunnel(cliConnection plugin.CliConnection, toService MysqlService, throughApp string, localPort int) {
	tunnelSpec := strconv.Itoa(localPort) + ":" + toService.Hostname + ":" + toService.Port
	fmt.Printf("ssh %s -N -L %s \n", throughApp, tunnelSpec)
	args, err := cliConnection.CliCommand("ssh", throughApp, "-N", "-L", tunnelSpec)

	if err != nil {
		panic(fmt.Errorf("SSH tunnel failed: %s %s", args, err))
	}
}
