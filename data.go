package main

type DLData struct {
	ImageURL string
	ImageID  string
}

type FBUser struct {
	ID       string `json:"id"`
	About    string `json:"about"`
	Awards   string `json:"awards"`
	CanPost  bool   `json:"can_post"`
	Category string `json:"category"`
	Checkins int    `json:"checkins"`
	Cover    struct {
		CoverID string `json:"cover_id"`
		OffsetX int    `json:"offset_x"`
		OffsetY int    `json:"offset_y"`
		Source  string `json:"source"`
		ID      string `json:"id"`
	} `json:"cover"`
	HasAddedApp     bool   `json:"has_added_app"`
	IsCommunityPage bool   `json:"is_community_page"`
	IsPublished     bool   `json:"is_published"`
	Likes           int    `json:"likes"`
	Link            string `json:"link"`
	Location        struct {
		Zip string `json:"zip"`
	} `json:"location"`
	Name    string `json:"name"`
	Parking struct {
		Lot    int `json:"lot"`
		Street int `json:"street"`
		Valet  int `json:"valet"`
	} `json:"parking"`
	PersonalInfo      string `json:"personal_info"`
	PersonalInterests string `json:"personal_interests"`
	TalkingAboutCount int    `json:"talking_about_count"`
	Username          string `json:"username"`
	WereHereCount     int    `json:"were_here_count"`
}

type FBPhotos struct {
	Data []struct {
		ID          string `json:"id"`
		CreatedTime string `json:"created_time"`
		From        struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		Height int    `json:"height"`
		Icon   string `json:"icon"`
		Images []struct {
			Height int    `json:"height"`
			Source string `json:"source"`
			Width  int    `json:"width"`
		} `json:"images"`
		Link    string `json:"link"`
		Name    string `json:"name"`
		Picture string `json:"picture"`
		Place   struct {
			Name     string `json:"name"`
			Location struct {
				City      string  `json:"city"`
				Country   string  `json:"country"`
				Latitude  float64 `json:"latitude"`
				Longitude float64 `json:"longitude"`
				Street    string  `json:"street"`
				Zip       string  `json:"zip"`
			} `json:"location"`
			ID string `json:"id"`
		} `json:"place"`
		Source      string `json:"source"`
		UpdatedTime string `json:"updated_time"`
		Width       int    `json:"width"`
		Tags        struct {
			Data []struct {
				ID          string  `json:"id"`
				Name        string  `json:"name"`
				CreatedTime string  `json:"created_time"`
				X           float64 `json:"x"`
				Y           float64 `json:"y"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					Before string `json:"before"`
					After  string `json:"after"`
				} `json:"cursors"`
			} `json:"paging"`
		} `json:"tags"`
		Likes struct {
			Data []struct {
				ID   string `json:"id"`
				Name string `json:"name"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					Before string `json:"before"`
					After  string `json:"after"`
				} `json:"cursors"`
				Next string `json:"next"`
			} `json:"paging"`
		} `json:"likes"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			Before string `json:"before"`
			After  string `json:"after"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}

type FBAlbums struct {
	Data []struct {
		CanUpload bool `json:"can_upload"`
		Comments  struct {
			Data []struct {
				CanRemove   bool   `json:"can_remove"`
				CreatedTime string `json:"created_time"`
				From        struct {
					ID   string `json:"id"`
					Name string `json:"name"`
				} `json:"from"`
				ID        string `json:"id"`
				LikeCount int    `json:"like_count"`
				Message   string `json:"message"`
				UserLikes bool   `json:"user_likes"`
			} `json:"data"`
			Paging struct {
				Cursors struct {
					After  string `json:"after"`
					Before string `json:"before"`
				} `json:"cursors"`
			} `json:"paging"`
		} `json:"comments"`
		Count       int    `json:"count"`
		CoverPhoto  string `json:"cover_photo"`
		CreatedTime string `json:"created_time"`
		From        struct {
			ID   string `json:"id"`
			Name string `json:"name"`
		} `json:"from"`
		ID          string `json:"id"`
		Link        string `json:"link"`
		Name        string `json:"name"`
		Privacy     string `json:"privacy"`
		Type        string `json:"type"`
		UpdatedTime string `json:"updated_time"`
	} `json:"data"`
	Paging struct {
		Cursors struct {
			After  string `json:"after"`
			Before string `json:"before"`
		} `json:"cursors"`
		Next string `json:"next"`
	} `json:"paging"`
}
