package db

import "context"

func (b *UserDB) Create(ctx context.Context, user User) error {
	stmt, err := b.db.Prepare("INSERT INTO users (Name, LastName, Email, Mobile, Birthday) VALUES ($1, $2, $3, $4, $5)")
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.ExecContext(ctx, user.Name, user.LastName, user.Email, user.Mobile, user.Birthday)
	if err != nil {
		return err
	}

	return nil
}

func (b *UserDB) Delete(ctx context.Context, user User) error {
	_, err := b.db.ExecContext(ctx, "DELETE FROM users WHERE Email = $1", user.Email)
	if err != nil {
		return err
	}
	return nil
}
