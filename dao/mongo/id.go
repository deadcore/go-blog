package mongo

type EntityWithId interface {
	SetId(id string)
}