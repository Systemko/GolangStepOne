/*
Реализовать загрузку страниц с пяти разных сайтов с помощью горутин
Вывести в консоль URL сайта и размер первоначальной страницы в байтах

Подсказки, что нужно использовать:

import (

	"net/http"
	"io/ioutil"

)

// Получить URL:
response, err := http.Get("www.example.com")

defer response.Body.Close()

// Получить тело ответа
body, err := ioutil.ReadAll(response.Body)

// Рекомендую создать соответствующую структуру Page с полями URL и Size
*/
package main

import (
	"fmt"
	"io"
	"net/http"
)

type Page struct {
	url  string
	size int
}

func responseSize(page *Page, ch chan int) {
	response, err := http.Get(page.url)

	if err == nil {
		defer response.Body.Close()
		body, err := io.ReadAll(response.Body)

		if err == nil {
			ch <- len(body)
		}
	}
}

func main() {
	pages := []*Page{
		{url: "https://ya.ru"},
		{url: "https://google.ru"},
		{url: "https://vk.com"},
		{url: "https://dzen.ru"},
		{url: "https://github.com"},
	}

	ch := make(chan int, len(pages))

	for _, page := range pages {
		go responseSize(page, ch)
	}

	for _, page := range pages {
		page.size = <-ch
		fmt.Printf("Размер страницы %s - %d байт.\n", page.url, page.size)
	}
}
