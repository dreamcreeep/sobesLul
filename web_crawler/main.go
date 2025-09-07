package main

import "time"

// Условие:
// - у нас есть web-crawler, который скачивает страницы с популярных social media платформ
// - скачанные вед страницы отпраыляются в очередь сообщений
// - наша задача - предоставить API, через которое можно возвращать список сайтов и кол-во результатов,
//	скачанных для этих сайтов за последние 10 минут

// facebook.com 542000
// twitter.com 321000
// linkedin.com 289000
// github.com 256000

type Solution interface {
	OnPageDownloaded(host string, timestamp time.Time)
	Count(host string) int // last 10 minutes
}
