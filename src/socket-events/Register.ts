import { Socket } from 'socket.io';
import isemail from 'isemail';
import { Events, IRegisterDetails, IUser } from '../types';
import { harmonyServer, config } from '..';
import { sign } from '../promisified/jwt';
import { User } from '../schema/userSchema';

interface IValidationResult {
  valid: boolean;
  message?: string;
}

function validPassword(password: string): IValidationResult {
  if (password.length < 5) {
    return { valid: false, message: 'Password must be at least 5 characters.' };
  }
  if (password.length > 30) {
    return {
      valid: false,
      message: 'Password cannot be longer than 30 characters.'
    };
  }
  if (!/[a-z]/.test(password)) {
    return { valid: false, message: 'Password must contain lowercase letters' };
  }
  if (!/[0-9]/.test(password)) {
    return { valid: false, message: 'Password must contain numbers' };
  }
  return { valid: true };
}

export default function onRegister(socket: Socket) {
  socket.on(Events.REGISTER, (data: IRegisterDetails) => {
    if (
      typeof data.email === 'string' &&
      typeof data.password === 'string' &&
      typeof data.username === 'string'
    ) {
      if (isemail.validate(data.email, { errorLevel: true }) <= 5) {
        if (validPassword(data.password).valid) {
          User.findOne({ email: data.email }, (err, user: IUser) => {
            if (!user) {
              harmonyServer.Database.register(
                data.email,
                data.password,
                data.username
              )
                .then(() => {
                  sign(
                    {
                      email: data.email
                    },
                    config.config.jwtsecret,
                    { expiresIn: '7d' }
                  )
                    .then(token => {
                      socket.emit(Events.REGISTER, token);
                    })
                    .catch(err => {
                      console.log(err);
                      socket.emit(
                        Events.REGISTER_ERROR,
                        'Sorry, but the API is having a stroke right now'
                      );
                    });
                })
                .catch(err => {
                  console.log(err);
                  socket.emit(
                    Events.REGISTER_ERROR,
                    'Sorry, but the API is having a stroke right now'
                  );
                });
            } else {
              socket.emit(Events.REGISTER_ERROR, 'Email already registered');
            }
          }).catch(() => {
            socket.emit(
              Events.REGISTER_ERROR,
              'Sorry, but the API is having a stroke right now'
            );
          });
        } else {
          socket.emit(
            Events.REGISTER_ERROR,
            validPassword(data.password).message
          );
        }
      } else {
        socket.emit(Events.REGISTER_ERROR, 'Invalid Email.');
      }
    } else {
      socket.emit(Events.REGISTER_ERROR, 'Missing email or password.');
    }
  });
}
