package bot

func (b *Bot) chopaLools(command []string) string {

	switch command[1] {
	case "нюхай":
		return ChopaSniff(command)

	case "лизни", "лижи":
		return ChopaLick(command)

	case "соси":
		return ChopaSuck(command)

	case "ешь":
		return ChopaEat(command)

	case "инфа":
		return ChopaInf()

	default:
		return "чего бля"

	}
}
