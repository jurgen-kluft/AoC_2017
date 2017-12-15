package download

import "os"
import "fmt"
import "net/http"

import "github.com/levigross/grequests"

// Exists reports whether the named file or directory exists.
func Exists(name string) bool {
	if _, err := os.Stat(name); err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

// Get will download a file from URL to filename
func get(URL string, filename string) {
	if Exists(filename) == false {
		ro := &grequests.RequestOptions{
			Cookies: []*http.Cookie{
				{
					Name:     "_ga",
					Value:    "GA1.2.291960124.1512269074",
					HttpOnly: true,
					Secure:   false,
				}, {
					Name:     "session",
					Value:    "53616c7465645f5fbd6db5327319842f6476b01d3c729cf34953166ea746ac2ae982b6c2f72313def34f592e40b1b636",
					HttpOnly: true,
					Secure:   false,
				}, {
					Name:     "_gid",
					Value:    "GA1.2.2028937355.1513227847",
					HttpOnly: true,
					Secure:   false,
				},
			},
		}

		resp, err := grequests.Get(URL, ro)
		if err != nil {
			fmt.Printf("%v\n", err.Error())
		} else if resp.Error != nil {
			fmt.Printf("%v\n", resp.Error.Error())
		} else {
			if resp.StatusCode != 404 && resp.StatusCode != 500 {
				resp.DownloadToFile(filename)
			} else {
				fmt.Printf("Bad HTTP response %v\n", resp.StatusCode)
			}

		}
	}
}

func GetInputForDay(day int) {
	URL := fmt.Sprintf("http://adventofcode.com/2017/day/%d/input?_ga=A,session=A,_gid=A", day)
	filename := fmt.Sprintf("day%d/input.text", day)
	get(URL, filename)

}
