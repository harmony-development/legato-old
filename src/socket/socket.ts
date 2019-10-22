import io from 'socket.io-client';
import { IProfileUpdate } from '../types';

export const Events = {
  PROFILE_UPDATE: 'PROFILE_UPDATE',
  GET_USER_DATA: 'GET_USER_DATA',
  GET_USER_DATA_ERROR: 'GET_USER_DATA_ERROR',
  MESSAGE: 'MESSAGE',
  LOGIN: 'LOGIN',
  LOGIN_ERROR: 'LOGIN_ERROR',
  REGISTER: 'REGISTER',
  REGISTER_ERROR: 'REGISTER_ERROR',
  INVALIDATE_SESSION: 'INVALIDATE_SESSION'
};

export class HarmonyConnection {
  connection: SocketIOClient.Socket;

  constructor() {
    this.connection = io('0.0.0.0:4000');

    this.connection.on(Events.INVALIDATE_SESSION, () => {
      console.log('Session Invalid');
    });
  }

  register(email: string, username: string, password: string): void {
    this.connection.emit(Events.REGISTER, { email, username, password });
  }

  login(email: string, password: string): void {
    this.connection.emit(Events.LOGIN, { email, password });
  }

  saveProfile(newUser: IProfileUpdate): void {
    this.connection.emit(Events.PROFILE_UPDATE, newUser);
  }

  getUserData(): void {
    this.connection.emit(Events.GET_USER_DATA, { token: localStorage.getItem('token') as string });
  }
}
