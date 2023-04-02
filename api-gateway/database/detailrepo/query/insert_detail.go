package query

const InsertDetailSql = `
	INSERT INTO details(long, width, height, color)
	VALUES ($1, $2, $3, $4)
	RETURNING id;
`
