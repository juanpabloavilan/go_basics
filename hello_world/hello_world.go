package hello_world

const (
	englishPrefix = "Hello, "
	frenchPrefix  = "Bonjour, "
	spanishPrefix = "Hola, "

	spanishLang = "Spanish"
	frenchLang  = "French"
)

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}

	var prefix string = englishPrefix

	switch language {
	case spanishLang:
		prefix = spanishPrefix
	case frenchLang:
		prefix = frenchPrefix
	}
	return prefix + name
}
