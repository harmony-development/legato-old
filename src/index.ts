import chalk from 'chalk';
import { Server } from './Server';
import { Config } from './Config';

const PORT = 4000;

export const harmonyServer = new Server(PORT);
export const config = new Config();

harmonyServer.open().then(() => {
  console.log(chalk.green(`Listening on port ${PORT}`));
});
