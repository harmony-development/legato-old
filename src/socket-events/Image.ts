import { Events, IImageData } from '../types';
import { harmonyServer } from '..';
import { Socket } from 'socket.io';

export default function onImage(socket: Socket) {
  socket.on(Events.IMAGE, (data: IImageData) => {
    if (harmonyServer.getUsers()[socket.id]) {
      harmonyServer.getSocketServer().emit(Events.IMAGE, {
        author: harmonyServer.getUsers()[socket.id].name,
        data
      });
    } else {
      harmonyServer.getSocketServer().emit(Events.IMAGE, {
        author: harmonyServer.getUsers()[socket.id].name,
        data
      });
    }
  });
}
