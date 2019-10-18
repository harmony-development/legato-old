import { harmonyServer } from '..';
import { Socket } from 'socket.io';
import { Events } from '../types';

export default function onDisconnect(socket: Socket) {
  socket.on(Events.DISCONNECT, () => {
    if (harmonyServer.getUsers()[socket.id]) {
      harmonyServer.emit('MESSAGE', {
        author: harmonyServer.getUsers()[socket.id].name,
        message: 'has left the channel'
      });
      delete harmonyServer.getUsers()[socket.id]; // free up RAM
    } else {
      harmonyServer.emit('MESSAGE', {
        author: 'Anonymous User',
        message: 'has left the channel'
      });
    }
  });
}
