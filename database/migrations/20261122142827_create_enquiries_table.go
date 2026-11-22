package migrations

import (
	"github.com/goravel/framework/contracts/database/schema"

	"goravel/app/facades"
)

type M20261122142827CreateEnquiriesTable struct{}

// Signature The unique signature for the migration.
func (r *M20261122142827CreateEnquiriesTable) Signature() string {
	return "20261122142827_create_enquiries_table"
}

// Up Run the migrations.
func (r *M20261122142827CreateEnquiriesTable) Up() error {
	if !facades.Schema().HasTable("enquiries") {
		return facades.Schema().Create("enquiries", func(table schema.Blueprint) {
			table.ID()
			table.String("name")
			table.String("email")
			table.String("phone")
			table.String("service").Nullable()
			table.Text("message")
			table.TimestampsTz()
		})
	}

	return nil
}

// Down Reverse the migrations.
func (r *M20261122142827CreateEnquiriesTable) Down() error {
	return facades.Schema().DropIfExists("enquiries")
}
