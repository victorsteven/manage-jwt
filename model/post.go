package model

type Post struct {
	ID          uint64 `json:"id"`
	UserID      uint64 `json:"user_id"`
	Title       string `json:"title"`
	Description string `json:"description"`
}

type PostDB struct {
	DB map[uint64]*Post
}

//seed one post, just to make sure we always have a post to test with:
func (s *PostDB) SeedPost() {
	post := &Post{}
	post.ID = 1
	post.UserID = 1
	post.Title = "Nice Post"
	post.Description = "Example Post"
	s.DB = make(map[uint64]*Post)
	s.DB[post.ID] = post
}

func (s *PostDB) Create(post *Post) (*Post, error) {
	//Seed the database first, so that the post created now will have an id of 2 instead of 1
	s.SeedPost()

	post.ID = uint64(len(s.DB) + 1)
	//if no record have been inserted the map yet
	if s.DB == nil {
		s.DB = make(map[uint64]*Post)
		s.DB[post.ID] = post
	} else {
		s.DB[post.ID] = post
	}
	return post, nil
}
