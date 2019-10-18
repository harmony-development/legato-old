import http from 'http';
import express from 'express';
import socketio from 'socket.io';
import onMessage from './socket-events/Message';
import { IUserData, Events, IMessage, ILoginData, EventData } from './types';
import onDisconnect from './socket-events/Disconnect';
import onUsernameUpdate from './socket-events/UsernameUpdate';
import onLogin from './socket-events/Login';

export class Server {
  private app: Express.Application;
  private HTTPServer: http.Server;
  private SocketServer: SocketIO.Server;
  private port: number;
  private users: IUserData;

  constructor(port: number) {
    this.app = express();
    this.HTTPServer = http.createServer(this.app);
    this.SocketServer = socketio(this.HTTPServer);
    this.SocketServer.on('connection', socket => {
      onMessage(socket);
      onDisconnect(socket);
      onLogin(socket);
      onUsernameUpdate(socket);
    });
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
      this.HTTPServer.listen(this.port, () => {
        resolve();
      });
    });
  };
}
