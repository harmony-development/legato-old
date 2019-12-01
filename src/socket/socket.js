import { EventEmitter } from 'fbemitter';

export class HarmonyConnection {
    connection;
    emitter = new EventEmitter();

    constructor() {
        this.connection = new WebSocket('ws://localhost:8080/api/socket/');
        this.connection.addEventListener('message', (event) => {
            const message = event.data;
            const parsed = JSON.parse(message);
            this.emitter.emit(parsed.type, parsed.data);
        });
        this.connection.addEventListener('close', () => {
            this.emitter.emit('close');
        });
        this.connection.addEventListener('error', () => {
            this.emitter.emit('error');
        });
        this.connection.addEventListener('open', () => {
            this.emitter.emit('open');
        });
    }
}
