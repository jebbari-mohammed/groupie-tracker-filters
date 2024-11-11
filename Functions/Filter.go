package functions

import (
	"net/http"
	"strconv"
	"strings"
	"text/template"
)

func Filter(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		ErrorHandler(w, r, 405)
		return
	}
	err, errr := Fitch_Global(w, r, Url_Artists), Fitch_Global(w, r, Url_Locations)
	if err != nil || errr != nil {
		ErrorHandler(w, r, 500)
		return
	}
	var sum_cd []int
	max := r.FormValue("id")
	min := r.FormValue("id2")
	max_int, _ := strconv.Atoi(max)
	min_int, _ := strconv.Atoi(min)
	if min_int > max_int {
		ErrorHandler(w,r,400)
		return
	}

	for _, artist := range Artists {
		if min_int <= artist.CreationDate && max_int >= artist.CreationDate {
			sum_cd = append(sum_cd, artist.Id)
		}
	}

	var sum_FA []int
	max_FA := r.FormValue("id3")
	min_FA := r.FormValue("id4")
	max_FA_Int, _ := strconv.Atoi(max_FA)
	min_FA_Int, _ := strconv.Atoi(min_FA)
	if min_FA_Int> max_FA_Int {
		ErrorHandler(w,r,400)
		return
	}


	for _, artist := range Artists {
		f, _ := strconv.Atoi(artist.FirstAlbum[6:])

		if min_FA_Int <= f && max_FA_Int >= f {
			sum_FA = append(sum_FA, artist.Id)
		}

	}

	c1 := r.Form["c1"]

	var mem []int
	for _, conv := range c1 {
		m, _ := strconv.Atoi(conv)
		if m >= 1 && m <= 8 {
			mem = append(mem, m)
		} else {
			ErrorHandler(w, r, 404)
			return
		}
	}
	sum_mem := []int{}

	for _, artist := range Artists {
		for _, z := range mem {
			if z == len(artist.Members) {
				sum_mem = append(sum_mem, artist.Id)
			}
		}
	}

	filtre_loc := strings.ReplaceAll(r.FormValue("filter"), ", ", "-")
	var sum_loc []int

	fond := map[int]bool{}

	for _, artist := range Artists {
		for _, location := range Locations.Index {
			for _, loc := range location.Locatins {

				if filtre_loc == "" {
					continue
				}

				if artist.Id == location.Id && strings.Contains(strings.ToLower(loc), strings.ToLower(filtre_loc)) && !fond[artist.Id] {
					sum_loc = append(sum_loc, artist.Id)
					fond[artist.Id] = true

				}

			}
		}
	}

	if len(sum_loc) == 0 && len(filtre_loc) != 0 {
		ErrorHandler(w, r, 404)
		return
	}

	var filters []int

	for _, cd := range sum_cd {
		for _, FA := range sum_FA {
			if FA == cd {
				filters = append(filters, FA)
			}
		}
	}
	var xx []int

	if len(sum_mem) != 0 {
		for _, f := range filters {
			for _, m := range sum_mem {
				if f == m {
					xx = append(xx, m)
				}
			}
		}
	} else {
		xx = filters
	}

	var xx2 []int

	if len(sum_loc) != 0 {
		
			for _, f := range xx {
				for _, m := range sum_loc {
					if f == m {
						xx2 = append(xx2, m)
					}
				}
			} 
			} else {
		xx2 = xx
	}
	if len(xx2) == 0 {
		ErrorHandler(w, r, 404)
		return
	}

	var artt []Artist
	for _, artist := range Artists {
		for _, x := range xx2 {
			if x == artist.Id {
				artt = append(artt, artist)
			}
		}
	}

	tmpl, err := template.ParseFiles("Template/filter.html")
	if err != nil {
		ErrorHandler(w, r, 500)
		return
	}
	err = tmpl.Execute(w, artt)
	if err != nil {
		ErrorHandler(w, r, 500)
		return
	}
}
