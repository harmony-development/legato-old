import { harmonyServer } from '..';
import { Socket } from 'socket.io';
import { Events } from '../types';

export default function onDisconnect(socket: Socket) {
  socket.on(Events.DISCONNECT, () => {
    if (harmonyServer.getUsers()[socket.id]) {
      harmonyServer.sendMessage(
        harmonyServer.getUsers()[socket.id].name,
        'has left the channel'
      );
      delete harmonyServer.getUsers()[socket.id]; // free up RAM
    } else {
      harmonyServer.sendMessage('Anonymous User', 'has left the channel');
    }
  });
}
