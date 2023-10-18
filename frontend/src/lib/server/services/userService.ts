import type { User } from '../entities/entitites';

const userService = {
	create: async (name: string, email: string, bio: string, avatar: Blob) => {
		const res = await fetch('/api/user', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ name, email, bio, avatar })
		});
	},
	getById: async (id: string): Promise<User> => {
		const res = await fetch(`/api/user/${id}`);
		return (await res.json()) as User;
	}
};
// Create(ctx context.Context, user *entity.User) (string, error)
// DeleteUserByID(ctx context.Context, id string) error
// GetUserByID(ctx context.Context, id string) (*entity.User, error)
// GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
// ListAll(ctx context.Context, limit int) ([]*entity.User, error)
// UpdateUserByID(ctx context.Context, id string, user *entity.User) (*entity.User, error)
