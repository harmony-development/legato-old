import { Socket } from 'socket.io';
import { harmonyServer } from '..';
import { IMessage, Events } from '../types';

export default function onMessage(socket: Socket) {
  socket.on(Events.MESSAGE, (data: IMessage) => {
    if (data.message && typeof data.message == 'string') {
      harmonyServer.getSocketServer().emit(data.message.substring(0, 500));
    }
  });
}
