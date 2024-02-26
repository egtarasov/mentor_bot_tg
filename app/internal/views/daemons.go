package views

import "fmt"

func FeedBackForm(form string) string {
	msg := "Привет, это снова я!\n\n" +
		"Надеюсь, что твоя адаптация идет полным ходом \xF0\x9F\x94\xA5\n\n" +
		"Если будет не сложно, заполни вот эту форму %s, как появится время.\n\n" +
		"Я буду очень признателен \xF0\x9F\x98\x8A"

	return fmt.Sprintf(msg, form)
}
