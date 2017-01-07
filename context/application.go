package context

type ApplicationContext interface {
	DaoContext() DaoContext
	ServiceContext() ServiceContext
}

type applicationContext struct {
	daoContext     DaoContext
	serviceContext ServiceContext
}

func (c *applicationContext) DaoContext() DaoContext {
	return c.daoContext;
}

func (c *applicationContext) ServiceContext() ServiceContext {
	return c.serviceContext;
}

func NewApplicationContext() ApplicationContext {
	daoContext := NewMongoDaoContext("127.0.0.1", "khazix")
	return &applicationContext{
		daoContext: daoContext,
		serviceContext: NewServiceContext(daoContext),
	}
}