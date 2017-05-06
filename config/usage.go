package config

// Usage usage function
func Usage() string {
	return `
NAME:
	{{.Name}} - {{.Usage}}
	
USAGE:
	{{.HelpName}} {{if .VisibleFlags}}[options]{{end}}{{if .Commands}} command [command options]{{end}} {{if .ArgsUsage}}{{.ArgsUsage}}{{else}}[arguments...]{{end}}
{{if len .Authors}}
AUTHOR:
{{range .Authors}}{{ . }}{{end}}
{{end}}{{if .Commands}}
COMMANDS:
{{range .Commands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}{{end}}{{if .VisibleFlags}}
OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}{{end}}{{if .Copyright}}COPYRIGHT: {{.Copyright}}{{end}}{{if .Version}}
VERSION: 
	{{.Version}}{{end}}
`
}

// CommandUsage  Command Usage
func CommandUsage() string {
	return `
NAME:
	{{.Name}} - {{.Usage}}

USAGE:
	pic_gen {{.HelpName}} tempalte_name
	
COMMANDS:
{{range .SubCommands}}{{if not .HideHelp}}   {{join .Names ", "}}{{ "\t"}}{{.Usage}}{{ "\n" }}{{end}}{{end}}

OPTIONS:
	{{range .VisibleFlags}}{{.}}
	{{end}}
`
}
