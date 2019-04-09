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

// Gernate Factorization Problems
func genFactProblems(level int, sign bool, problemCount int) (problems []string, answers []string, seed int64) {

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

	var possibleValues []int
	if sign {
		for i := -5 * level; i <= 5*level; i++ {
			if i == 0 {
				continue
			}
			possibleValues = append(possibleValues, i)
		}
	} else {
		for i := 1; i <= 5*level; i++ {
			if i == 0 {
				continue
			}
			possibleValues = append(possibleValues, i)
		}
	}

	length := len(possibleValues)
	for i := 0; i < problemCount; i++ {

		a := possibleValues[rand.Intn(length)]
		b := possibleValues[rand.Intn(length)]
		c := possibleValues[rand.Intn(length)]
		d := possibleValues[rand.Intn(length)]

		// problems[i] = fmt.Sprintf("%sx^2+%sx+%d=", xi(a*b), xim(a*d+b*c), c*d)
		// answers[i] = fmt.Sprintf("%sx^2+%sx+%d=(%sx+%d)*(%sx+%d)", xi(a*b), xim(a*d+b*c), c*d, xi(a), c, xi(b), d)
		problems[i] = fmt.Sprintf("%sx^2%sx%+d=", xi(a*b), xis(a*d+b*c), c*d)
		answers[i] = fmt.Sprintf("%sx^2%sx%s=(%sx%+d)*(%sx%+d)", xi(a*b), xis(a*d+b*c), cm(c*d), xi(a), c, xi(b), d)
	}
	return
}

// Gernate Linear Euqation with 2 unknowns Problems
func genLe2uProblems(level int, sign bool, problemCount int) (problems []string, answers []string, seed int64) {

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

	var possibleValues []int
	if sign {
		for i := -5 * level; i <= 5*level; i++ {
			if i == 0 {
				continue
			}
			possibleValues = append(possibleValues, i)
		}
	} else {
		for i := 1; i <= 5*level; i++ {
			if i == 0 {
				continue
			}
			possibleValues = append(possibleValues, i)
		}
	}

	length := len(possibleValues)
	for i := 0; i < problemCount; i++ {

		a := possibleValues[rand.Intn(length)]
		b := possibleValues[rand.Intn(length)]
		c := possibleValues[rand.Intn(length)]
		d := possibleValues[rand.Intn(length)]

		x := possibleValues[rand.Intn(length)]
		y := possibleValues[rand.Intn(length)]
		// problems[i] = fmt.Sprintf("%sx^2+%sx+%d=", xi(a*b), xim(a*d+b*c), c*d)
		// answers[i] = fmt.Sprintf("%sx^2+%sx+%d=(%sx+%d)*(%sx+%d)", xi(a*b), xim(a*d+b*c), c*d, xi(a), c, xi(b), d)
		problems[i] = fmt.Sprintf("%sx%sy=%d \\\\ %sx%sy=%d", xi(a), xis(b), a*x+b*y, xi(c), xis(d), c*x+d*y)
		answers[i] = fmt.Sprintf("x=%d \\\\ y=%d ", x, y)
	}
	return
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	tplHTMLData, _ := ioutil.ReadFile("index.tmpl.html")
	contentHTML := string(tplHTMLData)
	fmt.Fprintf(w, contentHTML)
}

func factHandler(w http.ResponseWriter, r *http.Request) {

	tplHTMLData, err1 := ioutil.ReadFile("problems.tmpl.html")
	var newHTML string
	if err1 == nil {
		html := string(tplHTMLData)
		// fmt.Println(html)

		var problemsHTML, answersHTML string = "", ""
		problems, answers, _ := genFactProblems(1, false, 5)
		for index := 0; index < len(problems); index++ {
			// fmt.Println(problems[index])
			problemsHTML += fmt.Sprintf("$$%s$$\n", problems[index])
			answersHTML += fmt.Sprintf("$$%s$$\n", answers[index])
		}

		// fmt.Println(problemsHTML)
		// fmt.Println(answersHTML)
		newHTML = strings.Replace(html, "#PROBLEMS#", problemsHTML, -1)
		newHTML = strings.Replace(newHTML, "#ANSWERS#", answersHTML, -1)
		// fmt.Println(newHTML)
	}

	// write to disk
	// problemHtmlData := []byte(newHTML)
	// err2 := ioutil.WriteFile("problems.html", problemHtmlData, 0644)
	// if err2 == nil {
	// 	fmt.Println("Done.")
	// }

	fmt.Fprintf(w, newHTML)
}

func le2uHandler(w http.ResponseWriter, r *http.Request) {

	tplHTMLData, err1 := ioutil.ReadFile("problems.tmpl.html")
	var newHTML string
	if err1 == nil {
		html := string(tplHTMLData)
		// fmt.Println(html)

		var problemsHTML, answersHTML string = "", ""
		problems, answers, _ := genLe2uProblems(1, false, 5)
		// problems, answers, _ := genLe2uProblems(1, true, 5)
		for index := 0; index < len(problems); index++ {
			// fmt.Println(problems[index])
			problemsHTML += fmt.Sprintf("$$%d. \\begin{cases}%s\\end{cases}$$\n", index+1, problems[index])
			answersHTML += fmt.Sprintf("$$%d. \\begin{cases}%s\\end{cases}$$\n", index+1, answers[index])
		}

		// fmt.Println(problemsHTML)
		// fmt.Println(answersHTML)
		newHTML = strings.Replace(html, "#PROBLEMS#", problemsHTML, -1)
		newHTML = strings.Replace(newHTML, "#ANSWERS#", answersHTML, -1)
		// fmt.Println(newHTML)
	}

	// write to disk
	// problemHtmlData := []byte(newHTML)
	// err2 := ioutil.WriteFile("problems.html", problemHtmlData, 0644)
	// if err2 == nil {
	// 	fmt.Println("Done.")
	// }

	fmt.Fprintf(w, newHTML)
}

func main() {

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/fact", factHandler)
	http.HandleFunc("/le2u", le2uHandler)
	http.ListenAndServe(":9923", nil)

}
