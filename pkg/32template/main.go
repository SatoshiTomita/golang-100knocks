package main

import (
	"html/template"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// テンプレートの定義
		tmpl := `
			<!DOCTYPE html>
			<html lang="en">
			<head>
				<meta charset="UTF-8">
				<title>My Page</title>
			</head>
			<body>
				<h1>Hello, {{.Name}}!</h1>
				<h2>テンプレート内で関数呼び出し</h2>
				<p>
					{{$num := add 1 2}}
					{{$num}}
				</p>
				<h2>テンプレート内でif文</h2>
				<p>
					{{if eq .Name "John"}}
						Johnです
					{{else}}
						Johnではありません
					{{end}}
				</p>
				<h2>比較関数</h2>
				<pre>
					eq
					arg1 == arg2
					ne
					arg1 != arg2
					lt
					arg1 < arg2
					le
					arg1 <= arg2
					gt
					arg1 > arg2
					ge
					arg1 >= arg2
				</pre>
				<h2>テンプレート内でrange文</h2>
				<p>
					{{range $i, $v := .NameList}}
						{{$i}}: {{$v}}
					{{end}}
				</p>
				<h2>テンプレート内で文字列を結合</h2>
				<p>
					{{printf "Hello, %s %s" .Name "!"}}
				</p>
			</body>
			</html>
		`

		// データを作成
		data := struct {
			Name     string
			NameList []string
		}{
			Name:     "John",
			NameList: []string{"John", "Bob", "Alice"},
		}

		// テンプレートに関数を登録
		funcMap := template.FuncMap{
			"add": func(x, y int) int {
				return x + y
			},
		}

		// テンプレートをパース
		t, err := template.New("base").Funcs(funcMap).Parse(tmpl)
		if err != nil {
			log.Fatal(err)
		}

		// テンプレートにデータを埋め込んでレスポンスを書き込む
		err = t.Execute(w, data)
		if err != nil {
			log.Fatal(err)
		}
	})

	log.Println("Server started at :8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}
