---
id: schema-fields
title: Fields
---

## Quick Summary

Fields (or properties) in the schema are the attributes of the vertex. For example, a `User`
with 4 fields: `age`, `name`, `username` and `created_at`:

![re-fields-properties](https://entgo.io/assets/er_fields_properties.png)

Fields are returned from the schema using the `Fields` method. For example:

```go
package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age"),
		field.String("name"),
		field.String("username").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}
```

All fields are required by default, and can be set to optional using the `Optional` method.

## Types

The following types are currently supported by the framework:

- All Go numeric types. Like, `int`, `uint8`, `float64`, etc.
- `bool`
- `string`
- `time.Time`
- `[]byte` (only support by SQL dialects).

```go
package schema

import (
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)

// User schema.
type User struct {
	ent.Schema
}

// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Int("age").
			Positive(),
		field.Float("rank").
			Optional(),
		field.Bool("active").
			Default(false),
		field.String("name").
			Unique(),
		field.Time("created_at").
			Default(time.Now),
	}
}
```

To read more about how each type is mapped to its database-type, go to the [Migration](migrate.md) section.

## Default Values

**Non-unique** fields support default values using the `.Default` and `.UpdateDefault` methods.

```go
// Fields of the User.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.Time("created_at").
			Default(time.Now),
		field.Time("updated_at").
			Default(time.Now).
			UpdateDefault(time.Now),
	}
}
```

## Validators

A field validator is a function from type `func(T) error` that is defined in the schema
using the `Validate` method, and applied on the given field value before creating or updating
the entity.

The supported types for field validators are `string` and all numeric types.

```go
package schema

import (
	"errors"
	"regexp"
	"strings"
	"time"

	"github.com/facebookincubator/ent"
	"github.com/facebookincubator/ent/schema/field"
)


// Group schema.
type Group struct {
	ent.Schema
}

// Fields of the group.
func (Group) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			Match(regexp.MustCompile("[a-zA-Z_]+$")).
			Validate(func(s string) error {
				if strings.ToLower(s) == s {
					return errors.New("group name must begin with uppercase")
				}
				return nil
			}),
	}
}
```

## Builtin Validators

The framework provides a few builtin validators for each type:

- Numeric types:
  - `Positive()`
  - `Negative()`
  - `Min(i)` - Validate that the given value is > i.
  - `Max(i)` - Validate that the given value is < i.
  - `Range(i, j)` - Validate that the given value is within the range [i, j].

- `string`
  - `MinLen(i)`
  - `MaxLen(i)`
  - `Match(regexp.Regexp)`

## Optional

Optional fields are fields that are not required in the entity creation, and
will be set to nullable fields in the database.
Unlike edges, **fields are required by default**, and setting them to
optional should be done explicitly using the `Optional` method.


```go
// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("required_name"),
		field.String("optional_name").
			Optional(),
	}
}
```

## Nillable
Sometime, you want to be able to distinguish between the zero value of fields
and `nil`. For example, if the database column contains `0` or `NULL`. The `Nillable` option exists exactly for this.

If you have an `Optional` field of type `T`, setting it to `Nillable` will generate
a struct field with type `*T`. Hence, if the database returns `NULL` for this field,
the struct field will be `nil`. Otherwise, it will contains a pointer to the actual data.

For example, given this schema:
```go
// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("required_name"),
		field.String("optional_name").
			Optional(),
		field.String("nillable_name").
			Optional().
			Nillable(),
	}
}
```

The generated struct for the `User` entity will be as follows:

```go
// ent/user.go
package ent

// User entity.
type User struct {
	RequiredName string `json:"required_name,omitempty"`
	OptionalName string `json:"optional_name,omitempty"`
	NillableName string `json:"nillable_name,omitempty"`
}
```

## Immutable

Immutable fields are fields that can be set only in the creation of the entity.
i.e., no setters will be generated for the entity updater.

```go
// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.Time("created_at").
			Default(time.Now),
			Immutable(),
	}
}
```

## Uniqueness
Fields can be defined as unique using the `Unique` method.
Note that, unique fields cannot have default values.

```go
// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name"),
		field.String("nickname").
			Unique(),
	}
}
```

## Indexes
Indexes can be defined on multi fields and some types of edges as well.
However, you should note, that this is currently an SQL-only feature.

Read more about this in the [Indexes](schema-indexes.md) section.

## Struct Tags

Custom struct tags can be added to the generated entities using the `StructTag`
method. Note that, if this option was not provided, or provided and did not
contain the `json` tag, the default `json` tag will be added with the field name.

```go
// Fields of the user.
func (User) Fields() []ent.Field {
	return []ent.Field{
		field.String("name").
			StructTag(`gqlgen:"gql_name"`),
	}
}
```