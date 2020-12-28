package routes

import (
	b64 "encoding/base64"
	"net/http"

	"github.com/labstack/echo/v4"
)


type Album struct {
	Name      string `json:"name"`
	Path      string `json:"path"`
	Thumbnail string `json:"thumbnail"`
}

type Image struct {
	Name      string `json:"name"`
	Thumbnail string `json:"thumbnail"`
	Height    uint64 `json:"height"`
	Width     uint64 `json:"width"`
}

type List struct {
	Path   string  `json:"path"`
	Albums []Album `json:"albums"`
	Images []Image `json:"images"`
}

func ListHandler(ctx echo.Context) error {
	path := ctx.QueryParam("path")
	if path == "" {
		p := b64.URLEncoding.EncodeToString([]byte("/"))
		return ctx.Redirect(http.StatusTemporaryRedirect, "/list?path="+p)
	}

	sDec, err := b64.URLEncoding.DecodeString(path)
	if err != nil {
		return echo.ErrNotFound
	}

	return ctx.JSON(http.StatusOK, List{
		Path: string(sDec),
		Albums: []Album{
			{
				Name:      "album1",
				Path:      b64.URLEncoding.EncodeToString([]byte("/album1")),
				Thumbnail: "http://localhost:8080/thumbnail.jpg",
			}, {
				Name:      "album2",
				Path:      b64.URLEncoding.EncodeToString([]byte("/album2")),
				Thumbnail: "http://localhost:8080/thumbnail.jpg",
			},
		},
		Images: []Image{
			{
				Name:      "image1.jpg",
				Thumbnail: "http://localhost:8080/image1.jpg",
			},
		},
	})
}