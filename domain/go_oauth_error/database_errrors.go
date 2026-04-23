package gooautherror

type NotDatabaseConfigurateError struct{}

func (NotDatabaseConfigurateError) Error() string {
	return "NOT_DATABASE_CONFIGURATE_ERROR"
}
