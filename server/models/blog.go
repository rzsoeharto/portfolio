package models

import "time"

type BlogPost struct {
	ID        int
	Title     string
	Published time.Time
	Updated   time.Time
	Sections  []PostSection
}

type PostSection struct {
	ID          int
	BlogPostID  int
	SectionType string
	Content     string
}
