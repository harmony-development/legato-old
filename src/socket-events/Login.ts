import { Socket } from 'socket.io';
import { Events, ILoginData } from '../types';
import { harmonyServer } from '..';

export default function onLogin(socket: Socket) {
  socket.on(Events.LOGIN, (data: ILoginData) => {
    if (data.name) {
      if (harmonyServer.getUsers()[socket.id]) {
        harmonyServer.emit('MESSAGE', {
          author: harmonyServer.getUsers()[socket.id].name,
          message: `has joined the channel`
        });
        harmonyServer.getUsers()[socket.id].name = data.name;
      } else harmonyServer.getUsers()[socket.id] = { name: data.name };
    }
  });
}
