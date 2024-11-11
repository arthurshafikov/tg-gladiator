package helpers

import "fmt"

func StrikethroughText(text string) string {
	return fmt.Sprintf("<s>%s</s>", text)
}
