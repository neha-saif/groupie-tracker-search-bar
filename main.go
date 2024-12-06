package main

import (
	"fmt"
	"groupie/functions"
	"html/template" //text
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"
)

// hfgwhkuef

type Final struct {
	ID        int
	Image     string
	Artist    string
	Members   functions.Members
	AlbumYear string
	Album1    string
	Locations []string
}

func main() {
	// handler functions
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("style"))))
	http.HandleFunc("/homepagesearch", homepagesearch)
	http.HandleFunc("/result", result)
	http.HandleFunc("/", homepage)
	http.ListenAndServe(":8080", nil)
}
func homepagesearch(w http.ResponseWriter, r *http.Request) {
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Failed to parse form data", http.StatusBadRequest)
		return
	}
	// where form value is collected for artist name annd fed innto the relavent funnction
	artIdsug := r.FormValue("suggestions[]")

	fmt.Println("artIdsug:", artIdsug)
	var ids []int

	character, _ := functions.LoadData("https://groupietrackers.herokuapp.com/api/artists")

for _, data := range character {
		fmt.Print(string(data.Artist))
		if strings.Contains(string(data.Artist), strings.Title(artIdsug)){
			ids = append(ids, data.ID)

	} 
	for _, mem := range data.Members {
		if strings.Contains(mem,artIdsug) {
			ids = append(ids, data.ID)
		}
	if strings.Contains(data.Album1,artIdsug) {
		ids = append(ids, data.ID)
	}
	if strings.Contains(data.AlbumYear,artIdsug)  {
		ids = append(ids, data.ID)
	}
	
}
}

	// make thhis

	// for _, sug := range artIdsug {
	// 	parts := strings.Split(string(sug), ":")
	// 	if len(parts) > 1 {
	// 		id, err := strconv.Atoi(strings.TrimSpace(parts[len(parts)-1]))
	// 		if err == nil {
	// 			ids = append(ids, id)
	// 		}
	// 	}
	// }

	fmt.Println("ids:",ids)
	


// Filter char by IDs
filterchar := []functions.Data{}
	for _, id := range ids {
		for _, data := range character {
			if data.ID == id {
				filterchar = append(filterchar, data)
			}
		}
	}


	t, err := template.ParseFiles("index.html")
	if err != nil {
		http.Error(w, "Error parsing html", http.StatusInternalServerError)
		return
	}

	err = t.Execute(w, filterchar)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}

}

func homepage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		renderErrorPage(w, 404)
		return
	}
	character, _ := functions.LoadData("https://groupietrackers.herokuapp.com/api/artists")

	t, err := template.ParseFiles("index.html")
	if err != nil {
		renderErrorPage(w, 500)
		return
	}

	err = t.Execute(w, character)
	if err != nil {
		http.Error(w, "Error executing template", http.StatusInternalServerError)
		return
	}
}

func result(wr http.ResponseWriter, r *http.Request) {
	// For resultpage the request is always POST not GET
	if r.Method != http.MethodPost {
		renderErrorPage(wr, 405)
		return
	}

	// if url is not for result page error handle
	if r.URL.Path != "/result" {
		renderErrorPage(wr, 404)
		return
	}

	char, _ := functions.LoadData("https://groupietrackers.herokuapp.com/api/artists")
	artId := r.FormValue("artist")
	for _, ch := range artId {
		if ch != 10 && ch != 13 && (ch < 32 || ch > 126) {
			renderErrorPage(wr, 400)
			return
		}
	}
	artId = strings.Title(artId)

	if len(artId) > 2 {
		pattern := regexp.MustCompile(`:\d+`)
		artId = strings.TrimPrefix(pattern.FindString(artId), ":")
	}

	for i := 0; i < 52; i++ {
		if artId == char[i].Artist {
			artId = strconv.Itoa(char[i].ID)
		}
	}

	iint, err := strconv.Atoi(artId)
	if err != nil || iint <= 0 {
		renderErrorPage(wr, 500)
		return
	}
	i := iint - 1

	// Load artist data
	character, err := functions.LoadData("https://groupietrackers.herokuapp.com/api/artists")
	if err != nil {
		http.Error(wr, "Failed to load artist data", http.StatusInternalServerError)
		return
	}

	if len(character) == 0 {
		http.Error(wr, "No artist data available", http.StatusInternalServerError)
		return
	}

	charData, err := functions.LoadUrelles("https://groupietrackers.herokuapp.com/api/relation")
	if err != nil {
		http.Error(wr, "Failed to load relations data", http.StatusInternalServerError)
		return
	}

	if len(charData) == 0 {
		http.Error(wr, "No data available", http.StatusInternalServerError)
		return
	}

	// members := "No members available"
	// if len(character[i].Members) > 0 {
	// 	members = strings.Join(character[i].Members, ", ")
	// }

	var cdata []string
	x := 1
	d := ""
	for location, date := range charData[i].DatesLocations {
		d = strconv.Itoa(x) + ") " + strings.ReplaceAll(string(location), "-", ", ") + ": " + strings.Join(date, ", ")
		d = strings.ReplaceAll(d, "_", " ")
		cdata = append(cdata, d)
		x++
	}

	FFinal := Final{
		ID:        character[i].ID,
		Image:     character[i].Image,
		Artist:    character[i].Artist,
		Members:   functions.Members{},
		AlbumYear: character[i].AlbumYear,
		Album1:    character[i].Album1,
		Locations: cdata,
	}

	MMember := []string{}
	MMember = append(MMember, character[i].Members...)

	FFinal.Members = MMember

	t, err := template.ParseFiles("result.html")
	if err != nil {
		renderErrorPage(wr, 500)
		return
	}

	errr := t.Execute(wr, FFinal)
	if errr != nil {
		log.Println(errr)
		http.Error(wr, "Error executing template", http.StatusInternalServerError)
	}
}

func renderErrorPage(w http.ResponseWriter, code int) {

	// Set the status coded
	if code == 500 {
		w.WriteHeader(http.StatusInternalServerError)
	} else if code == 404 {
		w.WriteHeader(http.StatusNotFound)
	} else if code == 400 {
		w.WriteHeader(http.StatusBadRequest)
	}

	// Generate ASCII art for the error code with the "Standard" style

	// Parse and render the custom 404 template
	t, err := template.ParseFiles(fmt.Sprintf("style/%d.html", code))
	if err != nil {
		http.Error(w, fmt.Sprintf("Error parsing %d HTML", code), http.StatusInternalServerError)
		return
	}

	// Render the template with the result
	err = t.Execute(w, nil)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error executing %d template", code), http.StatusInternalServerError)
	}
}
