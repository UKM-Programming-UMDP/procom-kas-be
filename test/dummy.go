package test

func ClearAllDummy(config TestConfig) {
	DeleteKasSubmissionDummy(config)
	DeleteUserDummy(config)
	DeleteMonthDummy(config)
}

func CreateMonthDummy(config TestConfig) {
	config.Db.Exec("INSERT INTO month_schemas (id, year, month, show) VALUES (-1, 2000, 1, false)")
}
func DeleteMonthDummy(config TestConfig) {
	config.Db.Exec("DELETE FROM month_schemas WHERE year = 2000 AND month = 1")
}

func CreateUserDummy(config TestConfig) {
	CreateMonthDummy(config)
	config.Db.Exec("INSERT INTO user_schemas (id, npm, name, email, kas_payed, month_id) VALUES (-1, '1928476912', 'user_test', 'user_test@mail.com', 0, -1)")
}
func DeleteUserDummy(config TestConfig) {
	config.Db.Exec("DELETE FROM user_schemas WHERE npm = '1928476912'")
	DeleteMonthDummy(config)
}

func CreateKasSubmissionDummy(config TestConfig) {
	CreateUserDummy(config)
	config.Db.Exec("INSERT INTO kas_submission_schemas (id, user_npm, submission_id, payed_amount, status, note, evidence) VALUES (-1, '1928476912', 'test1', 100000, 'Pending', 'ini test', 'evidence.png')")
}
func DeleteKasSubmissionDummy(config TestConfig) {
	config.Db.Exec("DELETE FROM kas_submission_schemas WHERE user_npm = '1928476912'")
	DeleteUserDummy(config)
}

func CreateFinancialRequestDummy(config TestConfig) {
	CreateUserDummy(config)
	config.Db.Exec("INSERT INTO financial_request_schemas (id, request_id, amount, note, user_npm, status, type, target_provider, target_number evidence, transfered_evidence) VALUES (-1, 'test1', 100000, 'ini test', '1928476912', 'Pending', 'Transfer', 'Bank - TEST', '1234567890', 'evidence.png', 'transfered_evidence.png')")
}
func DeleteFinancialRequestDummy(config TestConfig) {
	config.Db.Exec("DELETE FROM financial_request_schemas WHERE user_npm = '1928476912'")
	DeleteUserDummy(config)
}
