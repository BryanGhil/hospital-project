package entity

type DefaultPageFilter struct {
	Page   int
	Limit  int
	Offset int
}

type GetPageResponse struct {
	Page      int
	Limit     int
	CountData int
	Data      any
}
