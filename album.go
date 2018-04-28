package main

import (
	"bytes"
	"encoding/json"
	"path"
	"strconv"
)

type albums struct {
	Albums []album `json:"albums"`

	store storage
}

type album struct {
	ID     int          `json:"id"`
	Images []albumImage `json:"images"`
	Title  string       `json:"title"`
}

type albumImage struct {
	FullSize  string `json:"fullsize"`
	Thumbnail string `json:"thumbnail"`
	Title     string `json:"title"`
}

func loadAlbumsFromStorage(s storage) (*albums, error) {
	tmp := &albums{}
	rc, err := s.ReadFile("albums.json")
	if err != nil {
		return nil, err
	}
	defer rc.Close()

	if err := json.NewDecoder(rc).Decode(tmp); err != nil {
		return nil, err
	}
	tmp.store = s

	return tmp, nil
}

func (a albums) SaveToStorage(s storage) error {
	buf := new(bytes.Buffer)
	if err := json.NewEncoder(buf).Encode(a); err != nil {
		return err
	}

	return s.WriteFile("albums.json", buf, "application/json")
}

func (a *albums) AddImage(name, title string) error {
	alb := album{}
	alblist := []album{}

	for _, al := range a.Albums {
		if al.ID == cfg.AlbumID {
			alb = al
			continue
		}
		alblist = append(alblist, al)
	}

	if alb.ID == 0 {
		// Newly initialized album
		alb.ID = cfg.AlbumID
	}

	// Overwrite title with new one
	alb.Title = cfg.AlbumTitle

	var duplicate bool
	for _, img := range alb.Images {
		if img.FullSize == path.Join("photos", strconv.Itoa(alb.ID), name) {
			// Duplicate, ignore
			duplicate = true
			break
		}
	}
	if duplicate {
		return nil
	}

	alb.Images = append(alb.Images, albumImage{
		FullSize:  path.Join("photos", strconv.Itoa(alb.ID), name),
		Thumbnail: path.Join("photos", strconv.Itoa(alb.ID), name+".thumb"),
		Title:     title,
	})

	alblist = append(alblist, alb)
	a.Albums = alblist

	return nil
}

func (a albums) HasGallery(id int) bool {
	for _, al := range a.Albums {
		if al.ID == id {
			return true
		}
	}

	return false
}
