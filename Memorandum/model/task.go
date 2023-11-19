package model

type Task struct {
	Uid       uint   `json:"uid"`
	Id        uint   `json:"id"`
	Title     string `json:"title"`
	Content   string `json:"content"`
	View      uint   `json:"view"`
	Status    int    `json:"status"`
	CreateAt  int64  `json:"create_at"`
	StartTime int64  `json:"start_time"`
	EndTime   int64  `json:"end_time"`
}

//func (Task *Task) TaskView() uint64 {
//	// 增加点击数
//	countStr, _ := cache.RedisClient.Get(cache.TaskViewKey(Task.Id)).Result()
//	count, _ := strconv.ParseUint(countStr, 10, 64)
//	return count
//}
//
//// AddView
//func (Task *Task) AddView() {
//	cache.RedisClient.Incr(cache.TaskViewKey(Task.Id))                      // 增加视频点击数
//	cache.RedisClient.ZIncrBy(cache.RankKey, 1, strconv.Itoa(int(Task.Id))) // 增加排行点击数
//}
