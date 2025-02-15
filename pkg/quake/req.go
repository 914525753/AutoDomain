package quake

import (
	"fmt"
	"github.com/shadowabi/AutoDomain_rebuild/config"
	"github.com/shadowabi/AutoDomain_rebuild/define"
	net2 "github.com/shadowabi/AutoDomain_rebuild/utils/response"
	"net/http"
	"strings"
	"time"
)

func QuakeRequest(client *http.Client, page int, reqStringList ...string) (respBody []string) {
	if len(reqStringList) != 0 {
		for _, reqString := range reqStringList {
			data := strings.NewReader(fmt.Sprintf("query=%s&start=%v&size=100&include=service.name&include=port&include=service.http.host", reqString, page))
			req, _ := http.NewRequest("POST", "https://quake.360.net/api/v3/search/quake_service", data)
			req.Header.Set("User-Agent", define.UserAgent)
			req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
			req.Header.Set("X-QuakeToken", config.C.QuakeKey)

			resp, err := client.Do(req)
			time.Sleep(500 * time.Millisecond)
			if err != nil {
				continue
			}
			respBody = append(respBody, net2.HandleResponse(resp))
			resp.Body.Close()
		}
	}
	return respBody
}
