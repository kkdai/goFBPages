package main

import (
	"encoding/json"
	"flag"
	"fmt"
	fb "github.com/huandu/facebook"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"sync"
)

var pageName = flag.String("n", "", "Facebook page name such as: scottiepippen")
var pageId = flag.String("id", "", "Facebook page id, if you know such as: 112743018776863")
var numOfWorkersPtr = flag.Int("c", 2, "the number of concurrent rename workers. default = 2")

var m sync.Mutex
var FileIndex int = 0

func GetFileIndex() (ret int) {
	m.Lock()
	ret = FileIndex
	FileIndex = FileIndex + 1
	m.Unlock()
	return ret
}

var TOKEN string

func init() {
	cfg := LoadConfig()
	TOKEN = cfg.Token
}

func DownloadWorker(destDir string, linkChan chan string, wg *sync.WaitGroup) {
	defer wg.Done()

	for target := range linkChan {
		var imageType string
		if strings.Contains(target, ".png") {
			imageType = ".png"
		} else {
			imageType = ".jpg"
		}

		resp, err := http.Get(target)
		if err != nil {
			log.Println("Http.Get\nerror: " + err.Error() + "\ntarget: " + target)
			continue
		}
		defer resp.Body.Close()

		m, _, err := image.Decode(resp.Body)
		if err != nil {
			log.Println("image.Decode\nerror: " + err.Error() + "\ntarget: " + target)
			continue
		}

		// Ignore small images
		bounds := m.Bounds()
		if bounds.Size().X > 300 && bounds.Size().Y > 300 {
			imgInfo := fmt.Sprintf("pic%04d", GetFileIndex())
			out, err := os.Create(destDir + "/" + imgInfo + imageType)
			if err != nil {
				log.Println("os.Create\nerror: %s", err)
				continue
			}
			defer out.Close()
			if imageType == ".png" {
				png.Encode(out, m)
			} else {
				jpeg.Encode(out, m, nil)
			}
		}
	}
}

func FindPhotoByAlbum(ownerName string, albumName string, albumId string, baseDir string) {
	res2, err := fb.Get("/"+albumId+"/photos", fb.Params{
		"access_token": TOKEN,
	})
	photoRet := FBPhotos{}

	jret2, _ := json.Marshal(res2)
	err = json.Unmarshal(jret2, &photoRet)
	if err != nil {
		fmt.Println(err)
	}

	dir := fmt.Sprintf("%v/%v/%v - %v", baseDir, ownerName, albumId, albumName)
	os.MkdirAll(dir, 0755)

	linkChan := make(chan string)
	wg := new(sync.WaitGroup)
	for i := 0; i < 1; i++ {
		wg.Add(1)
		go DownloadWorker(dir, linkChan, wg)
	}

	for _, v := range photoRet.Data {
		linkChan <- v.Source
	}

}

func main() {
	flag.Parse()
	var inputPage string
	if *pageName == "" && *pageId == "" {
		log.Fatalln("You need to input either -n=page or -id=pageid.")
	} else if *pageName != "" && *pageId != "" {
		log.Fatalln("You can only input either -n=page or -id=pageid.")
	} else if *pageName != "" {
		inputPage = *pageName
	} else {
		inputPage = *pageId
	}

	usr, _ := user.Current()
	baseDir := fmt.Sprintf("%v/Pictures/goFBPages", usr.HomeDir)
	res, err := fb.Get("/"+inputPage+"/albums", fb.Params{
		"access_token": TOKEN,
	})

	if err != nil {
		log.Fatalln("FB connect error, err=", err.Error())
	}

	albumRet := FBAlbums{}
	jret, _ := json.Marshal(res)
	err = json.Unmarshal(jret, &albumRet)
	if err != nil {
		fmt.Println(err)
	}

	for _, v := range albumRet.Data {
		fmt.Println("Starting downloading " + v.Name + "-" + v.From.Name)
		FindPhotoByAlbum(v.From.Name, v.Name, v.ID, baseDir)
	}
}
