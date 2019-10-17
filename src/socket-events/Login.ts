import { Socket } from 'socket.io';
import { Events, IConnectData } from '../types';
import { harmonyServer } from '..';

export default function onLogin(socket: Socket) {
  socket.on(Events.LOGIN, (data: IConnectData) => {
    if (data.name) {
      if (harmonyServer.getUsers()[socket.id])
        harmonyServer.getUsers()[socket.id].name = data.name;
      else harmonyServer.getUsers()[socket.id] = { name: data.name };
    }
  });
}
