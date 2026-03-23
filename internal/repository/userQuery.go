package repository

const (
	CreateUserQuery = `INSERT INTO account (
						   id, 
						   user_name, 
						   email, 
						   password, 
						   name, 
						   active
						   ) VALUES ($1, $2, $3, $4, $5, $6)`

	UpdateUserQuery = `UPDATE account SET
							user_name = $1,
							email = $2,
							password = $3,
							name = $4,
							active = $5
						WHERE id = $6`

	DeleteUserQuery = `DELETE FROM account WHERE id = $1`

	FindUserQuery = `SELECT
    					id,
    					user_name,
    					email,
    					name,
    					active
						FROM account WHERE id = $1`

	FindAllUserQuery = `SELECT
    						id,
    						user_name,
    						email,
    						name,
    						active
						FROM account`
)
