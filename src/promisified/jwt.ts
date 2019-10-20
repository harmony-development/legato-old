import jwt from 'jsonwebtoken';

export function sign(
  payload: string | object | Buffer,
  secretOrPrivateKey: jwt.Secret,
  options?: jwt.SignOptions
): Promise<string> {
  return new Promise<string>((resolve, reject) => {
    if (options) {
      jwt.sign(payload, secretOrPrivateKey, options, (err, token) => {
        if (err) {
          reject(err);
          return;
        }
        resolve(token);
      });
    } else {
      jwt.sign(payload, secretOrPrivateKey, (err, token) => {
        if (err) {
          reject(err);
          return;
        }
        resolve(token);
      });
    }
  });
}

interface IVerify {
  valid: boolean;
  decoded?: string | object;
}

export function verify(
  token: string,
  secretOrPublicKey: string,
  options?: jwt.VerifyOptions
): Promise<IVerify> {
  return new Promise<IVerify>((resolve, reject) => {
    jwt.verify(token, secretOrPublicKey, options, (err, decoded) => {
      if (decoded && !err) {
        resolve({ valid: true, decoded });
      } else {
        resolve({ valid: false });
      }
    });
  });
}
