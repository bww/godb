package godb

type Option func(d *Database) *Database

func OptionMaxOpenConns(v int) Option {
	return func(d *Database) *Database {
		d.db.SetMaxOpenConns(v)
		return d
	}
}

func OptionMaxIdleConns(v int) Option {
	return func(d *Database) *Database {
		d.db.SetMaxIdleConns(v)
		return d
	}
}
