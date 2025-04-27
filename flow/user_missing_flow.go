package flow

type UserMissingFlow string

const (
	MissingFlowDomain  UserMissingFlow = "domain"
	MissingFlowProject UserMissingFlow = "project"
)
