import chalk from 'chalk';
import { Server } from './Server';

export const harmonyServer = new Server(4000);

harmonyServer.open().then(() => {
  console.log(chalk.green('Successfully listening on port 4000'));
});
