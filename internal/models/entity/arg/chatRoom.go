package arg

type SplashArg struct {
	RoomId int `json:"room_id" form:"room_id" binding:"-"`
	Uid    int `json:"uid" form:"uid" binding:"-"`
}

type DeleteTopListArg struct {
	Uid int `json:"uid" form:"uid" binding:"required"`
}

type AddTopListArg struct {
	Uid         int `json:"uid" form:"uid" binding:"required"`
	IsTop       int `json:"is_top" form:"is_top" binding:"-"`
	MemberCount int `json:"member_count" form:"member_count" binding:"required"`
}
