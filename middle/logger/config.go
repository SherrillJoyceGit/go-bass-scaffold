package logger

import "errors"

type LogType string
type Network string

type Config struct {
	Network   Network
	Host      string
	ServiceId string
	LogType   LogType
	// Log Level
	// TimestampFormat
}

const Logrus LogType = "logrus"

// Known networks are "tcp", "tcp4" (IPv4-only), "tcp6" (IPv6-only),
// "udp", "udp4" (IPv4-only), "udp6" (IPv6-only), "ip", "ip4"
// (IPv4-only), "ip6" (IPv6-only), "unix", "unixgram" and
// "unixpacket".
const (
	tcp Network = "tcp"
	/*	tcp4 Network = "tcp4"
		tcp6 Network = "tcp6"
		udp  Network = "udp"
		udp4 Network = "udp6"
		udp6 Network = "tcp6"
		ip   Network = "ip"
		ip4  Network = "ip4"
		ip6  Network = "ip6"
		unix Network = "unix"
		unixgram   Network = "unixgram"
		unixpacket Network = "unixpacket"*/
)

var CurrentConfig = Config{}

func getCurrentConfig() (*Config, error) {
	if CurrentConfig.Network != tcp ||
		CurrentConfig.Host == "" ||
		CurrentConfig.ServiceId == "" ||
		CurrentConfig.LogType != Logrus {

		return nil, errors.New("current config is not be settled ! ")

	} else {
		return &CurrentConfig, nil
	}
}
