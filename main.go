package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"log"
	"net/http"
	"os"
	"os/user"
	"strings"
	"sync"

	fb "github.com/huandu/facebook"
)

var pageName = flag.String("n", "", "Facebook page name such as: scottiepippen")
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
	TOKEN = os.Getenv("FBTOKEN")
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
	photoRet := FBPhotos{}
	resPhoto := RunFBGraphAPI("/" + albumId + "/photos")
	ParseMapToStruct(resPhoto, &photoRet)
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

func ParseMapToStruct(inData interface{}, decodeStruct interface{}) {
	jret, _ := json.Marshal(inData)
	err := json.Unmarshal(jret, &decodeStruct)
	if err != nil {
		log.Fatalln(err)
	}
}

func RunFBGraphAPI(query string) (queryResult interface{}) {
	res, err := fb.Get(query, fb.Params{
		"access_token": TOKEN,
	})

	if err != nil {
		log.Fatalln("FB connect error, err=", err.Error())
	}
	return res
}

func main() {
	flag.Parse()
	var inputPage string
	if TOKEN == "" {
		log.Fatalln("Set your FB token as environment variables 'export FBTOKEN=XXXXXXX'")
	}

	if *pageName == "" {
		log.Fatalln("You need to input -n=Name_or_Id.")
	}
	inputPage = *pageName

	//Get system user folder
	usr, _ := user.Current()
	baseDir := fmt.Sprintf("%v/Pictures/goFBPages", usr.HomeDir)

	//Get User info
	resUser := RunFBGraphAPI("/" + inputPage)
	userRet := FBUser{}
	ParseMapToStruct(resUser, &userRet)

	//Get all albums
	resAlbums := RunFBGraphAPI("/" + inputPage + "/albums")
	albumRet := FBAlbums{}
	ParseMapToStruct(resAlbums, &albumRet)

	userFolderName := fmt.Sprintf("[%s]%s", userRet.Username, userRet.Name)
	for _, v := range albumRet.Data {
		fmt.Println("Starting download [" + v.Name + "]-" + v.From.Name)
		FindPhotoByAlbum(userFolderName, v.Name, v.ID, baseDir)
	}
}
