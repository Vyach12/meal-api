package business

type Meal struct {
	Id           int64
	MealName     string
	Category     string
	Area         string
	Instructions string
	MealThumb    string
	Tags         []string
	YoutubeUrl   string
	Ingredients  []string
	Measure      []string
	Source       string
	ImageSource  *string
}
