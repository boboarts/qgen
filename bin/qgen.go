package main

import (
	"fmt"
	"io/ioutil"
	"math/rand"
	"net/http"
	"strconv"
	"strings"
	"time"
)

func xi(x int) string {
	result := ""
	switch x {
	case 1:
		result = ""
	case -1:
		result = "-"
	default:
		// result = string(x)
		result = strconv.Itoa(x)
	}
	return result
}

func xis(x int) string {
	if x > 0 {
		return "+" + xi(x)
	}
	return xi(x)
}

func cm(x int) string {
	if x > 0 {
		return "+" + strconv.Itoa(x)
	}
	return strconv.Itoa(x)
}

func genProblems(level int, sign bool, problemCount int) (problems []string, answers []string, seed int64) {

	// problems = make([]string, problem_count)
	// answers = make([]string, problem_count)

	// seed = time.Now().UnixNano()
	// rand.Seed(seed)

	// for i := 0; i < problem_count; i++ {
	// 	a := rand.Intn(level) + 1
	// 	b := rand.Intn(level) + 1
	// 	c := rand.Intn(level) + 1
	// 	d := rand.Intn(level) + 1
	// 	problems[i] = fmt.Sprintf("%dx^2+%dx+%d=", a*b, a*d+b*c, c*d)
	// 	answers[i] = fmt.Sprintf("%dx^2+%dx+%d=(%dx+%d)*(%dx+%d)", a*b, a*d+b*c, c*d, a, c, b, d)
	// }

	problems = make([]string, problemCount)
	answers = make([]string, problemCount)

	seed = time.Now().UnixNano()
	rand.Seed(seed)

	var pvalues []int
	if sign {
		for i := -5 * level; i <= 5*level; i++ {
			if i == 0 {
				continue
			}
			pvalues = append(pvalues, i)
		}
	} else {
		for i := 1; i <= 5*level; i++ {
			if i == 0 {
				continue
			}
			pvalues = append(pvalues, i)
		}
	}

	length := len(pvalues)
	for i := 0; i < problemCount; i++ {

		a := pvalues[rand.Intn(length)]
		b := pvalues[rand.Intn(length)]
		c := pvalues[rand.Intn(length)]
		d := pvalues[rand.Intn(length)]

		// problems[i] = fmt.Sprintf("%sx^2+%sx+%d=", xi(a*b), xim(a*d+b*c), c*d)
		// answers[i] = fmt.Sprintf("%sx^2+%sx+%d=(%sx+%d)*(%sx+%d)", xi(a*b), xim(a*d+b*c), c*d, xi(a), c, xi(b), d)
		problems[i] = fmt.Sprintf("%sx^2%sx%+d=", xi(a*b), xis(a*d+b*c), c*d)
		answers[i] = fmt.Sprintf("%sx^2%sx%s=(%sx%+d)*(%sx%+d)", xi(a*b), xis(a*d+b*c), cm(c*d), xi(a), c, xi(b), d)
	}
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {

	tplHTMLData, err1 := ioutil.ReadFile("template.html")
	var newHtml string = ""
	if err1 == nil {
		html := string(tplHTMLData)
		// fmt.Println(html)

		var problemsHTML, answersHTML string = "", ""
		problems, answers, _ := genProblems(1, false, 15)
		for index := 0; index < len(problems); index++ {
			// fmt.Println(problems[index])
			problemsHTML += fmt.Sprintf("$$%s$$\n", problems[index])
			answersHTML += fmt.Sprintf("$$%s$$\n", answers[index])
		}

		// fmt.Println(problemsHTML)
		// fmt.Println(answersHTML)
		newHtml = strings.Replace(html, "#PROBLEMS#", problemsHTML, -1)
		newHtml = strings.Replace(newHtml, "#ANSWERS#", answersHTML, -1)
		// fmt.Println(newHtml)
	}

	// write to disk
	// problemHtmlData := []byte(newHtml)
	// err2 := ioutil.WriteFile("problems.html", problemHtmlData, 0644)
	// if err2 == nil {
	// 	fmt.Println("Done.")
	// }

	fmt.Fprintf(w, newHtml)
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.ListenAndServe(":9923", nil)

}
