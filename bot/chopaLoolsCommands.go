package bot

import (
	"fmt"
	"math/rand"
	"strings"
)

func ChopaSniff(command []string) string {
	return ChopaActions("нюхаю", command)
}

func ChopaLick(command []string) string {
	return ChopaActions("лижу", command)
}

func ChopaSuck(command []string) string {
	return ChopaActions("сосу", command)
}

func ChopaEat(command []string) string {
	return ChopaActions("ем", command)
}

func ChopaInf() string {
	return fmt.Sprintf("Инфа шансо %d%%", rand.Intn(101))
}

func ChopaActions(action string, command []string) string {
	toAction := command[2:]
	return fmt.Sprintf("*%s %s*", action, strings.Join(toAction, " "))
}
