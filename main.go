package main

import "encoding/json"
import "net/http"
import "fmt"

type Song struct {
    No		string
    Title	string
    Album	string
}

var Songs = []Song{
    Song{"1", "Papercut", "Hybrid Theory"},
    Song{"2", "Lying From You", "Meteora"},
    Song{"3", "With You", "Hybrid Theory"},
    Song{"4", "Faint", "Meteora"},
}

func songs(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var result, err = json.Marshal(Songs)

        if err != nil {
            http.Error(w, err.Error(), http.StatusInternalServerError)
            return
        }

        w.Write(result)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func song(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    if r.Method == "POST" {
        var no = r.FormValue("no")
        var result []byte
        var err error

		for _, each := range Songs {
            if each.No == no {
				result, err = json.Marshal(each)
                
                if err != nil {
                    http.Error(w, err.Error(), http.StatusInternalServerError)
                    return
                }

                w.Write(result)
                return
            }
        }

        http.Error(w, "Song not found", http.StatusBadRequest)
        return
    }

    http.Error(w, "", http.StatusBadRequest)
}

func main() {
    http.HandleFunc("/songs", songs)
    http.HandleFunc("/song", song)

    fmt.Println("starting web server at http://localhost:8080/")
    http.ListenAndServe(":8080", nil)
}

