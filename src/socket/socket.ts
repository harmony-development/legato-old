import io from 'socket.io-client';

export class HarmonyConnection {
  connection: SocketIOClient.Socket;

  constructor() {
    this.connection = io('0.0.0.0:4000');
  }
}
