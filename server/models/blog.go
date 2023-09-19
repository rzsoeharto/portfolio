package models

import "time"

type BlogPost struct {
	ID        int
	Title     string
	Author    string
	Published time.Time
	Sections  []PostSection
}

type PostSection struct {
	ID          int
	BlogPostID  int
	SectionType string
	Content     string
}
