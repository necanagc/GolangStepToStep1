package main

// import (
// 	"errors"
// 	"strings"
// 	"time"
// 	"unicode/utf8"
// )

// var (
// 	BadDayError = errors.New("исправь свой ответ, а лучше ложись поспать")
// )

// func currentDayOfTheWeek() string {
// 	var t = TimeNow()
// 	d := t.Weekday()
// 	switch d {
// 	case time.Monday:
// 		return "Понедельник"
// 	case time.Tuesday:
// 		return "Вторник"
// 	case time.Wednesday:
// 		return "Среда"
// 	case time.Thursday:
// 		return "Четверг"
// 	case time.Friday:
// 		return "Пятница"
// 	case time.Saturday:
// 		return "Суббота"
// 	case time.Sunday:
// 		return "Воскресенье"
// 	default:
// 		return ""
// 	}
// }

// func dayOrNight() string {
// 	var t = TimeNow()
// 	d := t.Hour()
// 	if d >= 10 && d <= 22 {
// 		return "День"
// 	} else {
// 		return "Ночь"
// 	}
// }

// func nextFriday() int {
// 	var t = TimeNow()
// 	return (5 - int(t.Weekday()) + 7) % 7
// }

// func CheckCurrentDayOfTheWeek(answer string) bool {

// 	return strings.ToLower(answer) == strings.ToLower(currentDayOfTheWeek())
// }

// func CheckNowDayOrNight(answer string) (bool, error) {

// 	if utf8.RuneCountInString(answer) > 4 || utf8.RuneCountInString(answer) < 4 {
// 		return false, BadDayError
// 	}

// 	return strings.ToLower(dayOrNight()) == strings.ToLower(answer), nil
// }
