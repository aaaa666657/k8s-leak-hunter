package event

import (
	"fmt"
	"net/http"
	"os"
)

var TG_Access_Token, _ = os.LookupEnv("TG_ACCESS_TOKEN")
var TG_Chat_ID, _ = os.LookupEnv("TG_CHAT_ID")

func SendNotify(message string) error {
	url := "https://api.telegram.org/bot" + TG_Access_Token + "/sendMessage"
	req, _ := http.NewRequest("GET", url, nil)
	q := req.URL.Query()
	q.Add("chat_id", TG_Chat_ID)
	q.Add("text", message)
	req.URL.RawQuery = q.Encode()
	fmt.Println(req.URL.String())
	//var resp *http.Response
	resp, _ := http.DefaultClient.Do(req)
	defer resp.Body.Close()

	return nil
}
