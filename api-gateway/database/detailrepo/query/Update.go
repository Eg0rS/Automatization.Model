package query

const UpdateDetailSql = `
	UPDATE details SET
	                   long = $1,
	                   width = $2,
	                   height = $3,
	                   color = $4,
	                   is_deleted = $5
	where id = $6
`
