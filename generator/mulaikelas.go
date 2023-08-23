package generator

import (
	"strconv"
	"strings"
)

func MulaiKelas(group_name string) (messages string) {
	listgroup := strings.Split(group_name, "-")
	_, err := strconv.Atoi(strings.TrimSpace(listgroup[0]))
	if (len(listgroup) >= 3) && (err == nil) {
		tokendosenpengganti := strings.Split(listgroup[2], "|")
		if len(tokendosenpengganti) > 1 {
			token := strings.Split(tokendosenpengganti[1], "@")
			messages = "iteung kelas luring dosen pengganti mulai mode tm passcode " + strings.TrimSpace(token[0])
		} else {
			messages = "iteung kelas luring mulai mode tm"
		}

	}
	return
}
