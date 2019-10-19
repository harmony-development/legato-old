import http from 'http';
import express from 'express';
import socketio from 'socket.io';
import onMessage from './socket-events/Message';
import { IUserData, IMessage, EventData } from './types';
import onDisconnect from './socket-events/Disconnect';
import onProfileUpdate from './socket-events/ProfileUpdate';
import onLogin from './socket-events/Login';

export class Server {
  app = express();
  HTTPServer: http.Server;
  SocketServer: SocketIO.Server;
  port: number;
  users: IUserData;

  constructor(port: number) {
    this.HTTPServer = http.createServer(this.app);
    this.SocketServer = socketio(this.HTTPServer);
    this.SocketServer.on('connection', socket => {
      onMessage(socket);
      onDisconnect(socket);
      onLogin(socket);
      onProfileUpdate(socket);
    });

    this.app.use(express.static('public'));

    this.port = port;
    this.HTTPServer.on('error', this.errorHandler);
    this.users = {};
  }

  private errorHandler(err: Error) {
    console.log(err.name);
  }

  updateName = (userID: string, name: string) => {
    if (this.users[userID]) {
      this.users[userID].name = name;
    }
  };

  emit(event: 'MESSAGE', data: IMessage): void;

  emit(event: string, data: EventData) {
    this.SocketServer.emit(event, data);
  }

  getUsers = () => {
    return this.users;
  };

  getSocketServer = () => {
    return this.SocketServer;
  };

  open = (): Promise<void> => {
    return new Promise((resolve, reject) => {
      this.HTTPServer.listen(this.port, '0.0.0.0', () => {
        resolve();
      });
    });
  };
}
