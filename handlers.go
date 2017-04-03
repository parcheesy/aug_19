package handlers

import (
	"html/template"
	"net/http"
	"time"

	"google.golang.org/appengine"
	"google.golang.org/appengine/datastore"
	"google.golang.org/appengine/user"
)

type Guest struct {
	Name string

	DateRSVPd time.Time

	PlusOne     bool
	PlusOneName string

	InvitedToSangeet   bool
	AttendingSangeet   bool
	InvitedToWedding   bool
	AttendingWedding   bool
	InvitedToReception bool
	AttendingReception bool

	Note string
}

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := appengine.NewContext(r)

	query := datastore.NewQuery("Guest")

    for t := q.Run(ctx); ; {
        var x Guest
        key, err := t.Next(&x)
        if err == datastore.Done {
            break
        }
        if err != nil {
            return
        }

    }

	t, _ := template.ParseFiles("templates/index.html")
	t.Execute(w, nil)
}
