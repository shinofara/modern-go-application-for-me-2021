// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"strings"

	"entgo.io/ent/dialect/sql"
	"github.com/shinofara/modern-go-application-for-me-2021/ent/auth"
	"github.com/shinofara/modern-go-application-for-me-2021/ent/user"
)

// User is the model entity for the User schema.
type User struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the UserQuery when eager-loading is set.
	Edges     UserEdges `json:"edges"`
	auth_user *int
}

// UserEdges holds the relations/edges for other nodes in the graph.
type UserEdges struct {
	// Auth holds the value of the auth edge.
	Auth *Auth `json:"auth,omitempty"`
	// CreateTasks holds the value of the create_tasks edge.
	CreateTasks []*Task `json:"create_tasks,omitempty"`
	// AssignTasks holds the value of the assign_tasks edge.
	AssignTasks []*Task `json:"assign_tasks,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [3]bool
}

// AuthOrErr returns the Auth value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e UserEdges) AuthOrErr() (*Auth, error) {
	if e.loadedTypes[0] {
		if e.Auth == nil {
			// The edge auth was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: auth.Label}
		}
		return e.Auth, nil
	}
	return nil, &NotLoadedError{edge: "auth"}
}

// CreateTasksOrErr returns the CreateTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) CreateTasksOrErr() ([]*Task, error) {
	if e.loadedTypes[1] {
		return e.CreateTasks, nil
	}
	return nil, &NotLoadedError{edge: "create_tasks"}
}

// AssignTasksOrErr returns the AssignTasks value or an error if the edge
// was not loaded in eager-loading.
func (e UserEdges) AssignTasksOrErr() ([]*Task, error) {
	if e.loadedTypes[2] {
		return e.AssignTasks, nil
	}
	return nil, &NotLoadedError{edge: "assign_tasks"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*User) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			values[i] = new(sql.NullInt64)
		case user.FieldName:
			values[i] = new(sql.NullString)
		case user.ForeignKeys[0]: // auth_user
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type User", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the User fields.
func (u *User) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case user.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			u.ID = int(value.Int64)
		case user.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				u.Name = value.String
			}
		case user.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field auth_user", value)
			} else if value.Valid {
				u.auth_user = new(int)
				*u.auth_user = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryAuth queries the "auth" edge of the User entity.
func (u *User) QueryAuth() *AuthQuery {
	return (&UserClient{config: u.config}).QueryAuth(u)
}

// QueryCreateTasks queries the "create_tasks" edge of the User entity.
func (u *User) QueryCreateTasks() *TaskQuery {
	return (&UserClient{config: u.config}).QueryCreateTasks(u)
}

// QueryAssignTasks queries the "assign_tasks" edge of the User entity.
func (u *User) QueryAssignTasks() *TaskQuery {
	return (&UserClient{config: u.config}).QueryAssignTasks(u)
}

// Update returns a builder for updating this User.
// Note that you need to call User.Unwrap() before calling this method if this User
// was returned from a transaction, and the transaction was committed or rolled back.
func (u *User) Update() *UserUpdateOne {
	return (&UserClient{config: u.config}).UpdateOne(u)
}

// Unwrap unwraps the User entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (u *User) Unwrap() *User {
	tx, ok := u.config.driver.(*txDriver)
	if !ok {
		panic("ent: User is not a transactional entity")
	}
	u.config.driver = tx.drv
	return u
}

// String implements the fmt.Stringer.
func (u *User) String() string {
	var builder strings.Builder
	builder.WriteString("User(")
	builder.WriteString(fmt.Sprintf("id=%v", u.ID))
	builder.WriteString(", name=")
	builder.WriteString(u.Name)
	builder.WriteByte(')')
	return builder.String()
}

// Users is a parsable slice of User.
type Users []*User

func (u Users) config(cfg config) {
	for _i := range u {
		u[_i].config = cfg
	}
}
