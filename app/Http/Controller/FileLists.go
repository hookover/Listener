package Controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/patrickmn/go-cache"
	"github.com/spf13/viper"
	"io/ioutil"
	"path"
	"strings"
	"time"
)

var Cache = cache.New(2*time.Minute, 5*time.Minute)

func GetBookList(c *gin.Context) {
	var mp3s []*Mp3Info

	if cacheData, found := Cache.Get("mp3s"); found {
		mp3s = cacheData.([]*Mp3Info)
	}
	if len(mp3s) == 0 {
		mp3s = getMp3Files()
		Cache.Set("mp3s", mp3s, 5*time.Minute)
	}

	c.JSON(200, gin.H{
		"code":    200,
		"message": "success",
		"data":    mp3s,
	})
}

type Mp3Info struct {
	Name string
	Path string
}

type FilePathList struct {
	data []*Mp3Info
}

func (list *FilePathList) Add(mp3 *Mp3Info) {
	list.data = append(list.data, mp3)
}
func (list *FilePathList) Get() []*Mp3Info {
	return list.data
}

func (list FilePathList) New() FilePathList {
	list.data = make([]*Mp3Info, 0, 10)
	return list
}

func getMp3Files() []*Mp3Info {
	data := &FilePathList{data: make([]*Mp3Info, 0, 10)}
	deepPath(fmt.Sprintf("%s/store", viper.Get("public")), data)

	for i, v := range data.data {
		data.data[i].Path = fmt.Sprintf("%s%s", "public/store", strings.Split(v.Path, "public/store")[1])
	}

	return data.Get()
}
func deepPath(pathName string, list *FilePathList) {
	dirList, err := ioutil.ReadDir(pathName)
	if err != nil {
		return
	}
	for _, v := range dirList {
		if v.IsDir() {
			deepPath(fmt.Sprintf("%s/%s", pathName, v.Name()), list)
		} else {
			if strings.ToLower(path.Ext(v.Name())) == ".mp3" {
				list.Add(&Mp3Info{Name: fmt.Sprintf("%s", strings.Split(v.Name(), ".")[0]), Path: fmt.Sprintf("%s/%s", pathName, v.Name())})
			}
		}
	}
}
