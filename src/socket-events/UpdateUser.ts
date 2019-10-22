import { Socket } from 'socket.io';
import { Events, IUserUpdateData, IToken } from '../types';
import { verify } from '../promisified/jwt';
import { config, harmonyServer } from '..';
import { User } from '../schema/userSchema';

export default function onUpdateUser(socket: Socket) {
  socket.on(Events.PROFILE_UPDATE, (data: IUserUpdateData) => {
    harmonyServer.Database.verifyToken(data.token)
      .then(userid => {
        User.findOne({ userid })
          .then(user => {
            if (user) {
              if (data.avatar) {
                user.avatar = data.avatar;
              }
              if (data.username) {
                user.username = data.username;
              }
              if (data.theme) {
                user.theme = data.theme;
              }
              user.save();
            } else {
              socket.emit(Events.PROFILE_UPDATE_ERROR, 'You do not exist');
            }
          })
          .catch(err => {
            console.log(err);
            socket.emit(
              Events.PROFILE_UPDATE_ERROR,
              'Failed to update profile'
            );
          });
      })
      .catch(() => {
        harmonyServer.SocketServer.emit(
          Events.INVALIDATE_SESSION,
          'Invalid token'
        );
      });
  });
}
