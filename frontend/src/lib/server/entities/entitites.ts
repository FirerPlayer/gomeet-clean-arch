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
	file: Blob;
	created: Date;
};

export type User = {
	id: string;
	avatar: Blob;
	name: string;
	email: string;
	bio: string;
	createdAt: Date;
	updatedAt: Date;
};
