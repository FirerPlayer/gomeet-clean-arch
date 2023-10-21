import type { User } from '$lib/server/entities/entitites';
import userService from '$lib/server/services/userService';
import type { LayoutServerLoad } from './$types';

export const load: LayoutServerLoad = async () => {
	// let newUser: User | null;
	// let error;
	// try {
	let newUser = (await userService.create('test', 'test', 'test', new Blob())) as User;
	// } catch (error) {
	// 	console.log(error);
	// 	newUser = null;
	// }
	// console.log(newUser);
	return { newUser };
};
