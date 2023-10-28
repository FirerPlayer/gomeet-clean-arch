export type Chat = {
	id: string;
	fromUser: string;
	toUsers: string[];
	createdAt: Date;
	updatedAt: Date;
};

export type Message = {
	chatId: string;
	content: string;
	file: Int8Array;
	created: Date;
};

export type User = {
	id: string;
	avatar: Int8Array;
	name: string;
	email: string;
	bio: string;
	createdAt: Date;
	updatedAt: Date;
};
