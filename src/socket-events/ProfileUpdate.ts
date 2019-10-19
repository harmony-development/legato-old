import _ from 'lodash';
import { Socket } from 'socket.io';
import { Events, IProfileUpdate } from '../types';
import { harmonyServer } from '..';

export default function onProfileUpdate(socket: Socket) {
  socket.on(Events.PROFILE_UPDATE, (data: IProfileUpdate) => {
    if (data.name && typeof data.name === 'string') {
      harmonyServer.getUsers()[socket.id].name = data.name.substring(0, 50);
    }
    if (data.icon && typeof data.icon === 'string') {
      harmonyServer.getUsers()[socket.id].icon = data.icon;
    }
  });
}
