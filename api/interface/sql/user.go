package sql

import (
	"context"
	"errors"

	"basic-service/domain"
	"basic-service/gen/db/model"
	"basic-service/gen/db/table"

	"braces.dev/errtrace"
	"github.com/go-jet/jet/v2/qrm"
	sqlite "github.com/go-jet/jet/v2/sqlite"
)

type UserRepository struct {
	db *SQLite
}

var (
	UserNotFoundErr = errors.New("user not found")
	UserExistsErr   = errors.New("user already exists")
)

func NewUserRepository(db *SQLite) *UserRepository {
	return &UserRepository{db: db}
}

type UserDataList struct {
	Total int
	Data  []domain.User
}

func (r *UserRepository) ListByRole(ctx context.Context, role domain.RoleType, page, pageSize int) (UserDataList, error) {
	// Validate pagination parameters
	if page < 1 {
		page = 1
	}
	if pageSize < 1 || pageSize > 100 {
		pageSize = 10 // Default page size
	}

	// First get the total count of users with this role
	countStmt := sqlite.SELECT(
		sqlite.COUNT(table.Users.ID).AS("total"),
	).FROM(
		table.Users,
	).WHERE(
		table.Users.Role.EQ(sqlite.Int32(int32(role))),
	)

	var total struct {
		Total int64
	}
	err := countStmt.QueryContext(ctx, r.db.db, &total)
	if err != nil {
		return UserDataList{}, errtrace.Wrap(err)
	}

	// Then get the paginated user data for this role
	offset := (page - 1) * pageSize
	stmt := table.Users.SELECT(
		table.Users.ID,
		table.Users.Email,
		table.Users.Name,
		table.Users.Role,
		table.Users.IsActive,
		table.Users.CreatedAt,
		table.Users.UpdatedAt,
		table.Users.Profile,
	).WHERE(
		table.Users.Role.EQ(sqlite.Int32(int32(role))),
	).ORDER_BY(
		table.Users.CreatedAt.DESC(),
	).LIMIT(
		int64(pageSize),
	).OFFSET(
		int64(offset),
	)

	var dbUsers []model.Users
	err = stmt.QueryContext(ctx, r.db.db, &dbUsers)
	if err != nil {
		return UserDataList{}, errtrace.Wrap(err)
	}

	// Convert model.Users to domain.User
	users := []domain.User{}
	for _, dbUser := range dbUsers {
		users = append(users, domain.User{
			ID:        dbUser.ID,
			Email:     dbUser.Email,
			Name:      dbUser.Name,
			Role:      domain.RoleType(dbUser.Role),
			IsActive:  dbUser.IsActive,
			Profile:   dbUser.Profile,
			CreatedAt: dbUser.CreatedAt,
			UpdatedAt: dbUser.UpdatedAt,
		})
	}

	return UserDataList{
		Total: int(total.Total),
		Data:  users,
	}, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, id string) (*domain.User, error) {
	stmt := sqlite.SELECT(
		table.Users.ID,
		table.Users.Email,
		table.Users.Password,
		table.Users.Name,
		table.Users.Role,
		table.Users.IsActive,
		table.Users.CreatedAt,
		table.Users.UpdatedAt,
		table.Users.Profile,
	).FROM(
		table.Users,
	).WHERE(
		table.Users.ID.EQ(sqlite.String(id)),
	)

	var dbUser model.Users
	err := stmt.QueryContext(ctx, r.db.db, &dbUser)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, errtrace.Wrap(UserNotFoundErr)
		}
		return nil, errtrace.Wrap(err)
	}

	return &domain.User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Profile:   dbUser.Profile,
		Password:  dbUser.Password,
		Name:      dbUser.Name,
		Role:      domain.RoleType(dbUser.Role),
		IsActive:  dbUser.IsActive,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
	}, nil
}

func (r *UserRepository) GetEmail(ctx context.Context, email string) (*domain.User, error) {
	stmt := sqlite.SELECT(
		table.Users.ID,
		table.Users.Email,
		table.Users.Password,
		table.Users.Name,
		table.Users.Role,
		table.Users.IsActive,
		table.Users.CreatedAt,
		table.Users.UpdatedAt,
		table.Users.Profile,
	).FROM(
		table.Users,
	).WHERE(
		table.Users.Email.EQ(sqlite.String(email)), // Using email as username
	)

	var dbUser model.Users
	err := stmt.QueryContext(ctx, r.db.db, &dbUser)
	if err != nil {
		if errors.Is(err, qrm.ErrNoRows) {
			return nil, errtrace.Wrap(UserNotFoundErr)
		}
		return nil, errtrace.Wrap(err)
	}

	return &domain.User{
		ID:        dbUser.ID,
		Email:     dbUser.Email,
		Password:  dbUser.Password,
		Name:      dbUser.Name,
		Role:      domain.RoleType(dbUser.Role),
		IsActive:  dbUser.IsActive,
		CreatedAt: dbUser.CreatedAt,
		UpdatedAt: dbUser.UpdatedAt,
		Profile:   dbUser.Profile,
	}, nil
}

func (r *UserRepository) UpdateUserState(ctx context.Context, id string, isActive bool) error {
	stmt := table.Users.UPDATE().
		SET(
			table.Users.IsActive.SET(sqlite.Bool(isActive)),
			table.Users.UpdatedAt.SET(sqlite.CURRENT_TIMESTAMP()),
		).WHERE(
		table.Users.ID.EQ(sqlite.String(id)),
	)

	_, err := stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}

func (r *UserRepository) Create(ctx context.Context, user *domain.User) error {
	// Check if user already exists
	existingUser, err := r.GetEmail(ctx, user.Email)
	if err != nil && !errors.Is(err, UserNotFoundErr) {
		return errtrace.Wrap(err)
	}
	if existingUser != nil {
		return errtrace.Wrap(UserExistsErr)
	}

	stmt := table.Users.INSERT(
		table.Users.ID,
		table.Users.Email,
		table.Users.Password,
		table.Users.Name,
		table.Users.Role,
		table.Users.IsActive,
		table.Users.Profile,
	).VALUES(
		sqlite.String(user.ID),
		sqlite.String(user.Email),
		sqlite.String(user.Password),
		sqlite.String(user.Name),
		sqlite.Int32(int32(user.Role)),
		sqlite.Bool(user.IsActive),
		sqlite.String(user.Profile),
	)

	_, err = stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}

func (r *UserRepository) Update(ctx context.Context, user *domain.User) error {
	stmt := table.Users.UPDATE().
		SET(
			table.Users.Email.SET(sqlite.String(user.Email)),
			table.Users.Name.SET(sqlite.String(user.Name)),
			table.Users.Profile.SET(sqlite.String(user.Profile)),
			table.Users.UpdatedAt.SET(sqlite.CURRENT_TIMESTAMP()),
		).WHERE(
		table.Users.ID.EQ(sqlite.String(user.ID)),
	)

	_, err := stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}

func (r *UserRepository) UpdatePassword(ctx context.Context, id, password string) error {
	stmt := table.Users.UPDATE().
		SET(
			table.Users.Password.SET(sqlite.String(password)),
			table.Users.UpdatedAt.SET(sqlite.CURRENT_TIMESTAMP()),
		).WHERE(
		table.Users.ID.EQ(sqlite.String(id)),
	)

	_, err := stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}

func (r *UserRepository) Delete(ctx context.Context, id string) error {
	stmt := table.Users.DELETE().
		WHERE(table.Users.ID.EQ(sqlite.String(id)))

	_, err := stmt.ExecContext(ctx, r.db.db)
	return errtrace.Wrap(err)
}
