package relations


type R_Tag_Blog struct {
	Blog_id string `json:"blog_id" bson:"blog_id"`
	Tag_id  string `json:"tag_id" bson:"tag_id"`
}
