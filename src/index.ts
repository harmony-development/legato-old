import http from 'http';
import express from 'express';

import socketio from 'socket.io';
import chalk from 'chalk';
import { IMessage, IUserData, IConnectData } from './types';
import { Server } from './Server';

export const harmonyServer = new Server(4000);

harmonyServer.open();

// const PORT = 4000;

// // export these two just in case we need it sometime
// export const app = express();
// export const httpServer = http.createServer(app);
// export const socketServer = socketio(httpServer);

// const clientData: IClientData = {};

// // dummy response
// app.get('/', (req, res) => {
//   res.json({ bacon: true });
// });

// socketServer.on('connection', socket => {
//   socket.on('disconnect', () => {
//     socketServer.emit('ClientDisconnectEvent', {
//       userid: clientData[socket.id] || socket.id
//     });
//   });

//   socket.on('ClientConnect', (data: IConnectData) => {
//     if (data.name) {
//       clientData[socket.id] = {
//         name: data.name.substr(0, 30)
//       };
//       socketServer.emit('ClientConnectEvent', {
//         userid: clientData[socket.id].name
//       });
//     }
//   });

//   socket.on('message', (message: IMessage) => {
//     if (message.message && clientData[socket.id]) {
//       socketServer.emit('message', {
//         user: clientData[socket.id].name,
//         message: message.message
//       });
//     }
//   });

//   socket.on('UsernameUpdate', (name: string) => {
//     if (clientData[socket.id]) {
//       const oldName = clientData[socket.id].name;
//       clientData[socket.id].name = name;
//       socketServer.emit('message', {
//         user: oldName,
//         message: ` has changed their name to ${clientData[socket.id].name}`
//       });
//     }
//   });
// });

// httpServer.listen(PORT, () => {
//   console.log(chalk.green(`Server listening on port ${PORT}`));
// });
