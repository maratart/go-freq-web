package main

import (
	"fmt"
	"net/http"

	"github.com/maratart/go-freq/freq"
)

const form = `
<div style="text-align: left; margin-top: 50px;">
<form action="post">
<textarea name="text" cols="80" rows="25"></textarea><br />
<input type="submit" value="Calculate">
</form>
</div>
`

const output = `
%s,%d<br />
`

func handler(w http.ResponseWriter, r *http.Request) {

	// set content-type for the answer page
	w.Header().Set("Content-Type", "text/html; charset=utf-8")

	// show form
	fmt.Fprintf(w, form)

	// get text
	// TODO: injection test
	text := r.FormValue("text")

	fw, err := freq.GetSortedFreq(text)
	if err != nil {
		fmt.Fprintf(w, "%s<br />", err)
	} else {
		n := len(fw)
		for i := n - 1; i >= 0; i-- {
			fmt.Fprintf(w, output, fw[i].Key, fw[i].Value)
		}
	}
	// fmt.Fprintf(w, "%s<br />", http.imeout)
}

func main() {
	// run server
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8002", nil)
}
