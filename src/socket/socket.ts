import { EventEmitter } from 'fbemitter';

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
    emitter = new EventEmitter();

    constructor() {
        this.connection = new WebSocket('ws://localhost:8080/api/socket/');
        this.connection.addEventListener('message', (event) => {
            const message = event.data;
            const parsed: ISocketEvent = JSON.parse(message);
        });
    }
}
