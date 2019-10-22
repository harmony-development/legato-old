import { Socket } from 'socket.io';
import { Events, IGetUserData, IToken, IUser } from '../types';
import { verify } from '../promisified/jwt';
import { config, harmonyServer } from '..';
import { User } from '../schema/userSchema';

export default function onGetUserData(socket: Socket) {
  socket.on(Events.GET_USER_DATA, (data: IGetUserData) => {
    if (data.token) {
      verify(data.token, config.config.jwtsecret)
        .then(result => {
          if (result.valid && result.decoded) {
            if ((result.decoded as IToken).userid) {
              User.findOne({ userid: (result.decoded as IToken).userid })
                .then(user => {
                  if (user) {
                    harmonyServer.getSocketServer().emit(Events.GET_USER_DATA, {
                      username: user.username,
                      avatar: user.avatar,
                      theme: user.theme
                    });
                  } else {
                    socket.emit(
                      Events.PROFILE_UPDATE_ERROR,
                      'You do not exist'
                    );
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
        .catch(() => {});
    } else {
      socket.emit(Events.GET_USER_DATA_ERROR, 'Invalid token');
    }
  });
}
