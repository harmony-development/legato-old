import { Socket } from 'socket.io';
import { harmonyServer } from '..';
import { IMessage, Events } from '../types';

export default function onMessage(socket: Socket) {
  socket.on(Events.MESSAGE, (data: IMessage) => {
    if (data.message && typeof data.message == 'string') {
      if (harmonyServer.getUsers()[socket.id]) {
        harmonyServer.sendMessage(
          harmonyServer.getUsers()[socket.id].name || socket.id,
          data.message.substring(0, 500)
        );
      } else {
        harmonyServer.sendMessage(socket.id, data.message.substring(0, 500));
      }
    }
  });
}
