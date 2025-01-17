// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AddressesColumns holds the columns for the "addresses" table.
	AddressesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true, SchemaType: map[string]string{"postgres": "SERIAL"}},
		{Name: "street", Type: field.TypeString},
		{Name: "district", Type: field.TypeString},
		{Name: "city", Type: field.TypeString},
		{Name: "country", Type: field.TypeString},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// AddressesTable holds the schema information for the "addresses" table.
	AddressesTable = &schema.Table{
		Name:       "addresses",
		Columns:    AddressesColumns,
		PrimaryKey: []*schema.Column{AddressesColumns[0]},
	}
	// UsersColumns holds the columns for the "users" table.
	UsersColumns = []*schema.Column{
		{Name: "id", Type: field.TypeUUID, Unique: true},
		{Name: "username", Type: field.TypeString, Unique: true},
		{Name: "email", Type: field.TypeString, Unique: true},
		{Name: "avatar", Type: field.TypeString, Nullable: true},
		{Name: "birth_day", Type: field.TypeString, Nullable: true},
		{Name: "first_name", Type: field.TypeString, Nullable: true},
		{Name: "last_name", Type: field.TypeString, Nullable: true},
		{Name: "phone_number", Type: field.TypeString, Nullable: true},
		{Name: "gender", Type: field.TypeInt, Nullable: true},
		{Name: "status", Type: field.TypeString},
		{Name: "address", Type: field.TypeString, Nullable: true},
		{Name: "street", Type: field.TypeString, Nullable: true},
		{Name: "district_id", Type: field.TypeString, Nullable: true},
		{Name: "city_id", Type: field.TypeString, Nullable: true},
		{Name: "country_code", Type: field.TypeString, Nullable: true},
		{Name: "hash_password", Type: field.TypeString},
		{Name: "verified", Type: field.TypeBool, Default: false},
		{Name: "created_at", Type: field.TypeTime},
		{Name: "updated_at", Type: field.TypeTime},
	}
	// UsersTable holds the schema information for the "users" table.
	UsersTable = &schema.Table{
		Name:       "users",
		Columns:    UsersColumns,
		PrimaryKey: []*schema.Column{UsersColumns[0]},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AddressesTable,
		UsersTable,
	}
)

func init() {
}
