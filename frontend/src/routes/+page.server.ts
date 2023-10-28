import type { User } from '$lib/server/entities/entitites';
import userService from '$lib/server/services/userService';
import type { Actions } from '@sveltejs/kit';

export const actions: Actions = {
	default: async ({ request, locals }) => {
		let newUserId = await userService.create('test', 'test', 'test', [] as unknown as Int8Array);
		if (newUserId instanceof Error) {
			console.log(newUserId);
			return;
		}
		let newUser = (await userService.getById(newUserId)) as User;
		locals.user = newUser;
		console.log(locals.user);
	}
};
