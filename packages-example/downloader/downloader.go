package downloader

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type InfoRequest struct {
	Bvids []string
}
type InfoResponse struct {
	Infos []VideoInfo
}

type VideoInfo struct {
	Code int `json:"code"`
	Data struct {
		Bvid  string `json:"bvid"`
		Title string `json:"title"`
		Desc  string `json:"desc"`
	} `json:"data"`
}

func BatchDownloadVideoInfo(request InfoRequest) (InfoResponse, error) {
	var response InfoResponse

	for _, bvid := range request.Bvids {
		var videoInfo VideoInfo

		resp, err := http.Get("http://api.bilibili.com/x/web-interface/view?bvid=" + bvid)
		if err != nil {
			return InfoResponse{}, err
		}
		respBytes, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			return InfoResponse{}, err
		}

		if err = json.Unmarshal(respBytes, &videoInfo); err != nil {
			return InfoResponse{}, err
		}

		if err = resp.Body.Close(); err != nil {
			return InfoResponse{}, err
		}
		response.Infos = append(response.Infos, videoInfo)

	}
	return response, nil
}
