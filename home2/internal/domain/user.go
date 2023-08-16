package domain

// User 领域对象， 是 DDD 中的entirely
// BO(business object)
type User struct {
	Id           int64
	Email        string
	Password     string
	NickName     string
	Birthday     string
	Introduction string
	CreateTime   string
	UpdateTime   string
}
