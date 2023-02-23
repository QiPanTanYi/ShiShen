package main

import (
	"bytes"
	"encoding/json"

	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

type Response struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg,omitempty"`
}
type User struct {
	Avatar          string `json:"avatar"`                   // 用户头像
	BackgroundImage string `json:"background_image"`         // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`           // 喜欢数
	Id              int64  `json:"id,omitempty"`             // 用户id
	Name            string `json:"name,omitempty"`           // 用户名称
	FollowCount     int64  `json:"follow_count,omitempty"`   // 关注总数
	FollowerCount   int64  `json:"follower_count,omitempty"` // 粉丝总数
	IsFollow        bool   `json:"is_follow,omitempty"`      // true-已关注，false-未关注
	Signature       string `json:"signature"`                // 个人简介
	TotalFavorited  string `json:"total_favorited"`          // 获赞数量
	WorkCount       int64  `json:"work_count"`               // 作品数
}
type Video struct {
	Id            int64  `json:"id,omitempty"`                       // 视频唯一标识
	Author        User   `json:"author"`                             // 视频作者信息
	PlayUrl       string `json:"play_url" json:"play_url,omitempty"` // 视频播放地址
	CoverUrl      string `json:"cover_url,omitempty"`                // 视频封面地址
	FavoriteCount int64  `json:"favorite_count,omitempty"`           // 视频的点赞总数
	CommentCount  int64  `json:"comment_count,omitempty"`            // 视频的评论总数
	IsFavorite    bool   `json:"is_favorite,omitempty"`              // true-已点赞，false-未点赞
	Title         string `json:"title"`                              // 视频标题
}

type FeedResponse struct {
	Response
	VideoList []Video `json:"video_list,omitempty"`
	NextTime  int64   `json:"next_time,omitempty"`
}

var DemoVideos = []Video{
	{
		Id:            1,
		Author:        DemoUser,
		PlayUrl:       "https://www.w3schools.com/html/movie.mp4",
		CoverUrl:      "https://cdn.pixabay.com/photo/2016/03/27/18/10/bear-1283347_1280.jpg",
		FavoriteCount: 0,
		CommentCount:  0,
		IsFavorite:    false,
		Title:         "666",
	},
}

var DemoUser = User{
	Avatar:          "http://175.178.223.138:99/u1.jpg",
	BackgroundImage: "http://175.178.223.138:99/u1.jpg",
	FavoriteCount:   0,
	Id:              1,
	Name:            "TestUser",
	FollowCount:     0,
	FollowerCount:   0,
	IsFollow:        false,
	Signature:       "选择大于努力",
	TotalFavorited:  "0",
	WorkCount:       1,
}

func main() {

	url := "github.com/RaymondCode/simple-demo/douyin/feed/"
	method := "POST"
	data := FeedResponse{
		Response:  Response{StatusCode: 0},
		VideoList: DemoVideos,
		NextTime:  time.Now().Unix(),
	}
	jsonStr, _ := json.Marshal(data)
	client := &http.Client{}
	req, err := http.NewRequest(method, url, bytes.NewBuffer(jsonStr))

	if err != nil {
		fmt.Println(err)
		return
	}
	req.Header.Add("User-Agent", "Apifox/1.0.0 (https://www.apifox.cn)")

	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
