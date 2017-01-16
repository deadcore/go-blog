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

func NewApplicationContext(configuration Configuration) ApplicationContext {
	daoContext := NewDaoContext(configuration)
	return &applicationContext{
		daoContext: daoContext,
		serviceContext: NewServiceContext(daoContext),
	}
}