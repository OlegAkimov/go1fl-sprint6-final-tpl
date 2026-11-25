package service

import (
	"strings"

	"github.com/Yandex-Practicum/go1fl-sprint6-final/pkg/morse"
)

// DetectAndConvert определяет, является ли входная строка кодом Морзе или текстом,
// и возвращает конвертированную строку.
func DetectAndConvert(input string) (string, error) {
	trimmed := strings.TrimSpace(input)
	if len(trimmed) == 0 {
		return "", nil
	}

	// Если строка содержит только символы Морзе (., -, пробелы) и содержит хотя бы одну точку или тире — это Морзе.
	isMorse := strings.ContainsAny(trimmed, ".-") && !strings.ContainsAny(trimmed, "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789юяыэъюягдкмнпсфхцчшщ")

	if isMorse {
		return morse.ToText(trimmed), nil
	}

	return morse.ToMorse(trimmed), nil
}
