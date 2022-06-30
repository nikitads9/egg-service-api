package repository

import (
	"context"
	"fmt"

	sq "github.com/Masterminds/squirrel"
	"github.com/jmoiron/sqlx"
	"github.com/nikitads9/egg-service-api/internal/app/model"
	"github.com/nikitads9/egg-service-api/internal/app/repository/table"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type IEggNutritionRepository interface {
	AddMeal(ctx context.Context, meal *model.MealInfo) (int64, error)
	GetList(ctx context.Context, userId int64) ([]*model.MealInfo, error)
	GetMeal(ctx context.Context, id int64, userId int64) (*model.MealInfo, error)
	RemoveMeal(ctx context.Context, id int64, userId int64) (int64, error)
	UpdateMeal(ctx context.Context, meal *model.MealInfo) error
}
type mealRepository struct {
	db *sqlx.DB
}

func NewEggNutritionRepository(db *sqlx.DB) IEggNutritionRepository {
	return &mealRepository{
		db: db,
	}
}

func (n *mealRepository) AddMeal(ctx context.Context, meal *model.MealInfo) (int64, error) {
	builder := sq.Insert(table.MealsTable).
		PlaceholderFormat(sq.Dollar).
		Columns("title, content").
		Values(meal.Title, meal.Content).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := n.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return 0, err
	}

	return id, nil
}

func (n *mealRepository) GetList(ctx context.Context, userId int64) ([]*model.MealInfo, error) {
	builder := sq.Select("title, content").
		PlaceholderFormat(sq.Dollar).
		From(table.MealsTable)

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []*model.MealInfo

	err = n.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	return res, nil
}

func (n *mealRepository) GetMeal(ctx context.Context, id int64, userId int64) (*model.MealInfo, error) {
	builder := sq.Select("id, title, content").
		PlaceholderFormat(sq.Dollar).
		From(table.MealsTable).
		Where(sq.Eq{"id": id})

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, err
	}

	var res []*model.MealInfo

	err = n.db.SelectContext(ctx, &res, query, args...)
	if err != nil {
		return nil, err
	}

	if len(res) == 0 {
		return nil, status.Error(codes.NotFound, "ebanyi rot etogo kasino")
	}

	return res[0], nil
}

func (n *mealRepository) RemoveMeal(ctx context.Context, id int64, userId int64) (int64, error) {
	builder := sq.Delete(table.MealsTable).
		PlaceholderFormat(sq.Dollar).
		Where(sq.Eq{"id": id}).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return 0, err
	}

	row, err := n.db.QueryContext(ctx, query, args...)
	if err != nil {
		return 0, err
	}
	defer row.Close()

	row.Next()
	var removed int64
	err = row.Scan(&removed)
	if err != nil {
		return 0, err
	}

	return removed, nil
}

func (n *mealRepository) UpdateMeal(ctx context.Context, meal *model.MealInfo) error {
	builder := sq.Update(table.MealsTable).
		Set("title", meal.Title).
		Set("content", meal.Content).
		Where(sq.Eq{"id": meal.Id}).
		PlaceholderFormat(sq.Dollar).
		Suffix("returning id")

	query, args, err := builder.ToSql()
	if err != nil {
		return err
	}

	row, err := n.db.QueryContext(ctx, query, args...)
	if err != nil {
		return err
	}
	defer row.Close()

	var id int64
	row.Next()
	err = row.Scan(&id)
	if err != nil {
		return err
	}

	fmt.Printf("edited meal with id %v\n", id)

	return nil
}
