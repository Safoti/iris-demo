package datamodels

/**
 * @Author safoti
 * @Date Created in 2022/7/13
 * @Description
 **/
type Movie struct {
	ID     int64  `json:"id"`
	Name   string `json:"name"`
	Year   int    `json:"year"`
	Genre  string `json:"genre"`
	Poster string `json:"poster"`
}