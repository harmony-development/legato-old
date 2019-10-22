import { Socket } from 'socket.io';
import { Events, IMessageData, IToken } from '../types';
import { verify } from '../promisified/jwt';
import { config, harmonyServer } from '..';
import { User } from '../schema/userSchema';

export default function onMessage(socket: Socket) {
  socket.on(Events.MESSAGE, (data: IMessageData) => {
    harmonyServer.Database.verifyToken(data.token)
      .then(userid => {
        User.findOne({ userid })
          .then(user => {
            if (user) {
              harmonyServer.getSocketServer().emit(Events.MESSAGE, {
                username: user.username,
                message: data.message,
                avatar: user.avatar,
                files: data.files
              });
            }
          })
          .catch(err => {
            console.log(err);
          });
      })
      .then(() => {
        harmonyServer
          .getSocketServer()
          .emit(Events.INVALIDATE_SESSION, 'Invalid token');
      });
  });
}
