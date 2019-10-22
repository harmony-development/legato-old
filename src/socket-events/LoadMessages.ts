import { Socket } from 'socket.io';
import { Events } from '../types';

export default function onLoadMessages(socket: Socket) {
  socket.on(Events.LOAD_MESSAGES, () => {});
}
