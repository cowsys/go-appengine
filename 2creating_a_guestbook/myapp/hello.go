package hello

import (
	"html/template"
	"net/http"
	"time"

	"appengine"
	"appengine/datastore"
	"appengine/user"
)

type Greeting struct {
	Author  string
	Content string
	Date    time.Time
}

func init() {
	http.HandleFunc("/", root)
	http.HandleFunc("/sign", sign)
}

func guestbookKey(c appengine.Context) *datastore.Key {
	// The string "default_guestbook" here could be varied to have multiple guestbooks.
	// https://cloud.google.com/appengine/docs/go/datastore/reference#NewKey
	//// func NewKey(c context.Context, kind, stringID string, intID int64, parent *Key) *Key
	////
	//// NewKey creates a new key. kind cannot be empty.
	//// Either one or both of stringID and intID must be zero.

	//// If both are zero, the key returned is incomplete.
	//// parent must either be a complete key or nil.
	return datastore.NewKey(c, "Guestbook", "default_guestbook", 0, nil)
}

func root(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)

	// First the function constructs a Query value that requests Greeting objects that are descendants of the root guestbook key,
	// in Date-descending order, with a limit of 10 objects.
	//
	//                                  まだここの理解度低い
	//                                  where句のようなものだとは思うけど
	q := datastore.NewQuery("Greeting").Ancestor(guestbookKey(c)).Order("-Date").Limit(10)
	greetings := make([]Greeting, 0, 10)

	if _, err := q.GetAll(c, &greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	if err := guestbookTemplate.Execute(w, greetings); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

var guestbookTemplate = template.Must(template.New("book").Parse(`
<html lang="en">
<head>
	<title>Go Guestbook</title>
</head>
<body>
	{{range .}}
		{{with .Author}}
			<p><b>{{.}}</b> wrote:</p>
		{{else}}
			<p>An anonymous person wrote:</p>
		{{end}}
		<pre>{{.Content}}</pre>
	{{end}}
	<form action="/sign" method="post">
		<div><textarea name="content" cols="60" rows="3"></textarea></div>
		<div><input type="submit" value="Sign Guestboook"></div>
	</form>
</body>
</html>
`))

func sign(w http.ResponseWriter, r *http.Request) {
	c := appengine.NewContext(r)
	g := Greeting{
		Content: r.FormValue("content"),
		Date:    time.Now(),
	}

	if u := user.Current(c); u != nil {
		g.Author = u.String()
	}

	// We set the same parent key on every Greeting entity to ensure
	// each Greeting is in the same entity group.

	// Queries across the single entity group will be consistent.

	// However, the write rate to a single entity group
	// should be limited to ~1/second.
	key := datastore.NewIncompleteKey(c, "Greeting", guestbookKey(c))
	_, err := datastore.Put(c, key, &g)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	http.Redirect(w, r, "/", http.StatusFound)
}
