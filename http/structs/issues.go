// This file was generated from JSON Schema using quicktype, do not modify it directly.
// To parse and unparse this JSON data, add this code to your project and do:
//
//    issues, err := UnmarshalIssues(bytes)
//    bytes, err = issues.Marshal()

package structs

import (
	"encoding/csv"
	"encoding/json"
	"os"
	"strconv"
)

type Issues []Issue

func UnmarshalIssuesJSON(data []byte) (Issues, error) {
	var r Issues
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *Issues) MarshalJSON() ([]byte, error) {
	return json.Marshal(r)
}

func (r *Issues) MarshalCSV(file *os.File) error {
	writer := csv.NewWriter(file)
	var rows [][]string
	for _, iss := range *r {
		rows = append(rows, iss.GenericCSVRow())
	}
	_ = writer.WriteAll(rows)
	return nil
}

type Issue struct {
	ID                   int                  `json:"id"`
	Iid                  int                  `json:"iid"`
	ProjectID            int                  `json:"project_id"`
	Title                string               `json:"title"`
	Description          *string              `json:"description"`
	State                string               `json:"state"`
	CreatedAt            string               `json:"created_at"`
	UpdatedAt            string               `json:"updated_at"`
	ClosedAt             *string              `json:"closed_at"`
	ClosedBy             *Assignee            `json:"closed_by"`
	Labels               []string             `json:"labels"`
	Milestone            interface{}          `json:"milestone"`
	Assignees            []Assignee           `json:"assignees"`
	Author               Assignee             `json:"author"`
	Type                 string               `json:"type"`
	Assignee             Assignee             `json:"assignee"`
	UserNotesCount       int                  `json:"user_notes_count"`
	MergeRequestsCount   int                  `json:"merge_requests_count"`
	Upvotes              int                  `json:"upvotes"`
	Downvotes            int                  `json:"downvotes"`
	DueDate              *string              `json:"due_date"`
	Confidential         bool                 `json:"confidential"`
	DiscussionLocked     bool                 `json:"discussion_locked"`
	IssueType            string               `json:"issue_type"`
	WebURL               string               `json:"web_url"`
	TimeStats            TimeStats            `json:"time_stats"`
	TaskCompletionStatus TaskCompletionStatus `json:"task_completion_status"`
	HasTasks             bool                 `json:"has_tasks"`
	Links                Links                `json:"_links"`
	References           References           `json:"references"`
	MovedToID            int                  `json:"moved_to_id"`
	ServiceDeskReplyTo   interface{}          `json:"service_desk_reply_to"`
}

func (r *Issue) GenericCSVRow() []string {
	return []string{
		strconv.Itoa(r.ID),
		r.Author.Name,
	}
}

type Assignee struct {
	ID        int64  `json:"id"`
	Name      string `json:"name"`
	Username  string `json:"username"`
	State     string `json:"state"`
	AvatarURL string `json:"avatar_url"`
	WebURL    string `json:"web_url"`
}

type Links struct {
	Self       string `json:"self"`
	Notes      string `json:"notes"`
	AwardEmoji string `json:"award_emoji"`
	Project    string `json:"project"`
}

type References struct {
	Short    string `json:"short"`
	Relative string `json:"relative"`
	Full     string `json:"full"`
}

type TaskCompletionStatus struct {
	Count          int64 `json:"count"`
	CompletedCount int64 `json:"completed_count"`
}

type TimeStats struct {
	TimeEstimate        int64   `json:"time_estimate"`
	TotalTimeSpent      int64   `json:"total_time_spent"`
	HumanTimeEstimate   *string `json:"human_time_estimate"`
	HumanTotalTimeSpent *string `json:"human_total_time_spent"`
}
