import type { Chat, Message } from '$lib/server/entities/entitites';
import { BASE_URL } from '$lib/server/services/userService';
function getChatService(address: string, onMessage: (msg: Message) => void) {
	const ws = new WebSocket(`${BASE_URL}${address}`);

	ws.onmessage = (ev) => {
		console.log(ev.data);
		onMessage(ev.data as Message);
	};

	const chatService = {
		create: (fromUser: string, toUsers: string[]) => {
			ws.onopen = (ev) => {
				ws.send(JSON.stringify({ fromUser, toUsers }));
			};
		},
		sendMessage: (chatId: string, content: string, file: Blob) => {
			ws.send(JSON.stringify({ chatId, content, file }));
		}
	};
	return chatService;
}
