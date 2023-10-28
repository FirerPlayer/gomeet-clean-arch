import type { User } from '../entities/entitites';

export const BASE_URL = 'http://localhost:8080';

const userService = {
	create: async (
		name: string,
		email: string,
		bio: string,
		avatar: Int8Array
	): Promise<string | Error> => {
		const res = await fetch(BASE_URL + '/api/user', {
			method: 'POST',
			headers: {
				'Content-Type': 'application/json'
			},
			body: JSON.stringify({ name, email, bio, avatar })
		});
		if (!res.ok) {
			return new Error(JSON.stringify({ status: res.statusText, error: await res.json() }));
		}
		return (await res.json()).id as string;
	},
	getById: async (id: string): Promise<User | Error> => {
		const res = await fetch(`${BASE_URL}/api/user/details?id=${id}`);
		if (!res.ok) {
			return new Error(JSON.stringify({ status: res.statusText, error: await res.json() }));
		}
		return (await res.json()) as User;
	},
	getByEmail: async (email: string): Promise<User | Error> => {
		const res = await fetch(`${BASE_URL}/api/user/details?email=${email}`);
		if (!res.ok) {
			return new Error(JSON.stringify({ status: res.statusText, error: await res.json() }));
		}
		return (await res.json()) as User;
	},
	listAll: async (limit: number = 20): Promise<User[] | Error> => {
		const res = await fetch(`${BASE_URL}/api/user/all?limit=${limit}`);
		if (!res.ok) {
			return new Error(JSON.stringify({ status: res.statusText, error: await res.json() }));
		}
		return (await res.json()) as User[];
	},
	deleteById: async (id: string): Promise<string | Error> => {
		const res = await fetch(`${BASE_URL}/api/user?id=${id}`, {
			method: 'DELETE'
		});
		if (!res.ok) {
			return new Error(JSON.stringify({ status: res.statusText, error: await res.json() }));
		}
		return res.statusText;
	}
};

export default userService;
// Create(ctx context.Context, user *entity.User) (string, error)
// DeleteUserByID(ctx context.Context, id string) error
// GetUserByID(ctx context.Context, id string) (*entity.User, error)
// GetUserByEmail(ctx context.Context, email string) (*entity.User, error)
// ListAll(ctx context.Context, limit int) ([]*entity.User, error)
// UpdateUserByID(ctx context.Context, id string, user *entity.User) (*entity.User, error)
