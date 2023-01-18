package src

type PublicInf struct {
	Name          string `json:"name"`
	Id            int    `json:"id"`
	FollowCount   int    `json:"follow_count"`
	FollowerCount int    `json:"follower_count"`
}

type User struct {
	PublicInf
	FollowList   []User
	FollowerList []User
	LikeList     []Video
	PostList     []Video
}

type Video struct {
	Id int `json:"id"`
	Title  string `json:"title"`
	Author PublicInf
	PlayUrl string `json:"play_url"`
	LikeCount int `json:"favorite_count"`
	CommentCount int `json:"comment_count"`
}

type Comment struct {
	VideoBelong int `json:"-"`  //编码JSON时忽略该字段
	Id int `json:"id"`
	Author PublicInf
	Content string  `json:"content"`
	CreateDate string `json:"create_date"`
}

type Message struct {
	ToUser int `json:"-"`
	FromUser int `json:"-"`
	Content string `json:"content"`
	CreateTime string `json:"create_time"`
}

type RespHead struct {
	StatusCode int `json:"status_code"`
	StatusMsg string `json:"status_msg"`
}