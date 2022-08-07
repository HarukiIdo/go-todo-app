package handler

import (
	"net/http"

	"github.com/HarukiIdo/go-todo-app/entity"
	"github.com/HarukiIdo/go-todo-app/store"
)

type ListTask struct {
	Store *store.TaskStore
}

type task struct {
	ID     entity.TaskID     `json:"id"`
	Title  string            `json:"title"`
	Status entity.TaskStatus `json:"status"`
}

// タスクの一覧を返すHTTPハンドラー
func (lt *ListTask) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	tasks := lt.Store.All()
	rsp := []task{}
	for _, t := range tasks {
		rsp = append(rsp, task{
			ID:     t.ID,
			Title:  t.Title,
			Status: t.Status,
		})
	}
	// fmt.Println("ok")
	RespondJSON(ctx, w, rsp, http.StatusOK)
}
