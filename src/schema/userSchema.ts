import mongoose, { Schema } from 'mongoose';
import bcrypt from 'bcrypt';
import { IUser } from '../types';

export const userSchema: Schema = new mongoose.Schema({
  userid: {
    unique: true,
    required: true,
    type: String
  },
  username: {
    unique: false,
    required: true,
    type: String
  },
  password: {
    unique: false,
    required: true,
    type: String
  }
});

userSchema.pre<IUser>('save', function(next) {
  bcrypt.hash(this.password, 10, (err, hash) => {
    if (err) {
      return next(err);
    }

    this.password = hash;

    next();
  });
});

export const User = mongoose.model<IUser>('User', userSchema);
