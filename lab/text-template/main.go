package main

import (
	"os"
	"text/template"
)

type User struct {
	Name  string
	Email string
	Admin bool
}

func main() {

	tpl := `Hello {{.Name}},
Your email is {{.Email}}.
{{if .Admin}}You are an admin.{{else}}You are a regular user.{{end}}
`

	t := template.Must(template.New("example").Parse(tpl))

	user := User{
		Name:  "Arsen",
		Email: "arsen@example.com",
		Admin: true,
	}

	t.Execute(os.Stdout, user)

}
