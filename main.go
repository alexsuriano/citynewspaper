package main

import (
	"fmt"
	"log"
	"os"

	"github.com/antchfx/htmlquery"
)

func main() {

	url := "https://www.jcnet.com.br/noticias/ultimas_noticias"
	var valor string

	doc, err := htmlquery.LoadURL(url)
	if err != nil {
		log.Fatal(err)
	}

	//list, err := htmlquery.QueryAll(doc, "//a")
	list, err := htmlquery.QueryAll(doc, "//a")
	if err != nil {
		log.Fatal(err)
	}

	for i, n := range list {
		a := htmlquery.FindOne(n, "//a")
		if a != nil {
			valor += fmt.Sprintf("%d %s(%s)\n", i, htmlquery.InnerText(a), htmlquery.SelectAttr(a, "href"))
		}
	}

	fmt.Println(valor)

	os.WriteFile("JCNET.txt", []byte(valor), 0644)
}

// <li class="loadMoreItem">
// <a href="/servicos/falecimentos/2021/10/778574-cacilda-cardamone-camargo.html">
// <div class="site">Falecimentos</div>
// <h3>
// CACILDA CARDAMONE CAMARGO </h3>
// <p class="linha-fina">
// Ocorrido em 21/10/2021, com 89 anos. Era viúva do sr JOÃO GOMES DE CAMARGO e deixa os filhos MARA L... </p>
// <div class="dt">21/10/2021 19:35 - <time class="timeago" datetime="2021-10-21T19:35:16Z" title="21/10/2021 19:35">há 7 horas</time></div>
// </a>
// </li>
