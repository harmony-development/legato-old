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
