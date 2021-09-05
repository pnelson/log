package log

var (
	_ Handler = &defaultHandler{}
	_ Handler = &textHandler{}
	_ Handler = &shellHandler{}
	_ Handler = &minimalShellHandler{}
	_ Handler = &discardHandler{}
)
