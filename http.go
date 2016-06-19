package main
import (
	"gopkg.in/urfave/cli.v2"
	"github.com/gorilla/mux"
	"net/http"
	"fmt"
	"path/filepath"
	"strings"
	"github.com/leonlau/initialser"
	"strconv"
	"errors"
//	"io"
//	"bytes"
	"io/ioutil"
)

var cmdHttp = &cli.Command{
	Name:"http",
	Usage:"start a http server",
	Action:runHttp,
}

const fileNamePathKey = "file_name"


func runHttp(c *cli.Context) error {

	//addr := fmt.Sprintf(":%d", c.Int("port"));
	addr := fmt.Sprintf(":%d", 9527);
	r := mux.NewRouter()
	r.HandleFunc("/", homeHandler);
	r.HandleFunc(fmt.Sprintf("/{%s}", fileNamePathKey), avatarHandler);
	println("app start ", addr)

	//append font path
	p, _ := filepath.Abs(".")
	initialser.AppendFontPath(p)

	return http.ListenAndServe(addr, r)

}




func homeHandler(w http.ResponseWriter, req *http.Request) {
	data, _ := ioutil.ReadFile("index.html")
	w.Write(data)
}
//avatarHandler server avatar
func avatarHandler(w http.ResponseWriter, req *http.Request) {
	// parse path name
	text, ext := parseFileName(req);
	encode := "png"
	switch ext {
	case ".ico":
	case ".svg":
		w.Header().Set("Content-Type", "image/svg+xml")
		setCacheControl(w)
		fmt.Fprint(w, initialser.NewAvatar(text).Svg())
		return;
	case ".jpg", ".jpeg":
		encode = "jpg"
		setCacheControl(w)
		w.Header().Set("Content-Type", "image/jpeg")
	case ".png", "":
		setCacheControl(w)
		w.Header().Set("Content-Type", "image/png")
	default:
		badReq(w, errors.New("not support ext " + ext))
		return;
	}
	avatar := initialser.NewAvatar(text)
	// parse query param to avatar
	err := parseParamTo(avatar, req);
	println(avatar.Key())
	if badReq(w, err) {
		return;
	}
	d, err := initialser.NewDrawer(avatar)
	if !badReq(w, err) {
		badReq(w, d.DrawToWriter(w, encode))
	}
}
//parseFileName parse url file name
func parseFileName(req *http.Request) (title string, ext string) {
	fileName := mux.Vars(req)[fileNamePathKey]
	ext = filepath.Ext(fileName)
	ext = strings.ToLower(ext)
	title = strings.TrimSuffix(fileName, ext)
	return
}


func setCacheControl(w http.ResponseWriter) {
	w.Header().Set("Cache-Control", "public, max-age=600")
}
//parseParam  ?bg=#dd00ff&s=200&f=宋体&fs=120&c=#020319
func parseParamTo(avatar *initialser.Avatar, req *http.Request) error {
	q := req.URL.Query()
	avatar.Font = ifBlankDefault(q.Get("f"), "Hiragino_Sans_GB_W3")
	avatar.Color = ifBlankDefault(q.Get("c"), avatar.Color)
	avatar.Background = ifBlankDefault(q.Get("bg"), avatar.Background)
	if q.Get("s") != "" {
		if size, err := strconv.Atoi(q.Get("s")); err != nil {
			return errors.New("s is not a valid int number")
		}else {
			avatar.Size = size
		}
	}
	if q.Get("fs") != "" {
		if fs, err := strconv.Atoi(q.Get("fs")); err != nil {
			return errors.New("fs is not a valid int number")
		}else {
			avatar.FontSize = fs
		}
	}
	return nil
}
//ifBlankDefault default value
func ifBlankDefault(str string, defStr string) string {
	if str == "" {
		return defStr
	}
	return str
}

//badReq err exist ,return bad request
func badReq(w http.ResponseWriter, err error) bool {
	if err == nil {
		return false
	}
	w.Header().Set("Content-Type", "text/plain")
	w.WriteHeader(http.StatusBadRequest)
	w.Write([]byte(err.Error()))
	return true
}



