import { Socket } from 'socket.io';
import { promisify } from 'util';
import { Events, ILoginDetails } from '../types';
import bcrypt from 'bcrypt';
import jwt from 'jsonwebtoken';
import { User } from '../schema/userSchema';
import { config } from '..';
import { sign } from '../promisified/jwt';

export default function onLogin(socket: Socket) {
  socket.on(Events.LOGIN, (data: ILoginDetails) => {
    if (data.email && data.password) {
      User.findOne({ email: data.email })
        .then(user => {
          if (user) {
            if (user.password) {
              bcrypt
                .compare(data.password, user.password)
                .then(success => {
                  if (success) {
                    sign(
                      {
                        userid: user.userid
                      },
                      config.config.jwtsecret,
                      { expiresIn: '7d' }
                    )
                      .then(token => {
                        socket.emit(Events.LOGIN, { token, theme: user.theme });
                      })
                      .catch(() => {
                        socket.emit(
                          Events.LOGIN_ERROR,
                          'Uhm. The API is having a stroke.'
                        );
                      });
                  } else {
                    socket.emit(
                      Events.LOGIN_ERROR,
                      'Invalid email or password'
                    );
                  }
                })
                .catch(() => {
                  socket.emit(
                    Events.LOGIN_ERROR,
                    'Uhm. The API is having a stroke.'
                  );
                });
            } else {
              socket.emit(
                Events.LOGIN_ERROR,
                'Somehow the password is missing in our records. Email support please!'
              );
            }
          } else {
            socket.emit(Events.LOGIN_ERROR, 'Invalid email or password');
          }
        })
        .catch(() => {
          socket.emit(Events.LOGIN_ERROR, 'Invalid email or password');
        });
    } else {
      socket.emit(Events.LOGIN_ERROR, 'Missing username or password');
    }
  });
}
