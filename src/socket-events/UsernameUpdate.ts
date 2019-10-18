import _ from 'lodash';
import { Socket } from 'socket.io';
import { Events, IUsernameUpdate } from '../types';
import { harmonyServer } from '..';

export default function onUsernameUpdate(socket: Socket) {
  socket.on(Events.USERNAME_UPDATE, (data: IUsernameUpdate) => {
    if (data.name && typeof data.name == 'string') {
      harmonyServer.emit('MESSAGE', {
        author: _.get(harmonyServer.getUsers()[socket.id], 'name', socket.id),
        message: `has changed their name to ${data.name.substring(0, 50)}`
      });
      harmonyServer.getUsers()[socket.id].name = data.name.substring(0, 50);
    }
  });
}
