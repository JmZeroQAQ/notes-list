package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Note struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Content string `json:"content"`
	Created_time string `json:"created_time"`
	Modified_time string `json:"modified_time"`
}

type Resp struct {
	Message string `json:"message"`
	Note *Note `json:"note"`
}

var notes  = []Note{
	{ID: "1", Title: "原神", Content: "你说的对，但是《原神》是由米哈游自主研发的一款全新开放世界冒险游戏。游戏发生在一个被称作「提瓦特」的幻想世界，在这里，被神选中的人将被授予「神之眼」，导引元素之力。", Created_time: "2023-05-15 14:30:45", Modified_time: "2025-05-15 14:30:45"},
	{ID: "2", Title: "喵喵喵", Content: "喵喵喵 喵喵 喵喵喵喵 喵喵喵？ 喵喵喵！", Created_time: "2024-06-11 11:00:45", Modified_time: "2024-07-11 10:00:45"},
	{ID: "3", Title: "原神", Content: "你说的对，但是《原神》是由米哈游自主研发的一款全新开放世界冒险游戏。游戏发生在一个被称作「提瓦特」的幻想世界，在这里，被神选中的人将被授予「神之眼」，导引元素之力。", Created_time: "2023-05-15 14:30:45", Modified_time: "2024-05-15 14:30:45"},
	{ID: "4", Title: "喵喵喵", Content: "喵喵喵 喵喵 喵喵喵喵 喵喵喵？ 喵喵喵！", Created_time: "2024-06-11 11:00:45", Modified_time: "2024-09-11 12:00:45"},
  {ID: "5", Title: "原神", Content: "你说的对，但是《原神》是由米哈游自主研发的一款全新开放世界冒险游戏。游戏发生在一个被称作「提瓦特」的幻想世界，在这里，被神选中的人将被授予「神之眼」，导引元素之力。", Created_time: "2023-05-15 14:30:45", Modified_time: "2025-02-15 14:30:45"},
	{ID: "6", Title: "喵喵喵", Content: "喵喵喵 喵喵 喵喵喵喵 喵喵喵？ 喵喵喵！", Created_time: "2024-06-11 11:00:45", Modified_time: "2024-06-13 11:00:45"},
  {ID: "7", Title: "原神", Content: "你说的对，但是《原神》是由米哈游自主研发的一款全新开放世界冒险游戏。游戏发生在一个被称作「提瓦特」的幻想世界，在这里，被神选中的人将被授予「神之眼」，导引元素之力。", Created_time: "2023-05-15 14:30:45", Modified_time: "2025-02-15 15:30:45"},
  {ID: "8", Title: "喵喵喵", Content: "喵喵喵 喵喵 喵喵喵喵 喵喵喵？ 喵喵喵！", Created_time: "2024-06-11 11:00:45", Modified_time: "2024-06-11 11:00:45"},
  {ID: "9", Title: "原神", Content: "你说的对，但是《原神》是由米哈游自主研发的一款全新开放世界冒险游戏。游戏发生在一个被称作「提瓦特」的幻想世界，在这里，被神选中的人将被授予「神之眼」，导引元素之力。", Created_time: "2023-05-15 14:30:45", Modified_time: "2023-05-15 14:30:45"},
  {ID: "10", Title: "这是笔记标题", Content: "我是笔记内容，接下来是很长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长长的内容", Created_time: "2025-04-16 21:11:45", Modified_time: "2025-04-16 21:11:45"},
}

// 设置允许跨域的中间件
func enableCORS(next http.Handler) http.Handler {
  return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
    w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

    next.ServeHTTP(w, r)
  })
}

func getNoteHandler(w http.ResponseWriter, r *http.Request) {
	// 获取 id
	id := r.URL.Path[len("/note/"):]

	var note *Note
	// 使用id查询note
	for _, n := range notes {
		if n.ID == id {
			note = &n
			break
		}
	}

	resp := new(Resp)
	if note == nil {
		resp.Message = "not found note"
	} else {
		resp.Message = "success"
		resp.Note = note
	}

	dat, err := json.Marshal(resp)
	if err != nil {
		fmt.Println("Marshal Error", err)
		w.WriteHeader(500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(200)
	w.Write(dat)
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/note/", getNoteHandler)

  handler := enableCORS(mux)

	http.ListenAndServe(":8080", handler)
}
