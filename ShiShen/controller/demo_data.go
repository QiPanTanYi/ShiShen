package controller

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

var DemoComments = []Comment{
	{
		Id:         1,
		User:       DemoUser,
		Content:    "Test Comment",
		CreateDate: "05-01",
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
