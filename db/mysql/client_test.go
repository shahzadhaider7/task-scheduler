package mysql

import (
	"fmt"
	"github.com/shahzadhaider7/task-scheduler/db"
	"github.com/shahzadhaider7/task-scheduler/models"
	"os"
	"reflect"
	"testing"
	"time"
)

func Test_client_AddTask(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "password")

	type args struct {
		task *models.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// test cases
		{
			name:    "success - add task in db",
			args:    args{task: &models.Task{Name: "meeting", CreatedAt: time.Now(), Status: "active", Data: "Hello"}},
			wantErr: false,
		},
		{
			name:    "fail - add invalid task in db",
			args:    args{task: &models.Task{ID: "4", Name: "coding", CreatedAt: time.Now(), Status: "active", Data: "Hello"}},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c, _ := NewClient(db.Option{})
			_, err := c.AddTask(tt.args.task)
			if (err != nil) != tt.wantErr {
				t.Errorf("AddTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
		})
	}
}

func Test_client_DeleteTask(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "password")

	c, _ := NewClient(db.Option{})
	task := &models.Task{Name: "Sleeping", CreatedAt: time.Now(), Status: "active", Data: "Hi"}
	_, _ = c.AddTask(task)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// test cases
		{
			name:    "success - delete task from db",
			args:    args{id: task.ID},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.DeleteTask(tt.args.id); (err != nil) != tt.wantErr {
				t.Errorf("DeleteTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_client_GetTask(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "password")

	c, _ := NewClient(db.Option{})
	task := &models.Task{Name: "party", CreatedAt: time.Now().UTC().Truncate(time.Minute), Status: "active", Data: "Bingo"}
	_, _ = c.AddTask(task)
	type args struct {
		id string
	}
	tests := []struct {
		name    string
		args    args
		want    *models.Task
		wantErr bool
	}{
		// test cases
		{
			name:    "success - get task from db",
			args:    args{id: task.ID},
			want:    task,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := c.GetTask(tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetTask() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetTask() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_client_UpdateSession(t *testing.T) {
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_USER", "root")
	os.Setenv("DB_PASS", "password")

	c, err := NewClient(db.Option{})
	if err != nil {
		fmt.Print(err, c)
	}
	type args struct {
		task *models.Task
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name:    "success - update task in db",
			args:    args{task: &models.Task{ID: "011f1949-1112-4214-8bac-1aa12594963d", Name: "Volla", CreatedAt: time.Now().UTC().Truncate(time.Minute), Status: "active", Data: "Hala Madrid"}},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := c.UpdateTask(tt.args.task); (err != nil) != tt.wantErr {
				t.Errorf("UpdateTask() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
