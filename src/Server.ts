import http from 'http';
import express from 'express';
import socketio from 'socket.io';

import registerRoute from './routes/register';
import loginRoute from './routes/login';

// import onMessage from './socket-events/Message';
// import onDisconnect from './socket-events/Disconnect';
// import onProfileUpdate from './socket-events/ProfileUpdate';
// import onLogin from './socket-events/Login';
import { HarmonyDB } from './HarmonyDB';

export class Server {
  app = express();
  HTTPServer: http.Server;
  SocketServer: SocketIO.Server;
  Database: HarmonyDB;
  port: number;

  constructor(port: number) {
    this.port = port;

    this.HTTPServer = http.createServer(this.app);
    this.HTTPServer.on('error', this.errorHandler);

    this.SocketServer = socketio(this.HTTPServer);
    this.SocketServer.on('connection', socket => {});

    this.Database = new HarmonyDB();

    this.app.use(express.static('public'));
    this.app.use('/api', registerRoute);
    this.app.use('/api', loginRoute);
  }

  private errorHandler(err: Error) {
    console.log(err.name);
  }

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
