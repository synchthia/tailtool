package utils

import (
	"strings"
)

var keyStr = []string{"§", "&"}

func ColorText(orig string) string {
	converted := orig + "&r"
	for _, key := range keyStr {
		converted = strings.ReplaceAll(converted, key+"0", "\x1b[30m") // §0 black
		converted = strings.ReplaceAll(converted, key+"1", "\x1b[34m") // §1 dark_blue
		converted = strings.ReplaceAll(converted, key+"2", "\x1b[32m") // §2 dark_green
		converted = strings.ReplaceAll(converted, key+"3", "\x1b[36m") // §3 dark_aqua
		converted = strings.ReplaceAll(converted, key+"4", "\x1b[31m") // §4 dark_red
		converted = strings.ReplaceAll(converted, key+"5", "\x1b[35m") // §5 dark_purple
		converted = strings.ReplaceAll(converted, key+"6", "\x1b[33m") // §6 gold
		converted = strings.ReplaceAll(converted, key+"7", "\x1b[37m") // §7 gray
		converted = strings.ReplaceAll(converted, key+"8", "\x1b[90m") // §8 dark_gray
		converted = strings.ReplaceAll(converted, key+"9", "\x1b[94m") // §9 blue
		converted = strings.ReplaceAll(converted, key+"a", "\x1b[92m") // §a green
		converted = strings.ReplaceAll(converted, key+"b", "\x1b[96m") // §b aqua
		converted = strings.ReplaceAll(converted, key+"c", "\x1b[91m") // §c red
		converted = strings.ReplaceAll(converted, key+"d", "\x1b[95m") // §d light_purple
		converted = strings.ReplaceAll(converted, key+"e", "\x1b[93m") // §e yellow
		converted = strings.ReplaceAll(converted, key+"f", "\x1b[97m") // §f white

		converted = strings.ReplaceAll(converted, key+"k", "\x1b[6m") // Obfuscated
		converted = strings.ReplaceAll(converted, key+"l", "\x1b[1m") // Bold
		converted = strings.ReplaceAll(converted, key+"m", "\x1b[9m") // Strike
		converted = strings.ReplaceAll(converted, key+"n", "\x1b[4m") // Underline
		converted = strings.ReplaceAll(converted, key+"o", "\x1b[3m") // Italic
		converted = strings.ReplaceAll(converted, key+"r", "\x1b[0m") // §r reset
	}

	return converted
}

//func MiniMessage(orig string) string {
//    converted := orig + "\x1b[0m]"
//
//    strings.ReplaceAll()
//}
