package repository

const (
	addPostQuery = `
	INSERT INTO post (
		title,
		slug,
		content,
		image_url,
		category
	) VALUES (
		?,
		?,
		?,
		?,
		?
	)
`

	getAllPostQuery = `
	SELECT
		*
	FROM
		post
	ORDER BY
		id
`

	getPostQuery = `
	SELECT
		id,
		title,
		slug,
		content,
		image_url,
		category
	FROM
		post
	WHERE
		id=?
`

	getPostByTitle = `
		SELECT
			id,
			title,
			slug,
			content,
			image_url,
			category
		FROM
			post
		WHERE 
			title like ?
		ORDER BY 
			title
`
	getPostBySlug = `
		SELECT
			id,
			title,
			slug,
			content,
			image_url,
			category
		FROM
			post
		WHERE 
			slug like ?
		ORDER BY 
			slug
`

	updatePostQuery = `
	UPDATE
		post
	SET
		title = ?,
		slug = ?,
		content = ?,
		image_url = ?,
		category = ?
	WHERE
		id=?
`

	deletePostByID = `
	DELETE FROM 
		post 
	WHERE 
		id = ?
`
)
