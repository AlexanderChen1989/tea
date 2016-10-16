package main

import (
	"fmt"

	. "github.com/AlexanderChen1989/tea/view"
)

func App(props Props, children ...Node) Node {
	return Network(
		Props{"name": "isolated_nw"},
		Container(
			Props{
				"name":  "db",
				"image": "mysql:5.7",
				"environment": Props{
					"MYSQL_ROOT_PASSWORD": "wordpress",
					"MYSQL_DATABASE":      "wordpress",
					"MYSQL_USER":          "wordpress",
					"MYSQL_PASSWORD":      "wordpress",
				},
			},
		),
		Container(
			Props{
				"image":      "wordpress:latest",
				"depends_on": "db",
				"links":      []string{"db"},
				"ports":      []string{"8000:80"},
				"environment": Props{
					"WORDPRESS_DB_HOST":     "db:3306",
					"WORDPRESS_DB_PASSWORD": "wordpress",
				},
			},
		),
	)
}

func main() {
	fmt.Printf("%v\n", App(nil))
}
