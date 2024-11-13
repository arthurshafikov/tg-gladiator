package services

import (
	"strings"

	"github.com/arthurshafikov/tg-gladiator/internal/core/errors"
)

type ValidatorService struct {
	logger Logger
}

func newValidatorService(
	logger Logger,
) *ValidatorService {
	return &ValidatorService{
		logger: logger,
	}
}

func (s *ValidatorService) SanitizeHeroName(heroName string) (string, error) {
	heroName = s.replaceTelegramSpecialCharacters(heroName)

	heroNameLen := len([]rune(heroName))
	if heroNameLen < 1 {
		return "", errors.ErrEmpty
	}

	if heroNameLen > 50 {
		return "", errors.ErrTooLong(50)
	}

	return heroName, nil
}

func (s *ValidatorService) replaceTelegramSpecialCharacters(str string) string {
	if strings.Contains(str, "<") {
		str = strings.ReplaceAll(str, "<", "&lt;")
	}
	if strings.Contains(str, ">") {
		str = strings.ReplaceAll(str, ">", "&gt;")
	}
	if strings.Contains(str, "&") {
		str = strings.ReplaceAll(str, "&", "&amp;")
	}

	return str
}
