import { Socket } from 'socket.io';
import { Events, IUserUpdateData, IToken } from '../types';
import { verify } from '../promisified/jwt';
import { config } from '..';
import { User } from '../schema/userSchema';

export default function onUpdateUser(socket: Socket) {
  socket.on(Events.PROFILE_UPDATE, (data: IUserUpdateData) => {
    if (typeof data.token === 'string') {
      verify(data.token, config.config.jwtsecret)
        .then(result => {
          if (result.valid) {
            if ((result.decoded as IToken).userid) {
              User.findOne({ userid: (result.decoded as IToken).userid })
                .then(user => {
                  if (user) {
                    if (data.avatar) {
                      user.avatar = data.avatar;
                    }
                    if (data.username) {
                      user.username = data.username;
                    }
                    user.save();
                  } else {
                    socket.emit(
                      Events.PROFILE_UPDATE_ERROR,
                      'You do not exist'
                    );
                  }
                })
                .catch(err => {
                  console.log(err);
                  socket.emit(
                    Events.PROFILE_UPDATE_ERROR,
                    'Failed to update profile'
                  );
                });
            }
          }
        })
        .catch(() => {
          socket.emit(Events.PROFILE_UPDATE_ERROR, 'Missing Token');
        });
    }
  });
}
