import { Socket } from 'socket.io';
import { Events, IUsernameUpdate } from '../types';
import { harmonyServer } from '..';

export default function onUsernameUpdate(socket: Socket) {
  socket.on(Events.USERNAME_UPDATE, (data: IUsernameUpdate) => {
    if (data.name && typeof data.name == 'string') {
      harmonyServer.sendMessage(
        harmonyServer.getUsers()[socket.id].name || socket.id,
        `updated their username to ${data.name.substring(0, 50)}`
      );
      harmonyServer.getUsers()[socket.id].name = data.name.substring(0, 50);
    }
  });
}
