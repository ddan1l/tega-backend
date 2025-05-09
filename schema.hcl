schema "public" {}

enum "action_type" {
  schema = schema.public
  values = ["create", "read", "update", "delete"]
}

enum "resource_type" {
  schema = schema.public
  values = ["project", "task", "user", "user2"]
}

enum "condition_operator" {
  schema = schema.public
  values = ["eq", "neq", "contains", "startsWith"]
}

table "policy_actions" {
  schema = schema.public
  column "action" {
    type = enum.action_type 
  }
}