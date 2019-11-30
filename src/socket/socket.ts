interface ISocketEvent {
    type: string;
    data: unknown;
}

interface IDeauth {
    message: string;
}

interface IMessage {
    guild: string;
    userid: string;
    message: string;
}

export class HarmonyConnection {
    connection: WebSocket;

    constructor() {
        this.connection = new WebSocket('ws://localhost:8080/api/socket/');
        this.connection.addEventListener('message', function(event) {
            const message = event.data;
            const parsed: ISocketEvent = JSON.parse(message);
            switch (parsed.type) {
                case 'Deauth': {
                    localStorage.removeItem('token');
                    break;
                }
                case 'Message': {
                    const parsedMessage = parsed.data as IMessage;
                    console.log(`Message in ${parsedMessage.guild} from ${parsedMessage.userid} with message ${parsedMessage.message}`);
                    break;
                }
                default: {
                    console.log(`Unknown event received : ${parsed.type}`);
                }
            }
        });
    }
}
