package user

const (
	addUserQuery = `
	INSERT INTO user (
		name,
		email,
		password,
		description
	) VALUES (
		$1,
		$2,
		$3,
		$4
	) returning id
`
	getUserQuery = `
	SELECT
		name,
		email,
		password,
		description
	FROM
		user
	WHERE
		id=$1
`
)
