package video

import (
	"encoding/json"
	"fmt"

	"github.com/shiyuexiaoyu/dyopen/open/context"
	"github.com/shiyuexiaoyu/dyopen/util"
)

const (
	// 上传视频
	videoUploadURL string = "https://open.douyin.com/video/upload?access_token=%s&open_id=%s"
	// 初始化分片上传
	videoPartInitURL string = "https://open.douyin.com/video/part/init?access_token=%s&open_id=%s"
	// 分片上传
	videoPartUploadURL string = "https://open.douyin.com/video/part/upload?access_token=%s&open_id=%s"
	// 分片完成上传
	videoPartCompleteURL string = "https://open.douyin.com/video/part/complete?access_token=%s&open_id=%s&upload_id=%s"
	// 创建视频
	videoCreateURL string = "https://open.douyin.com/video/create?access_token=%s&open_id=%s"
	// 删除视频
	videoDeleteURL string = "https://open.douyin.com/video/delete?access_token=%s&open_id=%s"
	// 视频列表
	videoListURL string = "https://open.douyin.com/video/list?open_id=%s&cursor=%d&count=%d"
	// 视频数据
	videoDataURL string = "https://open.douyin.com/video/data?access_token=%s&open_id=%s"
)

// Video 视频
type Video struct {
	*context.Context
}

// NewVideo .
func NewVideo(context *context.Context) *Video {
	video := new(Video)
	video.Context = context
	return video
}

// Info 视频信息.
type Info struct {
	util.CommonError

	Video struct {
		VideoID string `json:"video_id"`
		Height  int64  `json:"height"`
		Width   int64  `json:"width"`
	} `json:"video"`
}

type uploadVideoRes struct {
	Message string `json:"message"`
	Data    Info   `json:"data"`
}

// Upload 视频上传.
// refer: https://open.douyin.com/platform/doc/6848798087398295555
func (video *Video) Upload(openid, accessToken string, filename string) (videoInfo Info, err error) {

	uri := fmt.Sprintf(videoUploadURL, accessToken, openid)
	var response []byte
	response, err = util.PostFile("video", filename, uri)
	if err != nil {
		return
	}

	var result uploadVideoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Upload error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	videoInfo = result.Data
	return
}

// PartInfo .
type PartInfo struct {
	util.CommonError

	UploadID string `json:"upload_id"`
}

type partInfoRes struct {
	Message string   `json:"message"`
	Data    PartInfo `json:"data"`
}

