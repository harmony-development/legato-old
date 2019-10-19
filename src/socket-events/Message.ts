import _ from 'lodash';
import { Socket } from 'socket.io';
import { harmonyServer } from '..';
import { IMessage, Events } from '../types';

export default function onMessage(socket: Socket) {
  socket.on(Events.MESSAGE, (data: IMessage) => {
    if (data.message && typeof data.message == 'string') {
      harmonyServer.emit('MESSAGE', {
        ...data,
        icon: _.get(harmonyServer.getUsers()[socket.id], 'icon', '')
      });
    }
  });
}
