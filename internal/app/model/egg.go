package model

type MealInfo struct {
	Id        int64   `db:"id"`
	UserId    int64   `db:"user_id"`
	MealDate  string  `db:"meal_date"`
	Weight    float32 `db:"weight"`
	Proteins  float32 `db:"proteins"`
	Fat       float32 `db:"fat"`
	Carbo     float32 `db:"carbo"`
	Nutrition float32 `db:"nutrition"`
}
