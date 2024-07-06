package repository

import (
	"context"
	"fmt"
	"net/http"

	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/constants"
	ctxlogger "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/ctxLogger"
	customerrors "github.com/Suhaan-Bhandary/go-api-template/internal/pkg/customErrors"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/dto"
	"github.com/Suhaan-Bhandary/go-api-template/internal/pkg/helpers"
	"github.com/Suhaan-Bhandary/go-api-template/internal/repository"
	"github.com/jmoiron/sqlx"

	sq "github.com/Masterminds/squirrel"
	"github.com/go-sql-driver/mysql"
)

type userStore struct {
	BaseRepository
}

func NewUserRepo(db *sqlx.DB) repository.UserStorer {
	return &userStore{
		BaseRepository: BaseRepository{DB: db},
	}
}

func (userStore *userStore) ListUsersPaginated(ctx context.Context, tx repository.Transaction, filters dto.ListUsersPaginatedRequest) (dto.PaginatedUsers, error) {
	ctxlogger.Info(ctx, "Starting list users paginated in repository")
	defer ctxlogger.Info(ctx, "Ending list users paginated in repository")

	queryExecutor := userStore.initiateQueryExecutor(tx)

	// ----- Get the count
	totalRecordsQueryBuilder := sq.Select("count(*) as count").From("users")

	// Filtering
	if filters.SearchValue != "" {
		totalRecordsQueryBuilder = totalRecordsQueryBuilder.Where(sq.Like{"email": "%" + filters.SearchValue + "%"})
	}

	countQuery, countArgs, err := totalRecordsQueryBuilder.ToSql()
	if err != nil {
		return dto.PaginatedUsers{}, fmt.Errorf("failed to build SQL query: %w", err)
	}

	totalRecords := 0
	err = queryExecutor.QueryRowx(countQuery, countArgs...).Scan(&totalRecords)
	if err != nil {
		return dto.PaginatedUsers{}, fmt.Errorf("failed to execute query: %w", err)
	}

	// ----- Get the paginated data
	queryBuilder := sq.Select("id", "email", "created_at", "updated_at").From("users")

	// Pagination
	queryBuilder = helpers.SetQueryBuilderOffsetAndLimit(queryBuilder, filters.Page, filters.Limit)

	// Filtering
	if filters.SearchValue != "" {
		queryBuilder = queryBuilder.Where(sq.Like{"email": "%" + filters.SearchValue + "%"})
	}

	query, args, err := queryBuilder.ToSql()
	if err != nil {
		return dto.PaginatedUsers{}, fmt.Errorf("failed to build SQL query: %w", err)
	}

	users := make([]dto.User, 0)
	err = sqlx.Select(queryExecutor, &users, query, args...)
	if err != nil {
		return dto.PaginatedUsers{}, fmt.Errorf("failed to execute query: %w", err)
	}

	return dto.PaginatedUsers{
		Users:      users,
		Pagination: helpers.GetPaginationMetaData(filters.Page, filters.Limit, totalRecords),
	}, nil
}

func (userStore *userStore) CreateUser(ctx context.Context, tx repository.Transaction, user repository.User) error {
	ctxlogger.Info(ctx, "Starting create user in repository")
	defer ctxlogger.Info(ctx, "Ending create user in repository")

	queryExecutor := userStore.initiateQueryExecutor(tx)

	query := `
		INSERT INTO users (id, email, password, created_at, updated_at) 
		VALUES (:id, :email, :password, :created_at, :updated_at)
	`

	query, values, err := queryExecutor.BindNamed(query, user)
	if err != nil {
		return fmt.Errorf("failed to insert user: %w", err)
	}

	_, err = queryExecutor.Exec(query, values...)
	if err != nil {
		sqlErr, ok := err.(*mysql.MySQLError)
		if ok && sqlErr.Number == constants.MYSQL_KEY_EXITS {
			return customerrors.Error{
				Code:          http.StatusBadRequest,
				CustomMessage: "User with the email already exits",
				InternalError: fmt.Errorf("create user email error, err: %s", err),
			}
		}

		return fmt.Errorf("failed to insert user: %w", err)
	}

	return nil
}
