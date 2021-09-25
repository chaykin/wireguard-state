package connectionstate

import (
	"log"
	"os/exec"
	"strconv"
	"strings"
	"wireguard-state/config"
)

const (
	kib = 1024 << (10 * iota) // 2^20
	mib
	gib
	tib
)

type ConnectionState struct {
	Up bool
	Rx string
	Tx string
}

func (state *ConnectionState) Update(config *config.WgsConfig) {
	wgStateResult := runCmd(config.WgStateCommand)
	if wgStateResult != "" {
		rxResult := runCmd(config.RxCommand)
		txResult := runCmd(config.TxCommand)

		if rxResult != "" && txResult != "" {
			state.Up = true
			state.Rx = formatBytes(rxResult)
			state.Tx = formatBytes(txResult)

			return
		}
	}

	state.Up = false
	state.Rx = ""
	state.Tx = ""
}

func runCmd(cmd string) string {
	result, err := exec.Command("bash", "-c", cmd).Output()
	if err != nil {
		return ""
	}

	return string(result)
}

func formatBytes(bytes string) string {
	num, err := strconv.ParseFloat(strings.TrimSuffix(bytes, "\n"), 64)
	if err != nil {
		log.Println("Error convert value:", err)
		return bytes
	}

	switch {
	case num >= tib:
		return strconv.FormatFloat(num/tib, 'f', 2, 64) + " TiB"
	case num >= gib:
		return strconv.FormatFloat(num/gib, 'f', 2, 64) + " GiB"
	case num >= mib:
		return strconv.FormatFloat(num/mib, 'f', 2, 64) + " MiB"
	case num >= kib:
		return strconv.FormatFloat(num/kib, 'f', 2, 64) + " KiB"
	default:
		return strconv.FormatFloat(num, 'f', 0, 64) + " bytes"
	}
}
