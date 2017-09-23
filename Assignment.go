package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
)

//Total struct
type Total struct {
	Project   string   `json:"project"`
	Owner     string   `json:"owner"`
	Committer string   `json:"committer"`
	Commits   int      `json:"commits"`
	Languages []string `json:"language"`
}

//Repo struct
type Repo struct {
	Owner struct {
		Login string `json:"login"`
	}
	Project string `json:"html_url"`
}

//Committer struct
type Committer struct {
	Login         string `json:"login"`
	Contributions int    `json:"contributions"`
}

func getInfo(url string, t interface{}) error {

	r, err := http.Get(url)
	if err != nil {
		return err
	}
	defer r.Body.Close()

	fmt.Println(url)
	decoder := json.NewDecoder(r.Body)
	return decoder.Decode(&t)
}

func handler(w http.ResponseWriter, r *http.Request) {
	http.Header.Add(w.Header(), "content-type", "application/json")

	args := r.URL.Path
	part := strings.Split(args, "/")

	url := "https://api.github.com/repos/" + part[4] + "/" + part[5]
	commURL := url + "/contributors"
	langURL := url + "/languages"

	data1 := &Repo{}
	data2 := &[]Committer{}
	data3 := []string{}
	tot := &Total{}

	lang := make(map[string]interface{})

	err := getInfo(url, data1)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	err = getInfo(commURL, data2)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}

	err = getInfo(langURL, &lang)
	if err != nil {
		http.Error(w, http.StatusText(500), http.StatusInternalServerError)
		return
	}
	for i := range lang {
		data3 = append(data3, i)
	}

	tot.Owner = data1.Owner.Login
	tot.Project = data1.Project
	tot.Committer = (*data2)[0].Login
	tot.Commits = (*data2)[0].Contributions
	tot.Languages = data3

	json.NewEncoder(w).Encode(tot)
}

func main() {
	http.HandleFunc("/projectinfo/v1/", handler)
	http.ListenAndServe(":8080", nil)
}
