package query

const InsertDetailStageSql = `
	INSERT INTO detail_stage_versions(detail_id, stage_id, comment)
	VALUES ($1, $2, $3)
	RETURNING id;
`
