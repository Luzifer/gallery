package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"mime"
	"os"
	"path"
	"strconv"

	"github.com/Luzifer/rconfig"
	"github.com/cheggaaa/pb"
	log "github.com/sirupsen/logrus"
)

var (
	cfg = struct {
		AlbumID        int    `flag:"album-id" default:"" description:"ID of the album to import to (treated as sorting key!)" validate:"nonzero"`
		AlbumTitle     string `flag:"album-title" default:"" description:"Title of the album"`
		AllowExisting  bool   `flag:"allow-existing" default:"false" description:"Allow inserting into existing gallery"`
		FullSize       string `flag:"full-size" default:"1920x1080" description:"Images will get resized to fit into this without upscaling"`
		LogLevel       string `flag:"log-level" default:"info" description:"Log level (debug, info, warn, error, fatal)"`
		Storage        string `flag:"storage" default:"file://./output" description:"Storage to write results to"`
		ThumbnailSize  string `flag:"thumbnail-size" default:"500x500" description:"Size of the thumbnail to generate"`
		UpdateFrontend bool   `flag:"update-frontend" default:"false" description:"Upload interface files"`
		VersionAndExit bool   `flag:"version" default:"false" description:"Prints current version and exits"`
	}{}

	version = "dev"
)

func init() {
	if err := rconfig.ParseAndValidate(&cfg); err != nil {
		log.Fatalf("Unable to parse commandline options: %s", err)
	}

	if cfg.VersionAndExit {
		fmt.Printf("gallery %s\n", version)
		os.Exit(0)
	}

	if l, err := log.ParseLevel(cfg.LogLevel); err != nil {
		log.WithError(err).Fatal("Unable to parse log level")
	} else {
		log.SetLevel(l)
	}
}

func main() {
	args := rconfig.Args()[1:]
	if len(args) == 0 {
		fmt.Println("Usage: gallery [flags] <folder to import>")
		os.Exit(1)
	}

	store, err := getStorageByURI(cfg.Storage)
	if err != nil {
		log.WithError(err).Fatal("Could not open storage")
	}

	if cfg.UpdateFrontend {
		if err := uploadFrontend(store); err != nil {
			log.WithError(err).Fatal("Could not upload frontend files")
		}
	}

	albumList, err := loadAlbumsFromStorage(store)
	switch err {
	case nil:
	case errFileNotFound:
		albumList = &albums{}
	default:
		log.WithError(err).Fatal("Unable to load album list")
	}

	if albumList.HasGallery(cfg.AlbumID) && !cfg.AllowExisting {
		log.Fatal("Gallery with this ID already exists and allow-existing is not set")
	}

	imagePath := rconfig.Args()[1]

	files, err := ioutil.ReadDir(imagePath)
	if err != nil {
		log.WithError(err).Fatal("Could not list files in given directory")
	}

	bar := pb.StartNew(len(files))
	for _, file := range files {
		imgSrcPath := path.Join(imagePath, file.Name())
		flog := log.WithFields(log.Fields{
			"path": imgSrcPath,
		})

		thumb, full, err := scaleImage(imgSrcPath)
		if err != nil {
			flog.WithError(err).Error("Unable to resize image")
			continue
		}

		if err := store.WriteFile(
			path.Join("photos", strconv.Itoa(cfg.AlbumID), file.Name()+".thumb"),
			thumb,
			"image/jpeg",
		); err != nil {
			flog.WithError(err).Error("Thumbnail upload failed")
			continue
		}

		if err := store.WriteFile(
			path.Join("photos", strconv.Itoa(cfg.AlbumID), file.Name()),
			full,
			"image/jpeg",
		); err != nil {
			flog.WithError(err).Error("FullSize upload failed")
			continue
		}

		if err := albumList.AddImage(file.Name(), ""); err != nil { // FIXME (kahlers): Title?
			flog.WithError(err).Error("Could not update album")
			continue
		}

		bar.Increment()
	}
	bar.Finish()

	if err := albumList.SaveToStorage(store); err != nil {
		log.WithError(err).Fatal("Could not write album list to storage")
	}
}

func uploadFrontend(s storage) error {
	names, err := AssetDir("frontend")
	if err != nil {
		return err
	}

	for _, name := range names {
		if err := s.WriteFile(
			name,
			bytes.NewReader(MustAsset(path.Join("frontend", name))),
			mime.TypeByExtension(path.Ext(name)),
		); err != nil {
			return err
		}
	}

	return nil
}
