import { Socket } from 'socket.io';
import { Events, IMessageData, IToken } from '../types';
import { verify } from '../promisified/jwt';
import { config, harmonyServer } from '..';
import { User } from '../schema/userSchema';

export default function onMessage(socket: Socket) {
  socket.on(Events.MESSAGE, (data: IMessageData) => {
    verify(data.token, config.config.jwtsecret)
      .then(result => {
        if (result.valid && result.decoded) {
          if ((result.decoded as IToken).userid) {
            User.findOne({ userid: (result.decoded as IToken).userid })
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
          } else {
            socket.emit(Events.INVALIDATE_SESSION, 'invalid session token');
          }
        } else {
          socket.emit(Events.INVALIDATE_SESSION, 'invalid session token');
        }
      })
      .catch(() => {
        socket.emit(Events.INVALIDATE_SESSION, 'invalid session token');
      });
  });
}
