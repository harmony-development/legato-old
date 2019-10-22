import mongoose from 'mongoose';
import { User } from './schema/userSchema';
import chalk from 'chalk';
import { IUser, IToken, IMessage } from './types';
import { verify } from './promisified/jwt';
import { config } from '.';
import { Message } from './schema/messageSchema';

export class HarmonyDB {
  private db: mongoose.Connection;

  constructor() {
    mongoose.connect('mongodb://localhost/harmony', {
      useNewUrlParser: true,
      useCreateIndex: true,
      useUnifiedTopology: true
    });
    this.db = mongoose.connection;
    this.db.on('error', console.error.bind(console, 'Connection Error : '));
    this.db.once('open', () => {
      console.log(chalk.green(chalk.bold('Successfully connected to MongoDB')));
    });
  }

  async register(email: string, password: string, username: string) {
    const newUser: IUser = new User({
      username,
      password,
      email
    });
    return await newUser.save();
  }

  async addMessage(userid: string, message: string, files: string[]) {
    const newMessage: IMessage = new Message({
      author: userid,
      message,
      files
    });
    return await newMessage.save();
  }

  verifyToken(token: string): Promise<string> {
    return new Promise((resolve, reject) => {
      if (!token || typeof token !== 'string') {
        reject();
        return;
      }
      verify(token, config.config.jwtsecret)
        .then(result => {
          if (result.valid && result.decoded) {
            if ((result.decoded as IToken).userid) {
              resolve((result.decoded as IToken).userid);
            } else {
              reject();
              return;
            }
          } else {
            reject();
            return;
          }
        })
        .catch(() => {
          reject();
          return;
        });
    });
  }
}
