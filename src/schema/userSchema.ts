import mongoose, { Schema } from 'mongoose';
import bcrypt from 'bcrypt';
import jwt from 'jsonwebtoken';
import randomstring from 'crypto-random-string';
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
  email: {
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
  bcrypt
    .hash(this.password, 10)
    .then(hash => {
      this.password = hash;
      this.userid = randomstring({ length: 15 });
      next();
    })
    .catch(err => {
      next(err);
    });
});

export const User = mongoose.model<IUser>('User', userSchema);
