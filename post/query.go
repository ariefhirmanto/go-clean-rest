package post

const (
	addTicketQuery = `
	INSERT INTO ticket (
		name,
		description,
		price,
		rating,
		image_url
	) VALUES (
		$1,
		$2,
		$3,
		$4,
		$5
	) returning id
`
	getTicketQuery = `
	SELECT
		name,
		description,
		price,
		rating,
		image_url
	FROM
		ticket
	WHERE
		id=$1
`

	getTicketBatchQuery = `
	SELECT
		*
	FROM
		ticket
	LIMIT $1
	OFFSET $2
`

	updateProductQuery = `
	UPDATE
		ticket
	SET
		%s
	WHERE
		id=%d
`
)