// PartInit 初始化分片上传.
// refer: https://open.douyin.com/platform/doc/6848798087398393859
func (video *Video) PartInit(openid string, accessToken string) (partInfo PartInfo, err error) {

	uri := fmt.Sprintf(videoPartInitURL, accessToken, openid)
	var response []byte
	response, err = util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	var result partInfoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("PartInit error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	partInfo = result.Data
	return
}

type partVideoRes struct {
	Message string           `json:"message"`
	Data    util.CommonError `json:"data"`
}

// PartUpload 视频分片上传.
// refer: https://open.douyin.com/platform/doc/6848798087226460172
func (video *Video) PartUpload(openid, accessToken string, uploadid string, partNumber int64, filename string) (err error) {

	uri := fmt.Sprintf(videoPartUploadURL, accessToken, openid)

	var response []byte
	response, err = util.PostFile("video", filename, uri)
	if err != nil {
		return
	}

	var result partVideoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Upload error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	return
}

// PartComplete 视频分片完成上传.
// refer: https://open.douyin.com/platform/doc/6848798087398361091
func (video *Video) PartComplete(openid, accessToken string, uploadid string) (videoInfo Info, err error) {

	uri := fmt.Sprintf(videoPartCompleteURL, accessToken, openid, uploadid)
	var response []byte
	response, err = util.HTTPPost(uri, "")
	if err != nil {
		return
	}

	var result uploadVideoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("PartComplete error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	videoInfo = result.Data
	return
}

// CreateVideoReq .
type CreateVideoReq struct {
	VideoID           string   `json:"video_id"`
	CoverTsp          float64  `json:"cover_tsp"`
	GameID            string   `json:"string"`
	PoiID             string   `json:"poi_id"`
	PoiName           string   `json:"poi_name"`
	Text              string   `json:"text"`
	MicroAppURL       string   `json:"micro_app_url"`
	MicroAppID        string   `json:"micro_app_id"`
	MicroAppTitle     string   `json:"micro_app_title"`
	AtUsers           []string `json:"at_users"`
	GameContent       string   `json:"game_content"`
	TimelinessKeyword string   `json:"timeliness_keyword"`
	TimelinessLabel   string   `json:"timeliness_label"`
	ArticleID         string   `json:"article_id"`
	ArticleTitle      string   `json:"article_title"`
}

// CreateInfo .
type CreateInfo struct {
	util.CommonError

	ItemID string `json:"item_id"`
}

type createRes struct {
	Message string     `json:"message"`
	Data    CreateInfo `json:"data"`
}

// Create 视频创建.
// refer: https://open.douyin.com/platform/doc/6848798087398328323
func (video *Video) Create(openid string, accessToken string, videoInfo *CreateVideoReq) (info CreateInfo, err error) {

	uri := fmt.Sprintf(videoCreateURL, accessToken, openid)
	var response []byte
	response, err = util.PostJSON(uri, videoInfo)
	if err != nil {
		return
	}

	var result createRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Create error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	info = result.Data
	return
}

type deleteVideoReq struct {
	ItemID string `json:"item_id"`
}

type deleteVideoRes struct {
	Message string           `json:"message"`
	Data    util.CommonError `json:"data"`
}

// Delete 视频删除
// refer: https://open.douyin.com/platform/doc/6848806536383383560#url
func (video *Video) Delete(openid, accessToken string, itemid string) (err error) {

	uri := fmt.Sprintf(videoCreateURL, accessToken, openid)

	rep := &deleteVideoReq{
		ItemID: itemid,
	}

	var response []byte
	response, err = util.PostJSON(uri, rep)
	if err != nil {
		return
	}

	var result deleteVideoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Delete error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	return
}

// ListInfo video list info.
type ListInfo struct {
	util.CommonError

	Total   int64 `json:"total"`
	Cursor  int64 `json:"cursor"`
	HasMore bool  `json:"has_more"`
	List    []struct {
		Statistics struct {
			CommentCount  int `json:"comment_count"`
			DiggCount     int `json:"digg_count"`
			DownloadCount int `json:"download_count"`
			ForwardCount  int `json:"forward_count"`
			PlayCount     int `json:"play_count"`
			ShareCount    int `json:"share_count"`
		} `json:"statistics"`
		MediaType   int    `json:"media_type"`
		ItemID      string `json:"item_id"`
		Title       string `json:"title"`
		Cover       string `json:"cover"`
		IsTop       bool   `json:"is_top"`
		CreateTime  int64  `json:"create_time"`
		IsReviewed  bool   `json:"is_reviewed"`
		VideoStatus int    `json:"video_status"`
		VideoId     string `json:"video_id"`
		ShareURL    string `json:"share_url"`
	} `json:"list"`
}

type listInfoRes struct {
	Message string   `json:"message"`
	Data    ListInfo `json:"data"`
}

// List .
func (video *Video) List(openid string, accessToken string, cursor, count int64) (info *ListInfo, err error) {
	uri := fmt.Sprintf(videoListURL, openid, cursor, count)
	var response []byte
	response, err = util.HTTPGetHeader(uri, map[string]string{
		"access-token": accessToken,
	})
	if err != nil {
		return
	}
	var result listInfoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("List error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	info = &result.Data
	return
}

// DataInfo video data info.
type DataInfo struct {
	util.CommonError

	List []struct {
		VideoStatus int `json:"video_status"` //表示视频状态。2:不适宜公开;4:审核中;5:公开视频
		Statistics  struct {
			CommentCount  int `json:"comment_count"`
			DiggCount     int `json:"digg_count"`
			DownloadCount int `json:"download_count"`
			ForwardCount  int `json:"forward_count"`
			PlayCount     int `json:"play_count"`
			ShareCount    int `json:"share_count"`
		} `json:"statistics"`
		Title      string `json:"title"`
		Cover      string `json:"cover"` //视频封面
		IsTop      bool   `json:"is_top"`
		CreateTime int64  `json:"create_time"` //视频创建时间戳
		ItemID     string `json:"item_id"`
		IsReviewed bool   `json:"is_reviewed"` //表示是否审核结束。审核通过或者失败都会返回true，审核中返回false。
		VideoId    string `json:"video_id"`
		ShareURL   string `json:"share_url"`
		MediaType  int    `json:"media_type"` // 媒体类型。2:图集;4:视频
	} `json:"list"`
}

// DataReq .
type DataReq struct {
	ItemIDS []string `json:"item_ids"`
}

type dataInfoRes struct {
	Message string   `json:"message"`
	Data    DataInfo `json:"data"`
}

// Data .
func (video *Video) Data(openid string, accessToken string, itemIDS []string) (info *DataInfo, err error) {

	uri := fmt.Sprintf(videoDataURL, accessToken, openid)
	req := &DataReq{
		ItemIDS: itemIDS,
	}

	var response []byte
	response, err = util.PostJSON(uri, req)
	if err != nil {
		return
	}
	var result dataInfoRes
	err = json.Unmarshal(response, &result)
	if err != nil {
		return
	}
	if result.Data.ErrCode != 0 {
		err = fmt.Errorf("Data error : errcode=%v , errmsg=%v", result.Data.ErrCode, result.Data.ErrMsg)
		return
	}
	info = &result.Data
	return
}
