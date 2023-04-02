package query

const DeleteDetailSql = `
	UPDATE details SET is_deleted = true where id = $1
`
