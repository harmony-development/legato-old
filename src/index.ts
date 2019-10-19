import chalk from 'chalk';
import { Server } from './Server';

const PORT = 4000;

export const harmonyServer = new Server(PORT);

harmonyServer.open().then(() => {
  console.log(chalk.green(`Listening on port ${PORT}`));
});
