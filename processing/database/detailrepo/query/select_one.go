package query

const SelectOneDetailSql = `
	SELECT * FROM details WHERE is_deleted = FALSE AND id = $1
`
