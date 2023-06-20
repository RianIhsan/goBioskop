package models


type Film struct {
  ID uint `json:"id" gorm:"primarykey"`
  Title string `json:"title"`
  Image string `json:"image"`
  Genre string `json:"genre"`
  Description string `json:"description"`
  Duration string `json:"duration"`
  Creator string `json:"creator"`
  ReleaseDate string `json:"releasedate"`
}

type FilmReq struct {
  Title string `json:"title" form:"title"`
  Image string `json:"image" form:"image"`
  Genre string `json:"genre" form:"genre"`
  Description string `json:"description" form:"description"`
  Duration string `json:"duration" form:"duration"`
  Creator string `json:"creator" form:"creator"`
  ReleaseDate string `json:"releasedate" form:"releasedate"`
}
