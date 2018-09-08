package pencere

type MenuList struct {
	Menus []*Menu
}

type Menu struct {
	Title     string `json:"title"`
	MenuItems []*MenuItem
}

type MenuItem struct {
	Title string `json:"title"`
}
