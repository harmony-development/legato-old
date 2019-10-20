import mongoose from 'mongoose';
import { User } from './schema/userSchema';
import chalk from 'chalk';
import { IUser } from './types';

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
      email,
      userid: new mongoose.mongo.ObjectId()
    });
    await newUser.save();
  }
}
