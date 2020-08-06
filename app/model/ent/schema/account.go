package schema

/*
  Account
  ID该字段内置于架构中，不需要声明。
  在基于 SQL 的数据库中，其类型默认为数据库中自动递增
  https://entgo.io/docs/getting-started/
*/
import (
	"time"

	"github.com/facebookincubator/ent"
	//"github.com/facebookincubator/ent/schema/edge"
	"github.com/facebookincubator/ent/schema/field"
)

// Account holds the schema definition for the Account entity.
type Account struct {
	ent.Schema
}

// Config of the Account.
func (Account) Config() ent.Config {
	return ent.Config{
		Table: "account",
	}
}

// Hooks of the Card.
func (Account) Hooks() []ent.Hook {
	return nil
}

// Fields of the Account.
func (Account) Fields() []ent.Field {
	return []ent.Field{
		field.String("account"),                    // 账户
		field.String("account_type"),               // 账户类型
		field.String("platform"),                   // 账户归属平台
		field.String("verify_type"),                // 校验方式
		field.String("password"),                   // 登录密码
		field.String("password_salt"),              // 密码盐值
		field.String("password_type"),              // 校验方式
		field.Int("user_id"),                       // 用户标识
		field.Int("role_id"),                       // 角色标识
		field.Int("status"),                        // 状态
		field.String("desc"),                       // 账户描述
		field.Int("oauth2_id"),                     // oauth2认证
		field.String("oauth2_token"),               // oauth2令牌
		field.Time("oauth2_time"),                  // oauth2创建时间
		field.String("token_fake"),                 // oauth2令牌
		field.String("creator"),                    // 创建人
		field.Time("created_at").Default(time.Now), // 创建时间
		field.Time("updated_at").Default(time.Now), // 更新时间
		field.Int("version").Default(1),            // 数据版本
		field.String("string_1"),                   // 备用字段
		field.String("string_2"),                   // 备用字段
		field.String("string_3"),                   // 备用字段
		field.Int("number_1"),                      // 备用字段
		field.Int("number_2"),                      // 备用字段
		field.Int("number_3"),                      // 备用字段
	}
}

// Edges of the Account.
func (Account) Edges() []ent.Edge {
	//	return []ent.Edge{
	//	}
	return nil
}
