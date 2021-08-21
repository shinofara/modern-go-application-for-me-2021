// Code generated by entc, DO NOT EDIT.

package ent

import (
	"fmt"
	"mygo/ent/auth"
	"mygo/ent/user"
	"strings"

	"entgo.io/ent/dialect/sql"
)

// Auth is the model entity for the Auth schema.
type Auth struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Email holds the value of the "email" field.
	Email string `json:"email,omitempty"`
	// Password holds the value of the "password" field.
	Password string `json:"password,omitempty"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the AuthQuery when eager-loading is set.
	Edges     AuthEdges `json:"edges"`
	auth_user *int
}

// AuthEdges holds the relations/edges for other nodes in the graph.
type AuthEdges struct {
	// User holds the value of the user edge.
	User *User `json:"user,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [1]bool
}

// UserOrErr returns the User value or an error if the edge
// was not loaded in eager-loading, or loaded but was not found.
func (e AuthEdges) UserOrErr() (*User, error) {
	if e.loadedTypes[0] {
		if e.User == nil {
			// The edge user was loaded in eager-loading,
			// but was not found.
			return nil, &NotFoundError{label: user.Label}
		}
		return e.User, nil
	}
	return nil, &NotLoadedError{edge: "user"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Auth) scanValues(columns []string) ([]interface{}, error) {
	values := make([]interface{}, len(columns))
	for i := range columns {
		switch columns[i] {
		case auth.FieldID:
			values[i] = new(sql.NullInt64)
		case auth.FieldEmail, auth.FieldPassword:
			values[i] = new(sql.NullString)
		case auth.ForeignKeys[0]: // auth_user
			values[i] = new(sql.NullInt64)
		default:
			return nil, fmt.Errorf("unexpected column %q for type Auth", columns[i])
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Auth fields.
func (a *Auth) assignValues(columns []string, values []interface{}) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case auth.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			a.ID = int(value.Int64)
		case auth.FieldEmail:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field email", values[i])
			} else if value.Valid {
				a.Email = value.String
			}
		case auth.FieldPassword:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field password", values[i])
			} else if value.Valid {
				a.Password = value.String
			}
		case auth.ForeignKeys[0]:
			if value, ok := values[i].(*sql.NullInt64); !ok {
				return fmt.Errorf("unexpected type %T for edge-field auth_user", value)
			} else if value.Valid {
				a.auth_user = new(int)
				*a.auth_user = int(value.Int64)
			}
		}
	}
	return nil
}

// QueryUser queries the "user" edge of the Auth entity.
func (a *Auth) QueryUser() *UserQuery {
	return (&AuthClient{config: a.config}).QueryUser(a)
}

// Update returns a builder for updating this Auth.
// Note that you need to call Auth.Unwrap() before calling this method if this Auth
// was returned from a transaction, and the transaction was committed or rolled back.
func (a *Auth) Update() *AuthUpdateOne {
	return (&AuthClient{config: a.config}).UpdateOne(a)
}

// Unwrap unwraps the Auth entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (a *Auth) Unwrap() *Auth {
	tx, ok := a.config.driver.(*txDriver)
	if !ok {
		panic("ent: Auth is not a transactional entity")
	}
	a.config.driver = tx.drv
	return a
}

// String implements the fmt.Stringer.
func (a *Auth) String() string {
	var builder strings.Builder
	builder.WriteString("Auth(")
	builder.WriteString(fmt.Sprintf("id=%v", a.ID))
	builder.WriteString(", email=")
	builder.WriteString(a.Email)
	builder.WriteString(", password=")
	builder.WriteString(a.Password)
	builder.WriteByte(')')
	return builder.String()
}

// Auths is a parsable slice of Auth.
type Auths []*Auth

func (a Auths) config(cfg config) {
	for _i := range a {
		a[_i].config = cfg
	}
}
