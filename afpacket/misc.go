package afpacket

import (
	"bytes"
	"os/exec"
	"strings"
)

func supportRingMmap() bool {
	return strings.HasPrefix(getOSVersion(), "3.")
}

func getOSVersion() string {
	cmd := exec.Command("uname", "-srio")
	cmd.Stdin = strings.NewReader("some input")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr
	err := cmd.Run()
	if err != nil {
		return ""
	}

	osInfo := out.String()
	for strings.Index(osInfo, "broken pipe") != -1 {
		return ""
	}

	osStr := strings.Replace(osInfo, "\n", "", -1)
	osStr = strings.Replace(osStr, "\r\n", "", -1)
	osArr := strings.Split(osStr, " ")
	return osArr[1]
}
