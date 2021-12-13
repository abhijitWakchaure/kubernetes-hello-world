package main

// Meta ...
type Meta struct {
	Host        string `json:"host"`
	ServiceType string `json:"serviceType"`
}

// Task ...
type Task struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

// RequestInsertTask ...
type RequestInsertTask struct {
	Name string `json:"name,required" binding:"min=3,required"`
}

// ResponseInsertTask ...
type ResponseInsertTask struct {
	Meta Meta `json:"meta,omitempty"`
	Task
}

// ResponseListTasks ...
type ResponseListTasks struct {
	Meta  Meta   `json:"meta,omitempty"`
	Tasks []Task `json:"tasks"`
}

// ResponseError ...
type ResponseError struct {
	Code  int    `json:"code"`
	Error string `json:"error"`
}
