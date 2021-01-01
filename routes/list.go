package routes

import (
	b64 "encoding/base64"
	"fmt"
	"image/jpeg"
	"io/ioutil"
	"net/http"
	"os"
	"path/filepath"
	"strings"

	"github.com/labstack/echo/v4"

	"github.com/anvari1313/photopic-server/config"
)

type Album struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Thumbnail string `json:"thumbnail"`
}

type ImageType string

const (
	ImageTypeImage ImageType = "image"
	VideoTypeImage ImageType = "video"
)

type Image struct {
	Name      string    `json:"name"`
	Thumbnail string    `json:"thumbnail"`
	URL       string    `json:"url"`
	Type      ImageType `json:"type"`
	Height    int       `json:"height"`
	Width     int       `json:"width"`
}

type List struct {
	Path   string  `json:"path"`
	Albums []Album `json:"albums"`
	Images []Image `json:"images"`
}

func iterate(path string) ([]Album, []Image, error) {
	directories := make([]Album, 0)
	files := make([]Image, 0)

	fileInfos, err := ioutil.ReadDir(path)
	if err != nil {
		return nil, nil, err
	}

	for _, fileInfo := range fileInfos {
		if fileInfo.IsDir() {
			directories = append(directories, Album{
				Name:      fileInfo.Name(),
				Path:      b64.URLEncoding.EncodeToString([]byte(path + fileInfo.Name() + "/")),
				Thumbnail: "thumbnail",
			})
		} else {
			switch filepath.Ext(fileInfo.Name()) {
			case ".png":
			case ".jpg":
				filePath := path + fileInfo.Name()
				file, err := os.Open(filePath)
				if err != nil {
					return nil, nil, err
				}
				img, err := jpeg.DecodeConfig(file)
				if err != nil {
					return nil, nil, fmt.Errorf("file %s with extension %s error: %w", filePath, filepath.Ext(filePath), err)
				}

				fmt.Println(filePath)

				files = append(files, Image{
					Name:      fileInfo.Name(),
					Thumbnail: config.C.URLPrefix + "/thumbnail" + strings.ReplaceAll(filePath, config.C.BasePath, ""),
					URL:       config.C.URLPrefix + "/static" + strings.ReplaceAll(filePath, config.C.BasePath, ""),
					Height:    img.Height,
					Width:     img.Width,
					Type:      ImageTypeImage,
				})
			default:
				fmt.Printf("extension `%s` is not known\n", filepath.Ext(fileInfo.Name()))
			}
		}
	}

	return directories, files, nil
}

func ListHandler(ctx echo.Context) error {
	path := ctx.QueryParam("path")
	if path == "" {
		p := b64.URLEncoding.EncodeToString([]byte(config.C.BasePath + "/"))
		return ctx.Redirect(http.StatusTemporaryRedirect, "/list?path="+p)
	}

	decPath, err := b64.URLEncoding.DecodeString(path)
	if err != nil {
		return echo.ErrNotFound
	}

	path = string(decPath)
	p := path
	dirs, files, err := iterate(p)
	if err != nil {
		return err
	}

	return ctx.JSON(http.StatusOK, List{
		Path:   strings.ReplaceAll(path, config.C.BasePath, ""),
		Albums: dirs,
		Images: files,
	})
}
