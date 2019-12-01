import { EventEmitter } from 'fbemitter';
import { IPacket } from '../types/socket';

export default class HarmonySocket {
    conn: WebSocket;
    events: EventEmitter;

    constructor() {
        this.conn = new WebSocket('ws://localhost:8080/api/socket');
        this.events = new EventEmitter();
        this.conn.addEventListener('open', () => this.events.emit('open'));
        this.conn.addEventListener('close', () => this.events.emit('close'));
        this.conn.addEventListener('error', () => this.events.emit('error'));
        this.conn.onmessage = (e: MessageEvent) => {
            const unprocessed = JSON.parse(e.data);
            if (typeof unprocessed['type'] === 'string' && typeof unprocessed['data'] === 'string') {
                const packet: IPacket = unprocessed;
                this.events.emit(packet.type, packet.data);
            } else {
                console.warn(`Unsupported packet received : ${unprocessed}`);
            }
        };
    }

    emitEvent(type: string, data: unknown) {
        this.conn.send(JSON.stringify({ type, data }));
    }

    login(email: string, password: string) {
        this.emitEvent('login', {
            email,
            password
        });
    }
}
