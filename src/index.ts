import http from 'http';
import express from 'express';

import socketio from 'socket.io';
import chalk from 'chalk';
import { IMessage, IUserData, IConnectData } from './types';
import { Server } from './Server';

export const harmonyServer = new Server(4000);

harmonyServer.open().then(() => {
  console.log(chalk.green('Successfully listening on port 4000'));
});
