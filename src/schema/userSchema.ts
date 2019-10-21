import mongoose, { Schema } from 'mongoose';
import bcrypt from 'bcrypt';
import randomstring from 'crypto-random-string';
import { IUser, ITheme } from '../types';

export const userSchema: Schema = new mongoose.Schema({
  userid: {
    unique: true,
    required: false,
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
  },
  avatar: {
    unique: false,
    required: false,
    type: String
  },
  theme: {
    unique: false,
    required: false,
    type: {
      primary: {
        light: String,
        dark: String,
        contrastText: String
      },
      secondary: {
        light: String,
        dark: String,
        contrastText: String
      },
      type: String
    }
  }
});

userSchema.pre<IUser>('save', function(next) {
  if (!this.isModified('password')) return next();

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
